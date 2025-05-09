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

// checks if the Kompute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Kompute{}

// Kompute A Kompute is a wrapper object for bare virtual machines. It consists of an instance, one to several attached volumes and 2 network adapters (a private one, a public one). This is the prefered way for creating virtual machines. IP addresses will be automatically assigned.
type Kompute struct {
	// The Kompute ID (auto-generated).
	Id *string `json:"id,omitempty"`
	// The Kompute name.
	Name string `json:"name"`
	// The Kompute description.
	Description *string `json:"description,omitempty"`
	// The Kompute memory size (in bytes).
	Memory int64 `json:"memory"`
	// The Kompute number of vCPUs.
	Vcpus int64 `json:"vcpus"`
	// The Kompute OS disk size (in bytes).
	Disk int64 `json:"disk"`
	// The Kompute extra data disk size (in bytes). If unspecified, no extra data disk will be assigned.
	DataDisk *int64 `json:"data_disk,omitempty"`
	// The Kompute assigned private IPv4 address (read-only).
	Ip *string `json:"ip,omitempty"`
}

type _Kompute Kompute

// NewKompute instantiates a new Kompute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKompute(name string, memory int64, vcpus int64, disk int64) *Kompute {
	this := Kompute{}
	this.Name = name
	this.Memory = memory
	this.Vcpus = vcpus
	this.Disk = disk
	var dataDisk int64 = 0
	this.DataDisk = &dataDisk
	return &this
}

// NewKomputeWithDefaults instantiates a new Kompute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKomputeWithDefaults() *Kompute {
	this := Kompute{}
	var dataDisk int64 = 0
	this.DataDisk = &dataDisk
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Kompute) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kompute) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Kompute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Kompute) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Kompute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Kompute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Kompute) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Kompute) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kompute) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Kompute) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Kompute) SetDescription(v string) {
	o.Description = &v
}

// GetMemory returns the Memory field value
func (o *Kompute) GetMemory() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Kompute) GetMemoryOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memory, true
}

// SetMemory sets field value
func (o *Kompute) SetMemory(v int64) {
	o.Memory = v
}

// GetVcpus returns the Vcpus field value
func (o *Kompute) GetVcpus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value
// and a boolean to check if the value has been set.
func (o *Kompute) GetVcpusOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vcpus, true
}

// SetVcpus sets field value
func (o *Kompute) SetVcpus(v int64) {
	o.Vcpus = v
}

// GetDisk returns the Disk field value
func (o *Kompute) GetDisk() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Disk
}

// GetDiskOk returns a tuple with the Disk field value
// and a boolean to check if the value has been set.
func (o *Kompute) GetDiskOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disk, true
}

// SetDisk sets field value
func (o *Kompute) SetDisk(v int64) {
	o.Disk = v
}

// GetDataDisk returns the DataDisk field value if set, zero value otherwise.
func (o *Kompute) GetDataDisk() int64 {
	if o == nil || IsNil(o.DataDisk) {
		var ret int64
		return ret
	}
	return *o.DataDisk
}

// GetDataDiskOk returns a tuple with the DataDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kompute) GetDataDiskOk() (*int64, bool) {
	if o == nil || IsNil(o.DataDisk) {
		return nil, false
	}
	return o.DataDisk, true
}

// HasDataDisk returns a boolean if a field has been set.
func (o *Kompute) HasDataDisk() bool {
	if o != nil && !IsNil(o.DataDisk) {
		return true
	}

	return false
}

// SetDataDisk gets a reference to the given int64 and assigns it to the DataDisk field.
func (o *Kompute) SetDataDisk(v int64) {
	o.DataDisk = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Kompute) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Kompute) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *Kompute) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Kompute) SetIp(v string) {
	o.Ip = &v
}

func (o Kompute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Kompute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["memory"] = o.Memory
	toSerialize["vcpus"] = o.Vcpus
	toSerialize["disk"] = o.Disk
	if !IsNil(o.DataDisk) {
		toSerialize["data_disk"] = o.DataDisk
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return toSerialize, nil
}

func (o *Kompute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"memory",
		"vcpus",
		"disk",
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

	varKompute := _Kompute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKompute)

	if err != nil {
		return err
	}

	*o = Kompute(varKompute)

	return err
}

type NullableKompute struct {
	value *Kompute
	isSet bool
}

func (v NullableKompute) Get() *Kompute {
	return v.value
}

func (v *NullableKompute) Set(val *Kompute) {
	v.value = val
	v.isSet = true
}

func (v NullableKompute) IsSet() bool {
	return v.isSet
}

func (v *NullableKompute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKompute(val *Kompute) *NullableKompute {
	return &NullableKompute{value: val, isSet: true}
}

func (v NullableKompute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKompute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


