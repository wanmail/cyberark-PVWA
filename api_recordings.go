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


// RecordingsApiService RecordingsApi service
type RecordingsApiService service

type ApiRecordingsGetRecordingRequest struct {
	ctx context.Context
	ApiService *RecordingsApiService
	recordingId string
	returnURL *string
}

func (r ApiRecordingsGetRecordingRequest) ReturnURL(returnURL string) ApiRecordingsGetRecordingRequest {
	r.returnURL = &returnURL
	return r
}

func (r ApiRecordingsGetRecordingRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RecordingsGetRecordingExecute(r)
}

/*
RecordingsGetRecording Method for RecordingsGetRecording

This method returns details of a recorded session.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param recordingId The unique ID of the PSM Recorded Session.
 @return ApiRecordingsGetRecordingRequest
*/
func (a *RecordingsApiService) RecordingsGetRecording(ctx context.Context, recordingId string) ApiRecordingsGetRecordingRequest {
	return ApiRecordingsGetRecordingRequest{
		ApiService: a,
		ctx: ctx,
		recordingId: recordingId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RecordingsApiService) RecordingsGetRecordingExecute(r ApiRecordingsGetRecordingRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingsApiService.RecordingsGetRecording")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Recordings/{recordingId}"
	localVarPath = strings.Replace(localVarPath, "{"+"recordingId"+"}", url.PathEscape(parameterToString(r.recordingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnURL != nil {
		localVarQueryParams.Add("returnURL", parameterToString(*r.returnURL, ""))
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

type ApiRecordingsGetRecordingActivitiesRequest struct {
	ctx context.Context
	ApiService *RecordingsApiService
	recordingId string
}

func (r ApiRecordingsGetRecordingActivitiesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RecordingsGetRecordingActivitiesExecute(r)
}

/*
RecordingsGetRecordingActivities Method for RecordingsGetRecordingActivities

This method returns activities of a recorded session.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param recordingId The unique ID of the PSM Recorded Session.
 @return ApiRecordingsGetRecordingActivitiesRequest
*/
func (a *RecordingsApiService) RecordingsGetRecordingActivities(ctx context.Context, recordingId string) ApiRecordingsGetRecordingActivitiesRequest {
	return ApiRecordingsGetRecordingActivitiesRequest{
		ApiService: a,
		ctx: ctx,
		recordingId: recordingId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RecordingsApiService) RecordingsGetRecordingActivitiesExecute(r ApiRecordingsGetRecordingActivitiesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingsApiService.RecordingsGetRecordingActivities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Recordings/{recordingId}/activities"
	localVarPath = strings.Replace(localVarPath, "{"+"recordingId"+"}", url.PathEscape(parameterToString(r.recordingId, "")), -1)

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

type ApiRecordingsGetRecordingFileValidityStatusRequest struct {
	ctx context.Context
	ApiService *RecordingsApiService
	recordingId string
}

func (r ApiRecordingsGetRecordingFileValidityStatusRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RecordingsGetRecordingFileValidityStatusExecute(r)
}

/*
RecordingsGetRecordingFileValidityStatus Method for RecordingsGetRecordingFileValidityStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param recordingId
 @return ApiRecordingsGetRecordingFileValidityStatusRequest
*/
func (a *RecordingsApiService) RecordingsGetRecordingFileValidityStatus(ctx context.Context, recordingId string) ApiRecordingsGetRecordingFileValidityStatusRequest {
	return ApiRecordingsGetRecordingFileValidityStatusRequest{
		ApiService: a,
		ctx: ctx,
		recordingId: recordingId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RecordingsApiService) RecordingsGetRecordingFileValidityStatusExecute(r ApiRecordingsGetRecordingFileValidityStatusRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingsApiService.RecordingsGetRecordingFileValidityStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Recordings/{recordingId}/valid"
	localVarPath = strings.Replace(localVarPath, "{"+"recordingId"+"}", url.PathEscape(parameterToString(r.recordingId, "")), -1)

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

type ApiRecordingsGetRecordingPropertiesRequest struct {
	ctx context.Context
	ApiService *RecordingsApiService
	recordingId string
	returnURL *string
}

func (r ApiRecordingsGetRecordingPropertiesRequest) ReturnURL(returnURL string) ApiRecordingsGetRecordingPropertiesRequest {
	r.returnURL = &returnURL
	return r
}

func (r ApiRecordingsGetRecordingPropertiesRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RecordingsGetRecordingPropertiesExecute(r)
}

/*
RecordingsGetRecordingProperties Method for RecordingsGetRecordingProperties

This method returns properties of a recorded session.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param recordingId The unique ID of the PSM Recorded Session.
 @return ApiRecordingsGetRecordingPropertiesRequest
*/
func (a *RecordingsApiService) RecordingsGetRecordingProperties(ctx context.Context, recordingId string) ApiRecordingsGetRecordingPropertiesRequest {
	return ApiRecordingsGetRecordingPropertiesRequest{
		ApiService: a,
		ctx: ctx,
		recordingId: recordingId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RecordingsApiService) RecordingsGetRecordingPropertiesExecute(r ApiRecordingsGetRecordingPropertiesRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingsApiService.RecordingsGetRecordingProperties")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Recordings/{recordingId}/properties"
	localVarPath = strings.Replace(localVarPath, "{"+"recordingId"+"}", url.PathEscape(parameterToString(r.recordingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnURL != nil {
		localVarQueryParams.Add("returnURL", parameterToString(*r.returnURL, ""))
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

type ApiRecordingsGetRecordingsRequest struct {
	ctx context.Context
	ApiService *RecordingsApiService
	returnURL *string
	limit *int64
	offset *int64
	useCache *bool
	sort *[]string
	search *string
	activities *string
	fromTime *int64
	toTime *int64
	safe *string
}

// The returned URL.
func (r ApiRecordingsGetRecordingsRequest) ReturnURL(returnURL string) ApiRecordingsGetRecordingsRequest {
	r.returnURL = &returnURL
	return r
}

// Determines the number of recordings that are returned in the list. The maximum value is defined in the MaxRecords property in Options -&gt; Privileged Session Management -&gt; General Setting -&gt; Search Properties.
func (r ApiRecordingsGetRecordingsRequest) Limit(limit int64) ApiRecordingsGetRecordingsRequest {
	r.limit = &limit
	return r
}

// Determines which recording results will be returned, according to a specific place in the returned list. This value defines the recording&#39;s place in the list and how many results will be skipped.
func (r ApiRecordingsGetRecordingsRequest) Offset(offset int64) ApiRecordingsGetRecordingsRequest {
	r.offset = &offset
	return r
}

func (r ApiRecordingsGetRecordingsRequest) UseCache(useCache bool) ApiRecordingsGetRecordingsRequest {
	r.useCache = &useCache
	return r
}

// The sort can be done by each property on the recording file: RiskScore, FileName, SafeName, FolderName, PSMVaultUserName, FromIP, RemoteMachine, Client, Protocol, AccountUserName, AccountAddress, AccountPlatformID, PSMStartTime, TicketID. The sort can be in ascending or descending order. To sort in descending order, specify &#39; - &#39; (dash) before the recording property by which to sort.
func (r ApiRecordingsGetRecordingsRequest) Sort(sort []string) ApiRecordingsGetRecordingsRequest {
	r.sort = &sort
	return r
}

// Returns recordings that are filtered by properties that contain the specified search text.
func (r ApiRecordingsGetRecordingsRequest) Search(search string) ApiRecordingsGetRecordingsRequest {
	r.search = &search
	return r
}

// Returns recordings with specific activities.
func (r ApiRecordingsGetRecordingsRequest) Activities(activities string) ApiRecordingsGetRecordingsRequest {
	r.activities = &activities
	return r
}

// Returns recordings from a specific date.
func (r ApiRecordingsGetRecordingsRequest) FromTime(fromTime int64) ApiRecordingsGetRecordingsRequest {
	r.fromTime = &fromTime
	return r
}

// Returns recordings before a specific date.
func (r ApiRecordingsGetRecordingsRequest) ToTime(toTime int64) ApiRecordingsGetRecordingsRequest {
	r.toTime = &toTime
	return r
}

// Returns recordings from a specific Safe.
func (r ApiRecordingsGetRecordingsRequest) Safe(safe string) ApiRecordingsGetRecordingsRequest {
	r.safe = &safe
	return r
}

func (r ApiRecordingsGetRecordingsRequest) Execute() (*SessionData, *http.Response, error) {
	return r.ApiService.RecordingsGetRecordingsExecute(r)
}

/*
RecordingsGetRecordings Method for RecordingsGetRecordings

This method returns the details of recordings of PSM, PSMP or OPM sessions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRecordingsGetRecordingsRequest
*/
func (a *RecordingsApiService) RecordingsGetRecordings(ctx context.Context) ApiRecordingsGetRecordingsRequest {
	return ApiRecordingsGetRecordingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SessionData
func (a *RecordingsApiService) RecordingsGetRecordingsExecute(r ApiRecordingsGetRecordingsRequest) (*SessionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SessionData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingsApiService.RecordingsGetRecordings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Recordings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.returnURL != nil {
		localVarQueryParams.Add("returnURL", parameterToString(*r.returnURL, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.useCache != nil {
		localVarQueryParams.Add("useCache", parameterToString(*r.useCache, ""))
	}
	if r.sort != nil {
		t := *r.sort
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("sort", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("sort", parameterToString(t, "multi"))
		}
	}
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.activities != nil {
		localVarQueryParams.Add("activities", parameterToString(*r.activities, ""))
	}
	if r.fromTime != nil {
		localVarQueryParams.Add("fromTime", parameterToString(*r.fromTime, ""))
	}
	if r.toTime != nil {
		localVarQueryParams.Add("toTime", parameterToString(*r.toTime, ""))
	}
	if r.safe != nil {
		localVarQueryParams.Add("safe", parameterToString(*r.safe, ""))
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

type ApiRecordingsPlayRecordingRequest struct {
	ctx context.Context
	ApiService *RecordingsApiService
	recordingId string
}

func (r ApiRecordingsPlayRecordingRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.RecordingsPlayRecordingExecute(r)
}

/*
RecordingsPlayRecording Method for RecordingsPlayRecording

This method returns data stream of a recorded session.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param recordingId The unique ID of the PSM Recorded Session.
 @return ApiRecordingsPlayRecordingRequest
*/
func (a *RecordingsApiService) RecordingsPlayRecording(ctx context.Context, recordingId string) ApiRecordingsPlayRecordingRequest {
	return ApiRecordingsPlayRecordingRequest{
		ApiService: a,
		ctx: ctx,
		recordingId: recordingId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *RecordingsApiService) RecordingsPlayRecordingExecute(r ApiRecordingsPlayRecordingRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordingsApiService.RecordingsPlayRecording")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/Recordings/{recordingId}/Play"
	localVarPath = strings.Replace(localVarPath, "{"+"recordingId"+"}", url.PathEscape(parameterToString(r.recordingId, "")), -1)

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
