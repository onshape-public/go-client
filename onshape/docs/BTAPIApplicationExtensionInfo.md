# BTAPIApplicationExtensionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionBody** | Pointer to **string** |  | [optional] 
**ActionType** | Pointer to **int32** |  | [optional] 
**ActionUrl** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExtensionContext** | Pointer to **int64** |  | [optional] 
**ExtensionLocation** | Pointer to **int64** |  | [optional] 
**HasIcon** | Pointer to **bool** |  | [optional] 
**HasPendingIcon** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** | URI to fetch complete information of the resource. | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Id of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**ParentAppPrimaryFormat** | Pointer to **string** |  | [optional] 
**ShowResponse** | Pointer to **bool** |  | [optional] 
**ViewRef** | Pointer to **string** | URI to visualize the resource in a webclient if applicable. | [optional] 

## Methods

### NewBTAPIApplicationExtensionInfo

`func NewBTAPIApplicationExtensionInfo() *BTAPIApplicationExtensionInfo`

NewBTAPIApplicationExtensionInfo instantiates a new BTAPIApplicationExtensionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTAPIApplicationExtensionInfoWithDefaults

`func NewBTAPIApplicationExtensionInfoWithDefaults() *BTAPIApplicationExtensionInfo`

NewBTAPIApplicationExtensionInfoWithDefaults instantiates a new BTAPIApplicationExtensionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionBody

`func (o *BTAPIApplicationExtensionInfo) GetActionBody() string`

GetActionBody returns the ActionBody field if non-nil, zero value otherwise.

### GetActionBodyOk

`func (o *BTAPIApplicationExtensionInfo) GetActionBodyOk() (*string, bool)`

GetActionBodyOk returns a tuple with the ActionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionBody

`func (o *BTAPIApplicationExtensionInfo) SetActionBody(v string)`

SetActionBody sets ActionBody field to given value.

### HasActionBody

`func (o *BTAPIApplicationExtensionInfo) HasActionBody() bool`

HasActionBody returns a boolean if a field has been set.

### GetActionType

`func (o *BTAPIApplicationExtensionInfo) GetActionType() int32`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *BTAPIApplicationExtensionInfo) GetActionTypeOk() (*int32, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *BTAPIApplicationExtensionInfo) SetActionType(v int32)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *BTAPIApplicationExtensionInfo) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionUrl

`func (o *BTAPIApplicationExtensionInfo) GetActionUrl() string`

GetActionUrl returns the ActionUrl field if non-nil, zero value otherwise.

### GetActionUrlOk

`func (o *BTAPIApplicationExtensionInfo) GetActionUrlOk() (*string, bool)`

GetActionUrlOk returns a tuple with the ActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionUrl

`func (o *BTAPIApplicationExtensionInfo) SetActionUrl(v string)`

SetActionUrl sets ActionUrl field to given value.

### HasActionUrl

`func (o *BTAPIApplicationExtensionInfo) HasActionUrl() bool`

HasActionUrl returns a boolean if a field has been set.

### GetApplicationId

`func (o *BTAPIApplicationExtensionInfo) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *BTAPIApplicationExtensionInfo) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *BTAPIApplicationExtensionInfo) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *BTAPIApplicationExtensionInfo) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetClientId

`func (o *BTAPIApplicationExtensionInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BTAPIApplicationExtensionInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BTAPIApplicationExtensionInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BTAPIApplicationExtensionInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDescription

`func (o *BTAPIApplicationExtensionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTAPIApplicationExtensionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTAPIApplicationExtensionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTAPIApplicationExtensionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtensionContext

`func (o *BTAPIApplicationExtensionInfo) GetExtensionContext() int64`

GetExtensionContext returns the ExtensionContext field if non-nil, zero value otherwise.

### GetExtensionContextOk

`func (o *BTAPIApplicationExtensionInfo) GetExtensionContextOk() (*int64, bool)`

GetExtensionContextOk returns a tuple with the ExtensionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionContext

`func (o *BTAPIApplicationExtensionInfo) SetExtensionContext(v int64)`

SetExtensionContext sets ExtensionContext field to given value.

### HasExtensionContext

`func (o *BTAPIApplicationExtensionInfo) HasExtensionContext() bool`

HasExtensionContext returns a boolean if a field has been set.

### GetExtensionLocation

`func (o *BTAPIApplicationExtensionInfo) GetExtensionLocation() int64`

GetExtensionLocation returns the ExtensionLocation field if non-nil, zero value otherwise.

### GetExtensionLocationOk

`func (o *BTAPIApplicationExtensionInfo) GetExtensionLocationOk() (*int64, bool)`

GetExtensionLocationOk returns a tuple with the ExtensionLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionLocation

`func (o *BTAPIApplicationExtensionInfo) SetExtensionLocation(v int64)`

SetExtensionLocation sets ExtensionLocation field to given value.

### HasExtensionLocation

`func (o *BTAPIApplicationExtensionInfo) HasExtensionLocation() bool`

HasExtensionLocation returns a boolean if a field has been set.

### GetHasIcon

`func (o *BTAPIApplicationExtensionInfo) GetHasIcon() bool`

GetHasIcon returns the HasIcon field if non-nil, zero value otherwise.

### GetHasIconOk

`func (o *BTAPIApplicationExtensionInfo) GetHasIconOk() (*bool, bool)`

GetHasIconOk returns a tuple with the HasIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIcon

`func (o *BTAPIApplicationExtensionInfo) SetHasIcon(v bool)`

SetHasIcon sets HasIcon field to given value.

### HasHasIcon

`func (o *BTAPIApplicationExtensionInfo) HasHasIcon() bool`

HasHasIcon returns a boolean if a field has been set.

### GetHasPendingIcon

`func (o *BTAPIApplicationExtensionInfo) GetHasPendingIcon() bool`

GetHasPendingIcon returns the HasPendingIcon field if non-nil, zero value otherwise.

### GetHasPendingIconOk

`func (o *BTAPIApplicationExtensionInfo) GetHasPendingIconOk() (*bool, bool)`

GetHasPendingIconOk returns a tuple with the HasPendingIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingIcon

`func (o *BTAPIApplicationExtensionInfo) SetHasPendingIcon(v bool)`

SetHasPendingIcon sets HasPendingIcon field to given value.

### HasHasPendingIcon

`func (o *BTAPIApplicationExtensionInfo) HasHasPendingIcon() bool`

HasHasPendingIcon returns a boolean if a field has been set.

### GetHref

`func (o *BTAPIApplicationExtensionInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTAPIApplicationExtensionInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTAPIApplicationExtensionInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTAPIApplicationExtensionInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetIconUrl

`func (o *BTAPIApplicationExtensionInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *BTAPIApplicationExtensionInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *BTAPIApplicationExtensionInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *BTAPIApplicationExtensionInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *BTAPIApplicationExtensionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTAPIApplicationExtensionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTAPIApplicationExtensionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTAPIApplicationExtensionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BTAPIApplicationExtensionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTAPIApplicationExtensionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTAPIApplicationExtensionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTAPIApplicationExtensionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentAppPrimaryFormat

`func (o *BTAPIApplicationExtensionInfo) GetParentAppPrimaryFormat() string`

GetParentAppPrimaryFormat returns the ParentAppPrimaryFormat field if non-nil, zero value otherwise.

### GetParentAppPrimaryFormatOk

`func (o *BTAPIApplicationExtensionInfo) GetParentAppPrimaryFormatOk() (*string, bool)`

GetParentAppPrimaryFormatOk returns a tuple with the ParentAppPrimaryFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAppPrimaryFormat

`func (o *BTAPIApplicationExtensionInfo) SetParentAppPrimaryFormat(v string)`

SetParentAppPrimaryFormat sets ParentAppPrimaryFormat field to given value.

### HasParentAppPrimaryFormat

`func (o *BTAPIApplicationExtensionInfo) HasParentAppPrimaryFormat() bool`

HasParentAppPrimaryFormat returns a boolean if a field has been set.

### GetShowResponse

`func (o *BTAPIApplicationExtensionInfo) GetShowResponse() bool`

GetShowResponse returns the ShowResponse field if non-nil, zero value otherwise.

### GetShowResponseOk

`func (o *BTAPIApplicationExtensionInfo) GetShowResponseOk() (*bool, bool)`

GetShowResponseOk returns a tuple with the ShowResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResponse

`func (o *BTAPIApplicationExtensionInfo) SetShowResponse(v bool)`

SetShowResponse sets ShowResponse field to given value.

### HasShowResponse

`func (o *BTAPIApplicationExtensionInfo) HasShowResponse() bool`

HasShowResponse returns a boolean if a field has been set.

### GetViewRef

`func (o *BTAPIApplicationExtensionInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTAPIApplicationExtensionInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTAPIApplicationExtensionInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTAPIApplicationExtensionInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


