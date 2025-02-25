# BTAppElementBulkModifyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | Pointer to **string** | The latest change id for the element, after the edit was committed. Deprecated in favor of elementChangeResults. | [optional] 
**DocumentMicroversionId** | Pointer to **string** | The latest change id for the document, after the edit was committed. | [optional] 
**ElementChangeResults** | Pointer to [**[]BTAppElementModifyInfo**](BTAppElementModifyInfo.md) | The results of editing each element affected by the edit. | [optional] 
**ElementId** | Pointer to **string** | The id of the edited element, if a single element was edited. Deprecated in favor of elementChangeResults. | [optional] 
**ElementIds** | Pointer to **[]string** | The ids of the edited elements. Deprecated in favor of elementChangeResults. | [optional] 
**ErrorCode** | Pointer to **int32** | &#x60;0: OK (healthy) | 1: INFO | 2: WARNING | 3: ERROR (dangling or view generation call failed) | 4: UNKNOWN&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 
**ParentChangeId** | Pointer to **string** | The latest change id for the element, before the edit was made. Deprecated in favor of elementChangeResults. | [optional] 
**ParentDocumentMicroversionId** | Pointer to **string** | The latest change id for the document, before the edit was made. | [optional] 
**PropertyEditsMerged** | Pointer to **bool** | Whether the properties of any edited application element were changed after the transaction was created. Deprecated in favor of elementChangeResults. | [optional] 
**TransactionId** | Pointer to **string** | The id of the transaction in which the edit was applied. Deprecated in favor of elementChangeResults. | [optional] 

## Methods

### NewBTAppElementBulkModifyInfo

`func NewBTAppElementBulkModifyInfo() *BTAppElementBulkModifyInfo`

NewBTAppElementBulkModifyInfo instantiates a new BTAppElementBulkModifyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementBulkModifyInfoWithDefaults

`func NewBTAppElementBulkModifyInfoWithDefaults() *BTAppElementBulkModifyInfo`

NewBTAppElementBulkModifyInfoWithDefaults instantiates a new BTAppElementBulkModifyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeId

`func (o *BTAppElementBulkModifyInfo) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *BTAppElementBulkModifyInfo) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *BTAppElementBulkModifyInfo) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *BTAppElementBulkModifyInfo) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetDocumentMicroversionId

`func (o *BTAppElementBulkModifyInfo) GetDocumentMicroversionId() string`

GetDocumentMicroversionId returns the DocumentMicroversionId field if non-nil, zero value otherwise.

### GetDocumentMicroversionIdOk

`func (o *BTAppElementBulkModifyInfo) GetDocumentMicroversionIdOk() (*string, bool)`

GetDocumentMicroversionIdOk returns a tuple with the DocumentMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversionId

`func (o *BTAppElementBulkModifyInfo) SetDocumentMicroversionId(v string)`

SetDocumentMicroversionId sets DocumentMicroversionId field to given value.

### HasDocumentMicroversionId

`func (o *BTAppElementBulkModifyInfo) HasDocumentMicroversionId() bool`

HasDocumentMicroversionId returns a boolean if a field has been set.

### GetElementChangeResults

`func (o *BTAppElementBulkModifyInfo) GetElementChangeResults() []BTAppElementModifyInfo`

GetElementChangeResults returns the ElementChangeResults field if non-nil, zero value otherwise.

### GetElementChangeResultsOk

`func (o *BTAppElementBulkModifyInfo) GetElementChangeResultsOk() (*[]BTAppElementModifyInfo, bool)`

GetElementChangeResultsOk returns a tuple with the ElementChangeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementChangeResults

`func (o *BTAppElementBulkModifyInfo) SetElementChangeResults(v []BTAppElementModifyInfo)`

SetElementChangeResults sets ElementChangeResults field to given value.

### HasElementChangeResults

`func (o *BTAppElementBulkModifyInfo) HasElementChangeResults() bool`

HasElementChangeResults returns a boolean if a field has been set.

### GetElementId

`func (o *BTAppElementBulkModifyInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTAppElementBulkModifyInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTAppElementBulkModifyInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTAppElementBulkModifyInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetElementIds

`func (o *BTAppElementBulkModifyInfo) GetElementIds() []string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *BTAppElementBulkModifyInfo) GetElementIdsOk() (*[]string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *BTAppElementBulkModifyInfo) SetElementIds(v []string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *BTAppElementBulkModifyInfo) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppElementBulkModifyInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementBulkModifyInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementBulkModifyInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementBulkModifyInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementBulkModifyInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementBulkModifyInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementBulkModifyInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementBulkModifyInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementBulkModifyInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementBulkModifyInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementBulkModifyInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementBulkModifyInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppElementBulkModifyInfo) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppElementBulkModifyInfo) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppElementBulkModifyInfo) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppElementBulkModifyInfo) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.

### GetParentDocumentMicroversionId

`func (o *BTAppElementBulkModifyInfo) GetParentDocumentMicroversionId() string`

GetParentDocumentMicroversionId returns the ParentDocumentMicroversionId field if non-nil, zero value otherwise.

### GetParentDocumentMicroversionIdOk

`func (o *BTAppElementBulkModifyInfo) GetParentDocumentMicroversionIdOk() (*string, bool)`

GetParentDocumentMicroversionIdOk returns a tuple with the ParentDocumentMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDocumentMicroversionId

`func (o *BTAppElementBulkModifyInfo) SetParentDocumentMicroversionId(v string)`

SetParentDocumentMicroversionId sets ParentDocumentMicroversionId field to given value.

### HasParentDocumentMicroversionId

`func (o *BTAppElementBulkModifyInfo) HasParentDocumentMicroversionId() bool`

HasParentDocumentMicroversionId returns a boolean if a field has been set.

### GetPropertyEditsMerged

`func (o *BTAppElementBulkModifyInfo) GetPropertyEditsMerged() bool`

GetPropertyEditsMerged returns the PropertyEditsMerged field if non-nil, zero value otherwise.

### GetPropertyEditsMergedOk

`func (o *BTAppElementBulkModifyInfo) GetPropertyEditsMergedOk() (*bool, bool)`

GetPropertyEditsMergedOk returns a tuple with the PropertyEditsMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyEditsMerged

`func (o *BTAppElementBulkModifyInfo) SetPropertyEditsMerged(v bool)`

SetPropertyEditsMerged sets PropertyEditsMerged field to given value.

### HasPropertyEditsMerged

`func (o *BTAppElementBulkModifyInfo) HasPropertyEditsMerged() bool`

HasPropertyEditsMerged returns a boolean if a field has been set.

### GetTransactionId

`func (o *BTAppElementBulkModifyInfo) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BTAppElementBulkModifyInfo) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BTAppElementBulkModifyInfo) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BTAppElementBulkModifyInfo) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


