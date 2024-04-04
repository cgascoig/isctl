package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/icza/dyno"
	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"

	intersight "github.com/cgascoig/intersight-simple-go/intersight"
)

type ResultOpt struct {
	SingleResult *bool
}

type IsctlClient struct {
	IntersightClient *intersight.Client
	IntersightConfig intersight.Config
}

func ReadBody(bodyFormat string, bodyParamMap interface{}) error {
	if bodyFormat == "json" {
		// Gather body from JSON on stdin.
		fmt.Println("Waiting for JSON body: ")
		err := json.NewDecoder(os.Stdin).Decode(bodyParamMap)
		if err != nil {
			return fmt.Errorf("error decoding JSON: %v", err)
		}

		log.Tracef("After JSON parse, bodyParamMap: %v", bodyParamMap)
	} else if bodyFormat == "yaml" {
		// Gather body from YAML on stdin.
		fmt.Println("Waiting for YAML body: ")
		err := yaml.NewDecoder(os.Stdin).Decode(bodyParamMap)
		if err != nil {
			return fmt.Errorf("error decoding YAML: %v", err)
		}

		log.Tracef("After YAML parse, bodyParamMap: %v", bodyParamMap)
	} else {
		return fmt.Errorf("unknown request body format: %s", bodyFormat)
	}

	return nil
}

type ComplexValue struct {
	js string
}

func (v *ComplexValue) String() string {
	if v == nil {
		return ""
	}
	return v.js
}

func (v *ComplexValue) Set(s string) error {
	v.js = s
	return nil
}

func (v *ComplexValue) Type() string {
	return "json"
}

// getMoid takes a "<objecttype>.List" structure, checks there was exactly 1 match and returns the Moid of that match
func GetMoid(res interface{}) (string, bool) {
	results, err := dyno.GetSlice(res, "Results")
	if err != nil {
		log.Tracef("getMoid: no Results field (original res: %+v)", res)
		return "", false
	}

	if len(results) != 1 {
		log.Tracef("getMoid: number of results doesn't exactly equal 1")
		return "", false
	}

	moid, err := dyno.GetString(results, 0, "Moid")
	if err != nil {
		log.Tracef("getMoid: single result doesn't have Moid field")
		return "", false
	}

	return moid, true
}
