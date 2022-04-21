// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Account account
//
// swagger:model account
type Account struct {

	// Created At
	//
	// Timestamp when the user was created
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// account_id
	//
	// This is the account id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// account_name
	//
	// This is the account name
	// Example: ABC
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// source
	//
	// This is the source
	// Enum: [openstack openshift aws]
	Source string `json:"source,omitempty"`

	// Updated At
	//
	// Timestamp when the user updated something
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// user_id
	//
	// This is the id of user
	// Example: 1
	// Required: true
	UserID *int32 `json:"user_id"`

	// user_name
	//
	// This is the name of user
	// Example: ABC
	// Read Only: true
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	return nil
}

var accountTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["openstack","openshift","aws"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountTypeSourcePropEnum = append(accountTypeSourcePropEnum, v)
	}
}

const (

	// AccountSourceOpenstack captures enum value "openstack"
	AccountSourceOpenstack string = "openstack"

	// AccountSourceOpenshift captures enum value "openshift"
	AccountSourceOpenshift string = "openshift"

	// AccountSourceAws captures enum value "aws"
	AccountSourceAws string = "aws"
)

// prop value enum
func (m *Account) validateSourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountTypeSourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this account based on the context it is used
func (m *Account) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdatedAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) contextValidateCreatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created_at", "body", strfmt.DateTime(m.CreatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Account) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Account) contextValidateUpdatedAt(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "updated_at", "body", strfmt.DateTime(m.UpdatedAt)); err != nil {
		return err
	}

	return nil
}

func (m *Account) contextValidateUserName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "user_name", "body", string(m.UserName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Account) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Account) UnmarshalBinary(b []byte) error {
	var res Account
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}