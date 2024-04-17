# \StepsTestingStepV1API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestingStepV1GetCompletedSteps**](StepsTestingStepV1API.md#TestingStepV1GetCompletedSteps) | **Get** /api/v1/steps/testing-step-v1/Completed/ | Retrieve all &#39;Completed&#39; &#39;testing-step-v1&#39; steps
[**TestingStepV1GetFailedSteps**](StepsTestingStepV1API.md#TestingStepV1GetFailedSteps) | **Get** /api/v1/steps/testing-step-v1/Failed/ | Retrieve all &#39;Failed&#39; &#39;testing-step-v1&#39; steps
[**TestingStepV1GetReadySteps**](StepsTestingStepV1API.md#TestingStepV1GetReadySteps) | **Get** /api/v1/steps/testing-step-v1/Ready/ | Retrieve all &#39;Ready&#39; &#39;testing-step-v1&#39; steps
[**TestingStepV1MarkStepAsCompleted**](StepsTestingStepV1API.md#TestingStepV1MarkStepAsCompleted) | **Put** /api/v1/steps/testing-step-v1/{id}/Completed/ | Mark the specified &#39;testing-step-v1&#39; step as &#39;Completed&#39;
[**TestingStepV1MarkStepAsFailed**](StepsTestingStepV1API.md#TestingStepV1MarkStepAsFailed) | **Put** /api/v1/steps/testing-step-v1/{id}/Failed/ | Mark the specified &#39;testing-step-v1&#39; step as &#39;Failed&#39;



## TestingStepV1GetCompletedSteps

> []TestingStepV1CompletedStepDto TestingStepV1GetCompletedSteps(ctx).Execute()

Retrieve all 'Completed' 'testing-step-v1' steps

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PartlyFluked/taskhub-test-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepsTestingStepV1API.TestingStepV1GetCompletedSteps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepsTestingStepV1API.TestingStepV1GetCompletedSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestingStepV1GetCompletedSteps`: []TestingStepV1CompletedStepDto
	fmt.Fprintf(os.Stdout, "Response from `StepsTestingStepV1API.TestingStepV1GetCompletedSteps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestingStepV1GetCompletedStepsRequest struct via the builder pattern


### Return type

[**[]TestingStepV1CompletedStepDto**](TestingStepV1CompletedStepDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestingStepV1GetFailedSteps

> []TestingStepV1FailedStepDto TestingStepV1GetFailedSteps(ctx).Execute()

Retrieve all 'Failed' 'testing-step-v1' steps

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PartlyFluked/taskhub-test-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepsTestingStepV1API.TestingStepV1GetFailedSteps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepsTestingStepV1API.TestingStepV1GetFailedSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestingStepV1GetFailedSteps`: []TestingStepV1FailedStepDto
	fmt.Fprintf(os.Stdout, "Response from `StepsTestingStepV1API.TestingStepV1GetFailedSteps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestingStepV1GetFailedStepsRequest struct via the builder pattern


### Return type

[**[]TestingStepV1FailedStepDto**](TestingStepV1FailedStepDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestingStepV1GetReadySteps

> []TestingStepV1ReadyStepDto TestingStepV1GetReadySteps(ctx).Execute()

Retrieve all 'Ready' 'testing-step-v1' steps

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PartlyFluked/taskhub-test-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepsTestingStepV1API.TestingStepV1GetReadySteps(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepsTestingStepV1API.TestingStepV1GetReadySteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestingStepV1GetReadySteps`: []TestingStepV1ReadyStepDto
	fmt.Fprintf(os.Stdout, "Response from `StepsTestingStepV1API.TestingStepV1GetReadySteps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTestingStepV1GetReadyStepsRequest struct via the builder pattern


### Return type

[**[]TestingStepV1ReadyStepDto**](TestingStepV1ReadyStepDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestingStepV1MarkStepAsCompleted

> TestingStepV1CompletedStepDto TestingStepV1MarkStepAsCompleted(ctx, id).Body(body).Execute()

Mark the specified 'testing-step-v1' step as 'Completed'

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PartlyFluked/taskhub-test-client"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the step to update
	body := map[string]interface{}{ ... } // map[string]interface{} | The new state of the step

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepsTestingStepV1API.TestingStepV1MarkStepAsCompleted(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepsTestingStepV1API.TestingStepV1MarkStepAsCompleted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestingStepV1MarkStepAsCompleted`: TestingStepV1CompletedStepDto
	fmt.Fprintf(os.Stdout, "Response from `StepsTestingStepV1API.TestingStepV1MarkStepAsCompleted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the step to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestingStepV1MarkStepAsCompletedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | The new state of the step | 

### Return type

[**TestingStepV1CompletedStepDto**](TestingStepV1CompletedStepDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestingStepV1MarkStepAsFailed

> TestingStepV1FailedStepDto TestingStepV1MarkStepAsFailed(ctx, id).TestingStepV1FailedResult(testingStepV1FailedResult).Execute()

Mark the specified 'testing-step-v1' step as 'Failed'

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PartlyFluked/taskhub-test-client"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the step to update
	testingStepV1FailedResult := *openapiclient.NewTestingStepV1FailedResult("Error_example") // TestingStepV1FailedResult | The new state of the step

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StepsTestingStepV1API.TestingStepV1MarkStepAsFailed(context.Background(), id).TestingStepV1FailedResult(testingStepV1FailedResult).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StepsTestingStepV1API.TestingStepV1MarkStepAsFailed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestingStepV1MarkStepAsFailed`: TestingStepV1FailedStepDto
	fmt.Fprintf(os.Stdout, "Response from `StepsTestingStepV1API.TestingStepV1MarkStepAsFailed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the step to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestingStepV1MarkStepAsFailedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testingStepV1FailedResult** | [**TestingStepV1FailedResult**](TestingStepV1FailedResult.md) | The new state of the step | 

### Return type

[**TestingStepV1FailedStepDto**](TestingStepV1FailedStepDto.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

