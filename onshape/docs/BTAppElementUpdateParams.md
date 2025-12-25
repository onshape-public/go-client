# BTAppElementUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The label that will appear in the document&#39;s edit history for this operation. If blank, a value will be auto-generated. | [optional] 
**JsonPatch** | Pointer to **string** | A json patch that will be applied to the application element&#39;s json data. The JSON patch format is as specified in RFC 6902 from the IETF. | [optional] 
**JsonTreeEdit** | Pointer to [**BTJEdit3734**](BTJEdit3734.md) |  | [optional] 
**ParentChangeId** | Pointer to **string** | The id of the last change made to this application element. This can be retrieved from the response for any app element modification endpoint. | [optional] 
**PropertyUpdates** | Pointer to [**[]BTMetadataPropertyUpdateParams**](BTMetadataPropertyUpdateParams.md) | Edits to be applied to the element&#39;s metadata. | [optional] 
**ReturnError** | Pointer to **bool** | If true, errors in request processing will be returned in a 200 response with a meaningful description. Otherwise, errors will result in a relevant HTTP error response. | [optional] 
**ReturnJsonDifferenceFormat** | Pointer to **string** | If specified, and jsonTreeEdit is non-empty, the json difference will be returned in this format, otherwise no json difference will be returned. | [optional] 
**TransactionId** | Pointer to **string** | The id of the transaction in which this operation should take place. Transaction ids can be generated through the AppElement startTransaction API. | [optional] 

## Methods

### NewBTAppElementUpdateParams

`func NewBTAppElementUpdateParams() *BTAppElementUpdateParams`

NewBTAppElementUpdateParams instantiates a new BTAppElementUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementUpdateParamsWithDefaults

`func NewBTAppElementUpdateParamsWithDefaults() *BTAppElementUpdateParams`

NewBTAppElementUpdateParamsWithDefaults instantiates a new BTAppElementUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BTAppElementUpdateParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAppElementUpdateParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAppElementUpdateParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAppElementUpdateParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJsonPatch

`func (o *BTAppElementUpdateParams) GetJsonPatch() string`

GetJsonPatch returns the JsonPatch field if non-nil, zero value otherwise.

### GetJsonPatchOk

`func (o *BTAppElementUpdateParams) GetJsonPatchOk() (*string, bool)`

GetJsonPatchOk returns a tuple with the JsonPatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPatch

`func (o *BTAppElementUpdateParams) SetJsonPatch(v string)`

SetJsonPatch sets JsonPatch field to given value.

### HasJsonPatch

`func (o *BTAppElementUpdateParams) HasJsonPatch() bool`

HasJsonPatch returns a boolean if a field has been set.

### GetJsonTreeEdit

`func (o *BTAppElementUpdateParams) GetJsonTreeEdit() BTJEdit3734`

GetJsonTreeEdit returns the JsonTreeEdit field if non-nil, zero value otherwise.

### GetJsonTreeEditOk

`func (o *BTAppElementUpdateParams) GetJsonTreeEditOk() (*BTJEdit3734, bool)`

GetJsonTreeEditOk returns a tuple with the JsonTreeEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonTreeEdit

`func (o *BTAppElementUpdateParams) SetJsonTreeEdit(v BTJEdit3734)`

SetJsonTreeEdit sets JsonTreeEdit field to given value.

### HasJsonTreeEdit

`func (o *BTAppElementUpdateParams) HasJsonTreeEdit() bool`

HasJsonTreeEdit returns a boolean if a field has been set.

### GetParentChangeId

`func (o *BTAppElementUpdateParams) GetParentChangeId() string`

GetParentChangeId returns the ParentChangeId field if non-nil, zero value otherwise.

### GetParentChangeIdOk

`func (o *BTAppElementUpdateParams) GetParentChangeIdOk() (*string, bool)`

GetParentChangeIdOk returns a tuple with the ParentChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChangeId

`func (o *BTAppElementUpdateParams) SetParentChangeId(v string)`

SetParentChangeId sets ParentChangeId field to given value.

### HasParentChangeId

`func (o *BTAppElementUpdateParams) HasParentChangeId() bool`

HasParentChangeId returns a boolean if a field has been set.

### GetPropertyUpdates

`func (o *BTAppElementUpdateParams) GetPropertyUpdates() []BTMetadataPropertyUpdateParams`

GetPropertyUpdates returns the PropertyUpdates field if non-nil, zero value otherwise.

### GetPropertyUpdatesOk

`func (o *BTAppElementUpdateParams) GetPropertyUpdatesOk() (*[]BTMetadataPropertyUpdateParams, bool)`

GetPropertyUpdatesOk returns a tuple with the PropertyUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUpdates

`func (o *BTAppElementUpdateParams) SetPropertyUpdates(v []BTMetadataPropertyUpdateParams)`

SetPropertyUpdates sets PropertyUpdates field to given value.

### HasPropertyUpdates

`func (o *BTAppElementUpdateParams) HasPropertyUpdates() bool`

HasPropertyUpdates returns a boolean if a field has been set.

### GetReturnError

`func (o *BTAppElementUpdateParams) GetReturnError() bool`

GetReturnError returns the ReturnError field if non-nil, zero value otherwise.

### GetReturnErrorOk

`func (o *BTAppElementUpdateParams) GetReturnErrorOk() (*bool, bool)`

GetReturnErrorOk returns a tuple with the ReturnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnError

`func (o *BTAppElementUpdateParams) SetReturnError(v bool)`

SetReturnError sets ReturnError field to given value.

### HasReturnError

`func (o *BTAppElementUpdateParams) HasReturnError() bool`

HasReturnError returns a boolean if a field has been set.

### GetReturnJsonDifferenceFormat

`func (o *BTAppElementUpdateParams) GetReturnJsonDifferenceFormat() string`

GetReturnJsonDifferenceFormat returns the ReturnJsonDifferenceFormat field if non-nil, zero value otherwise.

### GetReturnJsonDifferenceFormatOk

`func (o *BTAppElementUpdateParams) GetReturnJsonDifferenceFormatOk() (*string, bool)`

GetReturnJsonDifferenceFormatOk returns a tuple with the ReturnJsonDifferenceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnJsonDifferenceFormat

`func (o *BTAppElementUpdateParams) SetReturnJsonDifferenceFormat(v string)`

SetReturnJsonDifferenceFormat sets ReturnJsonDifferenceFormat field to given value.

### HasReturnJsonDifferenceFormat

`func (o *BTAppElementUpdateParams) HasReturnJsonDifferenceFormat() bool`

HasReturnJsonDifferenceFormat returns a boolean if a field has been set.

### GetTransactionId

`func (o *BTAppElementUpdateParams) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BTAppElementUpdateParams) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BTAppElementUpdateParams) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BTAppElementUpdateParams) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


