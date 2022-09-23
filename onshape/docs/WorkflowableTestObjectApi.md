# \WorkflowableTestObjectApi

All URIs are relative to *https://staging.dev.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowableTestObject**](WorkflowableTestObjectApi.md#CreateWorkflowableTestObject) | **Post** /workflowabletestobject/testobject/{wfid} | Update workflowable test object by workflow ID.
[**GetWorkflowableTestObject**](WorkflowableTestObjectApi.md#GetWorkflowableTestObject) | **Get** /workflowabletestobject/{oid} | Retrieve workflowable test object by object ID.
[**TransitionWorkflowableTestObject**](WorkflowableTestObjectApi.md#TransitionWorkflowableTestObject) | **Post** /workflowabletestobject/{oid}/{transition} | Update workflowable test object transition by object ID.
[**UpdateWorkflowableTestObject**](WorkflowableTestObjectApi.md#UpdateWorkflowableTestObject) | **Post** /workflowabletestobject/{oid} | Update workflowable test object by object ID.



## CreateWorkflowableTestObject

> BTWorkflowableTestObjectInfo CreateWorkflowableTestObject(ctx, wfid).Execute()

Update workflowable test object by workflow ID.

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
    wfid := "wfid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowableTestObjectApi.CreateWorkflowableTestObject(context.Background(), wfid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowableTestObjectApi.CreateWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowableTestObjectApi.CreateWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowableTestObject

> BTWorkflowableTestObjectInfo GetWorkflowableTestObject(ctx, oid).Execute()

Retrieve workflowable test object by object ID.

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
    oid := "oid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowableTestObjectApi.GetWorkflowableTestObject(context.Background(), oid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowableTestObjectApi.GetWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowableTestObjectApi.GetWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionWorkflowableTestObject

> BTWorkflowableTestObjectInfo TransitionWorkflowableTestObject(ctx, oid, transition).Execute()

Update workflowable test object transition by object ID.

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
    oid := "oid_example" // string | 
    transition := "transition_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowableTestObjectApi.TransitionWorkflowableTestObject(context.Background(), oid, transition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowableTestObjectApi.TransitionWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransitionWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowableTestObjectApi.TransitionWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 
**transition** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowableTestObject

> BTWorkflowableTestObjectInfo UpdateWorkflowableTestObject(ctx, oid).BTUpdateWorkflowableTestObjectParams(bTUpdateWorkflowableTestObjectParams).Execute()

Update workflowable test object by object ID.

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
    oid := "oid_example" // string | 
    bTUpdateWorkflowableTestObjectParams := *openapiclient.NewBTUpdateWorkflowableTestObjectParams() // BTUpdateWorkflowableTestObjectParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowableTestObjectApi.UpdateWorkflowableTestObject(context.Background(), oid).BTUpdateWorkflowableTestObjectParams(bTUpdateWorkflowableTestObjectParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowableTestObjectApi.UpdateWorkflowableTestObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowableTestObject`: BTWorkflowableTestObjectInfo
    fmt.Fprintf(os.Stdout, "Response from `WorkflowableTestObjectApi.UpdateWorkflowableTestObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowableTestObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTUpdateWorkflowableTestObjectParams** | [**BTUpdateWorkflowableTestObjectParams**](BTUpdateWorkflowableTestObjectParams.md) |  | 

### Return type

[**BTWorkflowableTestObjectInfo**](BTWorkflowableTestObjectInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

