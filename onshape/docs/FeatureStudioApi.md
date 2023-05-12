# \FeatureStudioApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFeatureStudio**](FeatureStudioApi.md#CreateFeatureStudio) | **Post** /featurestudios/d/{did}/w/{wid} | Create Feature Studio by document ID and workspace ID.
[**GetFeatureStudioContents**](FeatureStudioApi.md#GetFeatureStudioContents) | **Get** /featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid} | 
[**GetFeatureStudioSpecs**](FeatureStudioApi.md#GetFeatureStudioSpecs) | **Get** /featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid}/featurespecs | Retrieve Feature Studio specs by document ID, workspace or version or microversion ID, and tab ID.
[**UpdateFeatureStudioContents**](FeatureStudioApi.md#UpdateFeatureStudioContents) | **Post** /featurestudios/d/{did}/{wvm}/{wvmid}/e/{eid} | Update Feature Studio contents by document ID, workspace or version or microversion ID, and tab ID.



## CreateFeatureStudio

> BTDocumentElementInfo CreateFeatureStudio(ctx, did, wid).BTModelElementParams(bTModelElementParams).Execute()

Create Feature Studio by document ID and workspace ID.

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
    did := "did_example" // string | 
    wid := "wid_example" // string | 
    bTModelElementParams := *openapiclient.NewBTModelElementParams() // BTModelElementParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureStudioApi.CreateFeatureStudio(context.Background(), did, wid).BTModelElementParams(bTModelElementParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudioApi.CreateFeatureStudio``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFeatureStudio`: BTDocumentElementInfo
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudioApi.CreateFeatureStudio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeatureStudioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bTModelElementParams** | [**BTModelElementParams**](BTModelElementParams.md) |  | 

### Return type

[**BTDocumentElementInfo**](BTDocumentElementInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureStudioContents

> BTFeatureStudioContents2239 GetFeatureStudioContents(ctx, did, wvm, wvmid, eid).Execute()



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
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureStudioApi.GetFeatureStudioContents(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudioApi.GetFeatureStudioContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureStudioContents`: BTFeatureStudioContents2239
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudioApi.GetFeatureStudioContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureStudioContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTFeatureStudioContents2239**](BTFeatureStudioContents2239.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeatureStudioSpecs

> BTFeatureSpecsResponse664 GetFeatureStudioSpecs(ctx, did, wvm, wvmid, eid).Execute()

Retrieve Feature Studio specs by document ID, workspace or version or microversion ID, and tab ID.

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
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureStudioApi.GetFeatureStudioSpecs(context.Background(), did, wvm, wvmid, eid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudioApi.GetFeatureStudioSpecs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFeatureStudioSpecs`: BTFeatureSpecsResponse664
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudioApi.GetFeatureStudioSpecs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureStudioSpecsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**BTFeatureSpecsResponse664**](BTFeatureSpecsResponse664.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFeatureStudioContents

> BTFeatureStudioContents2239 UpdateFeatureStudioContents(ctx, did, wvm, wvmid, eid).BTFeatureStudioContents2239(bTFeatureStudioContents2239).Execute()

Update Feature Studio contents by document ID, workspace or version or microversion ID, and tab ID.

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
    did := "did_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    eid := "eid_example" // string | 
    bTFeatureStudioContents2239 := *openapiclient.NewBTFeatureStudioContents2239() // BTFeatureStudioContents2239 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FeatureStudioApi.UpdateFeatureStudioContents(context.Background(), did, wvm, wvmid, eid).BTFeatureStudioContents2239(bTFeatureStudioContents2239).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FeatureStudioApi.UpdateFeatureStudioContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFeatureStudioContents`: BTFeatureStudioContents2239
    fmt.Fprintf(os.Stdout, "Response from `FeatureStudioApi.UpdateFeatureStudioContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFeatureStudioContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTFeatureStudioContents2239** | [**BTFeatureStudioContents2239**](BTFeatureStudioContents2239.md) |  | 

### Return type

[**BTFeatureStudioContents2239**](BTFeatureStudioContents2239.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

