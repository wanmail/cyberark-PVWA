# \PSMServersApi

All URIs are relative to *https://10.40.0.212/PasswordVault*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PSMServersGetPSMServersController**](PSMServersApi.md#PSMServersGetPSMServersController) | **Get** /api/PSM/Servers | 



## PSMServersGetPSMServersController

> []PSMServer PSMServersGetPSMServersController(ctx).Execute()





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
    resp, r, err := apiClient.PSMServersApi.PSMServersGetPSMServersController(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PSMServersApi.PSMServersGetPSMServersController``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PSMServersGetPSMServersController`: []PSMServer
    fmt.Fprintf(os.Stdout, "Response from `PSMServersApi.PSMServersGetPSMServersController`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPSMServersGetPSMServersControllerRequest struct via the builder pattern


### Return type

[**[]PSMServer**](PSMServer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, application/xml, text/xml, multipart/form-data, application/vnd.cyberark.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

