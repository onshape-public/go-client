# \TeamApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Find**](TeamApi.md#Find) | **Get** /teams | Get a list of all teams the current user belongs to.
[**GetMembers**](TeamApi.md#GetMembers) | **Get** /teams/{tid}/members | Get a list of a team&#39;s members.
[**GetTeam**](TeamApi.md#GetTeam) | **Get** /teams/{tid} | Get team information by team ID.



## Find

> BTGlobalTreeNodeListResponseBTTeamInfo Find(ctx).Prefix(prefix).Uid(uid).CompanyId(companyId).IncludeCompanyOwnedTeams(includeCompanyOwnedTeams).Execute()

Get a list of all teams the current user belongs to.

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
    prefix := "prefix_example" // string |  (optional) (default to "")
    uid := "uid_example" // string |  (optional)
    companyId := "companyId_example" // string |  (optional)
    includeCompanyOwnedTeams := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.Find(context.Background()).Prefix(prefix).Uid(uid).CompanyId(companyId).IncludeCompanyOwnedTeams(includeCompanyOwnedTeams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.Find``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Find`: BTGlobalTreeNodeListResponseBTTeamInfo
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.Find`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string** |  | [default to &quot;&quot;]
 **uid** | **string** |  | 
 **companyId** | **string** |  | 
 **includeCompanyOwnedTeams** | **bool** |  | [default to true]

### Return type

[**BTGlobalTreeNodeListResponseBTTeamInfo**](BTGlobalTreeNodeListResponseBTTeamInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMembers

> BTListResponseBTTeamMemberInfo GetMembers(ctx, tid).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Q(q).Execute()

Get a list of a team's members.



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
    tid := "tid_example" // string | 
    sortColumn := "sortColumn_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)
    q := "q_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.GetMembers(context.Background(), tid).SortColumn(sortColumn).SortOrder(sortOrder).Offset(offset).Limit(limit).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.GetMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMembers`: BTListResponseBTTeamMemberInfo
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.GetMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortColumn** | **string** |  | 
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]
 **q** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTListResponseBTTeamMemberInfo**](BTListResponseBTTeamMemberInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> BTTeamInfo GetTeam(ctx, tid).Execute()

Get team information by team ID.

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
    tid := "tid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.GetTeam(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.GetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeam`: BTTeamInfo
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.GetTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTTeamInfo**](BTTeamInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

