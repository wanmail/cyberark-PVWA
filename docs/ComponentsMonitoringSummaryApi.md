# \ComponentsMonitoringSummaryApi

All URIs are relative to *https://localhost/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentsMonitoringSummaryGetComponentsMonitoringSummary**](ComponentsMonitoringSummaryApi.md#ComponentsMonitoringSummaryGetComponentsMonitoringSummary) | **Get** /api/ComponentsMonitoringSummary | 



## ComponentsMonitoringSummaryGetComponentsMonitoringSummary

> ComponentsMonitoringSummaryData ComponentsMonitoringSummaryGetComponentsMonitoringSummary(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ComponentsMonitoringSummaryApi.ComponentsMonitoringSummaryGetComponentsMonitoringSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ComponentsMonitoringSummaryApi.ComponentsMonitoringSummaryGetComponentsMonitoringSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComponentsMonitoringSummaryGetComponentsMonitoringSummary`: ComponentsMonitoringSummaryData
    fmt.Fprintf(os.Stdout, "Response from `ComponentsMonitoringSummaryApi.ComponentsMonitoringSummaryGetComponentsMonitoringSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiComponentsMonitoringSummaryGetComponentsMonitoringSummaryRequest struct via the builder pattern


### Return type

[**ComponentsMonitoringSummaryData**](ComponentsMonitoringSummaryData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

