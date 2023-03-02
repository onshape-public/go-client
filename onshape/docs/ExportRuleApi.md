# \ExportRuleApi

All URIs are relative to *https://cad.onshape.com/api/v5*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidRuleOptions**](ExportRuleApi.md#GetValidRuleOptions) | **Get** /exportrules/options/{cu}/{cuid} | Retrieve a list of the valid export rule options by object type and owner ID.



## GetValidRuleOptions

> BTExportRuleValidOptionsInfo GetValidRuleOptions(ctx, cu, cuid).Execute()

Retrieve a list of the valid export rule options by object type and owner ID.

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
    cu := "cu_example" // string | Indicates which of company (c) or user (u) id is specified below.
    cuid := "cuid_example" // string | The id of the company or user in which the operation should be performed.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExportRuleApi.GetValidRuleOptions(context.Background(), cu, cuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportRuleApi.GetValidRuleOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetValidRuleOptions`: BTExportRuleValidOptionsInfo
    fmt.Fprintf(os.Stdout, "Response from `ExportRuleApi.GetValidRuleOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cu** | **string** | Indicates which of company (c) or user (u) id is specified below. | 
**cuid** | **string** | The id of the company or user in which the operation should be performed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetValidRuleOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BTExportRuleValidOptionsInfo**](BTExportRuleValidOptionsInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

