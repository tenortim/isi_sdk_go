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

// SummaryDriveDriveItem summary drive drive item
// swagger:model SummaryDriveDriveItem
type SummaryDriveDriveItem struct {

	// The average operation latency.
	// Required: true
	AccessLatency *float64 `json:"access_latency"`

	// The rate of slow (long-latency) operations.
	// Required: true
	AccessSlow *float64 `json:"access_slow"`

	// The percentage of time the drive was busy.
	// Required: true
	Busy *float64 `json:"busy"`

	// The rate of bytes written.
	// Required: true
	BytesIn *float64 `json:"bytes_in"`

	// The rate of bytes read.
	// Required: true
	BytesOut *float64 `json:"bytes_out"`

	// Drive ID LNN:bay.
	// Required: true
	DriveID *string `json:"drive_id"`

	// The average time spent in the I/O scheduler.
	// Required: true
	IoschedLatency *float64 `json:"iosched_latency"`

	// The average length of the I/O scheduler queue.
	// Required: true
	IoschedQueue *float64 `json:"iosched_queue"`

	// Unix Epoch time in seconds of the request.
	// Required: true
	Time *int64 `json:"time"`

	// The type of the drive.
	// Required: true
	Type *string `json:"type"`

	// The percent of blocks used on the drive.
	// Required: true
	UsedBytesPercent *float64 `json:"used_bytes_percent"`

	// The number of inodes allocated on the drive.
	// Required: true
	UsedInodes *float64 `json:"used_inodes"`

	// The average size of write operations.
	// Required: true
	XferSizeIn *float64 `json:"xfer_size_in"`

	// The average size of read operations.
	// Required: true
	XferSizeOut *float64 `json:"xfer_size_out"`

	// The rate of write operations.
	// Required: true
	XfersIn *float64 `json:"xfers_in"`

	// The rate of read operations.
	// Required: true
	XfersOut *float64 `json:"xfers_out"`
}

// Validate validates this summary drive drive item
func (m *SummaryDriveDriveItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccessSlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBytesIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBytesOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDriveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoschedLatency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoschedQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedBytesPercent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsedInodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXferSizeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXferSizeOut(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXfersIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXfersOut(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SummaryDriveDriveItem) validateAccessLatency(formats strfmt.Registry) error {

	if err := validate.Required("access_latency", "body", m.AccessLatency); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateAccessSlow(formats strfmt.Registry) error {

	if err := validate.Required("access_slow", "body", m.AccessSlow); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateBusy(formats strfmt.Registry) error {

	if err := validate.Required("busy", "body", m.Busy); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateBytesIn(formats strfmt.Registry) error {

	if err := validate.Required("bytes_in", "body", m.BytesIn); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateBytesOut(formats strfmt.Registry) error {

	if err := validate.Required("bytes_out", "body", m.BytesOut); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateDriveID(formats strfmt.Registry) error {

	if err := validate.Required("drive_id", "body", m.DriveID); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateIoschedLatency(formats strfmt.Registry) error {

	if err := validate.Required("iosched_latency", "body", m.IoschedLatency); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateIoschedQueue(formats strfmt.Registry) error {

	if err := validate.Required("iosched_queue", "body", m.IoschedQueue); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateUsedBytesPercent(formats strfmt.Registry) error {

	if err := validate.Required("used_bytes_percent", "body", m.UsedBytesPercent); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateUsedInodes(formats strfmt.Registry) error {

	if err := validate.Required("used_inodes", "body", m.UsedInodes); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateXferSizeIn(formats strfmt.Registry) error {

	if err := validate.Required("xfer_size_in", "body", m.XferSizeIn); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateXferSizeOut(formats strfmt.Registry) error {

	if err := validate.Required("xfer_size_out", "body", m.XferSizeOut); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateXfersIn(formats strfmt.Registry) error {

	if err := validate.Required("xfers_in", "body", m.XfersIn); err != nil {
		return err
	}

	return nil
}

func (m *SummaryDriveDriveItem) validateXfersOut(formats strfmt.Registry) error {

	if err := validate.Required("xfers_out", "body", m.XfersOut); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SummaryDriveDriveItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SummaryDriveDriveItem) UnmarshalBinary(b []byte) error {
	var res SummaryDriveDriveItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
