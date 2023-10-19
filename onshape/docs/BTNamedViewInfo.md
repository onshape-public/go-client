# BTNamedViewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Angle** | Pointer to **float64** |  | [optional] 
**CameraViewport** | Pointer to **[]float64** |  | [optional] 
**Perspective** | Pointer to **bool** |  | [optional] 
**SectionPlanes** | Pointer to [**[]BTSectionPlaneInfo**](BTSectionPlaneInfo.md) |  | [optional] 
**SectionViewData** | Pointer to [**BTGraphicsSectionViewStateData4379**](BTGraphicsSectionViewStateData4379.md) |  | [optional] 
**ViewMatrix** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewBTNamedViewInfo

`func NewBTNamedViewInfo() *BTNamedViewInfo`

NewBTNamedViewInfo instantiates a new BTNamedViewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTNamedViewInfoWithDefaults

`func NewBTNamedViewInfoWithDefaults() *BTNamedViewInfo`

NewBTNamedViewInfoWithDefaults instantiates a new BTNamedViewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAngle

`func (o *BTNamedViewInfo) GetAngle() float64`

GetAngle returns the Angle field if non-nil, zero value otherwise.

### GetAngleOk

`func (o *BTNamedViewInfo) GetAngleOk() (*float64, bool)`

GetAngleOk returns a tuple with the Angle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngle

`func (o *BTNamedViewInfo) SetAngle(v float64)`

SetAngle sets Angle field to given value.

### HasAngle

`func (o *BTNamedViewInfo) HasAngle() bool`

HasAngle returns a boolean if a field has been set.

### GetCameraViewport

`func (o *BTNamedViewInfo) GetCameraViewport() []float64`

GetCameraViewport returns the CameraViewport field if non-nil, zero value otherwise.

### GetCameraViewportOk

`func (o *BTNamedViewInfo) GetCameraViewportOk() (*[]float64, bool)`

GetCameraViewportOk returns a tuple with the CameraViewport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraViewport

`func (o *BTNamedViewInfo) SetCameraViewport(v []float64)`

SetCameraViewport sets CameraViewport field to given value.

### HasCameraViewport

`func (o *BTNamedViewInfo) HasCameraViewport() bool`

HasCameraViewport returns a boolean if a field has been set.

### GetPerspective

`func (o *BTNamedViewInfo) GetPerspective() bool`

GetPerspective returns the Perspective field if non-nil, zero value otherwise.

### GetPerspectiveOk

`func (o *BTNamedViewInfo) GetPerspectiveOk() (*bool, bool)`

GetPerspectiveOk returns a tuple with the Perspective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerspective

`func (o *BTNamedViewInfo) SetPerspective(v bool)`

SetPerspective sets Perspective field to given value.

### HasPerspective

`func (o *BTNamedViewInfo) HasPerspective() bool`

HasPerspective returns a boolean if a field has been set.

### GetSectionPlanes

`func (o *BTNamedViewInfo) GetSectionPlanes() []BTSectionPlaneInfo`

GetSectionPlanes returns the SectionPlanes field if non-nil, zero value otherwise.

### GetSectionPlanesOk

`func (o *BTNamedViewInfo) GetSectionPlanesOk() (*[]BTSectionPlaneInfo, bool)`

GetSectionPlanesOk returns a tuple with the SectionPlanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionPlanes

`func (o *BTNamedViewInfo) SetSectionPlanes(v []BTSectionPlaneInfo)`

SetSectionPlanes sets SectionPlanes field to given value.

### HasSectionPlanes

`func (o *BTNamedViewInfo) HasSectionPlanes() bool`

HasSectionPlanes returns a boolean if a field has been set.

### GetSectionViewData

`func (o *BTNamedViewInfo) GetSectionViewData() BTGraphicsSectionViewStateData4379`

GetSectionViewData returns the SectionViewData field if non-nil, zero value otherwise.

### GetSectionViewDataOk

`func (o *BTNamedViewInfo) GetSectionViewDataOk() (*BTGraphicsSectionViewStateData4379, bool)`

GetSectionViewDataOk returns a tuple with the SectionViewData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionViewData

`func (o *BTNamedViewInfo) SetSectionViewData(v BTGraphicsSectionViewStateData4379)`

SetSectionViewData sets SectionViewData field to given value.

### HasSectionViewData

`func (o *BTNamedViewInfo) HasSectionViewData() bool`

HasSectionViewData returns a boolean if a field has been set.

### GetViewMatrix

`func (o *BTNamedViewInfo) GetViewMatrix() []float64`

GetViewMatrix returns the ViewMatrix field if non-nil, zero value otherwise.

### GetViewMatrixOk

`func (o *BTNamedViewInfo) GetViewMatrixOk() (*[]float64, bool)`

GetViewMatrixOk returns a tuple with the ViewMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatrix

`func (o *BTNamedViewInfo) SetViewMatrix(v []float64)`

SetViewMatrix sets ViewMatrix field to given value.

### HasViewMatrix

`func (o *BTNamedViewInfo) HasViewMatrix() bool`

HasViewMatrix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


