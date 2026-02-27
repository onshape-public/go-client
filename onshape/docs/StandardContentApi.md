# \StandardContentApi

All URIs are relative to *https://cad.onshape.com/api/v14*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParameterValuesForId**](StandardContentApi.md#GetParameterValuesForId) | **Get** /standardcontent/d/{did}/parametervalues | Gets all possible values for each of the standard content parameters.
[**GetStandardContentList**](StandardContentApi.md#GetStandardContentList) | **Get** /standardcontent/list | List all standard content.
[**SetCustomParameters**](StandardContentApi.md#SetCustomParameters) | **Post** /standardcontent/d/{did}/customparameters | Sets the part number and description for a standard content component.



## GetParameterValuesForId

> []BTStandardContentParameterDescriptor GetParameterValuesForId(ctx, did).Execute()

Gets all possible values for each of the standard content parameters.

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
    did := "did_example" // string | Document ID.

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.StandardContentApi.GetParameterValuesForId(context.Background(), did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandardContentApi.GetParameterValuesForId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetParameterValuesForId`: []BTStandardContentParameterDescriptor
    fmt.Fprintf(os.Stdout, "Response from `StandardContentApi.GetParameterValuesForId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParameterValuesForIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BTStandardContentParameterDescriptor**](BTStandardContentParameterDescriptor.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStandardContentList

> []BTStandardContentHierarchyItem GetStandardContentList(ctx).Standard(standard).Execute()

List all standard content.

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
    standard := "standard_example" // string | If provided, list only standard content corresponding to this standard. (optional) (default to "")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.StandardContentApi.GetStandardContentList(context.Background()).Standard(standard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandardContentApi.GetStandardContentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStandardContentList`: []BTStandardContentHierarchyItem
    fmt.Fprintf(os.Stdout, "Response from `StandardContentApi.GetStandardContentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStandardContentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standard** | **string** | If provided, list only standard content corresponding to this standard. | [default to &quot;&quot;]

### Return type

[**[]BTStandardContentHierarchyItem**](BTStandardContentHierarchyItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCustomParameters

> BTStandardContentSetCustomParametersBatchResponse SetCustomParameters(ctx, did).CompanyId(companyId).BTStandardContentSetCustomParametersBatchRequest(bTStandardContentSetCustomParametersBatchRequest).Execute()

Sets the part number and description for a standard content component.

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
    did := "did_example" // string | Document ID.
    companyId := "companyId_example" // string | The id of the company that owns the metadata to be modified. (optional) (default to "")
    bTStandardContentSetCustomParametersBatchRequest := *openapiclient.NewBTStandardContentSetCustomParametersBatchRequest() // BTStandardContentSetCustomParametersBatchRequest |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.StandardContentApi.SetCustomParameters(context.Background(), did).CompanyId(companyId).BTStandardContentSetCustomParametersBatchRequest(bTStandardContentSetCustomParametersBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandardContentApi.SetCustomParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCustomParameters`: BTStandardContentSetCustomParametersBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `StandardContentApi.SetCustomParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | Document ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCustomParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyId** | **string** | The id of the company that owns the metadata to be modified. | [default to &quot;&quot;]
 **bTStandardContentSetCustomParametersBatchRequest** | [**BTStandardContentSetCustomParametersBatchRequest**](BTStandardContentSetCustomParametersBatchRequest.md) |  | 

### Return type

[**BTStandardContentSetCustomParametersBatchResponse**](BTStandardContentSetCustomParametersBatchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

