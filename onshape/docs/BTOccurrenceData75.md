# BTOccurrenceData75

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**FeatureData** | Pointer to [**map[string]BTFeatureOccurrenceData775**](BTFeatureOccurrenceData775.md) |  | [optional] 
**ForceHighestQualityTessellation** | Pointer to **bool** |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**IsFixed** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**LockInfo** | Pointer to [**BTLockedSubAssembly4590**](BTLockedSubAssembly4590.md) |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Occurrence** | Pointer to [**BTOccurrence74**](BTOccurrence74.md) |  | [optional] 
**Transform** | Pointer to [**BTBSMatrix386**](BTBSMatrix386.md) |  | [optional] 

## Methods

### NewBTOccurrenceData75

`func NewBTOccurrenceData75() *BTOccurrenceData75`

NewBTOccurrenceData75 instantiates a new BTOccurrenceData75 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTOccurrenceData75WithDefaults

`func NewBTOccurrenceData75WithDefaults() *BTOccurrenceData75`

NewBTOccurrenceData75WithDefaults instantiates a new BTOccurrenceData75 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTOccurrenceData75) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTOccurrenceData75) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTOccurrenceData75) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTOccurrenceData75) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetFeatureData

`func (o *BTOccurrenceData75) GetFeatureData() map[string]BTFeatureOccurrenceData775`

GetFeatureData returns the FeatureData field if non-nil, zero value otherwise.

### GetFeatureDataOk

`func (o *BTOccurrenceData75) GetFeatureDataOk() (*map[string]BTFeatureOccurrenceData775, bool)`

GetFeatureDataOk returns a tuple with the FeatureData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureData

`func (o *BTOccurrenceData75) SetFeatureData(v map[string]BTFeatureOccurrenceData775)`

SetFeatureData sets FeatureData field to given value.

### HasFeatureData

`func (o *BTOccurrenceData75) HasFeatureData() bool`

HasFeatureData returns a boolean if a field has been set.

### GetForceHighestQualityTessellation

`func (o *BTOccurrenceData75) GetForceHighestQualityTessellation() bool`

GetForceHighestQualityTessellation returns the ForceHighestQualityTessellation field if non-nil, zero value otherwise.

### GetForceHighestQualityTessellationOk

`func (o *BTOccurrenceData75) GetForceHighestQualityTessellationOk() (*bool, bool)`

GetForceHighestQualityTessellationOk returns a tuple with the ForceHighestQualityTessellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceHighestQualityTessellation

`func (o *BTOccurrenceData75) SetForceHighestQualityTessellation(v bool)`

SetForceHighestQualityTessellation sets ForceHighestQualityTessellation field to given value.

### HasForceHighestQualityTessellation

`func (o *BTOccurrenceData75) HasForceHighestQualityTessellation() bool`

HasForceHighestQualityTessellation returns a boolean if a field has been set.

### GetHidden

`func (o *BTOccurrenceData75) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *BTOccurrenceData75) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *BTOccurrenceData75) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *BTOccurrenceData75) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetIsFixed

`func (o *BTOccurrenceData75) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *BTOccurrenceData75) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *BTOccurrenceData75) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.

### HasIsFixed

`func (o *BTOccurrenceData75) HasIsFixed() bool`

HasIsFixed returns a boolean if a field has been set.

### GetIsHidden

`func (o *BTOccurrenceData75) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *BTOccurrenceData75) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *BTOccurrenceData75) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *BTOccurrenceData75) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetLockInfo

`func (o *BTOccurrenceData75) GetLockInfo() BTLockedSubAssembly4590`

GetLockInfo returns the LockInfo field if non-nil, zero value otherwise.

### GetLockInfoOk

`func (o *BTOccurrenceData75) GetLockInfoOk() (*BTLockedSubAssembly4590, bool)`

GetLockInfoOk returns a tuple with the LockInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockInfo

`func (o *BTOccurrenceData75) SetLockInfo(v BTLockedSubAssembly4590)`

SetLockInfo sets LockInfo field to given value.

### HasLockInfo

`func (o *BTOccurrenceData75) HasLockInfo() bool`

HasLockInfo returns a boolean if a field has been set.

### GetNodeId

`func (o *BTOccurrenceData75) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *BTOccurrenceData75) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *BTOccurrenceData75) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *BTOccurrenceData75) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetOccurrence

`func (o *BTOccurrenceData75) GetOccurrence() BTOccurrence74`

GetOccurrence returns the Occurrence field if non-nil, zero value otherwise.

### GetOccurrenceOk

`func (o *BTOccurrenceData75) GetOccurrenceOk() (*BTOccurrence74, bool)`

GetOccurrenceOk returns a tuple with the Occurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrence

`func (o *BTOccurrenceData75) SetOccurrence(v BTOccurrence74)`

SetOccurrence sets Occurrence field to given value.

### HasOccurrence

`func (o *BTOccurrenceData75) HasOccurrence() bool`

HasOccurrence returns a boolean if a field has been set.

### GetTransform

`func (o *BTOccurrenceData75) GetTransform() BTBSMatrix386`

GetTransform returns the Transform field if non-nil, zero value otherwise.

### GetTransformOk

`func (o *BTOccurrenceData75) GetTransformOk() (*BTBSMatrix386, bool)`

GetTransformOk returns a tuple with the Transform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransform

`func (o *BTOccurrenceData75) SetTransform(v BTBSMatrix386)`

SetTransform sets Transform field to given value.

### HasTransform

`func (o *BTOccurrenceData75) HasTransform() bool`

HasTransform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


