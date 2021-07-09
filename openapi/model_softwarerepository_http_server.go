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

// SoftwarerepositoryHttpServer An external HTTP file server. This can represent the cisco.com website or a HTTP server in the user's datacenter.
type SoftwarerepositoryHttpServer struct {
	SoftwarerepositoryFileServer
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// HTTP/HTTPS link to the image. Accepted formats are HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.
	LocationLink *string `json:"LocationLink,omitempty"`
	// Password as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories.
	Password *string `json:"Password,omitempty"`
	// Username as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories.
	Username             *string `json:"Username,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SoftwarerepositoryHttpServer SoftwarerepositoryHttpServer

// NewSoftwarerepositoryHttpServer instantiates a new SoftwarerepositoryHttpServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryHttpServer(classId string, objectType string) *SoftwarerepositoryHttpServer {
	this := SoftwarerepositoryHttpServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// NewSoftwarerepositoryHttpServerWithDefaults instantiates a new SoftwarerepositoryHttpServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryHttpServerWithDefaults() *SoftwarerepositoryHttpServer {
	this := SoftwarerepositoryHttpServer{}
	var classId string = "softwarerepository.HttpServer"
	this.ClassId = classId
	var objectType string = "softwarerepository.HttpServer"
	this.ObjectType = objectType
	var isPasswordSet bool = false
	this.IsPasswordSet = &isPasswordSet
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryHttpServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryHttpServer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryHttpServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryHttpServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServer) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServer) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServer) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *SoftwarerepositoryHttpServer) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetLocationLink returns the LocationLink field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServer) GetLocationLink() string {
	if o == nil || o.LocationLink == nil {
		var ret string
		return ret
	}
	return *o.LocationLink
}

// GetLocationLinkOk returns a tuple with the LocationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServer) GetLocationLinkOk() (*string, bool) {
	if o == nil || o.LocationLink == nil {
		return nil, false
	}
	return o.LocationLink, true
}

// HasLocationLink returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServer) HasLocationLink() bool {
	if o != nil && o.LocationLink != nil {
		return true
	}

	return false
}

// SetLocationLink gets a reference to the given string and assigns it to the LocationLink field.
func (o *SoftwarerepositoryHttpServer) SetLocationLink(v string) {
	o.LocationLink = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServer) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServer) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServer) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *SoftwarerepositoryHttpServer) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SoftwarerepositoryHttpServer) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryHttpServer) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SoftwarerepositoryHttpServer) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SoftwarerepositoryHttpServer) SetUsername(v string) {
	o.Username = &v
}

func (o SoftwarerepositoryHttpServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSoftwarerepositoryFileServer, errSoftwarerepositoryFileServer := json.Marshal(o.SoftwarerepositoryFileServer)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}
	errSoftwarerepositoryFileServer = json.Unmarshal([]byte(serializedSoftwarerepositoryFileServer), &toSerialize)
	if errSoftwarerepositoryFileServer != nil {
		return []byte{}, errSoftwarerepositoryFileServer
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.LocationLink != nil {
		toSerialize["LocationLink"] = o.LocationLink
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SoftwarerepositoryHttpServer) UnmarshalJSON(bytes []byte) (err error) {
	type SoftwarerepositoryHttpServerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// HTTP/HTTPS link to the image. Accepted formats are HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image.
		LocationLink *string `json:"LocationLink,omitempty"`
		// Password as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories.
		Password *string `json:"Password,omitempty"`
		// Username as configured on the HTTP[S] server for user authentication. It is generally required to authenticate user provided HTTP[S] based software repositories.
		Username *string `json:"Username,omitempty"`
	}

	varSoftwarerepositoryHttpServerWithoutEmbeddedStruct := SoftwarerepositoryHttpServerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryHttpServerWithoutEmbeddedStruct)
	if err == nil {
		varSoftwarerepositoryHttpServer := _SoftwarerepositoryHttpServer{}
		varSoftwarerepositoryHttpServer.ClassId = varSoftwarerepositoryHttpServerWithoutEmbeddedStruct.ClassId
		varSoftwarerepositoryHttpServer.ObjectType = varSoftwarerepositoryHttpServerWithoutEmbeddedStruct.ObjectType
		varSoftwarerepositoryHttpServer.IsPasswordSet = varSoftwarerepositoryHttpServerWithoutEmbeddedStruct.IsPasswordSet
		varSoftwarerepositoryHttpServer.LocationLink = varSoftwarerepositoryHttpServerWithoutEmbeddedStruct.LocationLink
		varSoftwarerepositoryHttpServer.Password = varSoftwarerepositoryHttpServerWithoutEmbeddedStruct.Password
		varSoftwarerepositoryHttpServer.Username = varSoftwarerepositoryHttpServerWithoutEmbeddedStruct.Username
		*o = SoftwarerepositoryHttpServer(varSoftwarerepositoryHttpServer)
	} else {
		return err
	}

	varSoftwarerepositoryHttpServer := _SoftwarerepositoryHttpServer{}

	err = json.Unmarshal(bytes, &varSoftwarerepositoryHttpServer)
	if err == nil {
		o.SoftwarerepositoryFileServer = varSoftwarerepositoryHttpServer.SoftwarerepositoryFileServer
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "LocationLink")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "Username")

		// remove fields from embedded structs
		reflectSoftwarerepositoryFileServer := reflect.ValueOf(o.SoftwarerepositoryFileServer)
		for i := 0; i < reflectSoftwarerepositoryFileServer.Type().NumField(); i++ {
			t := reflectSoftwarerepositoryFileServer.Type().Field(i)

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

type NullableSoftwarerepositoryHttpServer struct {
	value *SoftwarerepositoryHttpServer
	isSet bool
}

func (v NullableSoftwarerepositoryHttpServer) Get() *SoftwarerepositoryHttpServer {
	return v.value
}

func (v *NullableSoftwarerepositoryHttpServer) Set(val *SoftwarerepositoryHttpServer) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryHttpServer) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryHttpServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryHttpServer(val *SoftwarerepositoryHttpServer) *NullableSoftwarerepositoryHttpServer {
	return &NullableSoftwarerepositoryHttpServer{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryHttpServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryHttpServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
