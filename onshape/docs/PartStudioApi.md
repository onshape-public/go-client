# \PartStudioApi

All URIs are relative to *https://cad.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPartStudioFeature**](PartStudioApi.md#AddPartStudioFeature) | **Post** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Add feature to the feature list for a Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**ComparePartStudios**](PartStudioApi.md#ComparePartStudios) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/compare | Compare Part Studios by document ID, workspace or version or microversion ID, and tab ID.
[**CreatePartStudio**](PartStudioApi.md#CreatePartStudio) | **Post** /partstudios/d/{did}/w/{wid} | Create Part Studio by document ID and workspace ID.
[**CreatePartStudioTranslation**](PartStudioApi.md#CreatePartStudioTranslation) | **Post** /partstudios/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create Part Studio translation by document ID, workspace or version ID, and tab ID.
[**DeletePartStudioFeature**](PartStudioApi.md#DeletePartStudioFeature) | **Delete** /partstudios/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Delete feature by document ID, workspace ID, tab ID, and feature ID.
[**EvalFeatureScript**](PartStudioApi.md#EvalFeatureScript) | **Post** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurescript | Evaluate FeatureScript for a Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**ExportParasolid**](PartStudioApi.md#ExportParasolid) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/parasolid | Export Part Studio to Parasolid by document ID, workspace or version or microversion ID, and tab ID.
[**ExportPartStudioGltf**](PartStudioApi.md#ExportPartStudioGltf) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/gltf | Export GLTF representation for parts in a Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**ExportPartStudioStl**](PartStudioApi.md#ExportPartStudioStl) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/stl | Export Part Studio to STL by document ID, workspace or version or microversion ID, and tab ID.
[**GetFeatureScriptRepresentation**](PartStudioApi.md#GetFeatureScriptRepresentation) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurescriptrepresentation | Retrieve FeatureScript representation of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**GetFeatureScriptTable**](PartStudioApi.md#GetFeatureScriptTable) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/fstable | Retrieve FeatureScript table of the Part Studio or part by document ID, workspace or version or microversion ID, and tab ID.
[**GetPartStudioBodyDetails**](PartStudioApi.md#GetPartStudioBodyDetails) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/bodydetails | Retrieve an array of body details by document ID, workspace or version or microversion ID, and tab ID.
[**GetPartStudioBoundingBoxes**](PartStudioApi.md#GetPartStudioBoundingBoxes) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/boundingboxes | Retrieve an array of Mass properties of parts or a Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**GetPartStudioEdges**](PartStudioApi.md#GetPartStudioEdges) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/tessellatededges | Retrieve tessellated edges of the parts in the Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**GetPartStudioFaces**](PartStudioApi.md#GetPartStudioFaces) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/tessellatedfaces | 
[**GetPartStudioFeatureSpecs**](PartStudioApi.md#GetPartStudioFeatureSpecs) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | Retrieve feature specifications of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**GetPartStudioFeatures**](PartStudioApi.md#GetPartStudioFeatures) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Retrieve a feature list of parts or a Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**GetPartStudioMassProperties**](PartStudioApi.md#GetPartStudioMassProperties) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/massproperties | Retrieve mass properties of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**GetPartStudioNamedViews**](PartStudioApi.md#GetPartStudioNamedViews) | **Get** /partstudios/d/{did}/e/{eid}/namedViews | 
[**GetPartStudioShadedViews**](PartStudioApi.md#GetPartStudioShadedViews) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/shadedviews | Retrieve shaded views of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.
[**TranslateIds**](PartStudioApi.md#TranslateIds) | **Post** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/idtranslations | Create Part Studio ID translation by document ID, workspace or version or microversion ID, and tab ID.
[**UpdateFeatures**](PartStudioApi.md#UpdateFeatures) | **Post** /partstudios/d/{did}/w/{wid}/e/{eid}/features/updates | Update features by document ID, workspace ID, and tab ID.
[**UpdatePartStudioFeature**](PartStudioApi.md#UpdatePartStudioFeature) | **Post** /partstudios/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Update feature by document ID, workspace ID, tab ID, and feature ID.
[**UpdateRollback**](PartStudioApi.md#UpdateRollback) | **Post** /partstudios/d/{did}/w/{wid}/e/{eid}/features/rollback | Update feature rollback by document ID, workspace ID, and tab ID.



## AddPartStudioFeature

> BTFeatureDefinitionResponse1617 AddPartStudioFeature(ctx, did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()

Add feature to the feature list for a Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    bTFeatureDefinitionCall1406 := *openapiclient.NewBTFeatureDefinitionCall1406() // BTFeatureDefinitionCall1406 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.AddPartStudioFeature(context.Background(), did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.AddPartStudioFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPartStudioFeature`: BTFeatureDefinitionResponse1617
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.AddPartStudioFeature`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddPartStudioFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTFeatureDefinitionCall1406** | [**BTFeatureDefinitionCall1406**](BTFeatureDefinitionCall1406.md) |  | 

### Return type

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse1617.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComparePartStudios

> BTRootDiffInfo ComparePartStudios(ctx, did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).SourceConfiguration(sourceConfiguration).TargetConfiguration(targetConfiguration).LinkDocumentId(linkDocumentId).Execute()

Compare Part Studios by document ID, workspace or version or microversion ID, and tab ID.

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
    workspaceId := "workspaceId_example" // string |  (optional)
    versionId := "versionId_example" // string |  (optional)
    microversionId := "microversionId_example" // string |  (optional)
    sourceConfiguration := "sourceConfiguration_example" // string |  (optional)
    targetConfiguration := "targetConfiguration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.ComparePartStudios(context.Background(), did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).SourceConfiguration(sourceConfiguration).TargetConfiguration(targetConfiguration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.ComparePartStudios``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ComparePartStudios`: BTRootDiffInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.ComparePartStudios`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiComparePartStudiosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **workspaceId** | **string** |  | 
 **versionId** | **string** |  | 
 **microversionId** | **string** |  | 
 **sourceConfiguration** | **string** |  | 
 **targetConfiguration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTRootDiffInfo**](BTRootDiffInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudio

> BTDocumentElementInfo CreatePartStudio(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create Part Studio by document ID and workspace ID.

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
    bTModelElementParams := *openapiclient.NewBTModelElementParams() // BTModelElementParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.CreatePartStudio(context.Background(), did, wid).BTModelElementParams(bTModelElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.CreatePartStudio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartStudio`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.CreatePartStudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wid** | **string** | Workspace ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartStudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTModelElementParams** | [**BTModelElementParams**](BTModelElementParams.md) |  | 

### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudioTranslation

> BTTranslationRequestInfo CreatePartStudioTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create Part Studio translation by document ID, workspace or version ID, and tab ID.

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
    wv := "wv_example" // string | One of w or v corresponding to whether a workspace or version was specified.
    wvid := "wvid_example" // string | Workspace (w) or Version (v) ID.
    eid := "eid_example" // string | Element ID.
    bTTranslateFormatParams := *openapiclient.NewBTTranslateFormatParams() // BTTranslateFormatParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.CreatePartStudioTranslation(context.Background(), did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.CreatePartStudioTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartStudioTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.CreatePartStudioTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wv** | **string** | One of w or v corresponding to whether a workspace or version was specified. | 
**wvid** | **string** | Workspace (w) or Version (v) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartStudioTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTTranslateFormatParams** | [**BTTranslateFormatParams**](BTTranslateFormatParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePartStudioFeature

> BTFeatureApiBase1430 DeletePartStudioFeature(ctx, did, wid, eid, fid).Execute()

Delete feature by document ID, workspace ID, tab ID, and feature ID.

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
    fid := "fid_example" // string | The id of the feature being updated. This id should be URL encoded and must match the featureId found in the serialized structure

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.DeletePartStudioFeature(context.Background(), did, wid, eid, fid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.DeletePartStudioFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePartStudioFeature`: BTFeatureApiBase1430
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.DeletePartStudioFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wid** | **string** | Workspace ID. | 
**eid** | **string** | Element ID. | 
**fid** | **string** | The id of the feature being updated. This id should be URL encoded and must match the featureId found in the serialized structure | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePartStudioFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTFeatureApiBase1430**](BTFeatureApiBase1430.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvalFeatureScript

> BTFeatureScriptEvalResponse1859 EvalFeatureScript(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).BTFeatureScriptEvalCall2377(bTFeatureScriptEvalCall2377).Execute()

Evaluate FeatureScript for a Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    bTFeatureScriptEvalCall2377 := *openapiclient.NewBTFeatureScriptEvalCall2377() // BTFeatureScriptEvalCall2377 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.EvalFeatureScript(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).BTFeatureScriptEvalCall2377(bTFeatureScriptEvalCall2377).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.EvalFeatureScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EvalFeatureScript`: BTFeatureScriptEvalResponse1859
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.EvalFeatureScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvalFeatureScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **bTFeatureScriptEvalCall2377** | [**BTFeatureScriptEvalCall2377**](BTFeatureScriptEvalCall2377.md) |  | 

### Return type

[**BTFeatureScriptEvalResponse1859**](BTFeatureScriptEvalResponse1859.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportParasolid

> ExportParasolid(ctx, did, wvm, wvmid, eid).PartIds(partIds).Version(version).IncludeExportIds(includeExportIds).Configuration(configuration).LinkDocumentId(linkDocumentId).BinaryExport(binaryExport).Execute()

Export Part Studio to Parasolid by document ID, workspace or version or microversion ID, and tab ID.

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
    partIds := "partIds_example" // string | IDs of the parts to retrieve. Repeat query param to add more than one (i.e. partId=JHK&partId=JHD). May not be combined with other ID filters (optional)
    version := "version_example" // string | Parasolid version (optional) (default to "0")
    includeExportIds := true // bool | Whether topology ids should be exported as parasolid attributes (optional) (default to false)
    configuration := "configuration_example" // string | Configuration string. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)
    binaryExport := true // bool | Whether to use binary parasolid format instead of text (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.ExportParasolid(context.Background(), did, wvm, wvmid, eid).PartIds(partIds).Version(version).IncludeExportIds(includeExportIds).Configuration(configuration).LinkDocumentId(linkDocumentId).BinaryExport(binaryExport).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.ExportParasolid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiExportParasolidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **partIds** | **string** | IDs of the parts to retrieve. Repeat query param to add more than one (i.e. partId&#x3D;JHK&amp;partId&#x3D;JHD). May not be combined with other ID filters | 
 **version** | **string** | Parasolid version | [default to &quot;0&quot;]
 **includeExportIds** | **bool** | Whether topology ids should be exported as parasolid attributes | [default to false]
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 
 **binaryExport** | **bool** | Whether to use binary parasolid format instead of text | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPartStudioGltf

> HttpFile ExportPartStudioGltf(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).OutputSeparateFaceNodes(outputSeparateFaceNodes).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).Execute()

Export GLTF representation for parts in a Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    partId := []string{"Inner_example"} // []string |  (optional)
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    precomputedLevelOfDetail := "precomputedLevelOfDetail_example" // string |  (optional)
    outputSeparateFaceNodes := true // bool |  (optional) (default to false)
    faceId := []string{"Inner_example"} // []string |  (optional)
    outputFaceAppearances := true // bool |  (optional) (default to false)
    maxFacetWidth := float64(1.2) // float64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.ExportPartStudioGltf(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).OutputSeparateFaceNodes(outputSeparateFaceNodes).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.ExportPartStudioGltf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPartStudioGltf`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.ExportPartStudioGltf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPartStudioGltfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **partId** | **[]string** |  | 
 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **precomputedLevelOfDetail** | **string** |  | 
 **outputSeparateFaceNodes** | **bool** |  | [default to false]
 **faceId** | **[]string** |  | 
 **outputFaceAppearances** | **bool** |  | [default to false]
 **maxFacetWidth** | **float64** |  | 

### Return type

[**HttpFile**](HttpFile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: model/gltf-binary;qs=0.08

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPartStudioStl

> ExportPartStudioStl(ctx, did, wvm, wvmid, eid).PartIds(partIds).Mode(mode).Grouping(grouping).Scale(scale).Units(units).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).MinFacetWidth(minFacetWidth).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Export Part Studio to STL by document ID, workspace or version or microversion ID, and tab ID.

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
    partIds := "partIds_example" // string | IDs of the parts to retrieve. Repeat query param to add more than one (i.e. partId=JHK&partId=JHD). May not be combined with other ID filters (optional)
    mode := "mode_example" // string | Type of file: text, binary (optional) (default to "text")
    grouping := true // bool | Whether parts should be exported as a group or individually in a .zip file (optional) (default to true)
    scale := float64(1.2) // float64 | Scale for measurements. (optional) (default to 1)
    units := "units_example" // string | Name of base unit (meter, centimeter, millimeter, inch, foot, or yard) (optional) (default to "inch")
    angleTolerance := float64(1.2) // float64 | Angle tolerance (in radians). This specifies the limit on the sum of the angular deviations of a tessellation chord from the tangent vectors at two chord endpoints. The specified value must be less than PI/2. This parameter currently has a default value chosen based on the complexity of the parts being tessellated. (optional)
    chordTolerance := float64(1.2) // float64 | Chord tolerance (in meters). This specifies the limit on the maximum deviation of a tessellation chord from the true surface/edge. This parameter currently has a default value chosen based on the size and complexity of the parts being tessellated. (optional)
    maxFacetWidth := float64(1.2) // float64 | Max facet width. This specifies the limit on the size of any side of a tessellation facet. (optional)
    minFacetWidth := float64(1.2) // float64 | Max facet width. This specifies the limit on the size of any side of a tessellation facet. (optional)
    configuration := "configuration_example" // string | Configuration string. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.ExportPartStudioStl(context.Background(), did, wvm, wvmid, eid).PartIds(partIds).Mode(mode).Grouping(grouping).Scale(scale).Units(units).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).MinFacetWidth(minFacetWidth).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.ExportPartStudioStl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiExportPartStudioStlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **partIds** | **string** | IDs of the parts to retrieve. Repeat query param to add more than one (i.e. partId&#x3D;JHK&amp;partId&#x3D;JHD). May not be combined with other ID filters | 
 **mode** | **string** | Type of file: text, binary | [default to &quot;text&quot;]
 **grouping** | **bool** | Whether parts should be exported as a group or individually in a .zip file | [default to true]
 **scale** | **float64** | Scale for measurements. | [default to 1]
 **units** | **string** | Name of base unit (meter, centimeter, millimeter, inch, foot, or yard) | [default to &quot;inch&quot;]
 **angleTolerance** | **float64** | Angle tolerance (in radians). This specifies the limit on the sum of the angular deviations of a tessellation chord from the tangent vectors at two chord endpoints. The specified value must be less than PI/2. This parameter currently has a default value chosen based on the complexity of the parts being tessellated. | 
 **chordTolerance** | **float64** | Chord tolerance (in meters). This specifies the limit on the maximum deviation of a tessellation chord from the true surface/edge. This parameter currently has a default value chosen based on the size and complexity of the parts being tessellated. | 
 **maxFacetWidth** | **float64** | Max facet width. This specifies the limit on the size of any side of a tessellation facet. | 
 **minFacetWidth** | **float64** | Max facet width. This specifies the limit on the size of any side of a tessellation facet. | 
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureScriptRepresentation

> BTPModule234 GetFeatureScriptRepresentation(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).Execute()

Retrieve FeatureScript representation of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetFeatureScriptRepresentation(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetFeatureScriptRepresentation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureScriptRepresentation`: BTPModule234
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetFeatureScriptRepresentation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureScriptRepresentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 

### Return type

[**BTPModule234**](BTPModule234.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureScriptTable

> BTApiTableList1223 GetFeatureScriptTable(ctx, did, wvm, wvmid, eid).TableType(tableType).Configuration(configuration).TableNamespace(tableNamespace).TableParameters(tableParameters).PartId(partId).LinkDocumentId(linkDocumentId).Execute()

Retrieve FeatureScript table of the Part Studio or part by document ID, workspace or version or microversion ID, and tab ID.

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
    tableType := "tableType_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    tableNamespace := "tableNamespace_example" // string |  (optional)
    tableParameters := "tableParameters_example" // string |  (optional)
    partId := "partId_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetFeatureScriptTable(context.Background(), did, wvm, wvmid, eid).TableType(tableType).Configuration(configuration).TableNamespace(tableNamespace).TableParameters(tableParameters).PartId(partId).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetFeatureScriptTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureScriptTable`: BTApiTableList1223
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetFeatureScriptTable`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFeatureScriptTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tableType** | **string** |  | 
 **configuration** | **string** |  | 
 **tableNamespace** | **string** |  | 
 **tableParameters** | **string** |  | 
 **partId** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTApiTableList1223**](BTApiTableList1223.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioBodyDetails

> BTExportModelBodiesResponse734 GetPartStudioBodyDetails(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).IncludeSurfaces(includeSurfaces).IncludeCompositeParts(includeCompositeParts).IncludeGeometricData(includeGeometricData).Execute()

Retrieve an array of body details by document ID, workspace or version or microversion ID, and tab ID.

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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    includeSurfaces := true // bool | Whether or not surfaces should be included in the response. (optional) (default to false)
    includeCompositeParts := true // bool | Whether or not composite parts should be included in the response. (optional) (default to false)
    includeGeometricData := true // bool | Whether or not geometric data should be included in the response. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioBodyDetails(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).IncludeSurfaces(includeSurfaces).IncludeCompositeParts(includeCompositeParts).IncludeGeometricData(includeGeometricData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioBodyDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioBodyDetails`: BTExportModelBodiesResponse734
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioBodyDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioBodyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **includeSurfaces** | **bool** | Whether or not surfaces should be included in the response. | [default to false]
 **includeCompositeParts** | **bool** | Whether or not composite parts should be included in the response. | [default to false]
 **includeGeometricData** | **bool** | Whether or not geometric data should be included in the response. | [default to true]

### Return type

[**BTExportModelBodiesResponse734**](BTExportModelBodiesResponse734.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioBoundingBoxes

> BTBoundingBoxInfo GetPartStudioBoundingBoxes(ctx, did, wvm, wvmid, eid).IncludeHidden(includeHidden).IncludeWireBodies(includeWireBodies).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Retrieve an array of Mass properties of parts or a Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    includeHidden := true // bool | Whether or not to include bounding boxes for hidden parts. (optional) (default to false)
    includeWireBodies := true // bool | Whether to include wire bodies in the bounding box. (optional) (default to true)
    configuration := "configuration_example" // string | Configuration string. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioBoundingBoxes(context.Background(), did, wvm, wvmid, eid).IncludeHidden(includeHidden).IncludeWireBodies(includeWireBodies).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioBoundingBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioBoundingBoxes`: BTBoundingBoxInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioBoundingBoxes`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPartStudioBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeHidden** | **bool** | Whether or not to include bounding boxes for hidden parts. | [default to false]
 **includeWireBodies** | **bool** | Whether to include wire bodies in the bounding box. | [default to true]
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**BTBoundingBoxInfo**](BTBoundingBoxInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioEdges

> BTExportTessellatedEdgesResponse327 GetPartStudioEdges(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).EdgeId(edgeId).Execute()

Retrieve tessellated edges of the parts in the Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    partId := []string{"Inner_example"} // []string |  (optional)
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    precomputedLevelOfDetail := "precomputedLevelOfDetail_example" // string |  (optional)
    edgeId := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioEdges(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).EdgeId(edgeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioEdges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioEdges`: BTExportTessellatedEdgesResponse327
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioEdges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **partId** | **[]string** |  | 
 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **precomputedLevelOfDetail** | **string** |  | 
 **edgeId** | **[]string** |  | 

### Return type

[**BTExportTessellatedEdgesResponse327**](BTExportTessellatedEdgesResponse327.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFaces

> BTExportTessellatedFacesResponse898 GetPartStudioFaces(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputIndexTable(outputIndexTable).OutputErrorFaces(outputErrorFaces).CombineCompositePartConstituents(combineCompositePartConstituents).Execute()



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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    partId := []string{"Inner_example"} // []string |  (optional)
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    precomputedLevelOfDetail := "precomputedLevelOfDetail_example" // string |  (optional)
    faceId := []string{"Inner_example"} // []string |  (optional)
    outputFaceAppearances := true // bool |  (optional) (default to false)
    maxFacetWidth := float64(1.2) // float64 |  (optional)
    outputVertexNormals := true // bool |  (optional) (default to false)
    outputFacetNormals := true // bool |  (optional) (default to true)
    outputTextureCoordinates := true // bool |  (optional) (default to false)
    outputIndexTable := true // bool |  (optional) (default to false)
    outputErrorFaces := true // bool |  (optional) (default to false)
    combineCompositePartConstituents := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioFaces(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputIndexTable(outputIndexTable).OutputErrorFaces(outputErrorFaces).CombineCompositePartConstituents(combineCompositePartConstituents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioFaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioFaces`: BTExportTessellatedFacesResponse898
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioFaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioFacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **partId** | **[]string** |  | 
 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **precomputedLevelOfDetail** | **string** |  | 
 **faceId** | **[]string** |  | 
 **outputFaceAppearances** | **bool** |  | [default to false]
 **maxFacetWidth** | **float64** |  | 
 **outputVertexNormals** | **bool** |  | [default to false]
 **outputFacetNormals** | **bool** |  | [default to true]
 **outputTextureCoordinates** | **bool** |  | [default to false]
 **outputIndexTable** | **bool** |  | [default to false]
 **outputErrorFaces** | **bool** |  | [default to false]
 **combineCompositePartConstituents** | **bool** |  | [default to false]

### Return type

[**BTExportTessellatedFacesResponse898**](BTExportTessellatedFacesResponse898.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFeatureSpecs

> BTFeatureSpecsResponse664 GetPartStudioFeatureSpecs(ctx, did, wvm, wvmid, eid).Execute()

Retrieve feature specifications of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioFeatureSpecs(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioFeatureSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioFeatureSpecs`: BTFeatureSpecsResponse664
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioFeatureSpecs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPartStudioFeatureSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTFeatureSpecsResponse664**](BTFeatureSpecsResponse664.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFeatures

> BTFeatureListResponse2457 GetPartStudioFeatures(ctx, did, wvm, wvmid, eid).IncludeGeometryIds(includeGeometryIds).FeatureId(featureId).LinkDocumentId(linkDocumentId).NoSketchGeometry(noSketchGeometry).Execute()

Retrieve a feature list of parts or a Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    includeGeometryIds := true // bool |  (optional) (default to true)
    featureId := []string{"Inner_example"} // []string | ID of a feature; repeat query param to add more than one (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)
    noSketchGeometry := true // bool | Whether or not to output simple sketch info without geometry (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioFeatures(context.Background(), did, wvm, wvmid, eid).IncludeGeometryIds(includeGeometryIds).FeatureId(featureId).LinkDocumentId(linkDocumentId).NoSketchGeometry(noSketchGeometry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioFeatures`: BTFeatureListResponse2457
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioFeatures`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPartStudioFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includeGeometryIds** | **bool** |  | [default to true]
 **featureId** | **[]string** | ID of a feature; repeat query param to add more than one | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 
 **noSketchGeometry** | **bool** | Whether or not to output simple sketch info without geometry | [default to false]

### Return type

[**BTFeatureListResponse2457**](BTFeatureListResponse2457.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioMassProperties

> BTMassPropertiesBulkInfo GetPartStudioMassProperties(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).MassAsGroup(massAsGroup).UseMassPropertyOverrides(useMassPropertyOverrides).Execute()

Retrieve mass properties of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    wvm := "wvm_example" // string | Indicates which of workspace id, version id, or document microversion id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version, or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    partId := []string{"Inner_example"} // []string |  (optional)
    massAsGroup := true // bool | If true, specified parts will be evaluated as a single object instead of individually (optional) (default to true)
    useMassPropertyOverrides := true // bool | If true, use the user mass property overrides when calculated mass properties (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioMassProperties(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).MassAsGroup(massAsGroup).UseMassPropertyOverrides(useMassPropertyOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioMassProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioMassProperties`: BTMassPropertiesBulkInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioMassProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioMassPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **partId** | **[]string** |  | 
 **massAsGroup** | **bool** | If true, specified parts will be evaluated as a single object instead of individually | [default to true]
 **useMassPropertyOverrides** | **bool** | If true, use the user mass property overrides when calculated mass properties | [default to false]

### Return type

[**BTMassPropertiesBulkInfo**](BTMassPropertiesBulkInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioNamedViews

> BTNamedViewsInfo GetPartStudioNamedViews(ctx, did, eid).LinkDocumentId(linkDocumentId).SkipPerspective(skipPerspective).IncludeSectionCutViews(includeSectionCutViews).Execute()



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
    eid := "eid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    skipPerspective := true // bool |  (optional) (default to true)
    includeSectionCutViews := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioNamedViews(context.Background(), did, eid).LinkDocumentId(linkDocumentId).SkipPerspective(skipPerspective).IncludeSectionCutViews(includeSectionCutViews).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioNamedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioNamedViews`: BTNamedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioNamedViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioNamedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **skipPerspective** | **bool** |  | [default to true]
 **includeSectionCutViews** | **bool** |  | [default to false]

### Return type

[**BTNamedViewsInfo**](BTNamedViewsInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioShadedViews

> BTShadedViewsInfo GetPartStudioShadedViews(ctx, did, wvm, wvmid, eid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).IncludeWires(includeWires).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Retrieve shaded views of the Part Studio by document ID, workspace or version or microversion ID, and tab ID.

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
    viewMatrix := "viewMatrix_example" // string | 12-number view matrix (comma-separated), or one of the following named views: top, bottom, front, back, left, right The 12 entries in the view matrix form three rows and four columns, which is a linear transformation applied to the model itself. The matrix's first three columns maps the coordinate axes of the model to the coordinate axes of the view, and the fourth column translates the origin (in meters). The view coordinates have x pointing right, y pointing up, and z pointing towards the viewer, while a front view of the model has x pointing right, y pointing away from the viewer, and z pointing up. For example, the identity matrix viewMatrix=1,0,0,0,0,1,0,0,0,0,1,0 corresponds to the top view, and viewMatrix=0.612,0.612,0,0,-0.354,0.354,0.707,0,0.707,-0.707,0.707,0 corresponds (approximately) to the isometric view. The first three columns of the view matrix should be orthonormal and have a positive determinant.  If this is not the case, view behavior may be undefined. (optional) (default to "front")
    outputHeight := int32(56) // int32 | Output image height (in pixels) (optional) (default to 500)
    outputWidth := int32(56) // int32 | Output image width (in pixels) (optional) (default to 500)
    pixelSize := float64(1.2) // float64 | Height and width represented by each pixel (in meters). If the value is 0, the display will be sized to fit the output image dimensions. (optional) (default to 0.003)
    edges := "edges_example" // string | The treatment to be applied to edges in the display. Options are show: show visible edges, hide: hide visible edges. (optional) (default to "show")
    showAllParts := true // bool | Whether or not all parts should be shown in the element, regardless of user setting. If false, the visibility setting made by the user will be reflected in the image. If true, all parts will be shown. (optional) (default to false)
    includeSurfaces := true // bool | Whether or not surfaces should be shown in the element. It is applicable only when showAllParts is true. If false, surfaces will be excluded. If true, all surfaces will be shown. (optional) (default to false)
    useAntiAliasing := true // bool | If true, an anti-aliasing factor will be used to smooth model boundaries in the final image result. If false, the image will be rasterized at the given resolution. Setting to true can have negative performance implications with respect to rendering time and memory usage. If a high-resolution image is requested and anti-aliasing is turned on, the server may not be able to fulfill the request. (optional) (default to false)
    includeWires := true // bool |  (optional) (default to false)
    configuration := "configuration_example" // string | Configuration string. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioShadedViews(context.Background(), did, wvm, wvmid, eid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).IncludeWires(includeWires).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.GetPartStudioShadedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartStudioShadedViews`: BTShadedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.GetPartStudioShadedViews`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPartStudioShadedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **viewMatrix** | **string** | 12-number view matrix (comma-separated), or one of the following named views: top, bottom, front, back, left, right The 12 entries in the view matrix form three rows and four columns, which is a linear transformation applied to the model itself. The matrix&#39;s first three columns maps the coordinate axes of the model to the coordinate axes of the view, and the fourth column translates the origin (in meters). The view coordinates have x pointing right, y pointing up, and z pointing towards the viewer, while a front view of the model has x pointing right, y pointing away from the viewer, and z pointing up. For example, the identity matrix viewMatrix&#x3D;1,0,0,0,0,1,0,0,0,0,1,0 corresponds to the top view, and viewMatrix&#x3D;0.612,0.612,0,0,-0.354,0.354,0.707,0,0.707,-0.707,0.707,0 corresponds (approximately) to the isometric view. The first three columns of the view matrix should be orthonormal and have a positive determinant.  If this is not the case, view behavior may be undefined. | [default to &quot;front&quot;]
 **outputHeight** | **int32** | Output image height (in pixels) | [default to 500]
 **outputWidth** | **int32** | Output image width (in pixels) | [default to 500]
 **pixelSize** | **float64** | Height and width represented by each pixel (in meters). If the value is 0, the display will be sized to fit the output image dimensions. | [default to 0.003]
 **edges** | **string** | The treatment to be applied to edges in the display. Options are show: show visible edges, hide: hide visible edges. | [default to &quot;show&quot;]
 **showAllParts** | **bool** | Whether or not all parts should be shown in the element, regardless of user setting. If false, the visibility setting made by the user will be reflected in the image. If true, all parts will be shown. | [default to false]
 **includeSurfaces** | **bool** | Whether or not surfaces should be shown in the element. It is applicable only when showAllParts is true. If false, surfaces will be excluded. If true, all surfaces will be shown. | [default to false]
 **useAntiAliasing** | **bool** | If true, an anti-aliasing factor will be used to smooth model boundaries in the final image result. If false, the image will be rasterized at the given resolution. Setting to true can have negative performance implications with respect to rendering time and memory usage. If a high-resolution image is requested and anti-aliasing is turned on, the server may not be able to fulfill the request. | [default to false]
 **includeWires** | **bool** |  | [default to false]
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**BTShadedViewsInfo**](BTShadedViewsInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslateIds

> BTIdTranslationInfo TranslateIds(ctx, did, wvm, wvmid, eid).BTIdTranslationParams(bTIdTranslationParams).Execute()

Create Part Studio ID translation by document ID, workspace or version or microversion ID, and tab ID.

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
    bTIdTranslationParams := *openapiclient.NewBTIdTranslationParams() // BTIdTranslationParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.TranslateIds(context.Background(), did, wvm, wvmid, eid).BTIdTranslationParams(bTIdTranslationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.TranslateIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslateIds`: BTIdTranslationInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.TranslateIds`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTranslateIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTIdTranslationParams** | [**BTIdTranslationParams**](BTIdTranslationParams.md) |  | 

### Return type

[**BTIdTranslationInfo**](BTIdTranslationInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatures

> BTUpdateFeaturesResponse1333 UpdateFeatures(ctx, did, wid, eid).BTUpdateFeaturesCall1748(bTUpdateFeaturesCall1748).Execute()

Update features by document ID, workspace ID, and tab ID.

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
    bTUpdateFeaturesCall1748 := *openapiclient.NewBTUpdateFeaturesCall1748() // BTUpdateFeaturesCall1748 | feature The serialized feature definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.UpdateFeatures(context.Background(), did, wid, eid).BTUpdateFeaturesCall1748(bTUpdateFeaturesCall1748).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.UpdateFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatures`: BTUpdateFeaturesResponse1333
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.UpdateFeatures`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateFeaturesCall1748** | [**BTUpdateFeaturesCall1748**](BTUpdateFeaturesCall1748.md) | feature The serialized feature definition | 

### Return type

[**BTUpdateFeaturesResponse1333**](BTUpdateFeaturesResponse1333.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartStudioFeature

> BTFeatureDefinitionResponse1617 UpdatePartStudioFeature(ctx, did, wid, eid, fid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()

Update feature by document ID, workspace ID, tab ID, and feature ID.

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
    fid := "fid_example" // string | The id of the feature being updated. This id should be URL encoded and must match the featureId found in the serialized structure
    bTFeatureDefinitionCall1406 := *openapiclient.NewBTFeatureDefinitionCall1406() // BTFeatureDefinitionCall1406 | feature The serialized feature definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.UpdatePartStudioFeature(context.Background(), did, wid, eid, fid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.UpdatePartStudioFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePartStudioFeature`: BTFeatureDefinitionResponse1617
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.UpdatePartStudioFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wid** | **string** | Workspace ID. | 
**eid** | **string** | Element ID. | 
**fid** | **string** | The id of the feature being updated. This id should be URL encoded and must match the featureId found in the serialized structure | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartStudioFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTFeatureDefinitionCall1406** | [**BTFeatureDefinitionCall1406**](BTFeatureDefinitionCall1406.md) | feature The serialized feature definition | 

### Return type

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse1617.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRollback

> BTSetFeatureRollbackResponse1042 UpdateRollback(ctx, did, wid, eid).Body(body).Execute()

Update feature rollback by document ID, workspace ID, and tab ID.

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
    body := "body_example" // string | The index at which the rollback index should be placed. Features  with entry index (0-based) higher than or equal to the value are rolled back. Value of -1 is treated  as an alias for \"end of feature list\". Otherwise the value must be in the range 0 to the number of  entries in the feature list

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartStudioApi.UpdateRollback(context.Background(), did, wid, eid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.UpdateRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRollback`: BTSetFeatureRollbackResponse1042
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.UpdateRollback`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** | The index at which the rollback index should be placed. Features  with entry index (0-based) higher than or equal to the value are rolled back. Value of -1 is treated  as an alias for \&quot;end of feature list\&quot;. Otherwise the value must be in the range 0 to the number of  entries in the feature list | 

### Return type

[**BTSetFeatureRollbackResponse1042**](BTSetFeatureRollbackResponse1042.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

