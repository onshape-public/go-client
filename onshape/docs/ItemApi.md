# \ItemApi

All URIs are relative to *https://cad.onshape.com/api/v14*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateItem**](ItemApi.md#CreateItem) | **Post** /items | Create a new item.
[**DeleteItem**](ItemApi.md#DeleteItem) | **Delete** /items/{iid} | Delete an item.
[**GetItem**](ItemApi.md#GetItem) | **Get** /items/{iid} | Get item by ID.
[**GetItems**](ItemApi.md#GetItems) | **Get** /items | Get all items owned by a company/classroom/enterprise.
[**UpdateItem**](ItemApi.md#UpdateItem) | **Post** /items/{iid} | Update an item.



## CreateItem

> BTItemInfo CreateItem(ctx).BTItemParams(bTItemParams).Execute()

Create a new item.

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
    bTItemParams := *openapiclient.NewBTItemParams() // BTItemParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ItemApi.CreateItem(context.Background()).BTItemParams(bTItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.CreateItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateItem`: BTItemInfo
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.CreateItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTItemParams** | [**BTItemParams**](BTItemParams.md) |  | 

### Return type

[**BTItemInfo**](BTItemInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteItem

> map[string]interface{} DeleteItem(ctx, iid).Execute()

Delete an item.



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
    iid := "iid_example" // string | ID of the item to delete.

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ItemApi.DeleteItem(context.Background(), iid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.DeleteItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteItem`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.DeleteItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iid** | **string** | ID of the item to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemRequest struct via the builder pattern


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


## GetItem

> BTItemInfo GetItem(ctx, iid).DocumentId(documentId).CompanyId(companyId).Execute()

Get item by ID.



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
    iid := "iid_example" // string | Item ID. 
    documentId := "documentId_example" // string | ID of any document owned by the company/classroom/enterprise. (optional)
    companyId := "companyId_example" // string | Company/classroom/enterprise ID. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ItemApi.GetItem(context.Background(), iid).DocumentId(documentId).CompanyId(companyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.GetItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItem`: BTItemInfo
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.GetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iid** | **string** | Item ID.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentId** | **string** | ID of any document owned by the company/classroom/enterprise. | 
 **companyId** | **string** | Company/classroom/enterprise ID. | 

### Return type

[**BTItemInfo**](BTItemInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> BTListResponseBTItemInfo GetItems(ctx).DocumentId(documentId).CompanyId(companyId).Q(q).PublishStates(publishStates).Classification(classification).Offset(offset).Limit(limit).Execute()

Get all items owned by a company/classroom/enterprise.



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
    documentId := "documentId_example" // string | ID of any document owned by the company/classroom/enterprise. (optional)
    companyId := "companyId_example" // string | Company/classroom/enterprise ID. (optional)
    q := "q_example" // string | Optional search string. (optional) (default to "")
    publishStates := "publishStates_example" // string | Refine returned items by publish states: `0: PENDING | 1: ACTIVE | 2: INACTIVE` (optional) (default to "")
    classification := "classification_example" // string | Refine returned items by classification. Classifications are managed in company/classroom/enterprise [Properties settings](https://cad.onshape.com/help/Content/Plans/items.htm#item-class). (optional) (default to "")
    offset := int32(56) // int32 | Number of entries to skip in the returned list. (optional) (default to 0)
    limit := int32(56) // int32 | The number of list entries to return in a single API call. (optional) (default to 30)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ItemApi.GetItems(context.Background()).DocumentId(documentId).CompanyId(companyId).Q(q).PublishStates(publishStates).Classification(classification).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.GetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItems`: BTListResponseBTItemInfo
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.GetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentId** | **string** | ID of any document owned by the company/classroom/enterprise. | 
 **companyId** | **string** | Company/classroom/enterprise ID. | 
 **q** | **string** | Optional search string. | [default to &quot;&quot;]
 **publishStates** | **string** | Refine returned items by publish states: &#x60;0: PENDING | 1: ACTIVE | 2: INACTIVE&#x60; | [default to &quot;&quot;]
 **classification** | **string** | Refine returned items by classification. Classifications are managed in company/classroom/enterprise [Properties settings](https://cad.onshape.com/help/Content/Plans/items.htm#item-class). | [default to &quot;&quot;]
 **offset** | **int32** | Number of entries to skip in the returned list. | [default to 0]
 **limit** | **int32** | The number of list entries to return in a single API call. | [default to 30]

### Return type

[**BTListResponseBTItemInfo**](BTListResponseBTItemInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItem

> BTItemInfo UpdateItem(ctx, iid).BTItemParams(bTItemParams).Execute()

Update an item.

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
    iid := "iid_example" // string | ID of the item to update.
    bTItemParams := *openapiclient.NewBTItemParams() // BTItemParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ItemApi.UpdateItem(context.Background(), iid).BTItemParams(bTItemParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ItemApi.UpdateItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateItem`: BTItemInfo
    fmt.Fprintf(os.Stdout, "Response from `ItemApi.UpdateItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iid** | **string** | ID of the item to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTItemParams** | [**BTItemParams**](BTItemParams.md) |  | 

### Return type

[**BTItemInfo**](BTItemInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

