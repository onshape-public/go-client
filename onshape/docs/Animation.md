# Animation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]AnimationChannel**](AnimationChannel.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Samplers** | Pointer to [**[]AnimationSampler**](AnimationSampler.md) |  | [optional] 

## Methods

### NewAnimation

`func NewAnimation() *Animation`

NewAnimation instantiates a new Animation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnimationWithDefaults

`func NewAnimationWithDefaults() *Animation`

NewAnimationWithDefaults instantiates a new Animation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *Animation) GetChannels() []AnimationChannel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *Animation) GetChannelsOk() (*[]AnimationChannel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *Animation) SetChannels(v []AnimationChannel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *Animation) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetExtensions

`func (o *Animation) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Animation) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Animation) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Animation) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Animation) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Animation) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Animation) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Animation) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetName

`func (o *Animation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Animation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Animation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Animation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSamplers

`func (o *Animation) GetSamplers() []AnimationSampler`

GetSamplers returns the Samplers field if non-nil, zero value otherwise.

### GetSamplersOk

`func (o *Animation) GetSamplersOk() (*[]AnimationSampler, bool)`

GetSamplersOk returns a tuple with the Samplers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplers

`func (o *Animation) SetSamplers(v []AnimationSampler)`

SetSamplers sets Samplers field to given value.

### HasSamplers

`func (o *Animation) HasSamplers() bool`

HasSamplers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


