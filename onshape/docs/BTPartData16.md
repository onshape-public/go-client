# BTPartData16

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BestAvailableTessellationSetting** | Pointer to [**GBTTessellationSettingEnum**](GBTTessellationSettingEnum.md) |  | [optional] 
**BoundsDiameter** | Pointer to **float64** |  | [optional] 
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**ClosedConstituentPartData** | Pointer to [**BTClosedConstituentPartData2911**](BTClosedConstituentPartData2911.md) |  | [optional] 
**CoarsePlanarFaceTriangleCount** | Pointer to **int32** |  | [optional] 
**CoarseTriangleCount** | Pointer to **int32** |  | [optional] 
**ConstituentBodyDeterministicIds** | Pointer to **[]string** |  | [optional] 
**CopyWithoutEntities** | Pointer to [**BTPartData16**](BTPartData16.md) |  | [optional] 
**EntityDIds** | Pointer to **[]string** |  | [optional] 
**EntityDeterministicIds** | Pointer to **[]string** |  | [optional] 
**FlattenedToUnflattenedEntitiesMapping** | Pointer to **map[string][]string** |  | [optional] 
**FlattenedToUnflattenedMapping** | Pointer to **map[string]string** |  | [optional] 
**HighBoxCorner** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**IsACopyForTessellationSettings** | Pointer to **bool** |  | [optional] 
**IsAssociatedWithFlat** | Pointer to **bool** |  | [optional] 
**IsBendCenterLineBody** | Pointer to **bool** |  | [optional] 
**IsClosedComposite** | Pointer to **bool** |  | [optional] 
**IsComposite** | Pointer to **bool** |  | [optional] 
**IsDeletion** | Pointer to **bool** |  | [optional] 
**IsEntitylessPartData** | Pointer to **bool** |  | [optional] 
**IsFlattenedSheetMetalBody** | Pointer to **bool** |  | [optional] 
**IsOpenComposite** | Pointer to **bool** |  | [optional] 
**LowBoxCorner** | Pointer to [**BTVector3d389**](BTVector3d389.md) |  | [optional] 
**OwnerFlattenedBodyId** | Pointer to **string** |  | [optional] 
**SheetMetalModelId** | Pointer to **string** |  | [optional] 
**SheetMetalOrderId** | Pointer to **string** |  | [optional] 
**ShouldAlwaysUseHighestQualityTessellation** | Pointer to **bool** |  | [optional] 
**TessellationSettings** | Pointer to **[]int32** |  | [optional] 
**TotalEntityCount** | Pointer to **int32** |  | [optional] 
**UserTessellationSetting** | Pointer to [**GBTTessellationSettingEnum**](GBTTessellationSettingEnum.md) |  | [optional] 

## Methods

### NewBTPartData16

`func NewBTPartData16() *BTPartData16`

NewBTPartData16 instantiates a new BTPartData16 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPartData16WithDefaults

`func NewBTPartData16WithDefaults() *BTPartData16`

NewBTPartData16WithDefaults instantiates a new BTPartData16 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBestAvailableTessellationSetting

`func (o *BTPartData16) GetBestAvailableTessellationSetting() GBTTessellationSettingEnum`

GetBestAvailableTessellationSetting returns the BestAvailableTessellationSetting field if non-nil, zero value otherwise.

### GetBestAvailableTessellationSettingOk

`func (o *BTPartData16) GetBestAvailableTessellationSettingOk() (*GBTTessellationSettingEnum, bool)`

GetBestAvailableTessellationSettingOk returns a tuple with the BestAvailableTessellationSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAvailableTessellationSetting

`func (o *BTPartData16) SetBestAvailableTessellationSetting(v GBTTessellationSettingEnum)`

SetBestAvailableTessellationSetting sets BestAvailableTessellationSetting field to given value.

### HasBestAvailableTessellationSetting

`func (o *BTPartData16) HasBestAvailableTessellationSetting() bool`

HasBestAvailableTessellationSetting returns a boolean if a field has been set.

### GetBoundsDiameter

`func (o *BTPartData16) GetBoundsDiameter() float64`

GetBoundsDiameter returns the BoundsDiameter field if non-nil, zero value otherwise.

### GetBoundsDiameterOk

`func (o *BTPartData16) GetBoundsDiameterOk() (*float64, bool)`

GetBoundsDiameterOk returns a tuple with the BoundsDiameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundsDiameter

`func (o *BTPartData16) SetBoundsDiameter(v float64)`

SetBoundsDiameter sets BoundsDiameter field to given value.

### HasBoundsDiameter

`func (o *BTPartData16) HasBoundsDiameter() bool`

HasBoundsDiameter returns a boolean if a field has been set.

### GetBtType

`func (o *BTPartData16) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPartData16) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPartData16) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPartData16) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetClosedConstituentPartData

`func (o *BTPartData16) GetClosedConstituentPartData() BTClosedConstituentPartData2911`

GetClosedConstituentPartData returns the ClosedConstituentPartData field if non-nil, zero value otherwise.

### GetClosedConstituentPartDataOk

`func (o *BTPartData16) GetClosedConstituentPartDataOk() (*BTClosedConstituentPartData2911, bool)`

GetClosedConstituentPartDataOk returns a tuple with the ClosedConstituentPartData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedConstituentPartData

`func (o *BTPartData16) SetClosedConstituentPartData(v BTClosedConstituentPartData2911)`

SetClosedConstituentPartData sets ClosedConstituentPartData field to given value.

### HasClosedConstituentPartData

`func (o *BTPartData16) HasClosedConstituentPartData() bool`

HasClosedConstituentPartData returns a boolean if a field has been set.

### GetCoarsePlanarFaceTriangleCount

`func (o *BTPartData16) GetCoarsePlanarFaceTriangleCount() int32`

GetCoarsePlanarFaceTriangleCount returns the CoarsePlanarFaceTriangleCount field if non-nil, zero value otherwise.

### GetCoarsePlanarFaceTriangleCountOk

`func (o *BTPartData16) GetCoarsePlanarFaceTriangleCountOk() (*int32, bool)`

GetCoarsePlanarFaceTriangleCountOk returns a tuple with the CoarsePlanarFaceTriangleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoarsePlanarFaceTriangleCount

`func (o *BTPartData16) SetCoarsePlanarFaceTriangleCount(v int32)`

SetCoarsePlanarFaceTriangleCount sets CoarsePlanarFaceTriangleCount field to given value.

### HasCoarsePlanarFaceTriangleCount

`func (o *BTPartData16) HasCoarsePlanarFaceTriangleCount() bool`

HasCoarsePlanarFaceTriangleCount returns a boolean if a field has been set.

### GetCoarseTriangleCount

`func (o *BTPartData16) GetCoarseTriangleCount() int32`

GetCoarseTriangleCount returns the CoarseTriangleCount field if non-nil, zero value otherwise.

### GetCoarseTriangleCountOk

`func (o *BTPartData16) GetCoarseTriangleCountOk() (*int32, bool)`

GetCoarseTriangleCountOk returns a tuple with the CoarseTriangleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoarseTriangleCount

`func (o *BTPartData16) SetCoarseTriangleCount(v int32)`

SetCoarseTriangleCount sets CoarseTriangleCount field to given value.

### HasCoarseTriangleCount

`func (o *BTPartData16) HasCoarseTriangleCount() bool`

HasCoarseTriangleCount returns a boolean if a field has been set.

### GetConstituentBodyDeterministicIds

`func (o *BTPartData16) GetConstituentBodyDeterministicIds() []string`

GetConstituentBodyDeterministicIds returns the ConstituentBodyDeterministicIds field if non-nil, zero value otherwise.

### GetConstituentBodyDeterministicIdsOk

`func (o *BTPartData16) GetConstituentBodyDeterministicIdsOk() (*[]string, bool)`

GetConstituentBodyDeterministicIdsOk returns a tuple with the ConstituentBodyDeterministicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituentBodyDeterministicIds

`func (o *BTPartData16) SetConstituentBodyDeterministicIds(v []string)`

SetConstituentBodyDeterministicIds sets ConstituentBodyDeterministicIds field to given value.

### HasConstituentBodyDeterministicIds

`func (o *BTPartData16) HasConstituentBodyDeterministicIds() bool`

HasConstituentBodyDeterministicIds returns a boolean if a field has been set.

### GetCopyWithoutEntities

`func (o *BTPartData16) GetCopyWithoutEntities() BTPartData16`

GetCopyWithoutEntities returns the CopyWithoutEntities field if non-nil, zero value otherwise.

### GetCopyWithoutEntitiesOk

`func (o *BTPartData16) GetCopyWithoutEntitiesOk() (*BTPartData16, bool)`

GetCopyWithoutEntitiesOk returns a tuple with the CopyWithoutEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyWithoutEntities

`func (o *BTPartData16) SetCopyWithoutEntities(v BTPartData16)`

SetCopyWithoutEntities sets CopyWithoutEntities field to given value.

### HasCopyWithoutEntities

`func (o *BTPartData16) HasCopyWithoutEntities() bool`

HasCopyWithoutEntities returns a boolean if a field has been set.

### GetEntityDIds

`func (o *BTPartData16) GetEntityDIds() []string`

GetEntityDIds returns the EntityDIds field if non-nil, zero value otherwise.

### GetEntityDIdsOk

`func (o *BTPartData16) GetEntityDIdsOk() (*[]string, bool)`

GetEntityDIdsOk returns a tuple with the EntityDIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDIds

`func (o *BTPartData16) SetEntityDIds(v []string)`

SetEntityDIds sets EntityDIds field to given value.

### HasEntityDIds

`func (o *BTPartData16) HasEntityDIds() bool`

HasEntityDIds returns a boolean if a field has been set.

### GetEntityDeterministicIds

`func (o *BTPartData16) GetEntityDeterministicIds() []string`

GetEntityDeterministicIds returns the EntityDeterministicIds field if non-nil, zero value otherwise.

### GetEntityDeterministicIdsOk

`func (o *BTPartData16) GetEntityDeterministicIdsOk() (*[]string, bool)`

GetEntityDeterministicIdsOk returns a tuple with the EntityDeterministicIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityDeterministicIds

`func (o *BTPartData16) SetEntityDeterministicIds(v []string)`

SetEntityDeterministicIds sets EntityDeterministicIds field to given value.

### HasEntityDeterministicIds

`func (o *BTPartData16) HasEntityDeterministicIds() bool`

HasEntityDeterministicIds returns a boolean if a field has been set.

### GetFlattenedToUnflattenedEntitiesMapping

`func (o *BTPartData16) GetFlattenedToUnflattenedEntitiesMapping() map[string][]string`

GetFlattenedToUnflattenedEntitiesMapping returns the FlattenedToUnflattenedEntitiesMapping field if non-nil, zero value otherwise.

### GetFlattenedToUnflattenedEntitiesMappingOk

`func (o *BTPartData16) GetFlattenedToUnflattenedEntitiesMappingOk() (*map[string][]string, bool)`

GetFlattenedToUnflattenedEntitiesMappingOk returns a tuple with the FlattenedToUnflattenedEntitiesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenedToUnflattenedEntitiesMapping

`func (o *BTPartData16) SetFlattenedToUnflattenedEntitiesMapping(v map[string][]string)`

SetFlattenedToUnflattenedEntitiesMapping sets FlattenedToUnflattenedEntitiesMapping field to given value.

### HasFlattenedToUnflattenedEntitiesMapping

`func (o *BTPartData16) HasFlattenedToUnflattenedEntitiesMapping() bool`

HasFlattenedToUnflattenedEntitiesMapping returns a boolean if a field has been set.

### GetFlattenedToUnflattenedMapping

`func (o *BTPartData16) GetFlattenedToUnflattenedMapping() map[string]string`

GetFlattenedToUnflattenedMapping returns the FlattenedToUnflattenedMapping field if non-nil, zero value otherwise.

### GetFlattenedToUnflattenedMappingOk

`func (o *BTPartData16) GetFlattenedToUnflattenedMappingOk() (*map[string]string, bool)`

GetFlattenedToUnflattenedMappingOk returns a tuple with the FlattenedToUnflattenedMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlattenedToUnflattenedMapping

`func (o *BTPartData16) SetFlattenedToUnflattenedMapping(v map[string]string)`

SetFlattenedToUnflattenedMapping sets FlattenedToUnflattenedMapping field to given value.

### HasFlattenedToUnflattenedMapping

`func (o *BTPartData16) HasFlattenedToUnflattenedMapping() bool`

HasFlattenedToUnflattenedMapping returns a boolean if a field has been set.

### GetHighBoxCorner

`func (o *BTPartData16) GetHighBoxCorner() BTVector3d389`

GetHighBoxCorner returns the HighBoxCorner field if non-nil, zero value otherwise.

### GetHighBoxCornerOk

`func (o *BTPartData16) GetHighBoxCornerOk() (*BTVector3d389, bool)`

GetHighBoxCornerOk returns a tuple with the HighBoxCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighBoxCorner

`func (o *BTPartData16) SetHighBoxCorner(v BTVector3d389)`

SetHighBoxCorner sets HighBoxCorner field to given value.

### HasHighBoxCorner

`func (o *BTPartData16) HasHighBoxCorner() bool`

HasHighBoxCorner returns a boolean if a field has been set.

### GetIsACopyForTessellationSettings

`func (o *BTPartData16) GetIsACopyForTessellationSettings() bool`

GetIsACopyForTessellationSettings returns the IsACopyForTessellationSettings field if non-nil, zero value otherwise.

### GetIsACopyForTessellationSettingsOk

`func (o *BTPartData16) GetIsACopyForTessellationSettingsOk() (*bool, bool)`

GetIsACopyForTessellationSettingsOk returns a tuple with the IsACopyForTessellationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsACopyForTessellationSettings

`func (o *BTPartData16) SetIsACopyForTessellationSettings(v bool)`

SetIsACopyForTessellationSettings sets IsACopyForTessellationSettings field to given value.

### HasIsACopyForTessellationSettings

`func (o *BTPartData16) HasIsACopyForTessellationSettings() bool`

HasIsACopyForTessellationSettings returns a boolean if a field has been set.

### GetIsAssociatedWithFlat

`func (o *BTPartData16) GetIsAssociatedWithFlat() bool`

GetIsAssociatedWithFlat returns the IsAssociatedWithFlat field if non-nil, zero value otherwise.

### GetIsAssociatedWithFlatOk

`func (o *BTPartData16) GetIsAssociatedWithFlatOk() (*bool, bool)`

GetIsAssociatedWithFlatOk returns a tuple with the IsAssociatedWithFlat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAssociatedWithFlat

`func (o *BTPartData16) SetIsAssociatedWithFlat(v bool)`

SetIsAssociatedWithFlat sets IsAssociatedWithFlat field to given value.

### HasIsAssociatedWithFlat

`func (o *BTPartData16) HasIsAssociatedWithFlat() bool`

HasIsAssociatedWithFlat returns a boolean if a field has been set.

### GetIsBendCenterLineBody

`func (o *BTPartData16) GetIsBendCenterLineBody() bool`

GetIsBendCenterLineBody returns the IsBendCenterLineBody field if non-nil, zero value otherwise.

### GetIsBendCenterLineBodyOk

`func (o *BTPartData16) GetIsBendCenterLineBodyOk() (*bool, bool)`

GetIsBendCenterLineBodyOk returns a tuple with the IsBendCenterLineBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBendCenterLineBody

`func (o *BTPartData16) SetIsBendCenterLineBody(v bool)`

SetIsBendCenterLineBody sets IsBendCenterLineBody field to given value.

### HasIsBendCenterLineBody

`func (o *BTPartData16) HasIsBendCenterLineBody() bool`

HasIsBendCenterLineBody returns a boolean if a field has been set.

### GetIsClosedComposite

`func (o *BTPartData16) GetIsClosedComposite() bool`

GetIsClosedComposite returns the IsClosedComposite field if non-nil, zero value otherwise.

### GetIsClosedCompositeOk

`func (o *BTPartData16) GetIsClosedCompositeOk() (*bool, bool)`

GetIsClosedCompositeOk returns a tuple with the IsClosedComposite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosedComposite

`func (o *BTPartData16) SetIsClosedComposite(v bool)`

SetIsClosedComposite sets IsClosedComposite field to given value.

### HasIsClosedComposite

`func (o *BTPartData16) HasIsClosedComposite() bool`

HasIsClosedComposite returns a boolean if a field has been set.

### GetIsComposite

`func (o *BTPartData16) GetIsComposite() bool`

GetIsComposite returns the IsComposite field if non-nil, zero value otherwise.

### GetIsCompositeOk

`func (o *BTPartData16) GetIsCompositeOk() (*bool, bool)`

GetIsCompositeOk returns a tuple with the IsComposite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComposite

`func (o *BTPartData16) SetIsComposite(v bool)`

SetIsComposite sets IsComposite field to given value.

### HasIsComposite

`func (o *BTPartData16) HasIsComposite() bool`

HasIsComposite returns a boolean if a field has been set.

### GetIsDeletion

`func (o *BTPartData16) GetIsDeletion() bool`

GetIsDeletion returns the IsDeletion field if non-nil, zero value otherwise.

### GetIsDeletionOk

`func (o *BTPartData16) GetIsDeletionOk() (*bool, bool)`

GetIsDeletionOk returns a tuple with the IsDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletion

`func (o *BTPartData16) SetIsDeletion(v bool)`

SetIsDeletion sets IsDeletion field to given value.

### HasIsDeletion

`func (o *BTPartData16) HasIsDeletion() bool`

HasIsDeletion returns a boolean if a field has been set.

### GetIsEntitylessPartData

`func (o *BTPartData16) GetIsEntitylessPartData() bool`

GetIsEntitylessPartData returns the IsEntitylessPartData field if non-nil, zero value otherwise.

### GetIsEntitylessPartDataOk

`func (o *BTPartData16) GetIsEntitylessPartDataOk() (*bool, bool)`

GetIsEntitylessPartDataOk returns a tuple with the IsEntitylessPartData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntitylessPartData

`func (o *BTPartData16) SetIsEntitylessPartData(v bool)`

SetIsEntitylessPartData sets IsEntitylessPartData field to given value.

### HasIsEntitylessPartData

`func (o *BTPartData16) HasIsEntitylessPartData() bool`

HasIsEntitylessPartData returns a boolean if a field has been set.

### GetIsFlattenedSheetMetalBody

`func (o *BTPartData16) GetIsFlattenedSheetMetalBody() bool`

GetIsFlattenedSheetMetalBody returns the IsFlattenedSheetMetalBody field if non-nil, zero value otherwise.

### GetIsFlattenedSheetMetalBodyOk

`func (o *BTPartData16) GetIsFlattenedSheetMetalBodyOk() (*bool, bool)`

GetIsFlattenedSheetMetalBodyOk returns a tuple with the IsFlattenedSheetMetalBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattenedSheetMetalBody

`func (o *BTPartData16) SetIsFlattenedSheetMetalBody(v bool)`

SetIsFlattenedSheetMetalBody sets IsFlattenedSheetMetalBody field to given value.

### HasIsFlattenedSheetMetalBody

`func (o *BTPartData16) HasIsFlattenedSheetMetalBody() bool`

HasIsFlattenedSheetMetalBody returns a boolean if a field has been set.

### GetIsOpenComposite

`func (o *BTPartData16) GetIsOpenComposite() bool`

GetIsOpenComposite returns the IsOpenComposite field if non-nil, zero value otherwise.

### GetIsOpenCompositeOk

`func (o *BTPartData16) GetIsOpenCompositeOk() (*bool, bool)`

GetIsOpenCompositeOk returns a tuple with the IsOpenComposite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenComposite

`func (o *BTPartData16) SetIsOpenComposite(v bool)`

SetIsOpenComposite sets IsOpenComposite field to given value.

### HasIsOpenComposite

`func (o *BTPartData16) HasIsOpenComposite() bool`

HasIsOpenComposite returns a boolean if a field has been set.

### GetLowBoxCorner

`func (o *BTPartData16) GetLowBoxCorner() BTVector3d389`

GetLowBoxCorner returns the LowBoxCorner field if non-nil, zero value otherwise.

### GetLowBoxCornerOk

`func (o *BTPartData16) GetLowBoxCornerOk() (*BTVector3d389, bool)`

GetLowBoxCornerOk returns a tuple with the LowBoxCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowBoxCorner

`func (o *BTPartData16) SetLowBoxCorner(v BTVector3d389)`

SetLowBoxCorner sets LowBoxCorner field to given value.

### HasLowBoxCorner

`func (o *BTPartData16) HasLowBoxCorner() bool`

HasLowBoxCorner returns a boolean if a field has been set.

### GetOwnerFlattenedBodyId

`func (o *BTPartData16) GetOwnerFlattenedBodyId() string`

GetOwnerFlattenedBodyId returns the OwnerFlattenedBodyId field if non-nil, zero value otherwise.

### GetOwnerFlattenedBodyIdOk

`func (o *BTPartData16) GetOwnerFlattenedBodyIdOk() (*string, bool)`

GetOwnerFlattenedBodyIdOk returns a tuple with the OwnerFlattenedBodyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerFlattenedBodyId

`func (o *BTPartData16) SetOwnerFlattenedBodyId(v string)`

SetOwnerFlattenedBodyId sets OwnerFlattenedBodyId field to given value.

### HasOwnerFlattenedBodyId

`func (o *BTPartData16) HasOwnerFlattenedBodyId() bool`

HasOwnerFlattenedBodyId returns a boolean if a field has been set.

### GetSheetMetalModelId

`func (o *BTPartData16) GetSheetMetalModelId() string`

GetSheetMetalModelId returns the SheetMetalModelId field if non-nil, zero value otherwise.

### GetSheetMetalModelIdOk

`func (o *BTPartData16) GetSheetMetalModelIdOk() (*string, bool)`

GetSheetMetalModelIdOk returns a tuple with the SheetMetalModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetMetalModelId

`func (o *BTPartData16) SetSheetMetalModelId(v string)`

SetSheetMetalModelId sets SheetMetalModelId field to given value.

### HasSheetMetalModelId

`func (o *BTPartData16) HasSheetMetalModelId() bool`

HasSheetMetalModelId returns a boolean if a field has been set.

### GetSheetMetalOrderId

`func (o *BTPartData16) GetSheetMetalOrderId() string`

GetSheetMetalOrderId returns the SheetMetalOrderId field if non-nil, zero value otherwise.

### GetSheetMetalOrderIdOk

`func (o *BTPartData16) GetSheetMetalOrderIdOk() (*string, bool)`

GetSheetMetalOrderIdOk returns a tuple with the SheetMetalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheetMetalOrderId

`func (o *BTPartData16) SetSheetMetalOrderId(v string)`

SetSheetMetalOrderId sets SheetMetalOrderId field to given value.

### HasSheetMetalOrderId

`func (o *BTPartData16) HasSheetMetalOrderId() bool`

HasSheetMetalOrderId returns a boolean if a field has been set.

### GetShouldAlwaysUseHighestQualityTessellation

`func (o *BTPartData16) GetShouldAlwaysUseHighestQualityTessellation() bool`

GetShouldAlwaysUseHighestQualityTessellation returns the ShouldAlwaysUseHighestQualityTessellation field if non-nil, zero value otherwise.

### GetShouldAlwaysUseHighestQualityTessellationOk

`func (o *BTPartData16) GetShouldAlwaysUseHighestQualityTessellationOk() (*bool, bool)`

GetShouldAlwaysUseHighestQualityTessellationOk returns a tuple with the ShouldAlwaysUseHighestQualityTessellation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldAlwaysUseHighestQualityTessellation

`func (o *BTPartData16) SetShouldAlwaysUseHighestQualityTessellation(v bool)`

SetShouldAlwaysUseHighestQualityTessellation sets ShouldAlwaysUseHighestQualityTessellation field to given value.

### HasShouldAlwaysUseHighestQualityTessellation

`func (o *BTPartData16) HasShouldAlwaysUseHighestQualityTessellation() bool`

HasShouldAlwaysUseHighestQualityTessellation returns a boolean if a field has been set.

### GetTessellationSettings

`func (o *BTPartData16) GetTessellationSettings() []int32`

GetTessellationSettings returns the TessellationSettings field if non-nil, zero value otherwise.

### GetTessellationSettingsOk

`func (o *BTPartData16) GetTessellationSettingsOk() (*[]int32, bool)`

GetTessellationSettingsOk returns a tuple with the TessellationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTessellationSettings

`func (o *BTPartData16) SetTessellationSettings(v []int32)`

SetTessellationSettings sets TessellationSettings field to given value.

### HasTessellationSettings

`func (o *BTPartData16) HasTessellationSettings() bool`

HasTessellationSettings returns a boolean if a field has been set.

### GetTotalEntityCount

`func (o *BTPartData16) GetTotalEntityCount() int32`

GetTotalEntityCount returns the TotalEntityCount field if non-nil, zero value otherwise.

### GetTotalEntityCountOk

`func (o *BTPartData16) GetTotalEntityCountOk() (*int32, bool)`

GetTotalEntityCountOk returns a tuple with the TotalEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntityCount

`func (o *BTPartData16) SetTotalEntityCount(v int32)`

SetTotalEntityCount sets TotalEntityCount field to given value.

### HasTotalEntityCount

`func (o *BTPartData16) HasTotalEntityCount() bool`

HasTotalEntityCount returns a boolean if a field has been set.

### GetUserTessellationSetting

`func (o *BTPartData16) GetUserTessellationSetting() GBTTessellationSettingEnum`

GetUserTessellationSetting returns the UserTessellationSetting field if non-nil, zero value otherwise.

### GetUserTessellationSettingOk

`func (o *BTPartData16) GetUserTessellationSettingOk() (*GBTTessellationSettingEnum, bool)`

GetUserTessellationSettingOk returns a tuple with the UserTessellationSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTessellationSetting

`func (o *BTPartData16) SetUserTessellationSetting(v GBTTessellationSettingEnum)`

SetUserTessellationSetting sets UserTessellationSetting field to given value.

### HasUserTessellationSetting

`func (o *BTPartData16) HasUserTessellationSetting() bool`

HasUserTessellationSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


