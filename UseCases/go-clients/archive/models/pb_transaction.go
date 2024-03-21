// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PbTransaction pb transaction
//
// swagger:model pbTransaction
type PbTransaction struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// dest Id
	DestID string `json:"destId,omitempty"`

	// input hex
	InputHex string `json:"inputHex,omitempty"`

	// input size
	InputSize int64 `json:"inputSize,omitempty"`

	// input type
	InputType int64 `json:"inputType,omitempty"`

	// signature hex
	SignatureHex string `json:"signatureHex,omitempty"`

	// source Id
	SourceID string `json:"sourceId,omitempty"`

	// tick number
	TickNumber int64 `json:"tickNumber,omitempty"`

	// tx Id
	TxID string `json:"txId,omitempty"`
}

// Validate validates this pb transaction
func (m *PbTransaction) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pb transaction based on context it is used
func (m *PbTransaction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PbTransaction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PbTransaction) UnmarshalBinary(b []byte) error {
	var res PbTransaction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}