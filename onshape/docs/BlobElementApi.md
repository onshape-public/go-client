# \BlobElementApi

All URIs are relative to *https://cad.onshape.com/api/v10*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlobTranslation**](BlobElementApi.md#CreateBlobTranslation) | **Post** /blobelements/d/{did}/{wv}/{wvid}/e/{eid}/translations | Export a blob element to another format.
[**DownloadFileWorkspace**](BlobElementApi.md#DownloadFileWorkspace) | **Get** /blobelements/d/{did}/w/{wid}/e/{eid} | Download a file from a blob element for the specified workspace/version/microversion.
[**UpdateUnits**](BlobElementApi.md#UpdateUnits) | **Post** /blobelements/d/{did}/w/{wid}/e/{eid}/units | Change the measurement units for the blob element.
[**UploadFileCreateElement**](BlobElementApi.md#UploadFileCreateElement) | **Post** /blobelements/d/{did}/w/{wid} | Upload a file and create a blob element from it.
[**UploadFileUpdateElement**](BlobElementApi.md#UploadFileUpdateElement) | **Post** /blobelements/d/{did}/w/{wid}/e/{eid} | Update a blob element by uploading a file.



## CreateBlobTranslation

> BTTranslationRequestInfo CreateBlobTranslation(ctx, did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).LinkDocumentId(linkDocumentId).Execute()

Export a blob element to another format.



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
    wv := "wv_example" // string | Indicates which of workspace (w) or version (v) id is specified below.
    wvid := "wvid_example" // string | The id of the workspace, version in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    bTTranslateFormatParams := *openapiclient.NewBTTranslateFormatParams("FormatName_example") // BTTranslateFormatParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.BlobElementApi.CreateBlobTranslation(context.Background(), did, wv, wvid, eid).BTTranslateFormatParams(bTTranslateFormatParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.CreateBlobTranslation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlobTranslation`: BTTranslationRequestInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.CreateBlobTranslation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wv** | **string** | Indicates which of workspace (w) or version (v) id is specified below. | 
**wvid** | **string** | The id of the workspace, version in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **bTTranslateFormatParams** | [**BTTranslateFormatParams**](BTTranslateFormatParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTTranslationRequestInfo**](BTTranslationRequestInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFileWorkspace

> HttpFile DownloadFileWorkspace(ctx, did, wid, eid).LinkDocumentId(linkDocumentId).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).Execute()

Download a file from a blob element for the specified workspace/version/microversion.



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
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    contentDisposition := "contentDisposition_example" // string | If \"attachment\", includes a Content-Disposition return header with the filename. (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Entity tag; an md5 checksum of the data in double quotes. If the data to download has the same checksum as this entity tag, a 304 'Not Modified' status will be returned. The entity tag is returned in the response headers as ETag. (optional)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.BlobElementApi.DownloadFileWorkspace(context.Background(), did, wid, eid).LinkDocumentId(linkDocumentId).ContentDisposition(contentDisposition).IfNoneMatch(ifNoneMatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.DownloadFileWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFileWorkspace`: HttpFile
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.DownloadFileWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **contentDisposition** | **string** | If \&quot;attachment\&quot;, includes a Content-Disposition return header with the filename. | 
 **ifNoneMatch** | **string** | Entity tag; an md5 checksum of the data in double quotes. If the data to download has the same checksum as this entity tag, a 304 &#39;Not Modified&#39; status will be returned. The entity tag is returned in the response headers as ETag. | 

### Return type

[**HttpFile**](HttpFile.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUnits

> BTDocumentElementProcessingInfo UpdateUnits(ctx, did, wid, eid).BTUpdateMeshUnitsParams(bTUpdateMeshUnitsParams).LinkDocumentId(linkDocumentId).Execute()

Change the measurement units for the blob element.

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
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    bTUpdateMeshUnitsParams := *openapiclient.NewBTUpdateMeshUnitsParams() // BTUpdateMeshUnitsParams | 
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.BlobElementApi.UpdateUnits(context.Background(), did, wid, eid).BTUpdateMeshUnitsParams(bTUpdateMeshUnitsParams).LinkDocumentId(linkDocumentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.UpdateUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUnits`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.UpdateUnits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTUpdateMeshUnitsParams** | [**BTUpdateMeshUnitsParams**](BTUpdateMeshUnitsParams.md) |  | 
 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileCreateElement

> BTDocumentElementProcessingInfo UploadFileCreateElement(ctx, did, wid).LinkDocumentId(linkDocumentId).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).ImportAppearances(importAppearances).ImportMaterialDensity(importMaterialDensity).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).UseIGESImportPostProcessing(useIGESImportPostProcessing).Execute()

Upload a file and create a blob element from it.



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
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    file := map[string]interface{}{ ... } // map[string]interface{} | The file to upload. (optional)
    allowFaultyParts := true // bool | If true, and a part doesn't pass Onshape validation, it will be imported with faults. (optional)
    createComposite := true // bool | Not supported for importing into a single part studio. (optional)
    createDrawingIfPossible := true // bool |  (optional)
    encodedFilename := "encodedFilename_example" // string | If the filename contains non-ASCII characters. Use this field to store the filename. (optional)
    extractAssemblyHierarchy := true // bool |  (optional)
    flattenAssemblies := true // bool | If the file is an assembly, or contains an assembly, setting this to True will import it as a Part Studio. In this case the assembly will be flattened to a set of parts in a Part Studio. There will be duplicate parts created whenever a part is instanced more than once. If False, it will be imported as an Assembly. (optional)
    formatName := "formatName_example" // string |  (optional)
    joinAdjacentSurfaces := true // bool |  (optional)
    locationElementId := "locationElementId_example" // string |  (optional)
    locationGroupId := "locationGroupId_example" // string |  (optional)
    locationPosition := int32(56) // int32 |  (optional) (default to -1)
    notifyUser := true // bool |  (optional) (default to true)
    ownerId := "ownerId_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)
    projectId := "projectId_example" // string |  (optional)
    public := true // bool |  (optional)
    onePartPerDoc := true // bool |  (optional) (default to false)
    splitAssembliesIntoMultipleDocuments := true // bool |  (optional) (default to false)
    storeInDocument := true // bool |  (optional) (default to true)
    translate := true // bool |  (optional) (default to true)
    unit := "unit_example" // string |  (optional) (default to "")
    uploadId := "uploadId_example" // string |  (optional)
    versionString := "versionString_example" // string |  (optional)
    importAppearances := true // bool | Face appearances defined on models will be imported. (optional) (default to true)
    importMaterialDensity := true // bool | Material density defined on models will be imported. (optional) (default to false)
    yAxisIsUp := true // bool | If the file was created in a system that orients with Y Axis Up, the models would by default be brought into Onshape (a Z Axis Up system) with a flipped coordinate system. Toggle this value to reorient the axis system to match Onshape and display the model with the coordinates you expect. (optional)
    importWithinDocument := true // bool |  (optional)
    useIGESImportPostProcessing := true // bool | Try getting optimized topology from IGES model. (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.BlobElementApi.UploadFileCreateElement(context.Background(), did, wid).LinkDocumentId(linkDocumentId).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).ImportAppearances(importAppearances).ImportMaterialDensity(importMaterialDensity).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).UseIGESImportPostProcessing(useIGESImportPostProcessing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.UploadFileCreateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileCreateElement`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.UploadFileCreateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileCreateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **file** | [**map[string]interface{}**](map[string]interface{}.md) | The file to upload. | 
 **allowFaultyParts** | **bool** | If true, and a part doesn&#39;t pass Onshape validation, it will be imported with faults. | 
 **createComposite** | **bool** | Not supported for importing into a single part studio. | 
 **createDrawingIfPossible** | **bool** |  | 
 **encodedFilename** | **string** | If the filename contains non-ASCII characters. Use this field to store the filename. | 
 **extractAssemblyHierarchy** | **bool** |  | 
 **flattenAssemblies** | **bool** | If the file is an assembly, or contains an assembly, setting this to True will import it as a Part Studio. In this case the assembly will be flattened to a set of parts in a Part Studio. There will be duplicate parts created whenever a part is instanced more than once. If False, it will be imported as an Assembly. | 
 **formatName** | **string** |  | 
 **joinAdjacentSurfaces** | **bool** |  | 
 **locationElementId** | **string** |  | 
 **locationGroupId** | **string** |  | 
 **locationPosition** | **int32** |  | [default to -1]
 **notifyUser** | **bool** |  | [default to true]
 **ownerId** | **string** |  | 
 **parentId** | **string** |  | 
 **projectId** | **string** |  | 
 **public** | **bool** |  | 
 **onePartPerDoc** | **bool** |  | [default to false]
 **splitAssembliesIntoMultipleDocuments** | **bool** |  | [default to false]
 **storeInDocument** | **bool** |  | [default to true]
 **translate** | **bool** |  | [default to true]
 **unit** | **string** |  | [default to &quot;&quot;]
 **uploadId** | **string** |  | 
 **versionString** | **string** |  | 
 **importAppearances** | **bool** | Face appearances defined on models will be imported. | [default to true]
 **importMaterialDensity** | **bool** | Material density defined on models will be imported. | [default to false]
 **yAxisIsUp** | **bool** | If the file was created in a system that orients with Y Axis Up, the models would by default be brought into Onshape (a Z Axis Up system) with a flipped coordinate system. Toggle this value to reorient the axis system to match Onshape and display the model with the coordinates you expect. | 
 **importWithinDocument** | **bool** |  | 
 **useIGESImportPostProcessing** | **bool** | Try getting optimized topology from IGES model. | [default to false]

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFileUpdateElement

> BTDocumentElementProcessingInfo UploadFileUpdateElement(ctx, did, wid, eid).LinkDocumentId(linkDocumentId).ParentChangeId(parentChangeId).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).ImportAppearances(importAppearances).ImportMaterialDensity(importMaterialDensity).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).UseIGESImportPostProcessing(useIGESImportPostProcessing).Execute()

Update a blob element by uploading a file.



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
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    parentChangeId := "parentChangeId_example" // string | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. (optional)
    file := map[string]interface{}{ ... } // map[string]interface{} | The file to upload. (optional)
    allowFaultyParts := true // bool | If true, and a part doesn't pass Onshape validation, it will be imported with faults. (optional)
    createComposite := true // bool | Not supported for importing into a single part studio. (optional)
    createDrawingIfPossible := true // bool |  (optional)
    encodedFilename := "encodedFilename_example" // string | If the filename contains non-ASCII characters. Use this field to store the filename. (optional)
    extractAssemblyHierarchy := true // bool |  (optional)
    flattenAssemblies := true // bool | If the file is an assembly, or contains an assembly, setting this to True will import it as a Part Studio. In this case the assembly will be flattened to a set of parts in a Part Studio. There will be duplicate parts created whenever a part is instanced more than once. If False, it will be imported as an Assembly. (optional)
    formatName := "formatName_example" // string |  (optional)
    joinAdjacentSurfaces := true // bool |  (optional)
    locationElementId := "locationElementId_example" // string |  (optional)
    locationGroupId := "locationGroupId_example" // string |  (optional)
    locationPosition := int32(56) // int32 |  (optional) (default to -1)
    notifyUser := true // bool |  (optional) (default to true)
    ownerId := "ownerId_example" // string |  (optional)
    parentId := "parentId_example" // string |  (optional)
    projectId := "projectId_example" // string |  (optional)
    public := true // bool |  (optional)
    onePartPerDoc := true // bool |  (optional) (default to false)
    splitAssembliesIntoMultipleDocuments := true // bool |  (optional) (default to false)
    storeInDocument := true // bool |  (optional) (default to true)
    translate := true // bool |  (optional) (default to true)
    unit := "unit_example" // string |  (optional) (default to "")
    uploadId := "uploadId_example" // string |  (optional)
    versionString := "versionString_example" // string |  (optional)
    importAppearances := true // bool | Face appearances defined on models will be imported. (optional) (default to true)
    importMaterialDensity := true // bool | Material density defined on models will be imported. (optional) (default to false)
    yAxisIsUp := true // bool | If the file was created in a system that orients with Y Axis Up, the models would by default be brought into Onshape (a Z Axis Up system) with a flipped coordinate system. Toggle this value to reorient the axis system to match Onshape and display the model with the coordinates you expect. (optional)
    importWithinDocument := true // bool |  (optional)
    useIGESImportPostProcessing := true // bool | Try getting optimized topology from IGES model. (optional) (default to false)

    apiConfiguration := openapiclient.NewAPIConfiguration()
    apiClient := openapiclient.NewAPIClient(apiConfiguration)
    resp, r, err := apiClient.BlobElementApi.UploadFileUpdateElement(context.Background(), did, wid, eid).LinkDocumentId(linkDocumentId).ParentChangeId(parentChangeId).File(file).AllowFaultyParts(allowFaultyParts).CreateComposite(createComposite).CreateDrawingIfPossible(createDrawingIfPossible).EncodedFilename(encodedFilename).ExtractAssemblyHierarchy(extractAssemblyHierarchy).FlattenAssemblies(flattenAssemblies).FormatName(formatName).JoinAdjacentSurfaces(joinAdjacentSurfaces).LocationElementId(locationElementId).LocationGroupId(locationGroupId).LocationPosition(locationPosition).NotifyUser(notifyUser).OwnerId(ownerId).ParentId(parentId).ProjectId(projectId).Public(public).OnePartPerDoc(onePartPerDoc).SplitAssembliesIntoMultipleDocuments(splitAssembliesIntoMultipleDocuments).StoreInDocument(storeInDocument).Translate(translate).Unit(unit).UploadId(uploadId).VersionString(versionString).ImportAppearances(importAppearances).ImportMaterialDensity(importMaterialDensity).YAxisIsUp(yAxisIsUp).ImportWithinDocument(importWithinDocument).UseIGESImportPostProcessing(useIGESImportPostProcessing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlobElementApi.UploadFileUpdateElement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFileUpdateElement`: BTDocumentElementProcessingInfo
    fmt.Fprintf(os.Stdout, "Response from `BlobElementApi.UploadFileUpdateElement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wid** | **string** | The id of the workspace in which to perform the operation. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileUpdateElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **parentChangeId** | **string** | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. | 
 **file** | [**map[string]interface{}**](map[string]interface{}.md) | The file to upload. | 
 **allowFaultyParts** | **bool** | If true, and a part doesn&#39;t pass Onshape validation, it will be imported with faults. | 
 **createComposite** | **bool** | Not supported for importing into a single part studio. | 
 **createDrawingIfPossible** | **bool** |  | 
 **encodedFilename** | **string** | If the filename contains non-ASCII characters. Use this field to store the filename. | 
 **extractAssemblyHierarchy** | **bool** |  | 
 **flattenAssemblies** | **bool** | If the file is an assembly, or contains an assembly, setting this to True will import it as a Part Studio. In this case the assembly will be flattened to a set of parts in a Part Studio. There will be duplicate parts created whenever a part is instanced more than once. If False, it will be imported as an Assembly. | 
 **formatName** | **string** |  | 
 **joinAdjacentSurfaces** | **bool** |  | 
 **locationElementId** | **string** |  | 
 **locationGroupId** | **string** |  | 
 **locationPosition** | **int32** |  | [default to -1]
 **notifyUser** | **bool** |  | [default to true]
 **ownerId** | **string** |  | 
 **parentId** | **string** |  | 
 **projectId** | **string** |  | 
 **public** | **bool** |  | 
 **onePartPerDoc** | **bool** |  | [default to false]
 **splitAssembliesIntoMultipleDocuments** | **bool** |  | [default to false]
 **storeInDocument** | **bool** |  | [default to true]
 **translate** | **bool** |  | [default to true]
 **unit** | **string** |  | [default to &quot;&quot;]
 **uploadId** | **string** |  | 
 **versionString** | **string** |  | 
 **importAppearances** | **bool** | Face appearances defined on models will be imported. | [default to true]
 **importMaterialDensity** | **bool** | Material density defined on models will be imported. | [default to false]
 **yAxisIsUp** | **bool** | If the file was created in a system that orients with Y Axis Up, the models would by default be brought into Onshape (a Z Axis Up system) with a flipped coordinate system. Toggle this value to reorient the axis system to match Onshape and display the model with the coordinates you expect. | 
 **importWithinDocument** | **bool** |  | 
 **useIGESImportPostProcessing** | **bool** | Try getting optimized topology from IGES model. | [default to false]

### Return type

[**BTDocumentElementProcessingInfo**](BTDocumentElementProcessingInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

