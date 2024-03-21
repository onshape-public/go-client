# BTDebugGeometry2059

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Appearance** | Pointer to [**BTGraphicsAppearance1152**](BTGraphicsAppearance1152.md) |  | [optional] 
**BelongsToFlattenedSheetMetalBody** | Pointer to **bool** |  | [optional] 
**BodyId** | Pointer to **string** |  | [optional] 
**Color** | Pointer to [**GBTDebugEntityColor**](GBTDebugEntityColor.md) |  | [optional] 
**DeterministicId** | Pointer to **string** |  | [optional] 
**SheetMetalModelId** | Pointer to **string** |  | [optional] 
**Style** | Pointer to [**GBTDebugEntityStyle**](GBTDebugEntityStyle.md) |  | [optional] 
**Tessellation** | Pointer to [**BTTessellatedGeometry2576**](BTTessellatedGeometry2576.md) |  | [optional] 

## Methods

### NewBTDebugGeometry2059

`func NewBTDebugGeometry2059() *BTDebugGeometry2059`

NewBTDebugGeometry2059 instantiates a new BTDebugGeometry2059 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDebugGeometry2059WithDefaults

`func NewBTDebugGeometry2059WithDefaults() *BTDebugGeometry2059`

NewBTDebugGeometry2059WithDefaults instantiates a new BTDebugGeometry2059 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTDebugGeometry2059) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTDebugGeometry2059) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTDebugGeometry2059) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTDebugGeometry2059) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetAppearance

`func (o *BTDebugGeometry2059) GetAppearance() BTGraphicsAppearance1152`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *BTDebugGeometry2059) GetAppearanceOk() (*BTGraphicsAppearance1152, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *BTDebugGeometry2059) SetAppearance(v BTGraphicsAppearance1152)`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *BTDebugGeometry2059) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### GetBelongsToFlattenedSheetMetalBody

`func (o *BTDebugGeometry2059) GetBelongsToFlattenedSheetMetalBody() bool`

GetBelongsToFlattenedSheetMetalBody returns the BelongsToFlattenedSheetMetalBody field if non-nil, zero value otherwise.

### GetBelongsToFlattenedSheetMetalBodyOk

`func (o *BTDebugGeometry2059) GetBelongsToFlattenedSheetMetalBodyOk() (*bool, bool)`

GetBelongsToFlattenedSheetMetalBodyOk returns a tuple with the BelongsToFlattenedSheetMetalBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelongsToFlattenedSheetMetalBody

`func (o *BTDebugGeometry2059) SetBelongsToFlattenedSheetMetalBody(v bool)`

SetBelongsToFlattenedSheetMetalBody sets BelongsToFlattenedSheetMetalBody field to given value.

### HasBelongsToFlattenedSheetMetalBody

`func (o *BTDebugGeometry2059) HasBelongsToFlattenedSheetMetalBody() bool`

HasBelongsToFlattenedSheetMetalBody returns a boolean if a field has been set.

### GetBodyId

`func (o *BTDebugGeometry2059) GetBodyId() string`

GetBodyId returns the BodyId field if non-nil, zero value otherwise.

### GetBodyIdOk

`func (o *BTDebugGeometry2059) GetBodyIdOk() (*string, bool)`

GetBodyIdOk returns a tuple with the BodyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyId

`func (o *BTDebugGeometry2059) SetBodyId(v string)`

SetBodyId sets BodyId field to given value.

### HasBodyId

`func (o *BTDebugGeometry2059) HasBodyId() bool`

HasBodyId returns a boolean if a field has been set.

### GetColor

`func (o *BTDebugGeometry2059) GetColor() GBTDebugEntityColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BTDebugGeometry2059) GetColorOk() (*GBTDebugEntityColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BTDebugGeometry2059) SetColor(v GBTDebugEntityColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *BTDebugGeometry2059) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDeterministicId

`func (o *BTDebugGeometry2059) GetDeterministicId() string`

GetDeterministicId returns the DeterministicId field if non-nil, zero value otherwise.

### GetDeterministicIdOk

`func (o *BTDebugGeometry2059) GetDeterministicIdOk() (*string, bool)`

GetDeterministicIdOk returns a tuple with the DeterministicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeterministicId

`func (o *BTDebugGeometry2059) SetDeterministicId(v string)`

SetDeterministicId sets DeterministicId field to given value.

### HasDeterministicId

`func (o *BTDebugGeometry2059) HasDeterministicId() bool`

HasDeterministicId returns a boolean if a field has been set.

### GetSheetMetalModelId

`func (o *BTDebugGeometry2059) GetSheetMetalModelId() string`

GetSheetMetalModelId returns the SheetMetalModelId field if non-nil, zero value otherwise.

### GetSheetMetalModelIdOk

`func (o *BTDebugGeometry2059) GetSheetMetalModelIdOk() (*string, bool)`

GetSheetMetalModelIdOk returns a tuple with the SheetMetalModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetMetalModelId

`func (o *BTDebugGeometry2059) SetSheetMetalModelId(v string)`

SetSheetMetalModelId sets SheetMetalModelId field to given value.

### HasSheetMetalModelId

`func (o *BTDebugGeometry2059) HasSheetMetalModelId() bool`

HasSheetMetalModelId returns a boolean if a field has been set.

### GetStyle

`func (o *BTDebugGeometry2059) GetStyle() GBTDebugEntityStyle`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *BTDebugGeometry2059) GetStyleOk() (*GBTDebugEntityStyle, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *BTDebugGeometry2059) SetStyle(v GBTDebugEntityStyle)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *BTDebugGeometry2059) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetTessellation

`func (o *BTDebugGeometry2059) GetTessellation() BTTessellatedGeometry2576`

GetTessellation returns the Tessellation field if non-nil, zero value otherwise.

### GetTessellationOk

`func (o *BTDebugGeometry2059) GetTessellationOk() (*BTTessellatedGeometry2576, bool)`

GetTessellationOk returns a tuple with the Tessellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTessellation

`func (o *BTDebugGeometry2059) SetTessellation(v BTTessellatedGeometry2576)`

SetTessellation sets Tessellation field to given value.

### HasTessellation

`func (o *BTDebugGeometry2059) HasTessellation() bool`

HasTessellation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


