# \EventApi

All URIs are relative to *https://cad.onshape.com/api/v8*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FireEvent**](EventApi.md#FireEvent) | **Post** /events | Fire an asynchronous event.



## FireEvent

> map[string]interface{} FireEvent(ctx).BTEventParams(bTEventParams).Execute()

Fire an asynchronous event.

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
    bTEventParams := openapiclient.BTEventParams{BTDocumentOpenEventParams: openapiclient.NewBTDocumentOpenEventParams()} // BTEventParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventApi.FireEvent(context.Background()).BTEventParams(bTEventParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventApi.FireEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FireEvent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EventApi.FireEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFireEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTEventParams** | [**BTEventParams**](BTEventParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/vnd.onshape.v2+json;charset=UTF-8;qs=0.2
- **Accept**: application/vnd.onshape.v2+json;charset=UTF-8;qs=0.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

