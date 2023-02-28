# BTAssemblyPatternInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Id of the pattern. | [optional] 
**Name** | Pointer to **string** | Name of the pattern. | [optional] 
**SeedToPatternInstances** | Pointer to **map[string][]string** | Mapping of seed to pattern instance ids. | [optional] 
**Suppressed** | Pointer to **bool** | If pattern is suppressed. | [optional] 
**Type** | Pointer to **string** | Pattern type. | [optional] 

## Methods

### NewBTAssemblyPatternInfo

`func NewBTAssemblyPatternInfo() *BTAssemblyPatternInfo`

NewBTAssemblyPatternInfo instantiates a new BTAssemblyPatternInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyPatternInfoWithDefaults

`func NewBTAssemblyPatternInfoWithDefaults() *BTAssemblyPatternInfo`

NewBTAssemblyPatternInfoWithDefaults instantiates a new BTAssemblyPatternInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BTAssemblyPatternInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTAssemblyPatternInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTAssemblyPatternInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTAssemblyPatternInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTAssemblyPatternInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAssemblyPatternInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAssemblyPatternInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAssemblyPatternInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSeedToPatternInstances

`func (o *BTAssemblyPatternInfo) GetSeedToPatternInstances() map[string][]string`

GetSeedToPatternInstances returns the SeedToPatternInstances field if non-nil, zero value otherwise.

### GetSeedToPatternInstancesOk

`func (o *BTAssemblyPatternInfo) GetSeedToPatternInstancesOk() (*map[string][]string, bool)`

GetSeedToPatternInstancesOk returns a tuple with the SeedToPatternInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedToPatternInstances

`func (o *BTAssemblyPatternInfo) SetSeedToPatternInstances(v map[string][]string)`

SetSeedToPatternInstances sets SeedToPatternInstances field to given value.

### HasSeedToPatternInstances

`func (o *BTAssemblyPatternInfo) HasSeedToPatternInstances() bool`

HasSeedToPatternInstances returns a boolean if a field has been set.

### GetSuppressed

`func (o *BTAssemblyPatternInfo) GetSuppressed() bool`

GetSuppressed returns the Suppressed field if non-nil, zero value otherwise.

### GetSuppressedOk

`func (o *BTAssemblyPatternInfo) GetSuppressedOk() (*bool, bool)`

GetSuppressedOk returns a tuple with the Suppressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressed

`func (o *BTAssemblyPatternInfo) SetSuppressed(v bool)`

SetSuppressed sets Suppressed field to given value.

### HasSuppressed

`func (o *BTAssemblyPatternInfo) HasSuppressed() bool`

HasSuppressed returns a boolean if a field has been set.

### GetType

`func (o *BTAssemblyPatternInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTAssemblyPatternInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTAssemblyPatternInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTAssemblyPatternInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


