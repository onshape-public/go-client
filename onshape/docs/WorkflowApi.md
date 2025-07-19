# \WorkflowApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnumerateObjectWorkflows**](WorkflowApi.md#EnumerateObjectWorkflows) | **Get** /workflow/companies/{cid}/objects | Enumerate all of a company&#39;s workflowable objects.
[**GetActiveWorkflows**](WorkflowApi.md#GetActiveWorkflows) | **Get** /workflow/active | Get all active workflows for the currently logged in user&#39;s company.
[**GetAllowedApprovers**](WorkflowApi.md#GetAllowedApprovers) | **Get** /workflow/c/{companyId}/approvers | Get all identities allowed to be approvers on a workflow object.
[**GetAuditLog**](WorkflowApi.md#GetAuditLog) | **Get** /workflow/obj/{objectId}/auditlog | Get all audit log entries for a workflowable object.
[**GetWorkflowById**](WorkflowApi.md#GetWorkflowById) | **Get** /workflow/obj/{objectId} | Lightweight information about the current state of a workflowable object like release package.



## EnumerateObjectWorkflows

> BTListResponseBTObjectWorkflowInfo EnumerateObjectWorkflows(ctx, cid).ObjectTypes(objectTypes).States(states).Limit(limit).ModifiedAfter(modifiedAfter).Execute()

Enumerate all of a company's workflowable objects.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    cid := "cid_example" // string | The company or enterprise ID that owns the resource.
    objectTypes := []openapiclient.BTAPIWorkflowableType{openapiclient.BTAPIWorkflowableType("RELEASE")} // []BTAPIWorkflowableType | Optionally filter for specific workflowable types. Defaults to RELEASE and OBSOLETION (optional)
    states := []string{"Inner_example"} // []string | Optionally filter for specific workflow states like PENDING, RELEASED (optional)
    limit := int32(56) // int32 | The number of list entries to return in a single API call. (optional) (default to 20)
    modifiedAfter := time.Now() // JSONTime | The earliest modification date of workflowable object to find. (optional) (default to "2000-01-01T00:00Z")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.WorkflowApi.EnumerateObjectWorkflows(context.Background(), cid).ObjectTypes(objectTypes).States(states).Limit(limit).ModifiedAfter(modifiedAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.EnumerateObjectWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnumerateObjectWorkflows`: BTListResponseBTObjectWorkflowInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.EnumerateObjectWorkflows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | The company or enterprise ID that owns the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnumerateObjectWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectTypes** | [**[]BTAPIWorkflowableType**](BTAPIWorkflowableType.md) | Optionally filter for specific workflowable types. Defaults to RELEASE and OBSOLETION | 
 **states** | **[]string** | Optionally filter for specific workflow states like PENDING, RELEASED | 
 **limit** | **int32** | The number of list entries to return in a single API call. | [default to 20]
 **modifiedAfter** | **JSONTime** | The earliest modification date of workflowable object to find. | [default to &quot;2000-01-01T00:00Z&quot;]

### Return type

[**BTListResponseBTObjectWorkflowInfo**](BTListResponseBTObjectWorkflowInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveWorkflows

> BTActiveWorkflowInfo GetActiveWorkflows(ctx).DocumentId(documentId).Execute()

Get all active workflows for the currently logged in user's company.



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
    documentId := "documentId_example" // string |  (optional) (default to "")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.WorkflowApi.GetActiveWorkflows(context.Background()).DocumentId(documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.GetActiveWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveWorkflows`: BTActiveWorkflowInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.GetActiveWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **documentId** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTActiveWorkflowInfo**](BTActiveWorkflowInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowedApprovers

> BTListResponseBTWorkflowObserverOptionInfo GetAllowedApprovers(ctx, companyId).Q(q).ExpandTeams(expandTeams).IncludeSelf(includeSelf).Execute()

Get all identities allowed to be approvers on a workflow object.



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
    companyId := "companyId_example" // string | 
    q := "q_example" // string |  (optional) (default to "")
    expandTeams := true // bool |  (optional) (default to true)
    includeSelf := true // bool |  (optional) (default to true)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.WorkflowApi.GetAllowedApprovers(context.Background(), companyId).Q(q).ExpandTeams(expandTeams).IncludeSelf(includeSelf).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.GetAllowedApprovers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllowedApprovers`: BTListResponseBTWorkflowObserverOptionInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.GetAllowedApprovers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**companyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowedApproversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  | [default to &quot;&quot;]
 **expandTeams** | **bool** |  | [default to true]
 **includeSelf** | **bool** |  | [default to true]

### Return type

[**BTListResponseBTWorkflowObserverOptionInfo**](BTListResponseBTWorkflowObserverOptionInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLog

> BTWorkflowAuditLogInfo GetAuditLog(ctx, objectId).Execute()

Get all audit log entries for a workflowable object.

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
    objectId := "objectId_example" // string | Id of a workflowable object like release package, task etc.

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.WorkflowApi.GetAuditLog(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.GetAuditLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuditLog`: BTWorkflowAuditLogInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.GetAuditLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | Id of a workflowable object like release package, task etc. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowAuditLogInfo**](BTWorkflowAuditLogInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowById

> BTObjectWorkflowInfo GetWorkflowById(ctx, objectId).Execute()

Lightweight information about the current state of a workflowable object like release package.



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
    objectId := "objectId_example" // string | Id of a workflowable object like release package, task etc.

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.WorkflowApi.GetWorkflowById(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowApi.GetWorkflowById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowById`: BTObjectWorkflowInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowApi.GetWorkflowById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | Id of a workflowable object like release package, task etc. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTObjectWorkflowInfo**](BTObjectWorkflowInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

