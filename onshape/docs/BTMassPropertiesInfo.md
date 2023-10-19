# BTMassPropertiesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Centroid** | Pointer to **[]float64** |  | [optional] 
**HasMass** | Pointer to **bool** |  | [optional] 
**Inertia** | Pointer to **[]float64** |  | [optional] 
**Mass** | Pointer to **[]float64** |  | [optional] 
**MassMissingCount** | Pointer to **int32** |  | [optional] 
**Periphery** | Pointer to **[]float64** |  | [optional] 
**PrincipalAxes** | Pointer to [**[]BTVector3d389**](BTVector3d389.md) |  | [optional] 
**PrincipalInertia** | Pointer to **[]float64** |  | [optional] 
**Volume** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewBTMassPropertiesInfo

`func NewBTMassPropertiesInfo() *BTMassPropertiesInfo`

NewBTMassPropertiesInfo instantiates a new BTMassPropertiesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMassPropertiesInfoWithDefaults

`func NewBTMassPropertiesInfoWithDefaults() *BTMassPropertiesInfo`

NewBTMassPropertiesInfoWithDefaults instantiates a new BTMassPropertiesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentroid

`func (o *BTMassPropertiesInfo) GetCentroid() []float64`

GetCentroid returns the Centroid field if non-nil, zero value otherwise.

### GetCentroidOk

`func (o *BTMassPropertiesInfo) GetCentroidOk() (*[]float64, bool)`

GetCentroidOk returns a tuple with the Centroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentroid

`func (o *BTMassPropertiesInfo) SetCentroid(v []float64)`

SetCentroid sets Centroid field to given value.

### HasCentroid

`func (o *BTMassPropertiesInfo) HasCentroid() bool`

HasCentroid returns a boolean if a field has been set.

### GetHasMass

`func (o *BTMassPropertiesInfo) GetHasMass() bool`

GetHasMass returns the HasMass field if non-nil, zero value otherwise.

### GetHasMassOk

`func (o *BTMassPropertiesInfo) GetHasMassOk() (*bool, bool)`

GetHasMassOk returns a tuple with the HasMass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMass

`func (o *BTMassPropertiesInfo) SetHasMass(v bool)`

SetHasMass sets HasMass field to given value.

### HasHasMass

`func (o *BTMassPropertiesInfo) HasHasMass() bool`

HasHasMass returns a boolean if a field has been set.

### GetInertia

`func (o *BTMassPropertiesInfo) GetInertia() []float64`

GetInertia returns the Inertia field if non-nil, zero value otherwise.

### GetInertiaOk

`func (o *BTMassPropertiesInfo) GetInertiaOk() (*[]float64, bool)`

GetInertiaOk returns a tuple with the Inertia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInertia

`func (o *BTMassPropertiesInfo) SetInertia(v []float64)`

SetInertia sets Inertia field to given value.

### HasInertia

`func (o *BTMassPropertiesInfo) HasInertia() bool`

HasInertia returns a boolean if a field has been set.

### GetMass

`func (o *BTMassPropertiesInfo) GetMass() []float64`

GetMass returns the Mass field if non-nil, zero value otherwise.

### GetMassOk

`func (o *BTMassPropertiesInfo) GetMassOk() (*[]float64, bool)`

GetMassOk returns a tuple with the Mass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMass

`func (o *BTMassPropertiesInfo) SetMass(v []float64)`

SetMass sets Mass field to given value.

### HasMass

`func (o *BTMassPropertiesInfo) HasMass() bool`

HasMass returns a boolean if a field has been set.

### GetMassMissingCount

`func (o *BTMassPropertiesInfo) GetMassMissingCount() int32`

GetMassMissingCount returns the MassMissingCount field if non-nil, zero value otherwise.

### GetMassMissingCountOk

`func (o *BTMassPropertiesInfo) GetMassMissingCountOk() (*int32, bool)`

GetMassMissingCountOk returns a tuple with the MassMissingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMassMissingCount

`func (o *BTMassPropertiesInfo) SetMassMissingCount(v int32)`

SetMassMissingCount sets MassMissingCount field to given value.

### HasMassMissingCount

`func (o *BTMassPropertiesInfo) HasMassMissingCount() bool`

HasMassMissingCount returns a boolean if a field has been set.

### GetPeriphery

`func (o *BTMassPropertiesInfo) GetPeriphery() []float64`

GetPeriphery returns the Periphery field if non-nil, zero value otherwise.

### GetPeripheryOk

`func (o *BTMassPropertiesInfo) GetPeripheryOk() (*[]float64, bool)`

GetPeripheryOk returns a tuple with the Periphery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriphery

`func (o *BTMassPropertiesInfo) SetPeriphery(v []float64)`

SetPeriphery sets Periphery field to given value.

### HasPeriphery

`func (o *BTMassPropertiesInfo) HasPeriphery() bool`

HasPeriphery returns a boolean if a field has been set.

### GetPrincipalAxes

`func (o *BTMassPropertiesInfo) GetPrincipalAxes() []BTVector3d389`

GetPrincipalAxes returns the PrincipalAxes field if non-nil, zero value otherwise.

### GetPrincipalAxesOk

`func (o *BTMassPropertiesInfo) GetPrincipalAxesOk() (*[]BTVector3d389, bool)`

GetPrincipalAxesOk returns a tuple with the PrincipalAxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalAxes

`func (o *BTMassPropertiesInfo) SetPrincipalAxes(v []BTVector3d389)`

SetPrincipalAxes sets PrincipalAxes field to given value.

### HasPrincipalAxes

`func (o *BTMassPropertiesInfo) HasPrincipalAxes() bool`

HasPrincipalAxes returns a boolean if a field has been set.

### GetPrincipalInertia

`func (o *BTMassPropertiesInfo) GetPrincipalInertia() []float64`

GetPrincipalInertia returns the PrincipalInertia field if non-nil, zero value otherwise.

### GetPrincipalInertiaOk

`func (o *BTMassPropertiesInfo) GetPrincipalInertiaOk() (*[]float64, bool)`

GetPrincipalInertiaOk returns a tuple with the PrincipalInertia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalInertia

`func (o *BTMassPropertiesInfo) SetPrincipalInertia(v []float64)`

SetPrincipalInertia sets PrincipalInertia field to given value.

### HasPrincipalInertia

`func (o *BTMassPropertiesInfo) HasPrincipalInertia() bool`

HasPrincipalInertia returns a boolean if a field has been set.

### GetVolume

`func (o *BTMassPropertiesInfo) GetVolume() []float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *BTMassPropertiesInfo) GetVolumeOk() (*[]float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *BTMassPropertiesInfo) SetVolume(v []float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *BTMassPropertiesInfo) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


