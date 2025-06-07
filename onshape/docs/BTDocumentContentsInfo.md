# BTDocumentContentsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elements** | Pointer to [**[]BTDocumentElementInfo**](BTDocumentElementInfo.md) | The elements (tabs) in the document. This does not include folders. | [optional] 
**Folders** | Pointer to [**BTElementGroup1458**](BTElementGroup1458.md) |  | [optional] 

## Methods

### NewBTDocumentContentsInfo

`func NewBTDocumentContentsInfo() *BTDocumentContentsInfo`

NewBTDocumentContentsInfo instantiates a new BTDocumentContentsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTDocumentContentsInfoWithDefaults

`func NewBTDocumentContentsInfoWithDefaults() *BTDocumentContentsInfo`

NewBTDocumentContentsInfoWithDefaults instantiates a new BTDocumentContentsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElements

`func (o *BTDocumentContentsInfo) GetElements() []BTDocumentElementInfo`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *BTDocumentContentsInfo) GetElementsOk() (*[]BTDocumentElementInfo, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *BTDocumentContentsInfo) SetElements(v []BTDocumentElementInfo)`

SetElements sets Elements field to given value.

### HasElements

`func (o *BTDocumentContentsInfo) HasElements() bool`

HasElements returns a boolean if a field has been set.

### GetFolders

`func (o *BTDocumentContentsInfo) GetFolders() BTElementGroup1458`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *BTDocumentContentsInfo) GetFoldersOk() (*BTElementGroup1458, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *BTDocumentContentsInfo) SetFolders(v BTElementGroup1458)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *BTDocumentContentsInfo) HasFolders() bool`

HasFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


