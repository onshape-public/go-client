# \ThumbnailApi

All URIs are relative to *https://cad.onshape.com/api/v9*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationThumbnails**](ThumbnailApi.md#DeleteApplicationThumbnails) | **Delete** /thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | Delete an element&#39;s thumbnail.
[**GetDocumentThumbnail**](ThumbnailApi.md#GetDocumentThumbnail) | **Get** /thumbnails/d/{did}/w/{wid} | Get the thumbnail info for a workspace.
[**GetDocumentThumbnailWithSize**](ThumbnailApi.md#GetDocumentThumbnailWithSize) | **Get** /thumbnails/d/{did}/w/{wid}/s/{sz} | Get the thumbnail image with the given size for a document.
[**GetElementThumbnail**](ThumbnailApi.md#GetElementThumbnail) | **Get** /thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | Get the thumbnail info structure for an element.
[**GetElementThumbnailWithApiConfiguration**](ThumbnailApi.md#GetElementThumbnailWithApiConfiguration) | **Get** /thumbnails/d/{did}/w/{wid}/e/{eid}/ac/{cid}/s/{sz} | Get the thumbnail image with the given configuration for an element.
[**GetElementThumbnailWithSize**](ThumbnailApi.md#GetElementThumbnailWithSize) | **Get** /thumbnails/d/{did}/{wv}/{wvid}/e/{eid}/s/{sz} | Get the thumbnail image with the given size for an element.
[**GetThumbnailForDocument**](ThumbnailApi.md#GetThumbnailForDocument) | **Get** /thumbnails/d/{did} | Get the thumbnail info for a document in the default workspace.
[**GetThumbnailForDocumentAndVersion**](ThumbnailApi.md#GetThumbnailForDocumentAndVersion) | **Get** /thumbnails/d/{did}/v/{vid} | Get the thumbnail info for a version of a document.
[**GetThumbnailForDocumentAndVersionOld**](ThumbnailApi.md#GetThumbnailForDocumentAndVersionOld) | **Get** /thumbnails/document/{did}/version/{vid} | This endpoint will be deprecated soon. Use &#x60;getThumbnailForDocumentAndVersion&#x60; instead.
[**GetThumbnailForDocumentOld**](ThumbnailApi.md#GetThumbnailForDocumentOld) | **Get** /thumbnails/document/{did} | This endpoint will be deprecated soon. Use &#x60;getThumbnailForDocument&#x60; instead.
[**SetApplicationElementThumbnail**](ThumbnailApi.md#SetApplicationElementThumbnail) | **Post** /thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | Set the thumbnail image for an application element.



## DeleteApplicationThumbnails

> map[string]interface{} DeleteApplicationThumbnails(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()

Delete an element's thumbnail.



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
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.DeleteApplicationThumbnails(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.DeleteApplicationThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApplicationThumbnails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.DeleteApplicationThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

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


## GetDocumentThumbnail

> BTThumbnailInfo GetDocumentThumbnail(ctx, did, wid).Execute()

Get the thumbnail info for a workspace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetDocumentThumbnail(context.Background(), did, wid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetDocumentThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentThumbnail`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetDocumentThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentThumbnailWithSize

> map[string]interface{} GetDocumentThumbnailWithSize(ctx, did, wid, sz).T(t).SkipDefaultImage(skipDefaultImage).Execute()

Get the thumbnail image with the given size for a document.



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
    sz := "300x300" // string | the generated thumbnail size in pixels, widthxheigth
    t := "t_example" // string | Cache Control key. If specified, the response header returned will tell the client to use cached thumbnails. (optional)
    skipDefaultImage := "skipDefaultImage_example" // string | Controls the return of the default image, if thumbnail is not available (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetDocumentThumbnailWithSize(context.Background(), did, wid, sz).T(t).SkipDefaultImage(skipDefaultImage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetDocumentThumbnailWithSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentThumbnailWithSize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetDocumentThumbnailWithSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**sz** | **string** | the generated thumbnail size in pixels, widthxheigth | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **t** | **string** | Cache Control key. If specified, the response header returned will tell the client to use cached thumbnails. | 
 **skipDefaultImage** | **string** | Controls the return of the default image, if thumbnail is not available | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnail

> BTThumbnailInfo GetElementThumbnail(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()

Get the thumbnail info structure for an element.



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
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetElementThumbnail(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetElementThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementThumbnail`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetElementThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnailWithApiConfiguration

> map[string]interface{} GetElementThumbnailWithApiConfiguration(ctx, did, wid, eid, cid, sz).LinkDocumentId(linkDocumentId).T(t).SkipDefaultImage(skipDefaultImage).RejectEmpty(rejectEmpty).RequireConfigMatch(requireConfigMatch).Execute()

Get the thumbnail image with the given configuration for an element.



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
    cid := "cid_example" // string | 
    sz := "300x300" // string | the generated thumbnail size in pixels, widthxheigth
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    t := "t_example" // string | Cache Control key. If specified, the response header returned will tell the client to use cached thumbnails. (optional)
    skipDefaultImage := "skipDefaultImage_example" // string | Controls the return of the default image, if thumbnail is not available (optional) (default to "")
    rejectEmpty := true // bool |  (optional) (default to false)
    requireConfigMatch := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetElementThumbnailWithApiConfiguration(context.Background(), did, wid, eid, cid, sz).LinkDocumentId(linkDocumentId).T(t).SkipDefaultImage(skipDefaultImage).RejectEmpty(rejectEmpty).RequireConfigMatch(requireConfigMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetElementThumbnailWithApiConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementThumbnailWithApiConfiguration`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetElementThumbnailWithApiConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** | The id of the element in which to perform the operation. | 
**cid** | **string** |  | 
**sz** | **string** | the generated thumbnail size in pixels, widthxheigth | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailWithApiConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **t** | **string** | Cache Control key. If specified, the response header returned will tell the client to use cached thumbnails. | 
 **skipDefaultImage** | **string** | Controls the return of the default image, if thumbnail is not available | [default to &quot;&quot;]
 **rejectEmpty** | **bool** |  | [default to false]
 **requireConfigMatch** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnailWithSize

> map[string]interface{} GetElementThumbnailWithSize(ctx, did, wv, wvid, eid, sz).LinkDocumentId(linkDocumentId).T(t).SkipDefaultImage(skipDefaultImage).RejectEmpty(rejectEmpty).Execute()

Get the thumbnail image with the given size for an element.

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
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    sz := "300x300" // string | the generated thumbnail size in pixels, widthxheigth
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    t := "t_example" // string | Cache Control key. If specified, the response header returned will tell the client to use cached thumbnails. (optional)
    skipDefaultImage := "skipDefaultImage_example" // string | Controls the return of the default image, if thumbnail is not available (optional) (default to "")
    rejectEmpty := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetElementThumbnailWithSize(context.Background(), did, wv, wvid, eid, sz).LinkDocumentId(linkDocumentId).T(t).SkipDefaultImage(skipDefaultImage).RejectEmpty(rejectEmpty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetElementThumbnailWithSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetElementThumbnailWithSize`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetElementThumbnailWithSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 
**sz** | **string** | the generated thumbnail size in pixels, widthxheigth | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **t** | **string** | Cache Control key. If specified, the response header returned will tell the client to use cached thumbnails. | 
 **skipDefaultImage** | **string** | Controls the return of the default image, if thumbnail is not available | [default to &quot;&quot;]
 **rejectEmpty** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocument

> BTThumbnailInfo GetThumbnailForDocument(ctx, did).Execute()

Get the thumbnail info for a document in the default workspace.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetThumbnailForDocument(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetThumbnailForDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThumbnailForDocument`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetThumbnailForDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocumentAndVersion

> BTThumbnailInfo GetThumbnailForDocumentAndVersion(ctx, did, vid).LinkDocumentId(linkDocumentId).Execute()

Get the thumbnail info for a version of a document.



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
    vid := "vid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetThumbnailForDocumentAndVersion(context.Background(), did, vid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetThumbnailForDocumentAndVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThumbnailForDocumentAndVersion`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetThumbnailForDocumentAndVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentAndVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkDocumentId** | **string** |  | 

### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocumentAndVersionOld

> BTThumbnailInfo GetThumbnailForDocumentAndVersionOld(ctx, did, vid).Execute()

This endpoint will be deprecated soon. Use `getThumbnailForDocumentAndVersion` instead.



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
    vid := "vid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetThumbnailForDocumentAndVersionOld(context.Background(), did, vid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetThumbnailForDocumentAndVersionOld``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThumbnailForDocumentAndVersionOld`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetThumbnailForDocumentAndVersionOld`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentAndVersionOldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocumentOld

> BTThumbnailInfo GetThumbnailForDocumentOld(ctx, did).Execute()

This endpoint will be deprecated soon. Use `getThumbnailForDocument` instead.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.GetThumbnailForDocumentOld(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.GetThumbnailForDocumentOld``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetThumbnailForDocumentOld`: BTThumbnailInfo
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.GetThumbnailForDocumentOld`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentOldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApplicationElementThumbnail

> map[string]interface{} SetApplicationElementThumbnail(ctx, did, wv, wvid, eid).BTApplicationElementThumbnailParamsArray(bTApplicationElementThumbnailParamsArray).LinkDocumentId(linkDocumentId).Overwrite(overwrite).Execute()

Set the thumbnail image for an application element.



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
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    bTApplicationElementThumbnailParamsArray := *openapiclient.NewBTApplicationElementThumbnailParamsArray() // BTApplicationElementThumbnailParamsArray | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    overwrite := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThumbnailApi.SetApplicationElementThumbnail(context.Background(), did, wv, wvid, eid).BTApplicationElementThumbnailParamsArray(bTApplicationElementThumbnailParamsArray).LinkDocumentId(linkDocumentId).Overwrite(overwrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailApi.SetApplicationElementThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetApplicationElementThumbnail`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ThumbnailApi.SetApplicationElementThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApplicationElementThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTApplicationElementThumbnailParamsArray** | [**BTApplicationElementThumbnailParamsArray**](BTApplicationElementThumbnailParamsArray.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **overwrite** | **bool** |  | [default to false]

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

