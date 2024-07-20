# \NumberingSchemeApi

All URIs are relative to *https://cad.onshape.com/api/v8*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NextNumbers**](NumberingSchemeApi.md#NextNumbers) | **Post** /numberingscheme/nextnumbers | Send the items to generate numbers for, and return the next valid available part numbers.



## NextNumbers

> map[string][]BTNextPartNumber NextNumbers(ctx).BTNextPartNumbersParam(bTNextPartNumbersParam).Cid(cid).Did(did).Execute()

Send the items to generate numbers for, and return the next valid available part numbers.

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
    bTNextPartNumbersParam := *openapiclient.NewBTNextPartNumbersParam() // BTNextPartNumbersParam | 
    cid := "cid_example" // string |  (optional)
    did := "did_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NumberingSchemeApi.NextNumbers(context.Background()).BTNextPartNumbersParam(bTNextPartNumbersParam).Cid(cid).Did(did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NumberingSchemeApi.NextNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NextNumbers`: map[string][]BTNextPartNumber
    fmt.Fprintf(os.Stdout, "Response from `NumberingSchemeApi.NextNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNextNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTNextPartNumbersParam** | [**BTNextPartNumbersParam**](BTNextPartNumbersParam.md) |  | 
 **cid** | **string** |  | 
 **did** | **string** |  | 

### Return type

[**map[string][]BTNextPartNumber**](array.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

