# \APIApplicationApi

All URIs are relative to *https://staging.dev.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAppSettings**](APIApplicationApi.md#DeleteAppSettings) | **Delete** /applications/clients/{cid}/settings/users/{uid} | Delete application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.
[**DeleteCompanyAppSettings**](APIApplicationApi.md#DeleteCompanyAppSettings) | **Delete** /applications/clients/{cid}/settings/companies/{cpid} | Delete company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.
[**GetApplicableExtensionsForClient**](APIApplicationApi.md#GetApplicableExtensionsForClient) | **Get** /applications/extensions/user/{uid}/client/{cid} | 
[**GetCompanyAppSettings**](APIApplicationApi.md#GetCompanyAppSettings) | **Get** /applications/clients/{cid}/settings/companies/{cpid} | Retrieve company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.
[**GetUserAppSettings**](APIApplicationApi.md#GetUserAppSettings) | **Get** /applications/clients/{cid}/settings/users/{uid} | Retrieve application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.
[**UpdateAppCompanySettings**](APIApplicationApi.md#UpdateAppCompanySettings) | **Post** /applications/clients/{cid}/settings/companies/{cpid} | Update or create company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.
[**UpdateAppSettings**](APIApplicationApi.md#UpdateAppSettings) | **Post** /applications/clients/{cid}/settings/users/{uid} | Update or create application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.



## DeleteAppSettings

> DeleteAppSettings(ctx, uid, cid).Key(key).Execute()

Delete application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

Delete company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyAppSettings

> BTUserAppSettingsInfo GetCompanyAppSettings(ctx, cpid, cid).Key(key).Execute()

Retrieve company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIApplicationApi.GetCompanyAppSettings(context.Background(), cpid, cid).Key(key).Execute()
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

Retrieve application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

Update or create company level settings for this application by client ID and company ID. This API may only be used with an OAuth token and only by the current user.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

Update or create application settings for a user by client ID and user ID. This API may only be used with an OAuth token and only by the current user.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

