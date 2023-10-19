# \UserApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSettings**](UserApi.md#GetUserSettings) | **Get** /users/{uid}/settings | Get the user settings for any user in your organization (admins only).
[**GetUserSettingsCurrentLoggedInUser**](UserApi.md#GetUserSettingsCurrentLoggedInUser) | **Get** /users/settings | Get the user settings for the signed-in user (i.e., you) for the current session.
[**Session**](UserApi.md#Session) | **Post** /users/session | Authenticate a user&#39;s Onshape credentials, and create a session.
[**SessionInfo**](UserApi.md#SessionInfo) | **Get** /users/sessioninfo | Get the session information for an authenticated (signed-in) user.



## GetUserSettings

> BTUserSettingsInfo GetUserSettings(ctx, uid).Includematerials(includematerials).Execute()

Get the user settings for any user in your organization (admins only).



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uid := "uid_example" // string | 
    includematerials := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetUserSettings(context.Background(), uid).Includematerials(includematerials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettings`: BTUserSettingsInfo
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserSettings`: %v\n", resp)
}
```

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

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserSettingsCurrentLoggedInUser

> BTUserSettingsInfo GetUserSettingsCurrentLoggedInUser(ctx).Includematerials(includematerials).Execute()

Get the user settings for the signed-in user (i.e., you) for the current session.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    includematerials := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetUserSettingsCurrentLoggedInUser(context.Background()).Includematerials(includematerials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserSettingsCurrentLoggedInUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserSettingsCurrentLoggedInUser`: BTUserSettingsInfo
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserSettingsCurrentLoggedInUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSettingsCurrentLoggedInUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includematerials** | **bool** |  | [default to true]

### Return type

[**BTUserSettingsInfo**](BTUserSettingsInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Session

> map[string]interface{} Session(ctx).BTLoginParams(bTLoginParams).Execute()

Authenticate a user's Onshape credentials, and create a session.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bTLoginParams := *openapiclient.NewBTLoginParams() // BTLoginParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.Session(context.Background()).BTLoginParams(bTLoginParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.Session``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Session`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserApi.Session`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTLoginParams** | [**BTLoginParams**](BTLoginParams.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionInfo

> BTUserOAuth2SummaryInfo SessionInfo(ctx).Execute()

Get the session information for an authenticated (signed-in) user.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.SessionInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.SessionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SessionInfo`: BTUserOAuth2SummaryInfo
    fmt.Fprintf(os.Stdout, "Response from `UserApi.SessionInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSessionInfoRequest struct via the builder pattern


### Return type

[**BTUserOAuth2SummaryInfo**](BTUserOAuth2SummaryInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

