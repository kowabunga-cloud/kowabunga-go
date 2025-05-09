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
)

// checks if the RegionSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionSubnet{}

// RegionSubnet A region/subnet map.
type RegionSubnet struct {
	// The region key.
	Key *string `json:"key,omitempty"`
	// The subnet ID.
	Value *string `json:"value,omitempty"`
}

// NewRegionSubnet instantiates a new RegionSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionSubnet() *RegionSubnet {
	this := RegionSubnet{}
	return &this
}

// NewRegionSubnetWithDefaults instantiates a new RegionSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionSubnetWithDefaults() *RegionSubnet {
	this := RegionSubnet{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *RegionSubnet) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSubnet) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *RegionSubnet) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *RegionSubnet) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RegionSubnet) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionSubnet) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RegionSubnet) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RegionSubnet) SetValue(v string) {
	o.Value = &v
}

func (o RegionSubnet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRegionSubnet struct {
	value *RegionSubnet
	isSet bool
}

func (v NullableRegionSubnet) Get() *RegionSubnet {
	return v.value
}

func (v *NullableRegionSubnet) Set(val *RegionSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionSubnet(val *RegionSubnet) *NullableRegionSubnet {
	return &NullableRegionSubnet{value: val, isSet: true}
}

func (v NullableRegionSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


