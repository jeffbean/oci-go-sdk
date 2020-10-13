// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"io"
	"net/http"
)

// GetRunLogRequest wrapper for the GetRunLog operation
type GetRunLogRequest struct {

	// The unique ID for the run
	RunId *string `mandatory:"true" contributesTo:"path" name:"runId"`

	// The name of the log. Avoid entering confidential information.
	Name *string `mandatory:"true" contributesTo:"path" name:"name"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRunLogRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRunLogRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRunLogRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetRunLogResponse wrapper for the GetRunLog operation
type GetRunLogResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control.
	// See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Indicates the size of the data as described in RFC 2616, section 14.13.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// Specifies the media type of the underlying data as described in RFC 2616, section 14.17.
	ContentType *string `presentIn:"header" name:"content-type"`

	// Indicates the encoding of the data, as described in RFC 2616, section 14.11.
	ContentEncoding *string `presentIn:"header" name:"content-encoding"`

	// The user-defined metadata for the log.
	OpcMeta map[string]string `presentIn:"header-collection" prefix:"opc-meta-"`
}

func (response GetRunLogResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRunLogResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
