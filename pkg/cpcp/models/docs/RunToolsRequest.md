# RunToolsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolCalls** | [**[]ToolCall**](ToolCall.md) |  | 

## Methods

### NewRunToolsRequest

`func NewRunToolsRequest(toolCalls []ToolCall, ) *RunToolsRequest`

NewRunToolsRequest instantiates a new RunToolsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunToolsRequestWithDefaults

`func NewRunToolsRequestWithDefaults() *RunToolsRequest`

NewRunToolsRequestWithDefaults instantiates a new RunToolsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolCalls

`func (o *RunToolsRequest) GetToolCalls() []ToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *RunToolsRequest) GetToolCallsOk() (*[]ToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *RunToolsRequest) SetToolCalls(v []ToolCall)`

SetToolCalls sets ToolCalls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


