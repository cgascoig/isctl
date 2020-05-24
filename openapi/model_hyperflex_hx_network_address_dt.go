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

// HyperflexHxNetworkAddressDt struct for HyperflexHxNetworkAddressDt
type HyperflexHxNetworkAddressDt struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	Address           *string `json:"Address,omitempty" yaml:"Address,omitempty"`
	Fqdn              *string `json:"Fqdn,omitempty" yaml:"Fqdn,omitempty"`
	Ip                *string `json:"Ip,omitempty" yaml:"Ip,omitempty"`
}

// NewHyperflexHxNetworkAddressDt instantiates a new HyperflexHxNetworkAddressDt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexHxNetworkAddressDt() *HyperflexHxNetworkAddressDt {
	this := HyperflexHxNetworkAddressDt{}
	return &this
}

// NewHyperflexHxNetworkAddressDtWithDefaults instantiates a new HyperflexHxNetworkAddressDt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexHxNetworkAddressDtWithDefaults() *HyperflexHxNetworkAddressDt {
	this := HyperflexHxNetworkAddressDt{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDt) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDt) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *HyperflexHxNetworkAddressDt) SetAddress(v string) {
	o.Address = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDt) GetFqdn() string {
	if o == nil || o.Fqdn == nil {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetFqdnOk() (*string, bool) {
	if o == nil || o.Fqdn == nil {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDt) HasFqdn() bool {
	if o != nil && o.Fqdn != nil {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *HyperflexHxNetworkAddressDt) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *HyperflexHxNetworkAddressDt) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexHxNetworkAddressDt) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *HyperflexHxNetworkAddressDt) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *HyperflexHxNetworkAddressDt) SetIp(v string) {
	o.Ip = &v
}

func (o HyperflexHxNetworkAddressDt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.Fqdn != nil {
		toSerialize["Fqdn"] = o.Fqdn
	}
	if o.Ip != nil {
		toSerialize["Ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableHyperflexHxNetworkAddressDt struct {
	value *HyperflexHxNetworkAddressDt
	isSet bool
}

func (v NullableHyperflexHxNetworkAddressDt) Get() *HyperflexHxNetworkAddressDt {
	return v.value
}

func (v *NullableHyperflexHxNetworkAddressDt) Set(val *HyperflexHxNetworkAddressDt) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexHxNetworkAddressDt) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexHxNetworkAddressDt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexHxNetworkAddressDt(val *HyperflexHxNetworkAddressDt) *NullableHyperflexHxNetworkAddressDt {
	return &NullableHyperflexHxNetworkAddressDt{value: val, isSet: true}
}

func (v NullableHyperflexHxNetworkAddressDt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexHxNetworkAddressDt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
