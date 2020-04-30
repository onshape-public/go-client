# \RevisionsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnumerateRevisions**](RevisionsApi.md#EnumerateRevisions) | **Get** /api/revisions/companies/{cid} | 
[**GetLatestInDocumentOrCompany**](RevisionsApi.md#GetLatestInDocumentOrCompany) | **Get** /api/revisions/{cd}/{cdid}/p/{pnum}/latest | 
[**GetRevisionHistoryInCompany**](RevisionsApi.md#GetRevisionHistoryInCompany) | **Get** /api/revisions/companies/{cid}/partnumber/{pnum} | 



## EnumerateRevisions

> BTListResponseBTRevisionInfo EnumerateRevisions(ctx, cid).ElementType(elementType).Limit(limit).Offset(offset).LatestOnly(latestOnly).After(after).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnumerateRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementType** | **int32** |  | 
 **limit** | **int32** |  | [default to 20]
 **offset** | **int32** |  | [default to 0]
 **latestOnly** | **bool** |  | [default to false]
 **after** | **int64** |  | 

### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestInDocumentOrCompany

> BTRevisionInfo GetLatestInDocumentOrCompany(ctx, cd, cdid, pnum).Et(et).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cd** | **string** |  | 
**cdid** | **string** |  | 
**pnum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestInDocumentOrCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **et** | **string** |  | 

### Return type

[**BTRevisionInfo**](BTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionHistoryInCompany

> BTListResponseBTRevisionInfo GetRevisionHistoryInCompany(ctx, cid, pnum).ElementType(elementType).FillApprovers(fillApprovers).FillExportPermission(fillExportPermission).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**pnum** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionHistoryInCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementType** | **string** |  | 
 **fillApprovers** | **bool** |  | [default to false]
 **fillExportPermission** | **bool** |  | [default to false]

### Return type

[**BTListResponseBTRevisionInfo**](BTListResponseBTRevisionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

