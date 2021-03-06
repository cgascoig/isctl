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
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// TelemetryDruidPostAggregator - struct for TelemetryDruidPostAggregator
type TelemetryDruidPostAggregator struct {
	TelemetryDruidArithmeticPostAggregator            *TelemetryDruidArithmeticPostAggregator
	TelemetryDruidConstantPostAggregator              *TelemetryDruidConstantPostAggregator
	TelemetryDruidFieldAccessorPostAggregator         *TelemetryDruidFieldAccessorPostAggregator
	TelemetryDruidGreatestLeastPostAggregator         *TelemetryDruidGreatestLeastPostAggregator
	TelemetryDruidHyperUniquePostAggregator           *TelemetryDruidHyperUniquePostAggregator
	TelemetryDruidThetaSketchEstimatePostAggregator   *TelemetryDruidThetaSketchEstimatePostAggregator
	TelemetryDruidThetaSketchOperationsPostAggregator *TelemetryDruidThetaSketchOperationsPostAggregator
}

// TelemetryDruidArithmeticPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidArithmeticPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidArithmeticPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidArithmeticPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidArithmeticPostAggregator: v}
}

// TelemetryDruidConstantPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidConstantPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidConstantPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidConstantPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidConstantPostAggregator: v}
}

// TelemetryDruidFieldAccessorPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidFieldAccessorPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidFieldAccessorPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidFieldAccessorPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidFieldAccessorPostAggregator: v}
}

// TelemetryDruidGreatestLeastPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidGreatestLeastPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidGreatestLeastPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidGreatestLeastPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidGreatestLeastPostAggregator: v}
}

// TelemetryDruidHyperUniquePostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidHyperUniquePostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidHyperUniquePostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidHyperUniquePostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidHyperUniquePostAggregator: v}
}

// TelemetryDruidThetaSketchEstimatePostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidThetaSketchEstimatePostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidThetaSketchEstimatePostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidThetaSketchEstimatePostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidThetaSketchEstimatePostAggregator: v}
}

// TelemetryDruidThetaSketchOperationsPostAggregatorAsTelemetryDruidPostAggregator is a convenience function that returns TelemetryDruidThetaSketchOperationsPostAggregator wrapped in TelemetryDruidPostAggregator
func TelemetryDruidThetaSketchOperationsPostAggregatorAsTelemetryDruidPostAggregator(v *TelemetryDruidThetaSketchOperationsPostAggregator) TelemetryDruidPostAggregator {
	return TelemetryDruidPostAggregator{TelemetryDruidThetaSketchOperationsPostAggregator: v}
}

// Unmarshl JSON data into one of the pointers in the struct
func (dst *TelemetryDruidPostAggregator) UnmarshalJSON(data []byte) error {
	var err error

	var unmarshaled map[string]interface{}
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		return err
	}
	if v, ok := unmarshaled["type"]; ok {
		switch v {
		case "arithmetic":
			var result *TelemetryDruidArithmeticPostAggregator = &TelemetryDruidArithmeticPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidArithmeticPostAggregator = result
			return nil
		case "constant":
			var result *TelemetryDruidConstantPostAggregator = &TelemetryDruidConstantPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidConstantPostAggregator = result
			return nil
		case "doubleGreatest":
			var result *TelemetryDruidGreatestLeastPostAggregator = &TelemetryDruidGreatestLeastPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidGreatestLeastPostAggregator = result
			return nil
		case "doubleLeast":
			var result *TelemetryDruidGreatestLeastPostAggregator = &TelemetryDruidGreatestLeastPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidGreatestLeastPostAggregator = result
			return nil
		case "fieldAccess":
			var result *TelemetryDruidFieldAccessorPostAggregator = &TelemetryDruidFieldAccessorPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidFieldAccessorPostAggregator = result
			return nil
		case "finalizingFieldAccess":
			var result *TelemetryDruidFieldAccessorPostAggregator = &TelemetryDruidFieldAccessorPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidFieldAccessorPostAggregator = result
			return nil
		case "hyperUniqueCardinality":
			var result *TelemetryDruidHyperUniquePostAggregator = &TelemetryDruidHyperUniquePostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHyperUniquePostAggregator = result
			return nil
		case "longGreatest":
			var result *TelemetryDruidGreatestLeastPostAggregator = &TelemetryDruidGreatestLeastPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidGreatestLeastPostAggregator = result
			return nil
		case "longLeast":
			var result *TelemetryDruidGreatestLeastPostAggregator = &TelemetryDruidGreatestLeastPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidGreatestLeastPostAggregator = result
			return nil
		case "telemetry.DruidArithmeticPostAggregator":
			var result *TelemetryDruidArithmeticPostAggregator = &TelemetryDruidArithmeticPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidArithmeticPostAggregator = result
			return nil
		case "telemetry.DruidConstantPostAggregator":
			var result *TelemetryDruidConstantPostAggregator = &TelemetryDruidConstantPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidConstantPostAggregator = result
			return nil
		case "telemetry.DruidFieldAccessorPostAggregator":
			var result *TelemetryDruidFieldAccessorPostAggregator = &TelemetryDruidFieldAccessorPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidFieldAccessorPostAggregator = result
			return nil
		case "telemetry.DruidGreatestLeastPostAggregator":
			var result *TelemetryDruidGreatestLeastPostAggregator = &TelemetryDruidGreatestLeastPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidGreatestLeastPostAggregator = result
			return nil
		case "telemetry.DruidHyperUniquePostAggregator":
			var result *TelemetryDruidHyperUniquePostAggregator = &TelemetryDruidHyperUniquePostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidHyperUniquePostAggregator = result
			return nil
		case "telemetry.DruidThetaSketchEstimatePostAggregator":
			var result *TelemetryDruidThetaSketchEstimatePostAggregator = &TelemetryDruidThetaSketchEstimatePostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidThetaSketchEstimatePostAggregator = result
			return nil
		case "telemetry.DruidThetaSketchOperationsPostAggregator":
			var result *TelemetryDruidThetaSketchOperationsPostAggregator = &TelemetryDruidThetaSketchOperationsPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidThetaSketchOperationsPostAggregator = result
			return nil
		case "thetaSketchEstimate":
			var result *TelemetryDruidThetaSketchEstimatePostAggregator = &TelemetryDruidThetaSketchEstimatePostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidThetaSketchEstimatePostAggregator = result
			return nil
		case "thetaSketchSetOp":
			var result *TelemetryDruidThetaSketchOperationsPostAggregator = &TelemetryDruidThetaSketchOperationsPostAggregator{}
			err = json.Unmarshal(data, result)
			if err != nil {
				return err
			}
			dst.TelemetryDruidThetaSketchOperationsPostAggregator = result
			return nil
		default:
			return fmt.Errorf("No oneOf model has 'type' equal to %s", v)
		}
	} else {
		return fmt.Errorf("Discriminator property 'type' not found in unmarshaled payload: %+v", unmarshaled)
	}

}

// Marshl data from the first non-nil pointers in the struct to JSON
func (src *TelemetryDruidPostAggregator) MarshalJSON() ([]byte, error) {
	if src.TelemetryDruidArithmeticPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidArithmeticPostAggregator)
	}

	if src.TelemetryDruidConstantPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidConstantPostAggregator)
	}

	if src.TelemetryDruidFieldAccessorPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidFieldAccessorPostAggregator)
	}

	if src.TelemetryDruidGreatestLeastPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidGreatestLeastPostAggregator)
	}

	if src.TelemetryDruidHyperUniquePostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidHyperUniquePostAggregator)
	}

	if src.TelemetryDruidThetaSketchEstimatePostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidThetaSketchEstimatePostAggregator)
	}

	if src.TelemetryDruidThetaSketchOperationsPostAggregator != nil {
		return json.Marshal(&src.TelemetryDruidThetaSketchOperationsPostAggregator)
	}

	return nil, nil // no data in oneOf schemas
}

// Marshl data from the first non-nil pointers in the struct to YAML
func (src *TelemetryDruidPostAggregator) MarshalYAML() ([]byte, error) {
	if src.TelemetryDruidArithmeticPostAggregator != nil {
		return yaml.Marshal(&src.TelemetryDruidArithmeticPostAggregator)
	}

	if src.TelemetryDruidConstantPostAggregator != nil {
		return yaml.Marshal(&src.TelemetryDruidConstantPostAggregator)
	}

	if src.TelemetryDruidFieldAccessorPostAggregator != nil {
		return yaml.Marshal(&src.TelemetryDruidFieldAccessorPostAggregator)
	}

	if src.TelemetryDruidGreatestLeastPostAggregator != nil {
		return yaml.Marshal(&src.TelemetryDruidGreatestLeastPostAggregator)
	}

	if src.TelemetryDruidHyperUniquePostAggregator != nil {
		return yaml.Marshal(&src.TelemetryDruidHyperUniquePostAggregator)
	}

	if src.TelemetryDruidThetaSketchEstimatePostAggregator != nil {
		return yaml.Marshal(&src.TelemetryDruidThetaSketchEstimatePostAggregator)
	}

	if src.TelemetryDruidThetaSketchOperationsPostAggregator != nil {
		return yaml.Marshal(&src.TelemetryDruidThetaSketchOperationsPostAggregator)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TelemetryDruidPostAggregator) GetActualInstance() interface{} {
	if obj.TelemetryDruidArithmeticPostAggregator != nil {
		return obj.TelemetryDruidArithmeticPostAggregator
	}

	if obj.TelemetryDruidConstantPostAggregator != nil {
		return obj.TelemetryDruidConstantPostAggregator
	}

	if obj.TelemetryDruidFieldAccessorPostAggregator != nil {
		return obj.TelemetryDruidFieldAccessorPostAggregator
	}

	if obj.TelemetryDruidGreatestLeastPostAggregator != nil {
		return obj.TelemetryDruidGreatestLeastPostAggregator
	}

	if obj.TelemetryDruidHyperUniquePostAggregator != nil {
		return obj.TelemetryDruidHyperUniquePostAggregator
	}

	if obj.TelemetryDruidThetaSketchEstimatePostAggregator != nil {
		return obj.TelemetryDruidThetaSketchEstimatePostAggregator
	}

	if obj.TelemetryDruidThetaSketchOperationsPostAggregator != nil {
		return obj.TelemetryDruidThetaSketchOperationsPostAggregator
	}

	// all schemas are nil
	return nil
}
