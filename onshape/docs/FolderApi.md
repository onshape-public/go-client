# \FolderApi

All URIs are relative to *https://cad.onshape.com/api/v9*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFolderAcl**](FolderApi.md#GetFolderAcl) | **Get** /folders/{fid}/acl | Get the Access Control List (ACL) for a folder to view permissions.
[**Share**](FolderApi.md#Share) | **Post** /folders/{fid}/share | Share folder with an entity.
[**UnShare**](FolderApi.md#UnShare) | **Delete** /folders/{fid}/share/{eid} | Remove permissions from the folder for the specified Access Control List (ACL) entry.



## GetFolderAcl

> BTAclInfo GetFolderAcl(ctx, fid).Execute()

Get the Access Control List (ACL) for a folder to view permissions.



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
    fid := "fid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.GetFolderAcl(context.Background(), fid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.GetFolderAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolderAcl`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.GetFolderAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTAclInfo**](BTAclInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Share

> BTAclInfo Share(ctx, fid).BTShareParams(bTShareParams).Execute()

Share folder with an entity.



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
    fid := "fid_example" // string | 
    bTShareParams := *openapiclient.NewBTShareParams() // BTShareParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.Share(context.Background(), fid).BTShareParams(bTShareParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.Share``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Share`: BTAclInfo
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.Share`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTShareParams** | [**BTShareParams**](BTShareParams.md) |  | 

### Return type

[**BTAclInfo**](BTAclInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnShare

> map[string]interface{} UnShare(ctx, fid, eid).EntryType(entryType).Execute()

Remove permissions from the folder for the specified Access Control List (ACL) entry.



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
    fid := "fid_example" // string | 
    eid := "eid_example" // string | 
    entryType := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FolderApi.UnShare(context.Background(), fid, eid).EntryType(entryType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderApi.UnShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnShare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FolderApi.UnShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **entryType** | **int32** |  | [default to 0]

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

