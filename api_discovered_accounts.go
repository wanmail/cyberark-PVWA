/*
Privileged Access Security REST API

Display the PVWA REST APIs below for a description of how to use them and try them out. Access information about additional REST APIs through the online documentation.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// DiscoveredAccountsApiService DiscoveredAccountsApi service
type DiscoveredAccountsApiService service

type ApiDiscoveredAccountsAddDiscoveredAccountRequest struct {
	ctx context.Context
	ApiService *DiscoveredAccountsApiService
	discoveredAccount *DiscoveredAccount
}

// The discovered account to add
func (r ApiDiscoveredAccountsAddDiscoveredAccountRequest) DiscoveredAccount(discoveredAccount DiscoveredAccount) ApiDiscoveredAccountsAddDiscoveredAccountRequest {
	r.discoveredAccount = &discoveredAccount
	return r
}

func (r ApiDiscoveredAccountsAddDiscoveredAccountRequest) Execute() (*DiscoveredAccountResponse, *http.Response, error) {
	return r.ApiService.DiscoveredAccountsAddDiscoveredAccountExecute(r)
}

/*
DiscoveredAccountsAddDiscoveredAccount Method for DiscoveredAccountsAddDiscoveredAccount

This RPC service adds newly discovered accounts and its dependencies.
The account is onboarded according to matching onboarding rules or added directly to the Pending Accounts list.

If the account already exists in the Pending Account list, it will be updated as needed.

If the onboarded process failed for any reason, the account will be added to the Pending Accounts list.

To run this web service, you need the following:

1. Owner of 'PasswordManager_Pending' Safe with the following permissions to add pending accounts:

a. Update Account Properties

b. Add Account

c. List Files

2. Owner of the target Safe of the on-boarding rule with the following permissions to on-board the account:

a. Update Account Properties

b. Add Account

c. Initiate CPM account management operations

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDiscoveredAccountsAddDiscoveredAccountRequest
*/
func (a *DiscoveredAccountsApiService) DiscoveredAccountsAddDiscoveredAccount(ctx context.Context) ApiDiscoveredAccountsAddDiscoveredAccountRequest {
	return ApiDiscoveredAccountsAddDiscoveredAccountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DiscoveredAccountResponse
func (a *DiscoveredAccountsApiService) DiscoveredAccountsAddDiscoveredAccountExecute(r ApiDiscoveredAccountsAddDiscoveredAccountRequest) (*DiscoveredAccountResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiscoveredAccountResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveredAccountsApiService.DiscoveredAccountsAddDiscoveredAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/DiscoveredAccounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.discoveredAccount == nil {
		return localVarReturnValue, nil, reportError("discoveredAccount is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/vnd.cyberark.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.discoveredAccount
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDiscoveredAccountsDeleteDiscoveredAccountsRequest struct {
	ctx context.Context
	ApiService *DiscoveredAccountsApiService
}

func (r ApiDiscoveredAccountsDeleteDiscoveredAccountsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.DiscoveredAccountsDeleteDiscoveredAccountsExecute(r)
}

/*
DiscoveredAccountsDeleteDiscoveredAccounts Method for DiscoveredAccountsDeleteDiscoveredAccounts

This method deletes all discovered accounts and related dependencies from the Pending Accounts list.

To run this web service, the user must be a member of the Vault admins group.
            
Note: When thousands of accounts are deleted, it may take a few minutes to complete the process.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDiscoveredAccountsDeleteDiscoveredAccountsRequest
*/
func (a *DiscoveredAccountsApiService) DiscoveredAccountsDeleteDiscoveredAccounts(ctx context.Context) ApiDiscoveredAccountsDeleteDiscoveredAccountsRequest {
	return ApiDiscoveredAccountsDeleteDiscoveredAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DiscoveredAccountsApiService) DiscoveredAccountsDeleteDiscoveredAccountsExecute(r ApiDiscoveredAccountsDeleteDiscoveredAccountsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveredAccountsApiService.DiscoveredAccountsDeleteDiscoveredAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/DiscoveredAccounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/vnd.cyberark.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDiscoveredAccountsGetDiscoveredAccountRequest struct {
	ctx context.Context
	ApiService *DiscoveredAccountsApiService
	id string
}

func (r ApiDiscoveredAccountsGetDiscoveredAccountRequest) Execute() (*GetDiscoveredAccount, *http.Response, error) {
	return r.ApiService.DiscoveredAccountsGetDiscoveredAccountExecute(r)
}

/*
DiscoveredAccountsGetDiscoveredAccount Method for DiscoveredAccountsGetDiscoveredAccount

This method returns information about a discovered account and its dependencies from the Pending Accounts list.
The discovered account is identified by its ID.

To run this web service, the user must be a member of the Vault admins group.

Note: Discovered accounts that were onboarded either manually or automatically, according to predefined rules, won't be returned using this method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The account's unique ID.
 @return ApiDiscoveredAccountsGetDiscoveredAccountRequest
*/
func (a *DiscoveredAccountsApiService) DiscoveredAccountsGetDiscoveredAccount(ctx context.Context, id string) ApiDiscoveredAccountsGetDiscoveredAccountRequest {
	return ApiDiscoveredAccountsGetDiscoveredAccountRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetDiscoveredAccount
func (a *DiscoveredAccountsApiService) DiscoveredAccountsGetDiscoveredAccountExecute(r ApiDiscoveredAccountsGetDiscoveredAccountRequest) (*GetDiscoveredAccount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDiscoveredAccount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveredAccountsApiService.DiscoveredAccountsGetDiscoveredAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/DiscoveredAccounts/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/vnd.cyberark.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDiscoveredAccountsGetDiscoveredAccountsRequest struct {
	ctx context.Context
	ApiService *DiscoveredAccountsApiService
	search *string
	searchType *string
	offset *int32
	limit *int32
	filter *string
}

// Searches according to REST standard (search&#x3D;\&quot;search word\&quot;).    Search is supported for userName and address.
func (r ApiDiscoveredAccountsGetDiscoveredAccountsRequest) Search(search string) ApiDiscoveredAccountsGetDiscoveredAccountsRequest {
	r.search = &search
	return r
}

// The type of search to perform.     The keyword can either be contained within the account property values, or at the beginning of the value  specified in the search parameter.    When using a keyword at the beginning of a value, performance is enhanced.    Valid values: contains (default) or startswith
func (r ApiDiscoveredAccountsGetDiscoveredAccountsRequest) SearchType(searchType string) ApiDiscoveredAccountsGetDiscoveredAccountsRequest {
	r.searchType = &searchType
	return r
}

// The offset of the first returned accounts into the list of results.
func (r ApiDiscoveredAccountsGetDiscoveredAccountsRequest) Offset(offset int32) ApiDiscoveredAccountsGetDiscoveredAccountsRequest {
	r.offset = &offset
	return r
}

// The maximum number of returned accounts.     If not specified, the server limits the results to 100.    The maximum number that can be specified is 1000.
func (r ApiDiscoveredAccountsGetDiscoveredAccountsRequest) Limit(limit int32) ApiDiscoveredAccountsGetDiscoveredAccountsRequest {
	r.limit = &limit
	return r
}

// Search for accounts filtered by:    • platformType - Get discovered accounts of a specific platform, using the platform type name.        Usage: filter&#x3D;platformType eq Windows Server Local            Type: string            Valid values:                o Windows Server Local                    o Windows Desktop Local                    o Windows Domain                    o Unix                    o Unix SSH Key                    o AWS                    o AWS Access Keys                    o Azure Password Management        • privileged - Get only privileged or non-privileged discovered accounts.        Usage: filter&#x3D;privileged eq true            Type: boolean            Valid values: true/false        • accountEnabled - Get only enabled or disabled discovered accounts.           Usage: filter&#x3D;accountEnabled eq true            Type: boolean            Valid values: true/false                  Note: To use more than one filter, you can use the AND operator.         filter&#x3D;platformType eq Windows Server Local AND privileged eq true
func (r ApiDiscoveredAccountsGetDiscoveredAccountsRequest) Filter(filter string) ApiDiscoveredAccountsGetDiscoveredAccountsRequest {
	r.filter = &filter
	return r
}

func (r ApiDiscoveredAccountsGetDiscoveredAccountsRequest) Execute() (*GetDiscoveredAccountsResponse, *http.Response, error) {
	return r.ApiService.DiscoveredAccountsGetDiscoveredAccountsExecute(r)
}

/*
DiscoveredAccountsGetDiscoveredAccounts Method for DiscoveredAccountsGetDiscoveredAccounts

This method returns a list of all the discovered accounts from the Pending Accounts list.

To run this web service, the user must be a member of the Vault admins group.

Note: Discovered accounts that were onboarded either manually or automatically, according to predefined rules, won't be returned using this method.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDiscoveredAccountsGetDiscoveredAccountsRequest
*/
func (a *DiscoveredAccountsApiService) DiscoveredAccountsGetDiscoveredAccounts(ctx context.Context) ApiDiscoveredAccountsGetDiscoveredAccountsRequest {
	return ApiDiscoveredAccountsGetDiscoveredAccountsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetDiscoveredAccountsResponse
func (a *DiscoveredAccountsApiService) DiscoveredAccountsGetDiscoveredAccountsExecute(r ApiDiscoveredAccountsGetDiscoveredAccountsRequest) (*GetDiscoveredAccountsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDiscoveredAccountsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DiscoveredAccountsApiService.DiscoveredAccountsGetDiscoveredAccounts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/DiscoveredAccounts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.searchType != nil {
		localVarQueryParams.Add("searchType", parameterToString(*r.searchType, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml", "multipart/form-data", "application/vnd.cyberark.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}