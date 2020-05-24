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
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// VirtualizationVmwareDatacenterRelationship - A relationship to the 'virtualization.VmwareDatacenter' resource, or the expanded 'virtualization.VmwareDatacenter' resource, or the 'null' value.
type VirtualizationVmwareDatacenterRelationship struct {
	MoMoRef                        *MoMoRef
	VirtualizationVmwareDatacenter *VirtualizationVmwareDatacenter
}

// MoMoRefAsVirtualizationVmwareDatacenterRelationship is a convenience function that returns MoMoRef wrapped in VirtualizationVmwareDatacenterRelationship
func MoMoRefAsVirtualizationVmwareDatacenterRelationship(v *MoMoRef) VirtualizationVmwareDatacenterRelationship {
	return VirtualizationVmwareDatacenterRelationship{MoMoRef: v}
}

// VirtualizationVmwareDatacenterAsVirtualizationVmwareDatacenterRelationship is a convenience function that returns VirtualizationVmwareDatacenter wrapped in VirtualizationVmwareDatacenterRelationship
func VirtualizationVmwareDatacenterAsVirtualizationVmwareDatacenterRelationship(v *VirtualizationVmwareDatacenter) VirtualizationVmwareDatacenterRelationship {
	return VirtualizationVmwareDatacenterRelationship{VirtualizationVmwareDatacenter: v}
}

// Unmarshl JSON data into one of the pointers in the struct
func (dst *VirtualizationVmwareDatacenterRelationship) UnmarshalJSON(data []byte) error {
	var err error

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["ClassId"]; ok {
		switch v {
		case "mo.MoRef":
			var result *MoMoRef = &MoMoRef{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.MoMoRef = result
			return nil
		case "virtualization.VmwareDatacenter":
			var result *VirtualizationVmwareDatacenter = &VirtualizationVmwareDatacenter{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.VirtualizationVmwareDatacenter = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'ClassId' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'ClassId' not found in unmarshaled payload: %+v", unmarshaled)
	}

}

// Marshl data from the first non-nil pointers in the struct to JSON
func (src *VirtualizationVmwareDatacenterRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.VirtualizationVmwareDatacenter != nil {
		return json.Marshal(&src.VirtualizationVmwareDatacenter)
	}

	return nil, nil // no data in oneOf schemas
}

// Marshl data from the first non-nil pointers in the struct to YAML
func (src *VirtualizationVmwareDatacenterRelationship) MarshalYAML() ([]byte, error) {
	if src.MoMoRef != nil {
		return yaml.Marshal(&src.MoMoRef)
	}

	if src.VirtualizationVmwareDatacenter != nil {
		return yaml.Marshal(&src.VirtualizationVmwareDatacenter)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VirtualizationVmwareDatacenterRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.VirtualizationVmwareDatacenter != nil {
		return obj.VirtualizationVmwareDatacenter
	}

	// all schemas are nil
	return nil
}
