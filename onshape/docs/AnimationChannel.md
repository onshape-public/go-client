# AnimationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Sampler** | Pointer to **int32** |  | [optional] 
**Target** | Pointer to [**AnimationChannelTarget**](AnimationChannelTarget.md) |  | [optional] 

## Methods

### NewAnimationChannel

`func NewAnimationChannel() *AnimationChannel`

NewAnimationChannel instantiates a new AnimationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnimationChannelWithDefaults

`func NewAnimationChannelWithDefaults() *AnimationChannel`

NewAnimationChannelWithDefaults instantiates a new AnimationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *AnimationChannel) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *AnimationChannel) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *AnimationChannel) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *AnimationChannel) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *AnimationChannel) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *AnimationChannel) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *AnimationChannel) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *AnimationChannel) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetSampler

`func (o *AnimationChannel) GetSampler() int32`

GetSampler returns the Sampler field if non-nil, zero value otherwise.

### GetSamplerOk

`func (o *AnimationChannel) GetSamplerOk() (*int32, bool)`

GetSamplerOk returns a tuple with the Sampler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampler

`func (o *AnimationChannel) SetSampler(v int32)`

SetSampler sets Sampler field to given value.

### HasSampler

`func (o *AnimationChannel) HasSampler() bool`

HasSampler returns a boolean if a field has been set.

### GetTarget

`func (o *AnimationChannel) GetTarget() AnimationChannelTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AnimationChannel) GetTargetOk() (*AnimationChannelTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AnimationChannel) SetTarget(v AnimationChannelTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AnimationChannel) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


