# BTAnnotationDisplayData3225

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotationId** | Pointer to **string** |  | [optional] 
**AnnotationPlane** | Pointer to [**BTCoordinateSystem387**](BTCoordinateSystem387.md) |  | [optional] 
**BasePlane** | Pointer to [**BTCoordinateSystem387**](BTCoordinateSystem387.md) |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**CharacteristicId** | Pointer to **string** |  | [optional] 
**DeterministicId** | Pointer to **string** |  | [optional] 
**DxdySegments** | Pointer to [**[]BTVector2d1812**](BTVector2d1812.md) |  | [optional] 
**IsConstrainedToPlane** | Pointer to **bool** |  | [optional] 
**IsDeletion** | Pointer to **bool** |  | [optional] 
**IsDerived** | Pointer to **bool** |  | [optional] 
**MainConstraintId** | Pointer to **string** |  | [optional] 
**MainFeatureId** | Pointer to **string** |  | [optional] 
**MainParameterId** | Pointer to **string** |  | [optional] 
**MainPartId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAnnotationDisplayData3225

`func NewBTAnnotationDisplayData3225() *BTAnnotationDisplayData3225`

NewBTAnnotationDisplayData3225 instantiates a new BTAnnotationDisplayData3225 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAnnotationDisplayData3225WithDefaults

`func NewBTAnnotationDisplayData3225WithDefaults() *BTAnnotationDisplayData3225`

NewBTAnnotationDisplayData3225WithDefaults instantiates a new BTAnnotationDisplayData3225 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotationId

`func (o *BTAnnotationDisplayData3225) GetAnnotationId() string`

GetAnnotationId returns the AnnotationId field if non-nil, zero value otherwise.

### GetAnnotationIdOk

`func (o *BTAnnotationDisplayData3225) GetAnnotationIdOk() (*string, bool)`

GetAnnotationIdOk returns a tuple with the AnnotationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationId

`func (o *BTAnnotationDisplayData3225) SetAnnotationId(v string)`

SetAnnotationId sets AnnotationId field to given value.

### HasAnnotationId

`func (o *BTAnnotationDisplayData3225) HasAnnotationId() bool`

HasAnnotationId returns a boolean if a field has been set.

### GetAnnotationPlane

`func (o *BTAnnotationDisplayData3225) GetAnnotationPlane() BTCoordinateSystem387`

GetAnnotationPlane returns the AnnotationPlane field if non-nil, zero value otherwise.

### GetAnnotationPlaneOk

`func (o *BTAnnotationDisplayData3225) GetAnnotationPlaneOk() (*BTCoordinateSystem387, bool)`

GetAnnotationPlaneOk returns a tuple with the AnnotationPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationPlane

`func (o *BTAnnotationDisplayData3225) SetAnnotationPlane(v BTCoordinateSystem387)`

SetAnnotationPlane sets AnnotationPlane field to given value.

### HasAnnotationPlane

`func (o *BTAnnotationDisplayData3225) HasAnnotationPlane() bool`

HasAnnotationPlane returns a boolean if a field has been set.

### GetBasePlane

`func (o *BTAnnotationDisplayData3225) GetBasePlane() BTCoordinateSystem387`

GetBasePlane returns the BasePlane field if non-nil, zero value otherwise.

### GetBasePlaneOk

`func (o *BTAnnotationDisplayData3225) GetBasePlaneOk() (*BTCoordinateSystem387, bool)`

GetBasePlaneOk returns a tuple with the BasePlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePlane

`func (o *BTAnnotationDisplayData3225) SetBasePlane(v BTCoordinateSystem387)`

SetBasePlane sets BasePlane field to given value.

### HasBasePlane

`func (o *BTAnnotationDisplayData3225) HasBasePlane() bool`

HasBasePlane returns a boolean if a field has been set.

### GetBtType

`func (o *BTAnnotationDisplayData3225) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTAnnotationDisplayData3225) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTAnnotationDisplayData3225) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTAnnotationDisplayData3225) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCharacteristicId

`func (o *BTAnnotationDisplayData3225) GetCharacteristicId() string`

GetCharacteristicId returns the CharacteristicId field if non-nil, zero value otherwise.

### GetCharacteristicIdOk

`func (o *BTAnnotationDisplayData3225) GetCharacteristicIdOk() (*string, bool)`

GetCharacteristicIdOk returns a tuple with the CharacteristicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacteristicId

`func (o *BTAnnotationDisplayData3225) SetCharacteristicId(v string)`

SetCharacteristicId sets CharacteristicId field to given value.

### HasCharacteristicId

`func (o *BTAnnotationDisplayData3225) HasCharacteristicId() bool`

HasCharacteristicId returns a boolean if a field has been set.

### GetDeterministicId

`func (o *BTAnnotationDisplayData3225) GetDeterministicId() string`

GetDeterministicId returns the DeterministicId field if non-nil, zero value otherwise.

### GetDeterministicIdOk

`func (o *BTAnnotationDisplayData3225) GetDeterministicIdOk() (*string, bool)`

GetDeterministicIdOk returns a tuple with the DeterministicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicId

`func (o *BTAnnotationDisplayData3225) SetDeterministicId(v string)`

SetDeterministicId sets DeterministicId field to given value.

### HasDeterministicId

`func (o *BTAnnotationDisplayData3225) HasDeterministicId() bool`

HasDeterministicId returns a boolean if a field has been set.

### GetDxdySegments

`func (o *BTAnnotationDisplayData3225) GetDxdySegments() []BTVector2d1812`

GetDxdySegments returns the DxdySegments field if non-nil, zero value otherwise.

### GetDxdySegmentsOk

`func (o *BTAnnotationDisplayData3225) GetDxdySegmentsOk() (*[]BTVector2d1812, bool)`

GetDxdySegmentsOk returns a tuple with the DxdySegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDxdySegments

`func (o *BTAnnotationDisplayData3225) SetDxdySegments(v []BTVector2d1812)`

SetDxdySegments sets DxdySegments field to given value.

### HasDxdySegments

`func (o *BTAnnotationDisplayData3225) HasDxdySegments() bool`

HasDxdySegments returns a boolean if a field has been set.

### GetIsConstrainedToPlane

`func (o *BTAnnotationDisplayData3225) GetIsConstrainedToPlane() bool`

GetIsConstrainedToPlane returns the IsConstrainedToPlane field if non-nil, zero value otherwise.

### GetIsConstrainedToPlaneOk

`func (o *BTAnnotationDisplayData3225) GetIsConstrainedToPlaneOk() (*bool, bool)`

GetIsConstrainedToPlaneOk returns a tuple with the IsConstrainedToPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConstrainedToPlane

`func (o *BTAnnotationDisplayData3225) SetIsConstrainedToPlane(v bool)`

SetIsConstrainedToPlane sets IsConstrainedToPlane field to given value.

### HasIsConstrainedToPlane

`func (o *BTAnnotationDisplayData3225) HasIsConstrainedToPlane() bool`

HasIsConstrainedToPlane returns a boolean if a field has been set.

### GetIsDeletion

`func (o *BTAnnotationDisplayData3225) GetIsDeletion() bool`

GetIsDeletion returns the IsDeletion field if non-nil, zero value otherwise.

### GetIsDeletionOk

`func (o *BTAnnotationDisplayData3225) GetIsDeletionOk() (*bool, bool)`

GetIsDeletionOk returns a tuple with the IsDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletion

`func (o *BTAnnotationDisplayData3225) SetIsDeletion(v bool)`

SetIsDeletion sets IsDeletion field to given value.

### HasIsDeletion

`func (o *BTAnnotationDisplayData3225) HasIsDeletion() bool`

HasIsDeletion returns a boolean if a field has been set.

### GetIsDerived

`func (o *BTAnnotationDisplayData3225) GetIsDerived() bool`

GetIsDerived returns the IsDerived field if non-nil, zero value otherwise.

### GetIsDerivedOk

`func (o *BTAnnotationDisplayData3225) GetIsDerivedOk() (*bool, bool)`

GetIsDerivedOk returns a tuple with the IsDerived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDerived

`func (o *BTAnnotationDisplayData3225) SetIsDerived(v bool)`

SetIsDerived sets IsDerived field to given value.

### HasIsDerived

`func (o *BTAnnotationDisplayData3225) HasIsDerived() bool`

HasIsDerived returns a boolean if a field has been set.

### GetMainConstraintId

`func (o *BTAnnotationDisplayData3225) GetMainConstraintId() string`

GetMainConstraintId returns the MainConstraintId field if non-nil, zero value otherwise.

### GetMainConstraintIdOk

`func (o *BTAnnotationDisplayData3225) GetMainConstraintIdOk() (*string, bool)`

GetMainConstraintIdOk returns a tuple with the MainConstraintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainConstraintId

`func (o *BTAnnotationDisplayData3225) SetMainConstraintId(v string)`

SetMainConstraintId sets MainConstraintId field to given value.

### HasMainConstraintId

`func (o *BTAnnotationDisplayData3225) HasMainConstraintId() bool`

HasMainConstraintId returns a boolean if a field has been set.

### GetMainFeatureId

`func (o *BTAnnotationDisplayData3225) GetMainFeatureId() string`

GetMainFeatureId returns the MainFeatureId field if non-nil, zero value otherwise.

### GetMainFeatureIdOk

`func (o *BTAnnotationDisplayData3225) GetMainFeatureIdOk() (*string, bool)`

GetMainFeatureIdOk returns a tuple with the MainFeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainFeatureId

`func (o *BTAnnotationDisplayData3225) SetMainFeatureId(v string)`

SetMainFeatureId sets MainFeatureId field to given value.

### HasMainFeatureId

`func (o *BTAnnotationDisplayData3225) HasMainFeatureId() bool`

HasMainFeatureId returns a boolean if a field has been set.

### GetMainParameterId

`func (o *BTAnnotationDisplayData3225) GetMainParameterId() string`

GetMainParameterId returns the MainParameterId field if non-nil, zero value otherwise.

### GetMainParameterIdOk

`func (o *BTAnnotationDisplayData3225) GetMainParameterIdOk() (*string, bool)`

GetMainParameterIdOk returns a tuple with the MainParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainParameterId

`func (o *BTAnnotationDisplayData3225) SetMainParameterId(v string)`

SetMainParameterId sets MainParameterId field to given value.

### HasMainParameterId

`func (o *BTAnnotationDisplayData3225) HasMainParameterId() bool`

HasMainParameterId returns a boolean if a field has been set.

### GetMainPartId

`func (o *BTAnnotationDisplayData3225) GetMainPartId() string`

GetMainPartId returns the MainPartId field if non-nil, zero value otherwise.

### GetMainPartIdOk

`func (o *BTAnnotationDisplayData3225) GetMainPartIdOk() (*string, bool)`

GetMainPartIdOk returns a tuple with the MainPartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainPartId

`func (o *BTAnnotationDisplayData3225) SetMainPartId(v string)`

SetMainPartId sets MainPartId field to given value.

### HasMainPartId

`func (o *BTAnnotationDisplayData3225) HasMainPartId() bool`

HasMainPartId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


