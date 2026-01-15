# \SketchApi

All URIs are relative to *https://cad.onshape.com/api/v13*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSketchBoundingBoxes**](SketchApi.md#GetSketchBoundingBoxes) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/boundingboxes | Get all bounding boxes for a sketch.
[**GetSketchInfo**](SketchApi.md#GetSketchInfo) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches | Get information for all sketches in Part Studio.
[**GetTessellatedEntities**](SketchApi.md#GetTessellatedEntities) | **Get** /partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/tessellatedentities | Get the tessellations of a sketch in a Part Studio.



## GetSketchBoundingBoxes

> BTBoundingBoxInfo GetSketchBoundingBoxes(ctx, did, wvm, wvmid, eid, sid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Get all bounding boxes for a sketch.

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
    sid := "sid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.SketchApi.GetSketchBoundingBoxes(context.Background(), did, wvm, wvmid, eid, sid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SketchApi.GetSketchBoundingBoxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSketchBoundingBoxes`: BTBoundingBoxInfo
    fmt.Fprintf(os.Stdout, "Response from `SketchApi.GetSketchBoundingBoxes`: %v\n", resp)
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
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSketchBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

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


## GetSketchInfo

> map[string]interface{} GetSketchInfo(ctx, did, wvm, wvmid, eid).Configuration(configuration).SketchId(sketchId).Output3D(output3D).CurvePoints(curvePoints).IncludeGeometry(includeGeometry).LinkDocumentId(linkDocumentId).Execute()

Get information for all sketches in Part Studio.

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
    configuration := "configuration_example" // string | URL-encoded string of configuration values (separated by `;`). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. (optional)
    sketchId := []string{"Inner_example"} // []string |  (optional)
    output3D := true // bool |  (optional) (default to false)
    curvePoints := true // bool |  (optional) (default to false)
    includeGeometry := true // bool |  (optional) (default to true)
    linkDocumentId := "linkDocumentId_example" // string | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.SketchApi.GetSketchInfo(context.Background(), did, wvm, wvmid, eid).Configuration(configuration).SketchId(sketchId).Output3D(output3D).CurvePoints(curvePoints).IncludeGeometry(includeGeometry).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SketchApi.GetSketchInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSketchInfo`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SketchApi.GetSketchInfo`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetSketchInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** | URL-encoded string of configuration values (separated by &#x60;;&#x60;). See the [Configurations API Guide](https://onshape-public.github.io/docs/api-adv/configs/) for details. | 
 **sketchId** | **[]string** |  | 
 **output3D** | **bool** |  | [default to false]
 **curvePoints** | **bool** |  | [default to false]
 **includeGeometry** | **bool** |  | [default to true]
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

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


## GetTessellatedEntities

> map[string]interface{} GetTessellatedEntities(ctx, did, wvm, wvmid, eid, sid).Configuration(configuration).EntityId(entityId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).LinkDocumentId(linkDocumentId).Execute()

Get the tessellations of a sketch in a Part Studio.



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
    sid := "sid_example" // string | 
    configuration := "configuration_example" // string |  (optional)
    entityId := []string{"Inner_example"} // []string |  (optional)
    angleTolerance := float64(1.2) // float64 |  (optional)
    chordTolerance := float64(1.2) // float64 |  (optional)
    linkDocumentId := "linkDocumentId_example" // string |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.SketchApi.GetTessellatedEntities(context.Background(), did, wvm, wvmid, eid, sid).Configuration(configuration).EntityId(entityId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SketchApi.GetTessellatedEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTessellatedEntities`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SketchApi.GetTessellatedEntities`: %v\n", resp)
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
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTessellatedEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **entityId** | **[]string** |  | 
 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **linkDocumentId** | **string** |  | 

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

