# BTWorkflowAuditLogEntryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalOverride** | Pointer to **bool** |  | [optional] 
**ApproverIds** | Pointer to **[]string** |  | [optional] 
**CommentId** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **JSONTime** |  | [optional] 
**EntryType** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FeatureScriptConsole** | Pointer to **string** |  | [optional] 
**FeatureScriptNotices** | Pointer to **[]string** |  | [optional] 
**FeatureScriptResponse** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**PropertyUpdates** | Pointer to [**[]BTPropertyUpdateInfo**](BTPropertyUpdateInfo.md) |  | [optional] 
**SupportCode** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**WorkflowAction** | Pointer to **string** |  | [optional] 
**WorkflowState** | Pointer to **string** |  | [optional] 
**WorkflowTransition** | Pointer to **string** |  | [optional] 

## Methods

### NewBTWorkflowAuditLogEntryInfo

`func NewBTWorkflowAuditLogEntryInfo() *BTWorkflowAuditLogEntryInfo`

NewBTWorkflowAuditLogEntryInfo instantiates a new BTWorkflowAuditLogEntryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowAuditLogEntryInfoWithDefaults

`func NewBTWorkflowAuditLogEntryInfoWithDefaults() *BTWorkflowAuditLogEntryInfo`

NewBTWorkflowAuditLogEntryInfoWithDefaults instantiates a new BTWorkflowAuditLogEntryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalOverride

`func (o *BTWorkflowAuditLogEntryInfo) GetApprovalOverride() bool`

GetApprovalOverride returns the ApprovalOverride field if non-nil, zero value otherwise.

### GetApprovalOverrideOk

`func (o *BTWorkflowAuditLogEntryInfo) GetApprovalOverrideOk() (*bool, bool)`

GetApprovalOverrideOk returns a tuple with the ApprovalOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalOverride

`func (o *BTWorkflowAuditLogEntryInfo) SetApprovalOverride(v bool)`

SetApprovalOverride sets ApprovalOverride field to given value.

### HasApprovalOverride

`func (o *BTWorkflowAuditLogEntryInfo) HasApprovalOverride() bool`

HasApprovalOverride returns a boolean if a field has been set.

### GetApproverIds

`func (o *BTWorkflowAuditLogEntryInfo) GetApproverIds() []string`

GetApproverIds returns the ApproverIds field if non-nil, zero value otherwise.

### GetApproverIdsOk

`func (o *BTWorkflowAuditLogEntryInfo) GetApproverIdsOk() (*[]string, bool)`

GetApproverIdsOk returns a tuple with the ApproverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverIds

`func (o *BTWorkflowAuditLogEntryInfo) SetApproverIds(v []string)`

SetApproverIds sets ApproverIds field to given value.

### HasApproverIds

`func (o *BTWorkflowAuditLogEntryInfo) HasApproverIds() bool`

HasApproverIds returns a boolean if a field has been set.

### GetCommentId

`func (o *BTWorkflowAuditLogEntryInfo) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *BTWorkflowAuditLogEntryInfo) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *BTWorkflowAuditLogEntryInfo) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *BTWorkflowAuditLogEntryInfo) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetDate

`func (o *BTWorkflowAuditLogEntryInfo) GetDate() JSONTime`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BTWorkflowAuditLogEntryInfo) GetDateOk() (*JSONTime, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BTWorkflowAuditLogEntryInfo) SetDate(v JSONTime)`

SetDate sets Date field to given value.

### HasDate

`func (o *BTWorkflowAuditLogEntryInfo) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEntryType

`func (o *BTWorkflowAuditLogEntryInfo) GetEntryType() int32`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *BTWorkflowAuditLogEntryInfo) GetEntryTypeOk() (*int32, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *BTWorkflowAuditLogEntryInfo) SetEntryType(v int32)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *BTWorkflowAuditLogEntryInfo) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetErrorMessage

`func (o *BTWorkflowAuditLogEntryInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BTWorkflowAuditLogEntryInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BTWorkflowAuditLogEntryInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BTWorkflowAuditLogEntryInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFeatureScriptConsole

`func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptConsole() string`

GetFeatureScriptConsole returns the FeatureScriptConsole field if non-nil, zero value otherwise.

### GetFeatureScriptConsoleOk

`func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptConsoleOk() (*string, bool)`

GetFeatureScriptConsoleOk returns a tuple with the FeatureScriptConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureScriptConsole

`func (o *BTWorkflowAuditLogEntryInfo) SetFeatureScriptConsole(v string)`

SetFeatureScriptConsole sets FeatureScriptConsole field to given value.

### HasFeatureScriptConsole

`func (o *BTWorkflowAuditLogEntryInfo) HasFeatureScriptConsole() bool`

HasFeatureScriptConsole returns a boolean if a field has been set.

### GetFeatureScriptNotices

`func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptNotices() []string`

GetFeatureScriptNotices returns the FeatureScriptNotices field if non-nil, zero value otherwise.

### GetFeatureScriptNoticesOk

`func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptNoticesOk() (*[]string, bool)`

GetFeatureScriptNoticesOk returns a tuple with the FeatureScriptNotices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureScriptNotices

`func (o *BTWorkflowAuditLogEntryInfo) SetFeatureScriptNotices(v []string)`

SetFeatureScriptNotices sets FeatureScriptNotices field to given value.

### HasFeatureScriptNotices

`func (o *BTWorkflowAuditLogEntryInfo) HasFeatureScriptNotices() bool`

HasFeatureScriptNotices returns a boolean if a field has been set.

### GetFeatureScriptResponse

`func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptResponse() map[string]interface{}`

GetFeatureScriptResponse returns the FeatureScriptResponse field if non-nil, zero value otherwise.

### GetFeatureScriptResponseOk

`func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptResponseOk() (*map[string]interface{}, bool)`

GetFeatureScriptResponseOk returns a tuple with the FeatureScriptResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureScriptResponse

`func (o *BTWorkflowAuditLogEntryInfo) SetFeatureScriptResponse(v map[string]interface{})`

SetFeatureScriptResponse sets FeatureScriptResponse field to given value.

### HasFeatureScriptResponse

`func (o *BTWorkflowAuditLogEntryInfo) HasFeatureScriptResponse() bool`

HasFeatureScriptResponse returns a boolean if a field has been set.

### GetId

`func (o *BTWorkflowAuditLogEntryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTWorkflowAuditLogEntryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTWorkflowAuditLogEntryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTWorkflowAuditLogEntryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectId

`func (o *BTWorkflowAuditLogEntryInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTWorkflowAuditLogEntryInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTWorkflowAuditLogEntryInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTWorkflowAuditLogEntryInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetPropertyUpdates

`func (o *BTWorkflowAuditLogEntryInfo) GetPropertyUpdates() []BTPropertyUpdateInfo`

GetPropertyUpdates returns the PropertyUpdates field if non-nil, zero value otherwise.

### GetPropertyUpdatesOk

`func (o *BTWorkflowAuditLogEntryInfo) GetPropertyUpdatesOk() (*[]BTPropertyUpdateInfo, bool)`

GetPropertyUpdatesOk returns a tuple with the PropertyUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUpdates

`func (o *BTWorkflowAuditLogEntryInfo) SetPropertyUpdates(v []BTPropertyUpdateInfo)`

SetPropertyUpdates sets PropertyUpdates field to given value.

### HasPropertyUpdates

`func (o *BTWorkflowAuditLogEntryInfo) HasPropertyUpdates() bool`

HasPropertyUpdates returns a boolean if a field has been set.

### GetSupportCode

`func (o *BTWorkflowAuditLogEntryInfo) GetSupportCode() string`

GetSupportCode returns the SupportCode field if non-nil, zero value otherwise.

### GetSupportCodeOk

`func (o *BTWorkflowAuditLogEntryInfo) GetSupportCodeOk() (*string, bool)`

GetSupportCodeOk returns a tuple with the SupportCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCode

`func (o *BTWorkflowAuditLogEntryInfo) SetSupportCode(v string)`

SetSupportCode sets SupportCode field to given value.

### HasSupportCode

`func (o *BTWorkflowAuditLogEntryInfo) HasSupportCode() bool`

HasSupportCode returns a boolean if a field has been set.

### GetUserId

`func (o *BTWorkflowAuditLogEntryInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BTWorkflowAuditLogEntryInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BTWorkflowAuditLogEntryInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BTWorkflowAuditLogEntryInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWorkflowAction

`func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowAction() string`

GetWorkflowAction returns the WorkflowAction field if non-nil, zero value otherwise.

### GetWorkflowActionOk

`func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowActionOk() (*string, bool)`

GetWorkflowActionOk returns a tuple with the WorkflowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowAction

`func (o *BTWorkflowAuditLogEntryInfo) SetWorkflowAction(v string)`

SetWorkflowAction sets WorkflowAction field to given value.

### HasWorkflowAction

`func (o *BTWorkflowAuditLogEntryInfo) HasWorkflowAction() bool`

HasWorkflowAction returns a boolean if a field has been set.

### GetWorkflowState

`func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowState() string`

GetWorkflowState returns the WorkflowState field if non-nil, zero value otherwise.

### GetWorkflowStateOk

`func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowStateOk() (*string, bool)`

GetWorkflowStateOk returns a tuple with the WorkflowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowState

`func (o *BTWorkflowAuditLogEntryInfo) SetWorkflowState(v string)`

SetWorkflowState sets WorkflowState field to given value.

### HasWorkflowState

`func (o *BTWorkflowAuditLogEntryInfo) HasWorkflowState() bool`

HasWorkflowState returns a boolean if a field has been set.

### GetWorkflowTransition

`func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowTransition() string`

GetWorkflowTransition returns the WorkflowTransition field if non-nil, zero value otherwise.

### GetWorkflowTransitionOk

`func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowTransitionOk() (*string, bool)`

GetWorkflowTransitionOk returns a tuple with the WorkflowTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowTransition

`func (o *BTWorkflowAuditLogEntryInfo) SetWorkflowTransition(v string)`

SetWorkflowTransition sets WorkflowTransition field to given value.

### HasWorkflowTransition

`func (o *BTWorkflowAuditLogEntryInfo) HasWorkflowTransition() bool`

HasWorkflowTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


