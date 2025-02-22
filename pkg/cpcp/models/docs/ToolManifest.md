# ToolManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Tool name | 
**Description** | **string** | Tool description | 
**Parameters** | **map[string]interface{}** | Description of input parameters (similar to OpenAI API) | 

## Methods

### NewToolManifest

`func NewToolManifest(name string, description string, parameters map[string]interface{}, ) *ToolManifest`

NewToolManifest instantiates a new ToolManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolManifestWithDefaults

`func NewToolManifestWithDefaults() *ToolManifest`

NewToolManifestWithDefaults instantiates a new ToolManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ToolManifest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolManifest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolManifest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ToolManifest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ToolManifest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ToolManifest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetParameters

`func (o *ToolManifest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ToolManifest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ToolManifest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


