# \TranslationApi

All URIs are relative to *https://cad.onshape.com/api/v14*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTranslation**](TranslationApi.md#CreateTranslation) | **Post** /translations/d/{did}/w/{wid} | Import or upload a CAD file into Onshape, and translate the data into parts or assemblies.
[**DeleteTranslation**](TranslationApi.md#DeleteTranslation) | **Delete** /translations/{tid} | Delete a translation request.
[**GetAllTranslatorFormats**](TranslationApi.md#GetAllTranslatorFormats) | **Get** /translations/translationformats | Get a list of formats this translation can use.
[**GetDocumentTranslations**](TranslationApi.md#GetDocumentTranslations) | **Get** /translations/d/{did} | Get information on an in-progress or completed translation by document ID.
[**GetTranslation**](TranslationApi.md#GetTranslation) | **Get** /translations/{tid} | Get information on an in-progress or completed translation by translation ID.



## CreateTranslation

> BTTranslationRequestImportInfo CreateTranslation(ctx, did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FlattenAssemblies(flattenAssemblies).FormatName(formatName).ImportAppearances(importAppearances).ImportMaterialDensity(importMaterialDensity).ImportWithinDocument(importWithinDocument).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).MakePublic(makePublic).NotifyUser(notifyUser).OnePartPerDoc(onePartPerDoc).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).RepointAppElementVersionRefs(repointAppElementVersionRefs).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).UseIGESImportPostProcessing(useIGESImportPostProcessing).VersionString(versionString).YAxisIsUp(yAxisIsUp).Execute()

Import or upload a CAD file into Onshape, and translate the data into parts or assemblies.



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
    allowFaultyParts := true // bool |  (optional)
    createComposite := true // bool |  (optional)
    createDrawingIfPossible := true // bool |  (optional)
    encodedFilename := "encodedFilename_example" // string |  (optional)
    extractAssemblyHierarchy := true // bool |  (optional)
    file := os.NewFile(1234, "some_file") // HttpFile | The file to upload. (optional)
    flattenAssemblies := true // bool |  (optional)
    formatName := "formatName_example" // string |  (optional)
    importAppearances := true // bool |  (optional)
    importMaterialDensity := true // bool |  (optional)
    importWithinDocument := true // bool |  (optional)
    joinAdjacentSurfaces := true // bool |  (optional)
    locationElementId := "locationElementId_example" // string |  (optional)
    locationGroupId := "locationGroupId_example" // string |  (optional)
    locationPosition := int32(56) // int32 |  (optional)
    makePublic := true // bool |  (optional)
    notifyUser := true // bool |  (optional)
    onePartPerDoc := true // bool |  (optional)
    ownerId := "ownerId_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)
    projectId := "projectId_example" // string |  (optional)
    repointAppElementVersionRefs := true // bool |  (optional)
    splitAssembliesIntoMultipleDocuments := true // bool |  (optional)
    storeInDocument := true // bool |  (optional)
    translate := true // bool |  (optional)
    unit := "unit_example" // string |  (optional)
    uploadId := "uploadId_example" // string |  (optional)
    useIGESImportPostProcessing := true // bool |  (optional)
    versionString := "versionString_example" // string |  (optional)
    yAxisIsUp := true // bool |  (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TranslationApi.CreateTranslation(context.Background(), did, wid).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).File(file).FlattenAssemblies(flattenAssemblies).FormatName(formatName).ImportAppearances(importAppearances).ImportMaterialDensity(importMaterialDensity).ImportWithinDocument(importWithinDocument).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).MakePublic(makePublic).NotifyUser(notifyUser).OnePartPerDoc(onePartPerDoc).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).RepointAppElementVersionRefs(repointAppElementVersionRefs).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).UseIGESImportPostProcessing(useIGESImportPostProcessing).VersionString(versionString).YAxisIsUp(yAxisIsUp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.CreateTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTranslation`: BTTranslationRequestImportInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.CreateTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **allowFaultyParts** | **bool** |  | 
 **createComposite** | **bool** |  | 
 **createDrawingIfPossible** | **bool** |  | 
 **encodedFilename** | **string** |  | 
 **extractAssemblyHierarchy** | **bool** |  | 
 **file** | **HttpFile** | The file to upload. | 
 **flattenAssemblies** | **bool** |  | 
 **formatName** | **string** |  | 
 **importAppearances** | **bool** |  | 
 **importMaterialDensity** | **bool** |  | 
 **importWithinDocument** | **bool** |  | 
 **joinAdjacentSurfaces** | **bool** |  | 
 **locationElementId** | **string** |  | 
 **locationGroupId** | **string** |  | 
 **locationPosition** | **int32** |  | 
 **makePublic** | **bool** |  | 
 **notifyUser** | **bool** |  | 
 **onePartPerDoc** | **bool** |  | 
 **ownerId** | **string** |  | 
 **parentId** | **string** |  | 
 **projectId** | **string** |  | 
 **repointAppElementVersionRefs** | **bool** |  | 
 **splitAssembliesIntoMultipleDocuments** | **bool** |  | 
 **storeInDocument** | **bool** |  | 
 **translate** | **bool** |  | 
 **unit** | **string** |  | 
 **uploadId** | **string** |  | 
 **useIGESImportPostProcessing** | **bool** |  | 
 **versionString** | **string** |  | 
 **yAxisIsUp** | **bool** |  | 

### Return type

[**BTTranslationRequestImportInfo**](BTTranslationRequestImportInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTranslation

> map[string]interface{} DeleteTranslation(ctx, tid).Execute()

Delete a translation request.

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
    resp, r, err := apiClient.TranslationApi.DeleteTranslation(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.DeleteTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTranslation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.DeleteTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTranslationRequest struct via the builder pattern


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


## GetAllTranslatorFormats

> []BTModelFormatFullInfo GetAllTranslatorFormats(ctx).Execute()

Get a list of formats this translation can use.



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

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TranslationApi.GetAllTranslatorFormats(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.GetAllTranslatorFormats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTranslatorFormats`: []BTModelFormatFullInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.GetAllTranslatorFormats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTranslatorFormatsRequest struct via the builder pattern


### Return type

[**[]BTModelFormatFullInfo**](BTModelFormatFullInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentTranslations

> BTListResponseBTTranslationRequestInfo GetDocumentTranslations(ctx, did).Offset(offset).Limit(limit).Execute()

Get information on an in-progress or completed translation by document ID.

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
    offset := int32(56) // int32 |  (optional) (default to 0)
    limit := int32(56) // int32 |  (optional) (default to 20)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.TranslationApi.GetDocumentTranslations(context.Background(), did).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.GetDocumentTranslations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentTranslations`: BTListResponseBTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.GetDocumentTranslations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentTranslationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** |  | [default to 0]
 **limit** | **int32** |  | [default to 20]

### Return type

[**BTListResponseBTTranslationRequestInfo**](BTListResponseBTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTranslation

> BTTranslationRequestInfo GetTranslation(ctx, tid).Execute()

Get information on an in-progress or completed translation by translation ID.



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
    resp, r, err := apiClient.TranslationApi.GetTranslation(context.Background(), tid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslationApi.GetTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `TranslationApi.GetTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

