/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-04-17T15:33:06-07:00.
 *
 * API version: 1.0.9-1628
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// SoftwarerepositoryOperatingSystemFile An operating system image that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository.
type SoftwarerepositoryOperatingSystemFile struct {
	SoftwarerepositoryFile `yaml:"SoftwarerepositoryFile,inline"`
	// The vendor or publisher of this file.
	Vendor  *string                                `json:"Vendor,omitempty" yaml:"Vendor,omitempty"`
	Catalog *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty" yaml:"Catalog,omitempty"`
}

// NewSoftwarerepositoryOperatingSystemFile instantiates a new SoftwarerepositoryOperatingSystemFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryOperatingSystemFile() *SoftwarerepositoryOperatingSystemFile {
	this := SoftwarerepositoryOperatingSystemFile{}
	return &this
}

// NewSoftwarerepositoryOperatingSystemFileWithDefaults instantiates a new SoftwarerepositoryOperatingSystemFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryOperatingSystemFileWithDefaults() *SoftwarerepositoryOperatingSystemFile {
	this := SoftwarerepositoryOperatingSystemFile{}
	return &this
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *SoftwarerepositoryOperatingSystemFile) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryOperatingSystemFile) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *SoftwarerepositoryOperatingSystemFile) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *SoftwarerepositoryOperatingSystemFile) SetVendor(v string) {
	o.Vendor = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwarerepositoryOperatingSystemFile) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryOperatingSystemFile) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwarerepositoryOperatingSystemFile) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwarerepositoryOperatingSystemFile) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwarerepositoryOperatingSystemFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFile, errSoftwarerepositoryFile := json.Marshal(o.SoftwarerepositoryFile)
	if errSoftwarerepositoryFile != nil {
		return []byte{}, errSoftwarerepositoryFile
	}
	errSoftwarerepositoryFile = json.Unmarshal([]byte(serializedSoftwarerepositoryFile), &toSerialize)
	if errSoftwarerepositoryFile != nil {
		return []byte{}, errSoftwarerepositoryFile
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwarerepositoryOperatingSystemFile struct {
	value *SoftwarerepositoryOperatingSystemFile
	isSet bool
}

func (v NullableSoftwarerepositoryOperatingSystemFile) Get() *SoftwarerepositoryOperatingSystemFile {
	return v.value
}

func (v *NullableSoftwarerepositoryOperatingSystemFile) Set(val *SoftwarerepositoryOperatingSystemFile) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryOperatingSystemFile) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryOperatingSystemFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryOperatingSystemFile(val *SoftwarerepositoryOperatingSystemFile) *NullableSoftwarerepositoryOperatingSystemFile {
	return &NullableSoftwarerepositoryOperatingSystemFile{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryOperatingSystemFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryOperatingSystemFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
