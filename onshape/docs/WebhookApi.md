# \WebhookApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhookApi.md#CreateWebhook) | **Post** /webhooks | Create a new webhook.
[**GetWebhook**](WebhookApi.md#GetWebhook) | **Get** /webhooks/{webhookid} | Get webhook info by webhook ID.
[**GetWebhooks**](WebhookApi.md#GetWebhooks) | **Get** /webhooks | Get a list of all webhooks registered by a user or company.
[**PingWebhook**](WebhookApi.md#PingWebhook) | **Post** /webhooks/{webhookid}/ping | Ping a webhook.
[**UnregisterWebhook**](WebhookApi.md#UnregisterWebhook) | **Delete** /webhooks/{webhookid} | Unregister a webhook.
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Post** /webhooks/{webhookid} | Update a webhook.



## CreateWebhook

> BTWebhookInfo CreateWebhook(ctx).BTWebhookParams(bTWebhookParams).Execute()

Create a new webhook.

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
    bTWebhookParams := *openapiclient.NewBTWebhookParams() // BTWebhookParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.CreateWebhook(context.Background()).BTWebhookParams(bTWebhookParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: BTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTWebhookParams** | [**BTWebhookParams**](BTWebhookParams.md) |  | 

### Return type

[**BTWebhookInfo**](BTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> BTWebhookInfo GetWebhook(ctx, webhookid).Execute()

Get webhook info by webhook ID.

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
    webhookid := "webhookid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhook(context.Background(), webhookid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: BTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWebhookInfo**](BTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> BTListResponseBTWebhookInfo GetWebhooks(ctx).Company(company).User(user).Offset(offset).Limit(limit).Execute()

Get a list of all webhooks registered by a user or company.

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
    company := "company_example" // string |  (optional) (default to "")
    user := "user_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.GetWebhooks(context.Background()).Company(company).User(user).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.GetWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooks`: BTListResponseBTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **company** | **string** |  | [default to &quot;&quot;]
 **user** | **string** |  | 
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTWebhookInfo**](BTListResponseBTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingWebhook

> map[string]interface{} PingWebhook(ctx, webhookid).Execute()

Ping a webhook.

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
    webhookid := "webhookid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.PingWebhook(context.Background(), webhookid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.PingWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingWebhook`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.PingWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UnregisterWebhook

> map[string]interface{} UnregisterWebhook(ctx, webhookid).BlockNotification(blockNotification).Execute()

Unregister a webhook.

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
    webhookid := "webhookid_example" // string | 
    blockNotification := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.UnregisterWebhook(context.Background(), webhookid).BlockNotification(blockNotification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.UnregisterWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterWebhook`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.UnregisterWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockNotification** | **bool** |  | [default to false]

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


## UpdateWebhook

> BTWebhookInfo UpdateWebhook(ctx, webhookid).BTWebhookParams(bTWebhookParams).Execute()

Update a webhook.



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
    webhookid := "webhookid_example" // string | 
    bTWebhookParams := *openapiclient.NewBTWebhookParams() // BTWebhookParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookApi.UpdateWebhook(context.Background(), webhookid).BTWebhookParams(bTWebhookParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: BTWebhookInfo
    fmt.Fprintf(os.Stdout, "Response from `WebhookApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTWebhookParams** | [**BTWebhookParams**](BTWebhookParams.md) |  | 

### Return type

[**BTWebhookInfo**](BTWebhookInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

