/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-01-11T18:30:19Z.
 *
 * API version: 1.0.9-3252
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
	"time"
)

// MoVersionContextAllOf Definition of the list of properties defined in 'mo.VersionContext', excluding properties defined in parent classes.
type MoVersionContextAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType    string    `json:"ObjectType" yaml:"ObjectType"`
	InterestedMos []MoMoRef `json:"InterestedMos,omitempty" yaml:"InterestedMos,omitempty"`
	RefMo         *MoMoRef  `json:"RefMo,omitempty" yaml:"RefMo,omitempty"`
	// The time this versioned Managed Object was created.
	Timestamp *time.Time `json:"Timestamp,omitempty" yaml:"Timestamp,omitempty"`
	// The version of the Managed Object, e.g. an incrementing number or a hash id.
	Version *string `json:"Version,omitempty" yaml:"Version,omitempty"`
	// Specifies type of version. Currently the only supported value is \"Configured\" that is used to keep track of snapshots of policies and profiles that are intended to be configured to target endpoints. * `Modified` - Version created every time an object is modified. * `Configured` - Version created every time an object is configured to the service profile. * `Deployed` - Version created for objects related to a service profile when it is deployed.
	VersionType *string `json:"VersionType,omitempty" yaml:"VersionType,omitempty"`
}

// NewMoVersionContextAllOf instantiates a new MoVersionContextAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoVersionContextAllOf(classId string, objectType string) *MoVersionContextAllOf {
	this := MoVersionContextAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var versionType string = "Modified"
	this.VersionType = &versionType
	return &this
}

// NewMoVersionContextAllOfWithDefaults instantiates a new MoVersionContextAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoVersionContextAllOfWithDefaults() *MoVersionContextAllOf {
	this := MoVersionContextAllOf{}
	var classId string = "mo.VersionContext"
	this.ClassId = classId
	var objectType string = "mo.VersionContext"
	this.ObjectType = objectType
	var versionType string = "Modified"
	this.VersionType = &versionType
	return &this
}

// GetClassId returns the ClassId field value
func (o *MoVersionContextAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *MoVersionContextAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *MoVersionContextAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *MoVersionContextAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetInterestedMos returns the InterestedMos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MoVersionContextAllOf) GetInterestedMos() []MoMoRef {
	if o == nil {
		var ret []MoMoRef
		return ret
	}
	return o.InterestedMos
}

// GetInterestedMosOk returns a tuple with the InterestedMos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MoVersionContextAllOf) GetInterestedMosOk() (*[]MoMoRef, bool) {
	if o == nil || o.InterestedMos == nil {
		return nil, false
	}
	return &o.InterestedMos, true
}

// HasInterestedMos returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasInterestedMos() bool {
	if o != nil && o.InterestedMos != nil {
		return true
	}

	return false
}

// SetInterestedMos gets a reference to the given []MoMoRef and assigns it to the InterestedMos field.
func (o *MoVersionContextAllOf) SetInterestedMos(v []MoMoRef) {
	o.InterestedMos = v
}

// GetRefMo returns the RefMo field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetRefMo() MoMoRef {
	if o == nil || o.RefMo == nil {
		var ret MoMoRef
		return ret
	}
	return *o.RefMo
}

// GetRefMoOk returns a tuple with the RefMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetRefMoOk() (*MoMoRef, bool) {
	if o == nil || o.RefMo == nil {
		return nil, false
	}
	return o.RefMo, true
}

// HasRefMo returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasRefMo() bool {
	if o != nil && o.RefMo != nil {
		return true
	}

	return false
}

// SetRefMo gets a reference to the given MoMoRef and assigns it to the RefMo field.
func (o *MoVersionContextAllOf) SetRefMo(v MoMoRef) {
	o.RefMo = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *MoVersionContextAllOf) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MoVersionContextAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *MoVersionContextAllOf) GetVersionType() string {
	if o == nil || o.VersionType == nil {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoVersionContextAllOf) GetVersionTypeOk() (*string, bool) {
	if o == nil || o.VersionType == nil {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *MoVersionContextAllOf) HasVersionType() bool {
	if o != nil && o.VersionType != nil {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *MoVersionContextAllOf) SetVersionType(v string) {
	o.VersionType = &v
}

func (o MoVersionContextAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.InterestedMos != nil {
		toSerialize["InterestedMos"] = o.InterestedMos
	}
	if o.RefMo != nil {
		toSerialize["RefMo"] = o.RefMo
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.VersionType != nil {
		toSerialize["VersionType"] = o.VersionType
	}
	return json.Marshal(toSerialize)
}

type NullableMoVersionContextAllOf struct {
	value *MoVersionContextAllOf
	isSet bool
}

func (v NullableMoVersionContextAllOf) Get() *MoVersionContextAllOf {
	return v.value
}

func (v *NullableMoVersionContextAllOf) Set(val *MoVersionContextAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMoVersionContextAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMoVersionContextAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoVersionContextAllOf(val *MoVersionContextAllOf) *NullableMoVersionContextAllOf {
	return &NullableMoVersionContextAllOf{value: val, isSet: true}
}

func (v NullableMoVersionContextAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoVersionContextAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
