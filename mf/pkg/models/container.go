// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Container container
// swagger:model Container
type Container struct {

	// Name for the container.
	ContainerName string `json:"container_name,omitempty"`

	// CPU.
	CPU int64 `json:"cpu,omitempty"`

	// The time that container fineshed to launch.
	FinishedTime string `json:"finished_time,omitempty"`

	// GPU.
	Gpu int64 `json:"gpu,omitempty"`

	// Image's name for the container.
	Image string `json:"image,omitempty"`

	// Image's Id for the container.
	ImageID string `json:"image_id,omitempty"`

	// Memory size.
	Memory string `json:"memory,omitempty"`

	// k8s namespaces.
	Namespace string `json:"namespace,omitempty"`

	// Name of pod that this container belong to.
	PodName string `json:"pod_name,omitempty"`

	// The time that container started to create.
	StartedTime string `json:"started_time,omitempty"`

	// Status of the container.
	Status string `json:"status,omitempty"`
}

// Validate validates this container
func (m *Container) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Container) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Container) UnmarshalBinary(b []byte) error {
	var res Container
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}