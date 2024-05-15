# \PropertiesTableTemplateAPI

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTableTemplate**](PropertiesTableTemplateAPI.md#CreateTableTemplate) | **Post** /tabletemplates | Create a new properties table template.
[**DeleteTableTemplate**](PropertiesTableTemplateAPI.md#DeleteTableTemplate) | **Delete** /tabletemplates/{tid} | Delete a properties table template.
[**GetByCompanyId**](PropertiesTableTemplateAPI.md#GetByCompanyId) | **Get** /tabletemplates/companies/{cid} | Get all properties table templates available for a company.
[**GetByDocumentId**](PropertiesTableTemplateAPI.md#GetByDocumentId) | **Get** /tabletemplates/d/{did} | Get all table templates that are available to use on the provided document.
[**GetTableTemplate**](PropertiesTableTemplateAPI.md#GetTableTemplate) | **Get** /tabletemplates/{tid} | Get a properties table template by template ID.



## CreateTableTemplate

> BTPropertiesTableTemplateInfo CreateTableTemplate(ctx).BTPropertiesTableTemplateParams(bTPropertiesTableTemplateParams).TemplateGroupId(templateGroupId).Execute()

Create a new properties table template.

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
    bTPropertiesTableTemplateParams := *openapiclient.NewBTPropertiesTableTemplateParams() // BTPropertiesTableTemplateParams | 
    templateGroupId := "templateGroupId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesTableTemplateAPI.CreateTableTemplate(context.Background()).BTPropertiesTableTemplateParams(bTPropertiesTableTemplateParams).TemplateGroupId(templateGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateAPI.CreateTableTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTableTemplate`: BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateAPI.CreateTableTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bTPropertiesTableTemplateParams** | [**BTPropertiesTableTemplateParams**](BTPropertiesTableTemplateParams.md) |  | 
 **templateGroupId** | **string** |  | 

### Return type

[**BTPropertiesTableTemplateInfo**](BTPropertiesTableTemplateInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTableTemplate

> map[string]interface{} DeleteTableTemplate(ctx, tid).Execute()

Delete a properties table template.

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
    tid := "tid_example" // string | The id of the template in which to perform the operation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesTableTemplateAPI.DeleteTableTemplate(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateAPI.DeleteTableTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTableTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateAPI.DeleteTableTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** | The id of the template in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByCompanyId

> []BTPropertiesTableTemplateInfo GetByCompanyId(ctx, cid).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()

Get all properties table templates available for a company.

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
    cid := "cid_example" // string | The id of the company in which to perform the operation.
    templateType := openapiclient.BTPropertiesTableTemplateType("BOM") // BTPropertiesTableTemplateType | Indicates filter for table templates: 0 (BOM) or 1 (Revision Table). (optional)
    onlyActive := true // bool |  (optional) (default to false)
    includeDefaults := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesTableTemplateAPI.GetByCompanyId(context.Background(), cid).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateAPI.GetByCompanyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByCompanyId`: []BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateAPI.GetByCompanyId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | The id of the company in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByCompanyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateType** | [**BTPropertiesTableTemplateType**](BTPropertiesTableTemplateType.md) | Indicates filter for table templates: 0 (BOM) or 1 (Revision Table). | 
 **onlyActive** | **bool** |  | [default to false]
 **includeDefaults** | **bool** |  | [default to true]

### Return type

[**[]BTPropertiesTableTemplateInfo**](BTPropertiesTableTemplateInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByDocumentId

> []BTPropertiesTableTemplateInfo GetByDocumentId(ctx, did).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()

Get all table templates that are available to use on the provided document.

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
    templateType := openapiclient.BTPropertiesTableTemplateType("BOM") // BTPropertiesTableTemplateType | Indicates filter for table templates: 0 (BOM) or 1 (Revision Table). (optional)
    onlyActive := true // bool |  (optional) (default to true)
    includeDefaults := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesTableTemplateAPI.GetByDocumentId(context.Background(), did).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateAPI.GetByDocumentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByDocumentId`: []BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateAPI.GetByDocumentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByDocumentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateType** | [**BTPropertiesTableTemplateType**](BTPropertiesTableTemplateType.md) | Indicates filter for table templates: 0 (BOM) or 1 (Revision Table). | 
 **onlyActive** | **bool** |  | [default to true]
 **includeDefaults** | **bool** |  | [default to true]

### Return type

[**[]BTPropertiesTableTemplateInfo**](BTPropertiesTableTemplateInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableTemplate

> BTPropertiesTableTemplateInfo GetTableTemplate(ctx, tid).Execute()

Get a properties table template by template ID.

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
    tid := "tid_example" // string | The id of the template in which to perform the operation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropertiesTableTemplateAPI.GetTableTemplate(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateAPI.GetTableTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableTemplate`: BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateAPI.GetTableTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** | The id of the template in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTableTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTPropertiesTableTemplateInfo**](BTPropertiesTableTemplateInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

