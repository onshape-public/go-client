# BTExportModelFace1363

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppearancePropertyNodeId** | Pointer to **string** | Identifies the application of the appearance. Faces that share a value were assigned an appearance together. | [optional] 
**Area** | Pointer to **float64** |  | [optional] 
**Box** | Pointer to [**BTBoundingBox1052**](BTBoundingBox1052.md) |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**DecalIdToDecal** | Pointer to [**map[string]BTDecal2404**](BTDecal2404.md) |  | [optional] 
**FaceProperties** | Pointer to [**BTExportModelProperties3216**](BTExportModelProperties3216.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Loops** | Pointer to [**[]BTExportModelLoop1182**](BTExportModelLoop1182.md) |  | [optional] 
**Orientation** | Pointer to **bool** |  | [optional] 
**Surface** | Pointer to [**BTSurfaceDescription1564**](BTSurfaceDescription1564.md) |  | [optional] 

## Methods

### NewBTExportModelFace1363

`func NewBTExportModelFace1363() *BTExportModelFace1363`

NewBTExportModelFace1363 instantiates a new BTExportModelFace1363 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTExportModelFace1363WithDefaults

`func NewBTExportModelFace1363WithDefaults() *BTExportModelFace1363`

NewBTExportModelFace1363WithDefaults instantiates a new BTExportModelFace1363 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearancePropertyNodeId

`func (o *BTExportModelFace1363) GetAppearancePropertyNodeId() string`

GetAppearancePropertyNodeId returns the AppearancePropertyNodeId field if non-nil, zero value otherwise.

### GetAppearancePropertyNodeIdOk

`func (o *BTExportModelFace1363) GetAppearancePropertyNodeIdOk() (*string, bool)`

GetAppearancePropertyNodeIdOk returns a tuple with the AppearancePropertyNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearancePropertyNodeId

`func (o *BTExportModelFace1363) SetAppearancePropertyNodeId(v string)`

SetAppearancePropertyNodeId sets AppearancePropertyNodeId field to given value.

### HasAppearancePropertyNodeId

`func (o *BTExportModelFace1363) HasAppearancePropertyNodeId() bool`

HasAppearancePropertyNodeId returns a boolean if a field has been set.

### GetArea

`func (o *BTExportModelFace1363) GetArea() float64`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *BTExportModelFace1363) GetAreaOk() (*float64, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *BTExportModelFace1363) SetArea(v float64)`

SetArea sets Area field to given value.

### HasArea

`func (o *BTExportModelFace1363) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetBox

`func (o *BTExportModelFace1363) GetBox() BTBoundingBox1052`

GetBox returns the Box field if non-nil, zero value otherwise.

### GetBoxOk

`func (o *BTExportModelFace1363) GetBoxOk() (*BTBoundingBox1052, bool)`

GetBoxOk returns a tuple with the Box field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBox

`func (o *BTExportModelFace1363) SetBox(v BTBoundingBox1052)`

SetBox sets Box field to given value.

### HasBox

`func (o *BTExportModelFace1363) HasBox() bool`

HasBox returns a boolean if a field has been set.

### GetBtType

`func (o *BTExportModelFace1363) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTExportModelFace1363) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTExportModelFace1363) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTExportModelFace1363) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDecalIdToDecal

`func (o *BTExportModelFace1363) GetDecalIdToDecal() map[string]BTDecal2404`

GetDecalIdToDecal returns the DecalIdToDecal field if non-nil, zero value otherwise.

### GetDecalIdToDecalOk

`func (o *BTExportModelFace1363) GetDecalIdToDecalOk() (*map[string]BTDecal2404, bool)`

GetDecalIdToDecalOk returns a tuple with the DecalIdToDecal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecalIdToDecal

`func (o *BTExportModelFace1363) SetDecalIdToDecal(v map[string]BTDecal2404)`

SetDecalIdToDecal sets DecalIdToDecal field to given value.

### HasDecalIdToDecal

`func (o *BTExportModelFace1363) HasDecalIdToDecal() bool`

HasDecalIdToDecal returns a boolean if a field has been set.

### GetFaceProperties

`func (o *BTExportModelFace1363) GetFaceProperties() BTExportModelProperties3216`

GetFaceProperties returns the FaceProperties field if non-nil, zero value otherwise.

### GetFacePropertiesOk

`func (o *BTExportModelFace1363) GetFacePropertiesOk() (*BTExportModelProperties3216, bool)`

GetFacePropertiesOk returns a tuple with the FaceProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceProperties

`func (o *BTExportModelFace1363) SetFaceProperties(v BTExportModelProperties3216)`

SetFaceProperties sets FaceProperties field to given value.

### HasFaceProperties

`func (o *BTExportModelFace1363) HasFaceProperties() bool`

HasFaceProperties returns a boolean if a field has been set.

### GetId

`func (o *BTExportModelFace1363) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTExportModelFace1363) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTExportModelFace1363) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTExportModelFace1363) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoops

`func (o *BTExportModelFace1363) GetLoops() []BTExportModelLoop1182`

GetLoops returns the Loops field if non-nil, zero value otherwise.

### GetLoopsOk

`func (o *BTExportModelFace1363) GetLoopsOk() (*[]BTExportModelLoop1182, bool)`

GetLoopsOk returns a tuple with the Loops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoops

`func (o *BTExportModelFace1363) SetLoops(v []BTExportModelLoop1182)`

SetLoops sets Loops field to given value.

### HasLoops

`func (o *BTExportModelFace1363) HasLoops() bool`

HasLoops returns a boolean if a field has been set.

### GetOrientation

`func (o *BTExportModelFace1363) GetOrientation() bool`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *BTExportModelFace1363) GetOrientationOk() (*bool, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *BTExportModelFace1363) SetOrientation(v bool)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *BTExportModelFace1363) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetSurface

`func (o *BTExportModelFace1363) GetSurface() BTSurfaceDescription1564`

GetSurface returns the Surface field if non-nil, zero value otherwise.

### GetSurfaceOk

`func (o *BTExportModelFace1363) GetSurfaceOk() (*BTSurfaceDescription1564, bool)`

GetSurfaceOk returns a tuple with the Surface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurface

`func (o *BTExportModelFace1363) SetSurface(v BTSurfaceDescription1564)`

SetSurface sets Surface field to given value.

### HasSurface

`func (o *BTExportModelFace1363) HasSurface() bool`

HasSurface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


