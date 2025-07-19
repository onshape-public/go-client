# BTModelFormatFullInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | Content-Type for this file format. | [optional] 
**CouldBeAssembly** | Pointer to **bool** | Indicates if this format could be an assembly. | [optional] 
**FileExtensions** | Pointer to **[]string** | Supported file extensions for this format. | [optional] 
**Name** | Pointer to **string** | Name of the format. | [optional] 
**TranslatorName** | Pointer to **string** | The name of the translator for the format. | [optional] 
**ValidDestinationFormat** | Pointer to **bool** | Indicates if this format is a valid destination format for translation. | [optional] 
**ValidSourceFormat** | Pointer to **bool** | Indicates if this format is a valid source format for translation. | [optional] 

## Methods

### NewBTModelFormatFullInfo

`func NewBTModelFormatFullInfo() *BTModelFormatFullInfo`

NewBTModelFormatFullInfo instantiates a new BTModelFormatFullInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBTModelFormatFullInfoWithDefaults

`func NewBTModelFormatFullInfoWithDefaults() *BTModelFormatFullInfo`

NewBTModelFormatFullInfoWithDefaults instantiates a new BTModelFormatFullInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *BTModelFormatFullInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BTModelFormatFullInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BTModelFormatFullInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BTModelFormatFullInfo) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetCouldBeAssembly

`func (o *BTModelFormatFullInfo) GetCouldBeAssembly() bool`

GetCouldBeAssembly returns the CouldBeAssembly field if non-nil, zero value otherwise.

### GetCouldBeAssemblyOk

`func (o *BTModelFormatFullInfo) GetCouldBeAssemblyOk() (*bool, bool)`

GetCouldBeAssemblyOk returns a tuple with the CouldBeAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouldBeAssembly

`func (o *BTModelFormatFullInfo) SetCouldBeAssembly(v bool)`

SetCouldBeAssembly sets CouldBeAssembly field to given value.

### HasCouldBeAssembly

`func (o *BTModelFormatFullInfo) HasCouldBeAssembly() bool`

HasCouldBeAssembly returns a boolean if a field has been set.

### GetFileExtensions

`func (o *BTModelFormatFullInfo) GetFileExtensions() []string`

GetFileExtensions returns the FileExtensions field if non-nil, zero value otherwise.

### GetFileExtensionsOk

`func (o *BTModelFormatFullInfo) GetFileExtensionsOk() (*[]string, bool)`

GetFileExtensionsOk returns a tuple with the FileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensions

`func (o *BTModelFormatFullInfo) SetFileExtensions(v []string)`

SetFileExtensions sets FileExtensions field to given value.

### HasFileExtensions

`func (o *BTModelFormatFullInfo) HasFileExtensions() bool`

HasFileExtensions returns a boolean if a field has been set.

### GetName

`func (o *BTModelFormatFullInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BTModelFormatFullInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BTModelFormatFullInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BTModelFormatFullInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTranslatorName

`func (o *BTModelFormatFullInfo) GetTranslatorName() string`

GetTranslatorName returns the TranslatorName field if non-nil, zero value otherwise.

### GetTranslatorNameOk

`func (o *BTModelFormatFullInfo) GetTranslatorNameOk() (*string, bool)`

GetTranslatorNameOk returns a tuple with the TranslatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatorName

`func (o *BTModelFormatFullInfo) SetTranslatorName(v string)`

SetTranslatorName sets TranslatorName field to given value.

### HasTranslatorName

`func (o *BTModelFormatFullInfo) HasTranslatorName() bool`

HasTranslatorName returns a boolean if a field has been set.

### GetValidDestinationFormat

`func (o *BTModelFormatFullInfo) GetValidDestinationFormat() bool`

GetValidDestinationFormat returns the ValidDestinationFormat field if non-nil, zero value otherwise.

### GetValidDestinationFormatOk

`func (o *BTModelFormatFullInfo) GetValidDestinationFormatOk() (*bool, bool)`

GetValidDestinationFormatOk returns a tuple with the ValidDestinationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDestinationFormat

`func (o *BTModelFormatFullInfo) SetValidDestinationFormat(v bool)`

SetValidDestinationFormat sets ValidDestinationFormat field to given value.

### HasValidDestinationFormat

`func (o *BTModelFormatFullInfo) HasValidDestinationFormat() bool`

HasValidDestinationFormat returns a boolean if a field has been set.

### GetValidSourceFormat

`func (o *BTModelFormatFullInfo) GetValidSourceFormat() bool`

GetValidSourceFormat returns the ValidSourceFormat field if non-nil, zero value otherwise.

### GetValidSourceFormatOk

`func (o *BTModelFormatFullInfo) GetValidSourceFormatOk() (*bool, bool)`

GetValidSourceFormatOk returns a tuple with the ValidSourceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidSourceFormat

`func (o *BTModelFormatFullInfo) SetValidSourceFormat(v bool)`

SetValidSourceFormat sets ValidSourceFormat field to given value.

### HasValidSourceFormat

`func (o *BTModelFormatFullInfo) HasValidSourceFormat() bool`

HasValidSourceFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


