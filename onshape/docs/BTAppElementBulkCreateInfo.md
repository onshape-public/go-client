# BTAppElementBulkCreateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentMicroversionId** | **string** | The latest document microversion, after the edit was committed. | 
**ElementIds** | Pointer to **[]string** | The ids of the created elements. | [optional] 
**ElementMicroversions** | Pointer to **[]string** | The microversion ids of the created elements, at creation time. | [optional] 
**ErrorCode** | Pointer to **int32** | &#x60;0: OK (healthy) | 1: INFO | 2: WARNING | 3: ERROR (dangling or view generation call failed) | 4: UNKNOWN&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** | A human-readable value for the error that occurred, if one occurred. | [optional] 
**ErrorValue** | Pointer to [**BTAppElementErrorCode**](BTAppElementErrorCode.md) |  | [optional] 

## Methods

### NewBTAppElementBulkCreateInfo

`func NewBTAppElementBulkCreateInfo(documentMicroversionId string, ) *BTAppElementBulkCreateInfo`

NewBTAppElementBulkCreateInfo instantiates a new BTAppElementBulkCreateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAppElementBulkCreateInfoWithDefaults

`func NewBTAppElementBulkCreateInfoWithDefaults() *BTAppElementBulkCreateInfo`

NewBTAppElementBulkCreateInfoWithDefaults instantiates a new BTAppElementBulkCreateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentMicroversionId

`func (o *BTAppElementBulkCreateInfo) GetDocumentMicroversionId() string`

GetDocumentMicroversionId returns the DocumentMicroversionId field if non-nil, zero value otherwise.

### GetDocumentMicroversionIdOk

`func (o *BTAppElementBulkCreateInfo) GetDocumentMicroversionIdOk() (*string, bool)`

GetDocumentMicroversionIdOk returns a tuple with the DocumentMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMicroversionId

`func (o *BTAppElementBulkCreateInfo) SetDocumentMicroversionId(v string)`

SetDocumentMicroversionId sets DocumentMicroversionId field to given value.


### GetElementIds

`func (o *BTAppElementBulkCreateInfo) GetElementIds() []string`

GetElementIds returns the ElementIds field if non-nil, zero value otherwise.

### GetElementIdsOk

`func (o *BTAppElementBulkCreateInfo) GetElementIdsOk() (*[]string, bool)`

GetElementIdsOk returns a tuple with the ElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementIds

`func (o *BTAppElementBulkCreateInfo) SetElementIds(v []string)`

SetElementIds sets ElementIds field to given value.

### HasElementIds

`func (o *BTAppElementBulkCreateInfo) HasElementIds() bool`

HasElementIds returns a boolean if a field has been set.

### GetElementMicroversions

`func (o *BTAppElementBulkCreateInfo) GetElementMicroversions() []string`

GetElementMicroversions returns the ElementMicroversions field if non-nil, zero value otherwise.

### GetElementMicroversionsOk

`func (o *BTAppElementBulkCreateInfo) GetElementMicroversionsOk() (*[]string, bool)`

GetElementMicroversionsOk returns a tuple with the ElementMicroversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementMicroversions

`func (o *BTAppElementBulkCreateInfo) SetElementMicroversions(v []string)`

SetElementMicroversions sets ElementMicroversions field to given value.

### HasElementMicroversions

`func (o *BTAppElementBulkCreateInfo) HasElementMicroversions() bool`

HasElementMicroversions returns a boolean if a field has been set.

### GetErrorCode

`func (o *BTAppElementBulkCreateInfo) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BTAppElementBulkCreateInfo) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BTAppElementBulkCreateInfo) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BTAppElementBulkCreateInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *BTAppElementBulkCreateInfo) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BTAppElementBulkCreateInfo) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BTAppElementBulkCreateInfo) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *BTAppElementBulkCreateInfo) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetErrorValue

`func (o *BTAppElementBulkCreateInfo) GetErrorValue() BTAppElementErrorCode`

GetErrorValue returns the ErrorValue field if non-nil, zero value otherwise.

### GetErrorValueOk

`func (o *BTAppElementBulkCreateInfo) GetErrorValueOk() (*BTAppElementErrorCode, bool)`

GetErrorValueOk returns a tuple with the ErrorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorValue

`func (o *BTAppElementBulkCreateInfo) SetErrorValue(v BTAppElementErrorCode)`

SetErrorValue sets ErrorValue field to given value.

### HasErrorValue

`func (o *BTAppElementBulkCreateInfo) HasErrorValue() bool`

HasErrorValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


