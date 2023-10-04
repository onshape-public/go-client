# BTElementLibrarySummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**LibraryId** | Pointer to **string** | The Id of the library -- unique across Onshape. | [optional] 
**LibraryVersion** | Pointer to **string** | The current version Id of the library. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**OwnerId** | Pointer to **string** | The owner Id of an element library (either Onshape, company, or user). | [optional] 
**OwnerType** | Pointer to **int32** | The type of library owner, Onshape, user, or company | [optional] 
**Purpose** | Pointer to **string** | The purpose string identifying the type of element library. | [optional] 
**SourceFolderId** | Pointer to **string** | The id of the root folder of the library | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTElementLibrarySummaryInfo

`func NewBTElementLibrarySummaryInfo() *BTElementLibrarySummaryInfo`

NewBTElementLibrarySummaryInfo instantiates a new BTElementLibrarySummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTElementLibrarySummaryInfoWithDefaults

`func NewBTElementLibrarySummaryInfoWithDefaults() *BTElementLibrarySummaryInfo`

NewBTElementLibrarySummaryInfoWithDefaults instantiates a new BTElementLibrarySummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BTElementLibrarySummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTElementLibrarySummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTElementLibrarySummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTElementLibrarySummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTElementLibrarySummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTElementLibrarySummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTElementLibrarySummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTElementLibrarySummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibraryId

`func (o *BTElementLibrarySummaryInfo) GetLibraryId() string`

GetLibraryId returns the LibraryId field if non-nil, zero value otherwise.

### GetLibraryIdOk

`func (o *BTElementLibrarySummaryInfo) GetLibraryIdOk() (*string, bool)`

GetLibraryIdOk returns a tuple with the LibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryId

`func (o *BTElementLibrarySummaryInfo) SetLibraryId(v string)`

SetLibraryId sets LibraryId field to given value.

### HasLibraryId

`func (o *BTElementLibrarySummaryInfo) HasLibraryId() bool`

HasLibraryId returns a boolean if a field has been set.

### GetLibraryVersion

`func (o *BTElementLibrarySummaryInfo) GetLibraryVersion() string`

GetLibraryVersion returns the LibraryVersion field if non-nil, zero value otherwise.

### GetLibraryVersionOk

`func (o *BTElementLibrarySummaryInfo) GetLibraryVersionOk() (*string, bool)`

GetLibraryVersionOk returns a tuple with the LibraryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryVersion

`func (o *BTElementLibrarySummaryInfo) SetLibraryVersion(v string)`

SetLibraryVersion sets LibraryVersion field to given value.

### HasLibraryVersion

`func (o *BTElementLibrarySummaryInfo) HasLibraryVersion() bool`

HasLibraryVersion returns a boolean if a field has been set.

### GetName

`func (o *BTElementLibrarySummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTElementLibrarySummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTElementLibrarySummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTElementLibrarySummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTElementLibrarySummaryInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTElementLibrarySummaryInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTElementLibrarySummaryInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTElementLibrarySummaryInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTElementLibrarySummaryInfo) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTElementLibrarySummaryInfo) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTElementLibrarySummaryInfo) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTElementLibrarySummaryInfo) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPurpose

`func (o *BTElementLibrarySummaryInfo) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *BTElementLibrarySummaryInfo) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *BTElementLibrarySummaryInfo) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *BTElementLibrarySummaryInfo) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSourceFolderId

`func (o *BTElementLibrarySummaryInfo) GetSourceFolderId() string`

GetSourceFolderId returns the SourceFolderId field if non-nil, zero value otherwise.

### GetSourceFolderIdOk

`func (o *BTElementLibrarySummaryInfo) GetSourceFolderIdOk() (*string, bool)`

GetSourceFolderIdOk returns a tuple with the SourceFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFolderId

`func (o *BTElementLibrarySummaryInfo) SetSourceFolderId(v string)`

SetSourceFolderId sets SourceFolderId field to given value.

### HasSourceFolderId

`func (o *BTElementLibrarySummaryInfo) HasSourceFolderId() bool`

HasSourceFolderId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTElementLibrarySummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTElementLibrarySummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTElementLibrarySummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTElementLibrarySummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


