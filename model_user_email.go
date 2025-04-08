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

// checks if the UserEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEmail{}

// UserEmail A Kowabunga user email.
type UserEmail struct {
	// The user email address used for login.
	Email string `json:"email"`
}

type _UserEmail UserEmail

// NewUserEmail instantiates a new UserEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEmail(email string) *UserEmail {
	this := UserEmail{}
	this.Email = email
	return &this
}

// NewUserEmailWithDefaults instantiates a new UserEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEmailWithDefaults() *UserEmail {
	this := UserEmail{}
	return &this
}

// GetEmail returns the Email field value
func (o *UserEmail) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UserEmail) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UserEmail) SetEmail(v string) {
	o.Email = v
}

func (o UserEmail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *UserEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varUserEmail := _UserEmail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserEmail)

	if err != nil {
		return err
	}

	*o = UserEmail(varUserEmail)

	return err
}

type NullableUserEmail struct {
	value *UserEmail
	isSet bool
}

func (v NullableUserEmail) Get() *UserEmail {
	return v.value
}

func (v *NullableUserEmail) Set(val *UserEmail) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEmail) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEmail(val *UserEmail) *NullableUserEmail {
	return &NullableUserEmail{value: val, isSet: true}
}

func (v NullableUserEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


