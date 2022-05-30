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
	"reflect"
)


// AutomaticOnboardingRulesApiService AutomaticOnboardingRulesApi service
type AutomaticOnboardingRulesApiService service

type ApiAutomaticOnboardingRulesDeleteAutomaticOnboardingRuleRequest struct {
	ctx context.Context
	ApiService *AutomaticOnboardingRulesApiService
	id int64
}

func (r ApiAutomaticOnboardingRulesDeleteAutomaticOnboardingRuleRequest) Execute() (*http.Response, error) {
	return r.ApiService.AutomaticOnboardingRulesDeleteAutomaticOnboardingRuleExecute(r)
}

/*
AutomaticOnboardingRulesDeleteAutomaticOnboardingRule Method for AutomaticOnboardingRulesDeleteAutomaticOnboardingRule

This method deletes an automatic onboarding rule from the Vault.
<para />
The user who runs this web service must belong to the following group: Vault Admins.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique ID of the rule to delete.
 @return ApiAutomaticOnboardingRulesDeleteAutomaticOnboardingRuleRequest
*/
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesDeleteAutomaticOnboardingRule(ctx context.Context, id int64) ApiAutomaticOnboardingRulesDeleteAutomaticOnboardingRuleRequest {
	return ApiAutomaticOnboardingRulesDeleteAutomaticOnboardingRuleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesDeleteAutomaticOnboardingRuleExecute(r ApiAutomaticOnboardingRulesDeleteAutomaticOnboardingRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutomaticOnboardingRulesApiService.AutomaticOnboardingRulesDeleteAutomaticOnboardingRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/AutomaticOnboardingRules/{id}"
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

type ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest struct {
	ctx context.Context
	ApiService *AutomaticOnboardingRulesApiService
	names *[]string
}

// List of names
func (r ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest) Names(names []string) ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest {
	r.names = &names
	return r
}

func (r ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest) Execute() (*AutomaticOnboardingRuleData, *http.Response, error) {
	return r.ApiService.AutomaticOnboardingRulesGetAutomaticOnboardingRulesExecute(r)
}

/*
AutomaticOnboardingRulesGetAutomaticOnboardingRules Method for AutomaticOnboardingRulesGetAutomaticOnboardingRules

This method returns information about all the defined onboarding rules.
The user who runs this web service must belong to the following group: Vault Admins

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest
*/
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesGetAutomaticOnboardingRules(ctx context.Context) ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest {
	return ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AutomaticOnboardingRuleData
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesGetAutomaticOnboardingRulesExecute(r ApiAutomaticOnboardingRulesGetAutomaticOnboardingRulesRequest) (*AutomaticOnboardingRuleData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AutomaticOnboardingRuleData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutomaticOnboardingRulesApiService.AutomaticOnboardingRulesGetAutomaticOnboardingRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/AutomaticOnboardingRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.names != nil {
		t := *r.names
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("names", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("names", parameterToString(t, "multi"))
		}
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

type ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest struct {
	ctx context.Context
	ApiService *AutomaticOnboardingRulesApiService
	data *AutomaticOnboardingRuleDataIn
}

// The onboarding rule to create.
func (r ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest) Data(data AutomaticOnboardingRuleDataIn) ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest {
	r.data = &data
	return r
}

func (r ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest) Execute() (*AutomaticOnboardingRuleData, *http.Response, error) {
	return r.ApiService.AutomaticOnboardingRulesPostAutomaticOnboardingRuleExecute(r)
}

/*
AutomaticOnboardingRulesPostAutomaticOnboardingRule Method for AutomaticOnboardingRulesPostAutomaticOnboardingRule

This method adds a new onboarding rule to the Vault that filters discovered local privileged pending accounts. When a discovered pending account matches a rule, it will automatically be onboarded to the Safe that is defined in the rule and the password will be reconciled. 

The Safe and the reconcile account must be created according to the rule’s definition before you run this API in order to onboard the pending account automatically. The reconcile account must be associated to the platform that is defined in the rule. 

The user who runs this web service must belong to the following group: Vault Admins.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest
*/
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesPostAutomaticOnboardingRule(ctx context.Context) ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest {
	return ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AutomaticOnboardingRuleData
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesPostAutomaticOnboardingRuleExecute(r ApiAutomaticOnboardingRulesPostAutomaticOnboardingRuleRequest) (*AutomaticOnboardingRuleData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AutomaticOnboardingRuleData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutomaticOnboardingRulesApiService.AutomaticOnboardingRulesPostAutomaticOnboardingRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/AutomaticOnboardingRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.data == nil {
		return localVarReturnValue, nil, reportError("data is required and must be specified")
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
	localVarPostBody = r.data
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

type ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest struct {
	ctx context.Context
	ApiService *AutomaticOnboardingRulesApiService
	id int64
	data *AutomaticOnboardingRuleDataIn
}

// The onboarding rule to create
func (r ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest) Data(data AutomaticOnboardingRuleDataIn) ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest {
	r.data = &data
	return r
}

func (r ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest) Execute() (*AutomaticOnboardingRuleData, *http.Response, error) {
	return r.ApiService.AutomaticOnboardingRulesUpdateAutomaticOnboardingRuleExecute(r)
}

/*
AutomaticOnboardingRulesUpdateAutomaticOnboardingRule Method for AutomaticOnboardingRulesUpdateAutomaticOnboardingRule

This method Updates an existing Automatic Onboarding Rule.
            
The user who runs this web service must belong to the following group: Vault Admins.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique ID of the rule to update
 @return ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest
*/
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesUpdateAutomaticOnboardingRule(ctx context.Context, id int64) ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest {
	return ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AutomaticOnboardingRuleData
func (a *AutomaticOnboardingRulesApiService) AutomaticOnboardingRulesUpdateAutomaticOnboardingRuleExecute(r ApiAutomaticOnboardingRulesUpdateAutomaticOnboardingRuleRequest) (*AutomaticOnboardingRuleData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AutomaticOnboardingRuleData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutomaticOnboardingRulesApiService.AutomaticOnboardingRulesUpdateAutomaticOnboardingRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/AutomaticOnboardingRules/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.data == nil {
		return localVarReturnValue, nil, reportError("data is required and must be specified")
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
	localVarPostBody = r.data
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