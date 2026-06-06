# BTWhereUsedItemInfoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultConfig** | Pointer to [**[]ConfigInfo**](ConfigInfo.md) | The resolved default configuration parameters for the queried element. Only populated when the configuration query parameter is set to &#39;default&#39;. | [optional] 
**Href** | Pointer to **string** | URI for current page of resources. | [optional] 
**Items** | Pointer to [**[]BTProductStructureItemInfo**](BTProductStructureItemInfo.md) | Array of items in the current page. | [optional] 
**Next** | Pointer to **string** | URI for next page of the resources if more are available. | [optional] 
**Previous** | Pointer to **string** | URI for previous page of the resources. | [optional] 
**Versions** | Pointer to [**[]BTWhereUsedVersionReferencedInfo**](BTWhereUsedVersionReferencedInfo.md) | List of document versions in which the queried item is referenced. Each entry contains the version info, referring document and element details, configuration, revision, and part identifiers. Only populated when includeVersionInfo is true. | [optional] 

## Methods

### NewBTWhereUsedItemInfoList

`func NewBTWhereUsedItemInfoList() *BTWhereUsedItemInfoList`

NewBTWhereUsedItemInfoList instantiates a new BTWhereUsedItemInfoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWhereUsedItemInfoListWithDefaults

`func NewBTWhereUsedItemInfoListWithDefaults() *BTWhereUsedItemInfoList`

NewBTWhereUsedItemInfoListWithDefaults instantiates a new BTWhereUsedItemInfoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultConfig

`func (o *BTWhereUsedItemInfoList) GetDefaultConfig() []ConfigInfo`

GetDefaultConfig returns the DefaultConfig field if non-nil, zero value otherwise.

### GetDefaultConfigOk

`func (o *BTWhereUsedItemInfoList) GetDefaultConfigOk() (*[]ConfigInfo, bool)`

GetDefaultConfigOk returns a tuple with the DefaultConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfig

`func (o *BTWhereUsedItemInfoList) SetDefaultConfig(v []ConfigInfo)`

SetDefaultConfig sets DefaultConfig field to given value.

### HasDefaultConfig

`func (o *BTWhereUsedItemInfoList) HasDefaultConfig() bool`

HasDefaultConfig returns a boolean if a field has been set.

### GetHref

`func (o *BTWhereUsedItemInfoList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTWhereUsedItemInfoList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTWhereUsedItemInfoList) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTWhereUsedItemInfoList) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *BTWhereUsedItemInfoList) GetItems() []BTProductStructureItemInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BTWhereUsedItemInfoList) GetItemsOk() (*[]BTProductStructureItemInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BTWhereUsedItemInfoList) SetItems(v []BTProductStructureItemInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *BTWhereUsedItemInfoList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNext

`func (o *BTWhereUsedItemInfoList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *BTWhereUsedItemInfoList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *BTWhereUsedItemInfoList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *BTWhereUsedItemInfoList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrevious

`func (o *BTWhereUsedItemInfoList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *BTWhereUsedItemInfoList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *BTWhereUsedItemInfoList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *BTWhereUsedItemInfoList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### GetVersions

`func (o *BTWhereUsedItemInfoList) GetVersions() []BTWhereUsedVersionReferencedInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BTWhereUsedItemInfoList) GetVersionsOk() (*[]BTWhereUsedVersionReferencedInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BTWhereUsedItemInfoList) SetVersions(v []BTWhereUsedVersionReferencedInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *BTWhereUsedItemInfoList) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


