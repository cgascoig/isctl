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

// WorkflowWebApi This models a single Web API request within a batch of requests that get executed within a single workflow task.
type WorkflowWebApi struct {
	WorkflowApi `yaml:"WorkflowApi,inline"`
	// If the target type is Endpoint, this property determines whether the request is to be handled as internal request or external request by the device connector.
	EndpointRequestType *string `json:"EndpointRequestType,omitempty" yaml:"EndpointRequestType,omitempty"`
	// Collection of key value pairs to set in the request header.
	Headers *map[string]interface{} `json:"Headers,omitempty" yaml:"Headers,omitempty"`
	// The HTTP method to be executed in the given URL (GET, POST, PUT, etc). If the value is not specified, GET will be used as default.
	Method *string `json:"Method,omitempty" yaml:"Method,omitempty"`
	// The type of the intersight object for which API request is to be made. The property is valid in case of Intersight API calls and the base url is automatically prepended based on the value.
	MoType *string `json:"MoType,omitempty" yaml:"MoType,omitempty"`
	// The accepted web protocol values are http and https.
	Protocol *string `json:"Protocol,omitempty" yaml:"Protocol,omitempty"`
	// If the web API is to be executed in a remote device connected to the Intersight through device connector, 'Endpoint' is expected as the value whereas if the API is an Intersight API, 'Local' is expected as the value.
	TargetType *string `json:"TargetType,omitempty" yaml:"TargetType,omitempty"`
	// The URL of the resource in the target to which the API request is made.
	Url *string `json:"Url,omitempty" yaml:"Url,omitempty"`
}

// NewWorkflowWebApi instantiates a new WorkflowWebApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWebApi() *WorkflowWebApi {
	this := WorkflowWebApi{}
	var endpointRequestType string = "Internal"
	this.EndpointRequestType = &endpointRequestType
	var targetType string = "Endpoint"
	this.TargetType = &targetType
	return &this
}

// NewWorkflowWebApiWithDefaults instantiates a new WorkflowWebApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWebApiWithDefaults() *WorkflowWebApi {
	this := WorkflowWebApi{}
	var endpointRequestType string = "Internal"
	this.EndpointRequestType = &endpointRequestType
	var targetType string = "Endpoint"
	this.TargetType = &targetType
	return &this
}

// GetEndpointRequestType returns the EndpointRequestType field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetEndpointRequestType() string {
	if o == nil || o.EndpointRequestType == nil {
		var ret string
		return ret
	}
	return *o.EndpointRequestType
}

// GetEndpointRequestTypeOk returns a tuple with the EndpointRequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetEndpointRequestTypeOk() (*string, bool) {
	if o == nil || o.EndpointRequestType == nil {
		return nil, false
	}
	return o.EndpointRequestType, true
}

// HasEndpointRequestType returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasEndpointRequestType() bool {
	if o != nil && o.EndpointRequestType != nil {
		return true
	}

	return false
}

// SetEndpointRequestType gets a reference to the given string and assigns it to the EndpointRequestType field.
func (o *WorkflowWebApi) SetEndpointRequestType(v string) {
	o.EndpointRequestType = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetHeaders() map[string]interface{} {
	if o == nil || o.Headers == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetHeadersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *WorkflowWebApi) SetHeaders(v map[string]interface{}) {
	o.Headers = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *WorkflowWebApi) SetMethod(v string) {
	o.Method = &v
}

// GetMoType returns the MoType field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetMoType() string {
	if o == nil || o.MoType == nil {
		var ret string
		return ret
	}
	return *o.MoType
}

// GetMoTypeOk returns a tuple with the MoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetMoTypeOk() (*string, bool) {
	if o == nil || o.MoType == nil {
		return nil, false
	}
	return o.MoType, true
}

// HasMoType returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasMoType() bool {
	if o != nil && o.MoType != nil {
		return true
	}

	return false
}

// SetMoType gets a reference to the given string and assigns it to the MoType field.
func (o *WorkflowWebApi) SetMoType(v string) {
	o.MoType = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *WorkflowWebApi) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetTargetType() string {
	if o == nil || o.TargetType == nil {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetTargetTypeOk() (*string, bool) {
	if o == nil || o.TargetType == nil {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasTargetType() bool {
	if o != nil && o.TargetType != nil {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *WorkflowWebApi) SetTargetType(v string) {
	o.TargetType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WorkflowWebApi) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWebApi) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WorkflowWebApi) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WorkflowWebApi) SetUrl(v string) {
	o.Url = &v
}

func (o WorkflowWebApi) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowApi, errWorkflowApi := json.Marshal(o.WorkflowApi)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	errWorkflowApi = json.Unmarshal([]byte(serializedWorkflowApi), &toSerialize)
	if errWorkflowApi != nil {
		return []byte{}, errWorkflowApi
	}
	if o.EndpointRequestType != nil {
		toSerialize["EndpointRequestType"] = o.EndpointRequestType
	}
	if o.Headers != nil {
		toSerialize["Headers"] = o.Headers
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.MoType != nil {
		toSerialize["MoType"] = o.MoType
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.TargetType != nil {
		toSerialize["TargetType"] = o.TargetType
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowWebApi struct {
	value *WorkflowWebApi
	isSet bool
}

func (v NullableWorkflowWebApi) Get() *WorkflowWebApi {
	return v.value
}

func (v *NullableWorkflowWebApi) Set(val *WorkflowWebApi) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWebApi) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWebApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWebApi(val *WorkflowWebApi) *NullableWorkflowWebApi {
	return &NullableWorkflowWebApi{value: val, isSet: true}
}

func (v NullableWorkflowWebApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWebApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
