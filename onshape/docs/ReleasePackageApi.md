# \ReleasePackageApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObsoletionPackage**](ReleasePackageApi.md#CreateObsoletionPackage) | **Post** /releasepackages/obsoletion/{wfid} | Create an obsoletion package to make an existing revision obsolete.
[**CreateReleasePackage**](ReleasePackageApi.md#CreateReleasePackage) | **Post** /releasepackages/release/{wfid} | Create a new release package for one or more items.
[**GetReleasePackage**](ReleasePackageApi.md#GetReleasePackage) | **Get** /releasepackages/{rpid} | Get details about the specified release package.
[**UpdateReleasePackage**](ReleasePackageApi.md#UpdateReleasePackage) | **Post** /releasepackages/{rpid} | Update the release/obsoletion package/item properties.



## CreateObsoletionPackage

> map[string]interface{} CreateObsoletionPackage(ctx, wfid).RevisionId(revisionId).DebugMode(debugMode).Execute()

Create an obsoletion package to make an existing revision obsolete.

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
    wfid := "wfid_example" // string | 
    revisionId := "revisionId_example" // string | 
    debugMode := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.CreateObsoletionPackage(context.Background(), wfid).RevisionId(revisionId).DebugMode(debugMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.CreateObsoletionPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObsoletionPackage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.CreateObsoletionPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObsoletionPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revisionId** | **string** |  | 
 **debugMode** | **bool** |  | [default to false]

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


## CreateReleasePackage

> map[string]interface{} CreateReleasePackage(ctx, wfid).BTReleasePackageParams(bTReleasePackageParams).DebugMode(debugMode).Execute()

Create a new release package for one or more items.



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
    wfid := "wfid_example" // string | 
    bTReleasePackageParams := *openapiclient.NewBTReleasePackageParams() // BTReleasePackageParams | 
    debugMode := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.CreateReleasePackage(context.Background(), wfid).BTReleasePackageParams(bTReleasePackageParams).DebugMode(debugMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.CreateReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReleasePackage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.CreateReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTReleasePackageParams** | [**BTReleasePackageParams**](BTReleasePackageParams.md) |  | 
 **debugMode** | **bool** |  | [default to false]

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


## GetReleasePackage

> BTReleasePackageInfo GetReleasePackage(ctx, rpid).Detailed(detailed).Execute()

Get details about the specified release package.

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
    rpid := "rpid_example" // string | 
    detailed := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.GetReleasePackage(context.Background(), rpid).Detailed(detailed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.GetReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleasePackage`: BTReleasePackageInfo
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.GetReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detailed** | **bool** |  | [default to false]

### Return type

[**BTReleasePackageInfo**](BTReleasePackageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleasePackage

> BTReleasePackageInfo UpdateReleasePackage(ctx, rpid).BTUpdateReleasePackageParams(bTUpdateReleasePackageParams).Action(action).Wfaction(wfaction).Execute()

Update the release/obsoletion package/item properties.



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
    rpid := "rpid_example" // string | 
    bTUpdateReleasePackageParams := *openapiclient.NewBTUpdateReleasePackageParams() // BTUpdateReleasePackageParams | 
    action := "action_example" // string |  (optional) (default to "UPDATE")
    wfaction := "wfaction_example" // string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.UpdateReleasePackage(context.Background(), rpid).BTUpdateReleasePackageParams(bTUpdateReleasePackageParams).Action(action).Wfaction(wfaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.UpdateReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReleasePackage`: BTReleasePackageInfo
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.UpdateReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTUpdateReleasePackageParams** | [**BTUpdateReleasePackageParams**](BTUpdateReleasePackageParams.md) |  | 
 **action** | **string** |  | [default to &quot;UPDATE&quot;]
 **wfaction** | **string** |  | 

### Return type

[**BTReleasePackageInfo**](BTReleasePackageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

