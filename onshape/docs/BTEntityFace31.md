# BTEntityFace31

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** | Type of JSON object. | [optional] 
**CompressedUvs** | Pointer to [**BTImmutableByteArray**](BTImmutableByteArray.md) |  | [optional] 
**FlipComputedNormals** | Pointer to **bool** |  | [optional] 
**Indices** | Pointer to [**BTImmutableIntArray**](BTImmutableIntArray.md) |  | [optional] 
**IndicesStoredAsDifferences** | Pointer to **bool** |  | [optional] 
**IsPlanar** | Pointer to **bool** |  | [optional] 
**MaxPrincipleCurvatureMagnitudes** | Pointer to [**BTImmutableFloatArray**](BTImmutableFloatArray.md) |  | [optional] 
**MinPrincipleCurvatureMagnitudes** | Pointer to [**BTImmutableFloatArray**](BTImmutableFloatArray.md) |  | [optional] 
**Normals** | Pointer to [**BTImmutableFloatArray**](BTImmutableFloatArray.md) |  | [optional] 
**Points** | Pointer to [**BTImmutableFloatArray**](BTImmutableFloatArray.md) |  | [optional] 
**SurfaceParameters** | Pointer to [**BTImmutableDoubleArray**](BTImmutableDoubleArray.md) |  | [optional] 
**SurfaceType** | Pointer to [**GBTSurfaceType**](GBTSurfaceType.md) |  | [optional] 
**TextureCoordinates** | Pointer to [**BTImmutableFloatArray**](BTImmutableFloatArray.md) |  | [optional] 
**TriangleCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTEntityFace31

`func NewBTEntityFace31() *BTEntityFace31`

NewBTEntityFace31 instantiates a new BTEntityFace31 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTEntityFace31WithDefaults

`func NewBTEntityFace31WithDefaults() *BTEntityFace31`

NewBTEntityFace31WithDefaults instantiates a new BTEntityFace31 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTEntityFace31) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTEntityFace31) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTEntityFace31) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTEntityFace31) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetCompressedUvs

`func (o *BTEntityFace31) GetCompressedUvs() BTImmutableByteArray`

GetCompressedUvs returns the CompressedUvs field if non-nil, zero value otherwise.

### GetCompressedUvsOk

`func (o *BTEntityFace31) GetCompressedUvsOk() (*BTImmutableByteArray, bool)`

GetCompressedUvsOk returns a tuple with the CompressedUvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedUvs

`func (o *BTEntityFace31) SetCompressedUvs(v BTImmutableByteArray)`

SetCompressedUvs sets CompressedUvs field to given value.

### HasCompressedUvs

`func (o *BTEntityFace31) HasCompressedUvs() bool`

HasCompressedUvs returns a boolean if a field has been set.

### GetFlipComputedNormals

`func (o *BTEntityFace31) GetFlipComputedNormals() bool`

GetFlipComputedNormals returns the FlipComputedNormals field if non-nil, zero value otherwise.

### GetFlipComputedNormalsOk

`func (o *BTEntityFace31) GetFlipComputedNormalsOk() (*bool, bool)`

GetFlipComputedNormalsOk returns a tuple with the FlipComputedNormals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlipComputedNormals

`func (o *BTEntityFace31) SetFlipComputedNormals(v bool)`

SetFlipComputedNormals sets FlipComputedNormals field to given value.

### HasFlipComputedNormals

`func (o *BTEntityFace31) HasFlipComputedNormals() bool`

HasFlipComputedNormals returns a boolean if a field has been set.

### GetIndices

`func (o *BTEntityFace31) GetIndices() BTImmutableIntArray`

GetIndices returns the Indices field if non-nil, zero value otherwise.

### GetIndicesOk

`func (o *BTEntityFace31) GetIndicesOk() (*BTImmutableIntArray, bool)`

GetIndicesOk returns a tuple with the Indices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndices

`func (o *BTEntityFace31) SetIndices(v BTImmutableIntArray)`

SetIndices sets Indices field to given value.

### HasIndices

`func (o *BTEntityFace31) HasIndices() bool`

HasIndices returns a boolean if a field has been set.

### GetIndicesStoredAsDifferences

`func (o *BTEntityFace31) GetIndicesStoredAsDifferences() bool`

GetIndicesStoredAsDifferences returns the IndicesStoredAsDifferences field if non-nil, zero value otherwise.

### GetIndicesStoredAsDifferencesOk

`func (o *BTEntityFace31) GetIndicesStoredAsDifferencesOk() (*bool, bool)`

GetIndicesStoredAsDifferencesOk returns a tuple with the IndicesStoredAsDifferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicesStoredAsDifferences

`func (o *BTEntityFace31) SetIndicesStoredAsDifferences(v bool)`

SetIndicesStoredAsDifferences sets IndicesStoredAsDifferences field to given value.

### HasIndicesStoredAsDifferences

`func (o *BTEntityFace31) HasIndicesStoredAsDifferences() bool`

HasIndicesStoredAsDifferences returns a boolean if a field has been set.

### GetIsPlanar

`func (o *BTEntityFace31) GetIsPlanar() bool`

GetIsPlanar returns the IsPlanar field if non-nil, zero value otherwise.

### GetIsPlanarOk

`func (o *BTEntityFace31) GetIsPlanarOk() (*bool, bool)`

GetIsPlanarOk returns a tuple with the IsPlanar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlanar

`func (o *BTEntityFace31) SetIsPlanar(v bool)`

SetIsPlanar sets IsPlanar field to given value.

### HasIsPlanar

`func (o *BTEntityFace31) HasIsPlanar() bool`

HasIsPlanar returns a boolean if a field has been set.

### GetMaxPrincipleCurvatureMagnitudes

`func (o *BTEntityFace31) GetMaxPrincipleCurvatureMagnitudes() BTImmutableFloatArray`

GetMaxPrincipleCurvatureMagnitudes returns the MaxPrincipleCurvatureMagnitudes field if non-nil, zero value otherwise.

### GetMaxPrincipleCurvatureMagnitudesOk

`func (o *BTEntityFace31) GetMaxPrincipleCurvatureMagnitudesOk() (*BTImmutableFloatArray, bool)`

GetMaxPrincipleCurvatureMagnitudesOk returns a tuple with the MaxPrincipleCurvatureMagnitudes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrincipleCurvatureMagnitudes

`func (o *BTEntityFace31) SetMaxPrincipleCurvatureMagnitudes(v BTImmutableFloatArray)`

SetMaxPrincipleCurvatureMagnitudes sets MaxPrincipleCurvatureMagnitudes field to given value.

### HasMaxPrincipleCurvatureMagnitudes

`func (o *BTEntityFace31) HasMaxPrincipleCurvatureMagnitudes() bool`

HasMaxPrincipleCurvatureMagnitudes returns a boolean if a field has been set.

### GetMinPrincipleCurvatureMagnitudes

`func (o *BTEntityFace31) GetMinPrincipleCurvatureMagnitudes() BTImmutableFloatArray`

GetMinPrincipleCurvatureMagnitudes returns the MinPrincipleCurvatureMagnitudes field if non-nil, zero value otherwise.

### GetMinPrincipleCurvatureMagnitudesOk

`func (o *BTEntityFace31) GetMinPrincipleCurvatureMagnitudesOk() (*BTImmutableFloatArray, bool)`

GetMinPrincipleCurvatureMagnitudesOk returns a tuple with the MinPrincipleCurvatureMagnitudes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrincipleCurvatureMagnitudes

`func (o *BTEntityFace31) SetMinPrincipleCurvatureMagnitudes(v BTImmutableFloatArray)`

SetMinPrincipleCurvatureMagnitudes sets MinPrincipleCurvatureMagnitudes field to given value.

### HasMinPrincipleCurvatureMagnitudes

`func (o *BTEntityFace31) HasMinPrincipleCurvatureMagnitudes() bool`

HasMinPrincipleCurvatureMagnitudes returns a boolean if a field has been set.

### GetNormals

`func (o *BTEntityFace31) GetNormals() BTImmutableFloatArray`

GetNormals returns the Normals field if non-nil, zero value otherwise.

### GetNormalsOk

`func (o *BTEntityFace31) GetNormalsOk() (*BTImmutableFloatArray, bool)`

GetNormalsOk returns a tuple with the Normals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormals

`func (o *BTEntityFace31) SetNormals(v BTImmutableFloatArray)`

SetNormals sets Normals field to given value.

### HasNormals

`func (o *BTEntityFace31) HasNormals() bool`

HasNormals returns a boolean if a field has been set.

### GetPoints

`func (o *BTEntityFace31) GetPoints() BTImmutableFloatArray`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *BTEntityFace31) GetPointsOk() (*BTImmutableFloatArray, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *BTEntityFace31) SetPoints(v BTImmutableFloatArray)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *BTEntityFace31) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetSurfaceParameters

`func (o *BTEntityFace31) GetSurfaceParameters() BTImmutableDoubleArray`

GetSurfaceParameters returns the SurfaceParameters field if non-nil, zero value otherwise.

### GetSurfaceParametersOk

`func (o *BTEntityFace31) GetSurfaceParametersOk() (*BTImmutableDoubleArray, bool)`

GetSurfaceParametersOk returns a tuple with the SurfaceParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurfaceParameters

`func (o *BTEntityFace31) SetSurfaceParameters(v BTImmutableDoubleArray)`

SetSurfaceParameters sets SurfaceParameters field to given value.

### HasSurfaceParameters

`func (o *BTEntityFace31) HasSurfaceParameters() bool`

HasSurfaceParameters returns a boolean if a field has been set.

### GetSurfaceType

`func (o *BTEntityFace31) GetSurfaceType() GBTSurfaceType`

GetSurfaceType returns the SurfaceType field if non-nil, zero value otherwise.

### GetSurfaceTypeOk

`func (o *BTEntityFace31) GetSurfaceTypeOk() (*GBTSurfaceType, bool)`

GetSurfaceTypeOk returns a tuple with the SurfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurfaceType

`func (o *BTEntityFace31) SetSurfaceType(v GBTSurfaceType)`

SetSurfaceType sets SurfaceType field to given value.

### HasSurfaceType

`func (o *BTEntityFace31) HasSurfaceType() bool`

HasSurfaceType returns a boolean if a field has been set.

### GetTextureCoordinates

`func (o *BTEntityFace31) GetTextureCoordinates() BTImmutableFloatArray`

GetTextureCoordinates returns the TextureCoordinates field if non-nil, zero value otherwise.

### GetTextureCoordinatesOk

`func (o *BTEntityFace31) GetTextureCoordinatesOk() (*BTImmutableFloatArray, bool)`

GetTextureCoordinatesOk returns a tuple with the TextureCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextureCoordinates

`func (o *BTEntityFace31) SetTextureCoordinates(v BTImmutableFloatArray)`

SetTextureCoordinates sets TextureCoordinates field to given value.

### HasTextureCoordinates

`func (o *BTEntityFace31) HasTextureCoordinates() bool`

HasTextureCoordinates returns a boolean if a field has been set.

### GetTriangleCount

`func (o *BTEntityFace31) GetTriangleCount() int32`

GetTriangleCount returns the TriangleCount field if non-nil, zero value otherwise.

### GetTriangleCountOk

`func (o *BTEntityFace31) GetTriangleCountOk() (*int32, bool)`

GetTriangleCountOk returns a tuple with the TriangleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriangleCount

`func (o *BTEntityFace31) SetTriangleCount(v int32)`

SetTriangleCount sets TriangleCount field to given value.

### HasTriangleCount

`func (o *BTEntityFace31) HasTriangleCount() bool`

HasTriangleCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


