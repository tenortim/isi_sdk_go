// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventAlertConditionsAlertCondition event alert conditions alert condition
// swagger:model EventAlertConditionsAlert-Condition
type EventAlertConditionsAlertCondition struct {

	// Event Group categories to be alerted
	Categories []string `json:"categories"`

	// Channels for alert
	ChannelIds []int64 `json:"channel_ids"`

	// Trigger condition for alert
	// Enum: [NEW NEW EVENTS ONGOING SEVERITY INCREASE SEVERITY DECREASE RESOLVED]
	Condition string `json:"condition,omitempty"`

	// Event Group IDs to be alerted
	EventgroupIds []string `json:"eventgroup_ids"`

	// Unique identifier.
	ID string `json:"id,omitempty"`

	// Required with ONGOING condition only, period in seconds between alerts of ongoing conditions
	Interval int64 `json:"interval,omitempty"`

	// Required with NEW EVENTS condition only, limits the number of alerts sent as events are added
	Limit int64 `json:"limit,omitempty"`

	// Unique identifier.
	Name string `json:"name,omitempty"`

	// Severities to be alerted
	Severities []string `json:"severities"`

	// Any eventgroup lasting less than this many seconds is deemed transient and will not generate alerts under this condition.
	Transient int64 `json:"transient,omitempty"`
}

// Validate validates this event alert conditions alert condition
func (m *EventAlertConditionsAlertCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var eventAlertConditionsAlertConditionTypeConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEW","NEW EVENTS","ONGOING","SEVERITY INCREASE","SEVERITY DECREASE","RESOLVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventAlertConditionsAlertConditionTypeConditionPropEnum = append(eventAlertConditionsAlertConditionTypeConditionPropEnum, v)
	}
}

const (

	// EventAlertConditionsAlertConditionConditionNEW captures enum value "NEW"
	EventAlertConditionsAlertConditionConditionNEW string = "NEW"

	// EventAlertConditionsAlertConditionConditionNEWEVENTS captures enum value "NEW EVENTS"
	EventAlertConditionsAlertConditionConditionNEWEVENTS string = "NEW EVENTS"

	// EventAlertConditionsAlertConditionConditionONGOING captures enum value "ONGOING"
	EventAlertConditionsAlertConditionConditionONGOING string = "ONGOING"

	// EventAlertConditionsAlertConditionConditionSEVERITYINCREASE captures enum value "SEVERITY INCREASE"
	EventAlertConditionsAlertConditionConditionSEVERITYINCREASE string = "SEVERITY INCREASE"

	// EventAlertConditionsAlertConditionConditionSEVERITYDECREASE captures enum value "SEVERITY DECREASE"
	EventAlertConditionsAlertConditionConditionSEVERITYDECREASE string = "SEVERITY DECREASE"

	// EventAlertConditionsAlertConditionConditionRESOLVED captures enum value "RESOLVED"
	EventAlertConditionsAlertConditionConditionRESOLVED string = "RESOLVED"
)

// prop value enum
func (m *EventAlertConditionsAlertCondition) validateConditionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, eventAlertConditionsAlertConditionTypeConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EventAlertConditionsAlertCondition) validateCondition(formats strfmt.Registry) error {

	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	// value enum
	if err := m.validateConditionEnum("condition", "body", m.Condition); err != nil {
		return err
	}

	return nil
}

var eventAlertConditionsAlertConditionSeveritiesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["emergency","critical","warning","information"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventAlertConditionsAlertConditionSeveritiesItemsEnum = append(eventAlertConditionsAlertConditionSeveritiesItemsEnum, v)
	}
}

func (m *EventAlertConditionsAlertCondition) validateSeveritiesItemsEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, eventAlertConditionsAlertConditionSeveritiesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *EventAlertConditionsAlertCondition) validateSeverities(formats strfmt.Registry) error {

	if swag.IsZero(m.Severities) { // not required
		return nil
	}

	for i := 0; i < len(m.Severities); i++ {

		// value enum
		if err := m.validateSeveritiesItemsEnum("severities"+"."+strconv.Itoa(i), "body", m.Severities[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventAlertConditionsAlertCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventAlertConditionsAlertCondition) UnmarshalBinary(b []byte) error {
	var res EventAlertConditionsAlertCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
