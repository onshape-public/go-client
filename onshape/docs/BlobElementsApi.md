# \BlobElementsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlobTranslation**](BlobElementsApi.md#CreateBlobTranslation) | **Post** /api/blobelements/d/{did}/{wv}/{wvid}/e/{eid}/translations | Create Translation.
[**DownloadFileWorkspace**](BlobElementsApi.md#DownloadFileWorkspace) | **Get** /api/blobelements/d/{did}/w/{wid}/e/{eid} | Download File From Blob Element.
[**UpdateUnits**](BlobElementsApi.md#UpdateUnits) | **Post** /api/blobelements/d/{did}/w/{wid}/e/{eid}/units | Update Mesh Units.
[**UploadFileCreateElement**](BlobElementsApi.md#UploadFileCreateElement) | **Post** /api/blobelements/d/{did}/w/{wid} | Upload file to new element.
[**UploadFileUpdateElement**](BlobElementsApi.md#UploadFileUpdateElement) | **Post** /api/blobelements/d/{did}/w/{wid}/e/{eid} | Update Blob Element.



## CreateBlobTranslation

> BTTranslationRequestInfo CreateBlobTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).Execute()

Create Translation.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobTranslationRequest struct via the builder pattern


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


## DownloadFileWorkspace

> *os.File DownloadFileWorkspace(ctx, did, wid, eid).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).LinkDocumentId(linkDocumentId).Execute()

Download File From Blob Element.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentDisposition** | **string** |  | 
 **ifNoneMatch** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUnits

> BTDocumentElementProcessingInfo UpdateUnits(ctx, did, eid, wid).BTUpdateMeshUnitsParams(bTUpdateMeshUnitsParams).Execute()

Update Mesh Units.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateMeshUnitsParams** | [**BTUpdateMeshUnitsParams**](BTUpdateMeshUnitsParams.md) |  | 

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileCreateElement

> BTDocumentElementProcessingInfo UploadFileCreateElement(ctx, did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FileBodyWithDetails(fileBodyWithDetails).FileContentLength(fileContentLength).FileDetail(fileDetail).FlattenAssemblies(flattenAssemblies).FormatName(formatName).IsyAxisIsUp(isyAxisIsUp).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).OwnerType(ownerType).ParentId(parentId).ProjectId(projectId).Public(public).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).Execute()

Upload file to new element.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allowFaultyParts** | **bool** |  | 
 **createComposite** | **bool** |  | 
 **createDrawingIfPossible** | **bool** |  | 
 **encodedFilename** | **string** |  | 
 **extractAssemblyHierarchy** | **bool** |  | 
 **file** | ***os.File** |  | 
 **fileBodyWithDetails** | [**FormDataBodyPart**](FormDataBodyPart.md) |  | 
 **fileContentLength** | **int64** |  | 
 **fileDetail** | [**FormDataContentDisposition**](FormDataContentDisposition.md) |  | 
 **flattenAssemblies** | **bool** |  | 
 **formatName** | **string** |  | 
 **isyAxisIsUp** | **bool** |  | 
 **joinAdjacentSurfaces** | **bool** |  | 
 **locationElementId** | **string** |  | 
 **locationGroupId** | **string** |  | 
 **locationPosition** | **int32** |  | 
 **notifyUser** | **bool** |  | 
 **ownerId** | **string** |  | 
 **ownerType** | **string** |  | 
 **parentId** | **string** |  | 
 **projectId** | **string** |  | 
 **public** | **bool** |  | 
 **splitAssembliesIntoMultipleDocuments** | **bool** |  | 
 **storeInDocument** | **bool** |  | 
 **translate** | **bool** |  | 
 **unit** | **string** |  | 
 **uploadId** | **string** |  | 
 **versionString** | **string** |  | 

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileUpdateElement

> BTDocumentElementProcessingInfo UploadFileUpdateElement(ctx, did, eid, wid).ParentChangeId(parentChangeId).Execute()

Update Blob Element.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileUpdateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **parentChangeId** | **string** |  | 

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

