# BTSmartFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectType** | Pointer to **int32** |  | [optional] 
**Predicates** | Pointer to [**[]BTSmartFolderPredicateInfo**](BTSmartFolderPredicateInfo.md) |  | [optional] 
**Sort** | Pointer to [**[]BTSmartFolderSortInfo**](BTSmartFolderSortInfo.md) |  | [optional] 

## Methods

### NewBTSmartFolderInfo

`func NewBTSmartFolderInfo() *BTSmartFolderInfo`

NewBTSmartFolderInfo instantiates a new BTSmartFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTSmartFolderInfoWithDefaults

`func NewBTSmartFolderInfoWithDefaults() *BTSmartFolderInfo`

NewBTSmartFolderInfoWithDefaults instantiates a new BTSmartFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectType

`func (o *BTSmartFolderInfo) GetObjectType() int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTSmartFolderInfo) GetObjectTypeOk() (*int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTSmartFolderInfo) SetObjectType(v int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTSmartFolderInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPredicates

`func (o *BTSmartFolderInfo) GetPredicates() []BTSmartFolderPredicateInfo`

GetPredicates returns the Predicates field if non-nil, zero value otherwise.

### GetPredicatesOk

`func (o *BTSmartFolderInfo) GetPredicatesOk() (*[]BTSmartFolderPredicateInfo, bool)`

GetPredicatesOk returns a tuple with the Predicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicates

`func (o *BTSmartFolderInfo) SetPredicates(v []BTSmartFolderPredicateInfo)`

SetPredicates sets Predicates field to given value.

### HasPredicates

`func (o *BTSmartFolderInfo) HasPredicates() bool`

HasPredicates returns a boolean if a field has been set.

### GetSort

`func (o *BTSmartFolderInfo) GetSort() []BTSmartFolderSortInfo`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *BTSmartFolderInfo) GetSortOk() (*[]BTSmartFolderSortInfo, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *BTSmartFolderInfo) SetSort(v []BTSmartFolderSortInfo)`

SetSort sets Sort field to given value.

### HasSort

`func (o *BTSmartFolderInfo) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


