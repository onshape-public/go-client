# BTBaseEntityData33

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ConstructionPlane** | Pointer to **bool** |  | [optional] 
**CopyWithoutGeometry** | Pointer to [**BTBaseEntityData33**](BTBaseEntityData33.md) |  | [optional] 
**Decompressed** | Pointer to [**BTBaseEntityData33**](BTBaseEntityData33.md) |  | [optional] 
**Deletion** | Pointer to **bool** |  | [optional] 
**FeatureIds** | Pointer to **[]string** |  | [optional] 
**FromSketch** | Pointer to **bool** |  | [optional] 
**Geometries** | Pointer to [**[]BTEntityGeometry35**](BTEntityGeometry35.md) |  | [optional] 
**Origin** | Pointer to **bool** |  | [optional] 

## Methods

### NewBTBaseEntityData33

`func NewBTBaseEntityData33() *BTBaseEntityData33`

NewBTBaseEntityData33 instantiates a new BTBaseEntityData33 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBaseEntityData33WithDefaults

`func NewBTBaseEntityData33WithDefaults() *BTBaseEntityData33`

NewBTBaseEntityData33WithDefaults instantiates a new BTBaseEntityData33 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTBaseEntityData33) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTBaseEntityData33) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTBaseEntityData33) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTBaseEntityData33) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetConstructionPlane

`func (o *BTBaseEntityData33) GetConstructionPlane() bool`

GetConstructionPlane returns the ConstructionPlane field if non-nil, zero value otherwise.

### GetConstructionPlaneOk

`func (o *BTBaseEntityData33) GetConstructionPlaneOk() (*bool, bool)`

GetConstructionPlaneOk returns a tuple with the ConstructionPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstructionPlane

`func (o *BTBaseEntityData33) SetConstructionPlane(v bool)`

SetConstructionPlane sets ConstructionPlane field to given value.

### HasConstructionPlane

`func (o *BTBaseEntityData33) HasConstructionPlane() bool`

HasConstructionPlane returns a boolean if a field has been set.

### GetCopyWithoutGeometry

`func (o *BTBaseEntityData33) GetCopyWithoutGeometry() BTBaseEntityData33`

GetCopyWithoutGeometry returns the CopyWithoutGeometry field if non-nil, zero value otherwise.

### GetCopyWithoutGeometryOk

`func (o *BTBaseEntityData33) GetCopyWithoutGeometryOk() (*BTBaseEntityData33, bool)`

GetCopyWithoutGeometryOk returns a tuple with the CopyWithoutGeometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyWithoutGeometry

`func (o *BTBaseEntityData33) SetCopyWithoutGeometry(v BTBaseEntityData33)`

SetCopyWithoutGeometry sets CopyWithoutGeometry field to given value.

### HasCopyWithoutGeometry

`func (o *BTBaseEntityData33) HasCopyWithoutGeometry() bool`

HasCopyWithoutGeometry returns a boolean if a field has been set.

### GetDecompressed

`func (o *BTBaseEntityData33) GetDecompressed() BTBaseEntityData33`

GetDecompressed returns the Decompressed field if non-nil, zero value otherwise.

### GetDecompressedOk

`func (o *BTBaseEntityData33) GetDecompressedOk() (*BTBaseEntityData33, bool)`

GetDecompressedOk returns a tuple with the Decompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompressed

`func (o *BTBaseEntityData33) SetDecompressed(v BTBaseEntityData33)`

SetDecompressed sets Decompressed field to given value.

### HasDecompressed

`func (o *BTBaseEntityData33) HasDecompressed() bool`

HasDecompressed returns a boolean if a field has been set.

### GetDeletion

`func (o *BTBaseEntityData33) GetDeletion() bool`

GetDeletion returns the Deletion field if non-nil, zero value otherwise.

### GetDeletionOk

`func (o *BTBaseEntityData33) GetDeletionOk() (*bool, bool)`

GetDeletionOk returns a tuple with the Deletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletion

`func (o *BTBaseEntityData33) SetDeletion(v bool)`

SetDeletion sets Deletion field to given value.

### HasDeletion

`func (o *BTBaseEntityData33) HasDeletion() bool`

HasDeletion returns a boolean if a field has been set.

### GetFeatureIds

`func (o *BTBaseEntityData33) GetFeatureIds() []string`

GetFeatureIds returns the FeatureIds field if non-nil, zero value otherwise.

### GetFeatureIdsOk

`func (o *BTBaseEntityData33) GetFeatureIdsOk() (*[]string, bool)`

GetFeatureIdsOk returns a tuple with the FeatureIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureIds

`func (o *BTBaseEntityData33) SetFeatureIds(v []string)`

SetFeatureIds sets FeatureIds field to given value.

### HasFeatureIds

`func (o *BTBaseEntityData33) HasFeatureIds() bool`

HasFeatureIds returns a boolean if a field has been set.

### GetFromSketch

`func (o *BTBaseEntityData33) GetFromSketch() bool`

GetFromSketch returns the FromSketch field if non-nil, zero value otherwise.

### GetFromSketchOk

`func (o *BTBaseEntityData33) GetFromSketchOk() (*bool, bool)`

GetFromSketchOk returns a tuple with the FromSketch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromSketch

`func (o *BTBaseEntityData33) SetFromSketch(v bool)`

SetFromSketch sets FromSketch field to given value.

### HasFromSketch

`func (o *BTBaseEntityData33) HasFromSketch() bool`

HasFromSketch returns a boolean if a field has been set.

### GetGeometries

`func (o *BTBaseEntityData33) GetGeometries() []BTEntityGeometry35`

GetGeometries returns the Geometries field if non-nil, zero value otherwise.

### GetGeometriesOk

`func (o *BTBaseEntityData33) GetGeometriesOk() (*[]BTEntityGeometry35, bool)`

GetGeometriesOk returns a tuple with the Geometries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometries

`func (o *BTBaseEntityData33) SetGeometries(v []BTEntityGeometry35)`

SetGeometries sets Geometries field to given value.

### HasGeometries

`func (o *BTBaseEntityData33) HasGeometries() bool`

HasGeometries returns a boolean if a field has been set.

### GetOrigin

`func (o *BTBaseEntityData33) GetOrigin() bool`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BTBaseEntityData33) GetOriginOk() (*bool, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BTBaseEntityData33) SetOrigin(v bool)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *BTBaseEntityData33) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


