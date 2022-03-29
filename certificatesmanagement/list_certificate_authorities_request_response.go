// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificatesmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"net/http"
	"strings"
)

// ListCertificateAuthoritiesRequest wrapper for the ListCertificateAuthorities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListCertificateAuthorities.go.html to see an example of how to use ListCertificateAuthoritiesRequest.
type ListCertificateAuthoritiesRequest struct {

	// Unique Oracle-assigned identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the given compartment OCID.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter that returns only resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState ListCertificateAuthoritiesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns only resources that match the specified name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The OCID of the certificate authority (CA). If the parameter is set to null, the service lists all CAs.
	IssuerCertificateAuthorityId *string `mandatory:"false" contributesTo:"query" name:"issuerCertificateAuthorityId"`

	// The OCID of the certificate authority (CA). If the parameter is set to null, the service lists all CAs.
	CertificateAuthorityId *string `mandatory:"false" contributesTo:"query" name:"certificateAuthorityId"`

	// The field to sort by. You can specify only one sort order. The default
	// order for `EXPIRATIONDATE` and 'TIMECREATED' is descending. The default order for `NAME`
	// is ascending.
	SortBy ListCertificateAuthoritiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCertificateAuthoritiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificateAuthoritiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificateAuthoritiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificateAuthoritiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificateAuthoritiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCertificateAuthoritiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCertificateAuthoritiesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCertificateAuthoritiesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificateAuthoritiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCertificateAuthoritiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificateAuthoritiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCertificateAuthoritiesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCertificateAuthoritiesResponse wrapper for the ListCertificateAuthorities operation
type ListCertificateAuthoritiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CertificateAuthorityCollection instances
	CertificateAuthorityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCertificateAuthoritiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificateAuthoritiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificateAuthoritiesLifecycleStateEnum Enum with underlying type: string
type ListCertificateAuthoritiesLifecycleStateEnum string

// Set of constants representing the allowable values for ListCertificateAuthoritiesLifecycleStateEnum
const (
	ListCertificateAuthoritiesLifecycleStateCreating           ListCertificateAuthoritiesLifecycleStateEnum = "CREATING"
	ListCertificateAuthoritiesLifecycleStateActive             ListCertificateAuthoritiesLifecycleStateEnum = "ACTIVE"
	ListCertificateAuthoritiesLifecycleStateUpdating           ListCertificateAuthoritiesLifecycleStateEnum = "UPDATING"
	ListCertificateAuthoritiesLifecycleStateDeleting           ListCertificateAuthoritiesLifecycleStateEnum = "DELETING"
	ListCertificateAuthoritiesLifecycleStateDeleted            ListCertificateAuthoritiesLifecycleStateEnum = "DELETED"
	ListCertificateAuthoritiesLifecycleStateSchedulingDeletion ListCertificateAuthoritiesLifecycleStateEnum = "SCHEDULING_DELETION"
	ListCertificateAuthoritiesLifecycleStatePendingDeletion    ListCertificateAuthoritiesLifecycleStateEnum = "PENDING_DELETION"
	ListCertificateAuthoritiesLifecycleStateCancellingDeletion ListCertificateAuthoritiesLifecycleStateEnum = "CANCELLING_DELETION"
	ListCertificateAuthoritiesLifecycleStateFailed             ListCertificateAuthoritiesLifecycleStateEnum = "FAILED"
)

var mappingListCertificateAuthoritiesLifecycleStateEnum = map[string]ListCertificateAuthoritiesLifecycleStateEnum{
	"CREATING":            ListCertificateAuthoritiesLifecycleStateCreating,
	"ACTIVE":              ListCertificateAuthoritiesLifecycleStateActive,
	"UPDATING":            ListCertificateAuthoritiesLifecycleStateUpdating,
	"DELETING":            ListCertificateAuthoritiesLifecycleStateDeleting,
	"DELETED":             ListCertificateAuthoritiesLifecycleStateDeleted,
	"SCHEDULING_DELETION": ListCertificateAuthoritiesLifecycleStateSchedulingDeletion,
	"PENDING_DELETION":    ListCertificateAuthoritiesLifecycleStatePendingDeletion,
	"CANCELLING_DELETION": ListCertificateAuthoritiesLifecycleStateCancellingDeletion,
	"FAILED":              ListCertificateAuthoritiesLifecycleStateFailed,
}

var mappingListCertificateAuthoritiesLifecycleStateEnumLowerCase = map[string]ListCertificateAuthoritiesLifecycleStateEnum{
	"creating":            ListCertificateAuthoritiesLifecycleStateCreating,
	"active":              ListCertificateAuthoritiesLifecycleStateActive,
	"updating":            ListCertificateAuthoritiesLifecycleStateUpdating,
	"deleting":            ListCertificateAuthoritiesLifecycleStateDeleting,
	"deleted":             ListCertificateAuthoritiesLifecycleStateDeleted,
	"scheduling_deletion": ListCertificateAuthoritiesLifecycleStateSchedulingDeletion,
	"pending_deletion":    ListCertificateAuthoritiesLifecycleStatePendingDeletion,
	"cancelling_deletion": ListCertificateAuthoritiesLifecycleStateCancellingDeletion,
	"failed":              ListCertificateAuthoritiesLifecycleStateFailed,
}

// GetListCertificateAuthoritiesLifecycleStateEnumValues Enumerates the set of values for ListCertificateAuthoritiesLifecycleStateEnum
func GetListCertificateAuthoritiesLifecycleStateEnumValues() []ListCertificateAuthoritiesLifecycleStateEnum {
	values := make([]ListCertificateAuthoritiesLifecycleStateEnum, 0)
	for _, v := range mappingListCertificateAuthoritiesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateAuthoritiesLifecycleStateEnumStringValues Enumerates the set of values in String for ListCertificateAuthoritiesLifecycleStateEnum
func GetListCertificateAuthoritiesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"SCHEDULING_DELETION",
		"PENDING_DELETION",
		"CANCELLING_DELETION",
		"FAILED",
	}
}

// GetMappingListCertificateAuthoritiesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateAuthoritiesLifecycleStateEnum(val string) (ListCertificateAuthoritiesLifecycleStateEnum, bool) {
	enum, ok := mappingListCertificateAuthoritiesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificateAuthoritiesSortByEnum Enum with underlying type: string
type ListCertificateAuthoritiesSortByEnum string

// Set of constants representing the allowable values for ListCertificateAuthoritiesSortByEnum
const (
	ListCertificateAuthoritiesSortByName           ListCertificateAuthoritiesSortByEnum = "NAME"
	ListCertificateAuthoritiesSortByExpirationdate ListCertificateAuthoritiesSortByEnum = "EXPIRATIONDATE"
	ListCertificateAuthoritiesSortByTimecreated    ListCertificateAuthoritiesSortByEnum = "TIMECREATED"
)

var mappingListCertificateAuthoritiesSortByEnum = map[string]ListCertificateAuthoritiesSortByEnum{
	"NAME":           ListCertificateAuthoritiesSortByName,
	"EXPIRATIONDATE": ListCertificateAuthoritiesSortByExpirationdate,
	"TIMECREATED":    ListCertificateAuthoritiesSortByTimecreated,
}

var mappingListCertificateAuthoritiesSortByEnumLowerCase = map[string]ListCertificateAuthoritiesSortByEnum{
	"name":           ListCertificateAuthoritiesSortByName,
	"expirationdate": ListCertificateAuthoritiesSortByExpirationdate,
	"timecreated":    ListCertificateAuthoritiesSortByTimecreated,
}

// GetListCertificateAuthoritiesSortByEnumValues Enumerates the set of values for ListCertificateAuthoritiesSortByEnum
func GetListCertificateAuthoritiesSortByEnumValues() []ListCertificateAuthoritiesSortByEnum {
	values := make([]ListCertificateAuthoritiesSortByEnum, 0)
	for _, v := range mappingListCertificateAuthoritiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateAuthoritiesSortByEnumStringValues Enumerates the set of values in String for ListCertificateAuthoritiesSortByEnum
func GetListCertificateAuthoritiesSortByEnumStringValues() []string {
	return []string{
		"NAME",
		"EXPIRATIONDATE",
		"TIMECREATED",
	}
}

// GetMappingListCertificateAuthoritiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateAuthoritiesSortByEnum(val string) (ListCertificateAuthoritiesSortByEnum, bool) {
	enum, ok := mappingListCertificateAuthoritiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificateAuthoritiesSortOrderEnum Enum with underlying type: string
type ListCertificateAuthoritiesSortOrderEnum string

// Set of constants representing the allowable values for ListCertificateAuthoritiesSortOrderEnum
const (
	ListCertificateAuthoritiesSortOrderAsc  ListCertificateAuthoritiesSortOrderEnum = "ASC"
	ListCertificateAuthoritiesSortOrderDesc ListCertificateAuthoritiesSortOrderEnum = "DESC"
)

var mappingListCertificateAuthoritiesSortOrderEnum = map[string]ListCertificateAuthoritiesSortOrderEnum{
	"ASC":  ListCertificateAuthoritiesSortOrderAsc,
	"DESC": ListCertificateAuthoritiesSortOrderDesc,
}

var mappingListCertificateAuthoritiesSortOrderEnumLowerCase = map[string]ListCertificateAuthoritiesSortOrderEnum{
	"asc":  ListCertificateAuthoritiesSortOrderAsc,
	"desc": ListCertificateAuthoritiesSortOrderDesc,
}

// GetListCertificateAuthoritiesSortOrderEnumValues Enumerates the set of values for ListCertificateAuthoritiesSortOrderEnum
func GetListCertificateAuthoritiesSortOrderEnumValues() []ListCertificateAuthoritiesSortOrderEnum {
	values := make([]ListCertificateAuthoritiesSortOrderEnum, 0)
	for _, v := range mappingListCertificateAuthoritiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificateAuthoritiesSortOrderEnumStringValues Enumerates the set of values in String for ListCertificateAuthoritiesSortOrderEnum
func GetListCertificateAuthoritiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCertificateAuthoritiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificateAuthoritiesSortOrderEnum(val string) (ListCertificateAuthoritiesSortOrderEnum, bool) {
	enum, ok := mappingListCertificateAuthoritiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
