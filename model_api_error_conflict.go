/*
Kowabunga API documentation

Kvm Orchestrator With A BUNch of Goods Added

API version: 0.52.5
Contact: maintainers@kowabunga.cloud
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kowabunga

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiErrorConflict type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorConflict{}

// ApiErrorConflict struct for ApiErrorConflict
type ApiErrorConflict struct {
	Status int32 `json:"status"`
	Error string `json:"error"`
	Reason string `json:"reason"`
}

type _ApiErrorConflict ApiErrorConflict

// NewApiErrorConflict instantiates a new ApiErrorConflict object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorConflict(status int32, error_ string, reason string) *ApiErrorConflict {
	this := ApiErrorConflict{}
	this.Status = status
	this.Error = error_
	this.Reason = reason
	return &this
}

// NewApiErrorConflictWithDefaults instantiates a new ApiErrorConflict object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorConflictWithDefaults() *ApiErrorConflict {
	this := ApiErrorConflict{}
	return &this
}

// GetStatus returns the Status field value
func (o *ApiErrorConflict) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiErrorConflict) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiErrorConflict) SetStatus(v int32) {
	o.Status = v
}

// GetError returns the Error field value
func (o *ApiErrorConflict) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ApiErrorConflict) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ApiErrorConflict) SetError(v string) {
	o.Error = v
}

// GetReason returns the Reason field value
func (o *ApiErrorConflict) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ApiErrorConflict) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ApiErrorConflict) SetReason(v string) {
	o.Reason = v
}

func (o ApiErrorConflict) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorConflict) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["error"] = o.Error
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *ApiErrorConflict) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"error",
		"reason",
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

	varApiErrorConflict := _ApiErrorConflict{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiErrorConflict)

	if err != nil {
		return err
	}

	*o = ApiErrorConflict(varApiErrorConflict)

	return err
}

type NullableApiErrorConflict struct {
	value *ApiErrorConflict
	isSet bool
}

func (v NullableApiErrorConflict) Get() *ApiErrorConflict {
	return v.value
}

func (v *NullableApiErrorConflict) Set(val *ApiErrorConflict) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorConflict) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorConflict) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorConflict(val *ApiErrorConflict) *NullableApiErrorConflict {
	return &NullableApiErrorConflict{value: val, isSet: true}
}

func (v NullableApiErrorConflict) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorConflict) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


