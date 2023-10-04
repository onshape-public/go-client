# BTViewDataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Angle** | Pointer to **float64** |  | [optional] 
**CameraViewport** | Pointer to **[]float64** |  | [optional] 
**IsPerspective** | Pointer to **bool** |  | [optional] 
**ViewMatrix** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewBTViewDataInfo

`func NewBTViewDataInfo() *BTViewDataInfo`

NewBTViewDataInfo instantiates a new BTViewDataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTViewDataInfoWithDefaults

`func NewBTViewDataInfoWithDefaults() *BTViewDataInfo`

NewBTViewDataInfoWithDefaults instantiates a new BTViewDataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAngle

`func (o *BTViewDataInfo) GetAngle() float64`

GetAngle returns the Angle field if non-nil, zero value otherwise.

### GetAngleOk

`func (o *BTViewDataInfo) GetAngleOk() (*float64, bool)`

GetAngleOk returns a tuple with the Angle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngle

`func (o *BTViewDataInfo) SetAngle(v float64)`

SetAngle sets Angle field to given value.

### HasAngle

`func (o *BTViewDataInfo) HasAngle() bool`

HasAngle returns a boolean if a field has been set.

### GetCameraViewport

`func (o *BTViewDataInfo) GetCameraViewport() []float64`

GetCameraViewport returns the CameraViewport field if non-nil, zero value otherwise.

### GetCameraViewportOk

`func (o *BTViewDataInfo) GetCameraViewportOk() (*[]float64, bool)`

GetCameraViewportOk returns a tuple with the CameraViewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraViewport

`func (o *BTViewDataInfo) SetCameraViewport(v []float64)`

SetCameraViewport sets CameraViewport field to given value.

### HasCameraViewport

`func (o *BTViewDataInfo) HasCameraViewport() bool`

HasCameraViewport returns a boolean if a field has been set.

### GetIsPerspective

`func (o *BTViewDataInfo) GetIsPerspective() bool`

GetIsPerspective returns the IsPerspective field if non-nil, zero value otherwise.

### GetIsPerspectiveOk

`func (o *BTViewDataInfo) GetIsPerspectiveOk() (*bool, bool)`

GetIsPerspectiveOk returns a tuple with the IsPerspective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerspective

`func (o *BTViewDataInfo) SetIsPerspective(v bool)`

SetIsPerspective sets IsPerspective field to given value.

### HasIsPerspective

`func (o *BTViewDataInfo) HasIsPerspective() bool`

HasIsPerspective returns a boolean if a field has been set.

### GetViewMatrix

`func (o *BTViewDataInfo) GetViewMatrix() []float64`

GetViewMatrix returns the ViewMatrix field if non-nil, zero value otherwise.

### GetViewMatrixOk

`func (o *BTViewDataInfo) GetViewMatrixOk() (*[]float64, bool)`

GetViewMatrixOk returns a tuple with the ViewMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatrix

`func (o *BTViewDataInfo) SetViewMatrix(v []float64)`

SetViewMatrix sets ViewMatrix field to given value.

### HasViewMatrix

`func (o *BTViewDataInfo) HasViewMatrix() bool`

HasViewMatrix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


