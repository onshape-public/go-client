# \PartStudioApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPartStudioFeature**](PartStudioApi.md#AddPartStudioFeature) | **Post** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Add a feature to the Part Studio&#39;s Feature List.
[**ComparePartStudios**](PartStudioApi.md#ComparePartStudios) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/compare | Get the differences between two Part Studios in a single document.
[**CreatePartStudio**](PartStudioApi.md#CreatePartStudio) | **Post** /partstudios/d/{did}/w/{wid} | Create a new Part Studio in a document.
[**CreatePartStudioExportGltf**](PartStudioApi.md#CreatePartStudioExportGltf) | **Post** /partstudios/d/{did}/{wv}/{wvid}/e/{eid}/export/gltf | Asynchronously export a Part Studio to glTF.
[**CreatePartStudioExportObj**](PartStudioApi.md#CreatePartStudioExportObj) | **Post** /partstudios/d/{did}/{wv}/{wvid}/e/{eid}/export/obj | Asynchronously export a Part Studio to OBJ.
[**CreatePartStudioExportSolidworks**](PartStudioApi.md#CreatePartStudioExportSolidworks) | **Post** /partstudios/d/{did}/{wv}/{wvid}/e/{eid}/export/solidworks | Asynchronously export a Part Studio to Solidworks.
[**CreatePartStudioExportStep**](PartStudioApi.md#CreatePartStudioExportStep) | **Post** /partstudios/d/{did}/{wv}/{wvid}/e/{eid}/export/step | Asynchronously export a Part Studio to STEP.
[**CreatePartStudioTranslation**](PartStudioApi.md#CreatePartStudioTranslation) | **Post** /partstudios/d/{did}/{wv}/{wvid}/e/{eid}/translations | Asynchronously export a Part Studio to another format.
[**DeletePartStudioFeature**](PartStudioApi.md#DeletePartStudioFeature) | **Delete** /partstudios/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Delete a Part Studio feature.
[**EvalFeatureScript**](PartStudioApi.md#EvalFeatureScript) | **Post** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurescript | Evaluate the FeatureScript snippet for a Part Studio.
[**ExportParasolid**](PartStudioApi.md#ExportParasolid) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/parasolid | Synchronously export a Part Studio to a Parasolid file.
[**ExportPartStudioGltf**](PartStudioApi.md#ExportPartStudioGltf) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/gltf | Synchronously export a Part Studio to a glTF file.
[**ExportPartStudioStl**](PartStudioApi.md#ExportPartStudioStl) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/stl | Synchronously export a Part Studio to an STL file.
[**GetFeatureScriptRepresentation**](PartStudioApi.md#GetFeatureScriptRepresentation) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurescriptrepresentation | Get the FeatureScript representation of a Part Studio.
[**GetFeatureScriptTable**](PartStudioApi.md#GetFeatureScriptTable) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/fstable | Compute and return a FeatureScript table for a Part Studio.
[**GetPartStudioBodyDetails**](PartStudioApi.md#GetPartStudioBodyDetails) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/bodydetails | Get the body details for a Part Studio.
[**GetPartStudioBoundingBoxes**](PartStudioApi.md#GetPartStudioBoundingBoxes) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/boundingboxes | Get the bounding boxes for a Part Studio.
[**GetPartStudioEdges**](PartStudioApi.md#GetPartStudioEdges) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/tessellatededges | Get a list of all edges in a Part Studio.
[**GetPartStudioFaces**](PartStudioApi.md#GetPartStudioFaces) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/tessellatedfaces | Get a list of all faces in a Part Studio.
[**GetPartStudioFeatureSpecs**](PartStudioApi.md#GetPartStudioFeatureSpecs) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | Get the specs for a Part Studio feature.
[**GetPartStudioFeatures**](PartStudioApi.md#GetPartStudioFeatures) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Get a list of features instantiated in the Part Studio.
[**GetPartStudioMassProperties**](PartStudioApi.md#GetPartStudioMassProperties) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/massproperties | Get the mass properties for a Part Studio.
[**GetPartStudioNamedViews**](PartStudioApi.md#GetPartStudioNamedViews) | **Get** /partstudios/d/{did}/e/{eid}/namedViews | Get a list of all named views that exist in the Part Studio.
[**GetPartStudioShadedViews**](PartStudioApi.md#GetPartStudioShadedViews) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/shadedviews | Get a list of shaded views for a Part Studio.
[**TranslateIds**](PartStudioApi.md#TranslateIds) | **Post** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/idtranslations | Find corresponding deterministic IDs from a source document microversion at the target version.
[**UpdateFeatures**](PartStudioApi.md#UpdateFeatures) | **Post** /partstudios/d/{did}/w/{wid}/e/{eid}/features/updates | Update multiple features in a Part Studio
[**UpdatePartStudioFeature**](PartStudioApi.md#UpdatePartStudioFeature) | **Post** /partstudios/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Update the definition of a Part Studio feature.
[**UpdateRollback**](PartStudioApi.md#UpdateRollback) | **Post** /partstudios/d/{did}/w/{wid}/e/{eid}/features/rollback | Move the Feature List rollback bar in the Part Studio.



## AddPartStudioFeature

> BTFeatureDefinitionResponse1617 AddPartStudioFeature(ctx, did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()

Add a feature to the Part Studio's Feature List.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComparePartStudios

> BTRootDiffInfo ComparePartStudios(ctx, did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).SourceConfiguration(sourceConfiguration).TargetConfiguration(targetConfiguration).LinkDocumentId(linkDocumentId).Execute()

Get the differences between two Part Studios in a single document.

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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudio

> BTDocumentElementInfo CreatePartStudio(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create a new Part Studio in a document.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudioExportGltf

> BTTranslationRequestInfo CreatePartStudioExportGltf(ctx, did, wv, wvid, eid).BTBGltfExportParams(bTBGltfExportParams).Execute()

Asynchronously export a Part Studio to glTF.



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
    bTBGltfExportParams := *openapiclient.NewBTBGltfExportParams() // BTBGltfExportParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartStudioApi.CreatePartStudioExportGltf(context.Background(), did, wv, wvid, eid).BTBGltfExportParams(bTBGltfExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.CreatePartStudioExportGltf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartStudioExportGltf`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.CreatePartStudioExportGltf`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreatePartStudioExportGltfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBGltfExportParams** | [**BTBGltfExportParams**](BTBGltfExportParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudioExportObj

> BTTranslationRequestInfo CreatePartStudioExportObj(ctx, did, wv, wvid, eid).BTBObjExportParams(bTBObjExportParams).Execute()

Asynchronously export a Part Studio to OBJ.



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
    bTBObjExportParams := *openapiclient.NewBTBObjExportParams() // BTBObjExportParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartStudioApi.CreatePartStudioExportObj(context.Background(), did, wv, wvid, eid).BTBObjExportParams(bTBObjExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.CreatePartStudioExportObj``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartStudioExportObj`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.CreatePartStudioExportObj`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreatePartStudioExportObjRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBObjExportParams** | [**BTBObjExportParams**](BTBObjExportParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudioExportSolidworks

> BTTranslationRequestInfo CreatePartStudioExportSolidworks(ctx, did, wv, wvid, eid).BTBSolidworksExportParams(bTBSolidworksExportParams).Execute()

Asynchronously export a Part Studio to Solidworks.



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
    bTBSolidworksExportParams := *openapiclient.NewBTBSolidworksExportParams() // BTBSolidworksExportParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartStudioApi.CreatePartStudioExportSolidworks(context.Background(), did, wv, wvid, eid).BTBSolidworksExportParams(bTBSolidworksExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.CreatePartStudioExportSolidworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartStudioExportSolidworks`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.CreatePartStudioExportSolidworks`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreatePartStudioExportSolidworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBSolidworksExportParams** | [**BTBSolidworksExportParams**](BTBSolidworksExportParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudioExportStep

> BTTranslationRequestInfo CreatePartStudioExportStep(ctx, did, wv, wvid, eid).BTBStepExportParams(bTBStepExportParams).Execute()

Asynchronously export a Part Studio to STEP.



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
    bTBStepExportParams := *openapiclient.NewBTBStepExportParams() // BTBStepExportParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartStudioApi.CreatePartStudioExportStep(context.Background(), did, wv, wvid, eid).BTBStepExportParams(bTBStepExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.CreatePartStudioExportStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartStudioExportStep`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.CreatePartStudioExportStep`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreatePartStudioExportStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBStepExportParams** | [**BTBStepExportParams**](BTBStepExportParams.md) |  | 

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudioTranslation

> BTTranslationRequestInfo CreatePartStudioTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Asynchronously export a Part Studio to another format.



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
    bTTranslateFormatParams := *openapiclient.NewBTTranslateFormatParams("FormatName_example") // BTTranslateFormatParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePartStudioFeature

> BTFeatureApiBase1430 DeletePartStudioFeature(ctx, did, wid, eid, fid).Execute()

Delete a Part Studio feature.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvalFeatureScript

> BTFeatureScriptEvalResponse1859 EvalFeatureScript(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).BTFeatureScriptEvalCall2377(bTFeatureScriptEvalCall2377).Execute()

Evaluate the FeatureScript snippet for a Part Studio.

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
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    bTFeatureScriptEvalCall2377 := *openapiclient.NewBTFeatureScriptEvalCall2377() // BTFeatureScriptEvalCall2377 |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvalFeatureScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **bTFeatureScriptEvalCall2377** | [**BTFeatureScriptEvalCall2377**](BTFeatureScriptEvalCall2377.md) |  | 

### Return type

[**BTFeatureScriptEvalResponse1859**](BTFeatureScriptEvalResponse1859.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportParasolid

> ExportParasolid(ctx, did, wvm, wvmid, eid).PartIds(partIds).Version(version).IncludeExportIds(includeExportIds).Configuration(configuration).LinkDocumentId(linkDocumentId).BinaryExport(binaryExport).Execute()

Synchronously export a Part Studio to a Parasolid file.



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
    partIds := "partIds_example" // string | IDs of the parts to retrieve. Use comma-separated IDs for multiple parts (example: partIds=JHK,JHD). (optional)
    version := "version_example" // string | Parasolid version (optional) (default to "0")
    includeExportIds := true // bool | Whether topology ids should be exported as parasolid attributes (optional) (default to false)
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)
    binaryExport := true // bool | Whether to use binary parasolid format instead of text (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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




 **partIds** | **string** | IDs of the parts to retrieve. Use comma-separated IDs for multiple parts (example: partIds&#x3D;JHK,JHD). | 
 **version** | **string** | Parasolid version | [default to &quot;0&quot;]
 **includeExportIds** | **bool** | Whether topology ids should be exported as parasolid attributes | [default to false]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 
 **binaryExport** | **bool** | Whether to use binary parasolid format instead of text | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPartStudioGltf

> GlTF ExportPartStudioGltf(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).OutputSeparateFaceNodes(outputSeparateFaceNodes).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).Execute()

Synchronously export a Part Studio to a glTF file.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartStudioApi.ExportPartStudioGltf(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).OutputSeparateFaceNodes(outputSeparateFaceNodes).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartStudioApi.ExportPartStudioGltf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPartStudioGltf`: GlTF
    fmt.Fprintf(os.Stdout, "Response from `PartStudioApi.ExportPartStudioGltf`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiExportPartStudioGltfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
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

[**GlTF**](GlTF.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: model/gltf+json;charset=UTF-8;qs=0.08, model/gltf-binary;qs=0.08

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPartStudioStl

> ExportPartStudioStl(ctx, did, wvm, wvmid, eid).PartIds(partIds).Mode(mode).Grouping(grouping).Scale(scale).Units(units).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).MinFacetWidth(minFacetWidth).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Synchronously export a Part Studio to an STL file.



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
    partIds := "partIds_example" // string | IDs of the parts to retrieve. Use comma-separated IDs for multiple parts (example: partIds=JHK,JHD). (optional)
    mode := "mode_example" // string | Type of file: text, binary (optional) (default to "text")
    grouping := true // bool | Whether parts should be exported as a group or individually in a .zip file (optional) (default to true)
    scale := float64(1.2) // float64 | Scale for measurements. (optional) (default to 1)
    units := "units_example" // string | Units for the element: `METER` | `CENTIMETER` | `MILLIMETER` | `INCH` | `FOOT` | `YARD` (optional) (default to "inch")
    angleTolerance := float64(1.2) // float64 | Angle tolerance (in radians). This specifies the limit on the sum of the angular deviations of a tessellation chord from the tangent vectors at two chord endpoints. The specified value must be less than PI/2. This parameter currently has a default value chosen based on the complexity of the parts being tessellated. (optional)
    chordTolerance := float64(1.2) // float64 | Chord tolerance (in meters). This specifies the limit on the maximum deviation of a tessellation chord from the true surface/edge. This parameter currently has a default value chosen based on the size and complexity of the parts being tessellated. (optional)
    maxFacetWidth := float64(1.2) // float64 | Max facet width. This specifies the limit on the size of any side of a tessellation facet. (optional)
    minFacetWidth := float64(1.2) // float64 | Max facet width. This specifies the limit on the size of any side of a tessellation facet. (optional)
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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




 **partIds** | **string** | IDs of the parts to retrieve. Use comma-separated IDs for multiple parts (example: partIds&#x3D;JHK,JHD). | 
 **mode** | **string** | Type of file: text, binary | [default to &quot;text&quot;]
 **grouping** | **bool** | Whether parts should be exported as a group or individually in a .zip file | [default to true]
 **scale** | **float64** | Scale for measurements. | [default to 1]
 **units** | **string** | Units for the element: &#x60;METER&#x60; | &#x60;CENTIMETER&#x60; | &#x60;MILLIMETER&#x60; | &#x60;INCH&#x60; | &#x60;FOOT&#x60; | &#x60;YARD&#x60; | [default to &quot;inch&quot;]
 **angleTolerance** | **float64** | Angle tolerance (in radians). This specifies the limit on the sum of the angular deviations of a tessellation chord from the tangent vectors at two chord endpoints. The specified value must be less than PI/2. This parameter currently has a default value chosen based on the complexity of the parts being tessellated. | 
 **chordTolerance** | **float64** | Chord tolerance (in meters). This specifies the limit on the maximum deviation of a tessellation chord from the true surface/edge. This parameter currently has a default value chosen based on the size and complexity of the parts being tessellated. | 
 **maxFacetWidth** | **float64** | Max facet width. This specifies the limit on the size of any side of a tessellation facet. | 
 **minFacetWidth** | **float64** | Max facet width. This specifies the limit on the size of any side of a tessellation facet. | 
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureScriptRepresentation

> BTPModule234 GetFeatureScriptRepresentation(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).Execute()

Get the FeatureScript representation of a Part Studio.

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
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureScriptRepresentationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 

### Return type

[**BTPModule234**](BTPModule234.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureScriptTable

> BTApiTableList1223 GetFeatureScriptTable(ctx, did, wvm, wvmid, eid).TableType(tableType).Configuration(configuration).TableNamespace(tableNamespace).TableParameters(tableParameters).PartId(partId).LinkDocumentId(linkDocumentId).Execute()

Compute and return a FeatureScript table for a Part Studio.

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
    tableType := "tableType_example" // string | May be any standard table type (i.e., `holeTable` or `cutlistTable`) or custom table type name defined in FeatureScript.  If using a custom table type, `tableNamespace` must also be specified.
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    tableNamespace := "tableNamespace_example" // string | Namespace of the custom table in FS. Must be in the form of:  * `e{eid}::m{mid}` if the FS and PS tabs are in the same workspace.  * `d{did}::v{vid}::e{eid}::m{mid}` if the tabs are in different workspaces.   Obtain the microversion id (`{mid}`) with [this endpoint](#/Document/getElementsInDocument) called on the FS tab.  Leave blank if using a standard table. Required if using a custom `tableType`. (optional)
    tableParameters := "tableParameters_example" // string | Include all parameters for the table. i.e., `customBool=false` (optional)
    partId := "partId_example" // string | ID of the part to retrieve. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureScriptTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **tableType** | **string** | May be any standard table type (i.e., &#x60;holeTable&#x60; or &#x60;cutlistTable&#x60;) or custom table type name defined in FeatureScript.  If using a custom table type, &#x60;tableNamespace&#x60; must also be specified. | 
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **tableNamespace** | **string** | Namespace of the custom table in FS. Must be in the form of:  * &#x60;e{eid}::m{mid}&#x60; if the FS and PS tabs are in the same workspace.  * &#x60;d{did}::v{vid}::e{eid}::m{mid}&#x60; if the tabs are in different workspaces.   Obtain the microversion id (&#x60;{mid}&#x60;) with [this endpoint](#/Document/getElementsInDocument) called on the FS tab.  Leave blank if using a standard table. Required if using a custom &#x60;tableType&#x60;. | 
 **tableParameters** | **string** | Include all parameters for the table. i.e., &#x60;customBool&#x3D;false&#x60; | 
 **partId** | **string** | ID of the part to retrieve. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**BTApiTableList1223**](BTApiTableList1223.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioBodyDetails

> BTExportModelBodiesResponse734 GetPartStudioBodyDetails(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartIds(partIds).IncludeSurfaces(includeSurfaces).IncludeCompositeParts(includeCompositeParts).IncludeGeometricData(includeGeometricData).Execute()

Get the body details for a Part Studio.



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
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    partIds := []string{"Inner_example"} // []string | If specified, the response will only include body details for the specific parts as indicated here by their corresponding Id (optional)
    includeSurfaces := true // bool | Whether or not surfaces should be included in the response. (optional) (default to false)
    includeCompositeParts := true // bool | Whether or not composite parts should be included in the response. (optional) (default to false)
    includeGeometricData := true // bool | Whether or not geometric data should be included in the response. (optional) (default to true)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioBodyDetails(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartIds(partIds).IncludeSurfaces(includeSurfaces).IncludeCompositeParts(includeCompositeParts).IncludeGeometricData(includeGeometricData).Execute()
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
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioBodyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **partIds** | **[]string** | If specified, the response will only include body details for the specific parts as indicated here by their corresponding Id | 
 **includeSurfaces** | **bool** | Whether or not surfaces should be included in the response. | [default to false]
 **includeCompositeParts** | **bool** | Whether or not composite parts should be included in the response. | [default to false]
 **includeGeometricData** | **bool** | Whether or not geometric data should be included in the response. | [default to true]

### Return type

[**BTExportModelBodiesResponse734**](BTExportModelBodiesResponse734.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioBoundingBoxes

> BTBoundingBoxInfo GetPartStudioBoundingBoxes(ctx, did, wvm, wvmid, eid).IncludeHidden(includeHidden).IncludeWireBodies(includeWireBodies).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Get the bounding boxes for a Part Studio.



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
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**BTBoundingBoxInfo**](BTBoundingBoxInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioEdges

> BTExportTessellatedEdgesResponse327 GetPartStudioEdges(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).EdgeId(edgeId).Execute()

Get a list of all edges in a Part Studio.



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
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    partId := []string{"Inner_example"} // []string |  (optional)
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    precomputedLevelOfDetail := "precomputedLevelOfDetail_example" // string |  (optional)
    edgeId := []string{"Inner_example"} // []string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFaces

> BTExportTessellatedFacesResponse898 GetPartStudioFaces(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputIndexTable(outputIndexTable).OutputErrorFaces(outputErrorFaces).CombineCompositePartConstituents(combineCompositePartConstituents).Execute()

Get a list of all faces in a Part Studio.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioFacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFeatureSpecs

> BTFeatureSpecsResponse664 GetPartStudioFeatureSpecs(ctx, did, wvm, wvmid, eid).Execute()

Get the specs for a Part Studio feature.



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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFeatures

> BTFeatureListResponse2457 GetPartStudioFeatures(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).IncludeGeometryIds(includeGeometryIds).FeatureId(featureId).NoSketchGeometry(noSketchGeometry).Execute()

Get a list of features instantiated in the Part Studio.



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
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    includeGeometryIds := true // bool | If true, include the underlying geometry IDs in the feature definition. (optional) (default to true)
    featureId := []string{"Inner_example"} // []string | ID of a feature; repeat query param to add more than one (optional)
    noSketchGeometry := true // bool | Whether or not to output simple sketch info without geometry (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartStudioApi.GetPartStudioFeatures(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).IncludeGeometryIds(includeGeometryIds).FeatureId(featureId).NoSketchGeometry(noSketchGeometry).Execute()
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
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **includeGeometryIds** | **bool** | If true, include the underlying geometry IDs in the feature definition. | [default to true]
 **featureId** | **[]string** | ID of a feature; repeat query param to add more than one | 
 **noSketchGeometry** | **bool** | Whether or not to output simple sketch info without geometry | [default to false]

### Return type

[**BTFeatureListResponse2457**](BTFeatureListResponse2457.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioMassProperties

> BTMassPropertiesBulkInfo GetPartStudioMassProperties(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).PartId(partId).MassAsGroup(massAsGroup).UseMassPropertyOverrides(useMassPropertyOverrides).Execute()

Get the mass properties for a Part Studio.



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
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    partId := []string{"Inner_example"} // []string |  (optional)
    massAsGroup := true // bool | If true, specified parts will be evaluated as a single object instead of individually (optional) (default to true)
    useMassPropertyOverrides := true // bool | If true, use the user mass property overrides when calculated mass properties (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioMassPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **partId** | **[]string** |  | 
 **massAsGroup** | **bool** | If true, specified parts will be evaluated as a single object instead of individually | [default to true]
 **useMassPropertyOverrides** | **bool** | If true, use the user mass property overrides when calculated mass properties | [default to false]

### Return type

[**BTMassPropertiesBulkInfo**](BTMassPropertiesBulkInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioNamedViews

> BTNamedViewsInfo GetPartStudioNamedViews(ctx, did, eid).LinkDocumentId(linkDocumentId).SkipPerspective(skipPerspective).IncludeSectionCutViews(includeSectionCutViews).Execute()

Get a list of all named views that exist in the Part Studio.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioShadedViews

> BTShadedViewsInfo GetPartStudioShadedViews(ctx, did, wvm, wvmid, eid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).IncludeWires(includeWires).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Get a list of shaded views for a Part Studio.

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
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**BTShadedViewsInfo**](BTShadedViewsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslateIds

> BTIdTranslationInfo TranslateIds(ctx, did, wvm, wvmid, eid).BTIdTranslationParams(bTIdTranslationParams).Execute()

Find corresponding deterministic IDs from a source document microversion at the target version.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatures

> BTUpdateFeaturesResponse1333 UpdateFeatures(ctx, did, wid, eid).BTUpdateFeaturesCall1748(bTUpdateFeaturesCall1748).Execute()

Update multiple features in a Part Studio



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartStudioFeature

> BTFeatureDefinitionResponse1617 UpdatePartStudioFeature(ctx, did, wid, eid, fid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()

Update the definition of a Part Studio feature.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRollback

> BTSetFeatureRollbackResponse1042 UpdateRollback(ctx, did, wid, eid).Body(body).Execute()

Move the Feature List rollback bar in the Part Studio.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

