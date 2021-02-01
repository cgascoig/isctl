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

// PolicyinventoryJobInfoAllOf Definition of the list of properties defined in 'policyinventory.JobInfo', excluding properties defined in parent classes.
type PolicyinventoryJobInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Execution status of the inventory job. * `Scheduled` - Inventory job is marked as scheduled. * `Completed` - Inventory job is marked as completed. * `Error` - Inventory job has errored out.
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" yaml:"ExecutionStatus,omitempty"`
	// Last scheduled time of the inventory job.
	LastScheduledTime *time.Time `json:"LastScheduledTime,omitempty" yaml:"LastScheduledTime,omitempty"`
	// Policy ID for the inventory job.
	PolicyId *string `json:"PolicyId,omitempty" yaml:"PolicyId,omitempty"`
	// Policy name for the inventory job.
	PolicyName *string `json:"PolicyName,omitempty" yaml:"PolicyName,omitempty"`
}

// NewPolicyinventoryJobInfoAllOf instantiates a new PolicyinventoryJobInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyinventoryJobInfoAllOf(classId string, objectType string) *PolicyinventoryJobInfoAllOf {
	this := PolicyinventoryJobInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var executionStatus string = "Scheduled"
	this.ExecutionStatus = &executionStatus
	return &this
}

// NewPolicyinventoryJobInfoAllOfWithDefaults instantiates a new PolicyinventoryJobInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyinventoryJobInfoAllOfWithDefaults() *PolicyinventoryJobInfoAllOf {
	this := PolicyinventoryJobInfoAllOf{}
	var classId string = "policyinventory.JobInfo"
	this.ClassId = classId
	var objectType string = "policyinventory.JobInfo"
	this.ObjectType = objectType
	var executionStatus string = "Scheduled"
	this.ExecutionStatus = &executionStatus
	return &this
}

// GetClassId returns the ClassId field value
func (o *PolicyinventoryJobInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *PolicyinventoryJobInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *PolicyinventoryJobInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *PolicyinventoryJobInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PolicyinventoryJobInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PolicyinventoryJobInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetExecutionStatus returns the ExecutionStatus field value if set, zero value otherwise.
func (o *PolicyinventoryJobInfoAllOf) GetExecutionStatus() string {
	if o == nil || o.ExecutionStatus == nil {
		var ret string
		return ret
	}
	return *o.ExecutionStatus
}

// GetExecutionStatusOk returns a tuple with the ExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryJobInfoAllOf) GetExecutionStatusOk() (*string, bool) {
	if o == nil || o.ExecutionStatus == nil {
		return nil, false
	}
	return o.ExecutionStatus, true
}

// HasExecutionStatus returns a boolean if a field has been set.
func (o *PolicyinventoryJobInfoAllOf) HasExecutionStatus() bool {
	if o != nil && o.ExecutionStatus != nil {
		return true
	}

	return false
}

// SetExecutionStatus gets a reference to the given string and assigns it to the ExecutionStatus field.
func (o *PolicyinventoryJobInfoAllOf) SetExecutionStatus(v string) {
	o.ExecutionStatus = &v
}

// GetLastScheduledTime returns the LastScheduledTime field value if set, zero value otherwise.
func (o *PolicyinventoryJobInfoAllOf) GetLastScheduledTime() time.Time {
	if o == nil || o.LastScheduledTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastScheduledTime
}

// GetLastScheduledTimeOk returns a tuple with the LastScheduledTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryJobInfoAllOf) GetLastScheduledTimeOk() (*time.Time, bool) {
	if o == nil || o.LastScheduledTime == nil {
		return nil, false
	}
	return o.LastScheduledTime, true
}

// HasLastScheduledTime returns a boolean if a field has been set.
func (o *PolicyinventoryJobInfoAllOf) HasLastScheduledTime() bool {
	if o != nil && o.LastScheduledTime != nil {
		return true
	}

	return false
}

// SetLastScheduledTime gets a reference to the given time.Time and assigns it to the LastScheduledTime field.
func (o *PolicyinventoryJobInfoAllOf) SetLastScheduledTime(v time.Time) {
	o.LastScheduledTime = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PolicyinventoryJobInfoAllOf) GetPolicyId() string {
	if o == nil || o.PolicyId == nil {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryJobInfoAllOf) GetPolicyIdOk() (*string, bool) {
	if o == nil || o.PolicyId == nil {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PolicyinventoryJobInfoAllOf) HasPolicyId() bool {
	if o != nil && o.PolicyId != nil {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *PolicyinventoryJobInfoAllOf) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *PolicyinventoryJobInfoAllOf) GetPolicyName() string {
	if o == nil || o.PolicyName == nil {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyinventoryJobInfoAllOf) GetPolicyNameOk() (*string, bool) {
	if o == nil || o.PolicyName == nil {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *PolicyinventoryJobInfoAllOf) HasPolicyName() bool {
	if o != nil && o.PolicyName != nil {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *PolicyinventoryJobInfoAllOf) SetPolicyName(v string) {
	o.PolicyName = &v
}

func (o PolicyinventoryJobInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ExecutionStatus != nil {
		toSerialize["ExecutionStatus"] = o.ExecutionStatus
	}
	if o.LastScheduledTime != nil {
		toSerialize["LastScheduledTime"] = o.LastScheduledTime
	}
	if o.PolicyId != nil {
		toSerialize["PolicyId"] = o.PolicyId
	}
	if o.PolicyName != nil {
		toSerialize["PolicyName"] = o.PolicyName
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyinventoryJobInfoAllOf struct {
	value *PolicyinventoryJobInfoAllOf
	isSet bool
}

func (v NullablePolicyinventoryJobInfoAllOf) Get() *PolicyinventoryJobInfoAllOf {
	return v.value
}

func (v *NullablePolicyinventoryJobInfoAllOf) Set(val *PolicyinventoryJobInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyinventoryJobInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyinventoryJobInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyinventoryJobInfoAllOf(val *PolicyinventoryJobInfoAllOf) *NullablePolicyinventoryJobInfoAllOf {
	return &NullablePolicyinventoryJobInfoAllOf{value: val, isSet: true}
}

func (v NullablePolicyinventoryJobInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyinventoryJobInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
