# \ThumbnailsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationThumbnails**](ThumbnailsApi.md#DeleteApplicationThumbnails) | **Delete** /api/thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | 
[**GetConfiguredElementThumbnailWithSize**](ThumbnailsApi.md#GetConfiguredElementThumbnailWithSize) | **Get** /api/thumbnails/d/{did}/w/{wid}/e/{eid}/c/{cid}/s/{sz} | 
[**GetDocumentThumbnail**](ThumbnailsApi.md#GetDocumentThumbnail) | **Get** /api/thumbnails/d/{did}/w/{wid} | 
[**GetDocumentThumbnailWithSize**](ThumbnailsApi.md#GetDocumentThumbnailWithSize) | **Get** /api/thumbnails/d/{did}/w/{wid}/s/{sz} | 
[**GetElementThumbnail**](ThumbnailsApi.md#GetElementThumbnail) | **Get** /api/thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | 
[**GetElementThumbnailWithApiConfiguration**](ThumbnailsApi.md#GetElementThumbnailWithApiConfiguration) | **Get** /api/thumbnails/d/{did}/w/{wid}/e/{eid}/ac/{cid}/s/{sz} | 
[**GetElementThumbnailWithSize**](ThumbnailsApi.md#GetElementThumbnailWithSize) | **Get** /api/thumbnails/d/{did}/w/{wid}/e/{eid}/s/{sz} | 
[**GetThumbnailForDocument**](ThumbnailsApi.md#GetThumbnailForDocument) | **Get** /api/thumbnails/d/{did} | 
[**GetThumbnailForDocumentAndVersion**](ThumbnailsApi.md#GetThumbnailForDocumentAndVersion) | **Get** /api/thumbnails/d/{did}/v/{vid} | 
[**GetThumbnailForDocumentAndVersionOld**](ThumbnailsApi.md#GetThumbnailForDocumentAndVersionOld) | **Get** /api/thumbnails/document/{did}/version/{vid} | 
[**GetThumbnailForDocumentOld**](ThumbnailsApi.md#GetThumbnailForDocumentOld) | **Get** /api/thumbnails/document/{did} | 
[**SetApplicationElementThumbnail**](ThumbnailsApi.md#SetApplicationElementThumbnail) | **Post** /api/thumbnails/d/{did}/{wv}/{wvid}/e/{eid} | 



## DeleteApplicationThumbnails

> DeleteApplicationThumbnails(ctx, did, wv, wvid, eid).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationThumbnailsRequest struct via the builder pattern


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


## GetConfiguredElementThumbnailWithSize

> GetConfiguredElementThumbnailWithSize(ctx, did, wid, eid, cid, sz).T(t).RejectEmpty(rejectEmpty).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**cid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguredElementThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **t** | **string** |  | 
 **rejectEmpty** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentThumbnail

> BTThumbnailInfo GetDocumentThumbnail(ctx, did, wid).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentThumbnailWithSize

> GetDocumentThumbnailWithSize(ctx, did, wid, sz).T(t).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **t** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnail

> BTThumbnailInfo GetElementThumbnail(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** |  | 

### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnailWithApiConfiguration

> GetElementThumbnailWithApiConfiguration(ctx, did, wid, eid, cid, sz).T(t).RejectEmpty(rejectEmpty).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**cid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailWithApiConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **t** | **string** |  | 
 **rejectEmpty** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementThumbnailWithSize

> GetElementThumbnailWithSize(ctx, did, wid, eid, sz).T(t).RejectEmpty(rejectEmpty).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 
**sz** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementThumbnailWithSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **t** | **string** |  | 
 **rejectEmpty** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocument

> BTThumbnailInfo GetThumbnailForDocument(ctx, did).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThumbnailForDocumentAndVersion

> GetThumbnailForDocumentAndVersion(ctx, did, vid).LinkDocumentId(linkDocumentId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentAndVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetThumbnailForDocumentAndVersionOld

> GetThumbnailForDocumentAndVersionOld(ctx, did, vid).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentAndVersionOldRequest struct via the builder pattern


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


## GetThumbnailForDocumentOld

> BTThumbnailInfo GetThumbnailForDocumentOld(ctx, did).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThumbnailForDocumentOldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTThumbnailInfo**](BTThumbnailInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApplicationElementThumbnail

> SetApplicationElementThumbnail(ctx, did, wv, wvid, eid).BTApplicationElementThumbnailParamsArray(bTApplicationElementThumbnailParamsArray).Overwrite(overwrite).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApplicationElementThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTApplicationElementThumbnailParamsArray** | [**BTApplicationElementThumbnailParamsArray**](BTApplicationElementThumbnailParamsArray.md) |  | 
 **overwrite** | **bool** |  | [default to false]

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

