# Skin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**InverseBindMatrices** | Pointer to **int32** |  | [optional] 
**Joints** | Pointer to **[]int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Skeleton** | Pointer to **int32** |  | [optional] 

## Methods

### NewSkin

`func NewSkin() *Skin`

NewSkin instantiates a new Skin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkinWithDefaults

`func NewSkinWithDefaults() *Skin`

NewSkinWithDefaults instantiates a new Skin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *Skin) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Skin) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Skin) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Skin) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Skin) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Skin) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Skin) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Skin) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetInverseBindMatrices

`func (o *Skin) GetInverseBindMatrices() int32`

GetInverseBindMatrices returns the InverseBindMatrices field if non-nil, zero value otherwise.

### GetInverseBindMatricesOk

`func (o *Skin) GetInverseBindMatricesOk() (*int32, bool)`

GetInverseBindMatricesOk returns a tuple with the InverseBindMatrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseBindMatrices

`func (o *Skin) SetInverseBindMatrices(v int32)`

SetInverseBindMatrices sets InverseBindMatrices field to given value.

### HasInverseBindMatrices

`func (o *Skin) HasInverseBindMatrices() bool`

HasInverseBindMatrices returns a boolean if a field has been set.

### GetJoints

`func (o *Skin) GetJoints() []int32`

GetJoints returns the Joints field if non-nil, zero value otherwise.

### GetJointsOk

`func (o *Skin) GetJointsOk() (*[]int32, bool)`

GetJointsOk returns a tuple with the Joints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoints

`func (o *Skin) SetJoints(v []int32)`

SetJoints sets Joints field to given value.

### HasJoints

`func (o *Skin) HasJoints() bool`

HasJoints returns a boolean if a field has been set.

### GetName

`func (o *Skin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Skin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSkeleton

`func (o *Skin) GetSkeleton() int32`

GetSkeleton returns the Skeleton field if non-nil, zero value otherwise.

### GetSkeletonOk

`func (o *Skin) GetSkeletonOk() (*int32, bool)`

GetSkeletonOk returns a tuple with the Skeleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkeleton

`func (o *Skin) SetSkeleton(v int32)`

SetSkeleton sets Skeleton field to given value.

### HasSkeleton

`func (o *Skin) HasSkeleton() bool`

HasSkeleton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


