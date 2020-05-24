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

// NetworkElementAllOf Definition of the list of properties defined in 'network.Element', excluding properties defined in parent classes.
type NetworkElementAllOf struct {
	// The administrative state of the network Element inband management interface.
	AdminInbandInterfaceState *string `json:"AdminInbandInterfaceState,omitempty" yaml:"AdminInbandInterfaceState,omitempty"`
	FaultSummary              *int64  `json:"FaultSummary,omitempty" yaml:"FaultSummary,omitempty"`
	// The IP address of the network Element inband management interface.
	InbandIpAddress *string `json:"InbandIpAddress,omitempty" yaml:"InbandIpAddress,omitempty"`
	// The default gateway of the network Element inband management interface.
	InbandIpGateway *string `json:"InbandIpGateway,omitempty" yaml:"InbandIpGateway,omitempty"`
	// The network mask of the network Element inband management interface.
	InbandIpMask *string `json:"InbandIpMask,omitempty" yaml:"InbandIpMask,omitempty"`
	// The VLAN ID of the network Element inband management interface.
	InbandVlan *int64 `json:"InbandVlan,omitempty" yaml:"InbandVlan,omitempty"`
	// The IP address of the network Element out-of-band management interface.
	OutOfBandIpAddress *string `json:"OutOfBandIpAddress,omitempty" yaml:"OutOfBandIpAddress,omitempty"`
	// The default gateway of the network Element out-of-band management interface.
	OutOfBandIpGateway *string `json:"OutOfBandIpGateway,omitempty" yaml:"OutOfBandIpGateway,omitempty"`
	// The network mask of the network Element out-of-band management interface.
	OutOfBandIpMask *string `json:"OutOfBandIpMask,omitempty" yaml:"OutOfBandIpMask,omitempty"`
	// The IPv4 address of the network Element out-of-band management interface.
	OutOfBandIpv4Address *string `json:"OutOfBandIpv4Address,omitempty" yaml:"OutOfBandIpv4Address,omitempty"`
	// The default IPv4 gateway of the network Element out-of-band management interface.
	OutOfBandIpv4Gateway *string `json:"OutOfBandIpv4Gateway,omitempty" yaml:"OutOfBandIpv4Gateway,omitempty"`
	// The network mask of the network Element out-of-band management interface.
	OutOfBandIpv4Mask *string `json:"OutOfBandIpv4Mask,omitempty" yaml:"OutOfBandIpv4Mask,omitempty"`
	// The IPv6 address of the network Element out-of-band management interface.
	OutOfBandIpv6Address *string `json:"OutOfBandIpv6Address,omitempty" yaml:"OutOfBandIpv6Address,omitempty"`
	// The default IPv6 gateway of the network Element out-of-band management interface.
	OutOfBandIpv6Gateway *string `json:"OutOfBandIpv6Gateway,omitempty" yaml:"OutOfBandIpv6Gateway,omitempty"`
	// The network mask of the network Element out-of-band management interface.
	OutOfBandIpv6Prefix *string `json:"OutOfBandIpv6Prefix,omitempty" yaml:"OutOfBandIpv6Prefix,omitempty"`
	// The MAC address of the network Element out-of-band management interface.
	OutOfBandMac *string `json:"OutOfBandMac,omitempty" yaml:"OutOfBandMac,omitempty"`
	// The Switch Id of the network Element.
	SwitchId *string `json:"SwitchId,omitempty" yaml:"SwitchId,omitempty"`
	// An array of relationships to equipmentSwitchCard resources.
	Cards *[]EquipmentSwitchCardRelationship `json:"Cards,omitempty" yaml:"Cards,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	Fanmodules          *[]EquipmentFanModuleRelationship `json:"Fanmodules,omitempty" yaml:"Fanmodules,omitempty"`
	ManagementContoller *ManagementControllerRelationship `json:"ManagementContoller,omitempty" yaml:"ManagementContoller,omitempty"`
	ManagementEntity    *ManagementEntityRelationship     `json:"ManagementEntity,omitempty" yaml:"ManagementEntity,omitempty"`
	// An array of relationships to equipmentPsu resources.
	Psus                *[]EquipmentPsuRelationship          `json:"Psus,omitempty" yaml:"Psus,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
	TopSystem           *TopSystemRelationship               `json:"TopSystem,omitempty" yaml:"TopSystem,omitempty"`
	UcsmRunningFirmware *FirmwareRunningFirmwareRelationship `json:"UcsmRunningFirmware,omitempty" yaml:"UcsmRunningFirmware,omitempty"`
}

// NewNetworkElementAllOf instantiates a new NetworkElementAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkElementAllOf() *NetworkElementAllOf {
	this := NetworkElementAllOf{}
	return &this
}

// NewNetworkElementAllOfWithDefaults instantiates a new NetworkElementAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkElementAllOfWithDefaults() *NetworkElementAllOf {
	this := NetworkElementAllOf{}
	return &this
}

// GetAdminInbandInterfaceState returns the AdminInbandInterfaceState field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetAdminInbandInterfaceState() string {
	if o == nil || o.AdminInbandInterfaceState == nil {
		var ret string
		return ret
	}
	return *o.AdminInbandInterfaceState
}

// GetAdminInbandInterfaceStateOk returns a tuple with the AdminInbandInterfaceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetAdminInbandInterfaceStateOk() (*string, bool) {
	if o == nil || o.AdminInbandInterfaceState == nil {
		return nil, false
	}
	return o.AdminInbandInterfaceState, true
}

// HasAdminInbandInterfaceState returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasAdminInbandInterfaceState() bool {
	if o != nil && o.AdminInbandInterfaceState != nil {
		return true
	}

	return false
}

// SetAdminInbandInterfaceState gets a reference to the given string and assigns it to the AdminInbandInterfaceState field.
func (o *NetworkElementAllOf) SetAdminInbandInterfaceState(v string) {
	o.AdminInbandInterfaceState = &v
}

// GetFaultSummary returns the FaultSummary field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetFaultSummary() int64 {
	if o == nil || o.FaultSummary == nil {
		var ret int64
		return ret
	}
	return *o.FaultSummary
}

// GetFaultSummaryOk returns a tuple with the FaultSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetFaultSummaryOk() (*int64, bool) {
	if o == nil || o.FaultSummary == nil {
		return nil, false
	}
	return o.FaultSummary, true
}

// HasFaultSummary returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasFaultSummary() bool {
	if o != nil && o.FaultSummary != nil {
		return true
	}

	return false
}

// SetFaultSummary gets a reference to the given int64 and assigns it to the FaultSummary field.
func (o *NetworkElementAllOf) SetFaultSummary(v int64) {
	o.FaultSummary = &v
}

// GetInbandIpAddress returns the InbandIpAddress field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetInbandIpAddress() string {
	if o == nil || o.InbandIpAddress == nil {
		var ret string
		return ret
	}
	return *o.InbandIpAddress
}

// GetInbandIpAddressOk returns a tuple with the InbandIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetInbandIpAddressOk() (*string, bool) {
	if o == nil || o.InbandIpAddress == nil {
		return nil, false
	}
	return o.InbandIpAddress, true
}

// HasInbandIpAddress returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasInbandIpAddress() bool {
	if o != nil && o.InbandIpAddress != nil {
		return true
	}

	return false
}

// SetInbandIpAddress gets a reference to the given string and assigns it to the InbandIpAddress field.
func (o *NetworkElementAllOf) SetInbandIpAddress(v string) {
	o.InbandIpAddress = &v
}

// GetInbandIpGateway returns the InbandIpGateway field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetInbandIpGateway() string {
	if o == nil || o.InbandIpGateway == nil {
		var ret string
		return ret
	}
	return *o.InbandIpGateway
}

// GetInbandIpGatewayOk returns a tuple with the InbandIpGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetInbandIpGatewayOk() (*string, bool) {
	if o == nil || o.InbandIpGateway == nil {
		return nil, false
	}
	return o.InbandIpGateway, true
}

// HasInbandIpGateway returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasInbandIpGateway() bool {
	if o != nil && o.InbandIpGateway != nil {
		return true
	}

	return false
}

// SetInbandIpGateway gets a reference to the given string and assigns it to the InbandIpGateway field.
func (o *NetworkElementAllOf) SetInbandIpGateway(v string) {
	o.InbandIpGateway = &v
}

// GetInbandIpMask returns the InbandIpMask field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetInbandIpMask() string {
	if o == nil || o.InbandIpMask == nil {
		var ret string
		return ret
	}
	return *o.InbandIpMask
}

// GetInbandIpMaskOk returns a tuple with the InbandIpMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetInbandIpMaskOk() (*string, bool) {
	if o == nil || o.InbandIpMask == nil {
		return nil, false
	}
	return o.InbandIpMask, true
}

// HasInbandIpMask returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasInbandIpMask() bool {
	if o != nil && o.InbandIpMask != nil {
		return true
	}

	return false
}

// SetInbandIpMask gets a reference to the given string and assigns it to the InbandIpMask field.
func (o *NetworkElementAllOf) SetInbandIpMask(v string) {
	o.InbandIpMask = &v
}

// GetInbandVlan returns the InbandVlan field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetInbandVlan() int64 {
	if o == nil || o.InbandVlan == nil {
		var ret int64
		return ret
	}
	return *o.InbandVlan
}

// GetInbandVlanOk returns a tuple with the InbandVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetInbandVlanOk() (*int64, bool) {
	if o == nil || o.InbandVlan == nil {
		return nil, false
	}
	return o.InbandVlan, true
}

// HasInbandVlan returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasInbandVlan() bool {
	if o != nil && o.InbandVlan != nil {
		return true
	}

	return false
}

// SetInbandVlan gets a reference to the given int64 and assigns it to the InbandVlan field.
func (o *NetworkElementAllOf) SetInbandVlan(v int64) {
	o.InbandVlan = &v
}

// GetOutOfBandIpAddress returns the OutOfBandIpAddress field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpAddress() string {
	if o == nil || o.OutOfBandIpAddress == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpAddress
}

// GetOutOfBandIpAddressOk returns a tuple with the OutOfBandIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpAddressOk() (*string, bool) {
	if o == nil || o.OutOfBandIpAddress == nil {
		return nil, false
	}
	return o.OutOfBandIpAddress, true
}

// HasOutOfBandIpAddress returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpAddress() bool {
	if o != nil && o.OutOfBandIpAddress != nil {
		return true
	}

	return false
}

// SetOutOfBandIpAddress gets a reference to the given string and assigns it to the OutOfBandIpAddress field.
func (o *NetworkElementAllOf) SetOutOfBandIpAddress(v string) {
	o.OutOfBandIpAddress = &v
}

// GetOutOfBandIpGateway returns the OutOfBandIpGateway field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpGateway() string {
	if o == nil || o.OutOfBandIpGateway == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpGateway
}

// GetOutOfBandIpGatewayOk returns a tuple with the OutOfBandIpGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpGatewayOk() (*string, bool) {
	if o == nil || o.OutOfBandIpGateway == nil {
		return nil, false
	}
	return o.OutOfBandIpGateway, true
}

// HasOutOfBandIpGateway returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpGateway() bool {
	if o != nil && o.OutOfBandIpGateway != nil {
		return true
	}

	return false
}

// SetOutOfBandIpGateway gets a reference to the given string and assigns it to the OutOfBandIpGateway field.
func (o *NetworkElementAllOf) SetOutOfBandIpGateway(v string) {
	o.OutOfBandIpGateway = &v
}

// GetOutOfBandIpMask returns the OutOfBandIpMask field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpMask() string {
	if o == nil || o.OutOfBandIpMask == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpMask
}

// GetOutOfBandIpMaskOk returns a tuple with the OutOfBandIpMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpMaskOk() (*string, bool) {
	if o == nil || o.OutOfBandIpMask == nil {
		return nil, false
	}
	return o.OutOfBandIpMask, true
}

// HasOutOfBandIpMask returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpMask() bool {
	if o != nil && o.OutOfBandIpMask != nil {
		return true
	}

	return false
}

// SetOutOfBandIpMask gets a reference to the given string and assigns it to the OutOfBandIpMask field.
func (o *NetworkElementAllOf) SetOutOfBandIpMask(v string) {
	o.OutOfBandIpMask = &v
}

// GetOutOfBandIpv4Address returns the OutOfBandIpv4Address field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpv4Address() string {
	if o == nil || o.OutOfBandIpv4Address == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpv4Address
}

// GetOutOfBandIpv4AddressOk returns a tuple with the OutOfBandIpv4Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpv4AddressOk() (*string, bool) {
	if o == nil || o.OutOfBandIpv4Address == nil {
		return nil, false
	}
	return o.OutOfBandIpv4Address, true
}

// HasOutOfBandIpv4Address returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpv4Address() bool {
	if o != nil && o.OutOfBandIpv4Address != nil {
		return true
	}

	return false
}

// SetOutOfBandIpv4Address gets a reference to the given string and assigns it to the OutOfBandIpv4Address field.
func (o *NetworkElementAllOf) SetOutOfBandIpv4Address(v string) {
	o.OutOfBandIpv4Address = &v
}

// GetOutOfBandIpv4Gateway returns the OutOfBandIpv4Gateway field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpv4Gateway() string {
	if o == nil || o.OutOfBandIpv4Gateway == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpv4Gateway
}

// GetOutOfBandIpv4GatewayOk returns a tuple with the OutOfBandIpv4Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpv4GatewayOk() (*string, bool) {
	if o == nil || o.OutOfBandIpv4Gateway == nil {
		return nil, false
	}
	return o.OutOfBandIpv4Gateway, true
}

// HasOutOfBandIpv4Gateway returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpv4Gateway() bool {
	if o != nil && o.OutOfBandIpv4Gateway != nil {
		return true
	}

	return false
}

// SetOutOfBandIpv4Gateway gets a reference to the given string and assigns it to the OutOfBandIpv4Gateway field.
func (o *NetworkElementAllOf) SetOutOfBandIpv4Gateway(v string) {
	o.OutOfBandIpv4Gateway = &v
}

// GetOutOfBandIpv4Mask returns the OutOfBandIpv4Mask field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpv4Mask() string {
	if o == nil || o.OutOfBandIpv4Mask == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpv4Mask
}

// GetOutOfBandIpv4MaskOk returns a tuple with the OutOfBandIpv4Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpv4MaskOk() (*string, bool) {
	if o == nil || o.OutOfBandIpv4Mask == nil {
		return nil, false
	}
	return o.OutOfBandIpv4Mask, true
}

// HasOutOfBandIpv4Mask returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpv4Mask() bool {
	if o != nil && o.OutOfBandIpv4Mask != nil {
		return true
	}

	return false
}

// SetOutOfBandIpv4Mask gets a reference to the given string and assigns it to the OutOfBandIpv4Mask field.
func (o *NetworkElementAllOf) SetOutOfBandIpv4Mask(v string) {
	o.OutOfBandIpv4Mask = &v
}

// GetOutOfBandIpv6Address returns the OutOfBandIpv6Address field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpv6Address() string {
	if o == nil || o.OutOfBandIpv6Address == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpv6Address
}

// GetOutOfBandIpv6AddressOk returns a tuple with the OutOfBandIpv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpv6AddressOk() (*string, bool) {
	if o == nil || o.OutOfBandIpv6Address == nil {
		return nil, false
	}
	return o.OutOfBandIpv6Address, true
}

// HasOutOfBandIpv6Address returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpv6Address() bool {
	if o != nil && o.OutOfBandIpv6Address != nil {
		return true
	}

	return false
}

// SetOutOfBandIpv6Address gets a reference to the given string and assigns it to the OutOfBandIpv6Address field.
func (o *NetworkElementAllOf) SetOutOfBandIpv6Address(v string) {
	o.OutOfBandIpv6Address = &v
}

// GetOutOfBandIpv6Gateway returns the OutOfBandIpv6Gateway field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpv6Gateway() string {
	if o == nil || o.OutOfBandIpv6Gateway == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpv6Gateway
}

// GetOutOfBandIpv6GatewayOk returns a tuple with the OutOfBandIpv6Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpv6GatewayOk() (*string, bool) {
	if o == nil || o.OutOfBandIpv6Gateway == nil {
		return nil, false
	}
	return o.OutOfBandIpv6Gateway, true
}

// HasOutOfBandIpv6Gateway returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpv6Gateway() bool {
	if o != nil && o.OutOfBandIpv6Gateway != nil {
		return true
	}

	return false
}

// SetOutOfBandIpv6Gateway gets a reference to the given string and assigns it to the OutOfBandIpv6Gateway field.
func (o *NetworkElementAllOf) SetOutOfBandIpv6Gateway(v string) {
	o.OutOfBandIpv6Gateway = &v
}

// GetOutOfBandIpv6Prefix returns the OutOfBandIpv6Prefix field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandIpv6Prefix() string {
	if o == nil || o.OutOfBandIpv6Prefix == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandIpv6Prefix
}

// GetOutOfBandIpv6PrefixOk returns a tuple with the OutOfBandIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandIpv6PrefixOk() (*string, bool) {
	if o == nil || o.OutOfBandIpv6Prefix == nil {
		return nil, false
	}
	return o.OutOfBandIpv6Prefix, true
}

// HasOutOfBandIpv6Prefix returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandIpv6Prefix() bool {
	if o != nil && o.OutOfBandIpv6Prefix != nil {
		return true
	}

	return false
}

// SetOutOfBandIpv6Prefix gets a reference to the given string and assigns it to the OutOfBandIpv6Prefix field.
func (o *NetworkElementAllOf) SetOutOfBandIpv6Prefix(v string) {
	o.OutOfBandIpv6Prefix = &v
}

// GetOutOfBandMac returns the OutOfBandMac field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetOutOfBandMac() string {
	if o == nil || o.OutOfBandMac == nil {
		var ret string
		return ret
	}
	return *o.OutOfBandMac
}

// GetOutOfBandMacOk returns a tuple with the OutOfBandMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetOutOfBandMacOk() (*string, bool) {
	if o == nil || o.OutOfBandMac == nil {
		return nil, false
	}
	return o.OutOfBandMac, true
}

// HasOutOfBandMac returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasOutOfBandMac() bool {
	if o != nil && o.OutOfBandMac != nil {
		return true
	}

	return false
}

// SetOutOfBandMac gets a reference to the given string and assigns it to the OutOfBandMac field.
func (o *NetworkElementAllOf) SetOutOfBandMac(v string) {
	o.OutOfBandMac = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetSwitchId() string {
	if o == nil || o.SwitchId == nil {
		var ret string
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetSwitchIdOk() (*string, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given string and assigns it to the SwitchId field.
func (o *NetworkElementAllOf) SetSwitchId(v string) {
	o.SwitchId = &v
}

// GetCards returns the Cards field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetCards() []EquipmentSwitchCardRelationship {
	if o == nil || o.Cards == nil {
		var ret []EquipmentSwitchCardRelationship
		return ret
	}
	return *o.Cards
}

// GetCardsOk returns a tuple with the Cards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetCardsOk() (*[]EquipmentSwitchCardRelationship, bool) {
	if o == nil || o.Cards == nil {
		return nil, false
	}
	return o.Cards, true
}

// HasCards returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasCards() bool {
	if o != nil && o.Cards != nil {
		return true
	}

	return false
}

// SetCards gets a reference to the given []EquipmentSwitchCardRelationship and assigns it to the Cards field.
func (o *NetworkElementAllOf) SetCards(v []EquipmentSwitchCardRelationship) {
	o.Cards = &v
}

// GetFanmodules returns the Fanmodules field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetFanmodules() []EquipmentFanModuleRelationship {
	if o == nil || o.Fanmodules == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return *o.Fanmodules
}

// GetFanmodulesOk returns a tuple with the Fanmodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool) {
	if o == nil || o.Fanmodules == nil {
		return nil, false
	}
	return o.Fanmodules, true
}

// HasFanmodules returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasFanmodules() bool {
	if o != nil && o.Fanmodules != nil {
		return true
	}

	return false
}

// SetFanmodules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the Fanmodules field.
func (o *NetworkElementAllOf) SetFanmodules(v []EquipmentFanModuleRelationship) {
	o.Fanmodules = &v
}

// GetManagementContoller returns the ManagementContoller field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetManagementContoller() ManagementControllerRelationship {
	if o == nil || o.ManagementContoller == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.ManagementContoller
}

// GetManagementContollerOk returns a tuple with the ManagementContoller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetManagementContollerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.ManagementContoller == nil {
		return nil, false
	}
	return o.ManagementContoller, true
}

// HasManagementContoller returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasManagementContoller() bool {
	if o != nil && o.ManagementContoller != nil {
		return true
	}

	return false
}

// SetManagementContoller gets a reference to the given ManagementControllerRelationship and assigns it to the ManagementContoller field.
func (o *NetworkElementAllOf) SetManagementContoller(v ManagementControllerRelationship) {
	o.ManagementContoller = &v
}

// GetManagementEntity returns the ManagementEntity field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetManagementEntity() ManagementEntityRelationship {
	if o == nil || o.ManagementEntity == nil {
		var ret ManagementEntityRelationship
		return ret
	}
	return *o.ManagementEntity
}

// GetManagementEntityOk returns a tuple with the ManagementEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetManagementEntityOk() (*ManagementEntityRelationship, bool) {
	if o == nil || o.ManagementEntity == nil {
		return nil, false
	}
	return o.ManagementEntity, true
}

// HasManagementEntity returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasManagementEntity() bool {
	if o != nil && o.ManagementEntity != nil {
		return true
	}

	return false
}

// SetManagementEntity gets a reference to the given ManagementEntityRelationship and assigns it to the ManagementEntity field.
func (o *NetworkElementAllOf) SetManagementEntity(v ManagementEntityRelationship) {
	o.ManagementEntity = &v
}

// GetPsus returns the Psus field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetPsus() []EquipmentPsuRelationship {
	if o == nil || o.Psus == nil {
		var ret []EquipmentPsuRelationship
		return ret
	}
	return *o.Psus
}

// GetPsusOk returns a tuple with the Psus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool) {
	if o == nil || o.Psus == nil {
		return nil, false
	}
	return o.Psus, true
}

// HasPsus returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasPsus() bool {
	if o != nil && o.Psus != nil {
		return true
	}

	return false
}

// SetPsus gets a reference to the given []EquipmentPsuRelationship and assigns it to the Psus field.
func (o *NetworkElementAllOf) SetPsus(v []EquipmentPsuRelationship) {
	o.Psus = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NetworkElementAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetTopSystem returns the TopSystem field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetTopSystem() TopSystemRelationship {
	if o == nil || o.TopSystem == nil {
		var ret TopSystemRelationship
		return ret
	}
	return *o.TopSystem
}

// GetTopSystemOk returns a tuple with the TopSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetTopSystemOk() (*TopSystemRelationship, bool) {
	if o == nil || o.TopSystem == nil {
		return nil, false
	}
	return o.TopSystem, true
}

// HasTopSystem returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasTopSystem() bool {
	if o != nil && o.TopSystem != nil {
		return true
	}

	return false
}

// SetTopSystem gets a reference to the given TopSystemRelationship and assigns it to the TopSystem field.
func (o *NetworkElementAllOf) SetTopSystem(v TopSystemRelationship) {
	o.TopSystem = &v
}

// GetUcsmRunningFirmware returns the UcsmRunningFirmware field value if set, zero value otherwise.
func (o *NetworkElementAllOf) GetUcsmRunningFirmware() FirmwareRunningFirmwareRelationship {
	if o == nil || o.UcsmRunningFirmware == nil {
		var ret FirmwareRunningFirmwareRelationship
		return ret
	}
	return *o.UcsmRunningFirmware
}

// GetUcsmRunningFirmwareOk returns a tuple with the UcsmRunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkElementAllOf) GetUcsmRunningFirmwareOk() (*FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.UcsmRunningFirmware == nil {
		return nil, false
	}
	return o.UcsmRunningFirmware, true
}

// HasUcsmRunningFirmware returns a boolean if a field has been set.
func (o *NetworkElementAllOf) HasUcsmRunningFirmware() bool {
	if o != nil && o.UcsmRunningFirmware != nil {
		return true
	}

	return false
}

// SetUcsmRunningFirmware gets a reference to the given FirmwareRunningFirmwareRelationship and assigns it to the UcsmRunningFirmware field.
func (o *NetworkElementAllOf) SetUcsmRunningFirmware(v FirmwareRunningFirmwareRelationship) {
	o.UcsmRunningFirmware = &v
}

func (o NetworkElementAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminInbandInterfaceState != nil {
		toSerialize["AdminInbandInterfaceState"] = o.AdminInbandInterfaceState
	}
	if o.FaultSummary != nil {
		toSerialize["FaultSummary"] = o.FaultSummary
	}
	if o.InbandIpAddress != nil {
		toSerialize["InbandIpAddress"] = o.InbandIpAddress
	}
	if o.InbandIpGateway != nil {
		toSerialize["InbandIpGateway"] = o.InbandIpGateway
	}
	if o.InbandIpMask != nil {
		toSerialize["InbandIpMask"] = o.InbandIpMask
	}
	if o.InbandVlan != nil {
		toSerialize["InbandVlan"] = o.InbandVlan
	}
	if o.OutOfBandIpAddress != nil {
		toSerialize["OutOfBandIpAddress"] = o.OutOfBandIpAddress
	}
	if o.OutOfBandIpGateway != nil {
		toSerialize["OutOfBandIpGateway"] = o.OutOfBandIpGateway
	}
	if o.OutOfBandIpMask != nil {
		toSerialize["OutOfBandIpMask"] = o.OutOfBandIpMask
	}
	if o.OutOfBandIpv4Address != nil {
		toSerialize["OutOfBandIpv4Address"] = o.OutOfBandIpv4Address
	}
	if o.OutOfBandIpv4Gateway != nil {
		toSerialize["OutOfBandIpv4Gateway"] = o.OutOfBandIpv4Gateway
	}
	if o.OutOfBandIpv4Mask != nil {
		toSerialize["OutOfBandIpv4Mask"] = o.OutOfBandIpv4Mask
	}
	if o.OutOfBandIpv6Address != nil {
		toSerialize["OutOfBandIpv6Address"] = o.OutOfBandIpv6Address
	}
	if o.OutOfBandIpv6Gateway != nil {
		toSerialize["OutOfBandIpv6Gateway"] = o.OutOfBandIpv6Gateway
	}
	if o.OutOfBandIpv6Prefix != nil {
		toSerialize["OutOfBandIpv6Prefix"] = o.OutOfBandIpv6Prefix
	}
	if o.OutOfBandMac != nil {
		toSerialize["OutOfBandMac"] = o.OutOfBandMac
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}
	if o.Cards != nil {
		toSerialize["Cards"] = o.Cards
	}
	if o.Fanmodules != nil {
		toSerialize["Fanmodules"] = o.Fanmodules
	}
	if o.ManagementContoller != nil {
		toSerialize["ManagementContoller"] = o.ManagementContoller
	}
	if o.ManagementEntity != nil {
		toSerialize["ManagementEntity"] = o.ManagementEntity
	}
	if o.Psus != nil {
		toSerialize["Psus"] = o.Psus
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.TopSystem != nil {
		toSerialize["TopSystem"] = o.TopSystem
	}
	if o.UcsmRunningFirmware != nil {
		toSerialize["UcsmRunningFirmware"] = o.UcsmRunningFirmware
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkElementAllOf struct {
	value *NetworkElementAllOf
	isSet bool
}

func (v NullableNetworkElementAllOf) Get() *NetworkElementAllOf {
	return v.value
}

func (v *NullableNetworkElementAllOf) Set(val *NetworkElementAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkElementAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkElementAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkElementAllOf(val *NetworkElementAllOf) *NullableNetworkElementAllOf {
	return &NullableNetworkElementAllOf{value: val, isSet: true}
}

func (v NullableNetworkElementAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkElementAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
