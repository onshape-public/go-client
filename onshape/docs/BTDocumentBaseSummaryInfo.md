# BTDocumentBaseSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**ModifiedBy** | Pointer to [**BTUserBasicSummaryInfo**](BTUserBasicSummaryInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**StorageSize** | Pointer to **int64** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTDocumentBaseSummaryInfo

`func NewBTDocumentBaseSummaryInfo() *BTDocumentBaseSummaryInfo`

NewBTDocumentBaseSummaryInfo instantiates a new BTDocumentBaseSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentBaseSummaryInfoWithDefaults

`func NewBTDocumentBaseSummaryInfoWithDefaults() *BTDocumentBaseSummaryInfo`

NewBTDocumentBaseSummaryInfoWithDefaults instantiates a new BTDocumentBaseSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *BTDocumentBaseSummaryInfo) GetCreatedBy() BTUserBasicSummaryInfo`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BTDocumentBaseSummaryInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BTDocumentBaseSummaryInfo) SetCreatedBy(v BTUserBasicSummaryInfo)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *BTDocumentBaseSummaryInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHref

`func (o *BTDocumentBaseSummaryInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTDocumentBaseSummaryInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTDocumentBaseSummaryInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTDocumentBaseSummaryInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTDocumentBaseSummaryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTDocumentBaseSummaryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTDocumentBaseSummaryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTDocumentBaseSummaryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifiedBy

`func (o *BTDocumentBaseSummaryInfo) GetModifiedBy() BTUserBasicSummaryInfo`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BTDocumentBaseSummaryInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BTDocumentBaseSummaryInfo) SetModifiedBy(v BTUserBasicSummaryInfo)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *BTDocumentBaseSummaryInfo) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *BTDocumentBaseSummaryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTDocumentBaseSummaryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTDocumentBaseSummaryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTDocumentBaseSummaryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *BTDocumentBaseSummaryInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTDocumentBaseSummaryInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTDocumentBaseSummaryInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTDocumentBaseSummaryInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageSize

`func (o *BTDocumentBaseSummaryInfo) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *BTDocumentBaseSummaryInfo) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *BTDocumentBaseSummaryInfo) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *BTDocumentBaseSummaryInfo) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.

### GetViewRef

`func (o *BTDocumentBaseSummaryInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTDocumentBaseSummaryInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTDocumentBaseSummaryInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTDocumentBaseSummaryInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


