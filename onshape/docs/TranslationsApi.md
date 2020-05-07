# \TranslationsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranslation**](TranslationsApi.md#CreateTranslation) | **Post** /api/translations/d/{did}/w/{wid} | Create translation from upload.
[**DeleteTranslation**](TranslationsApi.md#DeleteTranslation) | **Delete** /api/translations/{tid} | 
[**GetAllTranslatorFormats**](TranslationsApi.md#GetAllTranslatorFormats) | **Get** /api/translations/translationformats | 
[**GetDocumentTranslations**](TranslationsApi.md#GetDocumentTranslations) | **Get** /api/translations/d/{did} | 
[**GetTranslation**](TranslationsApi.md#GetTranslation) | **Get** /api/translations/{tid} | 



## CreateTranslation

> BTTranslationRequestInfo CreateTranslation(ctx, did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FileBodyWithDetails(fileBodyWithDetails).FileContentLength(fileContentLength).FileDetail(fileDetail).FlattenAssemblies(flattenAssemblies).FormatName(formatName).IsyAxisIsUp(isyAxisIsUp).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).OwnerType(ownerType).ParentId(parentId).ProjectId(projectId).Public(public).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).Execute()

Create translation from upload.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationRequest struct via the builder pattern


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

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranslation

> DeleteTranslation(ctx, tid).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTranslationRequest struct via the builder pattern


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


## GetAllTranslatorFormats

> []BTModelFormatFullInfo GetAllTranslatorFormats(ctx).Execute()



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTranslatorFormatsRequest struct via the builder pattern


### Return type

[**[]BTModelFormatFullInfo**](BTModelFormatFullInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentTranslations

> BTListResponseBTTranslationRequestInfo GetDocumentTranslations(ctx, did).Offset(offset).Limit(limit).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTTranslationRequestInfo**](BTListResponseBTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslation

> BTTranslationRequestInfo GetTranslation(ctx, tid).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

