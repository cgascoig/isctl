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

// TelemetryDruidHavingNumericFilter The simplest having clause is a numeric filter. Numeric filters can be used as the base filters for more complex boolean expressions of filters.
type TelemetryDruidHavingNumericFilter struct {
	// The having filter type.
	Type string `json:"type"`
	// aggregate metric
	Aggregation string `json:"aggregation"`
	// null
	Value                float64 `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidHavingNumericFilter TelemetryDruidHavingNumericFilter

// NewTelemetryDruidHavingNumericFilter instantiates a new TelemetryDruidHavingNumericFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidHavingNumericFilter(type_ string, aggregation string, value float64) *TelemetryDruidHavingNumericFilter {
	this := TelemetryDruidHavingNumericFilter{}
	this.Type = type_
	this.Aggregation = aggregation
	this.Value = value
	return &this
}

// NewTelemetryDruidHavingNumericFilterWithDefaults instantiates a new TelemetryDruidHavingNumericFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidHavingNumericFilterWithDefaults() *TelemetryDruidHavingNumericFilter {
	this := TelemetryDruidHavingNumericFilter{}
	return &this
}

// GetType returns the Type field value
func (o *TelemetryDruidHavingNumericFilter) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingNumericFilter) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TelemetryDruidHavingNumericFilter) SetType(v string) {
	o.Type = v
}

// GetAggregation returns the Aggregation field value
func (o *TelemetryDruidHavingNumericFilter) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingNumericFilter) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *TelemetryDruidHavingNumericFilter) SetAggregation(v string) {
	o.Aggregation = v
}

// GetValue returns the Value field value
func (o *TelemetryDruidHavingNumericFilter) GetValue() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidHavingNumericFilter) GetValueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TelemetryDruidHavingNumericFilter) SetValue(v float64) {
	o.Value = v
}

func (o TelemetryDruidHavingNumericFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["aggregation"] = o.Aggregation
	}
	if true {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidHavingNumericFilter) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidHavingNumericFilter := _TelemetryDruidHavingNumericFilter{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidHavingNumericFilter); err == nil {
		*o = TelemetryDruidHavingNumericFilter(varTelemetryDruidHavingNumericFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "aggregation")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidHavingNumericFilter struct {
	value *TelemetryDruidHavingNumericFilter
	isSet bool
}

func (v NullableTelemetryDruidHavingNumericFilter) Get() *TelemetryDruidHavingNumericFilter {
	return v.value
}

func (v *NullableTelemetryDruidHavingNumericFilter) Set(val *TelemetryDruidHavingNumericFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidHavingNumericFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidHavingNumericFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidHavingNumericFilter(val *TelemetryDruidHavingNumericFilter) *NullableTelemetryDruidHavingNumericFilter {
	return &NullableTelemetryDruidHavingNumericFilter{value: val, isSet: true}
}

func (v NullableTelemetryDruidHavingNumericFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidHavingNumericFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
