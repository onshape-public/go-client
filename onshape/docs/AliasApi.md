# \AliasApi

All URIs are relative to *https://staging.dev.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AliasApi.md#CreateAlias) | **Post** /aliases | Create an alias that contains users and/or teams.
[**DeleteAlias**](AliasApi.md#DeleteAlias) | **Delete** /aliases/{aid} | Delete alias by alias ID.
[**GetAlias**](AliasApi.md#GetAlias) | **Get** /aliases/{aid} | Retrieve an alias by alias ID.
[**GetAliasMembers**](AliasApi.md#GetAliasMembers) | **Get** /aliases/{aid}/members | Retrieve all alias members by alias ID.
[**GetAliasesInCompany**](AliasApi.md#GetAliasesInCompany) | **Get** /aliases | Retrieve an array of aliases for the enterprise.
[**UpdateAlias**](AliasApi.md#UpdateAlias) | **Post** /aliases/{aid} | Update alias by alias ID.



## CreateAlias

> BTAliasInfo CreateAlias(ctx).BTAliasParams(bTAliasParams).Execute()

Create an alias that contains users and/or teams.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasApi.CreateAlias(context.Background()).BTAliasParams(bTAliasParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasApi.CreateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlias`: BTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasApi.CreateAlias`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlias

> map[string]interface{} DeleteAlias(ctx, aid).Execute()

Delete alias by alias ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasApi.DeleteAlias(context.Background(), aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasApi.DeleteAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlias`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AliasApi.DeleteAlias`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlias

> BTAliasInfo GetAlias(ctx, aid).Execute()

Retrieve an alias by alias ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasApi.GetAlias(context.Background(), aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasApi.GetAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlias`: BTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasApi.GetAlias`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAliasMembers

> BTListResponseBTAliasEntryInfo GetAliasMembers(ctx, aid).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()

Retrieve all alias members by alias ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasApi.GetAliasMembers(context.Background(), aid).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasApi.GetAliasMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAliasMembers`: BTListResponseBTAliasEntryInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasApi.GetAliasMembers`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAliasesInCompany

> BTListResponseBTAliasInfo GetAliasesInCompany(ctx).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()

Retrieve an array of aliases for the enterprise.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasApi.GetAliasesInCompany(context.Background()).Prefix(prefix).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasApi.GetAliasesInCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAliasesInCompany`: BTListResponseBTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasApi.GetAliasesInCompany`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlias

> BTAliasInfo UpdateAlias(ctx, aid).BTAliasParams(bTAliasParams).Execute()

Update alias by alias ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AliasApi.UpdateAlias(context.Background(), aid).BTAliasParams(bTAliasParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AliasApi.UpdateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlias`: BTAliasInfo
    fmt.Fprintf(os.Stdout, "Response from `AliasApi.UpdateAlias`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

