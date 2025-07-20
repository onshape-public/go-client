# \DrawingApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDrawingAppElement**](DrawingApi.md#CreateDrawingAppElement) | **Post** /drawings/d/{did}/w/{wid}/create | Create a new drawing in a document.
[**CreateDrawingTranslation**](DrawingApi.md#CreateDrawingTranslation) | **Post** /drawings/d/{did}/{wv}/{wvid}/e/{eid}/translations | Translate (export) a drawing to a different format.
[**GetDrawingTranslatorFormats**](DrawingApi.md#GetDrawingTranslatorFormats) | **Get** /drawings/d/{did}/w/{wid}/e/{eid}/translationformats | Get a list of all valid formats the drawing can be translated (exported) to.
[**GetDrawingViewJsonGeometry1**](DrawingApi.md#GetDrawingViewJsonGeometry1) | **Get** /drawings/d/{did}/{wvm}/{wvmid}/e/{eid}/views/{viewid}/jsongeometry | Get view geometry of a drawing view in JSON format.
[**GetDrawingViews1**](DrawingApi.md#GetDrawingViews1) | **Get** /drawings/d/{did}/{wvm}/{wvmid}/e/{eid}/views | Get details of all drawing views.
[**GetModificationStatus**](DrawingApi.md#GetModificationStatus) | **Get** /drawings/modify/status/{mrid} | Get the status of a drawing modification operation.
[**ModifyDrawing**](DrawingApi.md#ModifyDrawing) | **Post** /drawings/d/{did}/w/{wid}/e/{eid}/modify | Modify a drawing via JSON payload.



## CreateDrawingAppElement

> BTDocumentElementInfo CreateDrawingAppElement(ctx, did, wid).BTDrawingParams(bTDrawingParams).Execute()

Create a new drawing in a document.



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
    did := "did_example" // string | ID of the document in which to create the drawing.
    wid := "wid_example" // string | ID of the workspace in which to create the drawing.
    bTDrawingParams := *openapiclient.NewBTDrawingParams() // BTDrawingParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.DrawingApi.CreateDrawingAppElement(context.Background(), did, wid).BTDrawingParams(bTDrawingParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingApi.CreateDrawingAppElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDrawingAppElement`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingApi.CreateDrawingAppElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | ID of the document in which to create the drawing. | 
**wid** | **string** | ID of the workspace in which to create the drawing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDrawingAppElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTDrawingParams** | [**BTDrawingParams**](BTDrawingParams.md) |  | 

### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDrawingTranslation

> BTTranslationRequestInfo CreateDrawingTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Translate (export) a drawing to a different format.



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
    did := "did_example" // string | 
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | 
    bTTranslateFormatParams := *openapiclient.NewBTTranslateFormatParams("FormatName_example") // BTTranslateFormatParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.DrawingApi.CreateDrawingTranslation(context.Background(), did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingApi.CreateDrawingTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDrawingTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingApi.CreateDrawingTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDrawingTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTTranslateFormatParams** | [**BTTranslateFormatParams**](BTTranslateFormatParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrawingTranslatorFormats

> []BTModelFormatInfo GetDrawingTranslatorFormats(ctx, did, wid, eid).Execute()

Get a list of all valid formats the drawing can be translated (exported) to.



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
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    eid := "eid_example" // string | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.DrawingApi.GetDrawingTranslatorFormats(context.Background(), did, wid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingApi.GetDrawingTranslatorFormats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrawingTranslatorFormats`: []BTModelFormatInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingApi.GetDrawingTranslatorFormats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrawingTranslatorFormatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]BTModelFormatInfo**](BTModelFormatInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrawingViewJsonGeometry1

> map[string]interface{} GetDrawingViewJsonGeometry1(ctx, did, wvm, wvmid, eid, viewid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Scale(scale).Execute()

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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wvm := "wvm_example" // string | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    viewid := "viewid_example" // string | The id of the view in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    transactionId := "transactionId_example" // string | The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API. (optional) (default to "")
    changeId := "changeId_example" // string | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. (optional) (default to "")
    scale := float64(1.2) // float64 | Scale for measurements. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.DrawingApi.GetDrawingViewJsonGeometry1(context.Background(), did, wvm, wvmid, eid, viewid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Scale(scale).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingApi.GetDrawingViewJsonGeometry1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrawingViewJsonGeometry1`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DrawingApi.GetDrawingViewJsonGeometry1`: %v\n", resp)
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
**viewid** | **string** | The id of the view in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrawingViewJsonGeometry1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** | The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API. | [default to &quot;&quot;]
 **changeId** | **string** | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. | [default to &quot;&quot;]
 **scale** | **float64** | Scale for measurements. | 

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


## GetDrawingViews1

> BTAppArrayInfoBTAppDrawingViewInfo GetDrawingViews1(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()

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
    transactionId := "transactionId_example" // string | The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API. (optional) (default to "")
    changeId := "changeId_example" // string | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. (optional) (default to "")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.DrawingApi.GetDrawingViews1(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingApi.GetDrawingViews1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDrawingViews1`: BTAppArrayInfoBTAppDrawingViewInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingApi.GetDrawingViews1`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDrawingViews1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** | The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API. | [default to &quot;&quot;]
 **changeId** | **string** | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. | [default to &quot;&quot;]

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


## GetModificationStatus

> BTAppModificationRequestInfo GetModificationStatus(ctx, mrid).Execute()

Get the status of a drawing modification operation.

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
    mrid := "mrid_example" // string | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.DrawingApi.GetModificationStatus(context.Background(), mrid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingApi.GetModificationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetModificationStatus`: BTAppModificationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingApi.GetModificationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mrid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModificationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTAppModificationRequestInfo**](BTAppModificationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyDrawing

> BTAppModificationRequestInfo ModifyDrawing(ctx, did, wid, eid).BTDrawingModificationParams(bTDrawingModificationParams).LinkDocumentId(linkDocumentId).Execute()

Modify a drawing via JSON payload.



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
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    bTDrawingModificationParams := *openapiclient.NewBTDrawingModificationParams() // BTDrawingModificationParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.DrawingApi.ModifyDrawing(context.Background(), did, wid, eid).BTDrawingModificationParams(bTDrawingModificationParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DrawingApi.ModifyDrawing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyDrawing`: BTAppModificationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `DrawingApi.ModifyDrawing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyDrawingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTDrawingModificationParams** | [**BTDrawingModificationParams**](BTDrawingModificationParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTAppModificationRequestInfo**](BTAppModificationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

