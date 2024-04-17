# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.5.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/PartlyFluked/taskhub-test-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*StepsTestingStepV1API* | [**TestingStepV1GetCompletedSteps**](docs/StepsTestingStepV1API.md#testingstepv1getcompletedsteps) | **Get** /api/v1/steps/testing-step-v1/Completed/ | Retrieve all &#39;Completed&#39; &#39;testing-step-v1&#39; steps
*StepsTestingStepV1API* | [**TestingStepV1GetFailedSteps**](docs/StepsTestingStepV1API.md#testingstepv1getfailedsteps) | **Get** /api/v1/steps/testing-step-v1/Failed/ | Retrieve all &#39;Failed&#39; &#39;testing-step-v1&#39; steps
*StepsTestingStepV1API* | [**TestingStepV1GetReadySteps**](docs/StepsTestingStepV1API.md#testingstepv1getreadysteps) | **Get** /api/v1/steps/testing-step-v1/Ready/ | Retrieve all &#39;Ready&#39; &#39;testing-step-v1&#39; steps
*StepsTestingStepV1API* | [**TestingStepV1MarkStepAsCompleted**](docs/StepsTestingStepV1API.md#testingstepv1markstepascompleted) | **Put** /api/v1/steps/testing-step-v1/{id}/Completed/ | Mark the specified &#39;testing-step-v1&#39; step as &#39;Completed&#39;
*StepsTestingStepV1API* | [**TestingStepV1MarkStepAsFailed**](docs/StepsTestingStepV1API.md#testingstepv1markstepasfailed) | **Put** /api/v1/steps/testing-step-v1/{id}/Failed/ | Mark the specified &#39;testing-step-v1&#39; step as &#39;Failed&#39;


## Documentation For Models

 - [BaseStepDto](docs/BaseStepDto.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [StepStatus](docs/StepStatus.md)
 - [StepType](docs/StepType.md)
 - [TestingStepV1CompletedStepDto](docs/TestingStepV1CompletedStepDto.md)
 - [TestingStepV1FailedResult](docs/TestingStepV1FailedResult.md)
 - [TestingStepV1FailedStepDto](docs/TestingStepV1FailedStepDto.md)
 - [TestingStepV1ReadyStepDto](docs/TestingStepV1ReadyStepDto.md)
 - [TestingStepV1StepDto](docs/TestingStepV1StepDto.md)
 - [TestingStepV1StepStatus](docs/TestingStepV1StepStatus.md)
 - [TestingStepV1StepType](docs/TestingStepV1StepType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, openapi.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


