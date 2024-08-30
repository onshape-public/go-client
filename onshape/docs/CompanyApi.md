# \CompanyApi

All URIs are relative to *https://cad.onshape.com/api/v9*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToCompany**](CompanyApi.md#AddUserToCompany) | **Post** /companies/{cid}/users | Add a user to a company.
[**FindCompany**](CompanyApi.md#FindCompany) | **Get** /companies | Get all companies to which the specified user belongs.
[**GetCompany**](CompanyApi.md#GetCompany) | **Get** /companies/{cid} | Get company information by company ID.
[**GetDocumentsByName**](CompanyApi.md#GetDocumentsByName) | **Get** /companies/{cid}/documentsbyname | Get document by exact document name.
[**RemoveUserFromCompany**](CompanyApi.md#RemoveUserFromCompany) | **Delete** /companies/{cid}/users/{uid} | Remove a user from a company, company teams, and all the direct shares.
[**UpdateCompanyUser**](CompanyApi.md#UpdateCompanyUser) | **Post** /companies/{cid}/users/{uid} | Update the company&#39;s information for a user.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

> []BTDocumentSummaryInfo GetDocumentsByName(ctx, cid).Name(name).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompanyApi.GetDocumentsByName(context.Background(), cid).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompanyApi.GetDocumentsByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentsByName`: []BTDocumentSummaryInfo
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

[**[]BTDocumentSummaryInfo**](BTDocumentSummaryInfo.md)

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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
    cid := "cid_example" // string | 
    uid := "uid_example" // string | 
    bTCompanyUserParams := *openapiclient.NewBTCompanyUserParams() // BTCompanyUserParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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
**cid** | **string** |  | 
**uid** | **string** |  | 

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

