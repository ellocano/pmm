// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetSettingsReader is a Reader for the GetSettings structure.
type GetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSettingsOK creates a GetSettingsOK with default headers values
func NewGetSettingsOK() *GetSettingsOK {
	return &GetSettingsOK{}
}

/*GetSettingsOK handles this case with default header values.

A successful response.
*/
type GetSettingsOK struct {
	Payload *GetSettingsOKBody
}

func (o *GetSettingsOK) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Get][%d] getSettingsOk  %+v", 200, o.Payload)
}

func (o *GetSettingsOK) GetPayload() *GetSettingsOKBody {
	return o.Payload
}

func (o *GetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSettingsDefault creates a GetSettingsDefault with default headers values
func NewGetSettingsDefault(code int) *GetSettingsDefault {
	return &GetSettingsDefault{
		_statusCode: code,
	}
}

/*GetSettingsDefault handles this case with default header values.

An unexpected error response
*/
type GetSettingsDefault struct {
	_statusCode int

	Payload *GetSettingsDefaultBody
}

// Code gets the status code for the get settings default response
func (o *GetSettingsDefault) Code() int {
	return o._statusCode
}

func (o *GetSettingsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Get][%d] GetSettings default  %+v", o._statusCode, o.Payload)
}

func (o *GetSettingsDefault) GetPayload() *GetSettingsDefaultBody {
	return o.Payload
}

func (o *GetSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSettingsDefaultBody get settings default body
swagger:model GetSettingsDefaultBody
*/
type GetSettingsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this get settings default body
func (o *GetSettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetSettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBody get settings OK body
swagger:model GetSettingsOKBody
*/
type GetSettingsOKBody struct {

	// settings
	Settings *GetSettingsOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this get settings OK body
func (o *GetSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBody) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettings Settings represents PMM Server settings.
swagger:model GetSettingsOKBodySettings
*/
type GetSettingsOKBodySettings struct {

	// updates disabled
	UpdatesDisabled bool `json:"updates_disabled,omitempty"`

	// telemetry enabled
	TelemetryEnabled bool `json:"telemetry_enabled,omitempty"`

	// data retention
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// Prometheus AlertManager URL (e.g., https://username:password@1.2.3.4/path).
	AlertManagerURL string `json:"alert_manager_url,omitempty"`

	// alert manager rules
	AlertManagerRules string `json:"alert_manager_rules,omitempty"`

	// Security Threat Tool enabled
	SttEnabled bool `json:"stt_enabled,omitempty"`

	// Percona Platform user's email, if this PMM instance is linked to the Platform.
	PlatformEmail string `json:"platform_email,omitempty"`

	// metrics resolutions
	MetricsResolutions *GetSettingsOKBodySettingsMetricsResolutions `json:"metrics_resolutions,omitempty"`
}

// Validate validates this get settings OK body settings
func (o *GetSettingsOKBodySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSettingsOKBodySettings) validateMetricsResolutions(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettings) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSettingsOKBodySettingsMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model GetSettingsOKBodySettingsMetricsResolutions
*/
type GetSettingsOKBodySettingsMetricsResolutions struct {

	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this get settings OK body settings metrics resolutions
func (o *GetSettingsOKBodySettingsMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSettingsOKBodySettingsMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res GetSettingsOKBodySettingsMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
