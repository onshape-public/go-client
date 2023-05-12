# \MetadataApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVEOPStandardContentMetadata**](MetadataApi.md#GetVEOPStandardContentMetadata) | **Get** /metadata/standardcontent/d/{did}/v/{vid}/e/{eid}/p/{pid} | Retrieve metadata of a standard content part in a version by document ID, version ID, tab ID, owner ID, and part ID.
[**GetWMVEMetadata**](MetadataApi.md#GetWMVEMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | Retrieve metadata by document ID, workspace or version or microversion ID, and tab ID.
[**GetWMVEPMetadata**](MetadataApi.md#GetWMVEPMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/{iden}/{pid} | Retrieve metadata by document ID, workspace or version or microversion ID, tab ID, and Part ID.
[**GetWMVEPsMetadata**](MetadataApi.md#GetWMVEPsMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p | Retrieve metadata by document ID, workspace or version or microversion ID, and tab ID.
[**GetWMVEsMetadata**](MetadataApi.md#GetWMVEsMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e | Retrieve metadata by document ID and workspace or version or microversion ID.
[**GetWVMetadata**](MetadataApi.md#GetWVMetadata) | **Get** /metadata/d/{did}/{wv}/{wvid} | Retrieve workspace or version metadata by document ID and workspace or version ID.
[**UpdateVEOPStandardContentPartMetadata**](MetadataApi.md#UpdateVEOPStandardContentPartMetadata) | **Post** /metadata/standardcontent/d/{did} | Update metadata of a standard content part in a version by document ID, version ID, tab ID, owner ID, and part ID.
[**UpdateWVEMetadata**](MetadataApi.md#UpdateWVEMetadata) | **Post** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | Update workspace metadata by document ID, workspace or version or microversion ID, and tab ID.
[**UpdateWVEPMetadata**](MetadataApi.md#UpdateWVEPMetadata) | **Post** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/{iden}/{pid} | Update workspace metadata by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**UpdateWVMetadata**](MetadataApi.md#UpdateWVMetadata) | **Post** /metadata/d/{did}/{wv}/{wvid} | Update workspace or version metadata by document ID and workspace or version ID.



## GetVEOPStandardContentMetadata

> BTMetadataObjectInfo GetVEOPStandardContentMetadata(ctx, did, vid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()

Retrieve metadata of a standard content part in a version by document ID, version ID, tab ID, owner ID, and part ID.

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
    eid := "eid_example" // string | 
    pid := "pid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetVEOPStandardContentMetadata(context.Background(), did, vid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetVEOPStandardContentMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVEOPStandardContentMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetVEOPStandardContentMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 
**eid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVEOPStandardContentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **includeComputedProperties** | **bool** |  | [default to true]
 **includeComputedAssemblyProperties** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

### Return type

[**BTMetadataObjectInfo**](BTMetadataObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWMVEMetadata

> BTMetadataObjectInfo GetWMVEMetadata(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()

Retrieve metadata by document ID, workspace or version or microversion ID, and tab ID.

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
    configuration := "configuration_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    depth := "depth_example" // string |  (optional) (default to "1")
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetWMVEMetadata(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetWMVEMetadata`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWMVEMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **includeComputedProperties** | **bool** |  | [default to true]
 **includeComputedAssemblyProperties** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

### Return type

[**BTMetadataObjectInfo**](BTMetadataObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWMVEPMetadata

> BTMetadataObjectInfo GetWMVEPMetadata(ctx, did, wvm, wvmid, eid, iden, pid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).InferMetadataOwner(inferMetadataOwner).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()

Retrieve metadata by document ID, workspace or version or microversion ID, tab ID, and Part ID.

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
    iden := "iden_example" // string | Denotes whether the pid specified is a part id (p) or a part identity (pi).
    pid := "pid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetWMVEPMetadata(context.Background(), did, wvm, wvmid, eid, iden, pid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).InferMetadataOwner(inferMetadataOwner).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEPMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEPMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetWMVEPMetadata`: %v\n", resp)
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
**iden** | **string** | Denotes whether the pid specified is a part id (p) or a part identity (pi). | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEPMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **includeComputedProperties** | **bool** |  | [default to true]
 **includeComputedAssemblyProperties** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

### Return type

[**BTMetadataObjectInfo**](BTMetadataObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWMVEPsMetadata

> BTMetadataObjectListInfoBTMetadataPartInfo GetWMVEPsMetadata(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).InferMetadataOwner(inferMetadataOwner).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()

Retrieve metadata by document ID, workspace or version or microversion ID, and tab ID.

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
    configuration := "configuration_example" // string |  (optional) (default to "")
    inferMetadataOwner := true // bool |  (optional) (default to false)
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetWMVEPsMetadata(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).InferMetadataOwner(inferMetadataOwner).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEPsMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEPsMetadata`: BTMetadataObjectListInfoBTMetadataPartInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetWMVEPsMetadata`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetWMVEPsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **inferMetadataOwner** | **bool** |  | [default to false]
 **includeComputedProperties** | **bool** |  | [default to true]
 **includeComputedAssemblyProperties** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

### Return type

[**BTMetadataObjectListInfoBTMetadataPartInfo**](BTMetadataObjectListInfoBTMetadataPartInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWMVEsMetadata

> BTMetadataObjectListInfoBTMetadataElementInfo GetWMVEsMetadata(ctx, did, wvm, wvmid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()

Retrieve metadata by document ID and workspace or version or microversion ID.

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    depth := "depth_example" // string |  (optional) (default to "1")
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetWMVEsMetadata(context.Background(), did, wvm, wvmid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWMVEsMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEsMetadata`: BTMetadataObjectListInfoBTMetadataElementInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetWMVEsMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **includeComputedProperties** | **bool** |  | [default to true]
 **includeComputedAssemblyProperties** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

### Return type

[**BTMetadataObjectListInfoBTMetadataElementInfo**](BTMetadataObjectListInfoBTMetadataElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWVMetadata

> BTMetadataObjectInfo GetWVMetadata(ctx, did, wv, wvid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()

Retrieve workspace or version metadata by document ID and workspace or version ID.

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
    linkDocumentId := "linkDocumentId_example" // string |  (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    depth := "depth_example" // string |  (optional) (default to "1")
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetWVMetadata(context.Background(), did, wv, wvid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetWVMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWVMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetWVMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWVMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **includeComputedProperties** | **bool** |  | [default to true]
 **includeComputedAssemblyProperties** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

### Return type

[**BTMetadataObjectInfo**](BTMetadataObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVEOPStandardContentPartMetadata

> map[string]interface{} UpdateVEOPStandardContentPartMetadata(ctx, did).LinkDocumentId(linkDocumentId).Body(body).Execute()

Update metadata of a standard content part in a version by document ID, version ID, tab ID, owner ID, and part ID.

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
    linkDocumentId := "linkDocumentId_example" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.UpdateVEOPStandardContentPartMetadata(context.Background(), did).LinkDocumentId(linkDocumentId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateVEOPStandardContentPartMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVEOPStandardContentPartMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.UpdateVEOPStandardContentPartMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVEOPStandardContentPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkDocumentId** | **string** |  | 
 **body** | **string** |  | 

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


## UpdateWVEMetadata

> map[string]interface{} UpdateWVEMetadata(ctx, did, wvm, wvmid, eid).Body(body).Configuration(configuration).Execute()

Update workspace metadata by document ID, workspace or version or microversion ID, and tab ID.

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
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 
    body := "body_example" // string | 
    configuration := "configuration_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.UpdateWVEMetadata(context.Background(), did, wvm, wvmid, eid).Body(body).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateWVEMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWVEMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.UpdateWVEMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVEMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 
 **configuration** | **string** |  | 

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


## UpdateWVEPMetadata

> map[string]interface{} UpdateWVEPMetadata(ctx, did, wvm, wvmid, eid, iden, pid).Body(body).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).Execute()

Update workspace metadata by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    iden := "iden_example" // string | Denotes whether the pid specified is a part id (p) or a part identity (pi).
    pid := "pid_example" // string | 
    body := "body_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.UpdateWVEPMetadata(context.Background(), did, wvm, wvmid, eid, iden, pid).Body(body).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateWVEPMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWVEPMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.UpdateWVEPMetadata`: %v\n", resp)
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
**iden** | **string** | Denotes whether the pid specified is a part id (p) or a part identity (pi). | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVEPMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | **string** |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 

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


## UpdateWVMetadata

> map[string]interface{} UpdateWVMetadata(ctx, did, wv, wvid).Body(body).Execute()

Update workspace or version metadata by document ID and workspace or version ID.

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
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.UpdateWVMetadata(context.Background(), did, wv, wvid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateWVMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWVMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.UpdateWVMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

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

