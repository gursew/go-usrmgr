// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserPartial user partial
//
// an user partial object
//
// swagger:model user_partial
type UserPartial struct {

	// device_id
	//
	// device id of user
	// Example: 23244
	DeviceID int64 `json:"device_id,omitempty"`

	// phone
	//
	// phone number of user
	// Example: +919873323432
	// Max Length: 13
	// Min Length: 10
	Phone string `json:"phone,omitempty"`

	// tags
	//
	// tags of user
	// Example: ["tag1","tag2"]
	Tags []string `json:"tags"`
}

// Validate validates this user partial
func (m *UserPartial) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserPartial) validatePhone(formats strfmt.Registry) error {
	if swag.IsZero(m.Phone) { // not required
		return nil
	}

	if err := validate.MinLength("phone", "body", m.Phone, 10); err != nil {
		return err
	}

	if err := validate.MaxLength("phone", "body", m.Phone, 13); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user partial based on context it is used
func (m *UserPartial) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserPartial) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserPartial) UnmarshalBinary(b []byte) error {
	var res UserPartial
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
