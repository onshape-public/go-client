# \PropertiesTableTemplateApi

All URIs are relative to *https://cad.onshape.com/api/v6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTableTemplate**](PropertiesTableTemplateApi.md#CreateTableTemplate) | **Post** /tabletemplates | Create a new properties table template.
[**DeleteTableTemplate**](PropertiesTableTemplateApi.md#DeleteTableTemplate) | **Delete** /tabletemplates/{tid} | Delete a properties table template by template ID.
[**GetByCompanyId**](PropertiesTableTemplateApi.md#GetByCompanyId) | **Get** /tabletemplates/companies/{cid} | Retrieve the properties table templates by company ID.
[**GetByDocumentId**](PropertiesTableTemplateApi.md#GetByDocumentId) | **Get** /tabletemplates/d/{did} | Retrieve the properties table templates by document ID.
[**GetTableTemplate**](PropertiesTableTemplateApi.md#GetTableTemplate) | **Get** /tabletemplates/{tid} | Retrieve a properties table template by template ID.



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
    resp, r, err := apiClient.PropertiesTableTemplateApi.CreateTableTemplate(context.Background()).BTPropertiesTableTemplateParams(bTPropertiesTableTemplateParams).TemplateGroupId(templateGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateApi.CreateTableTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTableTemplate`: BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateApi.CreateTableTemplate`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTableTemplate

> map[string]interface{} DeleteTableTemplate(ctx, tid).Execute()

Delete a properties table template by template ID.

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
    resp, r, err := apiClient.PropertiesTableTemplateApi.DeleteTableTemplate(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateApi.DeleteTableTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTableTemplate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateApi.DeleteTableTemplate`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByCompanyId

> []BTPropertiesTableTemplateInfo GetByCompanyId(ctx, cid).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()

Retrieve the properties table templates by company ID.

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
    resp, r, err := apiClient.PropertiesTableTemplateApi.GetByCompanyId(context.Background(), cid).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateApi.GetByCompanyId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByCompanyId`: []BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateApi.GetByCompanyId`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByDocumentId

> []BTPropertiesTableTemplateInfo GetByDocumentId(ctx, did).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()

Retrieve the properties table templates by document ID.

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
    resp, r, err := apiClient.PropertiesTableTemplateApi.GetByDocumentId(context.Background(), did).TemplateType(templateType).OnlyActive(onlyActive).IncludeDefaults(includeDefaults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateApi.GetByDocumentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByDocumentId`: []BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateApi.GetByDocumentId`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTableTemplate

> BTPropertiesTableTemplateInfo GetTableTemplate(ctx, tid).Execute()

Retrieve a properties table template by template ID.

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
    resp, r, err := apiClient.PropertiesTableTemplateApi.GetTableTemplate(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropertiesTableTemplateApi.GetTableTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTableTemplate`: BTPropertiesTableTemplateInfo
    fmt.Fprintf(os.Stdout, "Response from `PropertiesTableTemplateApi.GetTableTemplate`: %v\n", resp)
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

[BasicAuth](../README.md#BasicAuth), [OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

