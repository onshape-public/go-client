# \UsersApi

All URIs are relative to *https://cad.onshape.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSettings**](UsersApi.md#GetUserSettings) | **Get** /api/users/{uid}/settings | 
[**GetUserSettingsCurrentLoggedInUser**](UsersApi.md#GetUserSettingsCurrentLoggedInUser) | **Get** /api/users/settings | 
[**SessionInfo**](UsersApi.md#SessionInfo) | **Get** /api/users/sessioninfo | 



## GetUserSettings

> BTUserSettingsInfo GetUserSettings(ctx, uid).Includematerials(includematerials).Execute()



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includematerials** | **bool** |  | [default to true]

### Return type

[**BTUserSettingsInfo**](BTUserSettingsInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsCurrentLoggedInUser

> BTUserSettingsInfo GetUserSettingsCurrentLoggedInUser(ctx).Includematerials(includematerials).Execute()



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsCurrentLoggedInUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includematerials** | **bool** |  | [default to true]

### Return type

[**BTUserSettingsInfo**](BTUserSettingsInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionInfo

> BTUserOAuth2SummaryInfo SessionInfo(ctx).Execute()



### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSessionInfoRequest struct via the builder pattern


### Return type

[**BTUserOAuth2SummaryInfo**](BTUserOAuth2SummaryInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.onshape.v1+json;charset=UTF-8;qs=0.1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

