# BTPartMetadataInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appearance** | Pointer to [**BTPartAppearanceInfo**](BTPartAppearanceInfo.md) |  | [optional] 
**BodyType** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**CustomProperties** | Pointer to **map[string]string** |  | [optional] 
**DefaultColorHash** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsFlattenedBody** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsMesh** | Pointer to **bool** |  | [optional] 
**Material** | Pointer to [**BTPartMaterialInfo**](BTPartMaterialInfo.md) |  | [optional] 
**MeshState** | Pointer to [**GBTMeshState**](GBTMeshState.md) |  | [optional] 
**MetadataMicroversion** | Pointer to **string** |  | [optional] 
**MicroversionId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ordinal** | Pointer to **int32** |  | [optional] 
**PartId** | Pointer to **string** |  | [optional] 
**PartIdentity** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** |  | [optional] 
**PartQuery** | Pointer to **string** |  | [optional] 
**ProductLine** | Pointer to **string** |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**PropertySourceTypes** | Pointer to **map[string]int32** | &#x60;0: AUTOMATIC&#x60; Set automatically, like a part name |  &#x60;1: MERGED&#x60; Merged from another Part Studio | &#x60;2: FEATURE&#x60; Custom feature | &#x60;3: UNCONFIGURED&#x60; | &#x60;4: CONFIGURED&#x60; |  &#x60;5: STANDARD_CONTENT&#x60; | &#x60;6: DEFAULT&#x60; Applied from metadata property configuration | &#x60;7: COMPUTED&#x60; Non-overriden, non-configured, computed property |  &#x60;8: COMPUTED_CONFIGURED&#x60; Property is computed in this configuration; may also be configured in other configurations  &#x60;9: IMPORT&#x60; Imported properties are handled separately | [optional] 
**ReferencingConfiguredPartNodeIds** | Pointer to **[]string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**BTMetadataStateType**](BTMetadataStateType.md) |  | [optional] 
**ThumbnailConfigurationId** | Pointer to **string** |  | [optional] 
**ThumbnailInfo** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**Title1** | Pointer to **string** |  | [optional] 
**Title2** | Pointer to **string** |  | [optional] 
**Title3** | Pointer to **string** |  | [optional] 
**UnflattenedPartId** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPartMetadataInfo

`func NewBTPartMetadataInfo() *BTPartMetadataInfo`

NewBTPartMetadataInfo instantiates a new BTPartMetadataInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPartMetadataInfoWithDefaults

`func NewBTPartMetadataInfoWithDefaults() *BTPartMetadataInfo`

NewBTPartMetadataInfoWithDefaults instantiates a new BTPartMetadataInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearance

`func (o *BTPartMetadataInfo) GetAppearance() BTPartAppearanceInfo`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *BTPartMetadataInfo) GetAppearanceOk() (*BTPartAppearanceInfo, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *BTPartMetadataInfo) SetAppearance(v BTPartAppearanceInfo)`

SetAppearance sets Appearance field to given value.

### HasAppearance

`func (o *BTPartMetadataInfo) HasAppearance() bool`

HasAppearance returns a boolean if a field has been set.

### GetBodyType

`func (o *BTPartMetadataInfo) GetBodyType() string`

GetBodyType returns the BodyType field if non-nil, zero value otherwise.

### GetBodyTypeOk

`func (o *BTPartMetadataInfo) GetBodyTypeOk() (*string, bool)`

GetBodyTypeOk returns a tuple with the BodyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyType

`func (o *BTPartMetadataInfo) SetBodyType(v string)`

SetBodyType sets BodyType field to given value.

### HasBodyType

`func (o *BTPartMetadataInfo) HasBodyType() bool`

HasBodyType returns a boolean if a field has been set.

### GetConfigurationId

`func (o *BTPartMetadataInfo) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *BTPartMetadataInfo) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *BTPartMetadataInfo) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *BTPartMetadataInfo) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetCustomProperties

`func (o *BTPartMetadataInfo) GetCustomProperties() map[string]string`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BTPartMetadataInfo) GetCustomPropertiesOk() (*map[string]string, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BTPartMetadataInfo) SetCustomProperties(v map[string]string)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BTPartMetadataInfo) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDefaultColorHash

`func (o *BTPartMetadataInfo) GetDefaultColorHash() string`

GetDefaultColorHash returns the DefaultColorHash field if non-nil, zero value otherwise.

### GetDefaultColorHashOk

`func (o *BTPartMetadataInfo) GetDefaultColorHashOk() (*string, bool)`

GetDefaultColorHashOk returns a tuple with the DefaultColorHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultColorHash

`func (o *BTPartMetadataInfo) SetDefaultColorHash(v string)`

SetDefaultColorHash sets DefaultColorHash field to given value.

### HasDefaultColorHash

`func (o *BTPartMetadataInfo) HasDefaultColorHash() bool`

HasDefaultColorHash returns a boolean if a field has been set.

### GetDescription

`func (o *BTPartMetadataInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTPartMetadataInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTPartMetadataInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTPartMetadataInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetElementId

`func (o *BTPartMetadataInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTPartMetadataInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTPartMetadataInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTPartMetadataInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetHref

`func (o *BTPartMetadataInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTPartMetadataInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTPartMetadataInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTPartMetadataInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTPartMetadataInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTPartMetadataInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTPartMetadataInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTPartMetadataInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFlattenedBody

`func (o *BTPartMetadataInfo) GetIsFlattenedBody() bool`

GetIsFlattenedBody returns the IsFlattenedBody field if non-nil, zero value otherwise.

### GetIsFlattenedBodyOk

`func (o *BTPartMetadataInfo) GetIsFlattenedBodyOk() (*bool, bool)`

GetIsFlattenedBodyOk returns a tuple with the IsFlattenedBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlattenedBody

`func (o *BTPartMetadataInfo) SetIsFlattenedBody(v bool)`

SetIsFlattenedBody sets IsFlattenedBody field to given value.

### HasIsFlattenedBody

`func (o *BTPartMetadataInfo) HasIsFlattenedBody() bool`

HasIsFlattenedBody returns a boolean if a field has been set.

### GetIsHidden

`func (o *BTPartMetadataInfo) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *BTPartMetadataInfo) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *BTPartMetadataInfo) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *BTPartMetadataInfo) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsMesh

`func (o *BTPartMetadataInfo) GetIsMesh() bool`

GetIsMesh returns the IsMesh field if non-nil, zero value otherwise.

### GetIsMeshOk

`func (o *BTPartMetadataInfo) GetIsMeshOk() (*bool, bool)`

GetIsMeshOk returns a tuple with the IsMesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMesh

`func (o *BTPartMetadataInfo) SetIsMesh(v bool)`

SetIsMesh sets IsMesh field to given value.

### HasIsMesh

`func (o *BTPartMetadataInfo) HasIsMesh() bool`

HasIsMesh returns a boolean if a field has been set.

### GetMaterial

`func (o *BTPartMetadataInfo) GetMaterial() BTPartMaterialInfo`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *BTPartMetadataInfo) GetMaterialOk() (*BTPartMaterialInfo, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *BTPartMetadataInfo) SetMaterial(v BTPartMaterialInfo)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *BTPartMetadataInfo) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### GetMeshState

`func (o *BTPartMetadataInfo) GetMeshState() GBTMeshState`

GetMeshState returns the MeshState field if non-nil, zero value otherwise.

### GetMeshStateOk

`func (o *BTPartMetadataInfo) GetMeshStateOk() (*GBTMeshState, bool)`

GetMeshStateOk returns a tuple with the MeshState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshState

`func (o *BTPartMetadataInfo) SetMeshState(v GBTMeshState)`

SetMeshState sets MeshState field to given value.

### HasMeshState

`func (o *BTPartMetadataInfo) HasMeshState() bool`

HasMeshState returns a boolean if a field has been set.

### GetMetadataMicroversion

`func (o *BTPartMetadataInfo) GetMetadataMicroversion() string`

GetMetadataMicroversion returns the MetadataMicroversion field if non-nil, zero value otherwise.

### GetMetadataMicroversionOk

`func (o *BTPartMetadataInfo) GetMetadataMicroversionOk() (*string, bool)`

GetMetadataMicroversionOk returns a tuple with the MetadataMicroversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataMicroversion

`func (o *BTPartMetadataInfo) SetMetadataMicroversion(v string)`

SetMetadataMicroversion sets MetadataMicroversion field to given value.

### HasMetadataMicroversion

`func (o *BTPartMetadataInfo) HasMetadataMicroversion() bool`

HasMetadataMicroversion returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTPartMetadataInfo) GetMicroversionId() string`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTPartMetadataInfo) GetMicroversionIdOk() (*string, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTPartMetadataInfo) SetMicroversionId(v string)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTPartMetadataInfo) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetName

`func (o *BTPartMetadataInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPartMetadataInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPartMetadataInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTPartMetadataInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrdinal

`func (o *BTPartMetadataInfo) GetOrdinal() int32`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *BTPartMetadataInfo) GetOrdinalOk() (*int32, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *BTPartMetadataInfo) SetOrdinal(v int32)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *BTPartMetadataInfo) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### GetPartId

`func (o *BTPartMetadataInfo) GetPartId() string`

GetPartId returns the PartId field if non-nil, zero value otherwise.

### GetPartIdOk

`func (o *BTPartMetadataInfo) GetPartIdOk() (*string, bool)`

GetPartIdOk returns a tuple with the PartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartId

`func (o *BTPartMetadataInfo) SetPartId(v string)`

SetPartId sets PartId field to given value.

### HasPartId

`func (o *BTPartMetadataInfo) HasPartId() bool`

HasPartId returns a boolean if a field has been set.

### GetPartIdentity

`func (o *BTPartMetadataInfo) GetPartIdentity() string`

GetPartIdentity returns the PartIdentity field if non-nil, zero value otherwise.

### GetPartIdentityOk

`func (o *BTPartMetadataInfo) GetPartIdentityOk() (*string, bool)`

GetPartIdentityOk returns a tuple with the PartIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartIdentity

`func (o *BTPartMetadataInfo) SetPartIdentity(v string)`

SetPartIdentity sets PartIdentity field to given value.

### HasPartIdentity

`func (o *BTPartMetadataInfo) HasPartIdentity() bool`

HasPartIdentity returns a boolean if a field has been set.

### GetPartNumber

`func (o *BTPartMetadataInfo) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *BTPartMetadataInfo) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *BTPartMetadataInfo) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *BTPartMetadataInfo) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetPartQuery

`func (o *BTPartMetadataInfo) GetPartQuery() string`

GetPartQuery returns the PartQuery field if non-nil, zero value otherwise.

### GetPartQueryOk

`func (o *BTPartMetadataInfo) GetPartQueryOk() (*string, bool)`

GetPartQueryOk returns a tuple with the PartQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartQuery

`func (o *BTPartMetadataInfo) SetPartQuery(v string)`

SetPartQuery sets PartQuery field to given value.

### HasPartQuery

`func (o *BTPartMetadataInfo) HasPartQuery() bool`

HasPartQuery returns a boolean if a field has been set.

### GetProductLine

`func (o *BTPartMetadataInfo) GetProductLine() string`

GetProductLine returns the ProductLine field if non-nil, zero value otherwise.

### GetProductLineOk

`func (o *BTPartMetadataInfo) GetProductLineOk() (*string, bool)`

GetProductLineOk returns a tuple with the ProductLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductLine

`func (o *BTPartMetadataInfo) SetProductLine(v string)`

SetProductLine sets ProductLine field to given value.

### HasProductLine

`func (o *BTPartMetadataInfo) HasProductLine() bool`

HasProductLine returns a boolean if a field has been set.

### GetProject

`func (o *BTPartMetadataInfo) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BTPartMetadataInfo) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BTPartMetadataInfo) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *BTPartMetadataInfo) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPropertySourceTypes

`func (o *BTPartMetadataInfo) GetPropertySourceTypes() map[string]int32`

GetPropertySourceTypes returns the PropertySourceTypes field if non-nil, zero value otherwise.

### GetPropertySourceTypesOk

`func (o *BTPartMetadataInfo) GetPropertySourceTypesOk() (*map[string]int32, bool)`

GetPropertySourceTypesOk returns a tuple with the PropertySourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySourceTypes

`func (o *BTPartMetadataInfo) SetPropertySourceTypes(v map[string]int32)`

SetPropertySourceTypes sets PropertySourceTypes field to given value.

### HasPropertySourceTypes

`func (o *BTPartMetadataInfo) HasPropertySourceTypes() bool`

HasPropertySourceTypes returns a boolean if a field has been set.

### GetReferencingConfiguredPartNodeIds

`func (o *BTPartMetadataInfo) GetReferencingConfiguredPartNodeIds() []string`

GetReferencingConfiguredPartNodeIds returns the ReferencingConfiguredPartNodeIds field if non-nil, zero value otherwise.

### GetReferencingConfiguredPartNodeIdsOk

`func (o *BTPartMetadataInfo) GetReferencingConfiguredPartNodeIdsOk() (*[]string, bool)`

GetReferencingConfiguredPartNodeIdsOk returns a tuple with the ReferencingConfiguredPartNodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingConfiguredPartNodeIds

`func (o *BTPartMetadataInfo) SetReferencingConfiguredPartNodeIds(v []string)`

SetReferencingConfiguredPartNodeIds sets ReferencingConfiguredPartNodeIds field to given value.

### HasReferencingConfiguredPartNodeIds

`func (o *BTPartMetadataInfo) HasReferencingConfiguredPartNodeIds() bool`

HasReferencingConfiguredPartNodeIds returns a boolean if a field has been set.

### GetRevision

`func (o *BTPartMetadataInfo) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BTPartMetadataInfo) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BTPartMetadataInfo) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BTPartMetadataInfo) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetState

`func (o *BTPartMetadataInfo) GetState() BTMetadataStateType`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTPartMetadataInfo) GetStateOk() (*BTMetadataStateType, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTPartMetadataInfo) SetState(v BTMetadataStateType)`

SetState sets State field to given value.

### HasState

`func (o *BTPartMetadataInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetThumbnailConfigurationId

`func (o *BTPartMetadataInfo) GetThumbnailConfigurationId() string`

GetThumbnailConfigurationId returns the ThumbnailConfigurationId field if non-nil, zero value otherwise.

### GetThumbnailConfigurationIdOk

`func (o *BTPartMetadataInfo) GetThumbnailConfigurationIdOk() (*string, bool)`

GetThumbnailConfigurationIdOk returns a tuple with the ThumbnailConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailConfigurationId

`func (o *BTPartMetadataInfo) SetThumbnailConfigurationId(v string)`

SetThumbnailConfigurationId sets ThumbnailConfigurationId field to given value.

### HasThumbnailConfigurationId

`func (o *BTPartMetadataInfo) HasThumbnailConfigurationId() bool`

HasThumbnailConfigurationId returns a boolean if a field has been set.

### GetThumbnailInfo

`func (o *BTPartMetadataInfo) GetThumbnailInfo() BTThumbnailInfo`

GetThumbnailInfo returns the ThumbnailInfo field if non-nil, zero value otherwise.

### GetThumbnailInfoOk

`func (o *BTPartMetadataInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool)`

GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailInfo

`func (o *BTPartMetadataInfo) SetThumbnailInfo(v BTThumbnailInfo)`

SetThumbnailInfo sets ThumbnailInfo field to given value.

### HasThumbnailInfo

`func (o *BTPartMetadataInfo) HasThumbnailInfo() bool`

HasThumbnailInfo returns a boolean if a field has been set.

### GetTitle1

`func (o *BTPartMetadataInfo) GetTitle1() string`

GetTitle1 returns the Title1 field if non-nil, zero value otherwise.

### GetTitle1Ok

`func (o *BTPartMetadataInfo) GetTitle1Ok() (*string, bool)`

GetTitle1Ok returns a tuple with the Title1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle1

`func (o *BTPartMetadataInfo) SetTitle1(v string)`

SetTitle1 sets Title1 field to given value.

### HasTitle1

`func (o *BTPartMetadataInfo) HasTitle1() bool`

HasTitle1 returns a boolean if a field has been set.

### GetTitle2

`func (o *BTPartMetadataInfo) GetTitle2() string`

GetTitle2 returns the Title2 field if non-nil, zero value otherwise.

### GetTitle2Ok

`func (o *BTPartMetadataInfo) GetTitle2Ok() (*string, bool)`

GetTitle2Ok returns a tuple with the Title2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle2

`func (o *BTPartMetadataInfo) SetTitle2(v string)`

SetTitle2 sets Title2 field to given value.

### HasTitle2

`func (o *BTPartMetadataInfo) HasTitle2() bool`

HasTitle2 returns a boolean if a field has been set.

### GetTitle3

`func (o *BTPartMetadataInfo) GetTitle3() string`

GetTitle3 returns the Title3 field if non-nil, zero value otherwise.

### GetTitle3Ok

`func (o *BTPartMetadataInfo) GetTitle3Ok() (*string, bool)`

GetTitle3Ok returns a tuple with the Title3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle3

`func (o *BTPartMetadataInfo) SetTitle3(v string)`

SetTitle3 sets Title3 field to given value.

### HasTitle3

`func (o *BTPartMetadataInfo) HasTitle3() bool`

HasTitle3 returns a boolean if a field has been set.

### GetUnflattenedPartId

`func (o *BTPartMetadataInfo) GetUnflattenedPartId() string`

GetUnflattenedPartId returns the UnflattenedPartId field if non-nil, zero value otherwise.

### GetUnflattenedPartIdOk

`func (o *BTPartMetadataInfo) GetUnflattenedPartIdOk() (*string, bool)`

GetUnflattenedPartIdOk returns a tuple with the UnflattenedPartId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnflattenedPartId

`func (o *BTPartMetadataInfo) SetUnflattenedPartId(v string)`

SetUnflattenedPartId sets UnflattenedPartId field to given value.

### HasUnflattenedPartId

`func (o *BTPartMetadataInfo) HasUnflattenedPartId() bool`

HasUnflattenedPartId returns a boolean if a field has been set.

### GetVendor

`func (o *BTPartMetadataInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *BTPartMetadataInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *BTPartMetadataInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *BTPartMetadataInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


