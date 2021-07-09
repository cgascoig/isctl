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

// WorkflowApi Intersight Orchestrator supports generic API workflow tasks that can execute an API given the request body and response parser specification. API type models a single API request within a batch of requests that get executed within a single workflow task.
type WorkflowApi struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// Asset target defines the remote target endpoints which are either managed by Intersight or their service APIs are invoked from Intersight. Generic API executor service Jasmine can invoke these remote APIs via its executors. Asset targets can be accessed directly for cloud targets or via an associated Intersight Assist. Prerequisite to use asset targets is to persist the target details. Asset target MoRef can be given as task input of type TargetDataType. Fusion determines and populates the target context with the Assist if associated with. It is set internally at the API level. In case of an associated assist, it is used by Assist to fetch the target details and form the API request to send to endpoints. In case of cloud asset targets, Jasmine fetched the target details from DB, forms the API request and sends it to the target.
	AssetTargetMoid *string `json:"AssetTargetMoid,omitempty"`
	// The optional request body that is sent as part of this API request. The request body can contain a golang template that can be populated with task input parameters and previous API output parameters.
	Body *string `json:"Body,omitempty"`
	// Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. The type of the content that gets passed as payload and response in this API. The supported values are json, xml, text.
	ContentType *string `json:"ContentType,omitempty"`
	// A description that task designer can add to individual API requests that explain  what the API call is about.
	Description *string `json:"Description,omitempty"`
	// Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. Optional input to specify the content type in case of error API response. This should be used if the content type of error response is different from that of the success response. If not specified, contentType input value is used to parse the error response.
	ErrorContentType *string `json:"ErrorContentType,omitempty"`
	// A user friendly label that task designers have given to the batch API request.
	Label *string `json:"Label,omitempty"`
	// A reference name for this API request within the batch API request. This name shall be used to map the API output parameters to subsequent API input parameters within a batch API task.
	Name *string `json:"Name,omitempty"`
	// All the possible outcomes of this API are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
	Outcomes interface{} `json:"Outcomes,omitempty"`
	// The optional grammar specification for parsing the response to extract the required values. The specification should have extraction specification specified for all the API output parameters.
	ResponseSpec interface{} `json:"ResponseSpec,omitempty"`
	// The skip expression, if provided, allows the batch API executor to skip the api execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
	SkipOnCondition *string `json:"SkipOnCondition,omitempty"`
	// The delay in seconds after which the API needs to be executed. By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution. Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out.
	StartDelay *int64 `json:"StartDelay,omitempty"`
	// The duration in seconds by which the API response is expected from the API target. If the end point does not respond for the API request within this timeout duration, the task will be marked as failed.
	Timeout              *int64 `json:"Timeout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowApi WorkflowApi

// NewWorkflowApi instantiates a new WorkflowApi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowApi(classId string, objectType string) *WorkflowApi {
	this := WorkflowApi{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowApiWithDefaults instantiates a new WorkflowApi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowApiWithDefaults() *WorkflowApi {
	this := WorkflowApi{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowApi) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowApi) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowApi) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowApi) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAssetTargetMoid returns the AssetTargetMoid field value if set, zero value otherwise.
func (o *WorkflowApi) GetAssetTargetMoid() string {
	if o == nil || o.AssetTargetMoid == nil {
		var ret string
		return ret
	}
	return *o.AssetTargetMoid
}

// GetAssetTargetMoidOk returns a tuple with the AssetTargetMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetAssetTargetMoidOk() (*string, bool) {
	if o == nil || o.AssetTargetMoid == nil {
		return nil, false
	}
	return o.AssetTargetMoid, true
}

// HasAssetTargetMoid returns a boolean if a field has been set.
func (o *WorkflowApi) HasAssetTargetMoid() bool {
	if o != nil && o.AssetTargetMoid != nil {
		return true
	}

	return false
}

// SetAssetTargetMoid gets a reference to the given string and assigns it to the AssetTargetMoid field.
func (o *WorkflowApi) SetAssetTargetMoid(v string) {
	o.AssetTargetMoid = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *WorkflowApi) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *WorkflowApi) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *WorkflowApi) SetBody(v string) {
	o.Body = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *WorkflowApi) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *WorkflowApi) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *WorkflowApi) SetContentType(v string) {
	o.ContentType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowApi) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowApi) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowApi) SetDescription(v string) {
	o.Description = &v
}

// GetErrorContentType returns the ErrorContentType field value if set, zero value otherwise.
func (o *WorkflowApi) GetErrorContentType() string {
	if o == nil || o.ErrorContentType == nil {
		var ret string
		return ret
	}
	return *o.ErrorContentType
}

// GetErrorContentTypeOk returns a tuple with the ErrorContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetErrorContentTypeOk() (*string, bool) {
	if o == nil || o.ErrorContentType == nil {
		return nil, false
	}
	return o.ErrorContentType, true
}

// HasErrorContentType returns a boolean if a field has been set.
func (o *WorkflowApi) HasErrorContentType() bool {
	if o != nil && o.ErrorContentType != nil {
		return true
	}

	return false
}

// SetErrorContentType gets a reference to the given string and assigns it to the ErrorContentType field.
func (o *WorkflowApi) SetErrorContentType(v string) {
	o.ErrorContentType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowApi) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowApi) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowApi) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowApi) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowApi) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowApi) SetName(v string) {
	o.Name = &v
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowApi) GetOutcomes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowApi) GetOutcomesOk() (*interface{}, bool) {
	if o == nil || o.Outcomes == nil {
		return nil, false
	}
	return &o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *WorkflowApi) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given interface{} and assigns it to the Outcomes field.
func (o *WorkflowApi) SetOutcomes(v interface{}) {
	o.Outcomes = v
}

// GetResponseSpec returns the ResponseSpec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowApi) GetResponseSpec() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResponseSpec
}

// GetResponseSpecOk returns a tuple with the ResponseSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowApi) GetResponseSpecOk() (*interface{}, bool) {
	if o == nil || o.ResponseSpec == nil {
		return nil, false
	}
	return &o.ResponseSpec, true
}

// HasResponseSpec returns a boolean if a field has been set.
func (o *WorkflowApi) HasResponseSpec() bool {
	if o != nil && o.ResponseSpec != nil {
		return true
	}

	return false
}

// SetResponseSpec gets a reference to the given interface{} and assigns it to the ResponseSpec field.
func (o *WorkflowApi) SetResponseSpec(v interface{}) {
	o.ResponseSpec = v
}

// GetSkipOnCondition returns the SkipOnCondition field value if set, zero value otherwise.
func (o *WorkflowApi) GetSkipOnCondition() string {
	if o == nil || o.SkipOnCondition == nil {
		var ret string
		return ret
	}
	return *o.SkipOnCondition
}

// GetSkipOnConditionOk returns a tuple with the SkipOnCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetSkipOnConditionOk() (*string, bool) {
	if o == nil || o.SkipOnCondition == nil {
		return nil, false
	}
	return o.SkipOnCondition, true
}

// HasSkipOnCondition returns a boolean if a field has been set.
func (o *WorkflowApi) HasSkipOnCondition() bool {
	if o != nil && o.SkipOnCondition != nil {
		return true
	}

	return false
}

// SetSkipOnCondition gets a reference to the given string and assigns it to the SkipOnCondition field.
func (o *WorkflowApi) SetSkipOnCondition(v string) {
	o.SkipOnCondition = &v
}

// GetStartDelay returns the StartDelay field value if set, zero value otherwise.
func (o *WorkflowApi) GetStartDelay() int64 {
	if o == nil || o.StartDelay == nil {
		var ret int64
		return ret
	}
	return *o.StartDelay
}

// GetStartDelayOk returns a tuple with the StartDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetStartDelayOk() (*int64, bool) {
	if o == nil || o.StartDelay == nil {
		return nil, false
	}
	return o.StartDelay, true
}

// HasStartDelay returns a boolean if a field has been set.
func (o *WorkflowApi) HasStartDelay() bool {
	if o != nil && o.StartDelay != nil {
		return true
	}

	return false
}

// SetStartDelay gets a reference to the given int64 and assigns it to the StartDelay field.
func (o *WorkflowApi) SetStartDelay(v int64) {
	o.StartDelay = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *WorkflowApi) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowApi) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *WorkflowApi) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *WorkflowApi) SetTimeout(v int64) {
	o.Timeout = &v
}

func (o WorkflowApi) MarshalJSON() ([]byte, error) {
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
	if o.AssetTargetMoid != nil {
		toSerialize["AssetTargetMoid"] = o.AssetTargetMoid
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.ContentType != nil {
		toSerialize["ContentType"] = o.ContentType
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.ErrorContentType != nil {
		toSerialize["ErrorContentType"] = o.ErrorContentType
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Outcomes != nil {
		toSerialize["Outcomes"] = o.Outcomes
	}
	if o.ResponseSpec != nil {
		toSerialize["ResponseSpec"] = o.ResponseSpec
	}
	if o.SkipOnCondition != nil {
		toSerialize["SkipOnCondition"] = o.SkipOnCondition
	}
	if o.StartDelay != nil {
		toSerialize["StartDelay"] = o.StartDelay
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowApi) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowApiWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// Asset target defines the remote target endpoints which are either managed by Intersight or their service APIs are invoked from Intersight. Generic API executor service Jasmine can invoke these remote APIs via its executors. Asset targets can be accessed directly for cloud targets or via an associated Intersight Assist. Prerequisite to use asset targets is to persist the target details. Asset target MoRef can be given as task input of type TargetDataType. Fusion determines and populates the target context with the Assist if associated with. It is set internally at the API level. In case of an associated assist, it is used by Assist to fetch the target details and form the API request to send to endpoints. In case of cloud asset targets, Jasmine fetched the target details from DB, forms the API request and sends it to the target.
		AssetTargetMoid *string `json:"AssetTargetMoid,omitempty"`
		// The optional request body that is sent as part of this API request. The request body can contain a golang template that can be populated with task input parameters and previous API output parameters.
		Body *string `json:"Body,omitempty"`
		// Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. The type of the content that gets passed as payload and response in this API. The supported values are json, xml, text.
		ContentType *string `json:"ContentType,omitempty"`
		// A description that task designer can add to individual API requests that explain  what the API call is about.
		Description *string `json:"Description,omitempty"`
		// Intersight Orchestrator, with the support of response parser specification, can extract the values from API responses and map them to task output parameters. The value extraction is supported for response content types XML, JSON and Text. Optional input to specify the content type in case of error API response. This should be used if the content type of error response is different from that of the success response. If not specified, contentType input value is used to parse the error response.
		ErrorContentType *string `json:"ErrorContentType,omitempty"`
		// A user friendly label that task designers have given to the batch API request.
		Label *string `json:"Label,omitempty"`
		// A reference name for this API request within the batch API request. This name shall be used to map the API output parameters to subsequent API input parameters within a batch API task.
		Name *string `json:"Name,omitempty"`
		// All the possible outcomes of this API are captured here. Outcomes property is a collection property of type workflow.Outcome objects. The outcomes can be mapped to the message to be shown. The outcomes are evaluated in the order they are given. At the end of the outcomes list, an catchall success/fail outcome can be added with condition as 'true'. This is an optional property and if not specified the task will be marked as success.
		Outcomes interface{} `json:"Outcomes,omitempty"`
		// The optional grammar specification for parsing the response to extract the required values. The specification should have extraction specification specified for all the API output parameters.
		ResponseSpec interface{} `json:"ResponseSpec,omitempty"`
		// The skip expression, if provided, allows the batch API executor to skip the api execution when the given expression evaluates to true. The expression is given as such a golang template that has to be evaluated to a final content true/false. The expression is an optional and in case not provided, the API will always be executed.
		SkipOnCondition *string `json:"SkipOnCondition,omitempty"`
		// The delay in seconds after which the API needs to be executed. By default, the given API is executed immediately. Specifying a start delay adds to the delay to execution. Start Delay is not supported for the first API in the Batch and cumulative delay of all the APIs in the Batch should not exceed the task time out.
		StartDelay *int64 `json:"StartDelay,omitempty"`
		// The duration in seconds by which the API response is expected from the API target. If the end point does not respond for the API request within this timeout duration, the task will be marked as failed.
		Timeout *int64 `json:"Timeout,omitempty"`
	}

	varWorkflowApiWithoutEmbeddedStruct := WorkflowApiWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowApiWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowApi := _WorkflowApi{}
		varWorkflowApi.ClassId = varWorkflowApiWithoutEmbeddedStruct.ClassId
		varWorkflowApi.ObjectType = varWorkflowApiWithoutEmbeddedStruct.ObjectType
		varWorkflowApi.AssetTargetMoid = varWorkflowApiWithoutEmbeddedStruct.AssetTargetMoid
		varWorkflowApi.Body = varWorkflowApiWithoutEmbeddedStruct.Body
		varWorkflowApi.ContentType = varWorkflowApiWithoutEmbeddedStruct.ContentType
		varWorkflowApi.Description = varWorkflowApiWithoutEmbeddedStruct.Description
		varWorkflowApi.ErrorContentType = varWorkflowApiWithoutEmbeddedStruct.ErrorContentType
		varWorkflowApi.Label = varWorkflowApiWithoutEmbeddedStruct.Label
		varWorkflowApi.Name = varWorkflowApiWithoutEmbeddedStruct.Name
		varWorkflowApi.Outcomes = varWorkflowApiWithoutEmbeddedStruct.Outcomes
		varWorkflowApi.ResponseSpec = varWorkflowApiWithoutEmbeddedStruct.ResponseSpec
		varWorkflowApi.SkipOnCondition = varWorkflowApiWithoutEmbeddedStruct.SkipOnCondition
		varWorkflowApi.StartDelay = varWorkflowApiWithoutEmbeddedStruct.StartDelay
		varWorkflowApi.Timeout = varWorkflowApiWithoutEmbeddedStruct.Timeout
		*o = WorkflowApi(varWorkflowApi)
	} else {
		return err
	}

	varWorkflowApi := _WorkflowApi{}

	err = json.Unmarshal(bytes, &varWorkflowApi)
	if err == nil {
		o.MoBaseComplexType = varWorkflowApi.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AssetTargetMoid")
		delete(additionalProperties, "Body")
		delete(additionalProperties, "ContentType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "ErrorContentType")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Outcomes")
		delete(additionalProperties, "ResponseSpec")
		delete(additionalProperties, "SkipOnCondition")
		delete(additionalProperties, "StartDelay")
		delete(additionalProperties, "Timeout")

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

type NullableWorkflowApi struct {
	value *WorkflowApi
	isSet bool
}

func (v NullableWorkflowApi) Get() *WorkflowApi {
	return v.value
}

func (v *NullableWorkflowApi) Set(val *WorkflowApi) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowApi) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowApi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowApi(val *WorkflowApi) *NullableWorkflowApi {
	return &NullableWorkflowApi{value: val, isSet: true}
}

func (v NullableWorkflowApi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowApi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
