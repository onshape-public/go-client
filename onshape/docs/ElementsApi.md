# \ElementsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyElementFromSourceDocument**](ElementsApi.md#CopyElementFromSourceDocument) | **Post** /api/elements/copyelement/{did}/workspace/{wid} | Copy Element
[**DecodeConfiguration**](ElementsApi.md#DecodeConfiguration) | **Get** /api/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configurationencodings/{cid} | Decode Configuration String
[**DeleteElement**](ElementsApi.md#DeleteElement) | **Delete** /api/elements/d/{did}/w/{wid}/e/{eid} | Delete Element
[**EncodeConfigurationMap**](ElementsApi.md#EncodeConfigurationMap) | **Post** /api/elements/d/{did}/e/{eid}/configurationencodings | Encode Configuration
[**GetConfiguration**](ElementsApi.md#GetConfiguration) | **Get** /api/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Get Configuration
[**GetElementTranslatorFormatsByVersionOrWorkspace**](ElementsApi.md#GetElementTranslatorFormatsByVersionOrWorkspace) | **Get** /api/elements/translatorFormats/{did}/{wv}/{wvid}/{eid} | Get Element Translator Formats
[**UpdateConfiguration**](ElementsApi.md#UpdateConfiguration) | **Post** /api/elements/d/{did}/{wvm}/{wvmid}/e/{eid}/configuration | Update Configuration
[**UpdateReferences**](ElementsApi.md#UpdateReferences) | **Post** /api/elements/d/{did}/w/{wid}/e/{eid}/updatereferences | Update or replace node references



## CopyElementFromSourceDocument

> BTDocumentElementInfo CopyElementFromSourceDocument(ctx, did, wid).BTCopyElementParams(bTCopyElementParams).Execute()

Copy Element

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyElementFromSourceDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTCopyElementParams** | [**BTCopyElementParams**](BTCopyElementParams.md) |  | 

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


## DecodeConfiguration

> BTConfigurationInfo DecodeConfiguration(ctx, did, wvm, wvmid, eid, cid).LinkDocumentId(linkDocumentId).IncludeDisplay(includeDisplay).ConfigurationIsId(configurationIsId).Execute()

Decode Configuration String

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecodeConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **linkDocumentId** | **string** |  | 
 **includeDisplay** | **bool** |  | [default to false]
 **configurationIsId** | **bool** |  | [default to false]

### Return type

[**BTConfigurationInfo**](BTConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteElement

> DeleteElement(ctx, did, wid, eid).Execute()

Delete Element

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteElementRequest struct via the builder pattern


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


## EncodeConfigurationMap

> BTEncodedConfigurationInfo EncodeConfigurationMap(ctx, did, eid).BTConfigurationParams(bTConfigurationParams).VersionId(versionId).LinkDocumentId(linkDocumentId).Execute()

Encode Configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncodeConfigurationMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTConfigurationParams** | [**BTConfigurationParams**](BTConfigurationParams.md) |  | 
 **versionId** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTEncodedConfigurationInfo**](BTEncodedConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> BTConfigurationInfo GetConfiguration(ctx, did, wvm, wvmid, eid).Execute()

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

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTConfigurationInfo**](BTConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetElementTranslatorFormatsByVersionOrWorkspace

> []BTModelFormatInfo GetElementTranslatorFormatsByVersionOrWorkspace(ctx, did, wv, wvid, eid).CheckContent(checkContent).Configuration(configuration).Execute()

Get Element Translator Formats

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetElementTranslatorFormatsByVersionOrWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **checkContent** | **bool** |  | [default to true]
 **configuration** | **string** |  | [default to &quot;&quot;]

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


## UpdateConfiguration

> BTConfigurationInfo UpdateConfiguration(ctx, did, wvm, wvmid, eid).Body(body).Execute()

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

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTConfigurationInfo**](BTConfigurationInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReferences

> UpdateReferences(ctx, did, wid, eid).BTUpdateReferenceParams(bTUpdateReferenceParams).Execute()

Update or replace node references

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateReferenceParams** | [**BTUpdateReferenceParams**](BTUpdateReferenceParams.md) |  | 

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

