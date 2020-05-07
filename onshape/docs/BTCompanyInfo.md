# BTCompanyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**BTAddressInfo**](BTAddressInfo.md) |  | [optional] 
**Admin** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DomainPrefix** | Pointer to **string** |  | [optional] 
**EnterpriseBaseUrl** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NoPublicDocuments** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Purchase** | Pointer to [**BTPurchaseInfo**](BTPurchaseInfo.md) |  | [optional] 
**SecondaryDomainPrefixes** | Pointer to **[]string** |  | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**ViewRef** | Pointer to **string** |  | [optional] 

## Methods

### NewBTCompanyInfo

`func NewBTCompanyInfo() *BTCompanyInfo`

NewBTCompanyInfo instantiates a new BTCompanyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTCompanyInfoWithDefaults

`func NewBTCompanyInfoWithDefaults() *BTCompanyInfo`

NewBTCompanyInfoWithDefaults instantiates a new BTCompanyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BTCompanyInfo) GetAddress() BTAddressInfo`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BTCompanyInfo) GetAddressOk() (*BTAddressInfo, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BTCompanyInfo) SetAddress(v BTAddressInfo)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BTCompanyInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAdmin

`func (o *BTCompanyInfo) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *BTCompanyInfo) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *BTCompanyInfo) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *BTCompanyInfo) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetDescription

`func (o *BTCompanyInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BTCompanyInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BTCompanyInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BTCompanyInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainPrefix

`func (o *BTCompanyInfo) GetDomainPrefix() string`

GetDomainPrefix returns the DomainPrefix field if non-nil, zero value otherwise.

### GetDomainPrefixOk

`func (o *BTCompanyInfo) GetDomainPrefixOk() (*string, bool)`

GetDomainPrefixOk returns a tuple with the DomainPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPrefix

`func (o *BTCompanyInfo) SetDomainPrefix(v string)`

SetDomainPrefix sets DomainPrefix field to given value.

### HasDomainPrefix

`func (o *BTCompanyInfo) HasDomainPrefix() bool`

HasDomainPrefix returns a boolean if a field has been set.

### GetEnterpriseBaseUrl

`func (o *BTCompanyInfo) GetEnterpriseBaseUrl() string`

GetEnterpriseBaseUrl returns the EnterpriseBaseUrl field if non-nil, zero value otherwise.

### GetEnterpriseBaseUrlOk

`func (o *BTCompanyInfo) GetEnterpriseBaseUrlOk() (*string, bool)`

GetEnterpriseBaseUrlOk returns a tuple with the EnterpriseBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseBaseUrl

`func (o *BTCompanyInfo) SetEnterpriseBaseUrl(v string)`

SetEnterpriseBaseUrl sets EnterpriseBaseUrl field to given value.

### HasEnterpriseBaseUrl

`func (o *BTCompanyInfo) HasEnterpriseBaseUrl() bool`

HasEnterpriseBaseUrl returns a boolean if a field has been set.

### GetHref

`func (o *BTCompanyInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BTCompanyInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BTCompanyInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BTCompanyInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *BTCompanyInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTCompanyInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTCompanyInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTCompanyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *BTCompanyInfo) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BTCompanyInfo) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BTCompanyInfo) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BTCompanyInfo) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *BTCompanyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTCompanyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTCompanyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTCompanyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNoPublicDocuments

`func (o *BTCompanyInfo) GetNoPublicDocuments() bool`

GetNoPublicDocuments returns the NoPublicDocuments field if non-nil, zero value otherwise.

### GetNoPublicDocumentsOk

`func (o *BTCompanyInfo) GetNoPublicDocumentsOk() (*bool, bool)`

GetNoPublicDocumentsOk returns a tuple with the NoPublicDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoPublicDocuments

`func (o *BTCompanyInfo) SetNoPublicDocuments(v bool)`

SetNoPublicDocuments sets NoPublicDocuments field to given value.

### HasNoPublicDocuments

`func (o *BTCompanyInfo) HasNoPublicDocuments() bool`

HasNoPublicDocuments returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTCompanyInfo) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTCompanyInfo) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTCompanyInfo) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTCompanyInfo) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPurchase

`func (o *BTCompanyInfo) GetPurchase() BTPurchaseInfo`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *BTCompanyInfo) GetPurchaseOk() (*BTPurchaseInfo, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *BTCompanyInfo) SetPurchase(v BTPurchaseInfo)`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *BTCompanyInfo) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### GetSecondaryDomainPrefixes

`func (o *BTCompanyInfo) GetSecondaryDomainPrefixes() []string`

GetSecondaryDomainPrefixes returns the SecondaryDomainPrefixes field if non-nil, zero value otherwise.

### GetSecondaryDomainPrefixesOk

`func (o *BTCompanyInfo) GetSecondaryDomainPrefixesOk() (*[]string, bool)`

GetSecondaryDomainPrefixesOk returns a tuple with the SecondaryDomainPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDomainPrefixes

`func (o *BTCompanyInfo) SetSecondaryDomainPrefixes(v []string)`

SetSecondaryDomainPrefixes sets SecondaryDomainPrefixes field to given value.

### HasSecondaryDomainPrefixes

`func (o *BTCompanyInfo) HasSecondaryDomainPrefixes() bool`

HasSecondaryDomainPrefixes returns a boolean if a field has been set.

### GetState

`func (o *BTCompanyInfo) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BTCompanyInfo) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BTCompanyInfo) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *BTCompanyInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *BTCompanyInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTCompanyInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTCompanyInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *BTCompanyInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetViewRef

`func (o *BTCompanyInfo) GetViewRef() string`

GetViewRef returns the ViewRef field if non-nil, zero value otherwise.

### GetViewRefOk

`func (o *BTCompanyInfo) GetViewRefOk() (*string, bool)`

GetViewRefOk returns a tuple with the ViewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRef

`func (o *BTCompanyInfo) SetViewRef(v string)`

SetViewRef sets ViewRef field to given value.

### HasViewRef

`func (o *BTCompanyInfo) HasViewRef() bool`

HasViewRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


