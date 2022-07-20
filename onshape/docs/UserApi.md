# \UserApi

All URIs are relative to *https://staging.dev.onshape.com/api/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSettings**](UserApi.md#GetUserSettings) | **Get** /users/{uid}/settings | Retrieve user settings by user ID.
[**GetUserSettingsCurrentLoggedInUser**](UserApi.md#GetUserSettingsCurrentLoggedInUser) | **Get** /users/settings | Get user settings for the currently signed-in user if there is one, or else return the default settings.
[**Session**](UserApi.md#Session) | **Post** /users/session | Check if current user is signed-in.Information returned depends on OAuth2ReadPII scope.
[**SessionInfo**](UserApi.md#SessionInfo) | **Get** /users/sessioninfo | Check if current user is signed-in. Information returned depends on OAuth2ReadPII scope.



## GetUserSettings

> BTUserSettingsInfo GetUserSettings(ctx, uid).Includematerials(includematerials).Execute()

Retrieve user settings by user ID.

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

Get user settings for the currently signed-in user if there is one, or else return the default settings.

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

Check if current user is signed-in.Information returned depends on OAuth2ReadPII scope.

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

Check if current user is signed-in. Information returned depends on OAuth2ReadPII scope.

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

