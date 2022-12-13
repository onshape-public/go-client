# \VariablesApi

All URIs are relative to *https://cad.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariableStudio**](VariablesApi.md#CreateVariableStudio) | **Post** /variables/d/{did}/w/{wid}/variablestudio | Create a variable studio
[**GetVariableStudioReferences**](VariablesApi.md#GetVariableStudioReferences) | **Get** /variables/d/{did}/{wv}/{wvid}/e/{eid}/variablestudioreferences | Retrieve the variable studio references from an element
[**GetVariableStudioScope**](VariablesApi.md#GetVariableStudioScope) | **Get** /variables/d/{did}/{wv}/{wvid}/e/{eid}/variablestudioscope | Get the scope of a variable studio
[**GetVariables**](VariablesApi.md#GetVariables) | **Get** /variables/d/{did}/{wv}/{wvid}/e/{eid}/variables | Retrieve the variables from a variable table
[**SetVariableStudioReferences**](VariablesApi.md#SetVariableStudioReferences) | **Post** /variables/d/{did}/w/{wid}/e/{eid}/variablestudioreferences | Set the variable studio references for an element
[**SetVariableStudioScope**](VariablesApi.md#SetVariableStudioScope) | **Post** /variables/d/{did}/w/{wid}/e/{eid}/variablestudioscope | Set the scope of a variable studio
[**SetVariables**](VariablesApi.md#SetVariables) | **Post** /variables/d/{did}/w/{wid}/e/{eid}/variables | Assign variables to a variable studio



## CreateVariableStudio

> BTDocumentElementInfo CreateVariableStudio(ctx, did, wid).BTModelElementParams(bTModelElementParams).LinkDocumentId(linkDocumentId).Execute()

Create a variable studio



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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    bTModelElementParams := *openapiclient.NewBTModelElementParams() // BTModelElementParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesApi.CreateVariableStudio(context.Background(), did, wid).BTModelElementParams(bTModelElementParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.CreateVariableStudio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariableStudio`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.CreateVariableStudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableStudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTModelElementParams** | [**BTModelElementParams**](BTModelElementParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableStudioReferences

> BTVariableStudioReferenceListInfo GetVariableStudioReferences(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()

Retrieve the variable studio references from an element



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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesApi.GetVariableStudioReferences(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.GetVariableStudioReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariableStudioReferences`: BTVariableStudioReferenceListInfo
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.GetVariableStudioReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableStudioReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTVariableStudioReferenceListInfo**](BTVariableStudioReferenceListInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariableStudioScope

> BTVariableStudioScopeInfo GetVariableStudioScope(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()

Get the scope of a variable studio



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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesApi.GetVariableStudioScope(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.GetVariableStudioScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariableStudioScope`: BTVariableStudioScopeInfo
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.GetVariableStudioScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableStudioScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTVariableStudioScopeInfo**](BTVariableStudioScopeInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariables

> BTVariableTableInfo GetVariables(ctx, did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).IncludeValuesAndReferencedVariables(includeValuesAndReferencedVariables).Execute()

Retrieve the variables from a variable table



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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wv := "wv_example" // string | 
    wvid := "wvid_example" // string | 
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    configuration := "configuration_example" // string |  (optional) (default to "")
    includeValuesAndReferencedVariables := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesApi.GetVariables(context.Background(), did, wv, wvid, eid).LinkDocumentId(linkDocumentId).Configuration(configuration).IncludeValuesAndReferencedVariables(includeValuesAndReferencedVariables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.GetVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariables`: BTVariableTableInfo
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.GetVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** |  | 
**wvid** | **string** |  | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **configuration** | **string** |  | [default to &quot;&quot;]
 **includeValuesAndReferencedVariables** | **bool** |  | [default to false]

### Return type

[**BTVariableTableInfo**](BTVariableTableInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVariableStudioReferences

> map[string]interface{} SetVariableStudioReferences(ctx, did, wid, eid).BTVariableStudioReferenceListInfo(bTVariableStudioReferenceListInfo).LinkDocumentId(linkDocumentId).Execute()

Set the variable studio references for an element



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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    eid := "eid_example" // string | 
    bTVariableStudioReferenceListInfo := *openapiclient.NewBTVariableStudioReferenceListInfo() // BTVariableStudioReferenceListInfo | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesApi.SetVariableStudioReferences(context.Background(), did, wid, eid).BTVariableStudioReferenceListInfo(bTVariableStudioReferenceListInfo).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.SetVariableStudioReferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetVariableStudioReferences`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.SetVariableStudioReferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVariableStudioReferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTVariableStudioReferenceListInfo** | [**BTVariableStudioReferenceListInfo**](BTVariableStudioReferenceListInfo.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVariableStudioScope

> map[string]interface{} SetVariableStudioScope(ctx, did, wid, eid).BTVariableStudioScopeInfo(bTVariableStudioScopeInfo).LinkDocumentId(linkDocumentId).Execute()

Set the scope of a variable studio



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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    eid := "eid_example" // string | 
    bTVariableStudioScopeInfo := *openapiclient.NewBTVariableStudioScopeInfo(false) // BTVariableStudioScopeInfo | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesApi.SetVariableStudioScope(context.Background(), did, wid, eid).BTVariableStudioScopeInfo(bTVariableStudioScopeInfo).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.SetVariableStudioScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetVariableStudioScope`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.SetVariableStudioScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVariableStudioScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTVariableStudioScopeInfo** | [**BTVariableStudioScopeInfo**](BTVariableStudioScopeInfo.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetVariables

> map[string]interface{} SetVariables(ctx, did, wid, eid).BTVariableParams(bTVariableParams).LinkDocumentId(linkDocumentId).Execute()

Assign variables to a variable studio



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
    did := "did_example" // string | The id of the document in which to perform the operation.
    wid := "wid_example" // string | The id of the workspace in which to perform the operation.
    eid := "eid_example" // string | 
    bTVariableParams := []openapiclient.BTVariableParams{*openapiclient.NewBTVariableParams("Expression_example", "Name_example", "Type_example")} // []BTVariableParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VariablesApi.SetVariables(context.Background(), did, wid, eid).BTVariableParams(bTVariableParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VariablesApi.SetVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetVariables`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VariablesApi.SetVariables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTVariableParams** | [**[]BTVariableParams**](BTVariableParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

