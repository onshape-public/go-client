# BTCloudStorageObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanMove** | Pointer to **bool** |  | [optional] 
**CloudStorageAccountId** | Pointer to **string** |  | [optional] 
**CloudStorageObjectId** | Pointer to **string** |  | [optional] 
**CloudStorageProvider** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **JSONTime** |  | [optional] 
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**CreatedById** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**IconLink** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsContainer** | Pointer to **bool** |  | [optional] 
**IsEnterpriseOwned** | Pointer to **bool** |  | [optional] 
**IsMutable** | Pointer to **bool** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**ModifiedAt** | Pointer to **JSONTime** |  | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**ModifiedById** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**BTOwnerInfo**](BTOwnerInfo.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**SizeBytes** | Pointer to **int64** |  | [optional] 
**ThumbnailInfo** | Pointer to [**BTThumbnailInfo**](BTThumbnailInfo.md) |  | [optional] 
**TreeHref** | Pointer to **string** |  | [optional] 
**UnparentHref** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 
**WebViewLink** | Pointer to **string** |  | [optional] 

## Methods

### NewBTCloudStorageObjectInfo

`func NewBTCloudStorageObjectInfo() *BTCloudStorageObjectInfo`

NewBTCloudStorageObjectInfo instantiates a new BTCloudStorageObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCloudStorageObjectInfoWithDefaults

`func NewBTCloudStorageObjectInfoWithDefaults() *BTCloudStorageObjectInfo`

NewBTCloudStorageObjectInfoWithDefaults instantiates a new BTCloudStorageObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanMove

`func (o *BTCloudStorageObjectInfo) GetCanMove() bool`

GetCanMove returns the CanMove field if non-nil, zero value otherwise.

### GetCanMoveOk

`func (o *BTCloudStorageObjectInfo) GetCanMoveOk() (*bool, bool)`

GetCanMoveOk returns a tuple with the CanMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMove

`func (o *BTCloudStorageObjectInfo) SetCanMove(v bool)`

SetCanMove sets CanMove field to given value.

### HasCanMove

`func (o *BTCloudStorageObjectInfo) HasCanMove() bool`

HasCanMove returns a boolean if a field has been set.

### GetCloudStorageAccountId

`func (o *BTCloudStorageObjectInfo) GetCloudStorageAccountId() string`

GetCloudStorageAccountId returns the CloudStorageAccountId field if non-nil, zero value otherwise.

### GetCloudStorageAccountIdOk

`func (o *BTCloudStorageObjectInfo) GetCloudStorageAccountIdOk() (*string, bool)`

GetCloudStorageAccountIdOk returns a tuple with the CloudStorageAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageAccountId

`func (o *BTCloudStorageObjectInfo) SetCloudStorageAccountId(v string)`

SetCloudStorageAccountId sets CloudStorageAccountId field to given value.

### HasCloudStorageAccountId

`func (o *BTCloudStorageObjectInfo) HasCloudStorageAccountId() bool`

HasCloudStorageAccountId returns a boolean if a field has been set.

### GetCloudStorageObjectId

`func (o *BTCloudStorageObjectInfo) GetCloudStorageObjectId() string`

GetCloudStorageObjectId returns the CloudStorageObjectId field if non-nil, zero value otherwise.

### GetCloudStorageObjectIdOk

`func (o *BTCloudStorageObjectInfo) GetCloudStorageObjectIdOk() (*string, bool)`

GetCloudStorageObjectIdOk returns a tuple with the CloudStorageObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageObjectId

`func (o *BTCloudStorageObjectInfo) SetCloudStorageObjectId(v string)`

SetCloudStorageObjectId sets CloudStorageObjectId field to given value.

### HasCloudStorageObjectId

`func (o *BTCloudStorageObjectInfo) HasCloudStorageObjectId() bool`

HasCloudStorageObjectId returns a boolean if a field has been set.

### GetCloudStorageProvider

`func (o *BTCloudStorageObjectInfo) GetCloudStorageProvider() int32`

GetCloudStorageProvider returns the CloudStorageProvider field if non-nil, zero value otherwise.

### GetCloudStorageProviderOk

`func (o *BTCloudStorageObjectInfo) GetCloudStorageProviderOk() (*int32, bool)`

GetCloudStorageProviderOk returns a tuple with the CloudStorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudStorageProvider

`func (o *BTCloudStorageObjectInfo) SetCloudStorageProvider(v int32)`

SetCloudStorageProvider sets CloudStorageProvider field to given value.

### HasCloudStorageProvider

`func (o *BTCloudStorageObjectInfo) HasCloudStorageProvider() bool`

HasCloudStorageProvider returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTCloudStorageObjectInfo) GetCreatedAt() JSONTime`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTCloudStorageObjectInfo) GetCreatedAtOk() (*JSONTime, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTCloudStorageObjectInfo) SetCreatedAt(v JSONTime)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTCloudStorageObjectInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *BTCloudStorageObjectInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTCloudStorageObjectInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTCloudStorageObjectInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTCloudStorageObjectInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *BTCloudStorageObjectInfo) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *BTCloudStorageObjectInfo) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *BTCloudStorageObjectInfo) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *BTCloudStorageObjectInfo) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetDescription

`func (o *BTCloudStorageObjectInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTCloudStorageObjectInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTCloudStorageObjectInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTCloudStorageObjectInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHref

`func (o *BTCloudStorageObjectInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTCloudStorageObjectInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTCloudStorageObjectInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTCloudStorageObjectInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetIconLink

`func (o *BTCloudStorageObjectInfo) GetIconLink() string`

GetIconLink returns the IconLink field if non-nil, zero value otherwise.

### GetIconLinkOk

`func (o *BTCloudStorageObjectInfo) GetIconLinkOk() (*string, bool)`

GetIconLinkOk returns a tuple with the IconLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconLink

`func (o *BTCloudStorageObjectInfo) SetIconLink(v string)`

SetIconLink sets IconLink field to given value.

### HasIconLink

`func (o *BTCloudStorageObjectInfo) HasIconLink() bool`

HasIconLink returns a boolean if a field has been set.

### GetId

`func (o *BTCloudStorageObjectInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTCloudStorageObjectInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTCloudStorageObjectInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTCloudStorageObjectInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsContainer

`func (o *BTCloudStorageObjectInfo) GetIsContainer() bool`

GetIsContainer returns the IsContainer field if non-nil, zero value otherwise.

### GetIsContainerOk

`func (o *BTCloudStorageObjectInfo) GetIsContainerOk() (*bool, bool)`

GetIsContainerOk returns a tuple with the IsContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsContainer

`func (o *BTCloudStorageObjectInfo) SetIsContainer(v bool)`

SetIsContainer sets IsContainer field to given value.

### HasIsContainer

`func (o *BTCloudStorageObjectInfo) HasIsContainer() bool`

HasIsContainer returns a boolean if a field has been set.

### GetIsEnterpriseOwned

`func (o *BTCloudStorageObjectInfo) GetIsEnterpriseOwned() bool`

GetIsEnterpriseOwned returns the IsEnterpriseOwned field if non-nil, zero value otherwise.

### GetIsEnterpriseOwnedOk

`func (o *BTCloudStorageObjectInfo) GetIsEnterpriseOwnedOk() (*bool, bool)`

GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnterpriseOwned

`func (o *BTCloudStorageObjectInfo) SetIsEnterpriseOwned(v bool)`

SetIsEnterpriseOwned sets IsEnterpriseOwned field to given value.

### HasIsEnterpriseOwned

`func (o *BTCloudStorageObjectInfo) HasIsEnterpriseOwned() bool`

HasIsEnterpriseOwned returns a boolean if a field has been set.

### GetIsMutable

`func (o *BTCloudStorageObjectInfo) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *BTCloudStorageObjectInfo) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *BTCloudStorageObjectInfo) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.

### HasIsMutable

`func (o *BTCloudStorageObjectInfo) HasIsMutable() bool`

HasIsMutable returns a boolean if a field has been set.

### GetMimeType

`func (o *BTCloudStorageObjectInfo) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *BTCloudStorageObjectInfo) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *BTCloudStorageObjectInfo) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *BTCloudStorageObjectInfo) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetModifiedAt

`func (o *BTCloudStorageObjectInfo) GetModifiedAt() JSONTime`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BTCloudStorageObjectInfo) GetModifiedAtOk() (*JSONTime, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BTCloudStorageObjectInfo) SetModifiedAt(v JSONTime)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *BTCloudStorageObjectInfo) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTCloudStorageObjectInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTCloudStorageObjectInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTCloudStorageObjectInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTCloudStorageObjectInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetModifiedById

`func (o *BTCloudStorageObjectInfo) GetModifiedById() string`

GetModifiedById returns the ModifiedById field if non-nil, zero value otherwise.

### GetModifiedByIdOk

`func (o *BTCloudStorageObjectInfo) GetModifiedByIdOk() (*string, bool)`

GetModifiedByIdOk returns a tuple with the ModifiedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedById

`func (o *BTCloudStorageObjectInfo) SetModifiedById(v string)`

SetModifiedById sets ModifiedById field to given value.

### HasModifiedById

`func (o *BTCloudStorageObjectInfo) HasModifiedById() bool`

HasModifiedById returns a boolean if a field has been set.

### GetName

`func (o *BTCloudStorageObjectInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTCloudStorageObjectInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTCloudStorageObjectInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTCloudStorageObjectInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *BTCloudStorageObjectInfo) GetOwner() BTOwnerInfo`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BTCloudStorageObjectInfo) GetOwnerOk() (*BTOwnerInfo, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BTCloudStorageObjectInfo) SetOwner(v BTOwnerInfo)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BTCloudStorageObjectInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentId

`func (o *BTCloudStorageObjectInfo) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTCloudStorageObjectInfo) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTCloudStorageObjectInfo) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTCloudStorageObjectInfo) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectId

`func (o *BTCloudStorageObjectInfo) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BTCloudStorageObjectInfo) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BTCloudStorageObjectInfo) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BTCloudStorageObjectInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResourceType

`func (o *BTCloudStorageObjectInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BTCloudStorageObjectInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BTCloudStorageObjectInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BTCloudStorageObjectInfo) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSizeBytes

`func (o *BTCloudStorageObjectInfo) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *BTCloudStorageObjectInfo) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *BTCloudStorageObjectInfo) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *BTCloudStorageObjectInfo) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetThumbnailInfo

`func (o *BTCloudStorageObjectInfo) GetThumbnailInfo() BTThumbnailInfo`

GetThumbnailInfo returns the ThumbnailInfo field if non-nil, zero value otherwise.

### GetThumbnailInfoOk

`func (o *BTCloudStorageObjectInfo) GetThumbnailInfoOk() (*BTThumbnailInfo, bool)`

GetThumbnailInfoOk returns a tuple with the ThumbnailInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailInfo

`func (o *BTCloudStorageObjectInfo) SetThumbnailInfo(v BTThumbnailInfo)`

SetThumbnailInfo sets ThumbnailInfo field to given value.

### HasThumbnailInfo

`func (o *BTCloudStorageObjectInfo) HasThumbnailInfo() bool`

HasThumbnailInfo returns a boolean if a field has been set.

### GetTreeHref

`func (o *BTCloudStorageObjectInfo) GetTreeHref() string`

GetTreeHref returns the TreeHref field if non-nil, zero value otherwise.

### GetTreeHrefOk

`func (o *BTCloudStorageObjectInfo) GetTreeHrefOk() (*string, bool)`

GetTreeHrefOk returns a tuple with the TreeHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreeHref

`func (o *BTCloudStorageObjectInfo) SetTreeHref(v string)`

SetTreeHref sets TreeHref field to given value.

### HasTreeHref

`func (o *BTCloudStorageObjectInfo) HasTreeHref() bool`

HasTreeHref returns a boolean if a field has been set.

### GetUnparentHref

`func (o *BTCloudStorageObjectInfo) GetUnparentHref() string`

GetUnparentHref returns the UnparentHref field if non-nil, zero value otherwise.

### GetUnparentHrefOk

`func (o *BTCloudStorageObjectInfo) GetUnparentHrefOk() (*string, bool)`

GetUnparentHrefOk returns a tuple with the UnparentHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnparentHref

`func (o *BTCloudStorageObjectInfo) SetUnparentHref(v string)`

SetUnparentHref sets UnparentHref field to given value.

### HasUnparentHref

`func (o *BTCloudStorageObjectInfo) HasUnparentHref() bool`

HasUnparentHref returns a boolean if a field has been set.

### GetViewRef

`func (o *BTCloudStorageObjectInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTCloudStorageObjectInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTCloudStorageObjectInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTCloudStorageObjectInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWebViewLink

`func (o *BTCloudStorageObjectInfo) GetWebViewLink() string`

GetWebViewLink returns the WebViewLink field if non-nil, zero value otherwise.

### GetWebViewLinkOk

`func (o *BTCloudStorageObjectInfo) GetWebViewLinkOk() (*string, bool)`

GetWebViewLinkOk returns a tuple with the WebViewLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebViewLink

`func (o *BTCloudStorageObjectInfo) SetWebViewLink(v string)`

SetWebViewLink sets WebViewLink field to given value.

### HasWebViewLink

`func (o *BTCloudStorageObjectInfo) HasWebViewLink() bool`

HasWebViewLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


