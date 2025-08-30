# \CompanyApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGlobalPermissionsForIdentity**](CompanyApi.md#AddGlobalPermissionsForIdentity) | **Post** /companies/{cid}/globalpermission/{type}/{id} | Add one or more global permissions to company user or team.
[**AddUserToCompany**](CompanyApi.md#AddUserToCompany) | **Post** /companies/{cid}/users | Add a user to a company.
[**ClearGlobalPermissions**](CompanyApi.md#ClearGlobalPermissions) | **Delete** /companies/{cid}/globalpermission/{type}/{id} | Remove global permissions for a company user or team.
[**FindCompany**](CompanyApi.md#FindCompany) | **Get** /companies | Get all companies to which the specified user belongs.
[**GetCompany**](CompanyApi.md#GetCompany) | **Get** /companies/{cid} | Get company information by company ID.
[**GetDocumentsByName**](CompanyApi.md#GetDocumentsByName) | **Get** /companies/{cid}/documentsbyname | Get document by exact document name.
[**RemoveUserFromCompany**](CompanyApi.md#RemoveUserFromCompany) | **Delete** /companies/{cid}/users/{uid} | Remove a user from a company, company teams, and all the direct shares.
[**UpdateCompanyUser**](CompanyApi.md#UpdateCompanyUser) | **Post** /companies/{cid}/users/{uid} | Update the company&#39;s information for a user.



## AddGlobalPermissionsForIdentity

> BTGlobalPermissionInfo AddGlobalPermissionsForIdentity(ctx, cid, type_, id).RequestBody(requestBody).Execute()

Add one or more global permissions to company user or team.



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
    cid := "cid_example" // string | Company ID
    type_ := int32(56) // int32 | `0`: USER | `1`: TEAM
    id := "id_example" // string | User ID or Team ID, depending on `type`
    requestBody := []int32{int32(123)} // []int32 | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.AddGlobalPermissionsForIdentity(context.Background(), cid, type_, id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.AddGlobalPermissionsForIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddGlobalPermissionsForIdentity`: BTGlobalPermissionInfo
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.AddGlobalPermissionsForIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | Company ID | 
**type_** | **int32** | &#x60;0&#x60;: USER | &#x60;1&#x60;: TEAM | 
**id** | **string** | User ID or Team ID, depending on &#x60;type&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGlobalPermissionsForIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestBody** | **[]int32** |  | 

### Return type

[**BTGlobalPermissionInfo**](BTGlobalPermissionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToCompany

> BTCompanyUserInfo AddUserToCompany(ctx, cid).BTCompanyUserParams(bTCompanyUserParams).Execute()

Add a user to a company.



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
    cid := "cid_example" // string | 
    bTCompanyUserParams := *openapiclient.NewBTCompanyUserParams() // BTCompanyUserParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.AddUserToCompany(context.Background(), cid).BTCompanyUserParams(bTCompanyUserParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.AddUserToCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToCompany`: BTCompanyUserInfo
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.AddUserToCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTCompanyUserParams** | [**BTCompanyUserParams**](BTCompanyUserParams.md) |  | 

### Return type

[**BTCompanyUserInfo**](BTCompanyUserInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearGlobalPermissions

> map[string]interface{} ClearGlobalPermissions(ctx, cid, type_, id).Permission(permission).Execute()

Remove global permissions for a company user or team.



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
    cid := "cid_example" // string | Company ID
    type_ := int32(56) // int32 | `0`: USER | `1`: TEAM
    id := "id_example" // string | User ID or Team ID, depending on `type`
    permission := []int32{int32(123)} // []int32 | List of global permissions to grant. See [Onshape Help: Global Permissions](https://cad.onshape.com/help/Content/Plans/global_permissions.htm#Assignin) for details on each of the available permissions.   * `0`: Manage role based access control   * `1`: Manage users, teams, and aliases   * `2`: Enterprise administrator   * `3`: Permanently delete   * `4`: Analytics administrator   * `5`: Invite guest users   * `6`: Create projects   * `7`: Approve releases   * `8`: Enable link sharing   * `9`: Create releases   * `10`: Allow access to the App Store   * `11`: Create documents and folders in the Enterprise root   * `12`: Allow access to public documents   * `17`: Manage non-geometric items   * `18`: Manage workflows   * `19`: Transfer documents out of Enterprise   * `20`: Sync to Arena   * `21`: Create tasks   * `22`: Manage standard content metadata   * `23`: Workspace protection permissions   * `24`: Import files   * `25`: Use revision tools  * `26`: Export files   (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.ClearGlobalPermissions(context.Background(), cid, type_, id).Permission(permission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.ClearGlobalPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearGlobalPermissions`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.ClearGlobalPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | Company ID | 
**type_** | **int32** | &#x60;0&#x60;: USER | &#x60;1&#x60;: TEAM | 
**id** | **string** | User ID or Team ID, depending on &#x60;type&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearGlobalPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **permission** | **[]int32** | List of global permissions to grant. See [Onshape Help: Global Permissions](https://cad.onshape.com/help/Content/Plans/global_permissions.htm#Assignin) for details on each of the available permissions.   * &#x60;0&#x60;: Manage role based access control   * &#x60;1&#x60;: Manage users, teams, and aliases   * &#x60;2&#x60;: Enterprise administrator   * &#x60;3&#x60;: Permanently delete   * &#x60;4&#x60;: Analytics administrator   * &#x60;5&#x60;: Invite guest users   * &#x60;6&#x60;: Create projects   * &#x60;7&#x60;: Approve releases   * &#x60;8&#x60;: Enable link sharing   * &#x60;9&#x60;: Create releases   * &#x60;10&#x60;: Allow access to the App Store   * &#x60;11&#x60;: Create documents and folders in the Enterprise root   * &#x60;12&#x60;: Allow access to public documents   * &#x60;17&#x60;: Manage non-geometric items   * &#x60;18&#x60;: Manage workflows   * &#x60;19&#x60;: Transfer documents out of Enterprise   * &#x60;20&#x60;: Sync to Arena   * &#x60;21&#x60;: Create tasks   * &#x60;22&#x60;: Manage standard content metadata   * &#x60;23&#x60;: Workspace protection permissions   * &#x60;24&#x60;: Import files   * &#x60;25&#x60;: Use revision tools  * &#x60;26&#x60;: Export files   | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCompany

> BTListResponseBTCompanyInfo FindCompany(ctx).Uid(uid).ActiveOnly(activeOnly).IncludeAll(includeAll).Execute()

Get all companies to which the specified user belongs.



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
    uid := "uid_example" // string |  (optional)
    activeOnly := true // bool |  (optional) (default to true)
    includeAll := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.FindCompany(context.Background()).Uid(uid).ActiveOnly(activeOnly).IncludeAll(includeAll).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.FindCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindCompany`: BTListResponseBTCompanyInfo
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.FindCompany`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **string** |  | 
 **activeOnly** | **bool** |  | [default to true]
 **includeAll** | **bool** |  | [default to false]

### Return type

[**BTListResponseBTCompanyInfo**](BTListResponseBTCompanyInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompany

> BTCompanyInfo GetCompany(ctx, cid).Execute()

Get company information by company ID.

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
    cid := "cid_example" // string | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.GetCompany(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.GetCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompany`: BTCompanyInfo
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.GetCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTCompanyInfo**](BTCompanyInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentsByName

> []BTGlobalTreeNodeSummaryInfo GetDocumentsByName(ctx, cid).Name(name).Execute()

Get document by exact document name.



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
    cid := "cid_example" // string | 
    name := "name_example" // string | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.GetDocumentsByName(context.Background(), cid).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.GetDocumentsByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentsByName`: []BTGlobalTreeNodeSummaryInfo
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.GetDocumentsByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentsByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** |  | 

### Return type

[**[]BTGlobalTreeNodeSummaryInfo**](BTGlobalTreeNodeSummaryInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromCompany

> map[string]interface{} RemoveUserFromCompany(ctx, cid, uid).RemoveFromTeams(removeFromTeams).RemoveDirectShares(removeDirectShares).Execute()

Remove a user from a company, company teams, and all the direct shares.

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
    cid := "cid_example" // string | 
    uid := "uid_example" // string | 
    removeFromTeams := true // bool |  (optional) (default to true)
    removeDirectShares := true // bool |  (optional) (default to true)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.RemoveUserFromCompany(context.Background(), cid, uid).RemoveFromTeams(removeFromTeams).RemoveDirectShares(removeDirectShares).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.RemoveUserFromCompany``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserFromCompany`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.RemoveUserFromCompany`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**uid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **removeFromTeams** | **bool** |  | [default to true]
 **removeDirectShares** | **bool** |  | [default to true]

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCompanyUser

> BTCompanyUserInfo UpdateCompanyUser(ctx, cid, uid).BTCompanyUserParams(bTCompanyUserParams).Execute()

Update the company's information for a user.



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
    cid := "cid_example" // string | Company ID
    uid := "uid_example" // string | User ID
    bTCompanyUserParams := *openapiclient.NewBTCompanyUserParams() // BTCompanyUserParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.CompanyApi.UpdateCompanyUser(context.Background(), cid, uid).BTCompanyUserParams(bTCompanyUserParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.UpdateCompanyUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCompanyUser`: BTCompanyUserInfo
    fmt.Fprintf(os.Stdout, "Response from `CompanyApi.UpdateCompanyUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | Company ID | 
**uid** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTCompanyUserParams** | [**BTCompanyUserParams**](BTCompanyUserParams.md) |  | 

### Return type

[**BTCompanyUserInfo**](BTCompanyUserInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

