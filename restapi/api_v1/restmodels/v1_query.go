// Code generated by go-swagger; DO NOT EDIT.

package restmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1Query Playing with semi-generalized query request.
// swagger:model v1Query

type V1Query struct {

	// follow
	Follow bool `json:"follow,omitempty"`

	// The following three options are exclusive
	Meta *V1MetaInfo `json:"meta,omitempty"`

	// n records
	NRecords int32 `json:"nRecords,omitempty"`

	// search type
	SearchType QuerySearchType `json:"searchType,omitempty"`

	// since
	Since string `json:"since,omitempty"`
}

/* polymorph v1Query follow false */

/* polymorph v1Query meta false */

/* polymorph v1Query nRecords false */

/* polymorph v1Query searchType false */

/* polymorph v1Query since false */

// Validate validates this v1 query
func (m *V1Query) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSearchType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1Query) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {

		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *V1Query) validateSearchType(formats strfmt.Registry) error {

	if swag.IsZero(m.SearchType) { // not required
		return nil
	}

	if err := m.SearchType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("searchType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1Query) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Query) UnmarshalBinary(b []byte) error {
	var res V1Query
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
