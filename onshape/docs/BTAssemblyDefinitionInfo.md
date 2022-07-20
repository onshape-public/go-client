# BTAssemblyDefinitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartStudioFeatures** | Pointer to [**[]BTAssemblyPsFeatureInfo**](BTAssemblyPsFeatureInfo.md) |  | [optional] 
**Parts** | Pointer to [**[]BTAssemblyPartInfo**](BTAssemblyPartInfo.md) |  | [optional] 
**RootAssembly** | Pointer to [**BTRootAssemblyInfo**](BTRootAssemblyInfo.md) |  | [optional] 
**SubAssemblies** | Pointer to [**[]BTSubAssemblyInfo**](BTSubAssemblyInfo.md) |  | [optional] 

## Methods

### NewBTAssemblyDefinitionInfo

`func NewBTAssemblyDefinitionInfo() *BTAssemblyDefinitionInfo`

NewBTAssemblyDefinitionInfo instantiates a new BTAssemblyDefinitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAssemblyDefinitionInfoWithDefaults

`func NewBTAssemblyDefinitionInfoWithDefaults() *BTAssemblyDefinitionInfo`

NewBTAssemblyDefinitionInfoWithDefaults instantiates a new BTAssemblyDefinitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartStudioFeatures

`func (o *BTAssemblyDefinitionInfo) GetPartStudioFeatures() []BTAssemblyPsFeatureInfo`

GetPartStudioFeatures returns the PartStudioFeatures field if non-nil, zero value otherwise.

### GetPartStudioFeaturesOk

`func (o *BTAssemblyDefinitionInfo) GetPartStudioFeaturesOk() (*[]BTAssemblyPsFeatureInfo, bool)`

GetPartStudioFeaturesOk returns a tuple with the PartStudioFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartStudioFeatures

`func (o *BTAssemblyDefinitionInfo) SetPartStudioFeatures(v []BTAssemblyPsFeatureInfo)`

SetPartStudioFeatures sets PartStudioFeatures field to given value.

### HasPartStudioFeatures

`func (o *BTAssemblyDefinitionInfo) HasPartStudioFeatures() bool`

HasPartStudioFeatures returns a boolean if a field has been set.

### GetParts

`func (o *BTAssemblyDefinitionInfo) GetParts() []BTAssemblyPartInfo`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *BTAssemblyDefinitionInfo) GetPartsOk() (*[]BTAssemblyPartInfo, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *BTAssemblyDefinitionInfo) SetParts(v []BTAssemblyPartInfo)`

SetParts sets Parts field to given value.

### HasParts

`func (o *BTAssemblyDefinitionInfo) HasParts() bool`

HasParts returns a boolean if a field has been set.

### GetRootAssembly

`func (o *BTAssemblyDefinitionInfo) GetRootAssembly() BTRootAssemblyInfo`

GetRootAssembly returns the RootAssembly field if non-nil, zero value otherwise.

### GetRootAssemblyOk

`func (o *BTAssemblyDefinitionInfo) GetRootAssemblyOk() (*BTRootAssemblyInfo, bool)`

GetRootAssemblyOk returns a tuple with the RootAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootAssembly

`func (o *BTAssemblyDefinitionInfo) SetRootAssembly(v BTRootAssemblyInfo)`

SetRootAssembly sets RootAssembly field to given value.

### HasRootAssembly

`func (o *BTAssemblyDefinitionInfo) HasRootAssembly() bool`

HasRootAssembly returns a boolean if a field has been set.

### GetSubAssemblies

`func (o *BTAssemblyDefinitionInfo) GetSubAssemblies() []BTSubAssemblyInfo`

GetSubAssemblies returns the SubAssemblies field if non-nil, zero value otherwise.

### GetSubAssembliesOk

`func (o *BTAssemblyDefinitionInfo) GetSubAssembliesOk() (*[]BTSubAssemblyInfo, bool)`

GetSubAssembliesOk returns a tuple with the SubAssemblies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAssemblies

`func (o *BTAssemblyDefinitionInfo) SetSubAssemblies(v []BTSubAssemblyInfo)`

SetSubAssemblies sets SubAssemblies field to given value.

### HasSubAssemblies

`func (o *BTAssemblyDefinitionInfo) HasSubAssemblies() bool`

HasSubAssemblies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


