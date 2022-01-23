// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cache Cache
//
// HAPRoxy Cache section
//
// swagger:model cache
type Cache struct {

	// max age
	MaxAge int64 `json:"max_age,omitempty"`

	// max object size
	MaxObjectSize int64 `json:"max_object_size,omitempty"`

	// max secondary entries
	MaxSecondaryEntries int64 `json:"max_secondary_entries,omitempty"`

	// name
	// Required: true
	// Pattern: ^[A-Za-z0-9-_.:]+$
	Name *string `json:"name"`

	// process vary
	ProcessVary *bool `json:"process_vary,omitempty"`

	// total max size
	// Maximum: 4095
	// Minimum: 1
	TotalMaxSize int64 `json:"total_max_size,omitempty"`
}

// Validate validates this cache
func (m *Cache) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalMaxSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cache) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[A-Za-z0-9-_.:]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Cache) validateTotalMaxSize(formats strfmt.Registry) error {

	if swag.IsZero(m.TotalMaxSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("total_max_size", "body", int64(m.TotalMaxSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("total_max_size", "body", int64(m.TotalMaxSize), 4095, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cache) UnmarshalBinary(b []byte) error {
	var res Cache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
