# \AppElementsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommitTransaction**](AppElementsApi.md#CommitTransaction) | **Post** /api/appelements/d/{did}/w/{wid}/e/{eid}/transactions/{tid} | Commit Transaction
[**CreateElement**](AppElementsApi.md#CreateElement) | **Post** /api/appelements/d/{did}/w/{wid} | Create Element.
[**CreateReference**](AppElementsApi.md#CreateReference) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references | Create Reference
[**DeleteAppElementContent**](AppElementsApi.md#DeleteAppElementContent) | **Delete** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/subelements/{sid} | Delete a Sub-element
[**DeleteReference**](AppElementsApi.md#DeleteReference) | **Delete** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Delete Reference
[**GetAppElementHistory**](AppElementsApi.md#GetAppElementHistory) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/history | Get History
[**GetSubElementContent**](AppElementsApi.md#GetSubElementContent) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Get Content
[**GetSubelementIds**](AppElementsApi.md#GetSubelementIds) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content/ids | Get Sub-element IDs
[**ResolveReference**](AppElementsApi.md#ResolveReference) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Resolve Reference
[**ResolveReferences**](AppElementsApi.md#ResolveReferences) | **Get** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/resolvereferences | Resolve references.
[**StartTransaction**](AppElementsApi.md#StartTransaction) | **Post** /api/appelements/d/{did}/w/{wid}/e/{eid}/transactions | Start Transaction
[**UpdateAppElement**](AppElementsApi.md#UpdateAppElement) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/content | Update Element
[**UpdateReference**](AppElementsApi.md#UpdateReference) | **Post** /api/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/references/{rid} | Update Reference



## CommitTransaction

> BTAppElementModifyInfo CommitTransaction(ctx, did, eid, wid, tid).BTAppElementCommitTransactionParams(bTAppElementCommitTransactionParams).Execute()

Commit Transaction

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommitTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementCommitTransactionParams** | [**BTAppElementCommitTransactionParams**](BTAppElementCommitTransactionParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateElement

> BTAppElementModifyInfo CreateElement(ctx, did, wid).BTAppElementParams(bTAppElementParams).Execute()

Create Element.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTAppElementParams** | [**BTAppElementParams**](BTAppElementParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReference

> BTAppElementReferenceInfo CreateReference(ctx, did, eid, wvm, wvmid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Create Reference

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementReferenceParams** | [**BTAppElementReferenceParams**](BTAppElementReferenceParams.md) |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAppElementContent

> BTAppElementModifyInfo DeleteAppElementContent(ctx, did, eid, wvm, wvmid, sid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete a Sub-element

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**sid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppElementContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReference

> BTAppElementReferenceInfo DeleteReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).Description(description).Execute()

Delete Reference

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **description** | **string** |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppElementHistory

> BTAppElementHistoryInfo GetAppElementHistory(ctx, did, eid, wvm, wvmid).Execute()

Get History

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppElementHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTAppElementHistoryInfo**](BTAppElementHistoryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubElementContent

> BTAppElementContentInfo GetSubElementContent(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).BaseChangeId(baseChangeId).SubelementId(subelementId).LinkDocumentId(linkDocumentId).Execute()

Get Content

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubElementContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 
 **baseChangeId** | **string** |  | 
 **subelementId** | **string** |  | 
 **linkDocumentId** | **string** |  | 

### Return type

[**BTAppElementContentInfo**](BTAppElementContentInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubelementIds

> BTAppElementModifyInfo GetSubelementIds(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ChangeId(changeId).Execute()

Get Sub-element IDs

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubelementIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **changeId** | **string** |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReference

> BTAppElementReferenceResolveInfo ResolveReference(ctx, did, eid, wvm, wvmid, rid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).Execute()

Resolve Reference

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **includeInternal** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 

### Return type

[**BTAppElementReferenceResolveInfo**](BTAppElementReferenceResolveInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveReferences

> BTAppElementReferencesResolveInfo ResolveReferences(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).IncludeInternal(includeInternal).LinkDocumentId(linkDocumentId).ReferenceIds(referenceIds).Execute()

Resolve references.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | 
 **parentChangeId** | **string** |  | 
 **includeInternal** | **bool** |  | [default to false]
 **linkDocumentId** | **string** |  | 
 **referenceIds** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppElementReferencesResolveInfo**](BTAppElementReferencesResolveInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTransaction

> BTAppElementModifyInfo StartTransaction(ctx, did, eid, wid).BTAppElementStartTransactionParams(bTAppElementStartTransactionParams).Execute()

Start Transaction

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAppElementStartTransactionParams** | [**BTAppElementStartTransactionParams**](BTAppElementStartTransactionParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppElement

> BTAppElementModifyInfo UpdateAppElement(ctx, did, eid, wvm, wvmid).BTAppElementUpdateParams(bTAppElementUpdateParams).Execute()

Update Element

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTAppElementUpdateParams** | [**BTAppElementUpdateParams**](BTAppElementUpdateParams.md) |  | 

### Return type

[**BTAppElementModifyInfo**](BTAppElementModifyInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReference

> BTAppElementReferenceInfo UpdateReference(ctx, did, eid, wvm, wvmid, rid).BTAppElementReferenceParams(bTAppElementReferenceParams).Execute()

Update Reference

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**rid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bTAppElementReferenceParams** | [**BTAppElementReferenceParams**](BTAppElementReferenceParams.md) |  | 

### Return type

[**BTAppElementReferenceInfo**](BTAppElementReferenceInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

