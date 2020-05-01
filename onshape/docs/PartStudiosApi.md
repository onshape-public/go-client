# \PartStudiosApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPartStudioFeature**](PartStudiosApi.md#AddPartStudioFeature) | **Post** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Add Feature
[**ComparePartStudios**](PartStudiosApi.md#ComparePartStudios) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/compare | Compare Part Studios
[**CreatePartStudio**](PartStudiosApi.md#CreatePartStudio) | **Post** /api/partstudios/d/{did}/w/{wid} | Create Part Studio
[**CreatePartStudioTranslation**](PartStudiosApi.md#CreatePartStudioTranslation) | **Post** /api/partstudios/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create Part Studio translation
[**DeletePartStudioFeature**](PartStudiosApi.md#DeletePartStudioFeature) | **Delete** /api/partstudios/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Delete Feature
[**EvalFeatureScript**](PartStudiosApi.md#EvalFeatureScript) | **Post** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurescript | Evaluate FeatureScript
[**ExportPS1**](PartStudiosApi.md#ExportPS1) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/parasolid | Export Part Studio to Parasolid
[**ExportStl1**](PartStudiosApi.md#ExportStl1) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/stl | Export Part Studio to STL
[**GetPartStudioBodyDetails**](PartStudiosApi.md#GetPartStudioBodyDetails) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/bodydetails | Array of body information
[**GetPartStudioBoundingBoxes**](PartStudiosApi.md#GetPartStudioBoundingBoxes) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/boundingboxes | Mass properties of parts or a PartStudio.
[**GetPartStudioConfiguration**](PartStudiosApi.md#GetPartStudioConfiguration) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Get Configuration
[**GetPartStudioFeatureSpecs**](PartStudiosApi.md#GetPartStudioFeatureSpecs) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | Get Feature Specs
[**GetPartStudioFeatures**](PartStudiosApi.md#GetPartStudioFeatures) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Get Feature List
[**GetPartStudioMassProperties**](PartStudiosApi.md#GetPartStudioMassProperties) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/massproperties | Mass properties of parts or a PartStudio.
[**GetPartStudioNamedViews**](PartStudiosApi.md#GetPartStudioNamedViews) | **Get** /api/partstudios/d/{did}/e/{eid}/namedViews | Get Named Views
[**GetPartStudioShadedViews**](PartStudiosApi.md#GetPartStudioShadedViews) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/shadedviews | Get Shaded Views
[**GetPartStudioTessellatedEdges**](PartStudiosApi.md#GetPartStudioTessellatedEdges) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/tessellatededges | Tesselated edges from a PartStudio.
[**GetPartStudioTessellatedFaces**](PartStudiosApi.md#GetPartStudioTessellatedFaces) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/tessellatedfaces | Tesselated faces of the parts in the Part Studio.
[**TranslateIds**](PartStudiosApi.md#TranslateIds) | **Post** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/idtranslations | Id Translations
[**UpdateFeatures**](PartStudiosApi.md#UpdateFeatures) | **Post** /api/partstudios/d/{did}/w/{wid}/e/{eid}/features/updates | Update Features
[**UpdatePartStudioConfiguration**](PartStudiosApi.md#UpdatePartStudioConfiguration) | **Post** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Update Configuration
[**UpdatePartStudioFeature**](PartStudiosApi.md#UpdatePartStudioFeature) | **Post** /api/partstudios/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Update Feature
[**UpdateRollback**](PartStudiosApi.md#UpdateRollback) | **Post** /api/partstudios/d/{did}/w/{wid}/e/{eid}/features/rollback | Update Feature Rollback



## AddPartStudioFeature

> BTFeatureDefinitionResponse1617 AddPartStudioFeature(ctx, did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()

Add Feature

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




 **bTFeatureDefinitionCall1406** | [**BTFeatureDefinitionCall1406**](BTFeatureDefinitionCall1406.md) | feature The serialized feature definition | 

### Return type

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse-1617.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/vnd.onshape.v2+json;charset=UTF-8;qs=0.2
- **Accept**: application/vnd.onshape.v2+json;charset=UTF-8;qs=0.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ComparePartStudios

> BTRootDiffInfo ComparePartStudios(ctx, did, wvm, wvmid, eid).WorkspaceId(workspaceId).VersionId(versionId).MicroversionId(microversionId).SourceConfiguration(sourceConfiguration).TargetConfiguration(targetConfiguration).LinkDocumentId(linkDocumentId).Execute()

Compare Part Studios

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudio

> BTDocumentElementInfo CreatePartStudio(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create Part Studio

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartStudioTranslation

> BTTranslationRequestInfo CreatePartStudioTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create Part Studio translation

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePartStudioFeature

> BTFeatureApiBase1430 DeletePartStudioFeature(ctx, did, wid, eid, fid).Execute()

Delete Feature

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

[**BTFeatureApiBase1430**](BTFeatureApiBase-1430.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvalFeatureScript

> BTFeatureScriptEvalResponse1859 EvalFeatureScript(ctx, did, wvm, wvmid, eid).Configuration(configuration).BTFeatureScriptEvalCall2377(bTFeatureScriptEvalCall2377).Execute()

Evaluate FeatureScript

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvalFeatureScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** | Configuration string. | 
 **bTFeatureScriptEvalCall2377** | [**BTFeatureScriptEvalCall2377**](BTFeatureScriptEvalCall2377.md) |  | 

### Return type

[**BTFeatureScriptEvalResponse1859**](BTFeatureScriptEvalResponse-1859.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPS1

> ExportPS1(ctx, did, wvm, wvmid, eid).PartIds(partIds).Version(version).IncludeExportIds(includeExportIds).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Export Part Studio to Parasolid

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPS1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **partIds** | **string** | IDs of the parts to retrieve. Repeat query param to add more than one (i.e. partId&#x3D;JHK&amp;partId&#x3D;JHD). May not be combined with other ID filters | 
 **version** | **string** | Parasolid version | [default to &quot;0&quot;]
 **includeExportIds** | **bool** | Whether topolgy ids should be exported as parasolid attributes | [default to false]
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportStl1

> ExportStl1(ctx, did, wvm, wvmid, eid).PartIds(partIds).Mode(mode).Grouping(grouping).Scale(scale).Units(units).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).MinFacetWidth(minFacetWidth).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Export Part Studio to STL

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportStl1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **partIds** | **string** | IDs of the parts to retrieve. Repeat query param to add more than one (i.e. partId&#x3D;JHK&amp;partId&#x3D;JHD). May not be combined with other ID filters | 
 **mode** | **string** | Type of file: text, binary | [default to &quot;text&quot;]
 **grouping** | **bool** | Whether parts should be exported as a group or individually in a .zip file | [default to true]
 **scale** | **float64** | Scale for measurements. | [default to 1.0]
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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioBodyDetails

> BTExportModelBodiesResponse734 GetPartStudioBodyDetails(ctx, did, wvm, wvmid, eid).Configuration(configuration).LinkDocumentId(linkDocumentId).RollbackBarIndex(rollbackBarIndex).Body(body).Execute()

Array of body information

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioBodyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **rollbackBarIndex** | **int32** |  | [default to -1]
 **body** | **string** |  | 

### Return type

[**BTExportModelBodiesResponse734**](BTExportModelBodiesResponse-734.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioBoundingBoxes

> BTBoundingBoxInfo GetPartStudioBoundingBoxes(ctx, did, wvm, wvmid, eid).IncludeHidden(includeHidden).IncludeWireBodies(includeWireBodies).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Mass properties of parts or a PartStudio.

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioConfiguration

> GetPartStudioConfiguration(ctx, did, wvm, wvmid, eid).Execute()

Get Configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFeatureSpecs

> BTFeatureSpecsResponse664 GetPartStudioFeatureSpecs(ctx, did, wvm, wvmid, eid).Execute()

Get Feature Specs

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

[**BTFeatureSpecsResponse664**](BTFeatureSpecsResponse-664.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioFeatures

> BTFeatureListResponse2457 GetPartStudioFeatures(ctx, did, wvm, wvmid, eid).FeatureId(featureId).LinkDocumentId(linkDocumentId).NoSketchGeometry(noSketchGeometry).Execute()

Get Feature List

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




 **featureId** | [**[]string**](string.md) | ID of a feature; repeat query param to add more than one | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 
 **noSketchGeometry** | **bool** | Whether or not to output simple sketch info without geometry | [default to false]

### Return type

[**BTFeatureListResponse2457**](BTFeatureListResponse-2457.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v2+json;charset=UTF-8;qs=0.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioMassProperties

> BTMassPropertiesBulkInfo GetPartStudioMassProperties(ctx, did, wvm, wvmid, eid).PartId(partId).MassAsGroup(massAsGroup).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Mass properties of parts or a PartStudio.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**wvm** | **string** | One of w or v or m corresponding to whether a workspace or version or microversion was entered. | 
**wvmid** | **string** | Workspace (w), Version (v) or Microversion (m) ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioMassPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **partId** | [**[]string**](string.md) | IDs of the parts to retrieve. Repeat query param to add more than one (i.e. partId&#x3D;JHK&amp;partId&#x3D;JHD). May not be combined with other ID filters | 
 **massAsGroup** | **bool** | If true, specified parts will be evaluated as a single object instead of individually | [default to true]
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**BTMassPropertiesBulkInfo**](BTMassPropertiesBulkInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioNamedViews

> BTNamedViewsInfo GetPartStudioNamedViews(ctx, did, eid).SkipPerspective(skipPerspective).IncludeSectionCutViews(includeSectionCutViews).LinkDocumentId(linkDocumentId).Execute()

Get Named Views

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 
**eid** | **string** | Element ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioNamedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipPerspective** | **bool** | Whether views with a perspective projection should be omitted. | [default to true]
 **includeSectionCutViews** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTNamedViewsInfo**](BTNamedViewsInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioShadedViews

> BTShadedViewsInfo GetPartStudioShadedViews(ctx, did, wvm, wvmid, eid).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Get Shaded Views

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
 **configuration** | **string** | Configuration string. | 
 **linkDocumentId** | **string** | Id of document that links to the document being accessed. This may provide additional access rights to the document. Allowed only with version (v) path parameter. | 

### Return type

[**BTShadedViewsInfo**](BTShadedViewsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioTessellatedEdges

> BTExportTessellatedEdgesResponse327 GetPartStudioTessellatedEdges(ctx, did, wvm, wvmid, eid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).PartId(partId).EdgeId(edgeId).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()

Tesselated edges from a PartStudio.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioTessellatedEdgesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **partId** | [**[]string**](string.md) |  | 
 **edgeId** | [**[]string**](string.md) |  | 
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTExportTessellatedEdgesResponse327**](BTExportTessellatedEdgesResponse-327.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartStudioTessellatedFaces

> BTExportTessellatedFacesResponse898 GetPartStudioTessellatedFaces(ctx, did, wvm, wvmid, eid).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).MaxFacetWidth(maxFacetWidth).OutputVertexNormals(outputVertexNormals).OutputFacetNormals(outputFacetNormals).OutputTextureCoordinates(outputTextureCoordinates).OutputFaceAppearances(outputFaceAppearances).OutputIndexTable(outputIndexTable).PartId(partId).FaceId(faceId).OutputErrorFaces(outputErrorFaces).Configuration(configuration).LinkDocumentId(linkDocumentId).Body(body).Execute()

Tesselated faces of the parts in the Part Studio.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartStudioTessellatedFacesRequest struct via the builder pattern


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
 **partId** | [**[]string**](string.md) |  | 
 **faceId** | [**[]string**](string.md) |  | 
 **outputErrorFaces** | **bool** |  | [default to false]
 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **body** | **string** |  | 

### Return type

[**BTExportTessellatedFacesResponse898**](BTExportTessellatedFacesResponse-898.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslateIds

> BTIdTranslationInfo TranslateIds(ctx, did, wvm, wvmid, eid).BTIdTranslationParams(bTIdTranslationParams).Execute()

Id Translations

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatures

> BTUpdateFeaturesResponse1333 UpdateFeatures(ctx, did, wid, eid).Body(body).Execute()

Update Features

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



 **body** | **string** | feature The serialized feature definition | 

### Return type

[**BTUpdateFeaturesResponse1333**](BTUpdateFeaturesResponse-1333.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartStudioConfiguration

> BTConfigurationResponse2019 UpdatePartStudioConfiguration(ctx, did, wvm, wvmid, eid).Body(body).Execute()

Update Configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartStudioConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTConfigurationResponse2019**](BTConfigurationResponse-2019.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartStudioFeature

> BTFeatureDefinitionResponse1617 UpdatePartStudioFeature(ctx, did, wid, eid, fid).Body(body).Execute()

Update Feature

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




 **body** | **string** | feature The serialized feature definition | 

### Return type

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse-1617.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRollback

> BTSetFeatureRollbackResponse1042 UpdateRollback(ctx, did, wid, eid).Body(body).Execute()

Update Feature Rollback

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

[**BTSetFeatureRollbackResponse1042**](BTSetFeatureRollbackResponse-1042.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

