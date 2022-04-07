// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserInfo user info
// swagger:model UserInfo
type UserInfo struct {

	// enable flag
	EnableFlag int8 `json:"enable_flag,omitempty"`

	// gid
	Gid int64 `json:"gid,omitempty"`

	// guid check
	GUIDCheck int8 `json:"guid_check,omitempty"`

	// id for the model.
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// remarks
	Remarks string `json:"remarks,omitempty"`

	// uid
	UID int64 `json:"uid,omitempty"`

	// user type
	UserType string `json:"user_type,omitempty"`
}

// Validate validates this user info
func (m *UserInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserInfo) UnmarshalBinary(b []byte) error {
	var res UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
