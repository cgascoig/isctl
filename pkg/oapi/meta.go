package oapi

import (
	_ "embed"
	"encoding/json"
	"sync"
)

type Meta []ClassMeta

type ClassMeta struct {
	AccessPrivileges    []interface{}      `json:"AccessPrivileges"`
	AccountMoid         string             `json:"AccountMoid"`
	AncestorClasses     []string           `json:"AncestorClasses"`
	Ancestors           []interface{}      `json:"Ancestors"`
	ClassId             string             `json:"ClassId"`
	CreateTime          string             `json:"CreateTime"`
	DisplayNameMetas    []interface{}      `json:"DisplayNameMetas"`
	DomainGroupMoid     string             `json:"DomainGroupMoid"`
	IdentityConstraints []interface{}      `json:"IdentityConstraints"`
	IsConcrete          bool               `json:"IsConcrete"`
	MetaType            string             `json:"MetaType"`
	ModTime             string             `json:"ModTime"`
	Moid                string             `json:"Moid"`
	Name                string             `json:"Name"`
	Namespace           string             `json:"Namespace"`
	ObjectType          string             `json:"ObjectType"`
	Owner               string             `json:"Owner"`
	Owners              []string           `json:"Owners"`
	ParentClass         string             `json:"ParentClass"`
	PermissionResources []interface{}      `json:"PermissionResources"`
	PermissionSupported bool               `json:"PermissionSupported"`
	Properties          []PropertyMeta     `json:"Properties"`
	RbacResource        bool               `json:"RbacResource"`
	Relationships       []RelationshipMeta `json:"Relationships"`
	ResourcePoolTypes   []interface{}      `json:"ResourcePoolTypes"`
	RestPath            string             `json:"RestPath"`
	SharedScope         string             `json:"SharedScope"`
	Tags                []interface{}      `json:"Tags"`
	Version             string             `json:"Version"`
}

type PropertyMeta struct {
	ApiAccess     string  `json:"ApiAccess"`
	ClassId       string  `json:"ClassId"`
	Default       any     `json:"Default"`
	IsCollection  bool    `json:"IsCollection"`
	IsComplexType bool    `json:"IsComplexType"`
	Kind          string  `json:"Kind"`
	Name          string  `json:"Name"`
	ObjectType    string  `json:"ObjectType"`
	OpSecurity    string  `json:"OpSecurity"`
	SearchWeight  float32 `json:"SearchWeight"`
	Type          string  `json:"Type"`
}

type RelationshipMeta struct {
	ApiAccess                string   `json:"ApiAccess"`
	ClassId                  string   `json:"ClassId"`
	Collection               bool     `json:"Collection"`
	Export                   bool     `json:"Export"`
	ExportWithPeer           bool     `json:"ExportWithPeer"`
	Name                     string   `json:"Name"`
	ObjectType               string   `json:"ObjectType"`
	PeerRelName              string   `json:"PeerRelName"`
	PeerSupportedObjectTypes []string `json:"PeerSupportedObjectTypes"`
	PeerSync                 bool     `json:"PeerSync"`
	Type                     string   `json:"Type"`
}

//go:embed "meta.json"
var metaData []byte

var (
	meta      *Meta
	metaMutex sync.Once
)

// GetMeta lazily loads and returns the metadata.
func GetMeta() (*Meta, error) {
	if meta != nil {
		return meta, nil
	}

	var err error
	metaMutex.Do(func() {
		var m Meta
		err = json.Unmarshal(metaData, &m)
		if err == nil {
			meta = &m
		}
	})
	return meta, err
}

// GetIdentityConstraints returns the identity constraint fields for a given class ID.
func (m *Meta) GetIdentityConstraints(classId string) []string {
	for _, cm := range *m {
		if cm.Name == classId {
			if len(cm.IdentityConstraints) == 0 {
				return []string{}
			}
			ic, ok := cm.IdentityConstraints[0].(map[string]interface{})
			if !ok {
				return []string{}
			}
			fields, ok := ic["Fields"].([]interface{})
			if !ok {
				return []string{}
			}
			res := make([]string, 0, len(fields))
			for _, f := range fields {
				if f, ok := f.(string); ok {
					res = append(res, f)
				}
			}
			return res
		}
	}
	return []string{}
}

// GetRefType returns the type of the relationship if the property is a reference, and true.
// Otherwise returns empty string and false.
func (m *Meta) GetRefType(classId, propName string) (string, bool) {
	for _, cm := range *m {
		if cm.Name == classId {
			for _, rel := range cm.Relationships {
				if rel.Name == propName {
					return rel.Type, true
				}
			}
			return "", false
		}
	}
	return "", false
}
