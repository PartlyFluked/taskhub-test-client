/*
testing-step

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TestingStepV1CompletedStepDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestingStepV1CompletedStepDto{}

// TestingStepV1CompletedStepDto struct for TestingStepV1CompletedStepDto
type TestingStepV1CompletedStepDto struct {
	Id string `json:"id"`
	Key string `json:"key"`
	Name string `json:"name"`
	JobId string `json:"jobId"`
	MilestoneId string `json:"milestoneId"`
	DependedOnBy []string `json:"dependedOnBy"`
	DependsOn []string `json:"dependsOn"`
	StepType TestingStepV1StepType `json:"stepType"`
	Parameters map[string]string `json:"parameters"`
	Status TestingStepV1StepStatus `json:"status"`
	Result map[string]interface{} `json:"result"`
}

type _TestingStepV1CompletedStepDto TestingStepV1CompletedStepDto

// NewTestingStepV1CompletedStepDto instantiates a new TestingStepV1CompletedStepDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestingStepV1CompletedStepDto(id string, key string, name string, jobId string, milestoneId string, dependedOnBy []string, dependsOn []string, stepType TestingStepV1StepType, parameters map[string]string, status TestingStepV1StepStatus, result map[string]interface{}) *TestingStepV1CompletedStepDto {
	this := TestingStepV1CompletedStepDto{}
	this.Id = id
	this.Key = key
	this.Name = name
	this.JobId = jobId
	this.MilestoneId = milestoneId
	this.DependedOnBy = dependedOnBy
	this.DependsOn = dependsOn
	this.StepType = stepType
	this.Parameters = parameters
	this.Status = status
	this.Result = result
	return &this
}

// NewTestingStepV1CompletedStepDtoWithDefaults instantiates a new TestingStepV1CompletedStepDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestingStepV1CompletedStepDtoWithDefaults() *TestingStepV1CompletedStepDto {
	this := TestingStepV1CompletedStepDto{}
	return &this
}

// GetId returns the Id field value
func (o *TestingStepV1CompletedStepDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TestingStepV1CompletedStepDto) SetId(v string) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *TestingStepV1CompletedStepDto) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TestingStepV1CompletedStepDto) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *TestingStepV1CompletedStepDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestingStepV1CompletedStepDto) SetName(v string) {
	o.Name = v
}

// GetJobId returns the JobId field value
func (o *TestingStepV1CompletedStepDto) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *TestingStepV1CompletedStepDto) SetJobId(v string) {
	o.JobId = v
}

// GetMilestoneId returns the MilestoneId field value
func (o *TestingStepV1CompletedStepDto) GetMilestoneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MilestoneId
}

// GetMilestoneIdOk returns a tuple with the MilestoneId field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetMilestoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MilestoneId, true
}

// SetMilestoneId sets field value
func (o *TestingStepV1CompletedStepDto) SetMilestoneId(v string) {
	o.MilestoneId = v
}

// GetDependedOnBy returns the DependedOnBy field value
func (o *TestingStepV1CompletedStepDto) GetDependedOnBy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DependedOnBy
}

// GetDependedOnByOk returns a tuple with the DependedOnBy field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetDependedOnByOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependedOnBy, true
}

// SetDependedOnBy sets field value
func (o *TestingStepV1CompletedStepDto) SetDependedOnBy(v []string) {
	o.DependedOnBy = v
}

// GetDependsOn returns the DependsOn field value
func (o *TestingStepV1CompletedStepDto) GetDependsOn() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetDependsOnOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependsOn, true
}

// SetDependsOn sets field value
func (o *TestingStepV1CompletedStepDto) SetDependsOn(v []string) {
	o.DependsOn = v
}

// GetStepType returns the StepType field value
func (o *TestingStepV1CompletedStepDto) GetStepType() TestingStepV1StepType {
	if o == nil {
		var ret TestingStepV1StepType
		return ret
	}

	return o.StepType
}

// GetStepTypeOk returns a tuple with the StepType field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetStepTypeOk() (*TestingStepV1StepType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepType, true
}

// SetStepType sets field value
func (o *TestingStepV1CompletedStepDto) SetStepType(v TestingStepV1StepType) {
	o.StepType = v
}

// GetParameters returns the Parameters field value
func (o *TestingStepV1CompletedStepDto) GetParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetParametersOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *TestingStepV1CompletedStepDto) SetParameters(v map[string]string) {
	o.Parameters = v
}

// GetStatus returns the Status field value
func (o *TestingStepV1CompletedStepDto) GetStatus() TestingStepV1StepStatus {
	if o == nil {
		var ret TestingStepV1StepStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetStatusOk() (*TestingStepV1StepStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TestingStepV1CompletedStepDto) SetStatus(v TestingStepV1StepStatus) {
	o.Status = v
}

// GetResult returns the Result field value
func (o *TestingStepV1CompletedStepDto) GetResult() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *TestingStepV1CompletedStepDto) GetResultOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *TestingStepV1CompletedStepDto) SetResult(v map[string]interface{}) {
	o.Result = v
}

func (o TestingStepV1CompletedStepDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestingStepV1CompletedStepDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["jobId"] = o.JobId
	toSerialize["milestoneId"] = o.MilestoneId
	toSerialize["dependedOnBy"] = o.DependedOnBy
	toSerialize["dependsOn"] = o.DependsOn
	toSerialize["stepType"] = o.StepType
	toSerialize["parameters"] = o.Parameters
	toSerialize["status"] = o.Status
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *TestingStepV1CompletedStepDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"key",
		"name",
		"jobId",
		"milestoneId",
		"dependedOnBy",
		"dependsOn",
		"stepType",
		"parameters",
		"status",
		"result",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTestingStepV1CompletedStepDto := _TestingStepV1CompletedStepDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTestingStepV1CompletedStepDto)

	if err != nil {
		return err
	}

	*o = TestingStepV1CompletedStepDto(varTestingStepV1CompletedStepDto)

	return err
}

type NullableTestingStepV1CompletedStepDto struct {
	value *TestingStepV1CompletedStepDto
	isSet bool
}

func (v NullableTestingStepV1CompletedStepDto) Get() *TestingStepV1CompletedStepDto {
	return v.value
}

func (v *NullableTestingStepV1CompletedStepDto) Set(val *TestingStepV1CompletedStepDto) {
	v.value = val
	v.isSet = true
}

func (v NullableTestingStepV1CompletedStepDto) IsSet() bool {
	return v.isSet
}

func (v *NullableTestingStepV1CompletedStepDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestingStepV1CompletedStepDto(val *TestingStepV1CompletedStepDto) *NullableTestingStepV1CompletedStepDto {
	return &NullableTestingStepV1CompletedStepDto{value: val, isSet: true}
}

func (v NullableTestingStepV1CompletedStepDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestingStepV1CompletedStepDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

