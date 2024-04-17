# BaseStepDto

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

## Methods

### NewBaseStepDto

`func NewBaseStepDto(id string, key string, name string, jobId string, milestoneId string, dependedOnBy []string, dependsOn []string, ) *BaseStepDto`

NewBaseStepDto instantiates a new BaseStepDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseStepDtoWithDefaults

`func NewBaseStepDtoWithDefaults() *BaseStepDto`

NewBaseStepDtoWithDefaults instantiates a new BaseStepDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseStepDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseStepDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseStepDto) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *BaseStepDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BaseStepDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BaseStepDto) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *BaseStepDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseStepDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseStepDto) SetName(v string)`

SetName sets Name field to given value.


### GetJobId

`func (o *BaseStepDto) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *BaseStepDto) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *BaseStepDto) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetMilestoneId

`func (o *BaseStepDto) GetMilestoneId() string`

GetMilestoneId returns the MilestoneId field if non-nil, zero value otherwise.

### GetMilestoneIdOk

`func (o *BaseStepDto) GetMilestoneIdOk() (*string, bool)`

GetMilestoneIdOk returns a tuple with the MilestoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestoneId

`func (o *BaseStepDto) SetMilestoneId(v string)`

SetMilestoneId sets MilestoneId field to given value.


### GetDependedOnBy

`func (o *BaseStepDto) GetDependedOnBy() []string`

GetDependedOnBy returns the DependedOnBy field if non-nil, zero value otherwise.

### GetDependedOnByOk

`func (o *BaseStepDto) GetDependedOnByOk() (*[]string, bool)`

GetDependedOnByOk returns a tuple with the DependedOnBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependedOnBy

`func (o *BaseStepDto) SetDependedOnBy(v []string)`

SetDependedOnBy sets DependedOnBy field to given value.


### GetDependsOn

`func (o *BaseStepDto) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *BaseStepDto) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *BaseStepDto) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


