# \AppAssociativeDataApi

All URIs are relative to *https://cad.onshape.com/api/v9*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyAssociativeData**](AppAssociativeDataApi.md#CopyAssociativeData) | **Post** /appelements/d/{did}/w/{wid}/e/{eid}/copyassociativedata | Copy associative data from one view to another.
[**DeleteAssociativeData**](AppAssociativeDataApi.md#DeleteAssociativeData) | **Delete** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | Delete the associative data from the specified app element.
[**GetAssociativeData**](AppAssociativeDataApi.md#GetAssociativeData) | **Get** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | Get the associative data for the specified app element.
[**PostAssociativeData**](AppAssociativeDataApi.md#PostAssociativeData) | **Post** /appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata | Set the associative data for the specified app element.



## CopyAssociativeData

> BTAppAssociativeDataArrayInfo CopyAssociativeData(ctx, did, wid, eid).BTAppElementParamsArrayBTCopyViewAssociativeDataParams(bTAppElementParamsArrayBTCopyViewAssociativeDataParams).Execute()

Copy associative data from one view to another.



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
    eid := "eid_example" // string | 
    bTAppElementParamsArrayBTCopyViewAssociativeDataParams := *openapiclient.NewBTAppElementParamsArrayBTCopyViewAssociativeDataParams() // BTAppElementParamsArrayBTCopyViewAssociativeDataParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAssociativeDataApi.CopyAssociativeData(context.Background(), did, wid, eid).BTAppElementParamsArrayBTCopyViewAssociativeDataParams(bTAppElementParamsArrayBTCopyViewAssociativeDataParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAssociativeDataApi.CopyAssociativeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyAssociativeData`: BTAppAssociativeDataArrayInfo
    fmt.Fprintf(os.Stdout, "Response from `AppAssociativeDataApi.CopyAssociativeData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**wid** | **string** |  | 
**eid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bTAppElementParamsArrayBTCopyViewAssociativeDataParams** | [**BTAppElementParamsArrayBTCopyViewAssociativeDataParams**](BTAppElementParamsArrayBTCopyViewAssociativeDataParams.md) |  | 

### Return type

[**BTAppAssociativeDataArrayInfo**](BTAppAssociativeDataArrayInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAssociativeData

> BTAppElementBasicInfo DeleteAssociativeData(ctx, did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).AssociativeDataId(associativeDataId).ExternalDocumentId(externalDocumentId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).ReferenceId(referenceId).Execute()

Delete the associative data from the specified app element.



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
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    transactionId := "transactionId_example" // string |  (optional) (default to "")
    parentChangeId := "parentChangeId_example" // string |  (optional) (default to "")
    associativeDataId := []string{"Inner_example"} // []string |  (optional)
    externalDocumentId := "externalDocumentId_example" // string |  (optional) (default to "")
    elementId := "elementId_example" // string |  (optional) (default to "")
    viewId := "viewId_example" // string |  (optional) (default to "")
    microversionId := "microversionId_example" // string |  (optional) (default to "")
    documentMicroversion := "documentMicroversion_example" // string |  (optional) (default to "")
    deterministicId := "deterministicId_example" // string |  (optional) (default to "")
    featureId := "featureId_example" // string |  (optional) (default to "")
    entityId := "entityId_example" // string |  (optional) (default to "")
    occurrenceId := "occurrenceId_example" // string |  (optional) (default to "")
    referenceId := "referenceId_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAssociativeDataApi.DeleteAssociativeData(context.Background(), did, eid, wvm, wvmid).TransactionId(transactionId).ParentChangeId(parentChangeId).AssociativeDataId(associativeDataId).ExternalDocumentId(externalDocumentId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).ReferenceId(referenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAssociativeDataApi.DeleteAssociativeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAssociativeData`: BTAppElementBasicInfo
    fmt.Fprintf(os.Stdout, "Response from `AppAssociativeDataApi.DeleteAssociativeData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **transactionId** | **string** |  | [default to &quot;&quot;]
 **parentChangeId** | **string** |  | [default to &quot;&quot;]
 **associativeDataId** | **[]string** |  | 
 **externalDocumentId** | **string** |  | [default to &quot;&quot;]
 **elementId** | **string** |  | [default to &quot;&quot;]
 **viewId** | **string** |  | [default to &quot;&quot;]
 **microversionId** | **string** |  | [default to &quot;&quot;]
 **documentMicroversion** | **string** |  | [default to &quot;&quot;]
 **deterministicId** | **string** |  | [default to &quot;&quot;]
 **featureId** | **string** |  | [default to &quot;&quot;]
 **entityId** | **string** |  | [default to &quot;&quot;]
 **occurrenceId** | **string** |  | [default to &quot;&quot;]
 **referenceId** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppElementBasicInfo**](BTAppElementBasicInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssociativeData

> BTAppAssociativeDataArrayInfo GetAssociativeData(ctx, did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).AssociativeDataId(associativeDataId).ExternalDocumentId(externalDocumentId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).ReturnIdTags(returnIdTags).ReferenceId(referenceId).Execute()

Get the associative data for the specified app element.



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
    wvm := "wvm_example" // string | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below.
    wvmid := "wvmid_example" // string | The id of the workspace, version or document microversion in which the operation should be performed.
    eid := "eid_example" // string | The id of the element in which to perform the operation.
    linkDocumentId := "linkDocumentId_example" // string | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. (optional) (default to "")
    transactionId := "transactionId_example" // string |  (optional) (default to "")
    changeId := "changeId_example" // string |  (optional) (default to "")
    associativeDataId := []string{"Inner_example"} // []string |  (optional)
    externalDocumentId := "externalDocumentId_example" // string |  (optional) (default to "")
    elementId := "elementId_example" // string |  (optional) (default to "")
    viewId := "viewId_example" // string |  (optional) (default to "")
    microversionId := "microversionId_example" // string |  (optional) (default to "")
    documentMicroversion := "documentMicroversion_example" // string |  (optional) (default to "")
    deterministicId := "deterministicId_example" // string |  (optional) (default to "")
    featureId := "featureId_example" // string |  (optional) (default to "")
    entityId := "entityId_example" // string |  (optional) (default to "")
    occurrenceId := "occurrenceId_example" // string |  (optional) (default to "")
    returnIdTags := true // bool |  (optional) (default to false)
    referenceId := "referenceId_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAssociativeDataApi.GetAssociativeData(context.Background(), did, wvm, wvmid, eid).LinkDocumentId(linkDocumentId).TransactionId(transactionId).ChangeId(changeId).AssociativeDataId(associativeDataId).ExternalDocumentId(externalDocumentId).ElementId(elementId).ViewId(viewId).MicroversionId(microversionId).DocumentMicroversion(documentMicroversion).DeterministicId(deterministicId).FeatureId(featureId).EntityId(entityId).OccurrenceId(occurrenceId).ReturnIdTags(returnIdTags).ReferenceId(referenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAssociativeDataApi.GetAssociativeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssociativeData`: BTAppAssociativeDataArrayInfo
    fmt.Fprintf(os.Stdout, "Response from `AppAssociativeDataApi.GetAssociativeData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** | The id of the document in which to perform the operation. | 
**wvm** | **string** | Indicates which of workspace (w), version (v), or document microversion (m) id is specified below. | 
**wvmid** | **string** | The id of the workspace, version or document microversion in which the operation should be performed. | 
**eid** | **string** | The id of the element in which to perform the operation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **linkDocumentId** | **string** | The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both. | [default to &quot;&quot;]
 **transactionId** | **string** |  | [default to &quot;&quot;]
 **changeId** | **string** |  | [default to &quot;&quot;]
 **associativeDataId** | **[]string** |  | 
 **externalDocumentId** | **string** |  | [default to &quot;&quot;]
 **elementId** | **string** |  | [default to &quot;&quot;]
 **viewId** | **string** |  | [default to &quot;&quot;]
 **microversionId** | **string** |  | [default to &quot;&quot;]
 **documentMicroversion** | **string** |  | [default to &quot;&quot;]
 **deterministicId** | **string** |  | [default to &quot;&quot;]
 **featureId** | **string** |  | [default to &quot;&quot;]
 **entityId** | **string** |  | [default to &quot;&quot;]
 **occurrenceId** | **string** |  | [default to &quot;&quot;]
 **returnIdTags** | **bool** |  | [default to false]
 **referenceId** | **string** |  | [default to &quot;&quot;]

### Return type

[**BTAppAssociativeDataArrayInfo**](BTAppAssociativeDataArrayInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAssociativeData

> BTAppAssociativeDataArrayInfo PostAssociativeData(ctx, did, eid, wvm, wvmid).Body(body).Execute()

Set the associative data for the specified app element.



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
    eid := "eid_example" // string | 
    wvm := "wvm_example" // string | 
    wvmid := "wvmid_example" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAssociativeDataApi.PostAssociativeData(context.Background(), did, eid, wvm, wvmid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAssociativeDataApi.PostAssociativeData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAssociativeData`: BTAppAssociativeDataArrayInfo
    fmt.Fprintf(os.Stdout, "Response from `AppAssociativeDataApi.PostAssociativeData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**did** | **string** |  | 
**eid** | **string** |  | 
**wvm** | **string** |  | 
**wvmid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAssociativeDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**BTAppAssociativeDataArrayInfo**](BTAppAssociativeDataArrayInfo.md)

### Authorization

[OAuth2](../README.md#OAuth2), [BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8; qs=0.09
- **Accept**: application/json;charset=UTF-8; qs=0.09

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

