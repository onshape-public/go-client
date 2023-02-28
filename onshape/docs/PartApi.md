# \PartApi

All URIs are relative to *https://cad.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportPS**](PartApi.md#ExportPS) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/parasolid | Export part to Parasolid by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**ExportPartGltf**](PartApi.md#ExportPartGltf) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/gltf | Retrieve GLTF for part by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**ExportStl**](PartApi.md#ExportStl) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/stl | Retrieve part STL by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**GetBendTable**](PartApi.md#GetBendTable) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/sheetmetal/bendtable | Retrieve sheet metal bend table by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**GetBodyDetails**](PartApi.md#GetBodyDetails) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/bodydetails | Retrieve part body details by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**GetBoundingBoxes**](PartApi.md#GetBoundingBoxes) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/boundingboxes | Retrieve part bounding boxes by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**GetEdges**](PartApi.md#GetEdges) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/tessellatededges | Retrieve tessellated edges of a part by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**GetFaces1**](PartApi.md#GetFaces1) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/tessellatedfaces | 
[**GetMassProperties**](PartApi.md#GetMassProperties) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/massproperties | Retrieve mass properties of a part document ID, workspace or version or microversion ID, tab ID, and part ID.
[**GetPartShadedViews**](PartApi.md#GetPartShadedViews) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/shadedviews | Retrieve shaded views of a part by document ID, workspace or version or microversion ID, tab ID, and part ID.
[**GetPartsWMV**](PartApi.md#GetPartsWMV) | **Get** /parts/d/{did}/{wvm}/{wvmid} | Retrieve a list of parts by document ID, and workspace or version or microversion ID.
[**GetPartsWMVE**](PartApi.md#GetPartsWMVE) | **Get** /parts/d/{did}/{wvm}/{wvmid}/e/{eid} | Retrieve a list of parts from a tab by document ID, workspace or version or microversion ID, and tab ID.



## ExportPS

> HttpFile ExportPS(ctx, did, wvm, wvmid, eid, partid).Version(version).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Export part to Parasolid by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    version := "version_example" // string |  (optional) (default to "0")
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.ExportPS(context.Background(), did, wvm, wvmid, eid, partid).Version(version).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.ExportPS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPS`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `PartApi.ExportPS`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **version** | **string** |  | [default to &quot;0&quot;]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**HttpFile**](HttpFile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPartGltf

> HttpFile ExportPartGltf(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).OutputSeparateFaceNodes(outputSeparateFaceNodes).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).Execute()

Retrieve GLTF for part by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    precomputedLevelOfDetail := "precomputedLevelOfDetail_example" // string |  (optional)
    outputSeparateFaceNodes := true // bool |  (optional) (default to false)
    faceId := []string{"Inner_example"} // []string |  (optional)
    outputFaceAppearances := true // bool |  (optional) (default to false)
    maxFacetWidth := float64(1.2) // float64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.ExportPartGltf(context.Background(), did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).OutputSeparateFaceNodes(outputSeparateFaceNodes).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.ExportPartGltf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPartGltf`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `PartApi.ExportPartGltf`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPartGltfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
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


## ExportStl

> map[string]interface{} ExportStl(ctx, did, wvm, wvmid, eid, partid).Mode(mode).Grouping(grouping).Scale(scale).Units(units).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).MinFacetWidth(minFacetWidth).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Retrieve part STL by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    mode := "mode_example" // string |  (optional) (default to "text")
    grouping := true // bool |  (optional) (default to true)
    scale := float64(1.2) // float64 |  (optional) (default to 1)
    units := "units_example" // string |  (optional) (default to "inch")
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    maxFacetWidth := float64(1.2) // float64 |  (optional)
    minFacetWidth := float64(1.2) // float64 |  (optional)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.ExportStl(context.Background(), did, wvm, wvmid, eid, partid).Mode(mode).Grouping(grouping).Scale(scale).Units(units).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).MinFacetWidth(minFacetWidth).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.ExportStl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportStl`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PartApi.ExportStl`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportStlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **mode** | **string** |  | [default to &quot;text&quot;]
 **grouping** | **bool** |  | [default to true]
 **scale** | **float64** |  | [default to 1]
 **units** | **string** |  | [default to &quot;inch&quot;]
 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **maxFacetWidth** | **float64** |  | 
 **minFacetWidth** | **float64** |  | 
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBendTable

> BTTableResponse1546 GetBendTable(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Execute()

Retrieve sheet metal bend table by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetBendTable(context.Background(), did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetBendTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBendTable`: BTTableResponse1546
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetBendTable`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBendTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** |  | 

### Return type

[**BTTableResponse1546**](BTTableResponse1546.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBodyDetails

> BTExportModelBodiesResponse734 GetBodyDetails(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).IncludeGeometricData(includeGeometricData).Execute()

Retrieve part body details by document ID, workspace or version or microversion ID, tab ID, and part ID.



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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    includeGeometricData := true // bool | Whether or not geometric data should be included in the response. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetBodyDetails(context.Background(), did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).IncludeGeometricData(includeGeometricData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetBodyDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBodyDetails`: BTExportModelBodiesResponse734
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetBodyDetails`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBodyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
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


## GetBoundingBoxes

> BTBoundingBoxInfo GetBoundingBoxes(ctx, did, wvm, wvmid, eid, partid).IncludeHidden(includeHidden).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Retrieve part bounding boxes by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    includeHidden := true // bool |  (optional) (default to false)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetBoundingBoxes(context.Background(), did, wvm, wvmid, eid, partid).IncludeHidden(includeHidden).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetBoundingBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoundingBoxes`: BTBoundingBoxInfo
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetBoundingBoxes`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **includeHidden** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

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


## GetEdges

> BTExportTessellatedEdgesResponse327 GetEdges(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).EdgeId(edgeId).Execute()

Retrieve tessellated edges of a part by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    precomputedLevelOfDetail := "precomputedLevelOfDetail_example" // string |  (optional)
    edgeId := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetEdges(context.Background(), did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).EdgeId(edgeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetEdges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEdges`: BTExportTessellatedEdgesResponse327
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetEdges`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
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


## GetFaces1

> BTExportTessellatedFacesResponse898 GetFaces1(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputIndexTable(outputIndexTable).OutputErrorFaces(outputErrorFaces).CombineCompositePartConstituents(combineCompositePartConstituents).Execute()





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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
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
    resp, r, err := apiClient.PartApi.GetFaces1(context.Background(), did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PrecomputedLevelOfDetail(precomputedLevelOfDetail).FaceId(faceId).OutputFaceAppearances(outputFaceAppearances).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputIndexTable(outputIndexTable).OutputErrorFaces(outputErrorFaces).CombineCompositePartConstituents(combineCompositePartConstituents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetFaces1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFaces1`: BTExportTessellatedFacesResponse898
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetFaces1`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFaces1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
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


## GetMassProperties

> BTMassPropertiesBulkInfo GetMassProperties(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).InferMetadataOwner(inferMetadataOwner).UseMassPropertyOverrides(useMassPropertyOverrides).Execute()

Retrieve mass properties of a part document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    rollbackBarIndex := int32(56) // int32 | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. (optional) (default to -1)
    elementMicroversionId := "elementMicroversionId_example" // string | A specific element microversion in which to evaluate the request. (optional)
    inferMetadataOwner := true // bool |  (optional) (default to true)
    useMassPropertyOverrides := true // bool | If true, use the user mass property overrides when calculated mass properties (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetMassProperties(context.Background(), did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Configuration(configuration).RollbackBarIndex(rollbackBarIndex).ElementMicroversionId(elementMicroversionId).InferMetadataOwner(inferMetadataOwner).UseMassPropertyOverrides(useMassPropertyOverrides).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetMassProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMassProperties`: BTMassPropertiesBulkInfo
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetMassProperties`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMassPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **rollbackBarIndex** | **int32** | Index specifying the location of the rollback bar when the call is evaluated. A -1 indicates that it should be at the end of the featurelist. | [default to -1]
 **elementMicroversionId** | **string** | A specific element microversion in which to evaluate the request. | 
 **inferMetadataOwner** | **bool** |  | [default to true]
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


## GetPartShadedViews

> BTShadedViewsInfo GetPartShadedViews(ctx, did, wvm, wvmid, eid, partid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).UseAntiAliasing(useAntiAliasing).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Retrieve shaded views of a part by document ID, workspace or version or microversion ID, tab ID, and part ID.

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
    partid := "partid_example" // string | 
    viewMatrix := "viewMatrix_example" // string |  (optional) (default to "front")
    outputHeight := int32(56) // int32 |  (optional) (default to 500)
    outputWidth := int32(56) // int32 |  (optional) (default to 500)
    pixelSize := float64(1.2) // float64 |  (optional) (default to 0.003)
    edges := "edges_example" // string |  (optional) (default to "show")
    useAntiAliasing := true // bool |  (optional) (default to false)
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetPartShadedViews(context.Background(), did, wvm, wvmid, eid, partid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).UseAntiAliasing(useAntiAliasing).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetPartShadedViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartShadedViews`: BTShadedViewsInfo
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetPartShadedViews`: %v\n", resp)
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
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartShadedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **viewMatrix** | **string** |  | [default to &quot;front&quot;]
 **outputHeight** | **int32** |  | [default to 500]
 **outputWidth** | **int32** |  | [default to 500]
 **pixelSize** | **float64** |  | [default to 0.003]
 **edges** | **string** |  | [default to &quot;show&quot;]
 **useAntiAliasing** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

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


## GetPartsWMV

> []BTPartMetadataInfo GetPartsWMV(ctx, did, wvm, wvmid).ElementId(elementId).LinkDocumentId(linkDocumentId).Configuration(configuration).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).IncludeFlatParts(includeFlatParts).Execute()

Retrieve a list of parts by document ID, and workspace or version or microversion ID.

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
    elementId := "elementId_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    withThumbnails := true // bool | Whether or not to include thumbnails (not supported for microversion) (optional) (default to false)
    includePropertyDefaults := true // bool | If true, include metadata schema property defaults in response (optional) (default to false)
    includeFlatParts := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetPartsWMV(context.Background(), did, wvm, wvmid).ElementId(elementId).LinkDocumentId(linkDocumentId).Configuration(configuration).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).IncludeFlatParts(includeFlatParts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetPartsWMV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartsWMV`: []BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetPartsWMV`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace id, version id, or document microversion id is specified below. | 
**wvmid** | **string** | The id of the workspace, version, or document microversion in which the operation should be performed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartsWMVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **elementId** | **string** | The id of the element in which to perform the operation. | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **withThumbnails** | **bool** | Whether or not to include thumbnails (not supported for microversion) | [default to false]
 **includePropertyDefaults** | **bool** | If true, include metadata schema property defaults in response | [default to false]
 **includeFlatParts** | **bool** |  | 

### Return type

[**[]BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartsWMVE

> []BTPartMetadataInfo GetPartsWMVE(ctx, did, wvm, wvmid, eid).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).IncludeFlatParts(includeFlatParts).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Retrieve a list of parts from a tab by document ID, workspace or version or microversion ID, and tab ID.

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
    withThumbnails := true // bool | Whether or not to include thumbnails (not supported for microversion) (optional) (default to false)
    includePropertyDefaults := true // bool | If true, include metadata schema property defaults in response (optional) (default to false)
    includeFlatParts := true // bool |  (optional)
    configuration := "configuration_example" // string | Configuration string. (optional)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PartApi.GetPartsWMVE(context.Background(), did, wvm, wvmid, eid).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).IncludeFlatParts(includeFlatParts).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartApi.GetPartsWMVE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartsWMVE`: []BTPartMetadataInfo
    fmt.Fprintf(os.Stdout, "Response from `PartApi.GetPartsWMVE`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetPartsWMVERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **withThumbnails** | **bool** | Whether or not to include thumbnails (not supported for microversion) | [default to false]
 **includePropertyDefaults** | **bool** | If true, include metadata schema property defaults in response | [default to false]
 **includeFlatParts** | **bool** |  | 
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**[]BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

