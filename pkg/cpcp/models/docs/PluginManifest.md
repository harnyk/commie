# PluginManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Plugin name | 
**Description** | **string** | Plugin description | 
**Version** | **string** | Plugin version | 
**Repository** | Pointer to **string** | Repository URL (if available) | [optional] 
**Website** | Pointer to **string** | Plugin website (if available) | [optional] 
**Commands** | **[]string** | List of available command names | 
**Tools** | [**[]ToolManifest**](ToolManifest.md) | Full tool manifests (similar to OpenAI API) | 
**SystemPrompt** | Pointer to **NullableString** | Plugin-specific system prompt | [optional] 

## Methods

### NewPluginManifest

`func NewPluginManifest(name string, description string, version string, commands []string, tools []ToolManifest, ) *PluginManifest`

NewPluginManifest instantiates a new PluginManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginManifestWithDefaults

`func NewPluginManifestWithDefaults() *PluginManifest`

NewPluginManifestWithDefaults instantiates a new PluginManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginManifest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginManifest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginManifest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PluginManifest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginManifest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginManifest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *PluginManifest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginManifest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginManifest) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetRepository

`func (o *PluginManifest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *PluginManifest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *PluginManifest) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *PluginManifest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetWebsite

`func (o *PluginManifest) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *PluginManifest) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *PluginManifest) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *PluginManifest) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetCommands

`func (o *PluginManifest) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *PluginManifest) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *PluginManifest) SetCommands(v []string)`

SetCommands sets Commands field to given value.


### GetTools

`func (o *PluginManifest) GetTools() []ToolManifest`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *PluginManifest) GetToolsOk() (*[]ToolManifest, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *PluginManifest) SetTools(v []ToolManifest)`

SetTools sets Tools field to given value.


### GetSystemPrompt

`func (o *PluginManifest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *PluginManifest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *PluginManifest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *PluginManifest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *PluginManifest) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *PluginManifest) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


