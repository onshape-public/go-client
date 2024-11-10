# \AliasAPI

All URIs are relative to *https://cad.onshape.com/api/v9*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AliasAPI.md#CreateAlias) | **Post** /aliases | Create an alias in your enterprise.
[**DeleteAlias**](AliasAPI.md#DeleteAlias) | **Delete** /aliases/{aid} | Delete an alias from your enterprise.
[**GetAlias**](AliasAPI.md#GetAlias) | **Get** /aliases/{aid} | Get an alias by ID.
[**GetAliasMembers**](AliasAPI.md#GetAliasMembers) | **Get** /aliases/{aid}/members | Get all users and teams assigned to an alias.
[**GetAliasesInCompany**](AliasAPI.md#GetAliasesInCompany) | **Get** /aliases | Get a list of all aliases that exist for your enterprise.
[**UpdateAlias**](AliasAPI.md#UpdateAlias) | **Post** /aliases/{aid} | Add, remove, replace, or rename entries in an alias list.



## CreateAlias

> BTAliasInfo CreateAlias(ctx).BTAliasParams(bTAliasParams).Execute()

Create an alias in your enterprise.



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
    bTAliasParams := *openapiclient.NewBTAliasParams() // BTAliasParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AliasAPI.CreateAlias(context.Background()).BTAliasParams(bTAliasParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.CreateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlias`: BTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasAPI.CreateAlias`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTAliasParams** | [**BTAliasParams**](BTAliasParams.md) |  | 

### Return type

[**BTAliasInfo**](BTAliasInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlias

> map[string]interface{} DeleteAlias(ctx, aid).Execute()

Delete an alias from your enterprise.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AliasAPI.DeleteAlias(context.Background(), aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.DeleteAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlias`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AliasAPI.DeleteAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAliasRequest struct via the builder pattern


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


## GetAlias

> BTAliasInfo GetAlias(ctx, aid).Execute()

Get an alias by ID.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AliasAPI.GetAlias(context.Background(), aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.GetAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlias`: BTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasAPI.GetAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTAliasInfo**](BTAliasInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAliasMembers

> BTListResponseBTAliasEntryInfo GetAliasMembers(ctx, aid).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()

Get all users and teams assigned to an alias.



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
    prefix := "prefix_example" // string |  (optional) (default to "")
    sortColumn := "sortColumn_example" // string |  (optional) (default to "name")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AliasAPI.GetAliasMembers(context.Background(), aid).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.GetAliasMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAliasMembers`: BTListResponseBTAliasEntryInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasAPI.GetAliasMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAliasMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **string** |  | [default to &quot;&quot;]
 **sortColumn** | **string** |  | [default to &quot;name&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTAliasEntryInfo**](BTListResponseBTAliasEntryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAliasesInCompany

> BTListResponseBTAliasInfo GetAliasesInCompany(ctx).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()

Get a list of all aliases that exist for your enterprise.

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
    prefix := "prefix_example" // string |  (optional) (default to "")
    sortColumn := "sortColumn_example" // string |  (optional) (default to "name")
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AliasAPI.GetAliasesInCompany(context.Background()).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.GetAliasesInCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAliasesInCompany`: BTListResponseBTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasAPI.GetAliasesInCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAliasesInCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** |  | [default to &quot;&quot;]
 **sortColumn** | **string** |  | [default to &quot;name&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTAliasInfo**](BTListResponseBTAliasInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlias

> BTAliasInfo UpdateAlias(ctx, aid).BTAliasParams(bTAliasParams).Execute()

Add, remove, replace, or rename entries in an alias list.



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
    bTAliasParams := *openapiclient.NewBTAliasParams() // BTAliasParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AliasAPI.UpdateAlias(context.Background(), aid).BTAliasParams(bTAliasParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasAPI.UpdateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlias`: BTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasAPI.UpdateAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTAliasParams** | [**BTAliasParams**](BTAliasParams.md) |  | 

### Return type

[**BTAliasInfo**](BTAliasInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

