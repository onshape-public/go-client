# \MetadataAPI

All URIs are relative to *https://cad.onshape.com/api/v9*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFullAssemblyMetadata**](MetadataAPI.md#GetFullAssemblyMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/assembly-debug | Get the metadata for an assembly, including supporting metadata.
[**GetVEOPStandardContentMetadata**](MetadataAPI.md#GetVEOPStandardContentMetadata) | **Get** /metadata/standardcontent/d/{did}/v/{vid}/e/{eid}/p/{pid} | Get the metadata for a standard content part.
[**GetWMVEMetadata**](MetadataAPI.md#GetWMVEMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | Get the metadata for an element.
[**GetWMVEPMetadata**](MetadataAPI.md#GetWMVEPMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/{iden}/{pid} | Get the metadata for a part.
[**GetWMVEPsMetadata**](MetadataAPI.md#GetWMVEPsMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p | Get the metadata for all parts in a document.
[**GetWMVEsMetadata**](MetadataAPI.md#GetWMVEsMetadata) | **Get** /metadata/d/{did}/{wvm}/{wvmid}/e | Get the metadata for all elements in a document.
[**GetWVMetadata**](MetadataAPI.md#GetWVMetadata) | **Get** /metadata/d/{did}/{wv}/{wvid} | Get the metadata for a workspace or version.
[**UpdateVEOPStandardContentPartMetadata**](MetadataAPI.md#UpdateVEOPStandardContentPartMetadata) | **Post** /metadata/standardcontent/d/{did} | Update the metadata for a standard content part.
[**UpdateWVEMetadata**](MetadataAPI.md#UpdateWVEMetadata) | **Post** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | Update the metadata for an element.
[**UpdateWVEPMetadata**](MetadataAPI.md#UpdateWVEPMetadata) | **Post** /metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/{iden}/{pid} | Update the metadata for a part.
[**UpdateWVMetadata**](MetadataAPI.md#UpdateWVMetadata) | **Post** /metadata/d/{did}/{wv}/{wvid} | Update the metadata for a workspace or version.



## GetFullAssemblyMetadata

> BTAssemblyItemMetadataInfo GetFullAssemblyMetadata(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()

Get the metadata for an assembly, including supporting metadata.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.GetFullAssemblyMetadata(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetFullAssemblyMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFullAssemblyMetadata`: BTAssemblyItemMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetFullAssemblyMetadata`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFullAssemblyMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | 

### Return type

[**BTAssemblyItemMetadataInfo**](BTAssemblyItemMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVEOPStandardContentMetadata

> BTMetadataObjectInfo GetVEOPStandardContentMetadata(ctx, did, vid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()

Get the metadata for a standard content part.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.GetVEOPStandardContentMetadata(context.Background(), did, vid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetVEOPStandardContentMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVEOPStandardContentMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetVEOPStandardContentMetadata`: %v\n", resp)
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

Get the metadata for an element.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.GetWMVEMetadata(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetWMVEMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetWMVEMetadata`: %v\n", resp)
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

Get the metadata for a part.



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
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    inferMetadataOwner := true // bool |  (optional) (default to false)
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.GetWMVEPMetadata(context.Background(), did, wvm, wvmid, eid, iden, pid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).InferMetadataOwner(inferMetadataOwner).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetWMVEPMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEPMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetWMVEPMetadata`: %v\n", resp)
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
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
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

Get the metadata for all parts in a document.



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
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional) (default to "")
    inferMetadataOwner := true // bool |  (optional) (default to false)
    includeComputedProperties := true // bool |  (optional) (default to true)
    includeComputedAssemblyProperties := true // bool |  (optional) (default to false)
    thumbnail := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.GetWMVEPsMetadata(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).InferMetadataOwner(inferMetadataOwner).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetWMVEPsMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEPsMetadata`: BTMetadataObjectListInfoBTMetadataPartInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetWMVEPsMetadata`: %v\n", resp)
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
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
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

Get the metadata for all elements in a document.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.GetWMVEsMetadata(context.Background(), did, wvm, wvmid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetWMVEsMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWMVEsMetadata`: BTMetadataObjectListInfoBTMetadataElementInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetWMVEsMetadata`: %v\n", resp)
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

Get the metadata for a workspace or version.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.GetWVMetadata(context.Background(), did, wv, wvid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).IncludeComputedProperties(includeComputedProperties).IncludeComputedAssemblyProperties(includeComputedAssemblyProperties).Thumbnail(thumbnail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.GetWVMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWVMetadata`: BTMetadataObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.GetWVMetadata`: %v\n", resp)
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

Update the metadata for a standard content part.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.UpdateVEOPStandardContentPartMetadata(context.Background(), did).LinkDocumentId(linkDocumentId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.UpdateVEOPStandardContentPartMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVEOPStandardContentPartMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.UpdateVEOPStandardContentPartMetadata`: %v\n", resp)
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

Update the metadata for an element.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.UpdateWVEMetadata(context.Background(), did, wvm, wvmid, eid).Body(body).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.UpdateWVEMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWVEMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.UpdateWVEMetadata`: %v\n", resp)
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

Update the metadata for a part.



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
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.UpdateWVEPMetadata(context.Background(), did, wvm, wvmid, eid, iden, pid).Body(body).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.UpdateWVEPMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWVEPMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.UpdateWVEPMetadata`: %v\n", resp)
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
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
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

Update the metadata for a workspace or version.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.MetadataAPI.UpdateWVMetadata(context.Background(), did, wv, wvid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataAPI.UpdateWVMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWVMetadata`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MetadataAPI.UpdateWVMetadata`: %v\n", resp)
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

