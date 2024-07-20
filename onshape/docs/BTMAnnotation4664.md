# BTMAnnotation4664

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**AnnotationType** | Pointer to [**GBTAnnotationType**](GBTAnnotationType.md) |  | [optional] 
**Parameters** | Pointer to [**[]BTMParameter1**](BTMParameter1.md) |  | [optional] 

## Methods

### NewBTMAnnotation4664

`func NewBTMAnnotation4664() *BTMAnnotation4664`

NewBTMAnnotation4664 instantiates a new BTMAnnotation4664 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTMAnnotation4664WithDefaults

`func NewBTMAnnotation4664WithDefaults() *BTMAnnotation4664`

NewBTMAnnotation4664WithDefaults instantiates a new BTMAnnotation4664 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTMAnnotation4664) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTMAnnotation4664) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTMAnnotation4664) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTMAnnotation4664) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetAnnotationType

`func (o *BTMAnnotation4664) GetAnnotationType() GBTAnnotationType`

GetAnnotationType returns the AnnotationType field if non-nil, zero value otherwise.

### GetAnnotationTypeOk

`func (o *BTMAnnotation4664) GetAnnotationTypeOk() (*GBTAnnotationType, bool)`

GetAnnotationTypeOk returns a tuple with the AnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationType

`func (o *BTMAnnotation4664) SetAnnotationType(v GBTAnnotationType)`

SetAnnotationType sets AnnotationType field to given value.

### HasAnnotationType

`func (o *BTMAnnotation4664) HasAnnotationType() bool`

HasAnnotationType returns a boolean if a field has been set.

### GetParameters

`func (o *BTMAnnotation4664) GetParameters() []BTMParameter1`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *BTMAnnotation4664) GetParametersOk() (*[]BTMParameter1, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *BTMAnnotation4664) SetParameters(v []BTMParameter1)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *BTMAnnotation4664) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


