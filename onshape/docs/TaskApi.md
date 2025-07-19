# \TaskApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTask**](TaskApi.md#CreateTask) | **Post** /tasks | Create a new task in a draft state.
[**GetActionItems**](TaskApi.md#GetActionItems) | **Get** /tasks | Lists tasks assigned to the specified user
[**GetTask**](TaskApi.md#GetTask) | **Get** /tasks/{tid} | Get a task by id.
[**TransitionTask**](TaskApi.md#TransitionTask) | **Post** /tasks/{tid}/{transition} | Execute a workflow transition.
[**UpdateTask**](TaskApi.md#UpdateTask) | **Post** /tasks/{tid} | Update the task and its properties.



## CreateTask

> BTTaskInfo CreateTask(ctx).BTCreateTaskParams(bTCreateTaskParams).Execute()

Create a new task in a draft state.

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
    bTCreateTaskParams := *openapiclient.NewBTCreateTaskParams("CompanyId_example") // BTCreateTaskParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TaskApi.CreateTask(context.Background()).BTCreateTaskParams(bTCreateTaskParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.CreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTask`: BTTaskInfo
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.CreateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTCreateTaskParams** | [**BTCreateTaskParams**](BTCreateTaskParams.md) |  | 

### Return type

[**BTTaskInfo**](BTTaskInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActionItems

> BTTaskListResponse GetActionItems(ctx).UserId(userId).Status(status).Role(role).Order(order).Type_(type_).DocumentId(documentId).Offset(offset).Limit(limit).Execute()

Lists tasks assigned to the specified user



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
    userId := "userId_example" // string |  (optional)
    status := int32(56) // int32 |  (optional) (default to 2)
    role := int32(56) // int32 |  (optional) (default to 1)
    order := int32(56) // int32 |  (optional) (default to 1)
    type_ := []string{"Inner_example"} // []string |  (optional)
    documentId := "documentId_example" // string |  (optional)
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 50)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TaskApi.GetActionItems(context.Background()).UserId(userId).Status(status).Role(role).Order(order).Type_(type_).DocumentId(documentId).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetActionItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActionItems`: BTTaskListResponse
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetActionItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** |  | 
 **status** | **int32** |  | [default to 2]
 **role** | **int32** |  | [default to 1]
 **order** | **int32** |  | [default to 1]
 **type_** | **[]string** |  | 
 **documentId** | **string** |  | 
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 50]

### Return type

[**BTTaskListResponse**](BTTaskListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> BTTaskInfo GetTask(ctx, tid).Execute()

Get a task by id.

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
    resp, r, err := apiClient.TaskApi.GetTask(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: BTTaskInfo
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTTaskInfo**](BTTaskInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransitionTask

> BTTaskInfo TransitionTask(ctx, tid, transition).Execute()

Execute a workflow transition.

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
    transition := "transition_example" // string | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TaskApi.TransitionTask(context.Background(), tid, transition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.TransitionTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransitionTask`: BTTaskInfo
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.TransitionTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 
**transition** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransitionTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTTaskInfo**](BTTaskInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> BTTaskInfo UpdateTask(ctx, tid).BTUpdateTaskParams(bTUpdateTaskParams).Execute()

Update the task and its properties.

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
    bTUpdateTaskParams := *openapiclient.NewBTUpdateTaskParams() // BTUpdateTaskParams | 

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TaskApi.UpdateTask(context.Background(), tid).BTUpdateTaskParams(bTUpdateTaskParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskApi.UpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTask`: BTTaskInfo
    fmt.Fprintf(os.Stdout, "Response from `TaskApi.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTUpdateTaskParams** | [**BTUpdateTaskParams**](BTUpdateTaskParams.md) |  | 

### Return type

[**BTTaskInfo**](BTTaskInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

