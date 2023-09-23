# BTUserMetricsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDocuments** | Pointer to **int64** |  | [optional] 
**HasRecentlyOpenedDocuments** | Pointer to **bool** |  | [optional] 
**HasSharedDocuments** | Pointer to **bool** |  | [optional] 
**HasTrashedDocuments** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**PrivateDocuments** | Pointer to **int64** |  | [optional] 
**PublicDocuments** | Pointer to **int64** |  | [optional] 
**SharedDocuments** | Pointer to **int64** |  | [optional] 
**UserAccountLimitsCrossed** | Pointer to **bool** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTUserMetricsInfo

`func NewBTUserMetricsInfo() *BTUserMetricsInfo`

NewBTUserMetricsInfo instantiates a new BTUserMetricsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTUserMetricsInfoWithDefaults

`func NewBTUserMetricsInfoWithDefaults() *BTUserMetricsInfo`

NewBTUserMetricsInfoWithDefaults instantiates a new BTUserMetricsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDocuments

`func (o *BTUserMetricsInfo) GetCreatedDocuments() int64`

GetCreatedDocuments returns the CreatedDocuments field if non-nil, zero value otherwise.

### GetCreatedDocumentsOk

`func (o *BTUserMetricsInfo) GetCreatedDocumentsOk() (*int64, bool)`

GetCreatedDocumentsOk returns a tuple with the CreatedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDocuments

`func (o *BTUserMetricsInfo) SetCreatedDocuments(v int64)`

SetCreatedDocuments sets CreatedDocuments field to given value.

### HasCreatedDocuments

`func (o *BTUserMetricsInfo) HasCreatedDocuments() bool`

HasCreatedDocuments returns a boolean if a field has been set.

### GetHasRecentlyOpenedDocuments

`func (o *BTUserMetricsInfo) GetHasRecentlyOpenedDocuments() bool`

GetHasRecentlyOpenedDocuments returns the HasRecentlyOpenedDocuments field if non-nil, zero value otherwise.

### GetHasRecentlyOpenedDocumentsOk

`func (o *BTUserMetricsInfo) GetHasRecentlyOpenedDocumentsOk() (*bool, bool)`

GetHasRecentlyOpenedDocumentsOk returns a tuple with the HasRecentlyOpenedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRecentlyOpenedDocuments

`func (o *BTUserMetricsInfo) SetHasRecentlyOpenedDocuments(v bool)`

SetHasRecentlyOpenedDocuments sets HasRecentlyOpenedDocuments field to given value.

### HasHasRecentlyOpenedDocuments

`func (o *BTUserMetricsInfo) HasHasRecentlyOpenedDocuments() bool`

HasHasRecentlyOpenedDocuments returns a boolean if a field has been set.

### GetHasSharedDocuments

`func (o *BTUserMetricsInfo) GetHasSharedDocuments() bool`

GetHasSharedDocuments returns the HasSharedDocuments field if non-nil, zero value otherwise.

### GetHasSharedDocumentsOk

`func (o *BTUserMetricsInfo) GetHasSharedDocumentsOk() (*bool, bool)`

GetHasSharedDocumentsOk returns a tuple with the HasSharedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSharedDocuments

`func (o *BTUserMetricsInfo) SetHasSharedDocuments(v bool)`

SetHasSharedDocuments sets HasSharedDocuments field to given value.

### HasHasSharedDocuments

`func (o *BTUserMetricsInfo) HasHasSharedDocuments() bool`

HasHasSharedDocuments returns a boolean if a field has been set.

### GetHasTrashedDocuments

`func (o *BTUserMetricsInfo) GetHasTrashedDocuments() bool`

GetHasTrashedDocuments returns the HasTrashedDocuments field if non-nil, zero value otherwise.

### GetHasTrashedDocumentsOk

`func (o *BTUserMetricsInfo) GetHasTrashedDocumentsOk() (*bool, bool)`

GetHasTrashedDocumentsOk returns a tuple with the HasTrashedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTrashedDocuments

`func (o *BTUserMetricsInfo) SetHasTrashedDocuments(v bool)`

SetHasTrashedDocuments sets HasTrashedDocuments field to given value.

### HasHasTrashedDocuments

`func (o *BTUserMetricsInfo) HasHasTrashedDocuments() bool`

HasHasTrashedDocuments returns a boolean if a field has been set.

### GetHref

`func (o *BTUserMetricsInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTUserMetricsInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTUserMetricsInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTUserMetricsInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTUserMetricsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTUserMetricsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTUserMetricsInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTUserMetricsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTUserMetricsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTUserMetricsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTUserMetricsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTUserMetricsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivateDocuments

`func (o *BTUserMetricsInfo) GetPrivateDocuments() int64`

GetPrivateDocuments returns the PrivateDocuments field if non-nil, zero value otherwise.

### GetPrivateDocumentsOk

`func (o *BTUserMetricsInfo) GetPrivateDocumentsOk() (*int64, bool)`

GetPrivateDocumentsOk returns a tuple with the PrivateDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDocuments

`func (o *BTUserMetricsInfo) SetPrivateDocuments(v int64)`

SetPrivateDocuments sets PrivateDocuments field to given value.

### HasPrivateDocuments

`func (o *BTUserMetricsInfo) HasPrivateDocuments() bool`

HasPrivateDocuments returns a boolean if a field has been set.

### GetPublicDocuments

`func (o *BTUserMetricsInfo) GetPublicDocuments() int64`

GetPublicDocuments returns the PublicDocuments field if non-nil, zero value otherwise.

### GetPublicDocumentsOk

`func (o *BTUserMetricsInfo) GetPublicDocumentsOk() (*int64, bool)`

GetPublicDocumentsOk returns a tuple with the PublicDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDocuments

`func (o *BTUserMetricsInfo) SetPublicDocuments(v int64)`

SetPublicDocuments sets PublicDocuments field to given value.

### HasPublicDocuments

`func (o *BTUserMetricsInfo) HasPublicDocuments() bool`

HasPublicDocuments returns a boolean if a field has been set.

### GetSharedDocuments

`func (o *BTUserMetricsInfo) GetSharedDocuments() int64`

GetSharedDocuments returns the SharedDocuments field if non-nil, zero value otherwise.

### GetSharedDocumentsOk

`func (o *BTUserMetricsInfo) GetSharedDocumentsOk() (*int64, bool)`

GetSharedDocumentsOk returns a tuple with the SharedDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDocuments

`func (o *BTUserMetricsInfo) SetSharedDocuments(v int64)`

SetSharedDocuments sets SharedDocuments field to given value.

### HasSharedDocuments

`func (o *BTUserMetricsInfo) HasSharedDocuments() bool`

HasSharedDocuments returns a boolean if a field has been set.

### GetUserAccountLimitsCrossed

`func (o *BTUserMetricsInfo) GetUserAccountLimitsCrossed() bool`

GetUserAccountLimitsCrossed returns the UserAccountLimitsCrossed field if non-nil, zero value otherwise.

### GetUserAccountLimitsCrossedOk

`func (o *BTUserMetricsInfo) GetUserAccountLimitsCrossedOk() (*bool, bool)`

GetUserAccountLimitsCrossedOk returns a tuple with the UserAccountLimitsCrossed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountLimitsCrossed

`func (o *BTUserMetricsInfo) SetUserAccountLimitsCrossed(v bool)`

SetUserAccountLimitsCrossed sets UserAccountLimitsCrossed field to given value.

### HasUserAccountLimitsCrossed

`func (o *BTUserMetricsInfo) HasUserAccountLimitsCrossed() bool`

HasUserAccountLimitsCrossed returns a boolean if a field has been set.

### GetViewRef

`func (o *BTUserMetricsInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTUserMetricsInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTUserMetricsInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTUserMetricsInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


