# \AssemblyApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFeature**](AssemblyApi.md#AddFeature) | **Post** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Add a feature to the assembly feature list.
[**CreateAssembly**](AssemblyApi.md#CreateAssembly) | **Post** /assemblies/d/{did}/w/{wid} | Create a new assembly tab in the document.
[**CreateAssemblyExportGltf**](AssemblyApi.md#CreateAssemblyExportGltf) | **Post** /assemblies/d/{did}/{wv}/{wvid}/e/{eid}/export/gltf | Export the assembly to glTF.
[**CreateAssemblyExportObj**](AssemblyApi.md#CreateAssemblyExportObj) | **Post** /assemblies/d/{did}/{wv}/{wvid}/e/{eid}/export/obj | Export the assembly to OBJ.
[**CreateAssemblyExportSolidworks**](AssemblyApi.md#CreateAssemblyExportSolidworks) | **Post** /assemblies/d/{did}/{wv}/{wvid}/e/{eid}/export/solidworks | Export the assembly to Solidworks.
[**CreateAssemblyExportStep**](AssemblyApi.md#CreateAssemblyExportStep) | **Post** /assemblies/d/{did}/{wv}/{wvid}/e/{eid}/export/step | Export the assembly to STEP.
[**CreateInstance**](AssemblyApi.md#CreateInstance) | **Post** /assemblies/d/{did}/w/{wid}/e/{eid}/instances | Insert an instance of a part, sketch, assembly, or Part Studio into an assembly.
[**DeleteFeature**](AssemblyApi.md#DeleteFeature) | **Delete** /assemblies/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Delete a feature from an assembly.
[**DeleteInstance**](AssemblyApi.md#DeleteInstance) | **Delete** /assemblies/d/{did}/w/{wid}/e/{eid}/instance/nodeid/{nid} | Delete an instance of an assembly.
[**GetAssemblyBoundingBoxes**](AssemblyApi.md#GetAssemblyBoundingBoxes) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/boundingboxes | Get bounding box information for the specified assembly.
[**GetAssemblyDefinition**](AssemblyApi.md#GetAssemblyDefinition) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid} | Get definition information for the specified assembly.
[**GetAssemblyMassProperties**](AssemblyApi.md#GetAssemblyMassProperties) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/massproperties | Get the mass properties for the assembly.
[**GetAssemblyShadedViews**](AssemblyApi.md#GetAssemblyShadedViews) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/shadedviews | Get an array of shaded view images for the document.
[**GetBillOfMaterials**](AssemblyApi.md#GetBillOfMaterials) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/bom | Get the Bill Of Materials (BOM) content for the specified assembly.
[**GetDisplayStates**](AssemblyApi.md#GetDisplayStates) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/displaystates | Get a list of display states for the specified assembly.
[**GetExplodedViews**](AssemblyApi.md#GetExplodedViews) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/explodedviews | Get a list of exploded views for the specified assembly.
[**GetFeatureSpecs**](AssemblyApi.md#GetFeatureSpecs) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | Get the feature spec definitions for an assembly.
[**GetFeatures**](AssemblyApi.md#GetFeatures) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Get the definitions of the specified features in an assembly.
[**GetMateValues**](AssemblyApi.md#GetMateValues) | **Get** /assemblies/d/{did}/{wv}/{wvid}/e/{eid}/matevalues | Get a list of mate values in the specified assembly.
[**GetNamedPositions**](AssemblyApi.md#GetNamedPositions) | **Get** /assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/namedpositions | Get a list of all named positions for the assembly.
[**GetNamedViews**](AssemblyApi.md#GetNamedViews) | **Get** /assemblies/d/{did}/e/{eid}/namedViews | Get the view data for all named views for the specified element.
[**GetOrCreateBillOfMaterialsElement**](AssemblyApi.md#GetOrCreateBillOfMaterialsElement) | **Post** /assemblies/d/{did}/w/{wid}/e/{eid}/bomelement | Gets the Bill Of Materials (BOM) for the specified assembly, or creates a BOM if none exist.
[**InsertTransformedInstances**](AssemblyApi.md#InsertTransformedInstances) | **Post** /assemblies/d/{did}/w/{wid}/e/{eid}/transformedinstances | Create new instances with transformation.
[**Modify**](AssemblyApi.md#Modify) | **Post** /assemblies/d/{did}/w/{wid}/e/{eid}/modify | Modify an assembly.
[**TransformOccurrences**](AssemblyApi.md#TransformOccurrences) | **Post** /assemblies/d/{did}/w/{wid}/e/{eid}/occurrencetransforms | Transform a list of assembly occurrences.
[**TranslateFormat**](AssemblyApi.md#TranslateFormat) | **Post** /assemblies/d/{did}/{wv}/{wvid}/e/{eid}/translations | Export the assembly to another format.
[**UpdateFeature**](AssemblyApi.md#UpdateFeature) | **Post** /assemblies/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Update an existing feature for an Assembly.
[**UpdateMateValues**](AssemblyApi.md#UpdateMateValues) | **Post** /assemblies/d/{did}/w/{wid}/e/{eid}/matevalues | Update mate values for the given mates in the specified assembly.



## AddFeature

> BTFeatureDefinitionResponse1617 AddFeature(ctx, did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()

Add a feature to the assembly feature list.

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
    resp, r, err := apiClient.AssemblyApi.AddFeature(context.Background(), did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.AddFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddFeature`: BTFeatureDefinitionResponse1617
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.AddFeature`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiAddFeatureRequest struct via the builder pattern


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


## CreateAssembly

> BTDocumentElementInfo CreateAssembly(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create a new assembly tab in the document.

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
    bTModelElementParams := *openapiclient.NewBTModelElementParams() // BTModelElementParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.CreateAssembly(context.Background(), did, wid).BTModelElementParams(bTModelElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.CreateAssembly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssembly`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.CreateAssembly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssemblyRequest struct via the builder pattern


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


## CreateAssemblyExportGltf

> BTTranslationRequestInfo CreateAssemblyExportGltf(ctx, did, wv, wvid, eid).BTBGltfExportParams(bTBGltfExportParams).Execute()

Export the assembly to glTF.



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
    resp, r, err := apiClient.AssemblyApi.CreateAssemblyExportGltf(context.Background(), did, wv, wvid, eid).BTBGltfExportParams(bTBGltfExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.CreateAssemblyExportGltf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssemblyExportGltf`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.CreateAssemblyExportGltf`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateAssemblyExportGltfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBGltfExportParams** | [**BTBGltfExportParams**](BTBGltfExportParams.md) |  | 

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


## CreateAssemblyExportObj

> BTTranslationRequestInfo CreateAssemblyExportObj(ctx, did, wv, wvid, eid).BTBObjExportParams(bTBObjExportParams).Execute()

Export the assembly to OBJ.



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
    resp, r, err := apiClient.AssemblyApi.CreateAssemblyExportObj(context.Background(), did, wv, wvid, eid).BTBObjExportParams(bTBObjExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.CreateAssemblyExportObj``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssemblyExportObj`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.CreateAssemblyExportObj`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateAssemblyExportObjRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBObjExportParams** | [**BTBObjExportParams**](BTBObjExportParams.md) |  | 

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


## CreateAssemblyExportSolidworks

> BTTranslationRequestInfo CreateAssemblyExportSolidworks(ctx, did, wv, wvid, eid).BTBSolidworksExportParams(bTBSolidworksExportParams).Execute()

Export the assembly to Solidworks.



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
    resp, r, err := apiClient.AssemblyApi.CreateAssemblyExportSolidworks(context.Background(), did, wv, wvid, eid).BTBSolidworksExportParams(bTBSolidworksExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.CreateAssemblyExportSolidworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssemblyExportSolidworks`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.CreateAssemblyExportSolidworks`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateAssemblyExportSolidworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBSolidworksExportParams** | [**BTBSolidworksExportParams**](BTBSolidworksExportParams.md) |  | 

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


## CreateAssemblyExportStep

> BTTranslationRequestInfo CreateAssemblyExportStep(ctx, did, wv, wvid, eid).BTBStepExportParams(bTBStepExportParams).Execute()

Export the assembly to STEP.



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
    resp, r, err := apiClient.AssemblyApi.CreateAssemblyExportStep(context.Background(), did, wv, wvid, eid).BTBStepExportParams(bTBStepExportParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.CreateAssemblyExportStep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssemblyExportStep`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.CreateAssemblyExportStep`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateAssemblyExportStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTBStepExportParams** | [**BTBStepExportParams**](BTBStepExportParams.md) |  | 

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


## CreateInstance

> map[string]interface{} CreateInstance(ctx, did, wid, eid).BTAssemblyInstanceDefinitionParams(bTAssemblyInstanceDefinitionParams).Execute()

Insert an instance of a part, sketch, assembly, or Part Studio into an assembly.



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
    bTAssemblyInstanceDefinitionParams := *openapiclient.NewBTAssemblyInstanceDefinitionParams("DocumentId_example") // BTAssemblyInstanceDefinitionParams |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.CreateInstance(context.Background(), did, wid, eid).BTAssemblyInstanceDefinitionParams(bTAssemblyInstanceDefinitionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.CreateInstance`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAssemblyInstanceDefinitionParams** | [**BTAssemblyInstanceDefinitionParams**](BTAssemblyInstanceDefinitionParams.md) |  | 

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


## DeleteFeature

> BTFeatureApiBase1430 DeleteFeature(ctx, did, wid, eid, fid).Execute()

Delete a feature from an assembly.

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
    fid := "fid_example" // string | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.DeleteFeature(context.Background(), did, wid, eid, fid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.DeleteFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFeature`: BTFeatureApiBase1430
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.DeleteFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureRequest struct via the builder pattern


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


## DeleteInstance

> map[string]interface{} DeleteInstance(ctx, did, eid, wid, nid).Execute()

Delete an instance of an assembly.

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
    eid := "eid_example" // string | 
    wid := "wid_example" // string | 
    nid := "nid_example" // string | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.DeleteInstance(context.Background(), did, eid, wid, nid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.DeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 
**nid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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


## GetAssemblyBoundingBoxes

> BTBoundingBoxInfo GetAssemblyBoundingBoxes(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).IncludeHidden(includeHidden).DisplayStateId(displayStateId).NamedPositionId(namedPositionId).IncludeSketches(includeSketches).Execute()

Get bounding box information for the specified assembly.

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
    explodedViewId := "explodedViewId_example" // string |  (optional)
    includeHidden := true // bool |  (optional)
    displayStateId := "displayStateId_example" // string | Call the [getDisplayStates](#/Assembly/getDisplayStates) endpoint to get display state ID(s). (optional)
    namedPositionId := "namedPositionId_example" // string |  (optional)
    includeSketches := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetAssemblyBoundingBoxes(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).IncludeHidden(includeHidden).DisplayStateId(displayStateId).NamedPositionId(namedPositionId).IncludeSketches(includeSketches).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetAssemblyBoundingBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssemblyBoundingBoxes`: BTBoundingBoxInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetAssemblyBoundingBoxes`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssemblyBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **explodedViewId** | **string** |  | 
 **includeHidden** | **bool** |  | 
 **displayStateId** | **string** | Call the [getDisplayStates](#/Assembly/getDisplayStates) endpoint to get display state ID(s). | 
 **namedPositionId** | **string** |  | 
 **includeSketches** | **bool** |  | [default to false]

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


## GetAssemblyDefinition

> BTAssemblyDefinitionInfo GetAssemblyDefinition(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).IncludeMateFeatures(includeMateFeatures).IncludeNonSolids(includeNonSolids).IncludeMateConnectors(includeMateConnectors).ExcludeSuppressed(excludeSuppressed).Execute()

Get definition information for the specified assembly.



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
    explodedViewId := "explodedViewId_example" // string |  (optional)
    includeMateFeatures := true // bool |  (optional) (default to false)
    includeNonSolids := true // bool |  (optional) (default to false)
    includeMateConnectors := true // bool |  (optional) (default to false)
    excludeSuppressed := true // bool | Whether or not to exclude suppressed instances/mate features in response (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetAssemblyDefinition(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).IncludeMateFeatures(includeMateFeatures).IncludeNonSolids(includeNonSolids).IncludeMateConnectors(includeMateConnectors).ExcludeSuppressed(excludeSuppressed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetAssemblyDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssemblyDefinition`: BTAssemblyDefinitionInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetAssemblyDefinition`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssemblyDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **explodedViewId** | **string** |  | 
 **includeMateFeatures** | **bool** |  | [default to false]
 **includeNonSolids** | **bool** |  | [default to false]
 **includeMateConnectors** | **bool** |  | [default to false]
 **excludeSuppressed** | **bool** | Whether or not to exclude suppressed instances/mate features in response | [default to false]

### Return type

[**BTAssemblyDefinitionInfo**](BTAssemblyDefinitionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssemblyMassProperties

> BTMassPropertiesInfo GetAssemblyMassProperties(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()

Get the mass properties for the assembly.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetAssemblyMassProperties(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetAssemblyMassProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssemblyMassProperties`: BTMassPropertiesInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetAssemblyMassProperties`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssemblyMassPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]

### Return type

[**BTMassPropertiesInfo**](BTMassPropertiesInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssemblyShadedViews

> BTShadedViewsInfo GetAssemblyShadedViews(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).IncludeWires(includeWires).DisplayStateId(displayStateId).NamedPositionId(namedPositionId).Execute()

Get an array of shaded view images for the document.

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
    explodedViewId := "explodedViewId_example" // string |  (optional)
    viewMatrix := "viewMatrix_example" // string |  (optional) (default to "front")
    outputHeight := int32(56) // int32 |  (optional) (default to 500)
    outputWidth := int32(56) // int32 |  (optional) (default to 500)
    pixelSize := float64(1.2) // float64 |  (optional) (default to 0.003)
    edges := "edges_example" // string |  (optional) (default to "show")
    showAllParts := true // bool |  (optional) (default to false)
    includeSurfaces := true // bool |  (optional) (default to true)
    useAntiAliasing := true // bool |  (optional) (default to false)
    includeWires := true // bool |  (optional) (default to false)
    displayStateId := "displayStateId_example" // string | Call the [getDisplayStates](#/Assembly/getDisplayStates) endpoint to get display state ID(s). (optional)
    namedPositionId := "namedPositionId_example" // string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetAssemblyShadedViews(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).IncludeWires(includeWires).DisplayStateId(displayStateId).NamedPositionId(namedPositionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetAssemblyShadedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssemblyShadedViews`: BTShadedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetAssemblyShadedViews`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetAssemblyShadedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **explodedViewId** | **string** |  | 
 **viewMatrix** | **string** |  | [default to &quot;front&quot;]
 **outputHeight** | **int32** |  | [default to 500]
 **outputWidth** | **int32** |  | [default to 500]
 **pixelSize** | **float64** |  | [default to 0.003]
 **edges** | **string** |  | [default to &quot;show&quot;]
 **showAllParts** | **bool** |  | [default to false]
 **includeSurfaces** | **bool** |  | [default to true]
 **useAntiAliasing** | **bool** |  | [default to false]
 **includeWires** | **bool** |  | [default to false]
 **displayStateId** | **string** | Call the [getDisplayStates](#/Assembly/getDisplayStates) endpoint to get display state ID(s). | 
 **namedPositionId** | **string** |  | 

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


## GetBillOfMaterials

> BTBillOfMaterialsInfo GetBillOfMaterials(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).BomColumnIds(bomColumnIds).Indented(indented).MultiLevel(multiLevel).GenerateIfAbsent(generateIfAbsent).TemplateId(templateId).IncludeExcluded(includeExcluded).OnlyVisibleColumns(onlyVisibleColumns).IgnoreSubassemblyBomBehavior(ignoreSubassemblyBomBehavior).IncludeItemMicroversions(includeItemMicroversions).IncludeTopLevelAssemblyRow(includeTopLevelAssemblyRow).Thumbnail(thumbnail).RespectSubassemblyBomBehavior(respectSubassemblyBomBehavior).Execute()

Get the Bill Of Materials (BOM) content for the specified assembly.



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
    bomColumnIds := []string{"Inner_example"} // []string | Ids of the columns to include, or all columns if empty. BOM column ids correspond to metadata property ids. (optional)
    indented := true // bool | Return the Structured BOM table with all rows collapsed, otherwise returns the Flattened BOM. (optional) (default to true)
    multiLevel := true // bool | Return the Structured BOM table with all rows expanded. Ignored if indented is false. (optional) (default to false)
    generateIfAbsent := true // bool | Return the BOM table data even if the BOM does not exist. If this is false and the BOM does not exist, a 404 status code will be returned. This option is highly recommended. (optional) (default to false)
    templateId := "templateId_example" // string | The id of the BOM table template to use when generating the table. (optional)
    includeExcluded := true // bool | Include items that have been excluded from the BOM table. (optional)
    onlyVisibleColumns := true // bool | Only return data for visible columns, instead of all possible columns. (optional)
    ignoreSubassemblyBomBehavior := true // bool | Ignore the 'Subassembly BOM behavior' property when constructing the BOM table. (optional)
    includeItemMicroversions := true // bool | Include element microversions and version metadata microversions in the JSON. (optional) (default to false)
    includeTopLevelAssemblyRow := true // bool | Include top-level assembly row when constructing the BOM table. (optional) (default to false)
    thumbnail := true // bool | Return thumbnail info (optional) (default to false)
    respectSubassemblyBomBehavior := true // bool | Force respect subassembly BOM Behavior Property (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetBillOfMaterials(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).BomColumnIds(bomColumnIds).Indented(indented).MultiLevel(multiLevel).GenerateIfAbsent(generateIfAbsent).TemplateId(templateId).IncludeExcluded(includeExcluded).OnlyVisibleColumns(onlyVisibleColumns).IgnoreSubassemblyBomBehavior(ignoreSubassemblyBomBehavior).IncludeItemMicroversions(includeItemMicroversions).IncludeTopLevelAssemblyRow(includeTopLevelAssemblyRow).Thumbnail(thumbnail).RespectSubassemblyBomBehavior(respectSubassemblyBomBehavior).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetBillOfMaterials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillOfMaterials`: BTBillOfMaterialsInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetBillOfMaterials`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetBillOfMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **bomColumnIds** | **[]string** | Ids of the columns to include, or all columns if empty. BOM column ids correspond to metadata property ids. | 
 **indented** | **bool** | Return the Structured BOM table with all rows collapsed, otherwise returns the Flattened BOM. | [default to true]
 **multiLevel** | **bool** | Return the Structured BOM table with all rows expanded. Ignored if indented is false. | [default to false]
 **generateIfAbsent** | **bool** | Return the BOM table data even if the BOM does not exist. If this is false and the BOM does not exist, a 404 status code will be returned. This option is highly recommended. | [default to false]
 **templateId** | **string** | The id of the BOM table template to use when generating the table. | 
 **includeExcluded** | **bool** | Include items that have been excluded from the BOM table. | 
 **onlyVisibleColumns** | **bool** | Only return data for visible columns, instead of all possible columns. | 
 **ignoreSubassemblyBomBehavior** | **bool** | Ignore the &#39;Subassembly BOM behavior&#39; property when constructing the BOM table. | 
 **includeItemMicroversions** | **bool** | Include element microversions and version metadata microversions in the JSON. | [default to false]
 **includeTopLevelAssemblyRow** | **bool** | Include top-level assembly row when constructing the BOM table. | [default to false]
 **thumbnail** | **bool** | Return thumbnail info | [default to false]
 **respectSubassemblyBomBehavior** | **bool** | Force respect subassembly BOM Behavior Property | [default to false]

### Return type

[**BTBillOfMaterialsInfo**](BTBillOfMaterialsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisplayStates

> []BTDisplayStateInfo GetDisplayStates(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Execute()

Get a list of display states for the specified assembly.

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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetDisplayStates(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetDisplayStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDisplayStates`: []BTDisplayStateInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetDisplayStates`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetDisplayStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**[]BTDisplayStateInfo**](BTDisplayStateInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExplodedViews

> []BTViewFeatureBaseInfo GetExplodedViews(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()

Get a list of exploded views for the specified assembly.

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
    explodedViewId := "explodedViewId_example" // string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetExplodedViews(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetExplodedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExplodedViews`: []BTViewFeatureBaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetExplodedViews`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetExplodedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **explodedViewId** | **string** |  | 

### Return type

[**[]BTViewFeatureBaseInfo**](BTViewFeatureBaseInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureSpecs

> BTFeatureSpecsResponse664 GetFeatureSpecs(ctx, did, wvm, wvmid, eid).Execute()

Get the feature spec definitions for an assembly.

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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetFeatureSpecs(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetFeatureSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureSpecs`: BTFeatureSpecsResponse664
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetFeatureSpecs`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFeatureSpecsRequest struct via the builder pattern


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


## GetFeatures

> BTAssemblyFeatureListResponse1174 GetFeatures(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).FeatureId(featureId).Execute()

Get the definitions of the specified features in an assembly.

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
    explodedViewId := "explodedViewId_example" // string |  (optional)
    featureId := []string{"Inner_example"} // []string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetFeatures(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).FeatureId(featureId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatures`: BTAssemblyFeatureListResponse1174
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetFeatures`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **explodedViewId** | **string** |  | 
 **featureId** | **[]string** |  | 

### Return type

[**BTAssemblyFeatureListResponse1174**](BTAssemblyFeatureListResponse1174.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMateValues

> BTAssemblyMateValuesInfo GetMateValues(ctx, did, wv, wvid, eid).Execute()

Get a list of mate values in the specified assembly.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetMateValues(context.Background(), did, wv, wvid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetMateValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMateValues`: BTAssemblyMateValuesInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetMateValues`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetMateValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTAssemblyMateValuesInfo**](BTAssemblyMateValuesInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedPositions

> []BTViewFeatureBaseInfo GetNamedPositions(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()

Get a list of all named positions for the assembly.

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
    explodedViewId := "explodedViewId_example" // string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.GetNamedPositions(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetNamedPositions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamedPositions`: []BTViewFeatureBaseInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetNamedPositions`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetNamedPositionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | [default to &quot;&quot;]
 **explodedViewId** | **string** |  | 

### Return type

[**[]BTViewFeatureBaseInfo**](BTViewFeatureBaseInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedViews

> BTNamedViewsInfo GetNamedViews(ctx, did, eid).LinkDocumentId(linkDocumentId).SkipPerspective(skipPerspective).IncludeSectionCutViews(includeSectionCutViews).Execute()

Get the view data for all named views for the specified element.

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
    resp, r, err := apiClient.AssemblyApi.GetNamedViews(context.Background(), did, eid).LinkDocumentId(linkDocumentId).SkipPerspective(skipPerspective).IncludeSectionCutViews(includeSectionCutViews).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetNamedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamedViews`: BTNamedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetNamedViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamedViewsRequest struct via the builder pattern


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


## GetOrCreateBillOfMaterialsElement

> BTDocumentElementInfo GetOrCreateBillOfMaterialsElement(ctx, did, wid, eid).Execute()

Gets the Bill Of Materials (BOM) for the specified assembly, or creates a BOM if none exist.

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
    resp, r, err := apiClient.AssemblyApi.GetOrCreateBillOfMaterialsElement(context.Background(), did, wid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.GetOrCreateBillOfMaterialsElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrCreateBillOfMaterialsElement`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.GetOrCreateBillOfMaterialsElement`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetOrCreateBillOfMaterialsElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertTransformedInstances

> BTAssemblyInsertTransformedInstancesResponse InsertTransformedInstances(ctx, did, eid, wid).BTAssemblyTransformedInstancesDefinitionParams(bTAssemblyTransformedInstancesDefinitionParams).Execute()

Create new instances with transformation.

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
    eid := "eid_example" // string | 
    wid := "wid_example" // string | 
    bTAssemblyTransformedInstancesDefinitionParams := *openapiclient.NewBTAssemblyTransformedInstancesDefinitionParams() // BTAssemblyTransformedInstancesDefinitionParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.InsertTransformedInstances(context.Background(), did, eid, wid).BTAssemblyTransformedInstancesDefinitionParams(bTAssemblyTransformedInstancesDefinitionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.InsertTransformedInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InsertTransformedInstances`: BTAssemblyInsertTransformedInstancesResponse
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.InsertTransformedInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInsertTransformedInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAssemblyTransformedInstancesDefinitionParams** | [**BTAssemblyTransformedInstancesDefinitionParams**](BTAssemblyTransformedInstancesDefinitionParams.md) |  | 

### Return type

[**BTAssemblyInsertTransformedInstancesResponse**](BTAssemblyInsertTransformedInstancesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Modify

> map[string]interface{} Modify(ctx, did, wid, eid).LinkDocumentId(linkDocumentId).BTAssemblyModificationParams(bTAssemblyModificationParams).Execute()

Modify an assembly.



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
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    bTAssemblyModificationParams := *openapiclient.NewBTAssemblyModificationParams() // BTAssemblyModificationParams |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.Modify(context.Background(), did, wid, eid).LinkDocumentId(linkDocumentId).BTAssemblyModificationParams(bTAssemblyModificationParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.Modify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Modify`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.Modify`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **bTAssemblyModificationParams** | [**BTAssemblyModificationParams**](BTAssemblyModificationParams.md) |  | 

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


## TransformOccurrences

> map[string]interface{} TransformOccurrences(ctx, did, eid, wid).BTAssemblyTransformDefinitionParams(bTAssemblyTransformDefinitionParams).Execute()

Transform a list of assembly occurrences.

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
    eid := "eid_example" // string | 
    wid := "wid_example" // string | 
    bTAssemblyTransformDefinitionParams := *openapiclient.NewBTAssemblyTransformDefinitionParams() // BTAssemblyTransformDefinitionParams |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.TransformOccurrences(context.Background(), did, eid, wid).BTAssemblyTransformDefinitionParams(bTAssemblyTransformDefinitionParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.TransformOccurrences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransformOccurrences`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.TransformOccurrences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransformOccurrencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAssemblyTransformDefinitionParams** | [**BTAssemblyTransformDefinitionParams**](BTAssemblyTransformDefinitionParams.md) |  | 

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


## TranslateFormat

> BTTranslationRequestInfo TranslateFormat(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Export the assembly to another format.



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
    resp, r, err := apiClient.AssemblyApi.TranslateFormat(context.Background(), did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.TranslateFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslateFormat`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.TranslateFormat`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiTranslateFormatRequest struct via the builder pattern


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


## UpdateFeature

> BTFeatureDefinitionResponse1617 UpdateFeature(ctx, did, wid, eid, fid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()

Update an existing feature for an Assembly.

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
    fid := "fid_example" // string | 
    bTFeatureDefinitionCall1406 := *openapiclient.NewBTFeatureDefinitionCall1406() // BTFeatureDefinitionCall1406 |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.UpdateFeature(context.Background(), did, wid, eid, fid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.UpdateFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeature`: BTFeatureDefinitionResponse1617
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.UpdateFeature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**fid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureRequest struct via the builder pattern


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


## UpdateMateValues

> BTAssemblyMateValuesInfo UpdateMateValues(ctx, did, wid, eid).BTAssemblyMateValuesInfo(bTAssemblyMateValuesInfo).Execute()

Update mate values for the given mates in the specified assembly.



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
    bTAssemblyMateValuesInfo := *openapiclient.NewBTAssemblyMateValuesInfo() // BTAssemblyMateValuesInfo | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.AssemblyApi.UpdateMateValues(context.Background(), did, wid, eid).BTAssemblyMateValuesInfo(bTAssemblyMateValuesInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssemblyApi.UpdateMateValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMateValues`: BTAssemblyMateValuesInfo
    fmt.Fprintf(os.Stdout, "Response from `AssemblyApi.UpdateMateValues`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiUpdateMateValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAssemblyMateValuesInfo** | [**BTAssemblyMateValuesInfo**](BTAssemblyMateValuesInfo.md) |  | 

### Return type

[**BTAssemblyMateValuesInfo**](BTAssemblyMateValuesInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

