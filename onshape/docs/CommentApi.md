# \CommentApi

All URIs are relative to *https://staging.dev.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAttachment**](CommentApi.md#AddAttachment) | **Post** /comments/{cid}/attachment | Update a user’s comment by comment ID.
[**CreateComment**](CommentApi.md#CreateComment) | **Post** /comments | Update a document with a new comment.
[**DeleteAttachments**](CommentApi.md#DeleteAttachments) | **Delete** /comments/{cid}/attachment | Delete an attachment from a comment by comment ID.
[**DeleteComment**](CommentApi.md#DeleteComment) | **Delete** /comments/{cid} | Delete a comment by comment ID.
[**GetAttachment**](CommentApi.md#GetAttachment) | **Get** /comments/{cid}/attachment/{fdid}.{ext} | Retrieve an attachment associated with a comment by comment ID and file document ID (and extension).
[**GetComment**](CommentApi.md#GetComment) | **Get** /comments/{cid} | Retrieve details for a comment by comment ID.
[**GetComments**](CommentApi.md#GetComments) | **Get** /comments | Retrieve a list of comments for a document.
[**Reopen**](CommentApi.md#Reopen) | **Post** /comments/{cid}/reopen | Reopen a resolved comment by comment ID.
[**Resolve**](CommentApi.md#Resolve) | **Post** /comments/{cid}/resolve | Resolve a comment by comment ID.
[**UpdateComment**](CommentApi.md#UpdateComment) | **Post** /comments/{cid} | Update a user’s comment by comment ID.



## AddAttachment

> BTCommentInfo AddAttachment(ctx, cid).File(file).Execute()

Update a user’s comment by comment ID.

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
    file := map[string]interface{}{ ... } // map[string]interface{} | The file to upload.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentApi.AddAttachment(context.Background(), cid).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.AddAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAttachment`: BTCommentInfo
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.AddAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | [**map[string]interface{}**](map[string]interface{}.md) | The file to upload. | 

### Return type

[**BTCommentInfo**](BTCommentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateComment

> BTCommentInfo CreateComment(ctx).BTCommentParams(bTCommentParams).Execute()

Update a document with a new comment.

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
    bTCommentParams := *openapiclient.NewBTCommentParams() // BTCommentParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentApi.CreateComment(context.Background()).BTCommentParams(bTCommentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.CreateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateComment`: BTCommentInfo
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.CreateComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTCommentParams** | [**BTCommentParams**](BTCommentParams.md) |  | 

### Return type

[**BTCommentInfo**](BTCommentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttachments

> map[string]interface{} DeleteAttachments(ctx, cid).Execute()

Delete an attachment from a comment by comment ID.

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
    resp, r, err := apiClient.CommentApi.DeleteAttachments(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.DeleteAttachments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAttachments`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.DeleteAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComment

> map[string]interface{} DeleteComment(ctx, cid).Execute()

Delete a comment by comment ID.

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
    resp, r, err := apiClient.CommentApi.DeleteComment(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.DeleteComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteComment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.DeleteComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachment

> map[string]interface{} GetAttachment(ctx, cid, fdid, ext).Execute()

Retrieve an attachment associated with a comment by comment ID and file document ID (and extension).

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
    fdid := "fdid_example" // string | 
    ext := "ext_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentApi.GetAttachment(context.Background(), cid, fdid, ext).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.GetAttachment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttachment`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 
**fdid** | **string** |  | 
**ext** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/_*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComment

> BTCommentInfo GetComment(ctx, cid).Execute()

Retrieve details for a comment by comment ID.

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
    resp, r, err := apiClient.CommentApi.GetComment(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.GetComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComment`: BTCommentInfo
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTCommentInfo**](BTCommentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComments

> BTListResponseBTCommentInfo GetComments(ctx).Did(did).ObjectType(objectType).Pid(pid).Eid(eid).Filter(filter).Resolved(resolved).Offset(offset).Limit(limit).Execute()

Retrieve a list of comments for a document.

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
    did := "did_example" // string |  (optional) (default to "")
    objectType := int32(56) // int32 |  (optional) (default to 6)
    pid := "pid_example" // string |  (optional) (default to "")
    eid := "eid_example" // string |  (optional) (default to "")
    filter := int32(56) // int32 |  (optional) (default to 0)
    resolved := true // bool |  (optional) (default to true)
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentApi.GetComments(context.Background()).Did(did).ObjectType(objectType).Pid(pid).Eid(eid).Filter(filter).Resolved(resolved).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.GetComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComments`: BTListResponseBTCommentInfo
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.GetComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **did** | **string** |  | [default to &quot;&quot;]
 **objectType** | **int32** |  | [default to 6]
 **pid** | **string** |  | [default to &quot;&quot;]
 **eid** | **string** |  | [default to &quot;&quot;]
 **filter** | **int32** |  | [default to 0]
 **resolved** | **bool** |  | [default to true]
 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTCommentInfo**](BTListResponseBTCommentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reopen

> BTCommentInfo Reopen(ctx, cid).Execute()

Reopen a resolved comment by comment ID.

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
    resp, r, err := apiClient.CommentApi.Reopen(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.Reopen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Reopen`: BTCommentInfo
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.Reopen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReopenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTCommentInfo**](BTCommentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Resolve

> BTCommentInfo Resolve(ctx, cid).Execute()

Resolve a comment by comment ID.

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
    resp, r, err := apiClient.CommentApi.Resolve(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.Resolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Resolve`: BTCommentInfo
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.Resolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTCommentInfo**](BTCommentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComment

> BTCommentInfo UpdateComment(ctx, cid).BTCommentParams(bTCommentParams).Execute()

Update a user’s comment by comment ID.

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
    bTCommentParams := *openapiclient.NewBTCommentParams() // BTCommentParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CommentApi.UpdateComment(context.Background(), cid).BTCommentParams(bTCommentParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommentApi.UpdateComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateComment`: BTCommentInfo
    fmt.Fprintf(os.Stdout, "Response from `CommentApi.UpdateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bTCommentParams** | [**BTCommentParams**](BTCommentParams.md) |  | 

### Return type

[**BTCommentInfo**](BTCommentInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

