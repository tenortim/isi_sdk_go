// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClusterTimeNode cluster time node
// swagger:model ClusterTimeNode
type ClusterTimeNode struct {

	// Error message, if the HTTP status returned from this node was not 200.
	Error string `json:"error,omitempty"`

	// Node ID of the node reporting this information.
	// Maximum: 4.294967295e+09
	// Minimum: 0
	ID *int64 `json:"id,omitempty"`

	// Logical node number of the node reporting this information.
	// Maximum: 4.294967295e+09
	// Minimum: 0
	Lnn *int64 `json:"lnn,omitempty"`

	// Status of the HTTP response from this node if not 200.  If 200, this field does not appear.
	// Maximum: 4.294967295e+09
	// Minimum: 0
	Status *int64 `json:"status,omitempty"`

	// The current time on the cluster as a UNIX epoch (seconds since 1/1/1970), as reported by this node.
	// Maximum: 9.223372036854776e+18
	// Minimum: -9.223372036854776e+18
	Time *int64 `json:"time,omitempty"`
}

// Validate validates this cluster time node
func (m *ClusterTimeNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLnn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterTimeNode) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("id", "body", int64(*m.ID), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("id", "body", int64(*m.ID), 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTimeNode) validateLnn(formats strfmt.Registry) error {

	if swag.IsZero(m.Lnn) { // not required
		return nil
	}

	if err := validate.MinimumInt("lnn", "body", int64(*m.Lnn), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("lnn", "body", int64(*m.Lnn), 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTimeNode) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := validate.MinimumInt("status", "body", int64(*m.Status), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("status", "body", int64(*m.Status), 4.294967295e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *ClusterTimeNode) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.MinimumInt("time", "body", int64(*m.Time), -9.223372036854776e+18, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("time", "body", int64(*m.Time), 9.223372036854776e+18, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterTimeNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterTimeNode) UnmarshalBinary(b []byte) error {
	var res ClusterTimeNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
