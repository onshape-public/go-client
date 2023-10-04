# BTPModule234AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BtType** | Pointer to **string** |  | [optional] 
**DeepImports** | Pointer to [**map[string][]BTImport**](array.md) |  | [optional] 
**Imports** | Pointer to [**[]BTPTopLevelImport285**](BTPTopLevelImport285.md) |  | [optional] 
**IsBlob** | Pointer to **bool** |  | [optional] 
**IsInternalModule** | Pointer to **bool** |  | [optional] 
**MayHaveImplicitImports** | Pointer to **bool** |  | [optional] 
**PathMap** | Pointer to [**map[string]BTMicroversionId366**](BTMicroversionId366.md) |  | [optional] 
**ToBeParsed** | Pointer to [**BTLazilyParsedFeatureScript**](BTLazilyParsedFeatureScript.md) |  | [optional] 
**TopLevel** | Pointer to [**[]BTPTopLevelNode286**](BTPTopLevelNode286.md) |  | [optional] 
**Version** | Pointer to [**BTPLiteralNumber258**](BTPLiteralNumber258.md) |  | [optional] 
**VersionNumber** | Pointer to **int32** |  | [optional] 

## Methods

### NewBTPModule234AllOf

`func NewBTPModule234AllOf() *BTPModule234AllOf`

NewBTPModule234AllOf instantiates a new BTPModule234AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTPModule234AllOfWithDefaults

`func NewBTPModule234AllOfWithDefaults() *BTPModule234AllOf`

NewBTPModule234AllOfWithDefaults instantiates a new BTPModule234AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBtType

`func (o *BTPModule234AllOf) GetBtType() string`

GetBtType returns the BtType field if non-nil, zero value otherwise.

### GetBtTypeOk

`func (o *BTPModule234AllOf) GetBtTypeOk() (*string, bool)`

GetBtTypeOk returns a tuple with the BtType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtType

`func (o *BTPModule234AllOf) SetBtType(v string)`

SetBtType sets BtType field to given value.

### HasBtType

`func (o *BTPModule234AllOf) HasBtType() bool`

HasBtType returns a boolean if a field has been set.

### GetDeepImports

`func (o *BTPModule234AllOf) GetDeepImports() map[string][]BTImport`

GetDeepImports returns the DeepImports field if non-nil, zero value otherwise.

### GetDeepImportsOk

`func (o *BTPModule234AllOf) GetDeepImportsOk() (*map[string][]BTImport, bool)`

GetDeepImportsOk returns a tuple with the DeepImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepImports

`func (o *BTPModule234AllOf) SetDeepImports(v map[string][]BTImport)`

SetDeepImports sets DeepImports field to given value.

### HasDeepImports

`func (o *BTPModule234AllOf) HasDeepImports() bool`

HasDeepImports returns a boolean if a field has been set.

### GetImports

`func (o *BTPModule234AllOf) GetImports() []BTPTopLevelImport285`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *BTPModule234AllOf) GetImportsOk() (*[]BTPTopLevelImport285, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *BTPModule234AllOf) SetImports(v []BTPTopLevelImport285)`

SetImports sets Imports field to given value.

### HasImports

`func (o *BTPModule234AllOf) HasImports() bool`

HasImports returns a boolean if a field has been set.

### GetIsBlob

`func (o *BTPModule234AllOf) GetIsBlob() bool`

GetIsBlob returns the IsBlob field if non-nil, zero value otherwise.

### GetIsBlobOk

`func (o *BTPModule234AllOf) GetIsBlobOk() (*bool, bool)`

GetIsBlobOk returns a tuple with the IsBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlob

`func (o *BTPModule234AllOf) SetIsBlob(v bool)`

SetIsBlob sets IsBlob field to given value.

### HasIsBlob

`func (o *BTPModule234AllOf) HasIsBlob() bool`

HasIsBlob returns a boolean if a field has been set.

### GetIsInternalModule

`func (o *BTPModule234AllOf) GetIsInternalModule() bool`

GetIsInternalModule returns the IsInternalModule field if non-nil, zero value otherwise.

### GetIsInternalModuleOk

`func (o *BTPModule234AllOf) GetIsInternalModuleOk() (*bool, bool)`

GetIsInternalModuleOk returns a tuple with the IsInternalModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInternalModule

`func (o *BTPModule234AllOf) SetIsInternalModule(v bool)`

SetIsInternalModule sets IsInternalModule field to given value.

### HasIsInternalModule

`func (o *BTPModule234AllOf) HasIsInternalModule() bool`

HasIsInternalModule returns a boolean if a field has been set.

### GetMayHaveImplicitImports

`func (o *BTPModule234AllOf) GetMayHaveImplicitImports() bool`

GetMayHaveImplicitImports returns the MayHaveImplicitImports field if non-nil, zero value otherwise.

### GetMayHaveImplicitImportsOk

`func (o *BTPModule234AllOf) GetMayHaveImplicitImportsOk() (*bool, bool)`

GetMayHaveImplicitImportsOk returns a tuple with the MayHaveImplicitImports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayHaveImplicitImports

`func (o *BTPModule234AllOf) SetMayHaveImplicitImports(v bool)`

SetMayHaveImplicitImports sets MayHaveImplicitImports field to given value.

### HasMayHaveImplicitImports

`func (o *BTPModule234AllOf) HasMayHaveImplicitImports() bool`

HasMayHaveImplicitImports returns a boolean if a field has been set.

### GetPathMap

`func (o *BTPModule234AllOf) GetPathMap() map[string]BTMicroversionId366`

GetPathMap returns the PathMap field if non-nil, zero value otherwise.

### GetPathMapOk

`func (o *BTPModule234AllOf) GetPathMapOk() (*map[string]BTMicroversionId366, bool)`

GetPathMapOk returns a tuple with the PathMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathMap

`func (o *BTPModule234AllOf) SetPathMap(v map[string]BTMicroversionId366)`

SetPathMap sets PathMap field to given value.

### HasPathMap

`func (o *BTPModule234AllOf) HasPathMap() bool`

HasPathMap returns a boolean if a field has been set.

### GetToBeParsed

`func (o *BTPModule234AllOf) GetToBeParsed() BTLazilyParsedFeatureScript`

GetToBeParsed returns the ToBeParsed field if non-nil, zero value otherwise.

### GetToBeParsedOk

`func (o *BTPModule234AllOf) GetToBeParsedOk() (*BTLazilyParsedFeatureScript, bool)`

GetToBeParsedOk returns a tuple with the ToBeParsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBeParsed

`func (o *BTPModule234AllOf) SetToBeParsed(v BTLazilyParsedFeatureScript)`

SetToBeParsed sets ToBeParsed field to given value.

### HasToBeParsed

`func (o *BTPModule234AllOf) HasToBeParsed() bool`

HasToBeParsed returns a boolean if a field has been set.

### GetTopLevel

`func (o *BTPModule234AllOf) GetTopLevel() []BTPTopLevelNode286`

GetTopLevel returns the TopLevel field if non-nil, zero value otherwise.

### GetTopLevelOk

`func (o *BTPModule234AllOf) GetTopLevelOk() (*[]BTPTopLevelNode286, bool)`

GetTopLevelOk returns a tuple with the TopLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevel

`func (o *BTPModule234AllOf) SetTopLevel(v []BTPTopLevelNode286)`

SetTopLevel sets TopLevel field to given value.

### HasTopLevel

`func (o *BTPModule234AllOf) HasTopLevel() bool`

HasTopLevel returns a boolean if a field has been set.

### GetVersion

`func (o *BTPModule234AllOf) GetVersion() BTPLiteralNumber258`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BTPModule234AllOf) GetVersionOk() (*BTPLiteralNumber258, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BTPModule234AllOf) SetVersion(v BTPLiteralNumber258)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BTPModule234AllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionNumber

`func (o *BTPModule234AllOf) GetVersionNumber() int32`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *BTPModule234AllOf) GetVersionNumberOk() (*int32, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *BTPModule234AllOf) SetVersionNumber(v int32)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *BTPModule234AllOf) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


