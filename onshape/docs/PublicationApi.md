# \PublicationApi

All URIs are relative to *https://cad.onshape.com/api/v8*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItemToPublication**](PublicationApi.md#AddItemToPublication) | **Post** /publications/{pid}/item | Add an item in a publication.
[**AddItemsToPublication**](PublicationApi.md#AddItemsToPublication) | **Post** /publications/{pid}/items | Add publication items in bulk.
[**CreatePublication**](PublicationApi.md#CreatePublication) | **Post** /publications | Create a new publication.
[**DeletePublication**](PublicationApi.md#DeletePublication) | **Delete** /publications/{pid} | Delete a publication.
[**DeletePublicationItem**](PublicationApi.md#DeletePublicationItem) | **Delete** /publications/{pid}/item/{iid} | Remove an item from a publication.
[**GetPublicationItems**](PublicationApi.md#GetPublicationItems) | **Get** /publications/{pid}/items | Get all items in a publication.
[**UpdatePublicationAttributes**](PublicationApi.md#UpdatePublicationAttributes) | **Post** /publications/{pid} | Update publication&#39;s attributes name, description, and notes.



## AddItemToPublication

> BTPublicationInfo AddItemToPublication(ctx, pid).BTPublicationItemParams(bTPublicationItemParams).Execute()

Add an item in a publication.

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
    pid := "pid_example" // string | Publication ID.
    bTPublicationItemParams := *openapiclient.NewBTPublicationItemParams() // BTPublicationItemParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationApi.AddItemToPublication(context.Background(), pid).BTPublicationItemParams(bTPublicationItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationApi.AddItemToPublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddItemToPublication`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationApi.AddItemToPublication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** | Publication ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddItemToPublicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTPublicationItemParams** | [**BTPublicationItemParams**](BTPublicationItemParams.md) |  | 

### Return type

[**BTPublicationInfo**](BTPublicationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddItemsToPublication

> BTPublicationInfo AddItemsToPublication(ctx, pid).BTPublicationBulkItemParams(bTPublicationBulkItemParams).Execute()

Add publication items in bulk.

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
    pid := "pid_example" // string | Publication ID.
    bTPublicationBulkItemParams := *openapiclient.NewBTPublicationBulkItemParams() // BTPublicationBulkItemParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationApi.AddItemsToPublication(context.Background(), pid).BTPublicationBulkItemParams(bTPublicationBulkItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationApi.AddItemsToPublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddItemsToPublication`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationApi.AddItemsToPublication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** | Publication ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddItemsToPublicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTPublicationBulkItemParams** | [**BTPublicationBulkItemParams**](BTPublicationBulkItemParams.md) |  | 

### Return type

[**BTPublicationInfo**](BTPublicationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePublication

> BTPublicationInfo CreatePublication(ctx).BTPublicationParams(bTPublicationParams).Execute()

Create a new publication.

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
    bTPublicationParams := *openapiclient.NewBTPublicationParams() // BTPublicationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationApi.CreatePublication(context.Background()).BTPublicationParams(bTPublicationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationApi.CreatePublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePublication`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationApi.CreatePublication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTPublicationParams** | [**BTPublicationParams**](BTPublicationParams.md) |  | 

### Return type

[**BTPublicationInfo**](BTPublicationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublication

> map[string]interface{} DeletePublication(ctx, pid).Forever(forever).Execute()

Delete a publication.

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
    pid := "pid_example" // string | Publication ID.
    forever := true // bool | If true, publication is deleted forever. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationApi.DeletePublication(context.Background(), pid).Forever(forever).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationApi.DeletePublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePublication`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublicationApi.DeletePublication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** | Publication ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forever** | **bool** | If true, publication is deleted forever. | [default to false]

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


## DeletePublicationItem

> map[string]interface{} DeletePublicationItem(ctx, pid, iid).Execute()

Remove an item from a publication.

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
    pid := "pid_example" // string | Publication ID.
    iid := "iid_example" // string | Publication item ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationApi.DeletePublicationItem(context.Background(), pid, iid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationApi.DeletePublicationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePublicationItem`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublicationApi.DeletePublicationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** | Publication ID. | 
**iid** | **string** | Publication item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicationItemRequest struct via the builder pattern


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


## GetPublicationItems

> BTPublicationInfo GetPublicationItems(ctx, pid).Execute()

Get all items in a publication.

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
    pid := "pid_example" // string | Publication ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationApi.GetPublicationItems(context.Background(), pid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationApi.GetPublicationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicationItems`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationApi.GetPublicationItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** | Publication ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTPublicationInfo**](BTPublicationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicationAttributes

> map[string]interface{} UpdatePublicationAttributes(ctx, pid).BTPublicationParams(bTPublicationParams).Execute()

Update publication's attributes name, description, and notes.

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
    pid := "pid_example" // string | Publication ID.
    bTPublicationParams := *openapiclient.NewBTPublicationParams() // BTPublicationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicationApi.UpdatePublicationAttributes(context.Background(), pid).BTPublicationParams(bTPublicationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationApi.UpdatePublicationAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePublicationAttributes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublicationApi.UpdatePublicationAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** | Publication ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicationAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTPublicationParams** | [**BTPublicationParams**](BTPublicationParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

