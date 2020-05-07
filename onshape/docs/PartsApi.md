# \PartsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportPS**](PartsApi.md#ExportPS) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/parasolid | Export Part to Parasolid.
[**GetBendTable**](PartsApi.md#GetBendTable) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/sheetmetal/bendtable | Get Sheet Metal Bend Table.
[**GetBodyDetails**](PartsApi.md#GetBodyDetails) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/bodydetails | 
[**GetBoundingBoxes**](PartsApi.md#GetBoundingBoxes) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/boundingboxes | 
[**GetEdges**](PartsApi.md#GetEdges) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/tessellatededges | Tessellated Edges
[**GetFaces1**](PartsApi.md#GetFaces1) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/tessellatedfaces | Get Tessellated Faces
[**GetMassProperties**](PartsApi.md#GetMassProperties) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/massproperties | 
[**GetPartMetadata**](PartsApi.md#GetPartMetadata) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/metadata | 
[**GetPartsWMV**](PartsApi.md#GetPartsWMV) | **Get** /api/parts/d/{did}/{wvm}/{wvmid} | Get list of parts
[**GetPartsWMVE**](PartsApi.md#GetPartsWMVE) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid} | Get parts from an element.
[**GetShadedViews1**](PartsApi.md#GetShadedViews1) | **Get** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/shadedviews | 
[**GetStandardContentPartMetadata**](PartsApi.md#GetStandardContentPartMetadata) | **Get** /api/parts/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/partid/{partid}/metadata | 
[**UpdatePartMetadata**](PartsApi.md#UpdatePartMetadata) | **Post** /api/parts/d/{did}/{wvm}/{wvmid}/e/{eid}/partid/{partid}/metadata | 
[**UpdateStandardContentPartMetadata**](PartsApi.md#UpdateStandardContentPartMetadata) | **Post** /api/parts/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/partid/{partid}/metadata | 



## ExportPS

> *os.File ExportPS(ctx, did, wvm, wvmid, eid, partid).Version(version).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Export Part to Parasolid.

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

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+octet-stream;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBendTable

> BTTableResponse1546 GetBendTable(ctx, did, wvm, wvmid, eid, partid).LinkDocumentId(linkDocumentId).Execute()

Get Sheet Metal Bend Table.

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

[**BTTableResponse1546**](BTTableResponse-1546.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBodyDetails

> BTExportModelBodiesResponse734 GetBodyDetails(ctx, did, wvm, wvmid, eid, partid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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

Other parameters are passed through a pointer to a apiGetBodyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTExportModelBodiesResponse734**](BTExportModelBodiesResponse-734.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoundingBoxes

> BTBoundingBoxInfo GetBoundingBoxes(ctx, did, wvm, wvmid, eid, partid).IncludeHidden(includeHidden).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEdges

> BTExportTessellatedEdgesResponse327 GetEdges(ctx, did, wvm, wvmid, eid, partid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).EdgeId(edgeId).Configuration(configuration).LinkDocumentId(linkDocumentId).Body(body).Execute()

Tessellated Edges

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

Other parameters are passed through a pointer to a apiGetEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **edgeId** | [**[]string**](string.md) |  | 
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **body** | **string** |  | 

### Return type

[**BTExportTessellatedEdgesResponse327**](BTExportTessellatedEdgesResponse-327.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFaces1

> BTExportTessellatedFacesResponse898 GetFaces1(ctx, did, wvm, wvmid, eid, partid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputFaceAppearances(outputFaceAppearances).OutputIndexTable(outputIndexTable).FaceId(faceId).Configuration(configuration).OutputErrorFaces(outputErrorFaces).LinkDocumentId(linkDocumentId).Body(body).Execute()

Get Tessellated Faces

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

Other parameters are passed through a pointer to a apiGetFaces1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **maxFacetWidth** | **float64** |  | 
 **outputVertexNormals** | **bool** |  | [default to false]
 **outputFacetNormals** | **bool** |  | [default to true]
 **outputTextureCoordinates** | **bool** |  | [default to false]
 **outputFaceAppearances** | **bool** |  | [default to false]
 **outputIndexTable** | **bool** |  | [default to false]
 **faceId** | [**[]string**](string.md) |  | 
 **configuration** | **string** |  | 
 **outputErrorFaces** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 
 **body** | **string** |  | 

### Return type

[**BTExportTessellatedFacesResponse898**](BTExportTessellatedFacesResponse-898.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMassProperties

> BTMassPropertiesBulkInfo GetMassProperties(ctx, did, wvm, wvmid, eid, partid).InferMetadataOwner(inferMetadataOwner).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()



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

Other parameters are passed through a pointer to a apiGetMassPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **inferMetadataOwner** | **bool** |  | [default to true]
 **linkDocumentId** | **string** |  | 
 **configuration** | **string** |  | 

### Return type

[**BTMassPropertiesBulkInfo**](BTMassPropertiesBulkInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartMetadata

> BTPartMetadataInfo GetPartMetadata(ctx, did, wvm, wvmid, eid, partid).InferMetadataOwner(inferMetadataOwner).IncludePropertyDefaults(includePropertyDefaults).FriendlyUserIds(friendlyUserIds).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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

Other parameters are passed through a pointer to a apiGetPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **inferMetadataOwner** | **bool** |  | [default to false]
 **includePropertyDefaults** | **bool** |  | [default to false]
 **friendlyUserIds** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartsWMV

> []BTPartMetadataInfo GetPartsWMV(ctx, did, wvm, wvmid).ElementId(elementId).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()

Get list of parts

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartsWMVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **elementId** | **string** | Element ID | 
 **withThumbnails** | **bool** | Whether or not to include thumbnails (not supported for microversion) | [default to false]
 **includePropertyDefaults** | **bool** | If true, include metadata schema property defaults in response | [default to false]
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 
 **configuration** | **string** | Configuration string. | 

### Return type

[**[]BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartsWMVE

> []BTPartMetadataInfo GetPartsWMVE(ctx, did, wvm, wvmid, eid).WithThumbnails(withThumbnails).IncludePropertyDefaults(includePropertyDefaults).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Get parts from an element.

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
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**[]BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShadedViews1

> BTShadedViewsInfo GetShadedViews1(ctx, did, wvm, wvmid, eid, partid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).UseAntiAliasing(useAntiAliasing).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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

Other parameters are passed through a pointer to a apiGetShadedViews1Request struct via the builder pattern


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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStandardContentPartMetadata

> BTPartMetadataInfo GetStandardContentPartMetadata(ctx, did, vid, eid, otype, oid, partid).IncludePropertyDefaults(includePropertyDefaults).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 
**eid** | **string** |  | 
**otype** | **string** |  | 
**oid** | **string** |  | 
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStandardContentPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **includePropertyDefaults** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartMetadata

> BTPartMetadataInfo UpdatePartMetadata(ctx, did, wvm, wvmid, eid, partid).BTWorkspacePartParams(bTWorkspacePartParams).Execute()



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

Other parameters are passed through a pointer to a apiUpdatePartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bTWorkspacePartParams** | [**BTWorkspacePartParams**](BTWorkspacePartParams.md) |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStandardContentPartMetadata

> BTPartMetadataInfo UpdateStandardContentPartMetadata(ctx, did, vid, eid, otype, oid, partid).LinkDocumentId(linkDocumentId).IncludePropertyDefaults(includePropertyDefaults).BTWorkspacePartParams(bTWorkspacePartParams).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 
**eid** | **string** |  | 
**otype** | **string** |  | 
**oid** | **string** |  | 
**partid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStandardContentPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **linkDocumentId** | **string** |  | 
 **includePropertyDefaults** | **bool** |  | [default to false]
 **bTWorkspacePartParams** | [**BTWorkspacePartParams**](BTWorkspacePartParams.md) |  | 

### Return type

[**BTPartMetadataInfo**](BTPartMetadataInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

