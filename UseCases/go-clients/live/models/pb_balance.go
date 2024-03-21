// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PbBalance pb balance
//
// swagger:model pbBalance
type PbBalance struct {

	// balance
	Balance string `json:"balance,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// latest incoming transfer tick
	LatestIncomingTransferTick int64 `json:"latestIncomingTransferTick,omitempty"`

	// latest outgoing transfer tick
	LatestOutgoingTransferTick int64 `json:"latestOutgoingTransferTick,omitempty"`

	// valid for tick
	ValidForTick int64 `json:"validForTick,omitempty"`
}

// Validate validates this pb balance
func (m *PbBalance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pb balance based on context it is used
func (m *PbBalance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PbBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbBalance) UnmarshalBinary(b []byte) error {
	var res PbBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}