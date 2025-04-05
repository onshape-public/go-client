# BTDocumentSearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentFilter** | Pointer to **int32** | Type of documents to search: &#x60;0: My Documents | 1: Created | 2: Shared | 3: Trash | 4: Public | 5: Recent | 6: By Owner | 7: By Company | 9: By Team&#x60; | [optional] 
**FoundIn** | Pointer to [**BTESVersionWorkspaceChoice**](BTESVersionWorkspaceChoice.md) |  | [optional] 
**Limit** | Pointer to **int32** | Number of results to return per page. Default value is 20 (also the maximum). | [optional] [default to 20]
**Offset** | Pointer to **int32** | Offset. Determines where search results begin. Default value is 0. | [optional] [default to 0]
**OwnerId** | Pointer to **string** | Owner ID. Can be a user ID, company ID, or team ID, depending on &#x60;ownerType&#x60;. | [optional] 
**ParentId** | Pointer to **string** | Search document parent Id  | [optional] 
**RawQuery** | Pointer to **string** | Search for documents that contain the given string in the name. Search is not case-sensitive. | [optional] 
**SortColumn** | Pointer to **string** | Column by which to sort search results. &#x60;name | modifiedAt | createdAt (default) | email | modifiedBy | promotedAt&#x60; | [optional] [default to "createdAt"]
**SortOrder** | Pointer to **string** | Type of documents to search: &#x60;0: My Documents | 1: Created | 2: Shared | 3: Trash | 4: Public | 5: Recent | 6: By Owner | 7: By Company | 9: By Team&#x60; | [optional] [default to "desc"]
**Type** | Pointer to **string** | Type of owner. &#x60;0: User | 1: Company | 2: Onshape&#x60;. If the owner is a teamId, leave this unspecified. | [optional] 
**When** | Pointer to [**BTESResultsFilter**](BTESResultsFilter.md) |  | [optional] 

## Methods

### NewBTDocumentSearchParams

`func NewBTDocumentSearchParams() *BTDocumentSearchParams`

NewBTDocumentSearchParams instantiates a new BTDocumentSearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentSearchParamsWithDefaults

`func NewBTDocumentSearchParamsWithDefaults() *BTDocumentSearchParams`

NewBTDocumentSearchParamsWithDefaults instantiates a new BTDocumentSearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentFilter

`func (o *BTDocumentSearchParams) GetDocumentFilter() int32`

GetDocumentFilter returns the DocumentFilter field if non-nil, zero value otherwise.

### GetDocumentFilterOk

`func (o *BTDocumentSearchParams) GetDocumentFilterOk() (*int32, bool)`

GetDocumentFilterOk returns a tuple with the DocumentFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentFilter

`func (o *BTDocumentSearchParams) SetDocumentFilter(v int32)`

SetDocumentFilter sets DocumentFilter field to given value.

### HasDocumentFilter

`func (o *BTDocumentSearchParams) HasDocumentFilter() bool`

HasDocumentFilter returns a boolean if a field has been set.

### GetFoundIn

`func (o *BTDocumentSearchParams) GetFoundIn() BTESVersionWorkspaceChoice`

GetFoundIn returns the FoundIn field if non-nil, zero value otherwise.

### GetFoundInOk

`func (o *BTDocumentSearchParams) GetFoundInOk() (*BTESVersionWorkspaceChoice, bool)`

GetFoundInOk returns a tuple with the FoundIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundIn

`func (o *BTDocumentSearchParams) SetFoundIn(v BTESVersionWorkspaceChoice)`

SetFoundIn sets FoundIn field to given value.

### HasFoundIn

`func (o *BTDocumentSearchParams) HasFoundIn() bool`

HasFoundIn returns a boolean if a field has been set.

### GetLimit

`func (o *BTDocumentSearchParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BTDocumentSearchParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BTDocumentSearchParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BTDocumentSearchParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *BTDocumentSearchParams) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BTDocumentSearchParams) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BTDocumentSearchParams) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BTDocumentSearchParams) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOwnerId

`func (o *BTDocumentSearchParams) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *BTDocumentSearchParams) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *BTDocumentSearchParams) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *BTDocumentSearchParams) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetParentId

`func (o *BTDocumentSearchParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *BTDocumentSearchParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *BTDocumentSearchParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *BTDocumentSearchParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRawQuery

`func (o *BTDocumentSearchParams) GetRawQuery() string`

GetRawQuery returns the RawQuery field if non-nil, zero value otherwise.

### GetRawQueryOk

`func (o *BTDocumentSearchParams) GetRawQueryOk() (*string, bool)`

GetRawQueryOk returns a tuple with the RawQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawQuery

`func (o *BTDocumentSearchParams) SetRawQuery(v string)`

SetRawQuery sets RawQuery field to given value.

### HasRawQuery

`func (o *BTDocumentSearchParams) HasRawQuery() bool`

HasRawQuery returns a boolean if a field has been set.

### GetSortColumn

`func (o *BTDocumentSearchParams) GetSortColumn() string`

GetSortColumn returns the SortColumn field if non-nil, zero value otherwise.

### GetSortColumnOk

`func (o *BTDocumentSearchParams) GetSortColumnOk() (*string, bool)`

GetSortColumnOk returns a tuple with the SortColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortColumn

`func (o *BTDocumentSearchParams) SetSortColumn(v string)`

SetSortColumn sets SortColumn field to given value.

### HasSortColumn

`func (o *BTDocumentSearchParams) HasSortColumn() bool`

HasSortColumn returns a boolean if a field has been set.

### GetSortOrder

`func (o *BTDocumentSearchParams) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *BTDocumentSearchParams) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *BTDocumentSearchParams) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *BTDocumentSearchParams) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetType

`func (o *BTDocumentSearchParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTDocumentSearchParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTDocumentSearchParams) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BTDocumentSearchParams) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWhen

`func (o *BTDocumentSearchParams) GetWhen() BTESResultsFilter`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *BTDocumentSearchParams) GetWhenOk() (*BTESResultsFilter, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *BTDocumentSearchParams) SetWhen(v BTESResultsFilter)`

SetWhen sets When field to given value.

### HasWhen

`func (o *BTDocumentSearchParams) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


