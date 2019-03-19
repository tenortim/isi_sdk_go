// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem filepool policy file matching pattern or criteria item and criteria item
// swagger:model FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem
type FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem struct {

	// Indicates whether the existence of an attribute indicates a match (valid only with 'type' = 'custom_attribute')
	AttributeExists bool `json:"attribute_exists,omitempty"`

	// True to match files recursively under the given path. (valid only with 'type' = 'path')
	BeginsWith bool `json:"begins_with,omitempty"`

	// True to indicate case sensitivity when comparing file attributes (valid only with 'type' = 'name' or 'type' = 'path')
	CaseSensitive bool `json:"case_sensitive,omitempty"`

	// File attribute field name to be compared in a custom comparison (valid only with 'type' = 'custom_attribute')
	Field string `json:"field,omitempty"`

	// The comparison operator to use while comparing an attribute with its value
	Operator string `json:"operator,omitempty"`

	// The file attribute to be compared to a given value
	// Required: true
	// Enum: [name path link_count accessed_time birth_time changed_time metadata_changed_time size file_type custom_attribute]
	Type *string `json:"type"`

	// Size unit value. One of 'B','KB','MB','GB','TB','PB','EB' (valid only with 'type' = 'size')
	Units string `json:"units,omitempty"`

	// Whether time units refer to a calendar date and time (e.g., Jun 3, 2009) or a relative duration (e.g., 2 weeks) (valid only with 'type' in {accessed_time, birth_time, changed_time or metadata_changed_time}
	UseRelativeTime bool `json:"use_relative_time,omitempty"`

	// The value to be compared against a file attribute
	Value string `json:"value,omitempty"`
}

// Validate validates this filepool policy file matching pattern or criteria item and criteria item
func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var filepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["name","path","link_count","accessed_time","birth_time","changed_time","metadata_changed_time","size","file_type","custom_attribute"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeTypePropEnum = append(filepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeTypePropEnum, v)
	}
}

const (

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeName captures enum value "name"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeName string = "name"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypePath captures enum value "path"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypePath string = "path"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeLinkCount captures enum value "link_count"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeLinkCount string = "link_count"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeAccessedTime captures enum value "accessed_time"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeAccessedTime string = "accessed_time"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeBirthTime captures enum value "birth_time"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeBirthTime string = "birth_time"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeChangedTime captures enum value "changed_time"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeChangedTime string = "changed_time"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeMetadataChangedTime captures enum value "metadata_changed_time"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeMetadataChangedTime string = "metadata_changed_time"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeSize captures enum value "size"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeSize string = "size"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeFileType captures enum value "file_type"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeFileType string = "file_type"

	// FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeCustomAttribute captures enum value "custom_attribute"
	FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeCustomAttribute string = "custom_attribute"
)

// prop value enum
func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, filepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItemTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem) UnmarshalBinary(b []byte) error {
	var res FilepoolPolicyFileMatchingPatternOrCriteriaItemAndCriteriaItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
