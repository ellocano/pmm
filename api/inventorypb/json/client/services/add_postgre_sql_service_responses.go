// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddPostgreSQLServiceReader is a Reader for the AddPostgreSQLService structure.
type AddPostgreSQLServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPostgreSQLServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPostgreSQLServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddPostgreSQLServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPostgreSQLServiceOK creates a AddPostgreSQLServiceOK with default headers values
func NewAddPostgreSQLServiceOK() *AddPostgreSQLServiceOK {
	return &AddPostgreSQLServiceOK{}
}

/*AddPostgreSQLServiceOK handles this case with default header values.

A successful response.
*/
type AddPostgreSQLServiceOK struct {
	Payload *AddPostgreSQLServiceOKBody
}

func (o *AddPostgreSQLServiceOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddPostgreSQL][%d] addPostgreSqlServiceOk  %+v", 200, o.Payload)
}

func (o *AddPostgreSQLServiceOK) GetPayload() *AddPostgreSQLServiceOKBody {
	return o.Payload
}

func (o *AddPostgreSQLServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLServiceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPostgreSQLServiceDefault creates a AddPostgreSQLServiceDefault with default headers values
func NewAddPostgreSQLServiceDefault(code int) *AddPostgreSQLServiceDefault {
	return &AddPostgreSQLServiceDefault{
		_statusCode: code,
	}
}

/*AddPostgreSQLServiceDefault handles this case with default header values.

An error response.
*/
type AddPostgreSQLServiceDefault struct {
	_statusCode int

	Payload *AddPostgreSQLServiceDefaultBody
}

// Code gets the status code for the add postgre SQL service default response
func (o *AddPostgreSQLServiceDefault) Code() int {
	return o._statusCode
}

func (o *AddPostgreSQLServiceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Services/AddPostgreSQL][%d] AddPostgreSQLService default  %+v", o._statusCode, o.Payload)
}

func (o *AddPostgreSQLServiceDefault) GetPayload() *AddPostgreSQLServiceDefaultBody {
	return o.Payload
}

func (o *AddPostgreSQLServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddPostgreSQLServiceBody add postgre SQL service body
swagger:model AddPostgreSQLServiceBody
*/
type AddPostgreSQLServiceBody struct {

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs. Required.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP). Required.
	Address string `json:"address,omitempty"`

	// Access port. Required.
	Port int64 `json:"port,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add postgre SQL service body
func (o *AddPostgreSQLServiceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLServiceDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddPostgreSQLServiceDefaultBody
*/
type AddPostgreSQLServiceDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add postgre SQL service default body
func (o *AddPostgreSQLServiceDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLServiceOKBody add postgre SQL service OK body
swagger:model AddPostgreSQLServiceOKBody
*/
type AddPostgreSQLServiceOKBody struct {

	// postgresql
	Postgresql *AddPostgreSQLServiceOKBodyPostgresql `json:"postgresql,omitempty"`
}

// Validate validates this add postgre SQL service OK body
func (o *AddPostgreSQLServiceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePostgresql(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLServiceOKBody) validatePostgresql(formats strfmt.Registry) error {

	if swag.IsZero(o.Postgresql) { // not required
		return nil
	}

	if o.Postgresql != nil {
		if err := o.Postgresql.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlServiceOk" + "." + "postgresql")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLServiceOKBodyPostgresql PostgreSQLService represents a generic PostgreSQL instance.
swagger:model AddPostgreSQLServiceOKBodyPostgresql
*/
type AddPostgreSQLServiceOKBodyPostgresql struct {

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`
}

// Validate validates this add postgre SQL service OK body postgresql
func (o *AddPostgreSQLServiceOKBodyPostgresql) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBodyPostgresql) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLServiceOKBodyPostgresql) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLServiceOKBodyPostgresql
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}