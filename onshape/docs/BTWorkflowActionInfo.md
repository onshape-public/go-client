# BTWorkflowActionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**AllowIfApprovers** | Pointer to **bool** |  | [optional] 
**AllowIfNoApprovers** | Pointer to **bool** |  | [optional] 
**AlwaysAllow** | Pointer to **bool** |  | [optional] 
**IsAdminOverride** | Pointer to **bool** |  | [optional] 
**IsApproverAction** | Pointer to **bool** |  | [optional] 
**IsCreatorOverride** | Pointer to **bool** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**RequiredProperties** | Pointer to **[]string** |  | [optional] 
**Tooltip** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**BTTransitionType**](BTTransitionType.md) |  | [optional] 
**UiHint** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWorkflowActionInfo

`func NewBTWorkflowActionInfo() *BTWorkflowActionInfo`

NewBTWorkflowActionInfo instantiates a new BTWorkflowActionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowActionInfoWithDefaults

`func NewBTWorkflowActionInfoWithDefaults() *BTWorkflowActionInfo`

NewBTWorkflowActionInfoWithDefaults instantiates a new BTWorkflowActionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BTWorkflowActionInfo) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BTWorkflowActionInfo) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BTWorkflowActionInfo) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *BTWorkflowActionInfo) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAllowIfApprovers

`func (o *BTWorkflowActionInfo) GetAllowIfApprovers() bool`

GetAllowIfApprovers returns the AllowIfApprovers field if non-nil, zero value otherwise.

### GetAllowIfApproversOk

`func (o *BTWorkflowActionInfo) GetAllowIfApproversOk() (*bool, bool)`

GetAllowIfApproversOk returns a tuple with the AllowIfApprovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIfApprovers

`func (o *BTWorkflowActionInfo) SetAllowIfApprovers(v bool)`

SetAllowIfApprovers sets AllowIfApprovers field to given value.

### HasAllowIfApprovers

`func (o *BTWorkflowActionInfo) HasAllowIfApprovers() bool`

HasAllowIfApprovers returns a boolean if a field has been set.

### GetAllowIfNoApprovers

`func (o *BTWorkflowActionInfo) GetAllowIfNoApprovers() bool`

GetAllowIfNoApprovers returns the AllowIfNoApprovers field if non-nil, zero value otherwise.

### GetAllowIfNoApproversOk

`func (o *BTWorkflowActionInfo) GetAllowIfNoApproversOk() (*bool, bool)`

GetAllowIfNoApproversOk returns a tuple with the AllowIfNoApprovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIfNoApprovers

`func (o *BTWorkflowActionInfo) SetAllowIfNoApprovers(v bool)`

SetAllowIfNoApprovers sets AllowIfNoApprovers field to given value.

### HasAllowIfNoApprovers

`func (o *BTWorkflowActionInfo) HasAllowIfNoApprovers() bool`

HasAllowIfNoApprovers returns a boolean if a field has been set.

### GetAlwaysAllow

`func (o *BTWorkflowActionInfo) GetAlwaysAllow() bool`

GetAlwaysAllow returns the AlwaysAllow field if non-nil, zero value otherwise.

### GetAlwaysAllowOk

`func (o *BTWorkflowActionInfo) GetAlwaysAllowOk() (*bool, bool)`

GetAlwaysAllowOk returns a tuple with the AlwaysAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysAllow

`func (o *BTWorkflowActionInfo) SetAlwaysAllow(v bool)`

SetAlwaysAllow sets AlwaysAllow field to given value.

### HasAlwaysAllow

`func (o *BTWorkflowActionInfo) HasAlwaysAllow() bool`

HasAlwaysAllow returns a boolean if a field has been set.

### GetIsAdminOverride

`func (o *BTWorkflowActionInfo) GetIsAdminOverride() bool`

GetIsAdminOverride returns the IsAdminOverride field if non-nil, zero value otherwise.

### GetIsAdminOverrideOk

`func (o *BTWorkflowActionInfo) GetIsAdminOverrideOk() (*bool, bool)`

GetIsAdminOverrideOk returns a tuple with the IsAdminOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdminOverride

`func (o *BTWorkflowActionInfo) SetIsAdminOverride(v bool)`

SetIsAdminOverride sets IsAdminOverride field to given value.

### HasIsAdminOverride

`func (o *BTWorkflowActionInfo) HasIsAdminOverride() bool`

HasIsAdminOverride returns a boolean if a field has been set.

### GetIsApproverAction

`func (o *BTWorkflowActionInfo) GetIsApproverAction() bool`

GetIsApproverAction returns the IsApproverAction field if non-nil, zero value otherwise.

### GetIsApproverActionOk

`func (o *BTWorkflowActionInfo) GetIsApproverActionOk() (*bool, bool)`

GetIsApproverActionOk returns a tuple with the IsApproverAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproverAction

`func (o *BTWorkflowActionInfo) SetIsApproverAction(v bool)`

SetIsApproverAction sets IsApproverAction field to given value.

### HasIsApproverAction

`func (o *BTWorkflowActionInfo) HasIsApproverAction() bool`

HasIsApproverAction returns a boolean if a field has been set.

### GetIsCreatorOverride

`func (o *BTWorkflowActionInfo) GetIsCreatorOverride() bool`

GetIsCreatorOverride returns the IsCreatorOverride field if non-nil, zero value otherwise.

### GetIsCreatorOverrideOk

`func (o *BTWorkflowActionInfo) GetIsCreatorOverrideOk() (*bool, bool)`

GetIsCreatorOverrideOk returns a tuple with the IsCreatorOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCreatorOverride

`func (o *BTWorkflowActionInfo) SetIsCreatorOverride(v bool)`

SetIsCreatorOverride sets IsCreatorOverride field to given value.

### HasIsCreatorOverride

`func (o *BTWorkflowActionInfo) HasIsCreatorOverride() bool`

HasIsCreatorOverride returns a boolean if a field has been set.

### GetLabel

`func (o *BTWorkflowActionInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BTWorkflowActionInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BTWorkflowActionInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BTWorkflowActionInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRequiredProperties

`func (o *BTWorkflowActionInfo) GetRequiredProperties() []string`

GetRequiredProperties returns the RequiredProperties field if non-nil, zero value otherwise.

### GetRequiredPropertiesOk

`func (o *BTWorkflowActionInfo) GetRequiredPropertiesOk() (*[]string, bool)`

GetRequiredPropertiesOk returns a tuple with the RequiredProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredProperties

`func (o *BTWorkflowActionInfo) SetRequiredProperties(v []string)`

SetRequiredProperties sets RequiredProperties field to given value.

### HasRequiredProperties

`func (o *BTWorkflowActionInfo) HasRequiredProperties() bool`

HasRequiredProperties returns a boolean if a field has been set.

### GetTooltip

`func (o *BTWorkflowActionInfo) GetTooltip() string`

GetTooltip returns the Tooltip field if non-nil, zero value otherwise.

### GetTooltipOk

`func (o *BTWorkflowActionInfo) GetTooltipOk() (*string, bool)`

GetTooltipOk returns a tuple with the Tooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltip

`func (o *BTWorkflowActionInfo) SetTooltip(v string)`

SetTooltip sets Tooltip field to given value.

### HasTooltip

`func (o *BTWorkflowActionInfo) HasTooltip() bool`

HasTooltip returns a boolean if a field has been set.

### GetType

`func (o *BTWorkflowActionInfo) GetType() BTTransitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTWorkflowActionInfo) GetTypeOk() (*BTTransitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTWorkflowActionInfo) SetType(v BTTransitionType)`

SetType sets Type field to given value.

### HasType

`func (o *BTWorkflowActionInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUiHint

`func (o *BTWorkflowActionInfo) GetUiHint() string`

GetUiHint returns the UiHint field if non-nil, zero value otherwise.

### GetUiHintOk

`func (o *BTWorkflowActionInfo) GetUiHintOk() (*string, bool)`

GetUiHintOk returns a tuple with the UiHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiHint

`func (o *BTWorkflowActionInfo) SetUiHint(v string)`

SetUiHint sets UiHint field to given value.

### HasUiHint

`func (o *BTWorkflowActionInfo) HasUiHint() bool`

HasUiHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


