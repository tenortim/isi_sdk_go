// Code generated by go-swagger; DO NOT EDIT.

package sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sync API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sync API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSyncJob Start a SyncIQ job.
*/
func (a *Client) CreateSyncJob(params *CreateSyncJobParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSyncJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSyncJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSyncJob",
		Method:             "POST",
		PathPattern:        "/platform/3/sync/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSyncJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSyncJobOK), nil

}

/*
CreateSyncPolicy Create a SyncIQ policy.
*/
func (a *Client) CreateSyncPolicy(params *CreateSyncPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSyncPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSyncPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSyncPolicy",
		Method:             "POST",
		PathPattern:        "/platform/3/sync/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSyncPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSyncPolicyOK), nil

}

/*
CreateSyncReportsRotateItem Rotate the records in the database(s).
*/
func (a *Client) CreateSyncReportsRotateItem(params *CreateSyncReportsRotateItemParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSyncReportsRotateItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSyncReportsRotateItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSyncReportsRotateItem",
		Method:             "POST",
		PathPattern:        "/platform/1/sync/reports-rotate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSyncReportsRotateItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSyncReportsRotateItemOK), nil

}

/*
CreateSyncRule Create a new SyncIQ performance rule.
*/
func (a *Client) CreateSyncRule(params *CreateSyncRuleParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSyncRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSyncRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSyncRule",
		Method:             "POST",
		PathPattern:        "/platform/3/sync/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSyncRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSyncRuleOK), nil

}

/*
DeleteSyncPolicies Delete all SyncIQ policies.
*/
func (a *Client) DeleteSyncPolicies(params *DeleteSyncPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSyncPoliciesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSyncPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSyncPolicies",
		Method:             "DELETE",
		PathPattern:        "/platform/3/sync/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSyncPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSyncPoliciesNoContent), nil

}

/*
DeleteSyncPolicy Delete a single SyncIQ policy.
*/
func (a *Client) DeleteSyncPolicy(params *DeleteSyncPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSyncPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSyncPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSyncPolicy",
		Method:             "DELETE",
		PathPattern:        "/platform/3/sync/policies/{SyncPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSyncPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSyncPolicyNoContent), nil

}

/*
DeleteSyncRule Delete a single SyncIQ performance rule.
*/
func (a *Client) DeleteSyncRule(params *DeleteSyncRuleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSyncRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSyncRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSyncRule",
		Method:             "DELETE",
		PathPattern:        "/platform/3/sync/rules/{SyncRuleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSyncRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSyncRuleNoContent), nil

}

/*
DeleteSyncRules Delete all SyncIQ performance rules or all rules of a specified type.
*/
func (a *Client) DeleteSyncRules(params *DeleteSyncRulesParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSyncRulesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSyncRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSyncRules",
		Method:             "DELETE",
		PathPattern:        "/platform/3/sync/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSyncRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSyncRulesNoContent), nil

}

/*
DeleteTargetPolicy Break the target association with this cluster for this policy.
*/
func (a *Client) DeleteTargetPolicy(params *DeleteTargetPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTargetPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTargetPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTargetPolicy",
		Method:             "DELETE",
		PathPattern:        "/platform/1/sync/target/policies/{TargetPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTargetPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTargetPolicyNoContent), nil

}

/*
GetHistoryCPU List cpu performance data.
*/
func (a *Client) GetHistoryCPU(params *GetHistoryCPUParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryCPUOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistoryCPUParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHistoryCpu",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/history/cpu",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHistoryCPUReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHistoryCPUOK), nil

}

/*
GetHistoryFile List file operations performance data.
*/
func (a *Client) GetHistoryFile(params *GetHistoryFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistoryFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHistoryFile",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/history/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHistoryFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHistoryFileOK), nil

}

/*
GetHistoryNetwork List network operations performance data.
*/
func (a *Client) GetHistoryNetwork(params *GetHistoryNetworkParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistoryNetworkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHistoryNetwork",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/history/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHistoryNetworkReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHistoryNetworkOK), nil

}

/*
GetHistoryWorker List worker performance data.
*/
func (a *Client) GetHistoryWorker(params *GetHistoryWorkerParams, authInfo runtime.ClientAuthInfoWriter) (*GetHistoryWorkerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHistoryWorkerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHistoryWorker",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/history/worker",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHistoryWorkerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHistoryWorkerOK), nil

}

/*
GetSyncJob View a single SyncIQ job.
*/
func (a *Client) GetSyncJob(params *GetSyncJobParams, authInfo runtime.ClientAuthInfoWriter) (*GetSyncJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncJob",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/jobs/{SyncJobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncJobOK), nil

}

/*
GetSyncLicense Retrieve license information.
*/
func (a *Client) GetSyncLicense(params *GetSyncLicenseParams, authInfo runtime.ClientAuthInfoWriter) (*GetSyncLicenseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncLicenseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncLicense",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/license",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncLicenseReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncLicenseOK), nil

}

/*
GetSyncPolicy View a single SyncIQ policy.
*/
func (a *Client) GetSyncPolicy(params *GetSyncPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetSyncPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncPolicy",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/policies/{SyncPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncPolicyOK), nil

}

/*
GetSyncReport View a single SyncIQ report.
*/
func (a *Client) GetSyncReport(params *GetSyncReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetSyncReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncReport",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/reports/{SyncReportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncReportOK), nil

}

/*
GetSyncReports Get a list of SyncIQ reports.  By default 10 reports are returned per policy, unless otherwise specified by 'reports_per_policy'.
*/
func (a *Client) GetSyncReports(params *GetSyncReportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSyncReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncReportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncReports",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncReportsOK), nil

}

/*
GetSyncRule View a single SyncIQ performance rule.
*/
func (a *Client) GetSyncRule(params *GetSyncRuleParams, authInfo runtime.ClientAuthInfoWriter) (*GetSyncRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncRule",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/rules/{SyncRuleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncRuleOK), nil

}

/*
GetSyncSettings Retrieve the global SyncIQ settings.
*/
func (a *Client) GetSyncSettings(params *GetSyncSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSyncSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSyncSettings",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSyncSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSyncSettingsOK), nil

}

/*
GetTargetPolicies List all SyncIQ target policies.
*/
func (a *Client) GetTargetPolicies(params *GetTargetPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*GetTargetPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTargetPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTargetPolicies",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/target/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTargetPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTargetPoliciesOK), nil

}

/*
GetTargetPolicy View a single SyncIQ target policy.
*/
func (a *Client) GetTargetPolicy(params *GetTargetPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetTargetPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTargetPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTargetPolicy",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/target/policies/{TargetPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTargetPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTargetPolicyOK), nil

}

/*
GetTargetReport View a single SyncIQ target report.
*/
func (a *Client) GetTargetReport(params *GetTargetReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetTargetReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTargetReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTargetReport",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/target/reports/{TargetReportId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTargetReportReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTargetReportOK), nil

}

/*
GetTargetReports Get a list of SyncIQ target reports.  By default 10 reports are returned per policy, unless otherwise specified by 'reports_per_policy'.
*/
func (a *Client) GetTargetReports(params *GetTargetReportsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTargetReportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTargetReportsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTargetReports",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/target/reports",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTargetReportsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTargetReportsOK), nil

}

/*
ListSyncJobs Get a list of SyncIQ jobs.
*/
func (a *Client) ListSyncJobs(params *ListSyncJobsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSyncJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSyncJobsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSyncJobs",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSyncJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSyncJobsOK), nil

}

/*
ListSyncPolicies List all SyncIQ policies.
*/
func (a *Client) ListSyncPolicies(params *ListSyncPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*ListSyncPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSyncPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSyncPolicies",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSyncPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSyncPoliciesOK), nil

}

/*
ListSyncReportsRotate Whether the rotation is still running or not.
*/
func (a *Client) ListSyncReportsRotate(params *ListSyncReportsRotateParams, authInfo runtime.ClientAuthInfoWriter) (*ListSyncReportsRotateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSyncReportsRotateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSyncReportsRotate",
		Method:             "GET",
		PathPattern:        "/platform/1/sync/reports-rotate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSyncReportsRotateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSyncReportsRotateOK), nil

}

/*
ListSyncRules List all SyncIQ performance rules.
*/
func (a *Client) ListSyncRules(params *ListSyncRulesParams, authInfo runtime.ClientAuthInfoWriter) (*ListSyncRulesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSyncRulesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSyncRules",
		Method:             "GET",
		PathPattern:        "/platform/3/sync/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSyncRulesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSyncRulesOK), nil

}

/*
UpdateSyncJob Perform an action (pause, cancel, etc...) on a single job.
*/
func (a *Client) UpdateSyncJob(params *UpdateSyncJobParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSyncJobNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSyncJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSyncJob",
		Method:             "PUT",
		PathPattern:        "/platform/3/sync/jobs/{SyncJobId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSyncJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSyncJobNoContent), nil

}

/*
UpdateSyncPolicy Modify a single SyncIQ policy.
*/
func (a *Client) UpdateSyncPolicy(params *UpdateSyncPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSyncPolicyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSyncPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSyncPolicy",
		Method:             "PUT",
		PathPattern:        "/platform/3/sync/policies/{SyncPolicyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSyncPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSyncPolicyNoContent), nil

}

/*
UpdateSyncRule Modify a single SyncIQ performance rule.
*/
func (a *Client) UpdateSyncRule(params *UpdateSyncRuleParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSyncRuleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSyncRuleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSyncRule",
		Method:             "PUT",
		PathPattern:        "/platform/3/sync/rules/{SyncRuleId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSyncRuleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSyncRuleNoContent), nil

}

/*
UpdateSyncSettings Modify the global SyncIQ settings.  All input fields are optional, but one or more must be supplied.
*/
func (a *Client) UpdateSyncSettings(params *UpdateSyncSettingsParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSyncSettingsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSyncSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSyncSettings",
		Method:             "PUT",
		PathPattern:        "/platform/3/sync/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSyncSettingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateSyncSettingsNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}