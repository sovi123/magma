// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// FlowRecord flow record
// swagger:model flow_record
type FlowRecord struct {

	// bytes rx
	BytesRx uint64 `json:"bytes_rx,omitempty"`

	// bytes tx
	BytesTx uint64 `json:"bytes_tx,omitempty"`

	// match
	Match *FlowMatch `json:"match,omitempty"`

	// pkts rx
	PktsRx uint64 `json:"pkts_rx,omitempty"`

	// pkts tx
	PktsTx uint64 `json:"pkts_tx,omitempty"`

	// subscriber id
	SubscriberID SubscriberID `json:"subscriber_id,omitempty"`
}

// Validate validates this flow record
func (m *FlowRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriberID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlowRecord) validateMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if m.Match != nil {
		if err := m.Match.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match")
			}
			return err
		}
	}

	return nil
}

func (m *FlowRecord) validateSubscriberID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriberID) { // not required
		return nil
	}

	if err := m.SubscriberID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("subscriber_id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowRecord) UnmarshalBinary(b []byte) error {
	var res FlowRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
