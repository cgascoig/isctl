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

// FirmwareDriverDistributable A device driver image distributed by Cisco.
type FirmwareDriverDistributable struct {
	FirmwareBaseDistributable
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The device type on which the driver is installable.
	Category *string `json:"Category,omitempty"`
	// Indicates in which directory path this driver will be added.
	Directory *string `json:"Directory,omitempty"`
	// The operating system name to which this driver is compatible.
	Osname *string `json:"Osname,omitempty"`
	// OS Version. It is populated as part of the image import operation.
	Osversion            *string                                `json:"Osversion,omitempty"`
	Catalog              *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareDriverDistributable FirmwareDriverDistributable

// NewFirmwareDriverDistributable instantiates a new FirmwareDriverDistributable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareDriverDistributable(classId string, objectType string) *FirmwareDriverDistributable {
	this := FirmwareDriverDistributable{}
	this.ClassId = classId
	this.ObjectType = objectType
	var importAction string = "None"
	this.ImportAction = &importAction
	var importState string = "ReadyForImport"
	this.ImportState = &importState
	var vendor string = "Cisco"
	this.Vendor = &vendor
	return &this
}

// NewFirmwareDriverDistributableWithDefaults instantiates a new FirmwareDriverDistributable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareDriverDistributableWithDefaults() *FirmwareDriverDistributable {
	this := FirmwareDriverDistributable{}
	var classId string = "firmware.DriverDistributable"
	this.ClassId = classId
	var objectType string = "firmware.DriverDistributable"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareDriverDistributable) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareDriverDistributable) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareDriverDistributable) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareDriverDistributable) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareDriverDistributable) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareDriverDistributable) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *FirmwareDriverDistributable) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDriverDistributable) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *FirmwareDriverDistributable) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *FirmwareDriverDistributable) SetCategory(v string) {
	o.Category = &v
}

// GetDirectory returns the Directory field value if set, zero value otherwise.
func (o *FirmwareDriverDistributable) GetDirectory() string {
	if o == nil || o.Directory == nil {
		var ret string
		return ret
	}
	return *o.Directory
}

// GetDirectoryOk returns a tuple with the Directory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDriverDistributable) GetDirectoryOk() (*string, bool) {
	if o == nil || o.Directory == nil {
		return nil, false
	}
	return o.Directory, true
}

// HasDirectory returns a boolean if a field has been set.
func (o *FirmwareDriverDistributable) HasDirectory() bool {
	if o != nil && o.Directory != nil {
		return true
	}

	return false
}

// SetDirectory gets a reference to the given string and assigns it to the Directory field.
func (o *FirmwareDriverDistributable) SetDirectory(v string) {
	o.Directory = &v
}

// GetOsname returns the Osname field value if set, zero value otherwise.
func (o *FirmwareDriverDistributable) GetOsname() string {
	if o == nil || o.Osname == nil {
		var ret string
		return ret
	}
	return *o.Osname
}

// GetOsnameOk returns a tuple with the Osname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDriverDistributable) GetOsnameOk() (*string, bool) {
	if o == nil || o.Osname == nil {
		return nil, false
	}
	return o.Osname, true
}

// HasOsname returns a boolean if a field has been set.
func (o *FirmwareDriverDistributable) HasOsname() bool {
	if o != nil && o.Osname != nil {
		return true
	}

	return false
}

// SetOsname gets a reference to the given string and assigns it to the Osname field.
func (o *FirmwareDriverDistributable) SetOsname(v string) {
	o.Osname = &v
}

// GetOsversion returns the Osversion field value if set, zero value otherwise.
func (o *FirmwareDriverDistributable) GetOsversion() string {
	if o == nil || o.Osversion == nil {
		var ret string
		return ret
	}
	return *o.Osversion
}

// GetOsversionOk returns a tuple with the Osversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDriverDistributable) GetOsversionOk() (*string, bool) {
	if o == nil || o.Osversion == nil {
		return nil, false
	}
	return o.Osversion, true
}

// HasOsversion returns a boolean if a field has been set.
func (o *FirmwareDriverDistributable) HasOsversion() bool {
	if o != nil && o.Osversion != nil {
		return true
	}

	return false
}

// SetOsversion gets a reference to the given string and assigns it to the Osversion field.
func (o *FirmwareDriverDistributable) SetOsversion(v string) {
	o.Osversion = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *FirmwareDriverDistributable) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareDriverDistributable) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *FirmwareDriverDistributable) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *FirmwareDriverDistributable) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o FirmwareDriverDistributable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Category != nil {
		toSerialize["Category"] = o.Category
	}
	if o.Directory != nil {
		toSerialize["Directory"] = o.Directory
	}
	if o.Osname != nil {
		toSerialize["Osname"] = o.Osname
	}
	if o.Osversion != nil {
		toSerialize["Osversion"] = o.Osversion
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareDriverDistributable) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareDriverDistributableWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The device type on which the driver is installable.
		Category *string `json:"Category,omitempty"`
		// Indicates in which directory path this driver will be added.
		Directory *string `json:"Directory,omitempty"`
		// The operating system name to which this driver is compatible.
		Osname *string `json:"Osname,omitempty"`
		// OS Version. It is populated as part of the image import operation.
		Osversion *string                                `json:"Osversion,omitempty"`
		Catalog   *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
	}

	varFirmwareDriverDistributableWithoutEmbeddedStruct := FirmwareDriverDistributableWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareDriverDistributableWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareDriverDistributable := _FirmwareDriverDistributable{}
		varFirmwareDriverDistributable.ClassId = varFirmwareDriverDistributableWithoutEmbeddedStruct.ClassId
		varFirmwareDriverDistributable.ObjectType = varFirmwareDriverDistributableWithoutEmbeddedStruct.ObjectType
		varFirmwareDriverDistributable.Category = varFirmwareDriverDistributableWithoutEmbeddedStruct.Category
		varFirmwareDriverDistributable.Directory = varFirmwareDriverDistributableWithoutEmbeddedStruct.Directory
		varFirmwareDriverDistributable.Osname = varFirmwareDriverDistributableWithoutEmbeddedStruct.Osname
		varFirmwareDriverDistributable.Osversion = varFirmwareDriverDistributableWithoutEmbeddedStruct.Osversion
		varFirmwareDriverDistributable.Catalog = varFirmwareDriverDistributableWithoutEmbeddedStruct.Catalog
		*o = FirmwareDriverDistributable(varFirmwareDriverDistributable)
	} else {
		return err
	}

	varFirmwareDriverDistributable := _FirmwareDriverDistributable{}

	err = json.Unmarshal(bytes, &varFirmwareDriverDistributable)
	if err == nil {
		o.FirmwareBaseDistributable = varFirmwareDriverDistributable.FirmwareBaseDistributable
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Category")
		delete(additionalProperties, "Directory")
		delete(additionalProperties, "Osname")
		delete(additionalProperties, "Osversion")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectFirmwareBaseDistributable := reflect.ValueOf(o.FirmwareBaseDistributable)
		for i := 0; i < reflectFirmwareBaseDistributable.Type().NumField(); i++ {
			t := reflectFirmwareBaseDistributable.Type().Field(i)

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

type NullableFirmwareDriverDistributable struct {
	value *FirmwareDriverDistributable
	isSet bool
}

func (v NullableFirmwareDriverDistributable) Get() *FirmwareDriverDistributable {
	return v.value
}

func (v *NullableFirmwareDriverDistributable) Set(val *FirmwareDriverDistributable) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareDriverDistributable) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareDriverDistributable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareDriverDistributable(val *FirmwareDriverDistributable) *NullableFirmwareDriverDistributable {
	return &NullableFirmwareDriverDistributable{value: val, isSet: true}
}

func (v NullableFirmwareDriverDistributable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareDriverDistributable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
