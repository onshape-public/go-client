# BTMassPropertiesInfoNull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Centroid** | Pointer to **[]float64** |  | [optional] 
**HasMass** | Pointer to **bool** |  | [optional] 
**Inertia** | Pointer to **[]float64** |  | [optional] 
**Mass** | Pointer to **[]float64** |  | [optional] 
**MassMissingCount** | Pointer to **int32** |  | [optional] 
**Periphery** | Pointer to **[]float64** |  | [optional] 
**PrincipalAxes** | Pointer to [**[]BTVector3d389**](BTVector3d-389.md) |  | [optional] 
**PrincipalInertia** | Pointer to **[]float64** |  | [optional] 
**Volume** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewBTMassPropertiesInfoNull

`func NewBTMassPropertiesInfoNull() *BTMassPropertiesInfoNull`

NewBTMassPropertiesInfoNull instantiates a new BTMassPropertiesInfoNull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMassPropertiesInfoNullWithDefaults

`func NewBTMassPropertiesInfoNullWithDefaults() *BTMassPropertiesInfoNull`

NewBTMassPropertiesInfoNullWithDefaults instantiates a new BTMassPropertiesInfoNull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCentroid

`func (o *BTMassPropertiesInfoNull) GetCentroid() []float64`

GetCentroid returns the Centroid field if non-nil, zero value otherwise.

### GetCentroidOk

`func (o *BTMassPropertiesInfoNull) GetCentroidOk() (*[]float64, bool)`

GetCentroidOk returns a tuple with the Centroid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentroid

`func (o *BTMassPropertiesInfoNull) SetCentroid(v []float64)`

SetCentroid sets Centroid field to given value.

### HasCentroid

`func (o *BTMassPropertiesInfoNull) HasCentroid() bool`

HasCentroid returns a boolean if a field has been set.

### GetHasMass

`func (o *BTMassPropertiesInfoNull) GetHasMass() bool`

GetHasMass returns the HasMass field if non-nil, zero value otherwise.

### GetHasMassOk

`func (o *BTMassPropertiesInfoNull) GetHasMassOk() (*bool, bool)`

GetHasMassOk returns a tuple with the HasMass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMass

`func (o *BTMassPropertiesInfoNull) SetHasMass(v bool)`

SetHasMass sets HasMass field to given value.

### HasHasMass

`func (o *BTMassPropertiesInfoNull) HasHasMass() bool`

HasHasMass returns a boolean if a field has been set.

### GetInertia

`func (o *BTMassPropertiesInfoNull) GetInertia() []float64`

GetInertia returns the Inertia field if non-nil, zero value otherwise.

### GetInertiaOk

`func (o *BTMassPropertiesInfoNull) GetInertiaOk() (*[]float64, bool)`

GetInertiaOk returns a tuple with the Inertia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInertia

`func (o *BTMassPropertiesInfoNull) SetInertia(v []float64)`

SetInertia sets Inertia field to given value.

### HasInertia

`func (o *BTMassPropertiesInfoNull) HasInertia() bool`

HasInertia returns a boolean if a field has been set.

### GetMass

`func (o *BTMassPropertiesInfoNull) GetMass() []float64`

GetMass returns the Mass field if non-nil, zero value otherwise.

### GetMassOk

`func (o *BTMassPropertiesInfoNull) GetMassOk() (*[]float64, bool)`

GetMassOk returns a tuple with the Mass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMass

`func (o *BTMassPropertiesInfoNull) SetMass(v []float64)`

SetMass sets Mass field to given value.

### HasMass

`func (o *BTMassPropertiesInfoNull) HasMass() bool`

HasMass returns a boolean if a field has been set.

### GetMassMissingCount

`func (o *BTMassPropertiesInfoNull) GetMassMissingCount() int32`

GetMassMissingCount returns the MassMissingCount field if non-nil, zero value otherwise.

### GetMassMissingCountOk

`func (o *BTMassPropertiesInfoNull) GetMassMissingCountOk() (*int32, bool)`

GetMassMissingCountOk returns a tuple with the MassMissingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMassMissingCount

`func (o *BTMassPropertiesInfoNull) SetMassMissingCount(v int32)`

SetMassMissingCount sets MassMissingCount field to given value.

### HasMassMissingCount

`func (o *BTMassPropertiesInfoNull) HasMassMissingCount() bool`

HasMassMissingCount returns a boolean if a field has been set.

### GetPeriphery

`func (o *BTMassPropertiesInfoNull) GetPeriphery() []float64`

GetPeriphery returns the Periphery field if non-nil, zero value otherwise.

### GetPeripheryOk

`func (o *BTMassPropertiesInfoNull) GetPeripheryOk() (*[]float64, bool)`

GetPeripheryOk returns a tuple with the Periphery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriphery

`func (o *BTMassPropertiesInfoNull) SetPeriphery(v []float64)`

SetPeriphery sets Periphery field to given value.

### HasPeriphery

`func (o *BTMassPropertiesInfoNull) HasPeriphery() bool`

HasPeriphery returns a boolean if a field has been set.

### GetPrincipalAxes

`func (o *BTMassPropertiesInfoNull) GetPrincipalAxes() []BTVector3d389`

GetPrincipalAxes returns the PrincipalAxes field if non-nil, zero value otherwise.

### GetPrincipalAxesOk

`func (o *BTMassPropertiesInfoNull) GetPrincipalAxesOk() (*[]BTVector3d389, bool)`

GetPrincipalAxesOk returns a tuple with the PrincipalAxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalAxes

`func (o *BTMassPropertiesInfoNull) SetPrincipalAxes(v []BTVector3d389)`

SetPrincipalAxes sets PrincipalAxes field to given value.

### HasPrincipalAxes

`func (o *BTMassPropertiesInfoNull) HasPrincipalAxes() bool`

HasPrincipalAxes returns a boolean if a field has been set.

### GetPrincipalInertia

`func (o *BTMassPropertiesInfoNull) GetPrincipalInertia() []float64`

GetPrincipalInertia returns the PrincipalInertia field if non-nil, zero value otherwise.

### GetPrincipalInertiaOk

`func (o *BTMassPropertiesInfoNull) GetPrincipalInertiaOk() (*[]float64, bool)`

GetPrincipalInertiaOk returns a tuple with the PrincipalInertia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalInertia

`func (o *BTMassPropertiesInfoNull) SetPrincipalInertia(v []float64)`

SetPrincipalInertia sets PrincipalInertia field to given value.

### HasPrincipalInertia

`func (o *BTMassPropertiesInfoNull) HasPrincipalInertia() bool`

HasPrincipalInertia returns a boolean if a field has been set.

### GetVolume

`func (o *BTMassPropertiesInfoNull) GetVolume() []float64`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *BTMassPropertiesInfoNull) GetVolumeOk() (*[]float64, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *BTMassPropertiesInfoNull) SetVolume(v []float64)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *BTMassPropertiesInfoNull) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


