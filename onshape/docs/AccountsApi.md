# \AccountsApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPurchaseNew**](AccountsApi.md#CancelPurchaseNew) | **Delete** /api/accounts/{aid}/purchases/{pid} | Cancel Recurring Subscription
[**ConsumePurchase**](AccountsApi.md#ConsumePurchase) | **Post** /api/accounts/purchases/{pid}/consume | Mark Purchase Consumed For User
[**GetPlanPurchases**](AccountsApi.md#GetPlanPurchases) | **Get** /api/accounts/plans/{planId}/purchases | Get Plan Purchases
[**GetPurchases**](AccountsApi.md#GetPurchases) | **Get** /api/accounts/purchases | Get User&#39;s Appstore Purchases.



## CancelPurchaseNew

> CancelPurchaseNew(ctx, aid, pid).CancelImmediately(cancelImmediately).Execute()

Cancel Recurring Subscription

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **string** |  | 
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPurchaseNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cancelImmediately** | **bool** |  | [default to false]

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


## ConsumePurchase

> BTPurchaseInfo ConsumePurchase(ctx, pid).BTPurchaseUserParams(bTPurchaseUserParams).Execute()

Mark Purchase Consumed For User

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConsumePurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTPurchaseUserParams** | [**BTPurchaseUserParams**](BTPurchaseUserParams.md) |  | 

### Return type

[**BTPurchaseInfo**](BTPurchaseInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanPurchases

> BTListResponseBTPurchaseInfo GetPlanPurchases(ctx, planId).Offset(offset).Limit(limit).Execute()

Get Plan Purchases

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTPurchaseInfo**](BTListResponseBTPurchaseInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPurchases

> []BTPurchaseInfo GetPurchases(ctx).All(all).OwnPurchaseOnly(ownPurchaseOnly).Execute()

Get User's Appstore Purchases.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPurchasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **bool** |  | [default to false]
 **ownPurchaseOnly** | **bool** |  | [default to false]

### Return type

[**[]BTPurchaseInfo**](BTPurchaseInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

