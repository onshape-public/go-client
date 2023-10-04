# BTExportModelBody1272

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**Closed** | Pointer to **bool** | If type &#x3D;&#x3D; COMPOSITE, indicates whether it is open or closed. | [optional] 
**ConstituentBodyIds** | Pointer to **[]string** |  | [optional] 
**ConsumedByComposite** | Pointer to **bool** | Indicates if there is a closed composite that consumes this body. | [optional] 
**Edges** | Pointer to [**[]BTExportModelEdge1782**](BTExportModelEdge1782.md) |  | [optional] 
**Faces** | Pointer to [**[]BTExportModelFace1363**](BTExportModelFace1363.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**BTExportBodyProperties3559**](BTExportBodyProperties3559.md) |  | [optional] 
**Type** | Pointer to [**GBTBodyType**](GBTBodyType.md) |  | [optional] 
**Vertices** | Pointer to [**[]BTExportModelVertex858**](BTExportModelVertex858.md) |  | [optional] 

## Methods

### NewBTExportModelBody1272

`func NewBTExportModelBody1272() *BTExportModelBody1272`

NewBTExportModelBody1272 instantiates a new BTExportModelBody1272 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTExportModelBody1272WithDefaults

`func NewBTExportModelBody1272WithDefaults() *BTExportModelBody1272`

NewBTExportModelBody1272WithDefaults instantiates a new BTExportModelBody1272 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTExportModelBody1272) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTExportModelBody1272) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTExportModelBody1272) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTExportModelBody1272) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetClosed

`func (o *BTExportModelBody1272) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *BTExportModelBody1272) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *BTExportModelBody1272) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *BTExportModelBody1272) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetConstituentBodyIds

`func (o *BTExportModelBody1272) GetConstituentBodyIds() []string`

GetConstituentBodyIds returns the ConstituentBodyIds field if non-nil, zero value otherwise.

### GetConstituentBodyIdsOk

`func (o *BTExportModelBody1272) GetConstituentBodyIdsOk() (*[]string, bool)`

GetConstituentBodyIdsOk returns a tuple with the ConstituentBodyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstituentBodyIds

`func (o *BTExportModelBody1272) SetConstituentBodyIds(v []string)`

SetConstituentBodyIds sets ConstituentBodyIds field to given value.

### HasConstituentBodyIds

`func (o *BTExportModelBody1272) HasConstituentBodyIds() bool`

HasConstituentBodyIds returns a boolean if a field has been set.

### GetConsumedByComposite

`func (o *BTExportModelBody1272) GetConsumedByComposite() bool`

GetConsumedByComposite returns the ConsumedByComposite field if non-nil, zero value otherwise.

### GetConsumedByCompositeOk

`func (o *BTExportModelBody1272) GetConsumedByCompositeOk() (*bool, bool)`

GetConsumedByCompositeOk returns a tuple with the ConsumedByComposite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedByComposite

`func (o *BTExportModelBody1272) SetConsumedByComposite(v bool)`

SetConsumedByComposite sets ConsumedByComposite field to given value.

### HasConsumedByComposite

`func (o *BTExportModelBody1272) HasConsumedByComposite() bool`

HasConsumedByComposite returns a boolean if a field has been set.

### GetEdges

`func (o *BTExportModelBody1272) GetEdges() []BTExportModelEdge1782`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *BTExportModelBody1272) GetEdgesOk() (*[]BTExportModelEdge1782, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *BTExportModelBody1272) SetEdges(v []BTExportModelEdge1782)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *BTExportModelBody1272) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetFaces

`func (o *BTExportModelBody1272) GetFaces() []BTExportModelFace1363`

GetFaces returns the Faces field if non-nil, zero value otherwise.

### GetFacesOk

`func (o *BTExportModelBody1272) GetFacesOk() (*[]BTExportModelFace1363, bool)`

GetFacesOk returns a tuple with the Faces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaces

`func (o *BTExportModelBody1272) SetFaces(v []BTExportModelFace1363)`

SetFaces sets Faces field to given value.

### HasFaces

`func (o *BTExportModelBody1272) HasFaces() bool`

HasFaces returns a boolean if a field has been set.

### GetId

`func (o *BTExportModelBody1272) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BTExportModelBody1272) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BTExportModelBody1272) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BTExportModelBody1272) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *BTExportModelBody1272) GetProperties() BTExportBodyProperties3559`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BTExportModelBody1272) GetPropertiesOk() (*BTExportBodyProperties3559, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BTExportModelBody1272) SetProperties(v BTExportBodyProperties3559)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *BTExportModelBody1272) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetType

`func (o *BTExportModelBody1272) GetType() GBTBodyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BTExportModelBody1272) GetTypeOk() (*GBTBodyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BTExportModelBody1272) SetType(v GBTBodyType)`

SetType sets Type field to given value.

### HasType

`func (o *BTExportModelBody1272) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVertices

`func (o *BTExportModelBody1272) GetVertices() []BTExportModelVertex858`

GetVertices returns the Vertices field if non-nil, zero value otherwise.

### GetVerticesOk

`func (o *BTExportModelBody1272) GetVerticesOk() (*[]BTExportModelVertex858, bool)`

GetVerticesOk returns a tuple with the Vertices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertices

`func (o *BTExportModelBody1272) SetVertices(v []BTExportModelVertex858)`

SetVertices sets Vertices field to given value.

### HasVertices

`func (o *BTExportModelBody1272) HasVertices() bool`

HasVertices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


