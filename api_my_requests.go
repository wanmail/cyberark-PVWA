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


// MyRequestsApiService MyRequestsApi service
type MyRequestsApiService service

type ApiMyRequestsCreateRequestRequest struct {
	ctx context.Context
	ApiService *MyRequestsApiService
	request *NewRequest
}

// The request to create.
func (r ApiMyRequestsCreateRequestRequest) Request(request NewRequest) ApiMyRequestsCreateRequestRequest {
	r.request = &request
	return r
}

func (r ApiMyRequestsCreateRequestRequest) Execute() (*Request, *http.Response, error) {
	return r.ApiService.MyRequestsCreateRequestExecute(r)
}

/*
MyRequestsCreateRequest Method for MyRequestsCreateRequest

This method creates an access request for a specific account. This account may be either a password account or an SSH Key account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMyRequestsCreateRequestRequest
*/
func (a *MyRequestsApiService) MyRequestsCreateRequest(ctx context.Context) ApiMyRequestsCreateRequestRequest {
	return ApiMyRequestsCreateRequestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Request
func (a *MyRequestsApiService) MyRequestsCreateRequestExecute(r ApiMyRequestsCreateRequestRequest) (*Request, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Request
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyRequestsApiService.MyRequestsCreateRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/MyRequests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, reportError("request is required and must be specified")
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
	localVarPostBody = r.request
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

type ApiMyRequestsDeleteMyRequestsRequest struct {
	ctx context.Context
	ApiService *MyRequestsApiService
	requestID string
	requestID2 *string
	safeName *string
	requestorUserName *string
	requestorReason *string
	userReason *string
	creationDate *int64
	operation *string
	expirationDate *int64
	operationType *int32
	accessType *string
	confirmationsLeft *int64
	accessFrom *int64
	accessTo *int64
	status *int32
	statusTitle *string
	invalidRequestReason *int32
	currentConfirmationLevel *int32
	requiredConfirmersCountLevel2 *int32
	ticketingSystemPropertiesName *string
	ticketingSystemPropertiesNumber *string
	ticketingSystemPropertiesStatus *string
	additionalInfo *map[string]interface{}
	accountDetailsAccountID *string
	accountDetailsProperties *map[string]interface{}
	confirmers *[]map[string]interface{}
}

// The request&#39;s unique ID, composed of the SafeName and internal RequestID.
func (r ApiMyRequestsDeleteMyRequestsRequest) RequestID2(requestID2 string) ApiMyRequestsDeleteMyRequestsRequest {
	r.requestID2 = &requestID2
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) SafeName(safeName string) ApiMyRequestsDeleteMyRequestsRequest {
	r.safeName = &safeName
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) RequestorUserName(requestorUserName string) ApiMyRequestsDeleteMyRequestsRequest {
	r.requestorUserName = &requestorUserName
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) RequestorReason(requestorReason string) ApiMyRequestsDeleteMyRequestsRequest {
	r.requestorReason = &requestorReason
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) UserReason(userReason string) ApiMyRequestsDeleteMyRequestsRequest {
	r.userReason = &userReason
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) CreationDate(creationDate int64) ApiMyRequestsDeleteMyRequestsRequest {
	r.creationDate = &creationDate
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) Operation(operation string) ApiMyRequestsDeleteMyRequestsRequest {
	r.operation = &operation
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) ExpirationDate(expirationDate int64) ApiMyRequestsDeleteMyRequestsRequest {
	r.expirationDate = &expirationDate
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) OperationType(operationType int32) ApiMyRequestsDeleteMyRequestsRequest {
	r.operationType = &operationType
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) AccessType(accessType string) ApiMyRequestsDeleteMyRequestsRequest {
	r.accessType = &accessType
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) ConfirmationsLeft(confirmationsLeft int64) ApiMyRequestsDeleteMyRequestsRequest {
	r.confirmationsLeft = &confirmationsLeft
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) AccessFrom(accessFrom int64) ApiMyRequestsDeleteMyRequestsRequest {
	r.accessFrom = &accessFrom
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) AccessTo(accessTo int64) ApiMyRequestsDeleteMyRequestsRequest {
	r.accessTo = &accessTo
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) Status(status int32) ApiMyRequestsDeleteMyRequestsRequest {
	r.status = &status
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) StatusTitle(statusTitle string) ApiMyRequestsDeleteMyRequestsRequest {
	r.statusTitle = &statusTitle
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) InvalidRequestReason(invalidRequestReason int32) ApiMyRequestsDeleteMyRequestsRequest {
	r.invalidRequestReason = &invalidRequestReason
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) CurrentConfirmationLevel(currentConfirmationLevel int32) ApiMyRequestsDeleteMyRequestsRequest {
	r.currentConfirmationLevel = &currentConfirmationLevel
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2 int32) ApiMyRequestsDeleteMyRequestsRequest {
	r.requiredConfirmersCountLevel2 = &requiredConfirmersCountLevel2
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) TicketingSystemPropertiesName(ticketingSystemPropertiesName string) ApiMyRequestsDeleteMyRequestsRequest {
	r.ticketingSystemPropertiesName = &ticketingSystemPropertiesName
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber string) ApiMyRequestsDeleteMyRequestsRequest {
	r.ticketingSystemPropertiesNumber = &ticketingSystemPropertiesNumber
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus string) ApiMyRequestsDeleteMyRequestsRequest {
	r.ticketingSystemPropertiesStatus = &ticketingSystemPropertiesStatus
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) AdditionalInfo(additionalInfo map[string]interface{}) ApiMyRequestsDeleteMyRequestsRequest {
	r.additionalInfo = &additionalInfo
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) AccountDetailsAccountID(accountDetailsAccountID string) ApiMyRequestsDeleteMyRequestsRequest {
	r.accountDetailsAccountID = &accountDetailsAccountID
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) AccountDetailsProperties(accountDetailsProperties map[string]interface{}) ApiMyRequestsDeleteMyRequestsRequest {
	r.accountDetailsProperties = &accountDetailsProperties
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) Confirmers(confirmers []map[string]interface{}) ApiMyRequestsDeleteMyRequestsRequest {
	r.confirmers = &confirmers
	return r
}

func (r ApiMyRequestsDeleteMyRequestsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.MyRequestsDeleteMyRequestsExecute(r)
}

/*
MyRequestsDeleteMyRequests Method for MyRequestsDeleteMyRequests

This method deletes a request made by a user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param requestID The request's unique ID, composed of the SafeName and internal RequestID.
 @return ApiMyRequestsDeleteMyRequestsRequest
*/
func (a *MyRequestsApiService) MyRequestsDeleteMyRequests(ctx context.Context, requestID string) ApiMyRequestsDeleteMyRequestsRequest {
	return ApiMyRequestsDeleteMyRequestsRequest{
		ApiService: a,
		ctx: ctx,
		requestID: requestID,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *MyRequestsApiService) MyRequestsDeleteMyRequestsExecute(r ApiMyRequestsDeleteMyRequestsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyRequestsApiService.MyRequestsDeleteMyRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/MyRequests/{RequestID}"
	localVarPath = strings.Replace(localVarPath, "{"+"RequestID"+"}", url.PathEscape(parameterToString(r.requestID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.requestID2 != nil {
		localVarQueryParams.Add("requestID", parameterToString(*r.requestID2, ""))
	}
	if r.safeName != nil {
		localVarQueryParams.Add("safeName", parameterToString(*r.safeName, ""))
	}
	if r.requestorUserName != nil {
		localVarQueryParams.Add("requestorUserName", parameterToString(*r.requestorUserName, ""))
	}
	if r.requestorReason != nil {
		localVarQueryParams.Add("requestorReason", parameterToString(*r.requestorReason, ""))
	}
	if r.userReason != nil {
		localVarQueryParams.Add("userReason", parameterToString(*r.userReason, ""))
	}
	if r.creationDate != nil {
		localVarQueryParams.Add("creationDate", parameterToString(*r.creationDate, ""))
	}
	if r.operation != nil {
		localVarQueryParams.Add("operation", parameterToString(*r.operation, ""))
	}
	if r.expirationDate != nil {
		localVarQueryParams.Add("expirationDate", parameterToString(*r.expirationDate, ""))
	}
	if r.operationType != nil {
		localVarQueryParams.Add("operationType", parameterToString(*r.operationType, ""))
	}
	if r.accessType != nil {
		localVarQueryParams.Add("accessType", parameterToString(*r.accessType, ""))
	}
	if r.confirmationsLeft != nil {
		localVarQueryParams.Add("confirmationsLeft", parameterToString(*r.confirmationsLeft, ""))
	}
	if r.accessFrom != nil {
		localVarQueryParams.Add("accessFrom", parameterToString(*r.accessFrom, ""))
	}
	if r.accessTo != nil {
		localVarQueryParams.Add("accessTo", parameterToString(*r.accessTo, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.statusTitle != nil {
		localVarQueryParams.Add("statusTitle", parameterToString(*r.statusTitle, ""))
	}
	if r.invalidRequestReason != nil {
		localVarQueryParams.Add("invalidRequestReason", parameterToString(*r.invalidRequestReason, ""))
	}
	if r.currentConfirmationLevel != nil {
		localVarQueryParams.Add("currentConfirmationLevel", parameterToString(*r.currentConfirmationLevel, ""))
	}
	if r.requiredConfirmersCountLevel2 != nil {
		localVarQueryParams.Add("requiredConfirmersCountLevel2", parameterToString(*r.requiredConfirmersCountLevel2, ""))
	}
	if r.ticketingSystemPropertiesName != nil {
		localVarQueryParams.Add("ticketingSystemProperties.name", parameterToString(*r.ticketingSystemPropertiesName, ""))
	}
	if r.ticketingSystemPropertiesNumber != nil {
		localVarQueryParams.Add("ticketingSystemProperties.number", parameterToString(*r.ticketingSystemPropertiesNumber, ""))
	}
	if r.ticketingSystemPropertiesStatus != nil {
		localVarQueryParams.Add("ticketingSystemProperties.status", parameterToString(*r.ticketingSystemPropertiesStatus, ""))
	}
	if r.additionalInfo != nil {
		localVarQueryParams.Add("additionalInfo", parameterToString(*r.additionalInfo, ""))
	}
	if r.accountDetailsAccountID != nil {
		localVarQueryParams.Add("accountDetails.accountID", parameterToString(*r.accountDetailsAccountID, ""))
	}
	if r.accountDetailsProperties != nil {
		localVarQueryParams.Add("accountDetails.properties", parameterToString(*r.accountDetailsProperties, ""))
	}
	if r.confirmers != nil {
		t := *r.confirmers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("confirmers", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("confirmers", parameterToString(t, "multi"))
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

type ApiMyRequestsGetRequestRequest struct {
	ctx context.Context
	ApiService *MyRequestsApiService
	requestID string
	requestID2 *string
	safeName *string
	requestorUserName *string
	requestorReason *string
	userReason *string
	creationDate *int64
	operation *string
	expirationDate *int64
	operationType *int32
	accessType *string
	confirmationsLeft *int64
	accessFrom *int64
	accessTo *int64
	status *int32
	statusTitle *string
	invalidRequestReason *int32
	currentConfirmationLevel *int32
	requiredConfirmersCountLevel2 *int32
	ticketingSystemPropertiesName *string
	ticketingSystemPropertiesNumber *string
	ticketingSystemPropertiesStatus *string
	additionalInfo *map[string]interface{}
	accountDetailsAccountID *string
	accountDetailsProperties *map[string]interface{}
	confirmers *[]map[string]interface{}
}

func (r ApiMyRequestsGetRequestRequest) RequestID2(requestID2 string) ApiMyRequestsGetRequestRequest {
	r.requestID2 = &requestID2
	return r
}

func (r ApiMyRequestsGetRequestRequest) SafeName(safeName string) ApiMyRequestsGetRequestRequest {
	r.safeName = &safeName
	return r
}

func (r ApiMyRequestsGetRequestRequest) RequestorUserName(requestorUserName string) ApiMyRequestsGetRequestRequest {
	r.requestorUserName = &requestorUserName
	return r
}

func (r ApiMyRequestsGetRequestRequest) RequestorReason(requestorReason string) ApiMyRequestsGetRequestRequest {
	r.requestorReason = &requestorReason
	return r
}

func (r ApiMyRequestsGetRequestRequest) UserReason(userReason string) ApiMyRequestsGetRequestRequest {
	r.userReason = &userReason
	return r
}

func (r ApiMyRequestsGetRequestRequest) CreationDate(creationDate int64) ApiMyRequestsGetRequestRequest {
	r.creationDate = &creationDate
	return r
}

func (r ApiMyRequestsGetRequestRequest) Operation(operation string) ApiMyRequestsGetRequestRequest {
	r.operation = &operation
	return r
}

func (r ApiMyRequestsGetRequestRequest) ExpirationDate(expirationDate int64) ApiMyRequestsGetRequestRequest {
	r.expirationDate = &expirationDate
	return r
}

func (r ApiMyRequestsGetRequestRequest) OperationType(operationType int32) ApiMyRequestsGetRequestRequest {
	r.operationType = &operationType
	return r
}

func (r ApiMyRequestsGetRequestRequest) AccessType(accessType string) ApiMyRequestsGetRequestRequest {
	r.accessType = &accessType
	return r
}

func (r ApiMyRequestsGetRequestRequest) ConfirmationsLeft(confirmationsLeft int64) ApiMyRequestsGetRequestRequest {
	r.confirmationsLeft = &confirmationsLeft
	return r
}

func (r ApiMyRequestsGetRequestRequest) AccessFrom(accessFrom int64) ApiMyRequestsGetRequestRequest {
	r.accessFrom = &accessFrom
	return r
}

func (r ApiMyRequestsGetRequestRequest) AccessTo(accessTo int64) ApiMyRequestsGetRequestRequest {
	r.accessTo = &accessTo
	return r
}

func (r ApiMyRequestsGetRequestRequest) Status(status int32) ApiMyRequestsGetRequestRequest {
	r.status = &status
	return r
}

func (r ApiMyRequestsGetRequestRequest) StatusTitle(statusTitle string) ApiMyRequestsGetRequestRequest {
	r.statusTitle = &statusTitle
	return r
}

func (r ApiMyRequestsGetRequestRequest) InvalidRequestReason(invalidRequestReason int32) ApiMyRequestsGetRequestRequest {
	r.invalidRequestReason = &invalidRequestReason
	return r
}

func (r ApiMyRequestsGetRequestRequest) CurrentConfirmationLevel(currentConfirmationLevel int32) ApiMyRequestsGetRequestRequest {
	r.currentConfirmationLevel = &currentConfirmationLevel
	return r
}

func (r ApiMyRequestsGetRequestRequest) RequiredConfirmersCountLevel2(requiredConfirmersCountLevel2 int32) ApiMyRequestsGetRequestRequest {
	r.requiredConfirmersCountLevel2 = &requiredConfirmersCountLevel2
	return r
}

func (r ApiMyRequestsGetRequestRequest) TicketingSystemPropertiesName(ticketingSystemPropertiesName string) ApiMyRequestsGetRequestRequest {
	r.ticketingSystemPropertiesName = &ticketingSystemPropertiesName
	return r
}

func (r ApiMyRequestsGetRequestRequest) TicketingSystemPropertiesNumber(ticketingSystemPropertiesNumber string) ApiMyRequestsGetRequestRequest {
	r.ticketingSystemPropertiesNumber = &ticketingSystemPropertiesNumber
	return r
}

func (r ApiMyRequestsGetRequestRequest) TicketingSystemPropertiesStatus(ticketingSystemPropertiesStatus string) ApiMyRequestsGetRequestRequest {
	r.ticketingSystemPropertiesStatus = &ticketingSystemPropertiesStatus
	return r
}

func (r ApiMyRequestsGetRequestRequest) AdditionalInfo(additionalInfo map[string]interface{}) ApiMyRequestsGetRequestRequest {
	r.additionalInfo = &additionalInfo
	return r
}

func (r ApiMyRequestsGetRequestRequest) AccountDetailsAccountID(accountDetailsAccountID string) ApiMyRequestsGetRequestRequest {
	r.accountDetailsAccountID = &accountDetailsAccountID
	return r
}

func (r ApiMyRequestsGetRequestRequest) AccountDetailsProperties(accountDetailsProperties map[string]interface{}) ApiMyRequestsGetRequestRequest {
	r.accountDetailsProperties = &accountDetailsProperties
	return r
}

func (r ApiMyRequestsGetRequestRequest) Confirmers(confirmers []map[string]interface{}) ApiMyRequestsGetRequestRequest {
	r.confirmers = &confirmers
	return r
}

func (r ApiMyRequestsGetRequestRequest) Execute() (*MyRequest, *http.Response, error) {
	return r.ApiService.MyRequestsGetRequestExecute(r)
}

/*
MyRequestsGetRequest Method for MyRequestsGetRequest

This method returns details of all the requests in My Requests list.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param requestID
 @return ApiMyRequestsGetRequestRequest
*/
func (a *MyRequestsApiService) MyRequestsGetRequest(ctx context.Context, requestID string) ApiMyRequestsGetRequestRequest {
	return ApiMyRequestsGetRequestRequest{
		ApiService: a,
		ctx: ctx,
		requestID: requestID,
	}
}

// Execute executes the request
//  @return MyRequest
func (a *MyRequestsApiService) MyRequestsGetRequestExecute(r ApiMyRequestsGetRequestRequest) (*MyRequest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MyRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyRequestsApiService.MyRequestsGetRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/MyRequests/{RequestID}"
	localVarPath = strings.Replace(localVarPath, "{"+"RequestID"+"}", url.PathEscape(parameterToString(r.requestID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.requestID2 != nil {
		localVarQueryParams.Add("requestID", parameterToString(*r.requestID2, ""))
	}
	if r.safeName != nil {
		localVarQueryParams.Add("safeName", parameterToString(*r.safeName, ""))
	}
	if r.requestorUserName != nil {
		localVarQueryParams.Add("requestorUserName", parameterToString(*r.requestorUserName, ""))
	}
	if r.requestorReason != nil {
		localVarQueryParams.Add("requestorReason", parameterToString(*r.requestorReason, ""))
	}
	if r.userReason != nil {
		localVarQueryParams.Add("userReason", parameterToString(*r.userReason, ""))
	}
	if r.creationDate != nil {
		localVarQueryParams.Add("creationDate", parameterToString(*r.creationDate, ""))
	}
	if r.operation != nil {
		localVarQueryParams.Add("operation", parameterToString(*r.operation, ""))
	}
	if r.expirationDate != nil {
		localVarQueryParams.Add("expirationDate", parameterToString(*r.expirationDate, ""))
	}
	if r.operationType != nil {
		localVarQueryParams.Add("operationType", parameterToString(*r.operationType, ""))
	}
	if r.accessType != nil {
		localVarQueryParams.Add("accessType", parameterToString(*r.accessType, ""))
	}
	if r.confirmationsLeft != nil {
		localVarQueryParams.Add("confirmationsLeft", parameterToString(*r.confirmationsLeft, ""))
	}
	if r.accessFrom != nil {
		localVarQueryParams.Add("accessFrom", parameterToString(*r.accessFrom, ""))
	}
	if r.accessTo != nil {
		localVarQueryParams.Add("accessTo", parameterToString(*r.accessTo, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.statusTitle != nil {
		localVarQueryParams.Add("statusTitle", parameterToString(*r.statusTitle, ""))
	}
	if r.invalidRequestReason != nil {
		localVarQueryParams.Add("invalidRequestReason", parameterToString(*r.invalidRequestReason, ""))
	}
	if r.currentConfirmationLevel != nil {
		localVarQueryParams.Add("currentConfirmationLevel", parameterToString(*r.currentConfirmationLevel, ""))
	}
	if r.requiredConfirmersCountLevel2 != nil {
		localVarQueryParams.Add("requiredConfirmersCountLevel2", parameterToString(*r.requiredConfirmersCountLevel2, ""))
	}
	if r.ticketingSystemPropertiesName != nil {
		localVarQueryParams.Add("ticketingSystemProperties.name", parameterToString(*r.ticketingSystemPropertiesName, ""))
	}
	if r.ticketingSystemPropertiesNumber != nil {
		localVarQueryParams.Add("ticketingSystemProperties.number", parameterToString(*r.ticketingSystemPropertiesNumber, ""))
	}
	if r.ticketingSystemPropertiesStatus != nil {
		localVarQueryParams.Add("ticketingSystemProperties.status", parameterToString(*r.ticketingSystemPropertiesStatus, ""))
	}
	if r.additionalInfo != nil {
		localVarQueryParams.Add("additionalInfo", parameterToString(*r.additionalInfo, ""))
	}
	if r.accountDetailsAccountID != nil {
		localVarQueryParams.Add("accountDetails.accountID", parameterToString(*r.accountDetailsAccountID, ""))
	}
	if r.accountDetailsProperties != nil {
		localVarQueryParams.Add("accountDetails.properties", parameterToString(*r.accountDetailsProperties, ""))
	}
	if r.confirmers != nil {
		t := *r.confirmers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("confirmers", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("confirmers", parameterToString(t, "multi"))
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

type ApiMyRequestsGetRequestsRequest struct {
	ctx context.Context
	ApiService *MyRequestsApiService
	onlyWaiting *bool
	expired *bool
}

// \&quot;Only requests waiting for approval will be listed.
func (r ApiMyRequestsGetRequestsRequest) OnlyWaiting(onlyWaiting bool) ApiMyRequestsGetRequestsRequest {
	r.onlyWaiting = &onlyWaiting
	return r
}

// Expired requests will be included in the list.
func (r ApiMyRequestsGetRequestsRequest) Expired(expired bool) ApiMyRequestsGetRequestsRequest {
	r.expired = &expired
	return r
}

func (r ApiMyRequestsGetRequestsRequest) Execute() (*MyRequest, *http.Response, error) {
	return r.ApiService.MyRequestsGetRequestsExecute(r)
}

/*
MyRequestsGetRequests Method for MyRequestsGetRequests

This method returns a list of the end user's requests.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMyRequestsGetRequestsRequest
*/
func (a *MyRequestsApiService) MyRequestsGetRequests(ctx context.Context) ApiMyRequestsGetRequestsRequest {
	return ApiMyRequestsGetRequestsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MyRequest
func (a *MyRequestsApiService) MyRequestsGetRequestsExecute(r ApiMyRequestsGetRequestsRequest) (*MyRequest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MyRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MyRequestsApiService.MyRequestsGetRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/MyRequests"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.onlyWaiting != nil {
		localVarQueryParams.Add("onlyWaiting", parameterToString(*r.onlyWaiting, ""))
	}
	if r.expired != nil {
		localVarQueryParams.Add("expired", parameterToString(*r.expired, ""))
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
