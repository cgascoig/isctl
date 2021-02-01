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
)

// KubernetesServiceStatus The Status field of the Kubernetes Service Object.
type KubernetesServiceStatus struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType   string                         `json:"ObjectType" yaml:"ObjectType"`
	LoadBalancer NullableKubernetesLoadBalancer `json:"LoadBalancer,omitempty" yaml:"LoadBalancer,omitempty"`
}

// NewKubernetesServiceStatus instantiates a new KubernetesServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesServiceStatus(classId string, objectType string) *KubernetesServiceStatus {
	this := KubernetesServiceStatus{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesServiceStatusWithDefaults instantiates a new KubernetesServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesServiceStatusWithDefaults() *KubernetesServiceStatus {
	this := KubernetesServiceStatus{}
	var classId string = "kubernetes.ServiceStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.ServiceStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesServiceStatus) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceStatus) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesServiceStatus) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesServiceStatus) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesServiceStatus) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesServiceStatus) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLoadBalancer returns the LoadBalancer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesServiceStatus) GetLoadBalancer() KubernetesLoadBalancer {
	if o == nil || o.LoadBalancer.Get() == nil {
		var ret KubernetesLoadBalancer
		return ret
	}
	return *o.LoadBalancer.Get()
}

// GetLoadBalancerOk returns a tuple with the LoadBalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesServiceStatus) GetLoadBalancerOk() (*KubernetesLoadBalancer, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoadBalancer.Get(), o.LoadBalancer.IsSet()
}

// HasLoadBalancer returns a boolean if a field has been set.
func (o *KubernetesServiceStatus) HasLoadBalancer() bool {
	if o != nil && o.LoadBalancer.IsSet() {
		return true
	}

	return false
}

// SetLoadBalancer gets a reference to the given NullableKubernetesLoadBalancer and assigns it to the LoadBalancer field.
func (o *KubernetesServiceStatus) SetLoadBalancer(v KubernetesLoadBalancer) {
	o.LoadBalancer.Set(&v)
}

// SetLoadBalancerNil sets the value for LoadBalancer to be an explicit nil
func (o *KubernetesServiceStatus) SetLoadBalancerNil() {
	o.LoadBalancer.Set(nil)
}

// UnsetLoadBalancer ensures that no value is present for LoadBalancer, not even an explicit nil
func (o *KubernetesServiceStatus) UnsetLoadBalancer() {
	o.LoadBalancer.Unset()
}

func (o KubernetesServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LoadBalancer.IsSet() {
		toSerialize["LoadBalancer"] = o.LoadBalancer.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesServiceStatus struct {
	value *KubernetesServiceStatus
	isSet bool
}

func (v NullableKubernetesServiceStatus) Get() *KubernetesServiceStatus {
	return v.value
}

func (v *NullableKubernetesServiceStatus) Set(val *KubernetesServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesServiceStatus(val *KubernetesServiceStatus) *NullableKubernetesServiceStatus {
	return &NullableKubernetesServiceStatus{value: val, isSet: true}
}

func (v NullableKubernetesServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
