/*
testing-step

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TestingStepV1StepType the model 'TestingStepV1StepType'
type TestingStepV1StepType string

// List of testing-step-v1_StepType
const (
	TESTING_STEP_V1 TestingStepV1StepType = "testing-step-v1"
)

// All allowed values of TestingStepV1StepType enum
var AllowedTestingStepV1StepTypeEnumValues = []TestingStepV1StepType{
	"testing-step-v1",
}

func (v *TestingStepV1StepType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestingStepV1StepType(value)
	for _, existing := range AllowedTestingStepV1StepTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestingStepV1StepType", value)
}

// NewTestingStepV1StepTypeFromValue returns a pointer to a valid TestingStepV1StepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestingStepV1StepTypeFromValue(v string) (*TestingStepV1StepType, error) {
	ev := TestingStepV1StepType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestingStepV1StepType: valid values are %v", v, AllowedTestingStepV1StepTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestingStepV1StepType) IsValid() bool {
	for _, existing := range AllowedTestingStepV1StepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to testing-step-v1_StepType value
func (v TestingStepV1StepType) Ptr() *TestingStepV1StepType {
	return &v
}

type NullableTestingStepV1StepType struct {
	value *TestingStepV1StepType
	isSet bool
}

func (v NullableTestingStepV1StepType) Get() *TestingStepV1StepType {
	return v.value
}

func (v *NullableTestingStepV1StepType) Set(val *TestingStepV1StepType) {
	v.value = val
	v.isSet = true
}

func (v NullableTestingStepV1StepType) IsSet() bool {
	return v.isSet
}

func (v *NullableTestingStepV1StepType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestingStepV1StepType(val *TestingStepV1StepType) *NullableTestingStepV1StepType {
	return &NullableTestingStepV1StepType{value: val, isSet: true}
}

func (v NullableTestingStepV1StepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestingStepV1StepType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
