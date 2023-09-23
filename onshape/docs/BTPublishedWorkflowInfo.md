# BTPublishedWorkflowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveState** | Pointer to **int32** |  | [optional] 
**CompanyId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DocumentId** | Pointer to **string** |  | [optional] 
**ElementId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageSrc** | Pointer to **string** |  | [optional] 
**IsPickable** | Pointer to **bool** |  | [optional] 
**Json** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **int32** |  | [optional] 
**OwnerType** | Pointer to **int32** |  | [optional] 
**PublishedDate** | Pointer to **JSONTime** | The date of publication of workflow | [optional] 
**UsesExternalPlm** | Pointer to **bool** | Whether the workflow connects to an external PLM service like Arena | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewBTPublishedWorkflowInfo

`func NewBTPublishedWorkflowInfo() *BTPublishedWorkflowInfo`

NewBTPublishedWorkflowInfo instantiates a new BTPublishedWorkflowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPublishedWorkflowInfoWithDefaults

`func NewBTPublishedWorkflowInfoWithDefaults() *BTPublishedWorkflowInfo`

NewBTPublishedWorkflowInfoWithDefaults instantiates a new BTPublishedWorkflowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveState

`func (o *BTPublishedWorkflowInfo) GetActiveState() int32`

GetActiveState returns the ActiveState field if non-nil, zero value otherwise.

### GetActiveStateOk

`func (o *BTPublishedWorkflowInfo) GetActiveStateOk() (*int32, bool)`

GetActiveStateOk returns a tuple with the ActiveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveState

`func (o *BTPublishedWorkflowInfo) SetActiveState(v int32)`

SetActiveState sets ActiveState field to given value.

### HasActiveState

`func (o *BTPublishedWorkflowInfo) HasActiveState() bool`

HasActiveState returns a boolean if a field has been set.

### GetCompanyId

`func (o *BTPublishedWorkflowInfo) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *BTPublishedWorkflowInfo) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *BTPublishedWorkflowInfo) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *BTPublishedWorkflowInfo) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetDescription

`func (o *BTPublishedWorkflowInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTPublishedWorkflowInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTPublishedWorkflowInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTPublishedWorkflowInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentId

`func (o *BTPublishedWorkflowInfo) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *BTPublishedWorkflowInfo) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *BTPublishedWorkflowInfo) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *BTPublishedWorkflowInfo) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetElementId

`func (o *BTPublishedWorkflowInfo) GetElementId() string`

GetElementId returns the ElementId field if non-nil, zero value otherwise.

### GetElementIdOk

`func (o *BTPublishedWorkflowInfo) GetElementIdOk() (*string, bool)`

GetElementIdOk returns a tuple with the ElementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementId

`func (o *BTPublishedWorkflowInfo) SetElementId(v string)`

SetElementId sets ElementId field to given value.

### HasElementId

`func (o *BTPublishedWorkflowInfo) HasElementId() bool`

HasElementId returns a boolean if a field has been set.

### GetId

`func (o *BTPublishedWorkflowInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTPublishedWorkflowInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTPublishedWorkflowInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTPublishedWorkflowInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageSrc

`func (o *BTPublishedWorkflowInfo) GetImageSrc() string`

GetImageSrc returns the ImageSrc field if non-nil, zero value otherwise.

### GetImageSrcOk

`func (o *BTPublishedWorkflowInfo) GetImageSrcOk() (*string, bool)`

GetImageSrcOk returns a tuple with the ImageSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSrc

`func (o *BTPublishedWorkflowInfo) SetImageSrc(v string)`

SetImageSrc sets ImageSrc field to given value.

### HasImageSrc

`func (o *BTPublishedWorkflowInfo) HasImageSrc() bool`

HasImageSrc returns a boolean if a field has been set.

### GetIsPickable

`func (o *BTPublishedWorkflowInfo) GetIsPickable() bool`

GetIsPickable returns the IsPickable field if non-nil, zero value otherwise.

### GetIsPickableOk

`func (o *BTPublishedWorkflowInfo) GetIsPickableOk() (*bool, bool)`

GetIsPickableOk returns a tuple with the IsPickable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPickable

`func (o *BTPublishedWorkflowInfo) SetIsPickable(v bool)`

SetIsPickable sets IsPickable field to given value.

### HasIsPickable

`func (o *BTPublishedWorkflowInfo) HasIsPickable() bool`

HasIsPickable returns a boolean if a field has been set.

### GetJson

`func (o *BTPublishedWorkflowInfo) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BTPublishedWorkflowInfo) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BTPublishedWorkflowInfo) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BTPublishedWorkflowInfo) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *BTPublishedWorkflowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTPublishedWorkflowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTPublishedWorkflowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTPublishedWorkflowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *BTPublishedWorkflowInfo) GetObjectType() int32`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *BTPublishedWorkflowInfo) GetObjectTypeOk() (*int32, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *BTPublishedWorkflowInfo) SetObjectType(v int32)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *BTPublishedWorkflowInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOwnerType

`func (o *BTPublishedWorkflowInfo) GetOwnerType() int32`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *BTPublishedWorkflowInfo) GetOwnerTypeOk() (*int32, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *BTPublishedWorkflowInfo) SetOwnerType(v int32)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *BTPublishedWorkflowInfo) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPublishedDate

`func (o *BTPublishedWorkflowInfo) GetPublishedDate() JSONTime`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *BTPublishedWorkflowInfo) GetPublishedDateOk() (*JSONTime, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *BTPublishedWorkflowInfo) SetPublishedDate(v JSONTime)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *BTPublishedWorkflowInfo) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetUsesExternalPlm

`func (o *BTPublishedWorkflowInfo) GetUsesExternalPlm() bool`

GetUsesExternalPlm returns the UsesExternalPlm field if non-nil, zero value otherwise.

### GetUsesExternalPlmOk

`func (o *BTPublishedWorkflowInfo) GetUsesExternalPlmOk() (*bool, bool)`

GetUsesExternalPlmOk returns a tuple with the UsesExternalPlm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesExternalPlm

`func (o *BTPublishedWorkflowInfo) SetUsesExternalPlm(v bool)`

SetUsesExternalPlm sets UsesExternalPlm field to given value.

### HasUsesExternalPlm

`func (o *BTPublishedWorkflowInfo) HasUsesExternalPlm() bool`

HasUsesExternalPlm returns a boolean if a field has been set.

### GetVersionId

`func (o *BTPublishedWorkflowInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *BTPublishedWorkflowInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *BTPublishedWorkflowInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *BTPublishedWorkflowInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


