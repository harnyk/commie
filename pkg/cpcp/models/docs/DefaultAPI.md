# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetManifest**](DefaultAPI.md#GetManifest) | **Post** /manifest | Retrieve plugin manifest
[**GetSystemPrompt**](DefaultAPI.md#GetSystemPrompt) | **Post** /get_system_prompt | Retrieve the system prompt of the plugin
[**RunCommand**](DefaultAPI.md#RunCommand) | **Post** /run_command | Execute a command
[**RunTools**](DefaultAPI.md#RunTools) | **Post** /run_tools | Execute multiple tools in parallel



## GetManifest

> PluginManifest GetManifest(ctx).Execute()

Retrieve plugin manifest

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/models"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetManifest(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManifest`: PluginManifest
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetManifest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManifestRequest struct via the builder pattern


### Return type

[**PluginManifest**](PluginManifest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemPrompt

> GetSystemPrompt200Response GetSystemPrompt(ctx).Execute()

Retrieve the system prompt of the plugin

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/models"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetSystemPrompt(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetSystemPrompt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSystemPrompt`: GetSystemPrompt200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetSystemPrompt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemPromptRequest struct via the builder pattern


### Return type

[**GetSystemPrompt200Response**](GetSystemPrompt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunCommand

> RunCommand200Response RunCommand(ctx).RunCommandRequest(runCommandRequest).Execute()

Execute a command

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/models"
)

func main() {
	runCommandRequest := *openapiclient.NewRunCommandRequest("Command_example") // RunCommandRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RunCommand(context.Background()).RunCommandRequest(runCommandRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RunCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunCommand`: RunCommand200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RunCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runCommandRequest** | [**RunCommandRequest**](RunCommandRequest.md) |  | 

### Return type

[**RunCommand200Response**](RunCommand200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunTools

> RunTools200Response RunTools(ctx).RunToolsRequest(runToolsRequest).Execute()

Execute multiple tools in parallel

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/models"
)

func main() {
	runToolsRequest := *openapiclient.NewRunToolsRequest([]openapiclient.ToolCall{*openapiclient.NewToolCall("ToolName_example", map[string]interface{}{"key": interface{}(123)})}) // RunToolsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RunTools(context.Background()).RunToolsRequest(runToolsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RunTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunTools`: RunTools200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RunTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runToolsRequest** | [**RunToolsRequest**](RunToolsRequest.md) |  | 

### Return type

[**RunTools200Response**](RunTools200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

