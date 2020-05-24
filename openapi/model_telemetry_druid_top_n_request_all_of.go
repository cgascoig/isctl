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

// TelemetryDruidTopNRequestAllOf struct for TelemetryDruidTopNRequestAllOf
type TelemetryDruidTopNRequestAllOf struct {
	DataSource TelemetryDruidDataSource `json:"dataSource" yaml:"dataSource"`
	// A JSON Object representing ISO-8601 Intervals. This defines the time ranges to run the query over.
	Intervals        []string                      `json:"intervals" yaml:"intervals"`
	Granularity      TelemetryDruidGranularity     `json:"granularity" yaml:"granularity"`
	Filter           *TelemetryDruidFilter         `json:"filter,omitempty" yaml:"filter,omitempty"`
	Aggregations     *TelemetryDruidAggregator     `json:"aggregations,omitempty" yaml:"aggregations,omitempty"`
	PostAggregations *TelemetryDruidPostAggregator `json:"postAggregations,omitempty" yaml:"postAggregations,omitempty"`
	Dimension        TelemetryDruidDimensionSpec   `json:"dimension" yaml:"dimension"`
	// An integer defining the N in the topN (i.e. how many results you want in the top list).
	Threshold int32                        `json:"threshold" yaml:"threshold"`
	Metric    TelemetryDruidTopNMetricSpec `json:"metric" yaml:"metric"`
	Context   *TelemetryDruidQueryContext  `json:"context,omitempty" yaml:"context,omitempty"`
}

// NewTelemetryDruidTopNRequestAllOf instantiates a new TelemetryDruidTopNRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidTopNRequestAllOf(dataSource TelemetryDruidDataSource, intervals []string, granularity TelemetryDruidGranularity, dimension TelemetryDruidDimensionSpec, threshold int32, metric TelemetryDruidTopNMetricSpec) *TelemetryDruidTopNRequestAllOf {
	this := TelemetryDruidTopNRequestAllOf{}
	this.DataSource = dataSource
	this.Intervals = intervals
	this.Granularity = granularity
	this.Dimension = dimension
	this.Threshold = threshold
	this.Metric = metric
	return &this
}

// NewTelemetryDruidTopNRequestAllOfWithDefaults instantiates a new TelemetryDruidTopNRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidTopNRequestAllOfWithDefaults() *TelemetryDruidTopNRequestAllOf {
	this := TelemetryDruidTopNRequestAllOf{}
	return &this
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidTopNRequestAllOf) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidTopNRequestAllOf) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetIntervals returns the Intervals field value
func (o *TelemetryDruidTopNRequestAllOf) GetIntervals() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetIntervalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Intervals, true
}

// SetIntervals sets field value
func (o *TelemetryDruidTopNRequestAllOf) SetIntervals(v []string) {
	o.Intervals = v
}

// GetGranularity returns the Granularity field value
func (o *TelemetryDruidTopNRequestAllOf) GetGranularity() TelemetryDruidGranularity {
	if o == nil {
		var ret TelemetryDruidGranularity
		return ret
	}

	return o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetGranularityOk() (*TelemetryDruidGranularity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Granularity, true
}

// SetGranularity sets field value
func (o *TelemetryDruidTopNRequestAllOf) SetGranularity(v TelemetryDruidGranularity) {
	o.Granularity = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequestAllOf) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequestAllOf) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidTopNRequestAllOf) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetAggregations returns the Aggregations field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequestAllOf) GetAggregations() TelemetryDruidAggregator {
	if o == nil || o.Aggregations == nil {
		var ret TelemetryDruidAggregator
		return ret
	}
	return *o.Aggregations
}

// GetAggregationsOk returns a tuple with the Aggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetAggregationsOk() (*TelemetryDruidAggregator, bool) {
	if o == nil || o.Aggregations == nil {
		return nil, false
	}
	return o.Aggregations, true
}

// HasAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequestAllOf) HasAggregations() bool {
	if o != nil && o.Aggregations != nil {
		return true
	}

	return false
}

// SetAggregations gets a reference to the given TelemetryDruidAggregator and assigns it to the Aggregations field.
func (o *TelemetryDruidTopNRequestAllOf) SetAggregations(v TelemetryDruidAggregator) {
	o.Aggregations = &v
}

// GetPostAggregations returns the PostAggregations field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequestAllOf) GetPostAggregations() TelemetryDruidPostAggregator {
	if o == nil || o.PostAggregations == nil {
		var ret TelemetryDruidPostAggregator
		return ret
	}
	return *o.PostAggregations
}

// GetPostAggregationsOk returns a tuple with the PostAggregations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetPostAggregationsOk() (*TelemetryDruidPostAggregator, bool) {
	if o == nil || o.PostAggregations == nil {
		return nil, false
	}
	return o.PostAggregations, true
}

// HasPostAggregations returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequestAllOf) HasPostAggregations() bool {
	if o != nil && o.PostAggregations != nil {
		return true
	}

	return false
}

// SetPostAggregations gets a reference to the given TelemetryDruidPostAggregator and assigns it to the PostAggregations field.
func (o *TelemetryDruidTopNRequestAllOf) SetPostAggregations(v TelemetryDruidPostAggregator) {
	o.PostAggregations = &v
}

// GetDimension returns the Dimension field value
func (o *TelemetryDruidTopNRequestAllOf) GetDimension() TelemetryDruidDimensionSpec {
	if o == nil {
		var ret TelemetryDruidDimensionSpec
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetDimensionOk() (*TelemetryDruidDimensionSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *TelemetryDruidTopNRequestAllOf) SetDimension(v TelemetryDruidDimensionSpec) {
	o.Dimension = v
}

// GetThreshold returns the Threshold field value
func (o *TelemetryDruidTopNRequestAllOf) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *TelemetryDruidTopNRequestAllOf) SetThreshold(v int32) {
	o.Threshold = v
}

// GetMetric returns the Metric field value
func (o *TelemetryDruidTopNRequestAllOf) GetMetric() TelemetryDruidTopNMetricSpec {
	if o == nil {
		var ret TelemetryDruidTopNMetricSpec
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetMetricOk() (*TelemetryDruidTopNMetricSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *TelemetryDruidTopNRequestAllOf) SetMetric(v TelemetryDruidTopNMetricSpec) {
	o.Metric = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidTopNRequestAllOf) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTopNRequestAllOf) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidTopNRequestAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidTopNRequestAllOf) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidTopNRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if true {
		toSerialize["intervals"] = o.Intervals
	}
	if true {
		toSerialize["granularity"] = o.Granularity
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Aggregations != nil {
		toSerialize["aggregations"] = o.Aggregations
	}
	if o.PostAggregations != nil {
		toSerialize["postAggregations"] = o.PostAggregations
	}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryDruidTopNRequestAllOf struct {
	value *TelemetryDruidTopNRequestAllOf
	isSet bool
}

func (v NullableTelemetryDruidTopNRequestAllOf) Get() *TelemetryDruidTopNRequestAllOf {
	return v.value
}

func (v *NullableTelemetryDruidTopNRequestAllOf) Set(val *TelemetryDruidTopNRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTopNRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTopNRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTopNRequestAllOf(val *TelemetryDruidTopNRequestAllOf) *NullableTelemetryDruidTopNRequestAllOf {
	return &NullableTelemetryDruidTopNRequestAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidTopNRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTopNRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
