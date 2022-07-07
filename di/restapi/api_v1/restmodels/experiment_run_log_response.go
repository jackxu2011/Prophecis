// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ExperimentRunLogResponse experiment run log response
// swagger:model ExperimentRunLogResponse
type ExperimentRunLogResponse struct {

	// log from line number
	Fromline int64 `json:"fromline,omitempty"`

	// prophecis_experiment_run entity
	ID int64 `json:"id,omitempty"`

	// log
	Log []string `json:"log"`
}

// Validate validates this experiment run log response
func (m *ExperimentRunLogResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExperimentRunLogResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExperimentRunLogResponse) UnmarshalBinary(b []byte) error {
	var res ExperimentRunLogResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}