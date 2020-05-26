# BTDocumentMergeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LibraryVersionMismatch** | Pointer to **bool** |  | [optional] 
**OverwrittenElements** | Pointer to [**[]BTDocumentElementInfo**](BTDocumentElementInfo.md) |  | [optional] 

## Methods

### NewBTDocumentMergeInfo

`func NewBTDocumentMergeInfo() *BTDocumentMergeInfo`

NewBTDocumentMergeInfo instantiates a new BTDocumentMergeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentMergeInfoWithDefaults

`func NewBTDocumentMergeInfoWithDefaults() *BTDocumentMergeInfo`

NewBTDocumentMergeInfoWithDefaults instantiates a new BTDocumentMergeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLibraryVersionMismatch

`func (o *BTDocumentMergeInfo) GetLibraryVersionMismatch() bool`

GetLibraryVersionMismatch returns the LibraryVersionMismatch field if non-nil, zero value otherwise.

### GetLibraryVersionMismatchOk

`func (o *BTDocumentMergeInfo) GetLibraryVersionMismatchOk() (*bool, bool)`

GetLibraryVersionMismatchOk returns a tuple with the LibraryVersionMismatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersionMismatch

`func (o *BTDocumentMergeInfo) SetLibraryVersionMismatch(v bool)`

SetLibraryVersionMismatch sets LibraryVersionMismatch field to given value.

### HasLibraryVersionMismatch

`func (o *BTDocumentMergeInfo) HasLibraryVersionMismatch() bool`

HasLibraryVersionMismatch returns a boolean if a field has been set.

### GetOverwrittenElements

`func (o *BTDocumentMergeInfo) GetOverwrittenElements() []BTDocumentElementInfo`

GetOverwrittenElements returns the OverwrittenElements field if non-nil, zero value otherwise.

### GetOverwrittenElementsOk

`func (o *BTDocumentMergeInfo) GetOverwrittenElementsOk() (*[]BTDocumentElementInfo, bool)`

GetOverwrittenElementsOk returns a tuple with the OverwrittenElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrittenElements

`func (o *BTDocumentMergeInfo) SetOverwrittenElements(v []BTDocumentElementInfo)`

SetOverwrittenElements sets OverwrittenElements field to given value.

### HasOverwrittenElements

`func (o *BTDocumentMergeInfo) HasOverwrittenElements() bool`

HasOverwrittenElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


