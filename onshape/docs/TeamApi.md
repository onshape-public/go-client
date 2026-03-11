# \TeamApi

All URIs are relative to *https://cad.onshape.com/api/v14*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Find**](TeamApi.md#Find) | **Get** /teams | Get a list of all teams the current user belongs to.
[**GetMembers**](TeamApi.md#GetMembers) | **Get** /teams/{tid}/members | Get a list of a team&#39;s members.
[**GetTeam**](TeamApi.md#GetTeam) | **Get** /teams/{tid} | Get team information by team ID.



## Find

> BTGlobalTreeNodeListResponseBTTeamInfo Find(ctx).Query(query).Filter(filter).Uid(uid).CompanyId(companyId).SortColumn(sortColumn).SortOrder(sortOrder).IncludeCompanyOwnedTeams(includeCompanyOwnedTeams).Offset(offset).Limit(limit).Execute()

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
    query := "query_example" // string |  (optional) (default to "")
    filter := int32(56) // int32 |  (optional) (default to 0)
    uid := "uid_example" // string |  (optional)
    companyId := "companyId_example" // string |  (optional)
    sortColumn := "sortColumn_example" // string |  (optional)
    sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
    includeCompanyOwnedTeams := true // bool |  (optional) (default to true)
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 100)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TeamApi.Find(context.Background()).Query(query).Filter(filter).Uid(uid).CompanyId(companyId).SortColumn(sortColumn).SortOrder(sortOrder).IncludeCompanyOwnedTeams(includeCompanyOwnedTeams).Offset(offset).Limit(limit).Execute()
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
 **query** | **string** |  | [default to &quot;&quot;]
 **filter** | **int32** |  | [default to 0]
 **uid** | **string** |  | 
 **companyId** | **string** |  | 
 **sortColumn** | **string** |  | 
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **includeCompanyOwnedTeams** | **bool** |  | [default to true]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 100]

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

> BTListResponseBTTeamMemberInfo GetMembers(ctx, tid).SortColumn(sortColumn).SortOrder(sortOrder).Q(q).Offset(offset).Limit(limit).Execute()

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
    q := "q_example" // string |  (optional) (default to "")
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TeamApi.GetMembers(context.Background(), tid).SortColumn(sortColumn).SortOrder(sortOrder).Q(q).Offset(offset).Limit(limit).Execute()
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
 **q** | **string** |  | [default to &quot;&quot;]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
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

