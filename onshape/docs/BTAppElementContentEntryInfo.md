# BTAppElementContentEntryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseContent** | Pointer to **string** |  | [optional] 
**Deltas** | Pointer to [**[]BTAppElementContentDeltaInfo**](BTAppElementContentDeltaInfo.md) |  | [optional] 
**SubelementId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTAppElementContentEntryInfo

`func NewBTAppElementContentEntryInfo() *BTAppElementContentEntryInfo`

NewBTAppElementContentEntryInfo instantiates a new BTAppElementContentEntryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementContentEntryInfoWithDefaults

`func NewBTAppElementContentEntryInfoWithDefaults() *BTAppElementContentEntryInfo`

NewBTAppElementContentEntryInfoWithDefaults instantiates a new BTAppElementContentEntryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseContent

`func (o *BTAppElementContentEntryInfo) GetBaseContent() string`

GetBaseContent returns the BaseContent field if non-nil, zero value otherwise.

### GetBaseContentOk

`func (o *BTAppElementContentEntryInfo) GetBaseContentOk() (*string, bool)`

GetBaseContentOk returns a tuple with the BaseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseContent

`func (o *BTAppElementContentEntryInfo) SetBaseContent(v string)`

SetBaseContent sets BaseContent field to given value.

### HasBaseContent

`func (o *BTAppElementContentEntryInfo) HasBaseContent() bool`

HasBaseContent returns a boolean if a field has been set.

### GetDeltas

`func (o *BTAppElementContentEntryInfo) GetDeltas() []BTAppElementContentDeltaInfo`

GetDeltas returns the Deltas field if non-nil, zero value otherwise.

### GetDeltasOk

`func (o *BTAppElementContentEntryInfo) GetDeltasOk() (*[]BTAppElementContentDeltaInfo, bool)`

GetDeltasOk returns a tuple with the Deltas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltas

`func (o *BTAppElementContentEntryInfo) SetDeltas(v []BTAppElementContentDeltaInfo)`

SetDeltas sets Deltas field to given value.

### HasDeltas

`func (o *BTAppElementContentEntryInfo) HasDeltas() bool`

HasDeltas returns a boolean if a field has been set.

### GetSubelementId

`func (o *BTAppElementContentEntryInfo) GetSubelementId() string`

GetSubelementId returns the SubelementId field if non-nil, zero value otherwise.

### GetSubelementIdOk

`func (o *BTAppElementContentEntryInfo) GetSubelementIdOk() (*string, bool)`

GetSubelementIdOk returns a tuple with the SubelementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelementId

`func (o *BTAppElementContentEntryInfo) SetSubelementId(v string)`

SetSubelementId sets SubelementId field to given value.

### HasSubelementId

`func (o *BTAppElementContentEntryInfo) HasSubelementId() bool`

HasSubelementId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


