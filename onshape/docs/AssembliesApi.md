# \AssembliesApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFeature**](AssembliesApi.md#AddFeature) | **Post** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/features | 
[**CreateAssembly**](AssembliesApi.md#CreateAssembly) | **Post** /api/assemblies/d/{did}/w/{wid} | Create Assembly
[**CreateInstance**](AssembliesApi.md#CreateInstance) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/instances | Create assembly instance
[**DeleteFeature**](AssembliesApi.md#DeleteFeature) | **Delete** /api/assemblies/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | Delete Feature
[**DeleteInstance**](AssembliesApi.md#DeleteInstance) | **Delete** /api/assemblies/d/{did}/w/{wid}/e/{eid}/instance/nodeid/{nid} | Delete assembly instance.
[**GetAssemblyBoundingBoxes**](AssembliesApi.md#GetAssemblyBoundingBoxes) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/boundingboxes | Bounding Boxes.
[**GetAssemblyDefinition**](AssembliesApi.md#GetAssemblyDefinition) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid} | Assembly Definition.
[**GetAssemblyShadedViews**](AssembliesApi.md#GetAssemblyShadedViews) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/shadedviews | 
[**GetBillOfMaterials**](AssembliesApi.md#GetBillOfMaterials) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/bom | Get Bill of Materials
[**GetFeatureSpecs**](AssembliesApi.md#GetFeatureSpecs) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | 
[**GetFeatures**](AssembliesApi.md#GetFeatures) | **Get** /api/assemblies/d/{did}/{wvm}/{wvmid}/e/{eid}/features | Get Feature List
[**GetNamedViews**](AssembliesApi.md#GetNamedViews) | **Get** /api/assemblies/d/{did}/e/{eid}/namedViews | 
[**GetOrCreateBillOfMaterialsElement**](AssembliesApi.md#GetOrCreateBillOfMaterialsElement) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/bomelement | Get or Create Bill of Materials Element
[**GetTranslatorFormats**](AssembliesApi.md#GetTranslatorFormats) | **Get** /api/assemblies/d/{did}/w/{wid}/e/{eid}/translationformats | Get Translation Formats
[**InsertTransformedInstances**](AssembliesApi.md#InsertTransformedInstances) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/transformedinstances | Create and transform assembly instances
[**TransformOccurrences**](AssembliesApi.md#TransformOccurrences) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/occurrencetransforms | Transform assembly occurrences.
[**TranslateFormat**](AssembliesApi.md#TranslateFormat) | **Post** /api/assemblies/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create Assembly translation.
[**UpdateFeature**](AssembliesApi.md#UpdateFeature) | **Post** /api/assemblies/d/{did}/w/{wid}/e/{eid}/features/featureid/{fid} | 



## AddFeature

> BTFeatureDefinitionResponse1617 AddFeature(ctx, did, wvm, wvmid, eid).BTFeatureDefinitionCall1406(bTFeatureDefinitionCall1406).Execute()



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

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse-1617.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAssembly

> BTDocumentElementInfo CreateAssembly(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create Assembly

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstance

> []BTOccurrence74 CreateInstance(ctx, did, wid, eid).BTAssemblyInstanceDefinitionParams(bTAssemblyInstanceDefinitionParams).Execute()

Create assembly instance

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

[**[]BTOccurrence74**](BTOccurrence-74.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFeature

> BTFeatureApiBase1430 DeleteFeature(ctx, did, wid, eid, fid).Execute()

Delete Feature

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

[**BTFeatureApiBase1430**](BTFeatureApiBase-1430.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, did, eid, wid, nid).Execute()

Delete assembly instance.

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

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssemblyBoundingBoxes

> BTBoundingBoxInfo GetAssemblyBoundingBoxes(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).IncludeHidden(includeHidden).DisplayStateId(displayStateId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()

Bounding Boxes.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssemblyBoundingBoxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 
 **includeHidden** | **bool** |  | 
 **displayStateId** | **string** |  | 
 **configuration** | **string** |  | 
 **explodedViewId** | **string** |  | 

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


## GetAssemblyDefinition

> BTAssemblyDefinitionInfo GetAssemblyDefinition(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).IncludeMateFeatures(includeMateFeatures).IncludeNonSolids(includeNonSolids).IncludeMateConnectors(includeMateConnectors).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()

Assembly Definition.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssemblyDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 
 **includeMateFeatures** | **bool** |  | 
 **includeNonSolids** | **bool** |  | 
 **includeMateConnectors** | **bool** |  | 
 **configuration** | **string** |  | 
 **explodedViewId** | **string** |  | 

### Return type

[**BTAssemblyDefinitionInfo**](BTAssemblyDefinitionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssemblyShadedViews

> BTShadedViewsInfo GetAssemblyShadedViews(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).ViewMatrix(viewMatrix).OutputHeight(outputHeight).OutputWidth(outputWidth).PixelSize(pixelSize).Edges(edges).ShowAllParts(showAllParts).IncludeSurfaces(includeSurfaces).UseAntiAliasing(useAntiAliasing).DisplayStateId(displayStateId).Configuration(configuration).ExplodedViewId(explodedViewId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssemblyShadedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 
 **viewMatrix** | **string** |  | [default to &quot;front&quot;]
 **outputHeight** | **int32** |  | [default to 500]
 **outputWidth** | **int32** |  | [default to 500]
 **pixelSize** | **float64** |  | [default to 0.003]
 **edges** | **string** |  | [default to &quot;show&quot;]
 **showAllParts** | **bool** |  | [default to false]
 **includeSurfaces** | **bool** |  | [default to true]
 **useAntiAliasing** | **bool** |  | [default to false]
 **displayStateId** | **string** |  | 
 **configuration** | **string** |  | 
 **explodedViewId** | **string** |  | 

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


## GetBillOfMaterials

> JsonNode GetBillOfMaterials(ctx, did, wvm, wvmid, eid).MetadataWorkspaceId(metadataWorkspaceId).BomColumnIds(bomColumnIds).Indented(indented).MultiLevel(multiLevel).GenerateIfAbsent(generateIfAbsent).LinkDocumentId(linkDocumentId).Configuration(configuration).Execute()

Get Bill of Materials

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillOfMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **metadataWorkspaceId** | **string** |  | [default to &quot;&quot;]
 **bomColumnIds** | [**[]string**](string.md) |  | 
 **indented** | **bool** |  | [default to true]
 **multiLevel** | **bool** |  | [default to false]
 **generateIfAbsent** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 
 **configuration** | **string** |  | 

### Return type

[**JsonNode**](JsonNode.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureSpecs

> BTFeatureSpecsResponse664 GetFeatureSpecs(ctx, did, wvm, wvmid, eid).Execute()



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

[**BTFeatureSpecsResponse664**](BTFeatureSpecsResponse-664.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatures

> BTAssemblyFeatureListResponse1174 GetFeatures(ctx, did, wvm, wvmid, eid).FeatureId(featureId).LinkDocumentId(linkDocumentId).Execute()

Get Feature List

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **featureId** | [**[]string**](string.md) |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTAssemblyFeatureListResponse1174**](BTAssemblyFeatureListResponse-1174.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v2+json;charset=UTF-8;qs=0.2

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedViews

> BTNamedViewsInfo GetNamedViews(ctx, did, eid).SkipPerspective(skipPerspective).LinkDocumentId(linkDocumentId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamedViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipPerspective** | **bool** |  | [default to true]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTNamedViewsInfo**](BTNamedViewsInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrCreateBillOfMaterialsElement

> BTDocumentElementInfo GetOrCreateBillOfMaterialsElement(ctx, did, wid, eid).Execute()

Get or Create Bill of Materials Element

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslatorFormats

> []BTModelFormatInfo GetTranslatorFormats(ctx, did, wid, eid).CheckContent(checkContent).Execute()

Get Translation Formats

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslatorFormatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **checkContent** | **bool** |  | [default to true]

### Return type

[**[]BTModelFormatInfo**](BTModelFormatInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertTransformedInstances

> BTAssemblyInsertTransformedInstancesResponse InsertTransformedInstances(ctx, did, eid, wid).BTAssemblyTransformedInstancesDefinitionParams(bTAssemblyTransformedInstancesDefinitionParams).Execute()

Create and transform assembly instances

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransformOccurrences

> TransformOccurrences(ctx, did, eid, wid).BTAssemblyTransformDefinitionParams(bTAssemblyTransformDefinitionParams).Execute()

Transform assembly occurrences.

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

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TranslateFormat

> BTTranslationRequestInfo TranslateFormat(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create Assembly translation.

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeature

> BTFeatureDefinitionResponse1617 UpdateFeature(ctx, did, wid, eid, fid).Body(body).Execute()



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




 **body** | **string** |  | 

### Return type

[**BTFeatureDefinitionResponse1617**](BTFeatureDefinitionResponse-1617.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

