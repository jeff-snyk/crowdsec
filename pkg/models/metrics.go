// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Metrics Metrics
//
// swagger:model Metrics
type Metrics struct {

	// the local version of crowdsec/apil
	// Required: true
	ApilVersion *string `json:"apil_version"`

	// bouncers
	// Required: true
	Bouncers []*MetricsSoftInfo `json:"bouncers"`

	// machines
	// Required: true
	Machines []*MetricsSoftInfo `json:"machines"`
}

// Validate validates this metrics
func (m *Metrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApilVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBouncers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMachines(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metrics) validateApilVersion(formats strfmt.Registry) error {

	if err := validate.Required("apil_version", "body", m.ApilVersion); err != nil {
		return err
	}

	return nil
}

func (m *Metrics) validateBouncers(formats strfmt.Registry) error {

	if err := validate.Required("bouncers", "body", m.Bouncers); err != nil {
		return err
	}

	for i := 0; i < len(m.Bouncers); i++ {
		if swag.IsZero(m.Bouncers[i]) { // not required
			continue
		}

		if m.Bouncers[i] != nil {
			if err := m.Bouncers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bouncers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Metrics) validateMachines(formats strfmt.Registry) error {

	if err := validate.Required("machines", "body", m.Machines); err != nil {
		return err
	}

	for i := 0; i < len(m.Machines); i++ {
		if swag.IsZero(m.Machines[i]) { // not required
			continue
		}

		if m.Machines[i] != nil {
			if err := m.Machines[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machines" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Metrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metrics) UnmarshalBinary(b []byte) error {
	var res Metrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
