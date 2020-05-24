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
	"time"
)

// OnpremUpgradePhaseAllOf Definition of the list of properties defined in 'onprem.UpgradePhase', excluding properties defined in parent classes.
type OnpremUpgradePhaseAllOf struct {
	// Elapsed time in seconds to complete the current upgrade phase.
	ElapsedTime *int64 `json:"ElapsedTime,omitempty" yaml:"ElapsedTime,omitempty"`
	// End date of the software upgrade phase.
	EndTime *time.Time `json:"EndTime,omitempty" yaml:"EndTime,omitempty"`
	// Indicates if the upgrade phase has failed or not.
	Failed *bool `json:"Failed,omitempty" yaml:"Failed,omitempty"`
	// Status message set during the upgrade phase.
	Message *string `json:"Message,omitempty" yaml:"Message,omitempty"`
	// Name of the upgrade phase.
	Name *string `json:"Name,omitempty" yaml:"Name,omitempty"`
	// Start date of the software upgrade phase.
	StartTime *time.Time `json:"StartTime,omitempty" yaml:"StartTime,omitempty"`
}

// NewOnpremUpgradePhaseAllOf instantiates a new OnpremUpgradePhaseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnpremUpgradePhaseAllOf() *OnpremUpgradePhaseAllOf {
	this := OnpremUpgradePhaseAllOf{}
	var name string = "init"
	this.Name = &name
	return &this
}

// NewOnpremUpgradePhaseAllOfWithDefaults instantiates a new OnpremUpgradePhaseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnpremUpgradePhaseAllOfWithDefaults() *OnpremUpgradePhaseAllOf {
	this := OnpremUpgradePhaseAllOf{}
	var name string = "init"
	this.Name = &name
	return &this
}

// GetElapsedTime returns the ElapsedTime field value if set, zero value otherwise.
func (o *OnpremUpgradePhaseAllOf) GetElapsedTime() int64 {
	if o == nil || o.ElapsedTime == nil {
		var ret int64
		return ret
	}
	return *o.ElapsedTime
}

// GetElapsedTimeOk returns a tuple with the ElapsedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhaseAllOf) GetElapsedTimeOk() (*int64, bool) {
	if o == nil || o.ElapsedTime == nil {
		return nil, false
	}
	return o.ElapsedTime, true
}

// HasElapsedTime returns a boolean if a field has been set.
func (o *OnpremUpgradePhaseAllOf) HasElapsedTime() bool {
	if o != nil && o.ElapsedTime != nil {
		return true
	}

	return false
}

// SetElapsedTime gets a reference to the given int64 and assigns it to the ElapsedTime field.
func (o *OnpremUpgradePhaseAllOf) SetElapsedTime(v int64) {
	o.ElapsedTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *OnpremUpgradePhaseAllOf) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhaseAllOf) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *OnpremUpgradePhaseAllOf) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *OnpremUpgradePhaseAllOf) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *OnpremUpgradePhaseAllOf) GetFailed() bool {
	if o == nil || o.Failed == nil {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhaseAllOf) GetFailedOk() (*bool, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *OnpremUpgradePhaseAllOf) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *OnpremUpgradePhaseAllOf) SetFailed(v bool) {
	o.Failed = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OnpremUpgradePhaseAllOf) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhaseAllOf) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OnpremUpgradePhaseAllOf) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OnpremUpgradePhaseAllOf) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OnpremUpgradePhaseAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhaseAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OnpremUpgradePhaseAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OnpremUpgradePhaseAllOf) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *OnpremUpgradePhaseAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnpremUpgradePhaseAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *OnpremUpgradePhaseAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *OnpremUpgradePhaseAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

func (o OnpremUpgradePhaseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ElapsedTime != nil {
		toSerialize["ElapsedTime"] = o.ElapsedTime
	}
	if o.EndTime != nil {
		toSerialize["EndTime"] = o.EndTime
	}
	if o.Failed != nil {
		toSerialize["Failed"] = o.Failed
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableOnpremUpgradePhaseAllOf struct {
	value *OnpremUpgradePhaseAllOf
	isSet bool
}

func (v NullableOnpremUpgradePhaseAllOf) Get() *OnpremUpgradePhaseAllOf {
	return v.value
}

func (v *NullableOnpremUpgradePhaseAllOf) Set(val *OnpremUpgradePhaseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOnpremUpgradePhaseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOnpremUpgradePhaseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnpremUpgradePhaseAllOf(val *OnpremUpgradePhaseAllOf) *NullableOnpremUpgradePhaseAllOf {
	return &NullableOnpremUpgradePhaseAllOf{value: val, isSet: true}
}

func (v NullableOnpremUpgradePhaseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnpremUpgradePhaseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
