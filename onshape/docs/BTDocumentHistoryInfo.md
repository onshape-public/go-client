# BTDocumentHistoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanBeRestored** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **JSONTime** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MicroversionId** | Pointer to **string** |  | [optional] 
**NextMicroversionId** | Pointer to **string** |  | [optional] 
**RestoreId** | Pointer to **string** | If this microversion is the result of a restore from another microversion, the restoreId will be the microversion Id of the original microversion that was restored. Otherwise this id will not be included within the response. | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewBTDocumentHistoryInfo

`func NewBTDocumentHistoryInfo() *BTDocumentHistoryInfo`

NewBTDocumentHistoryInfo instantiates a new BTDocumentHistoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentHistoryInfoWithDefaults

`func NewBTDocumentHistoryInfoWithDefaults() *BTDocumentHistoryInfo`

NewBTDocumentHistoryInfoWithDefaults instantiates a new BTDocumentHistoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanBeRestored

`func (o *BTDocumentHistoryInfo) GetCanBeRestored() bool`

GetCanBeRestored returns the CanBeRestored field if non-nil, zero value otherwise.

### GetCanBeRestoredOk

`func (o *BTDocumentHistoryInfo) GetCanBeRestoredOk() (*bool, bool)`

GetCanBeRestoredOk returns a tuple with the CanBeRestored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeRestored

`func (o *BTDocumentHistoryInfo) SetCanBeRestored(v bool)`

SetCanBeRestored sets CanBeRestored field to given value.

### HasCanBeRestored

`func (o *BTDocumentHistoryInfo) HasCanBeRestored() bool`

HasCanBeRestored returns a boolean if a field has been set.

### GetDate

`func (o *BTDocumentHistoryInfo) GetDate() JSONTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BTDocumentHistoryInfo) GetDateOk() (*JSONTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BTDocumentHistoryInfo) SetDate(v JSONTime)`

SetDate sets Date field to given value.

### HasDate

`func (o *BTDocumentHistoryInfo) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDescription

`func (o *BTDocumentHistoryInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTDocumentHistoryInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTDocumentHistoryInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTDocumentHistoryInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMicroversionId

`func (o *BTDocumentHistoryInfo) GetMicroversionId() string`

GetMicroversionId returns the MicroversionId field if non-nil, zero value otherwise.

### GetMicroversionIdOk

`func (o *BTDocumentHistoryInfo) GetMicroversionIdOk() (*string, bool)`

GetMicroversionIdOk returns a tuple with the MicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicroversionId

`func (o *BTDocumentHistoryInfo) SetMicroversionId(v string)`

SetMicroversionId sets MicroversionId field to given value.

### HasMicroversionId

`func (o *BTDocumentHistoryInfo) HasMicroversionId() bool`

HasMicroversionId returns a boolean if a field has been set.

### GetNextMicroversionId

`func (o *BTDocumentHistoryInfo) GetNextMicroversionId() string`

GetNextMicroversionId returns the NextMicroversionId field if non-nil, zero value otherwise.

### GetNextMicroversionIdOk

`func (o *BTDocumentHistoryInfo) GetNextMicroversionIdOk() (*string, bool)`

GetNextMicroversionIdOk returns a tuple with the NextMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMicroversionId

`func (o *BTDocumentHistoryInfo) SetNextMicroversionId(v string)`

SetNextMicroversionId sets NextMicroversionId field to given value.

### HasNextMicroversionId

`func (o *BTDocumentHistoryInfo) HasNextMicroversionId() bool`

HasNextMicroversionId returns a boolean if a field has been set.

### GetRestoreId

`func (o *BTDocumentHistoryInfo) GetRestoreId() string`

GetRestoreId returns the RestoreId field if non-nil, zero value otherwise.

### GetRestoreIdOk

`func (o *BTDocumentHistoryInfo) GetRestoreIdOk() (*string, bool)`

GetRestoreIdOk returns a tuple with the RestoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreId

`func (o *BTDocumentHistoryInfo) SetRestoreId(v string)`

SetRestoreId sets RestoreId field to given value.

### HasRestoreId

`func (o *BTDocumentHistoryInfo) HasRestoreId() bool`

HasRestoreId returns a boolean if a field has been set.

### GetUserId

`func (o *BTDocumentHistoryInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTDocumentHistoryInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTDocumentHistoryInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTDocumentHistoryInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsername

`func (o *BTDocumentHistoryInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BTDocumentHistoryInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BTDocumentHistoryInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BTDocumentHistoryInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


