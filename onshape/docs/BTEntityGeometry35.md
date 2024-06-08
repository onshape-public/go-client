# BTEntityGeometry35

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**Compressed** | Pointer to **bool** |  | [optional] 
**Decompressed** | Pointer to [**BTEntityGeometry35**](BTEntityGeometry35.md) |  | [optional] 
**Edge** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **int32** |  | [optional] 
**EstimatedMemoryUsageInBytes** | Pointer to **int32** |  | [optional] 
**Face** | Pointer to **bool** |  | [optional] 
**HasTessellationError** | Pointer to **bool** |  | [optional] 
**Point** | Pointer to **bool** |  | [optional] 
**SettingIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTEntityGeometry35

`func NewBTEntityGeometry35() *BTEntityGeometry35`

NewBTEntityGeometry35 instantiates a new BTEntityGeometry35 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTEntityGeometry35WithDefaults

`func NewBTEntityGeometry35WithDefaults() *BTEntityGeometry35`

NewBTEntityGeometry35WithDefaults instantiates a new BTEntityGeometry35 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTEntityGeometry35) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTEntityGeometry35) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTEntityGeometry35) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTEntityGeometry35) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCompressed

`func (o *BTEntityGeometry35) GetCompressed() bool`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *BTEntityGeometry35) GetCompressedOk() (*bool, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *BTEntityGeometry35) SetCompressed(v bool)`

SetCompressed sets Compressed field to given value.

### HasCompressed

`func (o *BTEntityGeometry35) HasCompressed() bool`

HasCompressed returns a boolean if a field has been set.

### GetDecompressed

`func (o *BTEntityGeometry35) GetDecompressed() BTEntityGeometry35`

GetDecompressed returns the Decompressed field if non-nil, zero value otherwise.

### GetDecompressedOk

`func (o *BTEntityGeometry35) GetDecompressedOk() (*BTEntityGeometry35, bool)`

GetDecompressedOk returns a tuple with the Decompressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecompressed

`func (o *BTEntityGeometry35) SetDecompressed(v BTEntityGeometry35)`

SetDecompressed sets Decompressed field to given value.

### HasDecompressed

`func (o *BTEntityGeometry35) HasDecompressed() bool`

HasDecompressed returns a boolean if a field has been set.

### GetEdge

`func (o *BTEntityGeometry35) GetEdge() bool`

GetEdge returns the Edge field if non-nil, zero value otherwise.

### GetEdgeOk

`func (o *BTEntityGeometry35) GetEdgeOk() (*bool, bool)`

GetEdgeOk returns a tuple with the Edge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdge

`func (o *BTEntityGeometry35) SetEdge(v bool)`

SetEdge sets Edge field to given value.

### HasEdge

`func (o *BTEntityGeometry35) HasEdge() bool`

HasEdge returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTEntityGeometry35) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTEntityGeometry35) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTEntityGeometry35) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTEntityGeometry35) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetEstimatedMemoryUsageInBytes

`func (o *BTEntityGeometry35) GetEstimatedMemoryUsageInBytes() int32`

GetEstimatedMemoryUsageInBytes returns the EstimatedMemoryUsageInBytes field if non-nil, zero value otherwise.

### GetEstimatedMemoryUsageInBytesOk

`func (o *BTEntityGeometry35) GetEstimatedMemoryUsageInBytesOk() (*int32, bool)`

GetEstimatedMemoryUsageInBytesOk returns a tuple with the EstimatedMemoryUsageInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMemoryUsageInBytes

`func (o *BTEntityGeometry35) SetEstimatedMemoryUsageInBytes(v int32)`

SetEstimatedMemoryUsageInBytes sets EstimatedMemoryUsageInBytes field to given value.

### HasEstimatedMemoryUsageInBytes

`func (o *BTEntityGeometry35) HasEstimatedMemoryUsageInBytes() bool`

HasEstimatedMemoryUsageInBytes returns a boolean if a field has been set.

### GetFace

`func (o *BTEntityGeometry35) GetFace() bool`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *BTEntityGeometry35) GetFaceOk() (*bool, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *BTEntityGeometry35) SetFace(v bool)`

SetFace sets Face field to given value.

### HasFace

`func (o *BTEntityGeometry35) HasFace() bool`

HasFace returns a boolean if a field has been set.

### GetHasTessellationError

`func (o *BTEntityGeometry35) GetHasTessellationError() bool`

GetHasTessellationError returns the HasTessellationError field if non-nil, zero value otherwise.

### GetHasTessellationErrorOk

`func (o *BTEntityGeometry35) GetHasTessellationErrorOk() (*bool, bool)`

GetHasTessellationErrorOk returns a tuple with the HasTessellationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTessellationError

`func (o *BTEntityGeometry35) SetHasTessellationError(v bool)`

SetHasTessellationError sets HasTessellationError field to given value.

### HasHasTessellationError

`func (o *BTEntityGeometry35) HasHasTessellationError() bool`

HasHasTessellationError returns a boolean if a field has been set.

### GetPoint

`func (o *BTEntityGeometry35) GetPoint() bool`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *BTEntityGeometry35) GetPointOk() (*bool, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *BTEntityGeometry35) SetPoint(v bool)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *BTEntityGeometry35) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetSettingIndex

`func (o *BTEntityGeometry35) GetSettingIndex() int32`

GetSettingIndex returns the SettingIndex field if non-nil, zero value otherwise.

### GetSettingIndexOk

`func (o *BTEntityGeometry35) GetSettingIndexOk() (*int32, bool)`

GetSettingIndexOk returns a tuple with the SettingIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingIndex

`func (o *BTEntityGeometry35) SetSettingIndex(v int32)`

SetSettingIndex sets SettingIndex field to given value.

### HasSettingIndex

`func (o *BTEntityGeometry35) HasSettingIndex() bool`

HasSettingIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


