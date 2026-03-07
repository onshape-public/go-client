# \AccountApi

All URIs are relative to *https://cad.onshape.com/api/v14*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPurchaseNew**](AccountApi.md#CancelPurchaseNew) | **Delete** /accounts/{aid}/purchases/{pid} | Cancel a recurring subscription for the specified account ID and purchase ID.
[**ConsumePurchase**](AccountApi.md#ConsumePurchase) | **Post** /accounts/purchases/{pid}/consume | Mark a purchase as consumed by the current user.
[**GetPlanPurchases**](AccountApi.md#GetPlanPurchases) | **Get** /accounts/plans/{planId}/purchases | Get a list of all app purchases made for the specified plan.
[**GetPurchases**](AccountApi.md#GetPurchases) | **Get** /accounts/purchases | Get a list of all app purchases made by the current user.



## CancelPurchaseNew

> map[string]interface{} CancelPurchaseNew(ctx, aid, pid).CancelImmediately(cancelImmediately).Execute()

Cancel a recurring subscription for the specified account ID and purchase ID.

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
    aid := "aid_example" // string | 
    pid := "pid_example" // string | 
    cancelImmediately := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AccountApi.CancelPurchaseNew(context.Background(), aid, pid).CancelImmediately(cancelImmediately).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CancelPurchaseNew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelPurchaseNew`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CancelPurchaseNew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPurchaseNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cancelImmediately** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConsumePurchase

> BTPurchaseInfo ConsumePurchase(ctx, pid).BTPurchaseIdentityParams(bTPurchaseIdentityParams).Execute()

Mark a purchase as consumed by the current user.

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
    pid := "pid_example" // string | 
    bTPurchaseIdentityParams := *openapiclient.NewBTPurchaseIdentityParams() // BTPurchaseIdentityParams |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AccountApi.ConsumePurchase(context.Background(), pid).BTPurchaseIdentityParams(bTPurchaseIdentityParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ConsumePurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConsumePurchase`: BTPurchaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ConsumePurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumePurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTPurchaseIdentityParams** | [**BTPurchaseIdentityParams**](BTPurchaseIdentityParams.md) |  | 

### Return type

[**BTPurchaseInfo**](BTPurchaseInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanPurchases

> BTListResponseBTPurchaseInfo GetPlanPurchases(ctx, planId).Offset(offset).Limit(limit).Execute()

Get a list of all app purchases made for the specified plan.

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
    planId := "planId_example" // string | 
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AccountApi.GetPlanPurchases(context.Background(), planId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetPlanPurchases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlanPurchases`: BTListResponseBTPurchaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetPlanPurchases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTPurchaseInfo**](BTListResponseBTPurchaseInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchases

> []BTPurchaseInfo GetPurchases(ctx).All(all).OwnPurchaseOnly(ownPurchaseOnly).Execute()

Get a list of all app purchases made by the current user.



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
    all := true // bool |  (optional) (default to false)
    ownPurchaseOnly := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AccountApi.GetPurchases(context.Background()).All(all).OwnPurchaseOnly(ownPurchaseOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetPurchases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPurchases`: []BTPurchaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetPurchases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** |  | [default to false]
 **ownPurchaseOnly** | **bool** |  | [default to false]

### Return type

[**[]BTPurchaseInfo**](BTPurchaseInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

