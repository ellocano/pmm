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

// ChangeSettingsReader is a Reader for the ChangeSettings structure.
type ChangeSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeSettingsOK creates a ChangeSettingsOK with default headers values
func NewChangeSettingsOK() *ChangeSettingsOK {
	return &ChangeSettingsOK{}
}

/*ChangeSettingsOK handles this case with default header values.

A successful response.
*/
type ChangeSettingsOK struct {
	Payload *ChangeSettingsOKBody
}

func (o *ChangeSettingsOK) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Change][%d] changeSettingsOk  %+v", 200, o.Payload)
}

func (o *ChangeSettingsOK) GetPayload() *ChangeSettingsOKBody {
	return o.Payload
}

func (o *ChangeSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeSettingsDefault creates a ChangeSettingsDefault with default headers values
func NewChangeSettingsDefault(code int) *ChangeSettingsDefault {
	return &ChangeSettingsDefault{
		_statusCode: code,
	}
}

/*ChangeSettingsDefault handles this case with default header values.

An unexpected error response.
*/
type ChangeSettingsDefault struct {
	_statusCode int

	Payload *ChangeSettingsDefaultBody
}

// Code gets the status code for the change settings default response
func (o *ChangeSettingsDefault) Code() int {
	return o._statusCode
}

func (o *ChangeSettingsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/Settings/Change][%d] ChangeSettings default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeSettingsDefault) GetPayload() *ChangeSettingsDefaultBody {
	return o.Payload
}

func (o *ChangeSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeSettingsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeSettingsBody change settings body
swagger:model ChangeSettingsBody
*/
type ChangeSettingsBody struct {

	// enable telemetry
	EnableTelemetry bool `json:"enable_telemetry,omitempty"`

	// disable telemetry
	DisableTelemetry bool `json:"disable_telemetry,omitempty"`

	// A number of full days for Prometheus and QAN data retention. Should have a suffix in JSON: 2592000s, 43200m, 720h.
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// External AlertManager URL (e.g., https://username:password@1.2.3.4/path).
	AlertManagerURL string `json:"alert_manager_url,omitempty"`

	// Remove external AlertManager URL.
	RemoveAlertManagerURL bool `json:"remove_alert_manager_url,omitempty"`

	// Custom alerting or recording rules.
	AlertManagerRules string `json:"alert_manager_rules,omitempty"`

	// Remove custom alerting or recording rules.
	RemoveAlertManagerRules bool `json:"remove_alert_manager_rules,omitempty"`

	// Enable Security Threat Tool.
	EnableStt bool `json:"enable_stt,omitempty"`

	// Disable Security Threat Tool.
	DisableStt bool `json:"disable_stt,omitempty"`

	// Enable Integrated Alerting.
	EnableAlerting bool `json:"enable_alerting,omitempty"`

	// Disable Integrated Alerting.
	DisableAlerting bool `json:"disable_alerting,omitempty"`

	// If true, removes Integrated Alerting email (SMTP) settings.
	RemoveEmailAlertingSettings bool `json:"remove_email_alerting_settings,omitempty"`

	// If true, removes Integrated Alerting Slack settings.
	RemoveSlackAlertingSettings bool `json:"remove_slack_alerting_settings,omitempty"`

	// PMM Server public address.
	PMMPublicAddress string `json:"pmm_public_address,omitempty"`

	// remove pmm public address
	RemovePMMPublicAddress bool `json:"remove_pmm_public_address,omitempty"`

	// email alerting settings
	EmailAlertingSettings *ChangeSettingsParamsBodyEmailAlertingSettings `json:"email_alerting_settings,omitempty"`

	// metrics resolutions
	MetricsResolutions *ChangeSettingsParamsBodyMetricsResolutions `json:"metrics_resolutions,omitempty"`

	// slack alerting settings
	SlackAlertingSettings *ChangeSettingsParamsBodySlackAlertingSettings `json:"slack_alerting_settings,omitempty"`

	// stt check intervals
	SttCheckIntervals *ChangeSettingsParamsBodySttCheckIntervals `json:"stt_check_intervals,omitempty"`
}

// Validate validates this change settings body
func (o *ChangeSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlackAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSttCheckIntervals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsBody) validateEmailAlertingSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.EmailAlertingSettings) { // not required
		return nil
	}

	if o.EmailAlertingSettings != nil {
		if err := o.EmailAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "email_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsBody) validateMetricsResolutions(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsBody) validateSlackAlertingSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.SlackAlertingSettings) { // not required
		return nil
	}

	if o.SlackAlertingSettings != nil {
		if err := o.SlackAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "slack_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsBody) validateSttCheckIntervals(formats strfmt.Registry) error {

	if swag.IsZero(o.SttCheckIntervals) { // not required
		return nil
	}

	if o.SttCheckIntervals != nil {
		if err := o.SttCheckIntervals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "stt_check_intervals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsDefaultBody change settings default body
swagger:model ChangeSettingsDefaultBody
*/
type ChangeSettingsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this change settings default body
func (o *ChangeSettingsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ChangeSettings default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBody change settings OK body
swagger:model ChangeSettingsOKBody
*/
type ChangeSettingsOKBody struct {

	// settings
	Settings *ChangeSettingsOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this change settings OK body
func (o *ChangeSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBody) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBodySettings Settings represents PMM Server settings.
swagger:model ChangeSettingsOKBodySettings
*/
type ChangeSettingsOKBodySettings struct {

	// True if updates are disabled.
	UpdatesDisabled bool `json:"updates_disabled,omitempty"`

	// True if telemetry is enabled.
	TelemetryEnabled bool `json:"telemetry_enabled,omitempty"`

	// data retention
	DataRetention string `json:"data_retention,omitempty"`

	// ssh key
	SSHKey string `json:"ssh_key,omitempty"`

	// aws partitions
	AWSPartitions []string `json:"aws_partitions"`

	// External AlertManager URL (e.g., https://username:password@1.2.3.4/path).
	AlertManagerURL string `json:"alert_manager_url,omitempty"`

	// Custom alerting or recording rules.
	AlertManagerRules string `json:"alert_manager_rules,omitempty"`

	// True if Security Threat Tool is enabled.
	SttEnabled bool `json:"stt_enabled,omitempty"`

	// Percona Platform user's email, if this PMM instance is linked to the Platform.
	PlatformEmail string `json:"platform_email,omitempty"`

	// True if DBaaS is enabled.
	DbaasEnabled bool `json:"dbaas_enabled,omitempty"`

	// True if Integrated Alerting is enabled.
	AlertingEnabled bool `json:"alerting_enabled,omitempty"`

	// PMM Server public address.
	PMMPublicAddress string `json:"pmm_public_address,omitempty"`

	// email alerting settings
	EmailAlertingSettings *ChangeSettingsOKBodySettingsEmailAlertingSettings `json:"email_alerting_settings,omitempty"`

	// metrics resolutions
	MetricsResolutions *ChangeSettingsOKBodySettingsMetricsResolutions `json:"metrics_resolutions,omitempty"`

	// slack alerting settings
	SlackAlertingSettings *ChangeSettingsOKBodySettingsSlackAlertingSettings `json:"slack_alerting_settings,omitempty"`

	// stt check intervals
	SttCheckIntervals *ChangeSettingsOKBodySettingsSttCheckIntervals `json:"stt_check_intervals,omitempty"`
}

// Validate validates this change settings OK body settings
func (o *ChangeSettingsOKBodySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmailAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMetricsResolutions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlackAlertingSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSttCheckIntervals(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeSettingsOKBodySettings) validateEmailAlertingSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.EmailAlertingSettings) { // not required
		return nil
	}

	if o.EmailAlertingSettings != nil {
		if err := o.EmailAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "email_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsOKBodySettings) validateMetricsResolutions(formats strfmt.Registry) error {

	if swag.IsZero(o.MetricsResolutions) { // not required
		return nil
	}

	if o.MetricsResolutions != nil {
		if err := o.MetricsResolutions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "metrics_resolutions")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsOKBodySettings) validateSlackAlertingSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.SlackAlertingSettings) { // not required
		return nil
	}

	if o.SlackAlertingSettings != nil {
		if err := o.SlackAlertingSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "slack_alerting_settings")
			}
			return err
		}
	}

	return nil
}

func (o *ChangeSettingsOKBodySettings) validateSttCheckIntervals(formats strfmt.Registry) error {

	if swag.IsZero(o.SttCheckIntervals) { // not required
		return nil
	}

	if o.SttCheckIntervals != nil {
		if err := o.SttCheckIntervals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeSettingsOk" + "." + "settings" + "." + "stt_check_intervals")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettings) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBodySettingsEmailAlertingSettings EmailAlertingSettings represents email (SMTP) configuration for Integrated Alerting.
swagger:model ChangeSettingsOKBodySettingsEmailAlertingSettings
*/
type ChangeSettingsOKBodySettingsEmailAlertingSettings struct {

	// SMTP From header field.
	From string `json:"from,omitempty"`

	// SMTP host and port.
	Smarthost string `json:"smarthost,omitempty"`

	// Hostname to identify to the SMTP server.
	Hello string `json:"hello,omitempty"`

	// Auth using CRAM-MD5, LOGIN and PLAIN.
	Username string `json:"username,omitempty"`

	// Auth using LOGIN and PLAIN.
	Password string `json:"password,omitempty"`

	// Auth using PLAIN.
	Identity string `json:"identity,omitempty"`

	// Auth using CRAM-MD5.
	Secret string `json:"secret,omitempty"`
}

// Validate validates this change settings OK body settings email alerting settings
func (o *ChangeSettingsOKBodySettingsEmailAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsEmailAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsEmailAlertingSettings) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettingsEmailAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBodySettingsMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model ChangeSettingsOKBodySettingsMetricsResolutions
*/
type ChangeSettingsOKBodySettingsMetricsResolutions struct {

	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this change settings OK body settings metrics resolutions
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettingsMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBodySettingsSlackAlertingSettings SlackAlertingSettings represents Slack configuration for Integrated Alerting.
swagger:model ChangeSettingsOKBodySettingsSlackAlertingSettings
*/
type ChangeSettingsOKBodySettingsSlackAlertingSettings struct {

	// Slack API (webhook) URL.
	URL string `json:"url,omitempty"`
}

// Validate validates this change settings OK body settings slack alerting settings
func (o *ChangeSettingsOKBodySettingsSlackAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsSlackAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsSlackAlertingSettings) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettingsSlackAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsOKBodySettingsSttCheckIntervals STTCheckIntervals represents intervals between STT checks.
swagger:model ChangeSettingsOKBodySettingsSttCheckIntervals
*/
type ChangeSettingsOKBodySettingsSttCheckIntervals struct {

	// Default check interval
	DefaultInterval string `json:"default_interval,omitempty"`

	// Interval for rare check runs
	RareInterval string `json:"rare_interval,omitempty"`

	// Interval for frequent check runs
	FrequentInterval string `json:"frequent_interval,omitempty"`
}

// Validate validates this change settings OK body settings stt check intervals
func (o *ChangeSettingsOKBodySettingsSttCheckIntervals) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsSttCheckIntervals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsOKBodySettingsSttCheckIntervals) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsOKBodySettingsSttCheckIntervals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsParamsBodyEmailAlertingSettings EmailAlertingSettings represents email (SMTP) configuration for Integrated Alerting.
swagger:model ChangeSettingsParamsBodyEmailAlertingSettings
*/
type ChangeSettingsParamsBodyEmailAlertingSettings struct {

	// SMTP From header field.
	From string `json:"from,omitempty"`

	// SMTP host and port.
	Smarthost string `json:"smarthost,omitempty"`

	// Hostname to identify to the SMTP server.
	Hello string `json:"hello,omitempty"`

	// Auth using CRAM-MD5, LOGIN and PLAIN.
	Username string `json:"username,omitempty"`

	// Auth using LOGIN and PLAIN.
	Password string `json:"password,omitempty"`

	// Auth using PLAIN.
	Identity string `json:"identity,omitempty"`

	// Auth using CRAM-MD5.
	Secret string `json:"secret,omitempty"`
}

// Validate validates this change settings params body email alerting settings
func (o *ChangeSettingsParamsBodyEmailAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyEmailAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyEmailAlertingSettings) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodyEmailAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsParamsBodyMetricsResolutions MetricsResolutions represents Prometheus exporters metrics resolutions.
swagger:model ChangeSettingsParamsBodyMetricsResolutions
*/
type ChangeSettingsParamsBodyMetricsResolutions struct {

	// High resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Hr string `json:"hr,omitempty"`

	// Medium resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Mr string `json:"mr,omitempty"`

	// Low resolution. Should have a suffix in JSON: 1s, 1m, 1h.
	Lr string `json:"lr,omitempty"`
}

// Validate validates this change settings params body metrics resolutions
func (o *ChangeSettingsParamsBodyMetricsResolutions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyMetricsResolutions) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodyMetricsResolutions) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodyMetricsResolutions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsParamsBodySlackAlertingSettings SlackAlertingSettings represents Slack configuration for Integrated Alerting.
swagger:model ChangeSettingsParamsBodySlackAlertingSettings
*/
type ChangeSettingsParamsBodySlackAlertingSettings struct {

	// Slack API (webhook) URL.
	URL string `json:"url,omitempty"`
}

// Validate validates this change settings params body slack alerting settings
func (o *ChangeSettingsParamsBodySlackAlertingSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodySlackAlertingSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodySlackAlertingSettings) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodySlackAlertingSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeSettingsParamsBodySttCheckIntervals STTCheckIntervals represents intervals between STT checks.
swagger:model ChangeSettingsParamsBodySttCheckIntervals
*/
type ChangeSettingsParamsBodySttCheckIntervals struct {

	// Default check interval
	DefaultInterval string `json:"default_interval,omitempty"`

	// Interval for rare check runs
	RareInterval string `json:"rare_interval,omitempty"`

	// Interval for frequent check runs
	FrequentInterval string `json:"frequent_interval,omitempty"`
}

// Validate validates this change settings params body stt check intervals
func (o *ChangeSettingsParamsBodySttCheckIntervals) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeSettingsParamsBodySttCheckIntervals) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeSettingsParamsBodySttCheckIntervals) UnmarshalBinary(b []byte) error {
	var res ChangeSettingsParamsBodySttCheckIntervals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
