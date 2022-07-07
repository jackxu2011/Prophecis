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

// ResultRequest result request
// swagger:model ResultRequest
type ResultRequest struct {

	// enable flag
	EnableFlag int8 `json:"enable_flag,omitempty"`

	// model id
	ModelID int64 `json:"model_id,omitempty"`

	// model version id
	ModelVersionID int64 `json:"model_version_id,omitempty"`

	// result msg
	// Required: true
	ResultMsg *string `json:"result_msg"`

	// training id
	// Required: true
	TrainingID *string `json:"training_id"`
}

// Validate validates this result request
func (m *ResultRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResultMsg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrainingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResultRequest) validateResultMsg(formats strfmt.Registry) error {

	if err := validate.Required("result_msg", "body", m.ResultMsg); err != nil {
		return err
	}

	return nil
}

func (m *ResultRequest) validateTrainingID(formats strfmt.Registry) error {

	if err := validate.Required("training_id", "body", m.TrainingID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResultRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResultRequest) UnmarshalBinary(b []byte) error {
	var res ResultRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}