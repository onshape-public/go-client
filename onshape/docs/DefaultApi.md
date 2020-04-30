# \DefaultApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowableTestObject**](DefaultApi.md#CreateWorkflowableTestObject) | **Post** /api/workflowabletestobject/testobject/{wfid} | 
[**DeleteAssociativeData**](DefaultApi.md#DeleteAssociativeData) | **Delete** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | 
[**GetAssociativeData**](DefaultApi.md#GetAssociativeData) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | 
[**GetLatestInDocument**](DefaultApi.md#GetLatestInDocument) | **Get** /api/insertables/d/{did}/latest | insertables for a document
[**GetMetadataSchema**](DefaultApi.md#GetMetadataSchema) | **Get** /api/metadataschema | 
[**GetProperties**](DefaultApi.md#GetProperties) | **Get** /api/metadataschema/properties | 
[**GetPropertyInfo**](DefaultApi.md#GetPropertyInfo) | **Get** /api/metadataschema/propertyinfo/{pid} | 
[**GetSchema**](DefaultApi.md#GetSchema) | **Get** /api/metadataschema/{sid} | 
[**GetSketchBoundingBoxes**](DefaultApi.md#GetSketchBoundingBoxes) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/boundingboxes | 
[**GetSketchInfo**](DefaultApi.md#GetSketchInfo) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches | 
[**GetTessellatedEntities**](DefaultApi.md#GetTessellatedEntities) | **Get** /api/partstudios/d/{did}/{wvm}/{wvmid}/e/{eid}/sketches/{sid}/tessellatedentities | 
[**GetWorkflowableTestObject**](DefaultApi.md#GetWorkflowableTestObject) | **Get** /api/workflowabletestobject/{oid} | 
[**PostAssociativeData**](DefaultApi.md#PostAssociativeData) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | 
[**TransitionWorkflowableTestObject**](DefaultApi.md#TransitionWorkflowableTestObject) | **Post** /api/workflowabletestobject/{oid}/{transition} | 
[**UpdateWorkflowableTestObject**](DefaultApi.md#UpdateWorkflowableTestObject) | **Post** /api/workflowabletestobject/{oid} | 



## CreateWorkflowableTestObject

> BTWorkflowableTestObjectInfo CreateWorkflowableTestObject(ctx, wfid).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssociativeData

> BTAppElementBasicInfo DeleteAssociativeData(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).AssociativeDataId(associativeDataId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | [default to &quot;&quot;]
 **parentChangeId** | **string** |  | [default to &quot;&quot;]
 **associativeDataId** | [**[]string**](string.md) |  | 
 **elementId** | **string** |  | [default to &quot;&quot;]
 **viewId** | **string** |  | [default to &quot;&quot;]
 **microversionId** | **string** |  | [default to &quot;&quot;]
 **documentMicroversion** | **string** |  | [default to &quot;&quot;]
 **deterministicId** | **string** |  | [default to &quot;&quot;]
 **featureId** | **string** |  | [default to &quot;&quot;]
 **entityId** | **string** |  | [default to &quot;&quot;]
 **occurrenceId** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppElementBasicInfo**](BTAppElementBasicInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssociativeData

> BTAppAssociativeDataInfoArray GetAssociativeData(ctx, did, wvm, wvmid, eid).TransactionId(transactionId).ChangeId(changeId).AssociativeDataId(associativeDataId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).ReturnIdTags(returnIdTags).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | [default to &quot;&quot;]
 **changeId** | **string** |  | [default to &quot;&quot;]
 **associativeDataId** | [**[]string**](string.md) |  | 
 **elementId** | **string** |  | [default to &quot;&quot;]
 **viewId** | **string** |  | [default to &quot;&quot;]
 **microversionId** | **string** |  | [default to &quot;&quot;]
 **documentMicroversion** | **string** |  | [default to &quot;&quot;]
 **deterministicId** | **string** |  | [default to &quot;&quot;]
 **featureId** | **string** |  | [default to &quot;&quot;]
 **entityId** | **string** |  | [default to &quot;&quot;]
 **occurrenceId** | **string** |  | [default to &quot;&quot;]
 **returnIdTags** | **bool** |  | [default to false]

### Return type

[**BTAppAssociativeDataInfoArray**](BTAppAssociativeDataInfoArray.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestInDocument

> BTListResponseBTInsertableInfo GetLatestInDocument(ctx, did).BetaCapabilityIds(betaCapabilityIds).IncludeParts(includeParts).IncludeSurfaces(includeSurfaces).IncludeWires(includeWires).IncludeSketches(includeSketches).IncludeReferenceFeatures(includeReferenceFeatures).IncludeAssemblies(includeAssemblies).IncludeFeatures(includeFeatures).IncludeFeatureStudios(includeFeatureStudios).IncludePartStudios(includePartStudios).IncludeBlobs(includeBlobs).IncludeMeshes(includeMeshes).IncludeFlattenedBodies(includeFlattenedBodies).AllowedBlobMimeTypes(allowedBlobMimeTypes).MaxFeatureScriptVersion(maxFeatureScriptVersion).IncludeApplications(includeApplications).AllowedApplicationMimeTypes(allowedApplicationMimeTypes).IncludeCompositeParts(includeCompositeParts).IncludeFSTables(includeFSTables).Execute()

insertables for a document

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaCapabilityIds** | [**[]string**](string.md) |  | 
 **includeParts** | **bool** |  | [default to false]
 **includeSurfaces** | **bool** |  | [default to false]
 **includeWires** | **bool** |  | [default to false]
 **includeSketches** | **bool** |  | [default to false]
 **includeReferenceFeatures** | **bool** |  | [default to false]
 **includeAssemblies** | **bool** |  | [default to false]
 **includeFeatures** | **bool** |  | [default to false]
 **includeFeatureStudios** | **bool** |  | [default to false]
 **includePartStudios** | **bool** |  | [default to false]
 **includeBlobs** | **bool** |  | [default to false]
 **includeMeshes** | **bool** |  | [default to false]
 **includeFlattenedBodies** | **bool** |  | [default to false]
 **allowedBlobMimeTypes** | **string** |  | [default to &quot;&quot;]
 **maxFeatureScriptVersion** | **int32** |  | [default to 0]
 **includeApplications** | **bool** |  | [default to false]
 **allowedApplicationMimeTypes** | **string** |  | [default to &quot;&quot;]
 **includeCompositeParts** | **bool** |  | [default to false]
 **includeFSTables** | **bool** |  | [default to false]

### Return type

[**BTListResponseBTInsertableInfo**](BTListResponseBTInsertableInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataSchema

> BTMetadataSchemaInfo GetMetadataSchema(ctx).ObjectType(objectType).OwnerId(ownerId).DocumentId(documentId).OwnerType(ownerType).Execute()



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectType** | **int32** |  | 
 **ownerId** | **string** |  | 
 **documentId** | **string** |  | 
 **ownerType** | **int32** |  | [default to 1]

### Return type

[**BTMetadataSchemaInfo**](BTMetadataSchemaInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProperties

> BTListResponseBTMetadataPropertySummaryInfo GetProperties(ctx).SchemaId(schemaId).OwnerId(ownerId).DocumentId(documentId).OwnerType(ownerType).ObjectType(objectType).Strict(strict).ActiveOnly(activeOnly).Offset(offset).Limit(limit).Execute()



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaId** | **string** |  | 
 **ownerId** | **string** |  | 
 **documentId** | **string** |  | 
 **ownerType** | **int32** |  | [default to 1]
 **objectType** | **int32** |  | 
 **strict** | **bool** |  | [default to false]
 **activeOnly** | **bool** |  | [default to false]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 200]

### Return type

[**BTListResponseBTMetadataPropertySummaryInfo**](BTListResponseBTMetadataPropertySummaryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPropertyInfo

> BTMetadataPropertyInfo GetPropertyInfo(ctx, pid).DocumentId(documentId).SchemaId(schemaId).OwnerId(ownerId).OwnerType(ownerType).ObjectType(objectType).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentId** | **string** |  | 
 **schemaId** | **string** |  | 
 **ownerId** | **string** |  | 
 **ownerType** | **int32** |  | [default to 1]
 **objectType** | **int32** |  | 

### Return type

[**BTMetadataPropertyInfo**](BTMetadataPropertyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> BTMetadataSchemaInfo GetSchema(ctx, sid).DocumentId(documentId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **documentId** | **string** |  | 

### Return type

[**BTMetadataSchemaInfo**](BTMetadataSchemaInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSketchBoundingBoxes

> BTBoundingBoxInfo GetSketchBoundingBoxes(ctx, did, wvm, wvmid, eid, sid).Configuration(configuration).LinkDocumentId(linkDocumentId).Execute()



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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSketchInfo

> GetSketchInfo(ctx, did, wvm, wvmid, eid).Configuration(configuration).SketchId(sketchId).Output3D(output3D).CurvePoints(curvePoints).IncludeGeometry(includeGeometry).LinkDocumentId(linkDocumentId).Execute()



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




 **configuration** | **string** |  | 
 **sketchId** | [**[]string**](string.md) |  | 
 **output3D** | **bool** |  | [default to false]
 **curvePoints** | **bool** |  | [default to false]
 **includeGeometry** | **bool** |  | [default to true]
 **linkDocumentId** | **string** |  | 

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


## GetTessellatedEntities

> GetTessellatedEntities(ctx, did, wvm, wvmid, eid, sid).Configuration(configuration).EntityId(entityId).AngleTolerance(angleTolerance).ChordTolerance(chordTolerance).LinkDocumentId(linkDocumentId).Execute()



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
 **entityId** | [**[]string**](string.md) |  | 
 **angleTolerance** | **float64** |  | 
 **chordTolerance** | **float64** |  | 
 **linkDocumentId** | **string** |  | 

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


## GetWorkflowableTestObject

> BTWorkflowableTestObjectInfo GetWorkflowableTestObject(ctx, oid).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAssociativeData

> BTAppAssociativeDataInfoArray PostAssociativeData(ctx, did, eid, wvm, wvmid).Body(body).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTAppAssociativeDataInfoArray**](BTAppAssociativeDataInfoArray.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionWorkflowableTestObject

> BTWorkflowableTestObjectInfo TransitionWorkflowableTestObject(ctx, oid, transition).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 
**transition** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowableTestObject

> BTWorkflowableTestObjectInfo UpdateWorkflowableTestObject(ctx, oid).BTUpdateWorkflowableTestObjectParams(bTUpdateWorkflowableTestObjectParams).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTUpdateWorkflowableTestObjectParams** | [**BTUpdateWorkflowableTestObjectParams**](BTUpdateWorkflowableTestObjectParams.md) |  | 

### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

