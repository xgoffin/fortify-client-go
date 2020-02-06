// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginMetaData plugin meta data holds attributes of the plugins installed in the application.
// swagger:model Plugin Meta Data
type PluginMetaData struct {

	// Version string of the SSC plugin api used to develop the plugin
	// Required: true
	APIVersion *string `json:"apiVersion"`

	// An integer used to tag the set of issue attributes provided by this plugin.
	// Required: true
	DataVersion *int32 `json:"dataVersion"`

	// Plugin description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Name of the scan engine supported by the plugin. Value is defined for parser plugins only
	// Read Only: true
	EngineType string `json:"engineType,omitempty"`

	// Plugin unique identifier
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Tracks whether this plugin instance was the most recently used of its kind
	// Required: true
	LastUsedOfKind *bool `json:"lastUsedOfKind"`

	// Additional configuration properties used by the plugin
	// Required: true
	PluginConfiguration []*PluginConfiguration `json:"pluginConfiguration"`

	// Identifier of the plugin, usually a fully-qualified classname. Non-unique when multiple versions of same plugin exist
	// Required: true
	PluginID *string `json:"pluginId"`

	// A string name for the plugin
	// Required: true
	PluginName *string `json:"pluginName"`

	// State of the plugin instance
	// Required: true
	// Enum: [STOPPED STARTING STARTED STOPPING FAILED_TO_START FAILED_TO_STOP UNKNOWN]
	PluginState *string `json:"pluginState"`

	// Denotes functionality of the plugin instance, such as scan parsing, bugtracker integration.
	// Required: true
	// Enum: [SCAN_PARSER BUG_TRACKER LEGACY_BUG_TRACKER]
	PluginType *string `json:"pluginType"`

	// A version string of the implementation code of the plugin
	// Required: true
	PluginVersion *string `json:"pluginVersion"`

	// Versions of the scan engine results supported by the plugin. Value is defined for parser plugins only
	// Read Only: true
	SupportedEngineVersions string `json:"supportedEngineVersions,omitempty"`

	// whether the plugin instance was installed by adding the jar to a special system folder
	// Required: true
	SystemInstalled *bool `json:"systemInstalled"`

	// Name of the company / organization that developed the plugin
	// Read Only: true
	VendorName string `json:"vendorName,omitempty"`

	// Plugin vendor's web site URL
	// Read Only: true
	VendorURL string `json:"vendorUrl,omitempty"`
}

// Validate validates this plugin meta data
func (m *PluginMetaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUsedOfKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePluginVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemInstalled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginMetaData) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("apiVersion", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *PluginMetaData) validateDataVersion(formats strfmt.Registry) error {

	if err := validate.Required("dataVersion", "body", m.DataVersion); err != nil {
		return err
	}

	return nil
}

func (m *PluginMetaData) validateLastUsedOfKind(formats strfmt.Registry) error {

	if err := validate.Required("lastUsedOfKind", "body", m.LastUsedOfKind); err != nil {
		return err
	}

	return nil
}

func (m *PluginMetaData) validatePluginConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("pluginConfiguration", "body", m.PluginConfiguration); err != nil {
		return err
	}

	for i := 0; i < len(m.PluginConfiguration); i++ {
		if swag.IsZero(m.PluginConfiguration[i]) { // not required
			continue
		}

		if m.PluginConfiguration[i] != nil {
			if err := m.PluginConfiguration[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pluginConfiguration" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginMetaData) validatePluginID(formats strfmt.Registry) error {

	if err := validate.Required("pluginId", "body", m.PluginID); err != nil {
		return err
	}

	return nil
}

func (m *PluginMetaData) validatePluginName(formats strfmt.Registry) error {

	if err := validate.Required("pluginName", "body", m.PluginName); err != nil {
		return err
	}

	return nil
}

var pluginMetaDataTypePluginStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STOPPED","STARTING","STARTED","STOPPING","FAILED_TO_START","FAILED_TO_STOP","UNKNOWN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pluginMetaDataTypePluginStatePropEnum = append(pluginMetaDataTypePluginStatePropEnum, v)
	}
}

const (

	// PluginMetaDataPluginStateSTOPPED captures enum value "STOPPED"
	PluginMetaDataPluginStateSTOPPED string = "STOPPED"

	// PluginMetaDataPluginStateSTARTING captures enum value "STARTING"
	PluginMetaDataPluginStateSTARTING string = "STARTING"

	// PluginMetaDataPluginStateSTARTED captures enum value "STARTED"
	PluginMetaDataPluginStateSTARTED string = "STARTED"

	// PluginMetaDataPluginStateSTOPPING captures enum value "STOPPING"
	PluginMetaDataPluginStateSTOPPING string = "STOPPING"

	// PluginMetaDataPluginStateFAILEDTOSTART captures enum value "FAILED_TO_START"
	PluginMetaDataPluginStateFAILEDTOSTART string = "FAILED_TO_START"

	// PluginMetaDataPluginStateFAILEDTOSTOP captures enum value "FAILED_TO_STOP"
	PluginMetaDataPluginStateFAILEDTOSTOP string = "FAILED_TO_STOP"

	// PluginMetaDataPluginStateUNKNOWN captures enum value "UNKNOWN"
	PluginMetaDataPluginStateUNKNOWN string = "UNKNOWN"
)

// prop value enum
func (m *PluginMetaData) validatePluginStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pluginMetaDataTypePluginStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PluginMetaData) validatePluginState(formats strfmt.Registry) error {

	if err := validate.Required("pluginState", "body", m.PluginState); err != nil {
		return err
	}

	// value enum
	if err := m.validatePluginStateEnum("pluginState", "body", *m.PluginState); err != nil {
		return err
	}

	return nil
}

var pluginMetaDataTypePluginTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SCAN_PARSER","BUG_TRACKER","LEGACY_BUG_TRACKER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pluginMetaDataTypePluginTypePropEnum = append(pluginMetaDataTypePluginTypePropEnum, v)
	}
}

const (

	// PluginMetaDataPluginTypeSCANPARSER captures enum value "SCAN_PARSER"
	PluginMetaDataPluginTypeSCANPARSER string = "SCAN_PARSER"

	// PluginMetaDataPluginTypeBUGTRACKER captures enum value "BUG_TRACKER"
	PluginMetaDataPluginTypeBUGTRACKER string = "BUG_TRACKER"

	// PluginMetaDataPluginTypeLEGACYBUGTRACKER captures enum value "LEGACY_BUG_TRACKER"
	PluginMetaDataPluginTypeLEGACYBUGTRACKER string = "LEGACY_BUG_TRACKER"
)

// prop value enum
func (m *PluginMetaData) validatePluginTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, pluginMetaDataTypePluginTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PluginMetaData) validatePluginType(formats strfmt.Registry) error {

	if err := validate.Required("pluginType", "body", m.PluginType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePluginTypeEnum("pluginType", "body", *m.PluginType); err != nil {
		return err
	}

	return nil
}

func (m *PluginMetaData) validatePluginVersion(formats strfmt.Registry) error {

	if err := validate.Required("pluginVersion", "body", m.PluginVersion); err != nil {
		return err
	}

	return nil
}

func (m *PluginMetaData) validateSystemInstalled(formats strfmt.Registry) error {

	if err := validate.Required("systemInstalled", "body", m.SystemInstalled); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginMetaData) UnmarshalBinary(b []byte) error {
	var res PluginMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
