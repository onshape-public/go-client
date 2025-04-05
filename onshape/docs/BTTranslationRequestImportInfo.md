# BTTranslationRequestImportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** |  | [optional] 
**FailureReason** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**RequestElementId** | Pointer to **string** |  | [optional] 
**RequestState** | Pointer to [**BTTranslationRequestState**](BTTranslationRequestState.md) |  | [optional] 
**ResultDocumentId** | Pointer to **string** |  | [optional] 
**ResultElementIds** | Pointer to **[]string** |  | [optional] 
**ResultExternalDataIds** | Pointer to **[]string** |  | [optional] 
**ResultWorkspaceId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTTranslationRequestImportInfo

`func NewBTTranslationRequestImportInfo() *BTTranslationRequestImportInfo`

NewBTTranslationRequestImportInfo instantiates a new BTTranslationRequestImportInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTTranslationRequestImportInfoWithDefaults

`func NewBTTranslationRequestImportInfoWithDefaults() *BTTranslationRequestImportInfo`

NewBTTranslationRequestImportInfoWithDefaults instantiates a new BTTranslationRequestImportInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *BTTranslationRequestImportInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTTranslationRequestImportInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTTranslationRequestImportInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTTranslationRequestImportInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetFailureReason

`func (o *BTTranslationRequestImportInfo) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *BTTranslationRequestImportInfo) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *BTTranslationRequestImportInfo) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *BTTranslationRequestImportInfo) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetHref

`func (o *BTTranslationRequestImportInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTTranslationRequestImportInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTTranslationRequestImportInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTTranslationRequestImportInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTTranslationRequestImportInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTTranslationRequestImportInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTTranslationRequestImportInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTTranslationRequestImportInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTTranslationRequestImportInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTTranslationRequestImportInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTTranslationRequestImportInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTTranslationRequestImportInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestElementId

`func (o *BTTranslationRequestImportInfo) GetRequestElementId() string`

GetRequestElementId returns the RequestElementId field if non-nil, zero value otherwise.

### GetRequestElementIdOk

`func (o *BTTranslationRequestImportInfo) GetRequestElementIdOk() (*string, bool)`

GetRequestElementIdOk returns a tuple with the RequestElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestElementId

`func (o *BTTranslationRequestImportInfo) SetRequestElementId(v string)`

SetRequestElementId sets RequestElementId field to given value.

### HasRequestElementId

`func (o *BTTranslationRequestImportInfo) HasRequestElementId() bool`

HasRequestElementId returns a boolean if a field has been set.

### GetRequestState

`func (o *BTTranslationRequestImportInfo) GetRequestState() BTTranslationRequestState`

GetRequestState returns the RequestState field if non-nil, zero value otherwise.

### GetRequestStateOk

`func (o *BTTranslationRequestImportInfo) GetRequestStateOk() (*BTTranslationRequestState, bool)`

GetRequestStateOk returns a tuple with the RequestState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestState

`func (o *BTTranslationRequestImportInfo) SetRequestState(v BTTranslationRequestState)`

SetRequestState sets RequestState field to given value.

### HasRequestState

`func (o *BTTranslationRequestImportInfo) HasRequestState() bool`

HasRequestState returns a boolean if a field has been set.

### GetResultDocumentId

`func (o *BTTranslationRequestImportInfo) GetResultDocumentId() string`

GetResultDocumentId returns the ResultDocumentId field if non-nil, zero value otherwise.

### GetResultDocumentIdOk

`func (o *BTTranslationRequestImportInfo) GetResultDocumentIdOk() (*string, bool)`

GetResultDocumentIdOk returns a tuple with the ResultDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultDocumentId

`func (o *BTTranslationRequestImportInfo) SetResultDocumentId(v string)`

SetResultDocumentId sets ResultDocumentId field to given value.

### HasResultDocumentId

`func (o *BTTranslationRequestImportInfo) HasResultDocumentId() bool`

HasResultDocumentId returns a boolean if a field has been set.

### GetResultElementIds

`func (o *BTTranslationRequestImportInfo) GetResultElementIds() []string`

GetResultElementIds returns the ResultElementIds field if non-nil, zero value otherwise.

### GetResultElementIdsOk

`func (o *BTTranslationRequestImportInfo) GetResultElementIdsOk() (*[]string, bool)`

GetResultElementIdsOk returns a tuple with the ResultElementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultElementIds

`func (o *BTTranslationRequestImportInfo) SetResultElementIds(v []string)`

SetResultElementIds sets ResultElementIds field to given value.

### HasResultElementIds

`func (o *BTTranslationRequestImportInfo) HasResultElementIds() bool`

HasResultElementIds returns a boolean if a field has been set.

### GetResultExternalDataIds

`func (o *BTTranslationRequestImportInfo) GetResultExternalDataIds() []string`

GetResultExternalDataIds returns the ResultExternalDataIds field if non-nil, zero value otherwise.

### GetResultExternalDataIdsOk

`func (o *BTTranslationRequestImportInfo) GetResultExternalDataIdsOk() (*[]string, bool)`

GetResultExternalDataIdsOk returns a tuple with the ResultExternalDataIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultExternalDataIds

`func (o *BTTranslationRequestImportInfo) SetResultExternalDataIds(v []string)`

SetResultExternalDataIds sets ResultExternalDataIds field to given value.

### HasResultExternalDataIds

`func (o *BTTranslationRequestImportInfo) HasResultExternalDataIds() bool`

HasResultExternalDataIds returns a boolean if a field has been set.

### GetResultWorkspaceId

`func (o *BTTranslationRequestImportInfo) GetResultWorkspaceId() string`

GetResultWorkspaceId returns the ResultWorkspaceId field if non-nil, zero value otherwise.

### GetResultWorkspaceIdOk

`func (o *BTTranslationRequestImportInfo) GetResultWorkspaceIdOk() (*string, bool)`

GetResultWorkspaceIdOk returns a tuple with the ResultWorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultWorkspaceId

`func (o *BTTranslationRequestImportInfo) SetResultWorkspaceId(v string)`

SetResultWorkspaceId sets ResultWorkspaceId field to given value.

### HasResultWorkspaceId

`func (o *BTTranslationRequestImportInfo) HasResultWorkspaceId() bool`

HasResultWorkspaceId returns a boolean if a field has been set.

### GetVersionId

`func (o *BTTranslationRequestImportInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTTranslationRequestImportInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTTranslationRequestImportInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTTranslationRequestImportInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetViewRef

`func (o *BTTranslationRequestImportInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTTranslationRequestImportInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTTranslationRequestImportInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTTranslationRequestImportInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *BTTranslationRequestImportInfo) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BTTranslationRequestImportInfo) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BTTranslationRequestImportInfo) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *BTTranslationRequestImportInfo) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


