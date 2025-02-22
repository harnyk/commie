# ToolCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolName** | **string** | Name of the tool to invoke | 
**Parameters** | **map[string]interface{}** | Input parameters for the tool | 

## Methods

### NewToolCall

`func NewToolCall(toolName string, parameters map[string]interface{}, ) *ToolCall`

NewToolCall instantiates a new ToolCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolCallWithDefaults

`func NewToolCallWithDefaults() *ToolCall`

NewToolCallWithDefaults instantiates a new ToolCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolName

`func (o *ToolCall) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *ToolCall) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *ToolCall) SetToolName(v string)`

SetToolName sets ToolName field to given value.


### GetParameters

`func (o *ToolCall) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ToolCall) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ToolCall) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


