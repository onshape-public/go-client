# Sampler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to [**AccessorModel**](AccessorModel.md) |  | [optional] 
**Interpolation** | Pointer to [**Interpolation**](Interpolation.md) |  | [optional] 
**Output** | Pointer to [**AccessorModel**](AccessorModel.md) |  | [optional] 

## Methods

### NewSampler

`func NewSampler() *Sampler`

NewSampler instantiates a new Sampler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamplerWithDefaults

`func NewSamplerWithDefaults() *Sampler`

NewSamplerWithDefaults instantiates a new Sampler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *Sampler) GetInput() AccessorModel`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Sampler) GetInputOk() (*AccessorModel, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Sampler) SetInput(v AccessorModel)`

SetInput sets Input field to given value.

### HasInput

`func (o *Sampler) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetInterpolation

`func (o *Sampler) GetInterpolation() Interpolation`

GetInterpolation returns the Interpolation field if non-nil, zero value otherwise.

### GetInterpolationOk

`func (o *Sampler) GetInterpolationOk() (*Interpolation, bool)`

GetInterpolationOk returns a tuple with the Interpolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolation

`func (o *Sampler) SetInterpolation(v Interpolation)`

SetInterpolation sets Interpolation field to given value.

### HasInterpolation

`func (o *Sampler) HasInterpolation() bool`

HasInterpolation returns a boolean if a field has been set.

### GetOutput

`func (o *Sampler) GetOutput() AccessorModel`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Sampler) GetOutputOk() (*AccessorModel, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Sampler) SetOutput(v AccessorModel)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Sampler) HasOutput() bool`

HasOutput returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


