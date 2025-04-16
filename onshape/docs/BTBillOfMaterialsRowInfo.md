# BTBillOfMaterialsRowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderIdToValue** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**IndentLevel** | Pointer to **int32** |  | [optional] 
**ItemSource** | Pointer to [**BTBillOfMaterialsItemSourceInfo**](BTBillOfMaterialsItemSourceInfo.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**RelatedOccurrences** | Pointer to **[]string** | Occurrence IDs in the assembly that refer to the part described by this BOM row. | [optional] 
**RowId** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTBillOfMaterialsRowInfo

`func NewBTBillOfMaterialsRowInfo() *BTBillOfMaterialsRowInfo`

NewBTBillOfMaterialsRowInfo instantiates a new BTBillOfMaterialsRowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBillOfMaterialsRowInfoWithDefaults

`func NewBTBillOfMaterialsRowInfoWithDefaults() *BTBillOfMaterialsRowInfo`

NewBTBillOfMaterialsRowInfoWithDefaults instantiates a new BTBillOfMaterialsRowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderIdToValue

`func (o *BTBillOfMaterialsRowInfo) GetHeaderIdToValue() map[string]interface{}`

GetHeaderIdToValue returns the HeaderIdToValue field if non-nil, zero value otherwise.

### GetHeaderIdToValueOk

`func (o *BTBillOfMaterialsRowInfo) GetHeaderIdToValueOk() (*map[string]interface{}, bool)`

GetHeaderIdToValueOk returns a tuple with the HeaderIdToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderIdToValue

`func (o *BTBillOfMaterialsRowInfo) SetHeaderIdToValue(v map[string]interface{})`

SetHeaderIdToValue sets HeaderIdToValue field to given value.

### HasHeaderIdToValue

`func (o *BTBillOfMaterialsRowInfo) HasHeaderIdToValue() bool`

HasHeaderIdToValue returns a boolean if a field has been set.

### GetHref

`func (o *BTBillOfMaterialsRowInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTBillOfMaterialsRowInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTBillOfMaterialsRowInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTBillOfMaterialsRowInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTBillOfMaterialsRowInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTBillOfMaterialsRowInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTBillOfMaterialsRowInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTBillOfMaterialsRowInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndentLevel

`func (o *BTBillOfMaterialsRowInfo) GetIndentLevel() int32`

GetIndentLevel returns the IndentLevel field if non-nil, zero value otherwise.

### GetIndentLevelOk

`func (o *BTBillOfMaterialsRowInfo) GetIndentLevelOk() (*int32, bool)`

GetIndentLevelOk returns a tuple with the IndentLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentLevel

`func (o *BTBillOfMaterialsRowInfo) SetIndentLevel(v int32)`

SetIndentLevel sets IndentLevel field to given value.

### HasIndentLevel

`func (o *BTBillOfMaterialsRowInfo) HasIndentLevel() bool`

HasIndentLevel returns a boolean if a field has been set.

### GetItemSource

`func (o *BTBillOfMaterialsRowInfo) GetItemSource() BTBillOfMaterialsItemSourceInfo`

GetItemSource returns the ItemSource field if non-nil, zero value otherwise.

### GetItemSourceOk

`func (o *BTBillOfMaterialsRowInfo) GetItemSourceOk() (*BTBillOfMaterialsItemSourceInfo, bool)`

GetItemSourceOk returns a tuple with the ItemSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSource

`func (o *BTBillOfMaterialsRowInfo) SetItemSource(v BTBillOfMaterialsItemSourceInfo)`

SetItemSource sets ItemSource field to given value.

### HasItemSource

`func (o *BTBillOfMaterialsRowInfo) HasItemSource() bool`

HasItemSource returns a boolean if a field has been set.

### GetName

`func (o *BTBillOfMaterialsRowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTBillOfMaterialsRowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTBillOfMaterialsRowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTBillOfMaterialsRowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRelatedOccurrences

`func (o *BTBillOfMaterialsRowInfo) GetRelatedOccurrences() []string`

GetRelatedOccurrences returns the RelatedOccurrences field if non-nil, zero value otherwise.

### GetRelatedOccurrencesOk

`func (o *BTBillOfMaterialsRowInfo) GetRelatedOccurrencesOk() (*[]string, bool)`

GetRelatedOccurrencesOk returns a tuple with the RelatedOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedOccurrences

`func (o *BTBillOfMaterialsRowInfo) SetRelatedOccurrences(v []string)`

SetRelatedOccurrences sets RelatedOccurrences field to given value.

### HasRelatedOccurrences

`func (o *BTBillOfMaterialsRowInfo) HasRelatedOccurrences() bool`

HasRelatedOccurrences returns a boolean if a field has been set.

### GetRowId

`func (o *BTBillOfMaterialsRowInfo) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *BTBillOfMaterialsRowInfo) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *BTBillOfMaterialsRowInfo) SetRowId(v string)`

SetRowId sets RowId field to given value.

### HasRowId

`func (o *BTBillOfMaterialsRowInfo) HasRowId() bool`

HasRowId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTBillOfMaterialsRowInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTBillOfMaterialsRowInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTBillOfMaterialsRowInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTBillOfMaterialsRowInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


