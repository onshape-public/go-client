# \ReleasePackageApi

All URIs are relative to *https://cad.onshape.com/api/v15*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObsoletionPackage**](ReleasePackageApi.md#CreateObsoletionPackage) | **Post** /releasepackages/obsoletion/{wfid} | Create an obsoletion candidate to make an existing revision obsolete.
[**CreateReleasePackage**](ReleasePackageApi.md#CreateReleasePackage) | **Post** /releasepackages/release/{wfid} | Create a new release candidate for one or more items.
[**GetReleasePackage**](ReleasePackageApi.md#GetReleasePackage) | **Get** /releasepackages/{rpid} | Get details about the specified release candidate.
[**UpdateReleasePackage**](ReleasePackageApi.md#UpdateReleasePackage) | **Post** /releasepackages/{rpid} | Update the release/obsoletion candidate/item properties.



## CreateObsoletionPackage

> map[string]interface{} CreateObsoletionPackage(ctx, wfid).RevisionId(revisionId).DebugMode(debugMode).Execute()

Create an obsoletion candidate to make an existing revision obsolete.

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
    wfid := "wfid_example" // string | Workflow ID. See [getActiveWorkflows](#/Workflow/getActiveWorkflows).
    revisionId := "revisionId_example" // string | ID of revision to obsolete. See [Revision](https://cad.onshape.com/glassworks/explorer/#/Revision).
    debugMode := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.CreateObsoletionPackage(context.Background(), wfid).RevisionId(revisionId).DebugMode(debugMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.CreateObsoletionPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObsoletionPackage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.CreateObsoletionPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** | Workflow ID. See [getActiveWorkflows](#/Workflow/getActiveWorkflows). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObsoletionPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revisionId** | **string** | ID of revision to obsolete. See [Revision](https://cad.onshape.com/glassworks/explorer/#/Revision). | 
 **debugMode** | **bool** |  | [default to false]

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


## CreateReleasePackage

> map[string]interface{} CreateReleasePackage(ctx, wfid).BTReleasePackageParams(bTReleasePackageParams).DebugMode(debugMode).Execute()

Create a new release candidate for one or more items.



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
    wfid := "wfid_example" // string | Workflow ID. See [getActiveWorkflows](#/Workflow/getActiveWorkflows).
    bTReleasePackageParams := *openapiclient.NewBTReleasePackageParams() // BTReleasePackageParams | 
    debugMode := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.CreateReleasePackage(context.Background(), wfid).BTReleasePackageParams(bTReleasePackageParams).DebugMode(debugMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.CreateReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReleasePackage`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.CreateReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** | Workflow ID. See [getActiveWorkflows](#/Workflow/getActiveWorkflows). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTReleasePackageParams** | [**BTReleasePackageParams**](BTReleasePackageParams.md) |  | 
 **debugMode** | **bool** |  | [default to false]

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleasePackage

> BTReleasePackageInfo GetReleasePackage(ctx, rpid).Detailed(detailed).Execute()

Get details about the specified release candidate.

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
    rpid := "rpid_example" // string | Release candidate package ID, as returned by [createReleasePackage](#/ReleasePackage/createReleasePackage).
    detailed := true // bool |  (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.GetReleasePackage(context.Background(), rpid).Detailed(detailed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.GetReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReleasePackage`: BTReleasePackageInfo
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.GetReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpid** | **string** | Release candidate package ID, as returned by [createReleasePackage](#/ReleasePackage/createReleasePackage). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detailed** | **bool** |  | [default to false]

### Return type

[**BTReleasePackageInfo**](BTReleasePackageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleasePackage

> BTReleasePackageInfo UpdateReleasePackage(ctx, rpid).BTUpdateReleasePackageParams(bTUpdateReleasePackageParams).Action(action).Wfaction(wfaction).Execute()

Update the release/obsoletion candidate/item properties.



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
    rpid := "rpid_example" // string | Release candidate package ID, as returned by [createReleasePackage](#/ReleasePackage/createReleasePackage).
    bTUpdateReleasePackageParams := *openapiclient.NewBTUpdateReleasePackageParams() // BTUpdateReleasePackageParams | 
    action := "action_example" // string | `UPDATE | ADD_ITEMS | REMOVE_ITEMS | SAVE_DRAFT` (optional) (default to "UPDATE")
    wfaction := "wfaction_example" // string | Workflow action to perform on the release candidate. When using the Onshape default release or obsoletion workflows, allowed values are:    `SUBMIT | CREATE_AND_RELEASE | RELEASE | REJECT | OBSOLETE | DISCARD | CREATE_AND_OBSOLETE`    * `DISCARD` can only be performed by the creator of the release candidate and is the only transition that can be performed even if items have errors.    * `CREATE_AND_RELEASE` and `CREATE_AND_OBSOLETE` can only be performed by creator if the [Release management settings](https://cad.onshape.com/help/Content/Plans/release_management_2.htm#rel_candidate_dialog) for the company allow release without approvers. If Release management settings restrict the approver list to a subset of company users, only those users can perform transitions.    When using a custom workflow, allowed values are determined by the workflow definition. See [getActiveWorkflows](#/Workflow/getActiveWorkflows) to get the list of active workflows and their transitions. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.ReleasePackageApi.UpdateReleasePackage(context.Background(), rpid).BTUpdateReleasePackageParams(bTUpdateReleasePackageParams).Action(action).Wfaction(wfaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReleasePackageApi.UpdateReleasePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReleasePackage`: BTReleasePackageInfo
    fmt.Fprintf(os.Stdout, "Response from `ReleasePackageApi.UpdateReleasePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rpid** | **string** | Release candidate package ID, as returned by [createReleasePackage](#/ReleasePackage/createReleasePackage). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleasePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTUpdateReleasePackageParams** | [**BTUpdateReleasePackageParams**](BTUpdateReleasePackageParams.md) |  | 
 **action** | **string** | &#x60;UPDATE | ADD_ITEMS | REMOVE_ITEMS | SAVE_DRAFT&#x60; | [default to &quot;UPDATE&quot;]
 **wfaction** | **string** | Workflow action to perform on the release candidate. When using the Onshape default release or obsoletion workflows, allowed values are:    &#x60;SUBMIT | CREATE_AND_RELEASE | RELEASE | REJECT | OBSOLETE | DISCARD | CREATE_AND_OBSOLETE&#x60;    * &#x60;DISCARD&#x60; can only be performed by the creator of the release candidate and is the only transition that can be performed even if items have errors.    * &#x60;CREATE_AND_RELEASE&#x60; and &#x60;CREATE_AND_OBSOLETE&#x60; can only be performed by creator if the [Release management settings](https://cad.onshape.com/help/Content/Plans/release_management_2.htm#rel_candidate_dialog) for the company allow release without approvers. If Release management settings restrict the approver list to a subset of company users, only those users can perform transitions.    When using a custom workflow, allowed values are determined by the workflow definition. See [getActiveWorkflows](#/Workflow/getActiveWorkflows) to get the list of active workflows and their transitions. | 

### Return type

[**BTReleasePackageInfo**](BTReleasePackageInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

