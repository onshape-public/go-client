# BTDocumentElementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccelerationUnits** | Pointer to **string** |  | [optional] 
**AngleUnits** | Pointer to **string** |  | [optional] 
**AngularVelocityUnits** | Pointer to **string** |  | [optional] 
**ApplicationTarget** | Pointer to [**BTApplicationTargetInfo**](BTApplicationTargetInfo.md) |  | [optional] 
**AreaUnits** | Pointer to **string** |  | [optional] 
**DataType** | Pointer to **string** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**ElementType** | Pointer to **string** |  | [optional] 
**EnergyUnits** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**ForceUnits** | Pointer to **string** |  | [optional] 
**ForeignDataId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LengthUnits** | Pointer to **string** |  | [optional] 
**MassUnits** | Pointer to **string** |  | [optional] 
**MicroversionId** | Pointer to **string** |  | [optional] 
**MomentUnits** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PressureUnits** | Pointer to **string** |  | [optional] 
**PrettyType** | Pointer to **string** |  | [optional] 
**SafeToShow** | Pointer to **bool** |  | [optional] 
**SpecifiedUnit** | Pointer to **string** |  | [optional] 
**ThumbnailInfo** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**Thumbnails** | Pointer to **string** |  | [optional] 
**TimeUnits** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Unupdatable** | Pointer to **bool** |  | [optional] 
**VolumeUnits** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to [**BTZipFileInfo**](BTZipFileInfo.md) |  | [optional] 

## Methods

### NewBTDocumentElementInfo

`func NewBTDocumentElementInfo() *BTDocumentElementInfo`

NewBTDocumentElementInfo instantiates a new BTDocumentElementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentElementInfoWithDefaults

`func NewBTDocumentElementInfoWithDefaults() *BTDocumentElementInfo`

NewBTDocumentElementInfoWithDefaults instantiates a new BTDocumentElementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccelerationUnits

`func (o *BTDocumentElementInfo) GetAccelerationUnits() string`

GetAccelerationUnits returns the AccelerationUnits field if non-nil, zero value otherwise.

### GetAccelerationUnitsOk

`func (o *BTDocumentElementInfo) GetAccelerationUnitsOk() (*string, bool)`

GetAccelerationUnitsOk returns a tuple with the AccelerationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccelerationUnits

`func (o *BTDocumentElementInfo) SetAccelerationUnits(v string)`

SetAccelerationUnits sets AccelerationUnits field to given value.

### HasAccelerationUnits

`func (o *BTDocumentElementInfo) HasAccelerationUnits() bool`

HasAccelerationUnits returns a boolean if a field has been set.

### GetAngleUnits

`func (o *BTDocumentElementInfo) GetAngleUnits() string`

GetAngleUnits returns the AngleUnits field if non-nil, zero value otherwise.

### GetAngleUnitsOk

`func (o *BTDocumentElementInfo) GetAngleUnitsOk() (*string, bool)`

GetAngleUnitsOk returns a tuple with the AngleUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngleUnits

`func (o *BTDocumentElementInfo) SetAngleUnits(v string)`

SetAngleUnits sets AngleUnits field to given value.

### HasAngleUnits

`func (o *BTDocumentElementInfo) HasAngleUnits() bool`

HasAngleUnits returns a boolean if a field has been set.

### GetAngularVelocityUnits

`func (o *BTDocumentElementInfo) GetAngularVelocityUnits() string`

GetAngularVelocityUnits returns the AngularVelocityUnits field if non-nil, zero value otherwise.

### GetAngularVelocityUnitsOk

`func (o *BTDocumentElementInfo) GetAngularVelocityUnitsOk() (*string, bool)`

GetAngularVelocityUnitsOk returns a tuple with the AngularVelocityUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngularVelocityUnits

`func (o *BTDocumentElementInfo) SetAngularVelocityUnits(v string)`

SetAngularVelocityUnits sets AngularVelocityUnits field to given value.

### HasAngularVelocityUnits

`func (o *BTDocumentElementInfo) HasAngularVelocityUnits() bool`

HasAngularVelocityUnits returns a boolean if a field has been set.

### GetApplicationTarget

`func (o *BTDocumentElementInfo) GetApplicationTarget() BTApplicationTargetInfo`

GetApplicationTarget returns the ApplicationTarget field if non-nil, zero value otherwise.

### GetApplicationTargetOk

`func (o *BTDocumentElementInfo) GetApplicationTargetOk() (*BTApplicationTargetInfo, bool)`

GetApplicationTargetOk returns a tuple with the ApplicationTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTarget

`func (o *BTDocumentElementInfo) SetApplicationTarget(v BTApplicationTargetInfo)`

SetApplicationTarget sets ApplicationTarget field to given value.

### HasApplicationTarget

`func (o *BTDocumentElementInfo) HasApplicationTarget() bool`

HasApplicationTarget returns a boolean if a field has been set.

### GetAreaUnits

`func (o *BTDocumentElementInfo) GetAreaUnits() string`

GetAreaUnits returns the AreaUnits field if non-nil, zero value otherwise.

### GetAreaUnitsOk

`func (o *BTDocumentElementInfo) GetAreaUnitsOk() (*string, bool)`

GetAreaUnitsOk returns a tuple with the AreaUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaUnits

`func (o *BTDocumentElementInfo) SetAreaUnits(v string)`

SetAreaUnits sets AreaUnits field to given value.

### HasAreaUnits

`func (o *BTDocumentElementInfo) HasAreaUnits() bool`

HasAreaUnits returns a boolean if a field has been set.

### GetDataType

`func (o *BTDocumentElementInfo) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *BTDocumentElementInfo) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *BTDocumentElementInfo) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *BTDocumentElementInfo) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetDeleted

`func (o *BTDocumentElementInfo) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *BTDocumentElementInfo) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *BTDocumentElementInfo) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *BTDocumentElementInfo) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetElementType

`func (o *BTDocumentElementInfo) GetElementType() string`

GetElementType returns the ElementType field if non-nil, zero value otherwise.

### GetElementTypeOk

`func (o *BTDocumentElementInfo) GetElementTypeOk() (*string, bool)`

GetElementTypeOk returns a tuple with the ElementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementType

`func (o *BTDocumentElementInfo) SetElementType(v string)`

SetElementType sets ElementType field to given value.

### HasElementType

`func (o *BTDocumentElementInfo) HasElementType() bool`

HasElementType returns a boolean if a field has been set.

### GetEnergyUnits

`func (o *BTDocumentElementInfo) GetEnergyUnits() string`

GetEnergyUnits returns the EnergyUnits field if non-nil, zero value otherwise.

### GetEnergyUnitsOk

`func (o *BTDocumentElementInfo) GetEnergyUnitsOk() (*string, bool)`

GetEnergyUnitsOk returns a tuple with the EnergyUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergyUnits

`func (o *BTDocumentElementInfo) SetEnergyUnits(v string)`

SetEnergyUnits sets EnergyUnits field to given value.

### HasEnergyUnits

`func (o *BTDocumentElementInfo) HasEnergyUnits() bool`

HasEnergyUnits returns a boolean if a field has been set.

### GetFilename

`func (o *BTDocumentElementInfo) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *BTDocumentElementInfo) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *BTDocumentElementInfo) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *BTDocumentElementInfo) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetForceUnits

`func (o *BTDocumentElementInfo) GetForceUnits() string`

GetForceUnits returns the ForceUnits field if non-nil, zero value otherwise.

### GetForceUnitsOk

`func (o *BTDocumentElementInfo) GetForceUnitsOk() (*string, bool)`

GetForceUnitsOk returns a tuple with the ForceUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnits

`func (o *BTDocumentElementInfo) SetForceUnits(v string)`

SetForceUnits sets ForceUnits field to given value.

### HasForceUnits

`func (o *BTDocumentElementInfo) HasForceUnits() bool`

HasForceUnits returns a boolean if a field has been set.

### GetForeignDataId

`func (o *BTDocumentElementInfo) GetForeignDataId() string`

GetForeignDataId returns the ForeignDataId field if non-nil, zero value otherwise.

### GetForeignDataIdOk

`func (o *BTDocumentElementInfo) GetForeignDataIdOk() (*string, bool)`

GetForeignDataIdOk returns a tuple with the ForeignDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignDataId

`func (o *BTDocumentElementInfo) SetForeignDataId(v string)`

SetForeignDataId sets ForeignDataId field to given value.

### HasForeignDataId

`func (o *BTDocumentElementInfo) HasForeignDataId() bool`

HasForeignDataId returns a boolean if a field has been set.

### GetId

`func (o *BTDocumentElementInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDocumentElementInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDocumentElementInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDocumentElementInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLengthUnits

`func (o *BTDocumentElementInfo) GetLengthUnits() string`

GetLengthUnits returns the LengthUnits field if non-nil, zero value otherwise.

### GetLengthUnitsOk

`func (o *BTDocumentElementInfo) GetLengthUnitsOk() (*string, bool)`

GetLengthUnitsOk returns a tuple with the LengthUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthUnits

`func (o *BTDocumentElementInfo) SetLengthUnits(v string)`

SetLengthUnits sets LengthUnits field to given value.

### HasLengthUnits

`func (o *BTDocumentElementInfo) HasLengthUnits() bool`

HasLengthUnits returns a boolean if a field has been set.

### GetMassUnits

`func (o *BTDocumentElementInfo) GetMassUnits() string`

GetMassUnits returns the MassUnits field if non-nil, zero value otherwise.

### GetMassUnitsOk

`func (o *BTDocumentElementInfo) GetMassUnitsOk() (*string, bool)`

GetMassUnitsOk returns a tuple with the MassUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMassUnits

`func (o *BTDocumentElementInfo) SetMassUnits(v string)`

SetMassUnits sets MassUnits field to given value.

### HasMassUnits

`func (o *BTDocumentElementInfo) HasMassUnits() bool`

HasMassUnits returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTDocumentElementInfo) GetMicroversionId() string`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTDocumentElementInfo) GetMicroversionIdOk() (*string, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTDocumentElementInfo) SetMicroversionId(v string)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTDocumentElementInfo) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetMomentUnits

`func (o *BTDocumentElementInfo) GetMomentUnits() string`

GetMomentUnits returns the MomentUnits field if non-nil, zero value otherwise.

### GetMomentUnitsOk

`func (o *BTDocumentElementInfo) GetMomentUnitsOk() (*string, bool)`

GetMomentUnitsOk returns a tuple with the MomentUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMomentUnits

`func (o *BTDocumentElementInfo) SetMomentUnits(v string)`

SetMomentUnits sets MomentUnits field to given value.

### HasMomentUnits

`func (o *BTDocumentElementInfo) HasMomentUnits() bool`

HasMomentUnits returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentElementInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentElementInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentElementInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentElementInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPressureUnits

`func (o *BTDocumentElementInfo) GetPressureUnits() string`

GetPressureUnits returns the PressureUnits field if non-nil, zero value otherwise.

### GetPressureUnitsOk

`func (o *BTDocumentElementInfo) GetPressureUnitsOk() (*string, bool)`

GetPressureUnitsOk returns a tuple with the PressureUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPressureUnits

`func (o *BTDocumentElementInfo) SetPressureUnits(v string)`

SetPressureUnits sets PressureUnits field to given value.

### HasPressureUnits

`func (o *BTDocumentElementInfo) HasPressureUnits() bool`

HasPressureUnits returns a boolean if a field has been set.

### GetPrettyType

`func (o *BTDocumentElementInfo) GetPrettyType() string`

GetPrettyType returns the PrettyType field if non-nil, zero value otherwise.

### GetPrettyTypeOk

`func (o *BTDocumentElementInfo) GetPrettyTypeOk() (*string, bool)`

GetPrettyTypeOk returns a tuple with the PrettyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrettyType

`func (o *BTDocumentElementInfo) SetPrettyType(v string)`

SetPrettyType sets PrettyType field to given value.

### HasPrettyType

`func (o *BTDocumentElementInfo) HasPrettyType() bool`

HasPrettyType returns a boolean if a field has been set.

### GetSafeToShow

`func (o *BTDocumentElementInfo) GetSafeToShow() bool`

GetSafeToShow returns the SafeToShow field if non-nil, zero value otherwise.

### GetSafeToShowOk

`func (o *BTDocumentElementInfo) GetSafeToShowOk() (*bool, bool)`

GetSafeToShowOk returns a tuple with the SafeToShow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeToShow

`func (o *BTDocumentElementInfo) SetSafeToShow(v bool)`

SetSafeToShow sets SafeToShow field to given value.

### HasSafeToShow

`func (o *BTDocumentElementInfo) HasSafeToShow() bool`

HasSafeToShow returns a boolean if a field has been set.

### GetSpecifiedUnit

`func (o *BTDocumentElementInfo) GetSpecifiedUnit() string`

GetSpecifiedUnit returns the SpecifiedUnit field if non-nil, zero value otherwise.

### GetSpecifiedUnitOk

`func (o *BTDocumentElementInfo) GetSpecifiedUnitOk() (*string, bool)`

GetSpecifiedUnitOk returns a tuple with the SpecifiedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedUnit

`func (o *BTDocumentElementInfo) SetSpecifiedUnit(v string)`

SetSpecifiedUnit sets SpecifiedUnit field to given value.

### HasSpecifiedUnit

`func (o *BTDocumentElementInfo) HasSpecifiedUnit() bool`

HasSpecifiedUnit returns a boolean if a field has been set.

### GetThumbnailInfo

`func (o *BTDocumentElementInfo) GetThumbnailInfo() BTThumbnailInfo`

GetThumbnailInfo returns the ThumbnailInfo field if non-nil, zero value otherwise.

### GetThumbnailInfoOk

`func (o *BTDocumentElementInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool)`

GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailInfo

`func (o *BTDocumentElementInfo) SetThumbnailInfo(v BTThumbnailInfo)`

SetThumbnailInfo sets ThumbnailInfo field to given value.

### HasThumbnailInfo

`func (o *BTDocumentElementInfo) HasThumbnailInfo() bool`

HasThumbnailInfo returns a boolean if a field has been set.

### GetThumbnails

`func (o *BTDocumentElementInfo) GetThumbnails() string`

GetThumbnails returns the Thumbnails field if non-nil, zero value otherwise.

### GetThumbnailsOk

`func (o *BTDocumentElementInfo) GetThumbnailsOk() (*string, bool)`

GetThumbnailsOk returns a tuple with the Thumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnails

`func (o *BTDocumentElementInfo) SetThumbnails(v string)`

SetThumbnails sets Thumbnails field to given value.

### HasThumbnails

`func (o *BTDocumentElementInfo) HasThumbnails() bool`

HasThumbnails returns a boolean if a field has been set.

### GetTimeUnits

`func (o *BTDocumentElementInfo) GetTimeUnits() string`

GetTimeUnits returns the TimeUnits field if non-nil, zero value otherwise.

### GetTimeUnitsOk

`func (o *BTDocumentElementInfo) GetTimeUnitsOk() (*string, bool)`

GetTimeUnitsOk returns a tuple with the TimeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnits

`func (o *BTDocumentElementInfo) SetTimeUnits(v string)`

SetTimeUnits sets TimeUnits field to given value.

### HasTimeUnits

`func (o *BTDocumentElementInfo) HasTimeUnits() bool`

HasTimeUnits returns a boolean if a field has been set.

### GetType

`func (o *BTDocumentElementInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTDocumentElementInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTDocumentElementInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTDocumentElementInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnupdatable

`func (o *BTDocumentElementInfo) GetUnupdatable() bool`

GetUnupdatable returns the Unupdatable field if non-nil, zero value otherwise.

### GetUnupdatableOk

`func (o *BTDocumentElementInfo) GetUnupdatableOk() (*bool, bool)`

GetUnupdatableOk returns a tuple with the Unupdatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnupdatable

`func (o *BTDocumentElementInfo) SetUnupdatable(v bool)`

SetUnupdatable sets Unupdatable field to given value.

### HasUnupdatable

`func (o *BTDocumentElementInfo) HasUnupdatable() bool`

HasUnupdatable returns a boolean if a field has been set.

### GetVolumeUnits

`func (o *BTDocumentElementInfo) GetVolumeUnits() string`

GetVolumeUnits returns the VolumeUnits field if non-nil, zero value otherwise.

### GetVolumeUnitsOk

`func (o *BTDocumentElementInfo) GetVolumeUnitsOk() (*string, bool)`

GetVolumeUnitsOk returns a tuple with the VolumeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeUnits

`func (o *BTDocumentElementInfo) SetVolumeUnits(v string)`

SetVolumeUnits sets VolumeUnits field to given value.

### HasVolumeUnits

`func (o *BTDocumentElementInfo) HasVolumeUnits() bool`

HasVolumeUnits returns a boolean if a field has been set.

### GetZip

`func (o *BTDocumentElementInfo) GetZip() BTZipFileInfo`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *BTDocumentElementInfo) GetZipOk() (*BTZipFileInfo, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *BTDocumentElementInfo) SetZip(v BTZipFileInfo)`

SetZip sets Zip field to given value.

### HasZip

`func (o *BTDocumentElementInfo) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


