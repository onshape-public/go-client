# \APIApplicationApi

All URIs are relative to *https://cad.onshape.com/api/v11*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppSettings**](APIApplicationApi.md#DeleteAppSettings) | **Delete** /applications/clients/{cid}/settings/users/{uid} | Delete a user&#39;s application preference settings.
[**DeleteCompanyAppSettings**](APIApplicationApi.md#DeleteCompanyAppSettings) | **Delete** /applications/clients/{cid}/settings/companies/{cpid} | Delete a company&#39;s application preference settings.
[**GetApplicableExtensionsForClient**](APIApplicationApi.md#GetApplicableExtensionsForClient) | **Get** /applications/extensions/user/{uid}/client/{cid} | Get a list of the client extensions the specified user has granted/accepted terms for.
[**GetCompanyAppSettings**](APIApplicationApi.md#GetCompanyAppSettings) | **Get** /applications/clients/{cid}/settings/companies/{cpid} | Get company-level preference settings for an application.
[**GetUserAppSettings**](APIApplicationApi.md#GetUserAppSettings) | **Get** /applications/clients/{cid}/settings/users/{uid} | Get user-level preference settings for an application.
[**UpdateAppCompanySettings**](APIApplicationApi.md#UpdateAppCompanySettings) | **Post** /applications/clients/{cid}/settings/companies/{cpid} | Update company preference settings for an application.
[**UpdateAppSettings**](APIApplicationApi.md#UpdateAppSettings) | **Post** /applications/clients/{cid}/settings/users/{uid} | Update a user&#39;s application preference settings.



## DeleteAppSettings

> DeleteAppSettings(ctx, uid, cid).Key(key).Execute()

Delete a user's application preference settings.



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
    uid := "uid_example" // string | 
    cid := "cid_example" // string | 
    key := []string{"Inner_example"} // []string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.APIApplicationApi.DeleteAppSettings(context.Background(), uid, cid).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIApplicationApi.DeleteAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCompanyAppSettings

> map[string]interface{} DeleteCompanyAppSettings(ctx, cpid, cid).Key(key).Execute()

Delete a company's application preference settings.



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
    cpid := "cpid_example" // string | 
    cid := "cid_example" // string | 
    key := []string{"Inner_example"} // []string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.APIApplicationApi.DeleteCompanyAppSettings(context.Background(), cpid, cid).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIApplicationApi.DeleteCompanyAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCompanyAppSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `APIApplicationApi.DeleteCompanyAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cpid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompanyAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **[]string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableExtensionsForClient

> []BTAPIApplicationExtensionInfo GetApplicableExtensionsForClient(ctx, uid, cid).ValidPurchases(validPurchases).Execute()

Get a list of the client extensions the specified user has granted/accepted terms for.

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
    uid := "uid_example" // string | 
    cid := "cid_example" // string | 
    validPurchases := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.APIApplicationApi.GetApplicableExtensionsForClient(context.Background(), uid, cid).ValidPurchases(validPurchases).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIApplicationApi.GetApplicableExtensionsForClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicableExtensionsForClient`: []BTAPIApplicationExtensionInfo
    fmt.Fprintf(os.Stdout, "Response from `APIApplicationApi.GetApplicableExtensionsForClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableExtensionsForClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validPurchases** | **bool** |  | [default to false]

### Return type

[**[]BTAPIApplicationExtensionInfo**](BTAPIApplicationExtensionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyAppSettings

> BTUserAppSettingsInfo GetCompanyAppSettings(ctx, cpid, cid).DocumentId(documentId).Key(key).Execute()

Get company-level preference settings for an application.



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
    cpid := "cpid_example" // string | 
    cid := "cid_example" // string | 
    documentId := "documentId_example" // string | A document owned by the company. Read access to this document allows read access to its owning company's settings. (optional)
    key := []string{"Inner_example"} // []string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.APIApplicationApi.GetCompanyAppSettings(context.Background(), cpid, cid).DocumentId(documentId).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIApplicationApi.GetCompanyAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyAppSettings`: BTUserAppSettingsInfo
    fmt.Fprintf(os.Stdout, "Response from `APIApplicationApi.GetCompanyAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cpid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **documentId** | **string** | A document owned by the company. Read access to this document allows read access to its owning company&#39;s settings. | 
 **key** | **[]string** |  | 

### Return type

[**BTUserAppSettingsInfo**](BTUserAppSettingsInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAppSettings

> BTUserAppSettingsInfo GetUserAppSettings(ctx, uid, cid).Key(key).Execute()

Get user-level preference settings for an application.



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
    uid := "uid_example" // string | 
    cid := "cid_example" // string | 
    key := []string{"Inner_example"} // []string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.APIApplicationApi.GetUserAppSettings(context.Background(), uid, cid).Key(key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIApplicationApi.GetUserAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAppSettings`: BTUserAppSettingsInfo
    fmt.Fprintf(os.Stdout, "Response from `APIApplicationApi.GetUserAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **key** | **[]string** |  | 

### Return type

[**BTUserAppSettingsInfo**](BTUserAppSettingsInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppCompanySettings

> map[string]interface{} UpdateAppCompanySettings(ctx, cpid, cid).BTUserAppSettingsParams(bTUserAppSettingsParams).Execute()

Update company preference settings for an application.



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
    cpid := "cpid_example" // string | 
    cid := "cid_example" // string | 
    bTUserAppSettingsParams := *openapiclient.NewBTUserAppSettingsParams() // BTUserAppSettingsParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.APIApplicationApi.UpdateAppCompanySettings(context.Background(), cpid, cid).BTUserAppSettingsParams(bTUserAppSettingsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIApplicationApi.UpdateAppCompanySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppCompanySettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `APIApplicationApi.UpdateAppCompanySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cpid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppCompanySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTUserAppSettingsParams** | [**BTUserAppSettingsParams**](BTUserAppSettingsParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppSettings

> map[string]interface{} UpdateAppSettings(ctx, uid, cid).BTUserAppSettingsParams(bTUserAppSettingsParams).Execute()

Update a user's application preference settings.



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
    uid := "uid_example" // string | 
    cid := "cid_example" // string | 
    bTUserAppSettingsParams := *openapiclient.NewBTUserAppSettingsParams() // BTUserAppSettingsParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.APIApplicationApi.UpdateAppSettings(context.Background(), uid, cid).BTUserAppSettingsParams(bTUserAppSettingsParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIApplicationApi.UpdateAppSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAppSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `APIApplicationApi.UpdateAppSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTUserAppSettingsParams** | [**BTUserAppSettingsParams**](BTUserAppSettingsParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

