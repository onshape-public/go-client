# BTBillOfMaterialsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BomSource** | Pointer to [**BTBillOfMaterialsSourceInfo**](BTBillOfMaterialsSourceInfo.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**FormatVersion** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**[]BTBillOfMaterialsHeaderInfo**](BTBillOfMaterialsHeaderInfo.md) |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]BTBillOfMaterialsRowInfo**](BTBillOfMaterialsRowInfo.md) |  | [optional] 
**TemplateId** | Pointer to **string** |  | [optional] 
**TopLevelAssemblyRow** | Pointer to [**BTBillOfMaterialsRowInfo**](BTBillOfMaterialsRowInfo.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTBillOfMaterialsInfo

`func NewBTBillOfMaterialsInfo() *BTBillOfMaterialsInfo`

NewBTBillOfMaterialsInfo instantiates a new BTBillOfMaterialsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTBillOfMaterialsInfoWithDefaults

`func NewBTBillOfMaterialsInfoWithDefaults() *BTBillOfMaterialsInfo`

NewBTBillOfMaterialsInfoWithDefaults instantiates a new BTBillOfMaterialsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBomSource

`func (o *BTBillOfMaterialsInfo) GetBomSource() BTBillOfMaterialsSourceInfo`

GetBomSource returns the BomSource field if non-nil, zero value otherwise.

### GetBomSourceOk

`func (o *BTBillOfMaterialsInfo) GetBomSourceOk() (*BTBillOfMaterialsSourceInfo, bool)`

GetBomSourceOk returns a tuple with the BomSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomSource

`func (o *BTBillOfMaterialsInfo) SetBomSource(v BTBillOfMaterialsSourceInfo)`

SetBomSource sets BomSource field to given value.

### HasBomSource

`func (o *BTBillOfMaterialsInfo) HasBomSource() bool`

HasBomSource returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BTBillOfMaterialsInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BTBillOfMaterialsInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BTBillOfMaterialsInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BTBillOfMaterialsInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFormatVersion

`func (o *BTBillOfMaterialsInfo) GetFormatVersion() string`

GetFormatVersion returns the FormatVersion field if non-nil, zero value otherwise.

### GetFormatVersionOk

`func (o *BTBillOfMaterialsInfo) GetFormatVersionOk() (*string, bool)`

GetFormatVersionOk returns a tuple with the FormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatVersion

`func (o *BTBillOfMaterialsInfo) SetFormatVersion(v string)`

SetFormatVersion sets FormatVersion field to given value.

### HasFormatVersion

`func (o *BTBillOfMaterialsInfo) HasFormatVersion() bool`

HasFormatVersion returns a boolean if a field has been set.

### GetHeaders

`func (o *BTBillOfMaterialsInfo) GetHeaders() []BTBillOfMaterialsHeaderInfo`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BTBillOfMaterialsInfo) GetHeadersOk() (*[]BTBillOfMaterialsHeaderInfo, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BTBillOfMaterialsInfo) SetHeaders(v []BTBillOfMaterialsHeaderInfo)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BTBillOfMaterialsInfo) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHref

`func (o *BTBillOfMaterialsInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTBillOfMaterialsInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTBillOfMaterialsInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTBillOfMaterialsInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTBillOfMaterialsInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTBillOfMaterialsInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTBillOfMaterialsInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTBillOfMaterialsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTBillOfMaterialsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTBillOfMaterialsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTBillOfMaterialsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTBillOfMaterialsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRows

`func (o *BTBillOfMaterialsInfo) GetRows() []BTBillOfMaterialsRowInfo`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *BTBillOfMaterialsInfo) GetRowsOk() (*[]BTBillOfMaterialsRowInfo, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *BTBillOfMaterialsInfo) SetRows(v []BTBillOfMaterialsRowInfo)`

SetRows sets Rows field to given value.

### HasRows

`func (o *BTBillOfMaterialsInfo) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTemplateId

`func (o *BTBillOfMaterialsInfo) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *BTBillOfMaterialsInfo) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *BTBillOfMaterialsInfo) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *BTBillOfMaterialsInfo) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTopLevelAssemblyRow

`func (o *BTBillOfMaterialsInfo) GetTopLevelAssemblyRow() BTBillOfMaterialsRowInfo`

GetTopLevelAssemblyRow returns the TopLevelAssemblyRow field if non-nil, zero value otherwise.

### GetTopLevelAssemblyRowOk

`func (o *BTBillOfMaterialsInfo) GetTopLevelAssemblyRowOk() (*BTBillOfMaterialsRowInfo, bool)`

GetTopLevelAssemblyRowOk returns a tuple with the TopLevelAssemblyRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelAssemblyRow

`func (o *BTBillOfMaterialsInfo) SetTopLevelAssemblyRow(v BTBillOfMaterialsRowInfo)`

SetTopLevelAssemblyRow sets TopLevelAssemblyRow field to given value.

### HasTopLevelAssemblyRow

`func (o *BTBillOfMaterialsInfo) HasTopLevelAssemblyRow() bool`

HasTopLevelAssemblyRow returns a boolean if a field has been set.

### GetType

`func (o *BTBillOfMaterialsInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTBillOfMaterialsInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTBillOfMaterialsInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTBillOfMaterialsInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetViewRef

`func (o *BTBillOfMaterialsInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTBillOfMaterialsInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTBillOfMaterialsInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTBillOfMaterialsInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


