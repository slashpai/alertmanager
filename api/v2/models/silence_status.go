// Code generated by go-swagger; DO NOT EDIT.

// Copyright Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// SilenceStatus silence status
//
// swagger:model silenceStatus
type SilenceStatus struct {

	// state
	// Required: true
	// Enum: [expired active pending]
	State *string `json:"state"`
}

// Validate validates this silence status
func (m *SilenceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var silenceStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["expired","active","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		silenceStatusTypeStatePropEnum = append(silenceStatusTypeStatePropEnum, v)
	}
}

const (

	// SilenceStatusStateExpired captures enum value "expired"
	SilenceStatusStateExpired string = "expired"

	// SilenceStatusStateActive captures enum value "active"
	SilenceStatusStateActive string = "active"

	// SilenceStatusStatePending captures enum value "pending"
	SilenceStatusStatePending string = "pending"
)

// prop value enum
func (m *SilenceStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, silenceStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SilenceStatus) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this silence status based on context it is used
func (m *SilenceStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SilenceStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SilenceStatus) UnmarshalBinary(b []byte) error {
	var res SilenceStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
