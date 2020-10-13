// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package vault

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"net/http"
)

// ListSecretsRequest wrapper for the ListSecrets operation
type ListSecretsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The secret name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can specify only one sort order. The default order for
	// `TIMECREATED` is descending. The default order for `NAME` is ascending.
	SortBy ListSecretsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSecretsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The OCID of the vault.
	VaultId *string `mandatory:"false" contributesTo:"query" name:"vaultId"`

	// A filter that returns only resources that match the specified lifecycle state. The state value is case-insensitive.
	LifecycleState SecretSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecretsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecretsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecretsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSecretsResponse wrapper for the ListSecrets operation
type ListSecretsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SecretSummary instances
	Items []SecretSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSecretsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecretsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecretsSortByEnum Enum with underlying type: string
type ListSecretsSortByEnum string

// Set of constants representing the allowable values for ListSecretsSortByEnum
const (
	ListSecretsSortByTimecreated ListSecretsSortByEnum = "TIMECREATED"
	ListSecretsSortByName        ListSecretsSortByEnum = "NAME"
)

var mappingListSecretsSortBy = map[string]ListSecretsSortByEnum{
	"TIMECREATED": ListSecretsSortByTimecreated,
	"NAME":        ListSecretsSortByName,
}

// GetListSecretsSortByEnumValues Enumerates the set of values for ListSecretsSortByEnum
func GetListSecretsSortByEnumValues() []ListSecretsSortByEnum {
	values := make([]ListSecretsSortByEnum, 0)
	for _, v := range mappingListSecretsSortBy {
		values = append(values, v)
	}
	return values
}

// ListSecretsSortOrderEnum Enum with underlying type: string
type ListSecretsSortOrderEnum string

// Set of constants representing the allowable values for ListSecretsSortOrderEnum
const (
	ListSecretsSortOrderAsc  ListSecretsSortOrderEnum = "ASC"
	ListSecretsSortOrderDesc ListSecretsSortOrderEnum = "DESC"
)

var mappingListSecretsSortOrder = map[string]ListSecretsSortOrderEnum{
	"ASC":  ListSecretsSortOrderAsc,
	"DESC": ListSecretsSortOrderDesc,
}

// GetListSecretsSortOrderEnumValues Enumerates the set of values for ListSecretsSortOrderEnum
func GetListSecretsSortOrderEnumValues() []ListSecretsSortOrderEnum {
	values := make([]ListSecretsSortOrderEnum, 0)
	for _, v := range mappingListSecretsSortOrder {
		values = append(values, v)
	}
	return values
}
