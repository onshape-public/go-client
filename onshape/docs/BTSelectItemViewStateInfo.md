# BTSelectItemViewStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSelectorId** | Pointer to **string** |  | [optional] 
**DocumentSelectors** | Pointer to [**[]BTDocumentSelectorInfo**](BTDocumentSelectorInfo.md) |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 

## Methods

### NewBTSelectItemViewStateInfo

`func NewBTSelectItemViewStateInfo() *BTSelectItemViewStateInfo`

NewBTSelectItemViewStateInfo instantiates a new BTSelectItemViewStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSelectItemViewStateInfoWithDefaults

`func NewBTSelectItemViewStateInfoWithDefaults() *BTSelectItemViewStateInfo`

NewBTSelectItemViewStateInfoWithDefaults instantiates a new BTSelectItemViewStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSelectorId

`func (o *BTSelectItemViewStateInfo) GetActiveSelectorId() string`

GetActiveSelectorId returns the ActiveSelectorId field if non-nil, zero value otherwise.

### GetActiveSelectorIdOk

`func (o *BTSelectItemViewStateInfo) GetActiveSelectorIdOk() (*string, bool)`

GetActiveSelectorIdOk returns a tuple with the ActiveSelectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSelectorId

`func (o *BTSelectItemViewStateInfo) SetActiveSelectorId(v string)`

SetActiveSelectorId sets ActiveSelectorId field to given value.

### HasActiveSelectorId

`func (o *BTSelectItemViewStateInfo) HasActiveSelectorId() bool`

HasActiveSelectorId returns a boolean if a field has been set.

### GetDocumentSelectors

`func (o *BTSelectItemViewStateInfo) GetDocumentSelectors() []BTDocumentSelectorInfo`

GetDocumentSelectors returns the DocumentSelectors field if non-nil, zero value otherwise.

### GetDocumentSelectorsOk

`func (o *BTSelectItemViewStateInfo) GetDocumentSelectorsOk() (*[]BTDocumentSelectorInfo, bool)`

GetDocumentSelectorsOk returns a tuple with the DocumentSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSelectors

`func (o *BTSelectItemViewStateInfo) SetDocumentSelectors(v []BTDocumentSelectorInfo)`

SetDocumentSelectors sets DocumentSelectors field to given value.

### HasDocumentSelectors

`func (o *BTSelectItemViewStateInfo) HasDocumentSelectors() bool`

HasDocumentSelectors returns a boolean if a field has been set.

### GetPurpose

`func (o *BTSelectItemViewStateInfo) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *BTSelectItemViewStateInfo) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *BTSelectItemViewStateInfo) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *BTSelectItemViewStateInfo) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


