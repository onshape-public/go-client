# \AppDrawingViewApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDrawingViewJsonGeometry**](AppDrawingViewApi.md#GetDrawingViewJsonGeometry) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/views/{viewid}/jsongeometry | Get view geometry of a drawing view in JSON format.
[**GetDrawingViews**](AppDrawingViewApi.md#GetDrawingViews) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/views | Get details of all drawing views.



## GetDrawingViewJsonGeometry

> map[string]interface{} GetDrawingViewJsonGeometry(ctx, viewid, did, wvm, wvmid, eid).TransactionId(transactionId).ChangeId(changeId).Scale(scale).Execute()

Get view geometry of a drawing view in JSON format.

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
    viewid := "viewid_example" // string | 
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional) (default to "")
    changeId := "changeId_example" // string |  (optional) (default to "")
    scale := float64(1.2) // float64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppDrawingViewApi.GetDrawingViewJsonGeometry(context.Background(), viewid, did, wvm, wvmid, eid).TransactionId(transactionId).ChangeId(changeId).Scale(scale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppDrawingViewApi.GetDrawingViewJsonGeometry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrawingViewJsonGeometry`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AppDrawingViewApi.GetDrawingViewJsonGeometry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewid** | **string** |  | 
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrawingViewJsonGeometryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | [default to &quot;&quot;]
 **changeId** | **string** |  | [default to &quot;&quot;]
 **scale** | **float64** |  | 

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


## GetDrawingViews

> BTAppArrayInfoBTAppDrawingViewInfo GetDrawingViews(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()

Get details of all drawing views.

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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wvm := "wvm_example" // string | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    transactionId := "transactionId_example" // string |  (optional) (default to "")
    changeId := "changeId_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppDrawingViewApi.GetDrawingViews(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppDrawingViewApi.GetDrawingViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrawingViews`: BTAppArrayInfoBTAppDrawingViewInfo
    fmt.Fprintf(os.Stdout, "Response from `AppDrawingViewApi.GetDrawingViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrawingViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** |  | [default to &quot;&quot;]
 **changeId** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppArrayInfoBTAppDrawingViewInfo**](BTAppArrayInfoBTAppDrawingViewInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

