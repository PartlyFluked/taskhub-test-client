# TestingStepV1ReadyStepDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | **string** |  | 
**Name** | **string** |  | 
**JobId** | **string** |  | 
**MilestoneId** | **string** |  | 
**DependedOnBy** | **[]string** |  | 
**DependsOn** | **[]string** |  | 
**StepType** | [**TestingStepV1StepType**](TestingStepV1StepType.md) |  | 
**Parameters** | **map[string]string** |  | 
**Status** | [**TestingStepV1StepStatus**](TestingStepV1StepStatus.md) |  | 
**Result** | **map[string]interface{}** |  | 

## Methods

### NewTestingStepV1ReadyStepDto

`func NewTestingStepV1ReadyStepDto(id string, key string, name string, jobId string, milestoneId string, dependedOnBy []string, dependsOn []string, stepType TestingStepV1StepType, parameters map[string]string, status TestingStepV1StepStatus, result map[string]interface{}, ) *TestingStepV1ReadyStepDto`

NewTestingStepV1ReadyStepDto instantiates a new TestingStepV1ReadyStepDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestingStepV1ReadyStepDtoWithDefaults

`func NewTestingStepV1ReadyStepDtoWithDefaults() *TestingStepV1ReadyStepDto`

NewTestingStepV1ReadyStepDtoWithDefaults instantiates a new TestingStepV1ReadyStepDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestingStepV1ReadyStepDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestingStepV1ReadyStepDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestingStepV1ReadyStepDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *TestingStepV1ReadyStepDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TestingStepV1ReadyStepDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TestingStepV1ReadyStepDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *TestingStepV1ReadyStepDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestingStepV1ReadyStepDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestingStepV1ReadyStepDto) SetName(v string)`

SetName sets Name field to given value.


### GetJobId

`func (o *TestingStepV1ReadyStepDto) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *TestingStepV1ReadyStepDto) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *TestingStepV1ReadyStepDto) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetMilestoneId

`func (o *TestingStepV1ReadyStepDto) GetMilestoneId() string`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *TestingStepV1ReadyStepDto) GetMilestoneIdOk() (*string, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *TestingStepV1ReadyStepDto) SetMilestoneId(v string)`

SetMilestoneId sets MilestoneId field to given value.


### GetDependedOnBy

`func (o *TestingStepV1ReadyStepDto) GetDependedOnBy() []string`

GetDependedOnBy returns the DependedOnBy field if non-nil, zero value otherwise.

### GetDependedOnByOk

`func (o *TestingStepV1ReadyStepDto) GetDependedOnByOk() (*[]string, bool)`

GetDependedOnByOk returns a tuple with the DependedOnBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependedOnBy

`func (o *TestingStepV1ReadyStepDto) SetDependedOnBy(v []string)`

SetDependedOnBy sets DependedOnBy field to given value.


### GetDependsOn

`func (o *TestingStepV1ReadyStepDto) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *TestingStepV1ReadyStepDto) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *TestingStepV1ReadyStepDto) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.


### GetStepType

`func (o *TestingStepV1ReadyStepDto) GetStepType() TestingStepV1StepType`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *TestingStepV1ReadyStepDto) GetStepTypeOk() (*TestingStepV1StepType, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *TestingStepV1ReadyStepDto) SetStepType(v TestingStepV1StepType)`

SetStepType sets StepType field to given value.


### GetParameters

`func (o *TestingStepV1ReadyStepDto) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestingStepV1ReadyStepDto) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestingStepV1ReadyStepDto) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetStatus

`func (o *TestingStepV1ReadyStepDto) GetStatus() TestingStepV1StepStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestingStepV1ReadyStepDto) GetStatusOk() (*TestingStepV1StepStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestingStepV1ReadyStepDto) SetStatus(v TestingStepV1StepStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *TestingStepV1ReadyStepDto) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestingStepV1ReadyStepDto) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestingStepV1ReadyStepDto) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *TestingStepV1ReadyStepDto) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *TestingStepV1ReadyStepDto) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


