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

// FirmwareHttpServer An external HTTP file server.
type FirmwareHttpServer struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.
	LocationLink *string `json:"LocationLink,omitempty" yaml:"LocationLink,omitempty"`
	// Mount option as configured on the HTTP server. Empty if nothing is configured.
	MountOptions *string `json:"MountOptions,omitempty" yaml:"MountOptions,omitempty"`
}

// NewFirmwareHttpServer instantiates a new FirmwareHttpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareHttpServer() *FirmwareHttpServer {
	this := FirmwareHttpServer{}
	return &this
}

// NewFirmwareHttpServerWithDefaults instantiates a new FirmwareHttpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareHttpServerWithDefaults() *FirmwareHttpServer {
	this := FirmwareHttpServer{}
	return &this
}

// GetLocationLink returns the LocationLink field value if set, zero value otherwise.
func (o *FirmwareHttpServer) GetLocationLink() string {
	if o == nil || o.LocationLink == nil {
		var ret string
		return ret
	}
	return *o.LocationLink
}

// GetLocationLinkOk returns a tuple with the LocationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServer) GetLocationLinkOk() (*string, bool) {
	if o == nil || o.LocationLink == nil {
		return nil, false
	}
	return o.LocationLink, true
}

// HasLocationLink returns a boolean if a field has been set.
func (o *FirmwareHttpServer) HasLocationLink() bool {
	if o != nil && o.LocationLink != nil {
		return true
	}

	return false
}

// SetLocationLink gets a reference to the given string and assigns it to the LocationLink field.
func (o *FirmwareHttpServer) SetLocationLink(v string) {
	o.LocationLink = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *FirmwareHttpServer) GetMountOptions() string {
	if o == nil || o.MountOptions == nil {
		var ret string
		return ret
	}
	return *o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServer) GetMountOptionsOk() (*string, bool) {
	if o == nil || o.MountOptions == nil {
		return nil, false
	}
	return o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *FirmwareHttpServer) HasMountOptions() bool {
	if o != nil && o.MountOptions != nil {
		return true
	}

	return false
}

// SetMountOptions gets a reference to the given string and assigns it to the MountOptions field.
func (o *FirmwareHttpServer) SetMountOptions(v string) {
	o.MountOptions = &v
}

func (o FirmwareHttpServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.LocationLink != nil {
		toSerialize["LocationLink"] = o.LocationLink
	}
	if o.MountOptions != nil {
		toSerialize["MountOptions"] = o.MountOptions
	}
	return json.Marshal(toSerialize)
}

type NullableFirmwareHttpServer struct {
	value *FirmwareHttpServer
	isSet bool
}

func (v NullableFirmwareHttpServer) Get() *FirmwareHttpServer {
	return v.value
}

func (v *NullableFirmwareHttpServer) Set(val *FirmwareHttpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareHttpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareHttpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareHttpServer(val *FirmwareHttpServer) *NullableFirmwareHttpServer {
	return &NullableFirmwareHttpServer{value: val, isSet: true}
}

func (v NullableFirmwareHttpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareHttpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
