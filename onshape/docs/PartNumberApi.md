# \PartNumberApi

All URIs are relative to *https://cad.onshape.com/api/v12*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateNextNumbers**](PartNumberApi.md#UpdateNextNumbers) | **Post** /partnumber/nextnumbers | Send the items to generate numbers for, and return the next valid available part numbers.



## UpdateNextNumbers

> map[string][]BTNextPartNumber UpdateNextNumbers(ctx).BTNextPartNumbersParam(bTNextPartNumbersParam).Cid(cid).Did(did).Execute()

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
    cid := "cid_example" // string | The ID of the company or enterprise that owns this resource and/or is making the request to Onshape. (optional)
    did := "did_example" // string | Document ID. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.PartNumberApi.UpdateNextNumbers(context.Background()).BTNextPartNumbersParam(bTNextPartNumbersParam).Cid(cid).Did(did).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartNumberApi.UpdateNextNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNextNumbers`: map[string][]BTNextPartNumber
    fmt.Fprintf(os.Stdout, "Response from `PartNumberApi.UpdateNextNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNextNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTNextPartNumbersParam** | [**BTNextPartNumbersParam**](BTNextPartNumbersParam.md) |  | 
 **cid** | **string** | The ID of the company or enterprise that owns this resource and/or is making the request to Onshape. | 
 **did** | **string** | Document ID. | 

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

