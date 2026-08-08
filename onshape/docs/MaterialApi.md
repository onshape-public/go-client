# \MaterialApi

All URIs are relative to *https://cad.onshape.com/api/v16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLibrary**](MaterialApi.md#GetLibrary) | **Get** /materials/libraries/d/{did}/{wvm}/{wvmid}/e/{eid} | Get a list of materials from a material library.
[**UpdateLibrary**](MaterialApi.md#UpdateLibrary) | **Post** /materials/libraries/d/{did}/w/{wid}/e/{eid}/csv | Updates an existing material library



## GetLibrary

> BTMaterialLibraryInfo GetLibrary(ctx, did, wvm, wvmid, eid).Execute()

Get a list of materials from a material library.



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
    did := "did_example" // string | Document ID.
    wvm := "wvm_example" // string | One of w or v or m corresponding to whether a workspace or version or microversion was entered.
    wvmid := "wvmid_example" // string | Workspace (w), Version (v) or Microversion (m) ID.
    eid := "eid_example" // string | Element ID.

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MaterialApi.GetLibrary(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterialApi.GetLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLibrary`: BTMaterialLibraryInfo
    fmt.Fprintf(os.Stdout, "Response from `MaterialApi.GetLibrary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTMaterialLibraryInfo**](BTMaterialLibraryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLibrary

> map[string]interface{} UpdateLibrary(ctx, did, wid, eid).File(file).Execute()

Updates an existing material library



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
    did := "did_example" // string | Document ID.
    wid := "wid_example" // string | Workspace ID.
    eid := "eid_example" // string | Element ID.
    file := os.NewFile(1234, "some_file") // HttpFile | The file to upload. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MaterialApi.UpdateLibrary(context.Background(), did, wid, eid).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaterialApi.UpdateLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLibrary`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MaterialApi.UpdateLibrary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wid** | **string** | Workspace ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **file** | **HttpFile** | The file to upload. | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

