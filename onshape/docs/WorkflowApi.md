# \WorkflowApi

All URIs are relative to *https://cad.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActiveWorkflows**](WorkflowApi.md#GetActiveWorkflows) | **Get** /workflow/active | Retrieve active workflow.
[**GetAllowedApprovers**](WorkflowApi.md#GetAllowedApprovers) | **Get** /workflow/c/{companyId}/approvers | 
[**GetAuditLog**](WorkflowApi.md#GetAuditLog) | **Get** /workflow/obj/{objectId}/auditlog | 



## GetActiveWorkflows

> BTActiveWorkflowInfo GetActiveWorkflows(ctx).DocumentId(documentId).Execute()

Retrieve active workflow.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowedApprovers

> BTListResponseBTWorkflowObserverOptionInfo GetAllowedApprovers(ctx, companyId).Q(q).ExpandTeams(expandTeams).IncludeSelf(includeSelf).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLog

> BTWorkflowAuditLogInfo GetAuditLog(ctx, objectId).Execute()



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
    objectId := "objectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
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
**objectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowAuditLogInfo**](BTWorkflowAuditLogInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

