# BTReleaseTransitionStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** |  | [optional] 
**LastStage** | Pointer to [**BTReleaseStage**](BTReleaseStage.md) |  | [optional] 
**LastUpdatedAt** | Pointer to **JSONTime** |  | [optional] 
**Retryable** | Pointer to **bool** |  | [optional] 
**SummaryState** | Pointer to [**BTReleaseSummaryState**](BTReleaseSummaryState.md) |  | [optional] 
**SupportCode** | Pointer to **string** |  | [optional] 

## Methods

### NewBTReleaseTransitionStatusInfo

`func NewBTReleaseTransitionStatusInfo() *BTReleaseTransitionStatusInfo`

NewBTReleaseTransitionStatusInfo instantiates a new BTReleaseTransitionStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTReleaseTransitionStatusInfoWithDefaults

`func NewBTReleaseTransitionStatusInfoWithDefaults() *BTReleaseTransitionStatusInfo`

NewBTReleaseTransitionStatusInfoWithDefaults instantiates a new BTReleaseTransitionStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *BTReleaseTransitionStatusInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BTReleaseTransitionStatusInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BTReleaseTransitionStatusInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BTReleaseTransitionStatusInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLastStage

`func (o *BTReleaseTransitionStatusInfo) GetLastStage() BTReleaseStage`

GetLastStage returns the LastStage field if non-nil, zero value otherwise.

### GetLastStageOk

`func (o *BTReleaseTransitionStatusInfo) GetLastStageOk() (*BTReleaseStage, bool)`

GetLastStageOk returns a tuple with the LastStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStage

`func (o *BTReleaseTransitionStatusInfo) SetLastStage(v BTReleaseStage)`

SetLastStage sets LastStage field to given value.

### HasLastStage

`func (o *BTReleaseTransitionStatusInfo) HasLastStage() bool`

HasLastStage returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *BTReleaseTransitionStatusInfo) GetLastUpdatedAt() JSONTime`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *BTReleaseTransitionStatusInfo) GetLastUpdatedAtOk() (*JSONTime, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *BTReleaseTransitionStatusInfo) SetLastUpdatedAt(v JSONTime)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *BTReleaseTransitionStatusInfo) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetRetryable

`func (o *BTReleaseTransitionStatusInfo) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *BTReleaseTransitionStatusInfo) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *BTReleaseTransitionStatusInfo) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *BTReleaseTransitionStatusInfo) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetSummaryState

`func (o *BTReleaseTransitionStatusInfo) GetSummaryState() BTReleaseSummaryState`

GetSummaryState returns the SummaryState field if non-nil, zero value otherwise.

### GetSummaryStateOk

`func (o *BTReleaseTransitionStatusInfo) GetSummaryStateOk() (*BTReleaseSummaryState, bool)`

GetSummaryStateOk returns a tuple with the SummaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryState

`func (o *BTReleaseTransitionStatusInfo) SetSummaryState(v BTReleaseSummaryState)`

SetSummaryState sets SummaryState field to given value.

### HasSummaryState

`func (o *BTReleaseTransitionStatusInfo) HasSummaryState() bool`

HasSummaryState returns a boolean if a field has been set.

### GetSupportCode

`func (o *BTReleaseTransitionStatusInfo) GetSupportCode() string`

GetSupportCode returns the SupportCode field if non-nil, zero value otherwise.

### GetSupportCodeOk

`func (o *BTReleaseTransitionStatusInfo) GetSupportCodeOk() (*string, bool)`

GetSupportCodeOk returns a tuple with the SupportCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCode

`func (o *BTReleaseTransitionStatusInfo) SetSupportCode(v string)`

SetSupportCode sets SupportCode field to given value.

### HasSupportCode

`func (o *BTReleaseTransitionStatusInfo) HasSupportCode() bool`

HasSupportCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


