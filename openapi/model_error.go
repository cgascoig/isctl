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

// Error struct for Error
type Error struct {
	// A value that is used to determine the nature of the error, and why it occurred.
	Code string `json:"code" yaml:"code"`
	// A language-independent identifier of the specific error message.
	MessageId *string `json:"messageId,omitempty" yaml:"messageId,omitempty"`
	// A human-readable description of the error, localized with the i18n standard. The language is determined from the \"Accept-Language\" request-header. The Accept-Language request-header restricts the set of natural languages that are preferred as a response to the request. See RFC 2616 for more details. Codes: 1. **InternalServerError**   An internal error occurred. 1. **InvalidMethod**         The HTTP method (POST, PUT...) is invalid for this API path. For example, a POST request was sent but this API path only supports GET. 1. **InvalidUrl**            The HTTP request contains an invalid URL. 1. **InvalidRequest**        The HTTP request contains an invalid or malformed message body. 1. **NotSupported**          The request is not supported for the specified REST resource. 1. **NotImplemented**        This API path is experimental and not implemented yet. 1. **NotFound**              The requested REST resource does not exist. 1. **AuthenticationFailure** The request lacks valid authentication credentials. 1. **UnauthorizedOperation** The client is not authorized to perform the operation, such as when the user has insufficient privileges. 1. **ValidationConflict**    The request contains conflicting attributes, such as two mutually exclusive attribute values. 1. **ServiceUnavailable**    See RFC 7231, status 503.
	Message string `json:"message" yaml:"message"`
	Cause   *Error `json:"cause,omitempty" yaml:"cause,omitempty"`
}

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError(code string, message string) *Error {
	this := Error{}
	this.Code = code
	this.Message = message
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetCode returns the Code field value
func (o *Error) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Error) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Error) SetCode(v string) {
	o.Code = v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *Error) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *Error) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *Error) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessage returns the Message field value
func (o *Error) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Error) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Error) SetMessage(v string) {
	o.Message = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *Error) GetCause() Error {
	if o == nil || o.Cause == nil {
		var ret Error
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetCauseOk() (*Error, bool) {
	if o == nil || o.Cause == nil {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *Error) HasCause() bool {
	if o != nil && o.Cause != nil {
		return true
	}

	return false
}

// SetCause gets a reference to the given Error and assigns it to the Cause field.
func (o *Error) SetCause(v Error) {
	o.Cause = &v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Cause != nil {
		toSerialize["cause"] = o.Cause
	}
	return json.Marshal(toSerialize)
}

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
