# \PublicationAPI

All URIs are relative to *https://cad.onshape.com/api/v9*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItemToPublication**](PublicationAPI.md#AddItemToPublication) | **Post** /publications/{pid}/item | Add an item in a publication.
[**AddItemsToPublication**](PublicationAPI.md#AddItemsToPublication) | **Post** /publications/{pid}/items | Add publication items in bulk.
[**CreatePublication**](PublicationAPI.md#CreatePublication) | **Post** /publications | Create a new publication.
[**DeletePublication**](PublicationAPI.md#DeletePublication) | **Delete** /publications/{pid} | Delete a publication.
[**DeletePublicationItem**](PublicationAPI.md#DeletePublicationItem) | **Delete** /publications/{pid}/item/{iid} | Remove an item from a publication.
[**GetPublicationItems**](PublicationAPI.md#GetPublicationItems) | **Get** /publications/{pid}/items | Get all items in a publication.
[**UpdatePublicationAttributes**](PublicationAPI.md#UpdatePublicationAttributes) | **Post** /publications/{pid} | Update publication&#39;s attributes name, description, and notes.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PublicationAPI.AddItemToPublication(context.Background(), pid).BTPublicationItemParams(bTPublicationItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationAPI.AddItemToPublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddItemToPublication`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationAPI.AddItemToPublication`: %v\n", resp)
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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PublicationAPI.AddItemsToPublication(context.Background(), pid).BTPublicationBulkItemParams(bTPublicationBulkItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationAPI.AddItemsToPublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddItemsToPublication`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationAPI.AddItemsToPublication`: %v\n", resp)
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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PublicationAPI.CreatePublication(context.Background()).BTPublicationParams(bTPublicationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationAPI.CreatePublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePublication`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationAPI.CreatePublication`: %v\n", resp)
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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PublicationAPI.DeletePublication(context.Background(), pid).Forever(forever).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationAPI.DeletePublication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePublication`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublicationAPI.DeletePublication`: %v\n", resp)
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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PublicationAPI.DeletePublicationItem(context.Background(), pid, iid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationAPI.DeletePublicationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePublicationItem`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublicationAPI.DeletePublicationItem`: %v\n", resp)
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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PublicationAPI.GetPublicationItems(context.Background(), pid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationAPI.GetPublicationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicationItems`: BTPublicationInfo
    fmt.Fprintf(os.Stdout, "Response from `PublicationAPI.GetPublicationItems`: %v\n", resp)
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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PublicationAPI.UpdatePublicationAttributes(context.Background(), pid).BTPublicationParams(bTPublicationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicationAPI.UpdatePublicationAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePublicationAttributes`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PublicationAPI.UpdatePublicationAttributes`: %v\n", resp)
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

