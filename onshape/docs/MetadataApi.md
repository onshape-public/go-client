# \MetadataApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVEOPStandardContentMetadata**](MetadataApi.md#GetVEOPStandardContentMetadata) | **Get** /api/metadata/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/p/{pid} | 
[**GetWMVEMetadata**](MetadataApi.md#GetWMVEMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | 
[**GetWMVEPMetadata**](MetadataApi.md#GetWMVEPMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p/{pid} | 
[**GetWMVEPsMetadata**](MetadataApi.md#GetWMVEPsMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p | 
[**GetWMVEsMetadata**](MetadataApi.md#GetWMVEsMetadata) | **Get** /api/metadata/d/{did}/{wvm}/{wvmid}/e | 
[**GetWVMetadata**](MetadataApi.md#GetWVMetadata) | **Get** /api/metadata/d/{did}/{wv}/{wvid} | 
[**UpdateVEOPStandardContentPartMetadata**](MetadataApi.md#UpdateVEOPStandardContentPartMetadata) | **Post** /api/metadata/standardcontent/d/{did}/v/{vid}/e/{eid}/{otype}/{oid}/p/{pid} | 
[**UpdateWVEMetadata**](MetadataApi.md#UpdateWVEMetadata) | **Post** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid} | 
[**UpdateWVEPMetadata**](MetadataApi.md#UpdateWVEPMetadata) | **Post** /api/metadata/d/{did}/{wvm}/{wvmid}/e/{eid}/p/{pid} | 
[**UpdateWVMetadata**](MetadataApi.md#UpdateWVMetadata) | **Post** /api/metadata/d/{did}/{wv}/{wvid} | 



## GetVEOPStandardContentMetadata

> GetVEOPStandardContentMetadata(ctx, did, vid, eid, otype, oid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).Thumbnail(thumbnail).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 
**eid** | **string** |  | 
**otype** | **string** |  | 
**oid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVEOPStandardContentMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEMetadata

> GetWMVEMetadata(ctx, did, wvm, wvmid, eid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEPMetadata

> GetWMVEPMetadata(ctx, did, wvm, wvmid, eid, pid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Thumbnail(thumbnail).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEPMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEPsMetadata

> GetWMVEPsMetadata(ctx, did, wvm, wvmid, eid).Configuration(configuration).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Thumbnail(thumbnail).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEPsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **configuration** | **string** |  | 
 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWMVEsMetadata

> GetWMVEsMetadata(ctx, did, wvm, wvmid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWMVEsMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **thumbnail** | **bool** |  | [default to false]

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


## GetWVMetadata

> GetWVMetadata(ctx, did, wv, wvid).LinkDocumentId(linkDocumentId).InferMetadataOwner(inferMetadataOwner).Depth(depth).Thumbnail(thumbnail).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWVMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** |  | 
 **inferMetadataOwner** | **bool** |  | [default to false]
 **depth** | **string** |  | [default to &quot;1&quot;]
 **thumbnail** | **bool** |  | [default to false]

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


## UpdateVEOPStandardContentPartMetadata

> UpdateVEOPStandardContentPartMetadata(ctx, did, vid, eid, otype, oid, pid).Body(body).LinkDocumentId(linkDocumentId).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**vid** | **string** |  | 
**eid** | **string** |  | 
**otype** | **string** |  | 
**oid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVEOPStandardContentPartMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | **string** |  | 
 **linkDocumentId** | **string** |  | 

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


## UpdateWVEMetadata

> UpdateWVEMetadata(ctx, did, wvm, wvmid, eid).Body(body).Configuration(configuration).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVEMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 
 **configuration** | **string** |  | 

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


## UpdateWVEPMetadata

> UpdateWVEPMetadata(ctx, did, wvm, wvmid, eid, pid, subResource).Body(body).Configuration(configuration).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 
**pid** | **string** |  | 
**subResource** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVEPMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | **string** |  | 
 **configuration** | **string** |  | 

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


## UpdateWVMetadata

> UpdateWVMetadata(ctx, did, wv, wvid).Body(body).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWVMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

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

