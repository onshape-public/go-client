# BTVersionOrWorkspaceMergeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMergeStrategy** | Pointer to [**BTMergeStrategy**](BTMergeStrategy.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MergeStrategyElementOverrides** | Pointer to [**map[string]BTMergeStrategy**](BTMergeStrategy.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBTVersionOrWorkspaceMergeInfo

`func NewBTVersionOrWorkspaceMergeInfo() *BTVersionOrWorkspaceMergeInfo`

NewBTVersionOrWorkspaceMergeInfo instantiates a new BTVersionOrWorkspaceMergeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTVersionOrWorkspaceMergeInfoWithDefaults

`func NewBTVersionOrWorkspaceMergeInfoWithDefaults() *BTVersionOrWorkspaceMergeInfo`

NewBTVersionOrWorkspaceMergeInfoWithDefaults instantiates a new BTVersionOrWorkspaceMergeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMergeStrategy

`func (o *BTVersionOrWorkspaceMergeInfo) GetDefaultMergeStrategy() BTMergeStrategy`

GetDefaultMergeStrategy returns the DefaultMergeStrategy field if non-nil, zero value otherwise.

### GetDefaultMergeStrategyOk

`func (o *BTVersionOrWorkspaceMergeInfo) GetDefaultMergeStrategyOk() (*BTMergeStrategy, bool)`

GetDefaultMergeStrategyOk returns a tuple with the DefaultMergeStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMergeStrategy

`func (o *BTVersionOrWorkspaceMergeInfo) SetDefaultMergeStrategy(v BTMergeStrategy)`

SetDefaultMergeStrategy sets DefaultMergeStrategy field to given value.

### HasDefaultMergeStrategy

`func (o *BTVersionOrWorkspaceMergeInfo) HasDefaultMergeStrategy() bool`

HasDefaultMergeStrategy returns a boolean if a field has been set.

### GetId

`func (o *BTVersionOrWorkspaceMergeInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTVersionOrWorkspaceMergeInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTVersionOrWorkspaceMergeInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTVersionOrWorkspaceMergeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMergeStrategyElementOverrides

`func (o *BTVersionOrWorkspaceMergeInfo) GetMergeStrategyElementOverrides() map[string]BTMergeStrategy`

GetMergeStrategyElementOverrides returns the MergeStrategyElementOverrides field if non-nil, zero value otherwise.

### GetMergeStrategyElementOverridesOk

`func (o *BTVersionOrWorkspaceMergeInfo) GetMergeStrategyElementOverridesOk() (*map[string]BTMergeStrategy, bool)`

GetMergeStrategyElementOverridesOk returns a tuple with the MergeStrategyElementOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStrategyElementOverrides

`func (o *BTVersionOrWorkspaceMergeInfo) SetMergeStrategyElementOverrides(v map[string]BTMergeStrategy)`

SetMergeStrategyElementOverrides sets MergeStrategyElementOverrides field to given value.

### HasMergeStrategyElementOverrides

`func (o *BTVersionOrWorkspaceMergeInfo) HasMergeStrategyElementOverrides() bool`

HasMergeStrategyElementOverrides returns a boolean if a field has been set.

### GetType

`func (o *BTVersionOrWorkspaceMergeInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTVersionOrWorkspaceMergeInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTVersionOrWorkspaceMergeInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTVersionOrWorkspaceMergeInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


