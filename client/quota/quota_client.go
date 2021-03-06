// Code generated by go-swagger; DO NOT EDIT.

package quota

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new quota API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quota API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateQuotaQuota Create a new quota.
*/
func (a *Client) CreateQuotaQuota(params *CreateQuotaQuotaParams, authInfo runtime.ClientAuthInfoWriter) (*CreateQuotaQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQuotaQuotaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createQuotaQuota",
		Method:             "POST",
		PathPattern:        "/platform/1/quota/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateQuotaQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateQuotaQuotaOK), nil

}

/*
CreateQuotaReport Create a new report. The type of this report is 'manual'; it is also sometimes called 'live' or 'ad-hoc'.
*/
func (a *Client) CreateQuotaReport(params *CreateQuotaReportParams, authInfo runtime.ClientAuthInfoWriter) (*CreateQuotaReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateQuotaReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createQuotaReport",
		Method:             "POST",
		PathPattern:        "/platform/1/quota/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateQuotaReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateQuotaReportOK), nil

}

/*
CreateSettingsMapping Create a new rule. The new rule must not conflict with an existing rule (e.g. match both the type and domain fields).
*/
func (a *Client) CreateSettingsMapping(params *CreateSettingsMappingParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSettingsMappingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSettingsMappingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSettingsMapping",
		Method:             "POST",
		PathPattern:        "/platform/1/quota/settings/mappings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSettingsMappingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSettingsMappingOK), nil

}

/*
CreateSettingsNotification Create a new global notification rule.
*/
func (a *Client) CreateSettingsNotification(params *CreateSettingsNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSettingsNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSettingsNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSettingsNotification",
		Method:             "POST",
		PathPattern:        "/platform/1/quota/settings/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSettingsNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSettingsNotificationOK), nil

}

/*
DeleteQuotaQuota Delete the quota.
*/
func (a *Client) DeleteQuotaQuota(params *DeleteQuotaQuotaParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteQuotaQuotaNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteQuotaQuotaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteQuotaQuota",
		Method:             "DELETE",
		PathPattern:        "/platform/1/quota/quotas/{QuotaQuotaId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteQuotaQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteQuotaQuotaNoContent), nil

}

/*
DeleteQuotaQuotas Delete all or matching quotas.
*/
func (a *Client) DeleteQuotaQuotas(params *DeleteQuotaQuotasParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteQuotaQuotasNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteQuotaQuotasParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteQuotaQuotas",
		Method:             "DELETE",
		PathPattern:        "/platform/1/quota/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteQuotaQuotasReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteQuotaQuotasNoContent), nil

}

/*
DeleteQuotaReport Delete the report.
*/
func (a *Client) DeleteQuotaReport(params *DeleteQuotaReportParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteQuotaReportNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteQuotaReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteQuotaReport",
		Method:             "DELETE",
		PathPattern:        "/platform/1/quota/reports/{QuotaReportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteQuotaReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteQuotaReportNoContent), nil

}

/*
DeleteSettingsMapping Delete the mapping.
*/
func (a *Client) DeleteSettingsMapping(params *DeleteSettingsMappingParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSettingsMappingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSettingsMappingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSettingsMapping",
		Method:             "DELETE",
		PathPattern:        "/platform/1/quota/settings/mappings/{SettingsMappingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSettingsMappingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSettingsMappingNoContent), nil

}

/*
DeleteSettingsMappings Delete all rules.
*/
func (a *Client) DeleteSettingsMappings(params *DeleteSettingsMappingsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSettingsMappingsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSettingsMappingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSettingsMappings",
		Method:             "DELETE",
		PathPattern:        "/platform/1/quota/settings/mappings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSettingsMappingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSettingsMappingsNoContent), nil

}

/*
DeleteSettingsNotification Delete the notification rule.
*/
func (a *Client) DeleteSettingsNotification(params *DeleteSettingsNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSettingsNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSettingsNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSettingsNotification",
		Method:             "DELETE",
		PathPattern:        "/platform/1/quota/settings/notifications/{SettingsNotificationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSettingsNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSettingsNotificationNoContent), nil

}

/*
DeleteSettingsNotifications Delete all rules.
*/
func (a *Client) DeleteSettingsNotifications(params *DeleteSettingsNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSettingsNotificationsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSettingsNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSettingsNotifications",
		Method:             "DELETE",
		PathPattern:        "/platform/1/quota/settings/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSettingsNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSettingsNotificationsNoContent), nil

}

/*
GetQuotaLicense Retrieve license information.
*/
func (a *Client) GetQuotaLicense(params *GetQuotaLicenseParams, authInfo runtime.ClientAuthInfoWriter) (*GetQuotaLicenseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuotaLicenseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getQuotaLicense",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/license",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQuotaLicenseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetQuotaLicenseOK), nil

}

/*
GetQuotaQuota Retrieve quota information.
*/
func (a *Client) GetQuotaQuota(params *GetQuotaQuotaParams, authInfo runtime.ClientAuthInfoWriter) (*GetQuotaQuotaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuotaQuotaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getQuotaQuota",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/quotas/{QuotaQuotaId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQuotaQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetQuotaQuotaOK), nil

}

/*
GetQuotaQuotasSummary Return summary information about quotas.
*/
func (a *Client) GetQuotaQuotasSummary(params *GetQuotaQuotasSummaryParams, authInfo runtime.ClientAuthInfoWriter) (*GetQuotaQuotasSummaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuotaQuotasSummaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getQuotaQuotasSummary",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/quotas-summary",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQuotaQuotasSummaryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetQuotaQuotasSummaryOK), nil

}

/*
GetQuotaReport Retrieve report data (XML) or contents (meta-data).
*/
func (a *Client) GetQuotaReport(params *GetQuotaReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetQuotaReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuotaReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getQuotaReport",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/reports/{QuotaReportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetQuotaReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetQuotaReportOK), nil

}

/*
GetSettingsMapping Retrieve the mapping information.
*/
func (a *Client) GetSettingsMapping(params *GetSettingsMappingParams, authInfo runtime.ClientAuthInfoWriter) (*GetSettingsMappingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsMappingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSettingsMapping",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/settings/mappings/{SettingsMappingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSettingsMappingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSettingsMappingOK), nil

}

/*
GetSettingsNotification Retrieve notification rule information.
*/
func (a *Client) GetSettingsNotification(params *GetSettingsNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*GetSettingsNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSettingsNotification",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/settings/notifications/{SettingsNotificationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSettingsNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSettingsNotificationOK), nil

}

/*
GetSettingsReports List all settings.
*/
func (a *Client) GetSettingsReports(params *GetSettingsReportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSettingsReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsReportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSettingsReports",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/settings/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSettingsReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSettingsReportsOK), nil

}

/*
ListQuotaQuotas List all or matching quotas. Can also be used to retrieve quota state from existing reports. For any query argument not supplied, the default behavior is return all.
*/
func (a *Client) ListQuotaQuotas(params *ListQuotaQuotasParams, authInfo runtime.ClientAuthInfoWriter) (*ListQuotaQuotasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQuotaQuotasParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listQuotaQuotas",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListQuotaQuotasReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListQuotaQuotasOK), nil

}

/*
ListQuotaReports List all or matching reports.
*/
func (a *Client) ListQuotaReports(params *ListQuotaReportsParams, authInfo runtime.ClientAuthInfoWriter) (*ListQuotaReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListQuotaReportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listQuotaReports",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListQuotaReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListQuotaReportsOK), nil

}

/*
ListSettingsMappings List all rules.
*/
func (a *Client) ListSettingsMappings(params *ListSettingsMappingsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSettingsMappingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSettingsMappingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSettingsMappings",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/settings/mappings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSettingsMappingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSettingsMappingsOK), nil

}

/*
ListSettingsNotifications List all rules.
*/
func (a *Client) ListSettingsNotifications(params *ListSettingsNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSettingsNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSettingsNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSettingsNotifications",
		Method:             "GET",
		PathPattern:        "/platform/1/quota/settings/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSettingsNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSettingsNotificationsOK), nil

}

/*
UpdateQuotaQuota Modify quota. All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateQuotaQuota(params *UpdateQuotaQuotaParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateQuotaQuotaNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateQuotaQuotaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateQuotaQuota",
		Method:             "PUT",
		PathPattern:        "/platform/1/quota/quotas/{QuotaQuotaId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateQuotaQuotaReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateQuotaQuotaNoContent), nil

}

/*
UpdateSettingsMapping Modify the mapping.
*/
func (a *Client) UpdateSettingsMapping(params *UpdateSettingsMappingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSettingsMappingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSettingsMappingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSettingsMapping",
		Method:             "PUT",
		PathPattern:        "/platform/1/quota/settings/mappings/{SettingsMappingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSettingsMappingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSettingsMappingNoContent), nil

}

/*
UpdateSettingsNotification Modify notification rule. All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateSettingsNotification(params *UpdateSettingsNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSettingsNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSettingsNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSettingsNotification",
		Method:             "PUT",
		PathPattern:        "/platform/1/quota/settings/notifications/{SettingsNotificationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSettingsNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSettingsNotificationNoContent), nil

}

/*
UpdateSettingsReports Modify one or more settings.
*/
func (a *Client) UpdateSettingsReports(params *UpdateSettingsReportsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSettingsReportsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSettingsReportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSettingsReports",
		Method:             "PUT",
		PathPattern:        "/platform/1/quota/settings/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSettingsReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSettingsReportsNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
