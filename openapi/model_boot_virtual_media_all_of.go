/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BootVirtualMediaAllOf Definition of the list of properties defined in 'boot.VirtualMedia', excluding properties defined in parent classes.
type BootVirtualMediaAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The subtype for the selected device type. * `None` - No sub type for virtual media. * `cimc-mapped-dvd` - The virtual media device is mapped to a virtual DVD device. * `cimc-mapped-hdd` - The virtual media device is mapped to a virtual HDD device. * `kvm-mapped-dvd` - A KVM mapped DVD virtual media device. * `kvm-mapped-hdd` - A KVM mapped HDD virtual media device. * `kvm-mapped-fdd` - A KVM mapped FDD virtual media device.
	Subtype              *string `json:"Subtype,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BootVirtualMediaAllOf BootVirtualMediaAllOf

// NewBootVirtualMediaAllOf instantiates a new BootVirtualMediaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootVirtualMediaAllOf(classId string, objectType string) *BootVirtualMediaAllOf {
	this := BootVirtualMediaAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// NewBootVirtualMediaAllOfWithDefaults instantiates a new BootVirtualMediaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootVirtualMediaAllOfWithDefaults() *BootVirtualMediaAllOf {
	this := BootVirtualMediaAllOf{}
	var classId string = "boot.VirtualMedia"
	this.ClassId = classId
	var objectType string = "boot.VirtualMedia"
	this.ObjectType = objectType
	var subtype string = "None"
	this.Subtype = &subtype
	return &this
}

// GetClassId returns the ClassId field value
func (o *BootVirtualMediaAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *BootVirtualMediaAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *BootVirtualMediaAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *BootVirtualMediaAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *BootVirtualMediaAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *BootVirtualMediaAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *BootVirtualMediaAllOf) GetSubtype() string {
	if o == nil || o.Subtype == nil {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootVirtualMediaAllOf) GetSubtypeOk() (*string, bool) {
	if o == nil || o.Subtype == nil {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *BootVirtualMediaAllOf) HasSubtype() bool {
	if o != nil && o.Subtype != nil {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *BootVirtualMediaAllOf) SetSubtype(v string) {
	o.Subtype = &v
}

func (o BootVirtualMediaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Subtype != nil {
		toSerialize["Subtype"] = o.Subtype
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BootVirtualMediaAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBootVirtualMediaAllOf := _BootVirtualMediaAllOf{}

	if err = json.Unmarshal(bytes, &varBootVirtualMediaAllOf); err == nil {
		*o = BootVirtualMediaAllOf(varBootVirtualMediaAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Subtype")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBootVirtualMediaAllOf struct {
	value *BootVirtualMediaAllOf
	isSet bool
}

func (v NullableBootVirtualMediaAllOf) Get() *BootVirtualMediaAllOf {
	return v.value
}

func (v *NullableBootVirtualMediaAllOf) Set(val *BootVirtualMediaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBootVirtualMediaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBootVirtualMediaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootVirtualMediaAllOf(val *BootVirtualMediaAllOf) *NullableBootVirtualMediaAllOf {
	return &NullableBootVirtualMediaAllOf{value: val, isSet: true}
}

func (v NullableBootVirtualMediaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootVirtualMediaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
