# Camera

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Orthographic** | Pointer to [**CameraOrthographic**](CameraOrthographic.md) |  | [optional] 
**Perspective** | Pointer to [**CameraPerspective**](CameraPerspective.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCamera

`func NewCamera() *Camera`

NewCamera instantiates a new Camera object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCameraWithDefaults

`func NewCameraWithDefaults() *Camera`

NewCameraWithDefaults instantiates a new Camera object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensions

`func (o *Camera) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Camera) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Camera) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Camera) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetExtras

`func (o *Camera) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *Camera) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *Camera) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *Camera) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### GetName

`func (o *Camera) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Camera) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Camera) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Camera) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrthographic

`func (o *Camera) GetOrthographic() CameraOrthographic`

GetOrthographic returns the Orthographic field if non-nil, zero value otherwise.

### GetOrthographicOk

`func (o *Camera) GetOrthographicOk() (*CameraOrthographic, bool)`

GetOrthographicOk returns a tuple with the Orthographic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrthographic

`func (o *Camera) SetOrthographic(v CameraOrthographic)`

SetOrthographic sets Orthographic field to given value.

### HasOrthographic

`func (o *Camera) HasOrthographic() bool`

HasOrthographic returns a boolean if a field has been set.

### GetPerspective

`func (o *Camera) GetPerspective() CameraPerspective`

GetPerspective returns the Perspective field if non-nil, zero value otherwise.

### GetPerspectiveOk

`func (o *Camera) GetPerspectiveOk() (*CameraPerspective, bool)`

GetPerspectiveOk returns a tuple with the Perspective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerspective

`func (o *Camera) SetPerspective(v CameraPerspective)`

SetPerspective sets Perspective field to given value.

### HasPerspective

`func (o *Camera) HasPerspective() bool`

HasPerspective returns a boolean if a field has been set.

### GetType

`func (o *Camera) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Camera) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Camera) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Camera) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


