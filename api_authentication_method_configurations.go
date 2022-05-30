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


// AuthenticationMethodConfigurationsApiService AuthenticationMethodConfigurationsApi service
type AuthenticationMethodConfigurationsApiService service

type ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest struct {
	ctx context.Context
	ApiService *AuthenticationMethodConfigurationsApiService
	authenticationMethod *AuthenticationMethod
}

// The authentication method information
func (r ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest) AuthenticationMethod(authenticationMethod AuthenticationMethod) ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest {
	r.authenticationMethod = &authenticationMethod
	return r
}

func (r ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest) Execute() (*AuthenticationMethod, *http.Response, error) {
	return r.ApiService.AuthenticationMethodConfigurationsAddConfigurationValueExecute(r)
}

/*
AuthenticationMethodConfigurationsAddConfigurationValue Method for AuthenticationMethodConfigurationsAddConfigurationValue

This method adds a new authentication method.
Any user who is a member of the Vault Admins group can run this web service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest
*/
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsAddConfigurationValue(ctx context.Context) ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest {
	return ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthenticationMethod
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsAddConfigurationValueExecute(r ApiAuthenticationMethodConfigurationsAddConfigurationValueRequest) (*AuthenticationMethod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticationMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationMethodConfigurationsApiService.AuthenticationMethodConfigurationsAddConfigurationValue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/AuthenticationMethods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authenticationMethod == nil {
		return localVarReturnValue, nil, reportError("authenticationMethod is required and must be specified")
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
	localVarPostBody = r.authenticationMethod
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

type ApiAuthenticationMethodConfigurationsDeleteAuthenticationMethodRequest struct {
	ctx context.Context
	ApiService *AuthenticationMethodConfigurationsApiService
	id string
}

func (r ApiAuthenticationMethodConfigurationsDeleteAuthenticationMethodRequest) Execute() (*http.Response, error) {
	return r.ApiService.AuthenticationMethodConfigurationsDeleteAuthenticationMethodExecute(r)
}

/*
AuthenticationMethodConfigurationsDeleteAuthenticationMethod Method for AuthenticationMethodConfigurationsDeleteAuthenticationMethod

This method deletes a specific authentication method.
Any user who is a member of the Vault Admins group can run this web service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The authentication method unique identifier.
 @return ApiAuthenticationMethodConfigurationsDeleteAuthenticationMethodRequest
*/
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsDeleteAuthenticationMethod(ctx context.Context, id string) ApiAuthenticationMethodConfigurationsDeleteAuthenticationMethodRequest {
	return ApiAuthenticationMethodConfigurationsDeleteAuthenticationMethodRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsDeleteAuthenticationMethodExecute(r ApiAuthenticationMethodConfigurationsDeleteAuthenticationMethodRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationMethodConfigurationsApiService.AuthenticationMethodConfigurationsDeleteAuthenticationMethod")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/AuthenticationMethods/{id}"
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAuthenticationMethodConfigurationsGetAuthenticationMethodRequest struct {
	ctx context.Context
	ApiService *AuthenticationMethodConfigurationsApiService
	id string
}

func (r ApiAuthenticationMethodConfigurationsGetAuthenticationMethodRequest) Execute() (*AuthenticationMethod, *http.Response, error) {
	return r.ApiService.AuthenticationMethodConfigurationsGetAuthenticationMethodExecute(r)
}

/*
AuthenticationMethodConfigurationsGetAuthenticationMethod Method for AuthenticationMethodConfigurationsGetAuthenticationMethod

This method returns a specific authentication method.
Any user who is a member of the Vault Admins group can run this web service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The authentication method unique identifier.
 @return ApiAuthenticationMethodConfigurationsGetAuthenticationMethodRequest
*/
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsGetAuthenticationMethod(ctx context.Context, id string) ApiAuthenticationMethodConfigurationsGetAuthenticationMethodRequest {
	return ApiAuthenticationMethodConfigurationsGetAuthenticationMethodRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AuthenticationMethod
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsGetAuthenticationMethodExecute(r ApiAuthenticationMethodConfigurationsGetAuthenticationMethodRequest) (*AuthenticationMethod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticationMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationMethodConfigurationsApiService.AuthenticationMethodConfigurationsGetAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/AuthenticationMethods/{id}"
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

type ApiAuthenticationMethodConfigurationsGetAuthenticationMethodsRequest struct {
	ctx context.Context
	ApiService *AuthenticationMethodConfigurationsApiService
}

func (r ApiAuthenticationMethodConfigurationsGetAuthenticationMethodsRequest) Execute() (*AuthenticationMethodsResponse, *http.Response, error) {
	return r.ApiService.AuthenticationMethodConfigurationsGetAuthenticationMethodsExecute(r)
}

/*
AuthenticationMethodConfigurationsGetAuthenticationMethods Method for AuthenticationMethodConfigurationsGetAuthenticationMethods

This method returns a list of all existing authentication methods.
Any user who is a member of the Vault Admins group can run this web service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthenticationMethodConfigurationsGetAuthenticationMethodsRequest
*/
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsGetAuthenticationMethods(ctx context.Context) ApiAuthenticationMethodConfigurationsGetAuthenticationMethodsRequest {
	return ApiAuthenticationMethodConfigurationsGetAuthenticationMethodsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthenticationMethodsResponse
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsGetAuthenticationMethodsExecute(r ApiAuthenticationMethodConfigurationsGetAuthenticationMethodsRequest) (*AuthenticationMethodsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticationMethodsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationMethodConfigurationsApiService.AuthenticationMethodConfigurationsGetAuthenticationMethods")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/AuthenticationMethods"

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

type ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest struct {
	ctx context.Context
	ApiService *AuthenticationMethodConfigurationsApiService
	id string
	value *AuthenticationMethodBase
}

// The authentication method information
func (r ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest) Value(value AuthenticationMethodBase) ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest {
	r.value = &value
	return r
}

func (r ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest) Execute() (*AuthenticationMethod, *http.Response, error) {
	return r.ApiService.AuthenticationMethodConfigurationsUpdateAuthenticationMethodExecute(r)
}

/*
AuthenticationMethodConfigurationsUpdateAuthenticationMethod Method for AuthenticationMethodConfigurationsUpdateAuthenticationMethod

This method updates the properties for a specific authentication method.
Any user who is a member of the Vault Admins group can run this web service.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The authentication method unique identifier.
 @return ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest
*/
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsUpdateAuthenticationMethod(ctx context.Context, id string) ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest {
	return ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AuthenticationMethod
func (a *AuthenticationMethodConfigurationsApiService) AuthenticationMethodConfigurationsUpdateAuthenticationMethodExecute(r ApiAuthenticationMethodConfigurationsUpdateAuthenticationMethodRequest) (*AuthenticationMethod, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthenticationMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationMethodConfigurationsApiService.AuthenticationMethodConfigurationsUpdateAuthenticationMethod")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Configuration/AuthenticationMethods/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.value == nil {
		return localVarReturnValue, nil, reportError("value is required and must be specified")
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
	localVarPostBody = r.value
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
