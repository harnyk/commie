# ToolOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolName** | **string** | Name of the executed tool | 
**Result** | **map[string]interface{}** | Tool execution results | 

## Methods

### NewToolOutput

`func NewToolOutput(toolName string, result map[string]interface{}, ) *ToolOutput`

NewToolOutput instantiates a new ToolOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolOutputWithDefaults

`func NewToolOutputWithDefaults() *ToolOutput`

NewToolOutputWithDefaults instantiates a new ToolOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolName

`func (o *ToolOutput) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *ToolOutput) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *ToolOutput) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### GetResult

`func (o *ToolOutput) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ToolOutput) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ToolOutput) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


