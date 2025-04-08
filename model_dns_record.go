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

// checks if the DnsRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRecord{}

// DnsRecord A DNS record.
type DnsRecord struct {
	// The DNS record ID (auto-generated).
	Id *string `json:"id,omitempty"`
	// The DNS record name.
	Name string `json:"name"`
	// The DNS record description.
	Description *string `json:"description,omitempty"`
	// The DNS record associated domain (inherited from associated project).
	Domain *string `json:"domain,omitempty"`
	// A list of IPv4 addresses to be associated to the record.
	Addresses []string `json:"addresses"`
}

type _DnsRecord DnsRecord

// NewDnsRecord instantiates a new DnsRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRecord(name string, addresses []string) *DnsRecord {
	this := DnsRecord{}
	this.Name = name
	this.Addresses = addresses
	return &this
}

// NewDnsRecordWithDefaults instantiates a new DnsRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRecordWithDefaults() *DnsRecord {
	this := DnsRecord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DnsRecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DnsRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DnsRecord) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DnsRecord) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DnsRecord) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DnsRecord) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DnsRecord) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecord) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DnsRecord) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DnsRecord) SetDescription(v string) {
	o.Description = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DnsRecord) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRecord) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DnsRecord) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DnsRecord) SetDomain(v string) {
	o.Domain = &v
}

// GetAddresses returns the Addresses field value
func (o *DnsRecord) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *DnsRecord) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *DnsRecord) SetAddresses(v []string) {
	o.Addresses = v
}

func (o DnsRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	toSerialize["addresses"] = o.Addresses
	return toSerialize, nil
}

func (o *DnsRecord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"addresses",
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

	varDnsRecord := _DnsRecord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDnsRecord)

	if err != nil {
		return err
	}

	*o = DnsRecord(varDnsRecord)

	return err
}

type NullableDnsRecord struct {
	value *DnsRecord
	isSet bool
}

func (v NullableDnsRecord) Get() *DnsRecord {
	return v.value
}

func (v *NullableDnsRecord) Set(val *DnsRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecord(val *DnsRecord) *NullableDnsRecord {
	return &NullableDnsRecord{value: val, isSet: true}
}

func (v NullableDnsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


