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

// IsConcreteClass returns whether the given classId corresponds to a concrete class.
// Returns false if the class is abstract or not found in the metadata.
func (m *Meta) IsConcreteClass(classId string) bool {
	for _, cm := range *m {
		if cm.Name == classId {
			return cm.IsConcrete
		}
	}
	return false
}

// GetConcreteImplementations returns the names of concrete classes that have
// abstractClassId in their AncestorClasses slice. This traverses the full
// inheritance tree, not just direct children.
func (m *Meta) GetConcreteImplementations(abstractClassId string) []string {
	var result []string
	for _, cm := range *m {
		if !cm.IsConcrete {
			continue
		}
		for _, ancestor := range cm.AncestorClasses {
			if ancestor == abstractClassId {
				result = append(result, cm.Name)
				break
			}
		}
	}
	return result
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

// GetClassMeta returns the ClassMeta for a given class ID, or nil if not found.
func (m *Meta) GetClassMeta(classId string) *ClassMeta {
	for i := range *m {
		if (*m)[i].Name == classId {
			return &(*m)[i]
		}
	}
	return nil
}

// GetWritablePropertyNames returns the names of properties with ApiAccess == "ReadWrite".
func (m *Meta) GetWritablePropertyNames(classId string) []string {
	cm := m.GetClassMeta(classId)
	if cm == nil {
		return nil
	}
	var names []string
	for _, prop := range cm.Properties {
		if prop.ApiAccess == "ReadWrite" {
			names = append(names, prop.Name)
		}
	}
	return names
}

// GetWritableRelationshipNames returns the names of relationships with ApiAccess == "ReadWrite".
func (m *Meta) GetWritableRelationshipNames(classId string) []string {
	cm := m.GetClassMeta(classId)
	if cm == nil {
		return nil
	}
	var names []string
	for _, rel := range cm.Relationships {
		if rel.ApiAccess == "ReadWrite" {
			names = append(names, rel.Name)
		}
	}
	return names
}
