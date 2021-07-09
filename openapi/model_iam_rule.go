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
	"reflect"
	"strings"
)

// IamRule Rule stores the information about the IP address. This information is used during the IP address validation of the incoming request.
type IamRule struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The flag represents if IP addresses in the rule is IPv4 or IPv6.
	IpV6 *bool `json:"IpV6,omitempty"`
	// The type of the IP address. Currently three types are supported, ie IP, CIDR range and IP range. * `Ip` - The IP address rule type is IP. * `Cidr` - The IP address rule type is CIDR range. * `IpRange` - The IP address rule type is IP range.
	RuleType             *string  `json:"RuleType,omitempty"`
	RuleValue            []string `json:"RuleValue,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamRule IamRule

// NewIamRule instantiates a new IamRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRule(classId string, objectType string) *IamRule {
	this := IamRule{}
	this.ClassId = classId
	this.ObjectType = objectType
	var ruleType string = "Ip"
	this.RuleType = &ruleType
	return &this
}

// NewIamRuleWithDefaults instantiates a new IamRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRuleWithDefaults() *IamRule {
	this := IamRule{}
	var classId string = "iam.Rule"
	this.ClassId = classId
	var objectType string = "iam.Rule"
	this.ObjectType = objectType
	var ruleType string = "Ip"
	this.RuleType = &ruleType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamRule) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamRule) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamRule) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamRule) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamRule) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamRule) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIpV6 returns the IpV6 field value if set, zero value otherwise.
func (o *IamRule) GetIpV6() bool {
	if o == nil || o.IpV6 == nil {
		var ret bool
		return ret
	}
	return *o.IpV6
}

// GetIpV6Ok returns a tuple with the IpV6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRule) GetIpV6Ok() (*bool, bool) {
	if o == nil || o.IpV6 == nil {
		return nil, false
	}
	return o.IpV6, true
}

// HasIpV6 returns a boolean if a field has been set.
func (o *IamRule) HasIpV6() bool {
	if o != nil && o.IpV6 != nil {
		return true
	}

	return false
}

// SetIpV6 gets a reference to the given bool and assigns it to the IpV6 field.
func (o *IamRule) SetIpV6(v bool) {
	o.IpV6 = &v
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *IamRule) GetRuleType() string {
	if o == nil || o.RuleType == nil {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRule) GetRuleTypeOk() (*string, bool) {
	if o == nil || o.RuleType == nil {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *IamRule) HasRuleType() bool {
	if o != nil && o.RuleType != nil {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *IamRule) SetRuleType(v string) {
	o.RuleType = &v
}

// GetRuleValue returns the RuleValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamRule) GetRuleValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RuleValue
}

// GetRuleValueOk returns a tuple with the RuleValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamRule) GetRuleValueOk() (*[]string, bool) {
	if o == nil || o.RuleValue == nil {
		return nil, false
	}
	return &o.RuleValue, true
}

// HasRuleValue returns a boolean if a field has been set.
func (o *IamRule) HasRuleValue() bool {
	if o != nil && o.RuleValue != nil {
		return true
	}

	return false
}

// SetRuleValue gets a reference to the given []string and assigns it to the RuleValue field.
func (o *IamRule) SetRuleValue(v []string) {
	o.RuleValue = v
}

func (o IamRule) MarshalJSON() ([]byte, error) {
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
	if o.IpV6 != nil {
		toSerialize["IpV6"] = o.IpV6
	}
	if o.RuleType != nil {
		toSerialize["RuleType"] = o.RuleType
	}
	if o.RuleValue != nil {
		toSerialize["RuleValue"] = o.RuleValue
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamRule) UnmarshalJSON(bytes []byte) (err error) {
	type IamRuleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The flag represents if IP addresses in the rule is IPv4 or IPv6.
		IpV6 *bool `json:"IpV6,omitempty"`
		// The type of the IP address. Currently three types are supported, ie IP, CIDR range and IP range. * `Ip` - The IP address rule type is IP. * `Cidr` - The IP address rule type is CIDR range. * `IpRange` - The IP address rule type is IP range.
		RuleType  *string  `json:"RuleType,omitempty"`
		RuleValue []string `json:"RuleValue,omitempty"`
	}

	varIamRuleWithoutEmbeddedStruct := IamRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varIamRuleWithoutEmbeddedStruct)
	if err == nil {
		varIamRule := _IamRule{}
		varIamRule.ClassId = varIamRuleWithoutEmbeddedStruct.ClassId
		varIamRule.ObjectType = varIamRuleWithoutEmbeddedStruct.ObjectType
		varIamRule.IpV6 = varIamRuleWithoutEmbeddedStruct.IpV6
		varIamRule.RuleType = varIamRuleWithoutEmbeddedStruct.RuleType
		varIamRule.RuleValue = varIamRuleWithoutEmbeddedStruct.RuleValue
		*o = IamRule(varIamRule)
	} else {
		return err
	}

	varIamRule := _IamRule{}

	err = json.Unmarshal(bytes, &varIamRule)
	if err == nil {
		o.MoBaseComplexType = varIamRule.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IpV6")
		delete(additionalProperties, "RuleType")
		delete(additionalProperties, "RuleValue")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamRule struct {
	value *IamRule
	isSet bool
}

func (v NullableIamRule) Get() *IamRule {
	return v.value
}

func (v *NullableIamRule) Set(val *IamRule) {
	v.value = val
	v.isSet = true
}

func (v NullableIamRule) IsSet() bool {
	return v.isSet
}

func (v *NullableIamRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamRule(val *IamRule) *NullableIamRule {
	return &NullableIamRule{value: val, isSet: true}
}

func (v NullableIamRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
