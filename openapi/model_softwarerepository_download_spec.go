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

// SoftwarerepositoryDownloadSpec The URL, certificate and other associated information required to download an image listed in an Intersight catalog.
type SoftwarerepositoryDownloadSpec struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The OAuth2 token that will be used during image download by the endpoint to authenticate with file server.
	AuthToken *string `json:"AuthToken,omitempty" yaml:"AuthToken,omitempty"`
	// The certificate of file server that will be used by the endpoint to validate the server before starting the file download.
	Certificate *string `json:"Certificate,omitempty" yaml:"Certificate,omitempty"`
	// The name of the firmware image.
	Filename *string `json:"Filename,omitempty" yaml:"Filename,omitempty"`
	// MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image, post download.
	Md5sum *string `json:"Md5sum,omitempty" yaml:"Md5sum,omitempty"`
	// The size (in bytes) of the firmware image.
	Size *int64 `json:"Size,omitempty" yaml:"Size,omitempty"`
	// The URL of this file in file server. The endpoint uses this URL to download the file from the file server.
	Url  *string                             `json:"Url,omitempty" yaml:"Url,omitempty"`
	File *SoftwarerepositoryFileRelationship `json:"File,omitempty" yaml:"File,omitempty"`
}

// NewSoftwarerepositoryDownloadSpec instantiates a new SoftwarerepositoryDownloadSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwarerepositoryDownloadSpec(classId string, objectType string) *SoftwarerepositoryDownloadSpec {
	this := SoftwarerepositoryDownloadSpec{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSoftwarerepositoryDownloadSpecWithDefaults instantiates a new SoftwarerepositoryDownloadSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwarerepositoryDownloadSpecWithDefaults() *SoftwarerepositoryDownloadSpec {
	this := SoftwarerepositoryDownloadSpec{}
	var classId string = "softwarerepository.DownloadSpec"
	this.ClassId = classId
	var objectType string = "softwarerepository.DownloadSpec"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SoftwarerepositoryDownloadSpec) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SoftwarerepositoryDownloadSpec) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SoftwarerepositoryDownloadSpec) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SoftwarerepositoryDownloadSpec) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpec) GetAuthToken() string {
	if o == nil || o.AuthToken == nil {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetAuthTokenOk() (*string, bool) {
	if o == nil || o.AuthToken == nil {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpec) HasAuthToken() bool {
	if o != nil && o.AuthToken != nil {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *SoftwarerepositoryDownloadSpec) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpec) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpec) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *SoftwarerepositoryDownloadSpec) SetCertificate(v string) {
	o.Certificate = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpec) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpec) HasFilename() bool {
	if o != nil && o.Filename != nil {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *SoftwarerepositoryDownloadSpec) SetFilename(v string) {
	o.Filename = &v
}

// GetMd5sum returns the Md5sum field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpec) GetMd5sum() string {
	if o == nil || o.Md5sum == nil {
		var ret string
		return ret
	}
	return *o.Md5sum
}

// GetMd5sumOk returns a tuple with the Md5sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetMd5sumOk() (*string, bool) {
	if o == nil || o.Md5sum == nil {
		return nil, false
	}
	return o.Md5sum, true
}

// HasMd5sum returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpec) HasMd5sum() bool {
	if o != nil && o.Md5sum != nil {
		return true
	}

	return false
}

// SetMd5sum gets a reference to the given string and assigns it to the Md5sum field.
func (o *SoftwarerepositoryDownloadSpec) SetMd5sum(v string) {
	o.Md5sum = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpec) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpec) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SoftwarerepositoryDownloadSpec) SetSize(v int64) {
	o.Size = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpec) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpec) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SoftwarerepositoryDownloadSpec) SetUrl(v string) {
	o.Url = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *SoftwarerepositoryDownloadSpec) GetFile() SoftwarerepositoryFileRelationship {
	if o == nil || o.File == nil {
		var ret SoftwarerepositoryFileRelationship
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwarerepositoryDownloadSpec) GetFileOk() (*SoftwarerepositoryFileRelationship, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *SoftwarerepositoryDownloadSpec) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given SoftwarerepositoryFileRelationship and assigns it to the File field.
func (o *SoftwarerepositoryDownloadSpec) SetFile(v SoftwarerepositoryFileRelationship) {
	o.File = &v
}

func (o SoftwarerepositoryDownloadSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AuthToken != nil {
		toSerialize["AuthToken"] = o.AuthToken
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}
	if o.Filename != nil {
		toSerialize["Filename"] = o.Filename
	}
	if o.Md5sum != nil {
		toSerialize["Md5sum"] = o.Md5sum
	}
	if o.Size != nil {
		toSerialize["Size"] = o.Size
	}
	if o.Url != nil {
		toSerialize["Url"] = o.Url
	}
	if o.File != nil {
		toSerialize["File"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwarerepositoryDownloadSpec struct {
	value *SoftwarerepositoryDownloadSpec
	isSet bool
}

func (v NullableSoftwarerepositoryDownloadSpec) Get() *SoftwarerepositoryDownloadSpec {
	return v.value
}

func (v *NullableSoftwarerepositoryDownloadSpec) Set(val *SoftwarerepositoryDownloadSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwarerepositoryDownloadSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwarerepositoryDownloadSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwarerepositoryDownloadSpec(val *SoftwarerepositoryDownloadSpec) *NullableSoftwarerepositoryDownloadSpec {
	return &NullableSoftwarerepositoryDownloadSpec{value: val, isSet: true}
}

func (v NullableSoftwarerepositoryDownloadSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwarerepositoryDownloadSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
