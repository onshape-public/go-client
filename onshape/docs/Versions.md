# Versions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableVersions** | Pointer to [**[]BTApiVersion**](BTApiVersion.md) |  | [optional] 
**SpecifiedVersion** | Pointer to [**BTApiVersion**](BTApiVersion.md) |  | [optional] 

## Methods

### NewVersions

`func NewVersions() *Versions`

NewVersions instantiates a new Versions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionsWithDefaults

`func NewVersionsWithDefaults() *Versions`

NewVersionsWithDefaults instantiates a new Versions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableVersions

`func (o *Versions) GetAvailableVersions() []BTApiVersion`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *Versions) GetAvailableVersionsOk() (*[]BTApiVersion, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *Versions) SetAvailableVersions(v []BTApiVersion)`

SetAvailableVersions sets AvailableVersions field to given value.

### HasAvailableVersions

`func (o *Versions) HasAvailableVersions() bool`

HasAvailableVersions returns a boolean if a field has been set.

### GetSpecifiedVersion

`func (o *Versions) GetSpecifiedVersion() BTApiVersion`

GetSpecifiedVersion returns the SpecifiedVersion field if non-nil, zero value otherwise.

### GetSpecifiedVersionOk

`func (o *Versions) GetSpecifiedVersionOk() (*BTApiVersion, bool)`

GetSpecifiedVersionOk returns a tuple with the SpecifiedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedVersion

`func (o *Versions) SetSpecifiedVersion(v BTApiVersion)`

SetSpecifiedVersion sets SpecifiedVersion field to given value.

### HasSpecifiedVersion

`func (o *Versions) HasSpecifiedVersion() bool`

HasSpecifiedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


