# TestingStepV1StepDto

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

## Methods

### NewTestingStepV1StepDto

`func NewTestingStepV1StepDto(id string, key string, name string, jobId string, milestoneId string, dependedOnBy []string, dependsOn []string, stepType TestingStepV1StepType, parameters map[string]string, status TestingStepV1StepStatus, ) *TestingStepV1StepDto`

NewTestingStepV1StepDto instantiates a new TestingStepV1StepDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestingStepV1StepDtoWithDefaults

`func NewTestingStepV1StepDtoWithDefaults() *TestingStepV1StepDto`

NewTestingStepV1StepDtoWithDefaults instantiates a new TestingStepV1StepDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TestingStepV1StepDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestingStepV1StepDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestingStepV1StepDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *TestingStepV1StepDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TestingStepV1StepDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TestingStepV1StepDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *TestingStepV1StepDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestingStepV1StepDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestingStepV1StepDto) SetName(v string)`

SetName sets Name field to given value.


### GetJobId

`func (o *TestingStepV1StepDto) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *TestingStepV1StepDto) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *TestingStepV1StepDto) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetMilestoneId

`func (o *TestingStepV1StepDto) GetMilestoneId() string`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *TestingStepV1StepDto) GetMilestoneIdOk() (*string, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *TestingStepV1StepDto) SetMilestoneId(v string)`

SetMilestoneId sets MilestoneId field to given value.


### GetDependedOnBy

`func (o *TestingStepV1StepDto) GetDependedOnBy() []string`

GetDependedOnBy returns the DependedOnBy field if non-nil, zero value otherwise.

### GetDependedOnByOk

`func (o *TestingStepV1StepDto) GetDependedOnByOk() (*[]string, bool)`

GetDependedOnByOk returns a tuple with the DependedOnBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependedOnBy

`func (o *TestingStepV1StepDto) SetDependedOnBy(v []string)`

SetDependedOnBy sets DependedOnBy field to given value.


### GetDependsOn

`func (o *TestingStepV1StepDto) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *TestingStepV1StepDto) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *TestingStepV1StepDto) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.


### GetStepType

`func (o *TestingStepV1StepDto) GetStepType() TestingStepV1StepType`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *TestingStepV1StepDto) GetStepTypeOk() (*TestingStepV1StepType, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *TestingStepV1StepDto) SetStepType(v TestingStepV1StepType)`

SetStepType sets StepType field to given value.


### GetParameters

`func (o *TestingStepV1StepDto) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *TestingStepV1StepDto) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *TestingStepV1StepDto) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetStatus

`func (o *TestingStepV1StepDto) GetStatus() TestingStepV1StepStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestingStepV1StepDto) GetStatusOk() (*TestingStepV1StepStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestingStepV1StepDto) SetStatus(v TestingStepV1StepStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


