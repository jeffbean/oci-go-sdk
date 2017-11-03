package common

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	DefaultHostUrlTemplate   = "%s.%s.oraclecloud.com"
	defaultScheme            = "https"
	defaultSDKMarker         = "Oracle-GoSDK"
	defaultUserAgentTemplate = "%s/%s (%s/%s; go/%s)" //SDK/SDKVersion (OS/OSVersion; Lang/LangVersion)
	defaultTimeout           = time.Second * 30
)

type RequestInterceptor func(*http.Request) error

// HttpRequestor wraps the execution of a http request, it is generally implemented by
// http.Client.Do, but can be customized for testing
type HttpRequestDispatcher interface {
	Do(req *http.Request) (*http.Response, error)
}

//BaseClient struct implements all basic operations to call oci web services.
type BaseClient struct {
	//HttpClient performs the http network operations
	HttpClient HttpRequestDispatcher

	//Signer performs auth operation
	Signer HttpRequestSigner

	//A request interceptor can be used to customize the request before signing and dispatching
	Interceptor RequestInterceptor

	//The region of the client
	Region Region

	//The host of the service
	Host string

	//The user agent
	UserAgent string

	//Base path for all operations of this client
	BasePath string
}

func defaultUserAgent() string {
	userAgent := fmt.Sprintf(defaultUserAgentTemplate, defaultSDKMarker, Version(), runtime.GOOS, runtime.GOARCH, runtime.Version())
	return userAgent
}

func newClientWithHttpDispatcher(dispatcher HttpRequestDispatcher, region Region, provider ConfigurationProvider) (client BaseClient) {
	signer := NewOCIRequestSigner(provider)
	client = BaseClient{
		UserAgent:   defaultUserAgent(),
		Region:      region,
		Interceptor: nil,
		Signer:      signer,
		HttpClient:  dispatcher,
	}
	return
}

func getDefaultHttpDispatcher() http.Client {
	httpClient := http.Client{
		Timeout:   defaultTimeout,
		Transport: &http.Transport{},
	}
	return httpClient
}

//Create a new client with a configuration provider, the configuration provider
//will be used for the default signer as well as reading the region
func NewClientWithConfig(configProvider ConfigurationProvider) (client BaseClient, err error) {
	regionstr, err := configProvider.Region()
	if err != nil {
		err = fmt.Errorf("can not create client, bad region configuration: %s", err.Error())
		return
	}
	region, err := StringToRegion(regionstr)
	if err != nil {
		err = fmt.Errorf("can not create client, bad region configuration: %s", err.Error())
		return
	}

	dispatcher := getDefaultHttpDispatcher()
	client = newClientWithHttpDispatcher(&dispatcher, region, configProvider)
	return
}

//Create a new default client for a given region
func NewClientForRegion(region Region) (client BaseClient) {
	dispatcher := getDefaultHttpDispatcher()
	provider := environmentConfigurationProvider{EnvironmentVariablePrefix: "TF_VAR"}
	return newClientWithHttpDispatcher(&dispatcher, region, provider)
}

func (client *BaseClient) prepareRequest(request *http.Request) (err error) {
	if client.UserAgent == "" {
		return fmt.Errorf("user agent can not be blank")
	}

	if request.Header == nil {
		request.Header = http.Header{}
	}
	request.Header.Set("User-Agent", client.UserAgent)

	if !strings.Contains(client.Host, "http") &&
		!strings.Contains(client.Host, "https") {
		client.Host = fmt.Sprintf("%s://%s", defaultScheme, client.Host)
	}

	clientUrl, err := url.Parse(client.Host)
	if err != nil {
		return fmt.Errorf("host is invalid. %s", err.Error())
	}
	request.URL.Host = clientUrl.Host
	request.URL.Scheme = clientUrl.Scheme
	currentPath := request.URL.Path
	request.URL.Path = path.Clean(fmt.Sprintf("/%s/%s", client.BasePath, currentPath))
	return
}

func (client BaseClient) intercept(request *http.Request) (err error) {
	if client.Interceptor != nil {
		err = client.Interceptor(request)
	}
	return
}

func checkForSuccessfulResponse(res *http.Response) error {
	familyStatusCode := res.StatusCode / 100
	if familyStatusCode == 4 || familyStatusCode == 5 {
		return newServiceFailureFromResponse(res)
	}
	return nil

}

func (client BaseClient) Call(ctx context.Context, request *http.Request) (response *http.Response, err error) {
	Debugln("Atempting to call downstream service")
	request = request.WithContext(ctx)

	err = client.prepareRequest(request)
	if err != nil {
		return
	}

	//Intercept
	err = client.intercept(request)
	if err != nil {
		return
	}

	//Sign the request
	err = client.Signer.Sign(request)
	if err != nil {
		return
	}

	IfDebug(func() {
		if dump, e := httputil.DumpRequest(request, true); e == nil {
			Logf("Dump Request %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	//Execute the http request
	response, err = client.HttpClient.Do(request)

	IfDebug(func() {
		if err != nil {
			Logln(err)
			return
		}

		if dump, e := httputil.DumpResponse(response, true); e == nil {
			Logf("Dump Response %v", string(dump))
		} else {
			Debugln(e)
		}
	})

	if err != nil {
		return
	}

	err = checkForSuccessfulResponse(response)
	return
}
