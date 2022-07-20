# BTWorkflowAuditLogInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyId** | Pointer to **string** |  | [optional] 
**DebugMicroversionId** | Pointer to **string** |  | [optional] 
**Entries** | Pointer to [**[]BTWorkflowAuditLogEntryInfo**](BTWorkflowAuditLogEntryInfo.md) |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **int32** |  | [optional] 
**PublishedWorkflowId** | Pointer to [**BTPublishedWorkflowId**](BTPublishedWorkflowId.md) |  | [optional] 

## Methods

### NewBTWorkflowAuditLogInfo

`func NewBTWorkflowAuditLogInfo() *BTWorkflowAuditLogInfo`

NewBTWorkflowAuditLogInfo instantiates a new BTWorkflowAuditLogInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTWorkflowAuditLogInfoWithDefaults

`func NewBTWorkflowAuditLogInfoWithDefaults() *BTWorkflowAuditLogInfo`

NewBTWorkflowAuditLogInfoWithDefaults instantiates a new BTWorkflowAuditLogInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyId

`func (o *BTWorkflowAuditLogInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTWorkflowAuditLogInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTWorkflowAuditLogInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTWorkflowAuditLogInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDebugMicroversionId

`func (o *BTWorkflowAuditLogInfo) GetDebugMicroversionId() string`

GetDebugMicroversionId returns the DebugMicroversionId field if non-nil, zero value otherwise.

### GetDebugMicroversionIdOk

`func (o *BTWorkflowAuditLogInfo) GetDebugMicroversionIdOk() (*string, bool)`

GetDebugMicroversionIdOk returns a tuple with the DebugMicroversionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugMicroversionId

`func (o *BTWorkflowAuditLogInfo) SetDebugMicroversionId(v string)`

SetDebugMicroversionId sets DebugMicroversionId field to given value.

### HasDebugMicroversionId

`func (o *BTWorkflowAuditLogInfo) HasDebugMicroversionId() bool`

HasDebugMicroversionId returns a boolean if a field has been set.

### GetEntries

`func (o *BTWorkflowAuditLogInfo) GetEntries() []BTWorkflowAuditLogEntryInfo`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *BTWorkflowAuditLogInfo) GetEntriesOk() (*[]BTWorkflowAuditLogEntryInfo, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *BTWorkflowAuditLogInfo) SetEntries(v []BTWorkflowAuditLogEntryInfo)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *BTWorkflowAuditLogInfo) HasEntries() bool`

HasEntries returns a boolean if a field has been set.

### GetObjectId

`func (o *BTWorkflowAuditLogInfo) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BTWorkflowAuditLogInfo) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BTWorkflowAuditLogInfo) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *BTWorkflowAuditLogInfo) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *BTWorkflowAuditLogInfo) GetObjectType() int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTWorkflowAuditLogInfo) GetObjectTypeOk() (*int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTWorkflowAuditLogInfo) SetObjectType(v int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTWorkflowAuditLogInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPublishedWorkflowId

`func (o *BTWorkflowAuditLogInfo) GetPublishedWorkflowId() BTPublishedWorkflowId`

GetPublishedWorkflowId returns the PublishedWorkflowId field if non-nil, zero value otherwise.

### GetPublishedWorkflowIdOk

`func (o *BTWorkflowAuditLogInfo) GetPublishedWorkflowIdOk() (*BTPublishedWorkflowId, bool)`

GetPublishedWorkflowIdOk returns a tuple with the PublishedWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedWorkflowId

`func (o *BTWorkflowAuditLogInfo) SetPublishedWorkflowId(v BTPublishedWorkflowId)`

SetPublishedWorkflowId sets PublishedWorkflowId field to given value.

### HasPublishedWorkflowId

`func (o *BTWorkflowAuditLogInfo) HasPublishedWorkflowId() bool`

HasPublishedWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


