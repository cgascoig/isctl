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

// AssetDeviceContractInformation Contains information about the Cisco device identified by a unique identifier like serial number. It also contains information about warranty, contract status, validity of the device. In future this object could be expanded to store Case, RMA, device topology details. We observe new asset.DeviceRegisteration and use it as a trigger for creating an instance of this object. Currently the data is restricted to Cisco Standalone IMC servers and Fabric Interconnects. Support for more product lines will be added in future.
type AssetDeviceContractInformation struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	Contract *AssetContractInformation `json:"Contract,omitempty" yaml:"Contract,omitempty"`
	// Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.
	ContractStatus *string `json:"ContractStatus,omitempty" yaml:"ContractStatus,omitempty"`
	// End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API.
	CoveredProductLineEndDate *string `json:"CoveredProductLineEndDate,omitempty" yaml:"CoveredProductLineEndDate,omitempty"`
	// Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases.
	DeviceId *string `json:"DeviceId,omitempty" yaml:"DeviceId,omitempty"`
	// Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future.
	DeviceType            *string                   `json:"DeviceType,omitempty" yaml:"DeviceType,omitempty"`
	EndCustomer           *AssetCustomerInformation `json:"EndCustomer,omitempty" yaml:"EndCustomer,omitempty"`
	EndUserGlobalUltimate *AssetGlobalUltimate      `json:"EndUserGlobalUltimate,omitempty" yaml:"EndUserGlobalUltimate,omitempty"`
	// Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs.
	IsValid *bool `json:"IsValid,omitempty" yaml:"IsValid,omitempty"`
	// Item type of this specific Cisco device. example \"Chassis\".
	ItemType *string `json:"ItemType,omitempty" yaml:"ItemType,omitempty"`
	// Maintenance purchase order number for the Cisco device.
	MaintenancePurchaseOrderNumber *string `json:"MaintenancePurchaseOrderNumber,omitempty" yaml:"MaintenancePurchaseOrderNumber,omitempty"`
	// Maintenance sales order number for the Cisco device.
	MaintenanceSalesOrderNumber *string `json:"MaintenanceSalesOrderNumber,omitempty" yaml:"MaintenanceSalesOrderNumber,omitempty"`
	// The platform type of the Cisco device.
	PlatformType *string                  `json:"PlatformType,omitempty" yaml:"PlatformType,omitempty"`
	Product      *AssetProductInformation `json:"Product,omitempty" yaml:"Product,omitempty"`
	// Purchase order number for the Cisco device. It is a unique number assigned for every purchase.
	PurchaseOrderNumber    *string              `json:"PurchaseOrderNumber,omitempty" yaml:"PurchaseOrderNumber,omitempty"`
	ResellerGlobalUltimate *AssetGlobalUltimate `json:"ResellerGlobalUltimate,omitempty" yaml:"ResellerGlobalUltimate,omitempty"`
	// Sales order number for the Cisco device. It is a unique number assigned for every sale.
	SalesOrderNumber *string `json:"SalesOrderNumber,omitempty" yaml:"SalesOrderNumber,omitempty"`
	// The type of service contract that covers the Cisco device.
	ServiceDescription *string `json:"ServiceDescription,omitempty" yaml:"ServiceDescription,omitempty"`
	// End date for the Cisco service contract that covers this Cisco device.
	ServiceEndDate *time.Time `json:"ServiceEndDate,omitempty" yaml:"ServiceEndDate,omitempty"`
	// The type of service contract that covers the Cisco device.
	ServiceLevel *string `json:"ServiceLevel,omitempty" yaml:"ServiceLevel,omitempty"`
	// The SKU of the service contract that covers the Cisco device.
	ServiceSku *string `json:"ServiceSku,omitempty" yaml:"ServiceSku,omitempty"`
	// Start date for the Cisco service contract that covers this Cisco device.
	ServiceStartDate *time.Time `json:"ServiceStartDate,omitempty" yaml:"ServiceStartDate,omitempty"`
	// Internal property used for triggering and tracking actions for contract information.
	StateContract *string `json:"StateContract,omitempty" yaml:"StateContract,omitempty"`
	// End date for the warranty that covers the Cisco device.
	WarrantyEndDate *string `json:"WarrantyEndDate,omitempty" yaml:"WarrantyEndDate,omitempty"`
	// Type of warranty that covers the Cisco device.
	WarrantyType     *string                              `json:"WarrantyType,omitempty" yaml:"WarrantyType,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewAssetDeviceContractInformation instantiates a new AssetDeviceContractInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetDeviceContractInformation() *AssetDeviceContractInformation {
	this := AssetDeviceContractInformation{}
	var contractStatus string = "Not Covered"
	this.ContractStatus = &contractStatus
	var deviceType string = "None"
	this.DeviceType = &deviceType
	var platformType string = ""
	this.PlatformType = &platformType
	var stateContract string = "Update"
	this.StateContract = &stateContract
	return &this
}

// NewAssetDeviceContractInformationWithDefaults instantiates a new AssetDeviceContractInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetDeviceContractInformationWithDefaults() *AssetDeviceContractInformation {
	this := AssetDeviceContractInformation{}
	var contractStatus string = "Not Covered"
	this.ContractStatus = &contractStatus
	var deviceType string = "None"
	this.DeviceType = &deviceType
	var platformType string = ""
	this.PlatformType = &platformType
	var stateContract string = "Update"
	this.StateContract = &stateContract
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetContract() AssetContractInformation {
	if o == nil || o.Contract == nil {
		var ret AssetContractInformation
		return ret
	}
	return *o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetContractOk() (*AssetContractInformation, bool) {
	if o == nil || o.Contract == nil {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasContract() bool {
	if o != nil && o.Contract != nil {
		return true
	}

	return false
}

// SetContract gets a reference to the given AssetContractInformation and assigns it to the Contract field.
func (o *AssetDeviceContractInformation) SetContract(v AssetContractInformation) {
	o.Contract = &v
}

// GetContractStatus returns the ContractStatus field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetContractStatus() string {
	if o == nil || o.ContractStatus == nil {
		var ret string
		return ret
	}
	return *o.ContractStatus
}

// GetContractStatusOk returns a tuple with the ContractStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetContractStatusOk() (*string, bool) {
	if o == nil || o.ContractStatus == nil {
		return nil, false
	}
	return o.ContractStatus, true
}

// HasContractStatus returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasContractStatus() bool {
	if o != nil && o.ContractStatus != nil {
		return true
	}

	return false
}

// SetContractStatus gets a reference to the given string and assigns it to the ContractStatus field.
func (o *AssetDeviceContractInformation) SetContractStatus(v string) {
	o.ContractStatus = &v
}

// GetCoveredProductLineEndDate returns the CoveredProductLineEndDate field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetCoveredProductLineEndDate() string {
	if o == nil || o.CoveredProductLineEndDate == nil {
		var ret string
		return ret
	}
	return *o.CoveredProductLineEndDate
}

// GetCoveredProductLineEndDateOk returns a tuple with the CoveredProductLineEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetCoveredProductLineEndDateOk() (*string, bool) {
	if o == nil || o.CoveredProductLineEndDate == nil {
		return nil, false
	}
	return o.CoveredProductLineEndDate, true
}

// HasCoveredProductLineEndDate returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasCoveredProductLineEndDate() bool {
	if o != nil && o.CoveredProductLineEndDate != nil {
		return true
	}

	return false
}

// SetCoveredProductLineEndDate gets a reference to the given string and assigns it to the CoveredProductLineEndDate field.
func (o *AssetDeviceContractInformation) SetCoveredProductLineEndDate(v string) {
	o.CoveredProductLineEndDate = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *AssetDeviceContractInformation) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *AssetDeviceContractInformation) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetEndCustomer returns the EndCustomer field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetEndCustomer() AssetCustomerInformation {
	if o == nil || o.EndCustomer == nil {
		var ret AssetCustomerInformation
		return ret
	}
	return *o.EndCustomer
}

// GetEndCustomerOk returns a tuple with the EndCustomer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetEndCustomerOk() (*AssetCustomerInformation, bool) {
	if o == nil || o.EndCustomer == nil {
		return nil, false
	}
	return o.EndCustomer, true
}

// HasEndCustomer returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasEndCustomer() bool {
	if o != nil && o.EndCustomer != nil {
		return true
	}

	return false
}

// SetEndCustomer gets a reference to the given AssetCustomerInformation and assigns it to the EndCustomer field.
func (o *AssetDeviceContractInformation) SetEndCustomer(v AssetCustomerInformation) {
	o.EndCustomer = &v
}

// GetEndUserGlobalUltimate returns the EndUserGlobalUltimate field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetEndUserGlobalUltimate() AssetGlobalUltimate {
	if o == nil || o.EndUserGlobalUltimate == nil {
		var ret AssetGlobalUltimate
		return ret
	}
	return *o.EndUserGlobalUltimate
}

// GetEndUserGlobalUltimateOk returns a tuple with the EndUserGlobalUltimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetEndUserGlobalUltimateOk() (*AssetGlobalUltimate, bool) {
	if o == nil || o.EndUserGlobalUltimate == nil {
		return nil, false
	}
	return o.EndUserGlobalUltimate, true
}

// HasEndUserGlobalUltimate returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasEndUserGlobalUltimate() bool {
	if o != nil && o.EndUserGlobalUltimate != nil {
		return true
	}

	return false
}

// SetEndUserGlobalUltimate gets a reference to the given AssetGlobalUltimate and assigns it to the EndUserGlobalUltimate field.
func (o *AssetDeviceContractInformation) SetEndUserGlobalUltimate(v AssetGlobalUltimate) {
	o.EndUserGlobalUltimate = &v
}

// GetIsValid returns the IsValid field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetIsValid() bool {
	if o == nil || o.IsValid == nil {
		var ret bool
		return ret
	}
	return *o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetIsValidOk() (*bool, bool) {
	if o == nil || o.IsValid == nil {
		return nil, false
	}
	return o.IsValid, true
}

// HasIsValid returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasIsValid() bool {
	if o != nil && o.IsValid != nil {
		return true
	}

	return false
}

// SetIsValid gets a reference to the given bool and assigns it to the IsValid field.
func (o *AssetDeviceContractInformation) SetIsValid(v bool) {
	o.IsValid = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetItemType() string {
	if o == nil || o.ItemType == nil {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetItemTypeOk() (*string, bool) {
	if o == nil || o.ItemType == nil {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasItemType() bool {
	if o != nil && o.ItemType != nil {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *AssetDeviceContractInformation) SetItemType(v string) {
	o.ItemType = &v
}

// GetMaintenancePurchaseOrderNumber returns the MaintenancePurchaseOrderNumber field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetMaintenancePurchaseOrderNumber() string {
	if o == nil || o.MaintenancePurchaseOrderNumber == nil {
		var ret string
		return ret
	}
	return *o.MaintenancePurchaseOrderNumber
}

// GetMaintenancePurchaseOrderNumberOk returns a tuple with the MaintenancePurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetMaintenancePurchaseOrderNumberOk() (*string, bool) {
	if o == nil || o.MaintenancePurchaseOrderNumber == nil {
		return nil, false
	}
	return o.MaintenancePurchaseOrderNumber, true
}

// HasMaintenancePurchaseOrderNumber returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasMaintenancePurchaseOrderNumber() bool {
	if o != nil && o.MaintenancePurchaseOrderNumber != nil {
		return true
	}

	return false
}

// SetMaintenancePurchaseOrderNumber gets a reference to the given string and assigns it to the MaintenancePurchaseOrderNumber field.
func (o *AssetDeviceContractInformation) SetMaintenancePurchaseOrderNumber(v string) {
	o.MaintenancePurchaseOrderNumber = &v
}

// GetMaintenanceSalesOrderNumber returns the MaintenanceSalesOrderNumber field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetMaintenanceSalesOrderNumber() string {
	if o == nil || o.MaintenanceSalesOrderNumber == nil {
		var ret string
		return ret
	}
	return *o.MaintenanceSalesOrderNumber
}

// GetMaintenanceSalesOrderNumberOk returns a tuple with the MaintenanceSalesOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetMaintenanceSalesOrderNumberOk() (*string, bool) {
	if o == nil || o.MaintenanceSalesOrderNumber == nil {
		return nil, false
	}
	return o.MaintenanceSalesOrderNumber, true
}

// HasMaintenanceSalesOrderNumber returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasMaintenanceSalesOrderNumber() bool {
	if o != nil && o.MaintenanceSalesOrderNumber != nil {
		return true
	}

	return false
}

// SetMaintenanceSalesOrderNumber gets a reference to the given string and assigns it to the MaintenanceSalesOrderNumber field.
func (o *AssetDeviceContractInformation) SetMaintenanceSalesOrderNumber(v string) {
	o.MaintenanceSalesOrderNumber = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *AssetDeviceContractInformation) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetProduct() AssetProductInformation {
	if o == nil || o.Product == nil {
		var ret AssetProductInformation
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetProductOk() (*AssetProductInformation, bool) {
	if o == nil || o.Product == nil {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasProduct() bool {
	if o != nil && o.Product != nil {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AssetProductInformation and assigns it to the Product field.
func (o *AssetDeviceContractInformation) SetProduct(v AssetProductInformation) {
	o.Product = &v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetPurchaseOrderNumber() string {
	if o == nil || o.PurchaseOrderNumber == nil {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || o.PurchaseOrderNumber == nil {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasPurchaseOrderNumber() bool {
	if o != nil && o.PurchaseOrderNumber != nil {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *AssetDeviceContractInformation) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetResellerGlobalUltimate returns the ResellerGlobalUltimate field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetResellerGlobalUltimate() AssetGlobalUltimate {
	if o == nil || o.ResellerGlobalUltimate == nil {
		var ret AssetGlobalUltimate
		return ret
	}
	return *o.ResellerGlobalUltimate
}

// GetResellerGlobalUltimateOk returns a tuple with the ResellerGlobalUltimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetResellerGlobalUltimateOk() (*AssetGlobalUltimate, bool) {
	if o == nil || o.ResellerGlobalUltimate == nil {
		return nil, false
	}
	return o.ResellerGlobalUltimate, true
}

// HasResellerGlobalUltimate returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasResellerGlobalUltimate() bool {
	if o != nil && o.ResellerGlobalUltimate != nil {
		return true
	}

	return false
}

// SetResellerGlobalUltimate gets a reference to the given AssetGlobalUltimate and assigns it to the ResellerGlobalUltimate field.
func (o *AssetDeviceContractInformation) SetResellerGlobalUltimate(v AssetGlobalUltimate) {
	o.ResellerGlobalUltimate = &v
}

// GetSalesOrderNumber returns the SalesOrderNumber field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetSalesOrderNumber() string {
	if o == nil || o.SalesOrderNumber == nil {
		var ret string
		return ret
	}
	return *o.SalesOrderNumber
}

// GetSalesOrderNumberOk returns a tuple with the SalesOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetSalesOrderNumberOk() (*string, bool) {
	if o == nil || o.SalesOrderNumber == nil {
		return nil, false
	}
	return o.SalesOrderNumber, true
}

// HasSalesOrderNumber returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasSalesOrderNumber() bool {
	if o != nil && o.SalesOrderNumber != nil {
		return true
	}

	return false
}

// SetSalesOrderNumber gets a reference to the given string and assigns it to the SalesOrderNumber field.
func (o *AssetDeviceContractInformation) SetSalesOrderNumber(v string) {
	o.SalesOrderNumber = &v
}

// GetServiceDescription returns the ServiceDescription field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetServiceDescription() string {
	if o == nil || o.ServiceDescription == nil {
		var ret string
		return ret
	}
	return *o.ServiceDescription
}

// GetServiceDescriptionOk returns a tuple with the ServiceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetServiceDescriptionOk() (*string, bool) {
	if o == nil || o.ServiceDescription == nil {
		return nil, false
	}
	return o.ServiceDescription, true
}

// HasServiceDescription returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasServiceDescription() bool {
	if o != nil && o.ServiceDescription != nil {
		return true
	}

	return false
}

// SetServiceDescription gets a reference to the given string and assigns it to the ServiceDescription field.
func (o *AssetDeviceContractInformation) SetServiceDescription(v string) {
	o.ServiceDescription = &v
}

// GetServiceEndDate returns the ServiceEndDate field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetServiceEndDate() time.Time {
	if o == nil || o.ServiceEndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ServiceEndDate
}

// GetServiceEndDateOk returns a tuple with the ServiceEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetServiceEndDateOk() (*time.Time, bool) {
	if o == nil || o.ServiceEndDate == nil {
		return nil, false
	}
	return o.ServiceEndDate, true
}

// HasServiceEndDate returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasServiceEndDate() bool {
	if o != nil && o.ServiceEndDate != nil {
		return true
	}

	return false
}

// SetServiceEndDate gets a reference to the given time.Time and assigns it to the ServiceEndDate field.
func (o *AssetDeviceContractInformation) SetServiceEndDate(v time.Time) {
	o.ServiceEndDate = &v
}

// GetServiceLevel returns the ServiceLevel field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetServiceLevel() string {
	if o == nil || o.ServiceLevel == nil {
		var ret string
		return ret
	}
	return *o.ServiceLevel
}

// GetServiceLevelOk returns a tuple with the ServiceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetServiceLevelOk() (*string, bool) {
	if o == nil || o.ServiceLevel == nil {
		return nil, false
	}
	return o.ServiceLevel, true
}

// HasServiceLevel returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasServiceLevel() bool {
	if o != nil && o.ServiceLevel != nil {
		return true
	}

	return false
}

// SetServiceLevel gets a reference to the given string and assigns it to the ServiceLevel field.
func (o *AssetDeviceContractInformation) SetServiceLevel(v string) {
	o.ServiceLevel = &v
}

// GetServiceSku returns the ServiceSku field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetServiceSku() string {
	if o == nil || o.ServiceSku == nil {
		var ret string
		return ret
	}
	return *o.ServiceSku
}

// GetServiceSkuOk returns a tuple with the ServiceSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetServiceSkuOk() (*string, bool) {
	if o == nil || o.ServiceSku == nil {
		return nil, false
	}
	return o.ServiceSku, true
}

// HasServiceSku returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasServiceSku() bool {
	if o != nil && o.ServiceSku != nil {
		return true
	}

	return false
}

// SetServiceSku gets a reference to the given string and assigns it to the ServiceSku field.
func (o *AssetDeviceContractInformation) SetServiceSku(v string) {
	o.ServiceSku = &v
}

// GetServiceStartDate returns the ServiceStartDate field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetServiceStartDate() time.Time {
	if o == nil || o.ServiceStartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ServiceStartDate
}

// GetServiceStartDateOk returns a tuple with the ServiceStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetServiceStartDateOk() (*time.Time, bool) {
	if o == nil || o.ServiceStartDate == nil {
		return nil, false
	}
	return o.ServiceStartDate, true
}

// HasServiceStartDate returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasServiceStartDate() bool {
	if o != nil && o.ServiceStartDate != nil {
		return true
	}

	return false
}

// SetServiceStartDate gets a reference to the given time.Time and assigns it to the ServiceStartDate field.
func (o *AssetDeviceContractInformation) SetServiceStartDate(v time.Time) {
	o.ServiceStartDate = &v
}

// GetStateContract returns the StateContract field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetStateContract() string {
	if o == nil || o.StateContract == nil {
		var ret string
		return ret
	}
	return *o.StateContract
}

// GetStateContractOk returns a tuple with the StateContract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetStateContractOk() (*string, bool) {
	if o == nil || o.StateContract == nil {
		return nil, false
	}
	return o.StateContract, true
}

// HasStateContract returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasStateContract() bool {
	if o != nil && o.StateContract != nil {
		return true
	}

	return false
}

// SetStateContract gets a reference to the given string and assigns it to the StateContract field.
func (o *AssetDeviceContractInformation) SetStateContract(v string) {
	o.StateContract = &v
}

// GetWarrantyEndDate returns the WarrantyEndDate field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetWarrantyEndDate() string {
	if o == nil || o.WarrantyEndDate == nil {
		var ret string
		return ret
	}
	return *o.WarrantyEndDate
}

// GetWarrantyEndDateOk returns a tuple with the WarrantyEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetWarrantyEndDateOk() (*string, bool) {
	if o == nil || o.WarrantyEndDate == nil {
		return nil, false
	}
	return o.WarrantyEndDate, true
}

// HasWarrantyEndDate returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasWarrantyEndDate() bool {
	if o != nil && o.WarrantyEndDate != nil {
		return true
	}

	return false
}

// SetWarrantyEndDate gets a reference to the given string and assigns it to the WarrantyEndDate field.
func (o *AssetDeviceContractInformation) SetWarrantyEndDate(v string) {
	o.WarrantyEndDate = &v
}

// GetWarrantyType returns the WarrantyType field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetWarrantyType() string {
	if o == nil || o.WarrantyType == nil {
		var ret string
		return ret
	}
	return *o.WarrantyType
}

// GetWarrantyTypeOk returns a tuple with the WarrantyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetWarrantyTypeOk() (*string, bool) {
	if o == nil || o.WarrantyType == nil {
		return nil, false
	}
	return o.WarrantyType, true
}

// HasWarrantyType returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasWarrantyType() bool {
	if o != nil && o.WarrantyType != nil {
		return true
	}

	return false
}

// SetWarrantyType gets a reference to the given string and assigns it to the WarrantyType field.
func (o *AssetDeviceContractInformation) SetWarrantyType(v string) {
	o.WarrantyType = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AssetDeviceContractInformation) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetDeviceContractInformation) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AssetDeviceContractInformation) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AssetDeviceContractInformation) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AssetDeviceContractInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Contract != nil {
		toSerialize["Contract"] = o.Contract
	}
	if o.ContractStatus != nil {
		toSerialize["ContractStatus"] = o.ContractStatus
	}
	if o.CoveredProductLineEndDate != nil {
		toSerialize["CoveredProductLineEndDate"] = o.CoveredProductLineEndDate
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.EndCustomer != nil {
		toSerialize["EndCustomer"] = o.EndCustomer
	}
	if o.EndUserGlobalUltimate != nil {
		toSerialize["EndUserGlobalUltimate"] = o.EndUserGlobalUltimate
	}
	if o.IsValid != nil {
		toSerialize["IsValid"] = o.IsValid
	}
	if o.ItemType != nil {
		toSerialize["ItemType"] = o.ItemType
	}
	if o.MaintenancePurchaseOrderNumber != nil {
		toSerialize["MaintenancePurchaseOrderNumber"] = o.MaintenancePurchaseOrderNumber
	}
	if o.MaintenanceSalesOrderNumber != nil {
		toSerialize["MaintenanceSalesOrderNumber"] = o.MaintenanceSalesOrderNumber
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.Product != nil {
		toSerialize["Product"] = o.Product
	}
	if o.PurchaseOrderNumber != nil {
		toSerialize["PurchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	if o.ResellerGlobalUltimate != nil {
		toSerialize["ResellerGlobalUltimate"] = o.ResellerGlobalUltimate
	}
	if o.SalesOrderNumber != nil {
		toSerialize["SalesOrderNumber"] = o.SalesOrderNumber
	}
	if o.ServiceDescription != nil {
		toSerialize["ServiceDescription"] = o.ServiceDescription
	}
	if o.ServiceEndDate != nil {
		toSerialize["ServiceEndDate"] = o.ServiceEndDate
	}
	if o.ServiceLevel != nil {
		toSerialize["ServiceLevel"] = o.ServiceLevel
	}
	if o.ServiceSku != nil {
		toSerialize["ServiceSku"] = o.ServiceSku
	}
	if o.ServiceStartDate != nil {
		toSerialize["ServiceStartDate"] = o.ServiceStartDate
	}
	if o.StateContract != nil {
		toSerialize["StateContract"] = o.StateContract
	}
	if o.WarrantyEndDate != nil {
		toSerialize["WarrantyEndDate"] = o.WarrantyEndDate
	}
	if o.WarrantyType != nil {
		toSerialize["WarrantyType"] = o.WarrantyType
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableAssetDeviceContractInformation struct {
	value *AssetDeviceContractInformation
	isSet bool
}

func (v NullableAssetDeviceContractInformation) Get() *AssetDeviceContractInformation {
	return v.value
}

func (v *NullableAssetDeviceContractInformation) Set(val *AssetDeviceContractInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetDeviceContractInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetDeviceContractInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetDeviceContractInformation(val *AssetDeviceContractInformation) *NullableAssetDeviceContractInformation {
	return &NullableAssetDeviceContractInformation{value: val, isSet: true}
}

func (v NullableAssetDeviceContractInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetDeviceContractInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
