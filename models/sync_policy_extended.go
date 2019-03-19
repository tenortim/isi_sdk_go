// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SyncPolicyExtended sync policy extended
// swagger:model SyncPolicyExtended
type SyncPolicyExtended struct {

	// If set to true, SyncIQ will perform failback configuration tasks during the next job run, rather than waiting to perform those tasks during the failback process. Performing these tasks ahead of time will increase the speed of failback operations.
	// Required: true
	AcceleratedFailback *bool `json:"accelerated_failback"`

	// If 'copy', source files will be copied to the target cluster.  If 'sync', the target directory will be made an image of the source directory:  Files and directories that have been deleted on the source, have been moved within the target directory, or no longer match the selection criteria will be deleted from the target directory.
	// Required: true
	// Enum: [copy sync]
	Action *string `json:"action"`

	// NOTE: This field should not be changed without the help of Isilon support.  Enable/disable UDP-based data transfer.
	// Required: true
	BurstMode *bool `json:"burst_mode"`

	// If true, retain previous source snapshot and incremental repstate, both of which are required for changelist creation.
	// Required: true
	Changelist *bool `json:"changelist"`

	// If true, the sync target performs cyclic redundancy checks (CRC) on the data as it is received.
	// Required: true
	CheckIntegrity *bool `json:"check_integrity"`

	// If set to deny, replicates all CloudPools smartlinks to the target cluster as smartlinks; if the target cluster does not support the smartlinks, the job will fail. If set to force, replicates all smartlinks to the target cluster as regular files. If set to allow, SyncIQ will attempt to replicate smartlinks to the target cluster as smartlinks; if the target cluster does not support the smartlinks, SyncIQ will replicate the smartlinks as regular files.
	// Required: true
	// Enum: [deny allow force]
	CloudDeepCopy *string `json:"cloud_deep_copy"`

	// NOTE: This field should not be changed without the help of Isilon support.  If true, the most recent run of this policy encountered an error and this policy will not start any more scheduled jobs until this field is manually set back to 'false'.
	// Required: true
	Conflicted *bool `json:"conflicted"`

	// User-assigned description of this sync policy.
	// Required: true
	Description *string `json:"description"`

	// NOTE: This field should not be changed without the help of Isilon support.  If true, the 7.2+ file splitting capability will be disabled.
	// Required: true
	DisableFileSplit *bool `json:"disable_file_split"`

	// NOTE: This field should not be changed without the help of Isilon support.  Enable/disable sync failover/failback.
	// Required: true
	DisableFofb *bool `json:"disable_fofb"`

	// NOTE: This field should not be changed without the help of Isilon support.  Enable/disable the 6.5+ STF based data transfer and uses only treewalk.
	// Required: true
	DisableStf *bool `json:"disable_stf"`

	// If true, jobs will be automatically run based on this policy, according to its schedule.
	// Required: true
	Enabled *bool `json:"enabled"`

	// NOTE: This field should not be changed without the help of Isilon support.  Continue sending files even with the corrupted filesystem.
	// Required: true
	ExpectedDataloss *bool `json:"expected_dataloss"`

	// A file matching pattern, organized as an OR'ed set of AND'ed file criteria, for example ((a AND b) OR (x AND y)) used to define a set of files with specific properties.  Policies of type 'sync' cannot use 'path' or time criteria in their matching patterns, but policies of type 'copy' can use all listed criteria.
	// Required: true
	FileMatchingPattern *ReportSubreportPolicyFileMatchingPattern `json:"file_matching_pattern"`

	// NOTE: This field should not be changed without the help of Isilon support.  Determines whether data is sent only through the subnet and pool specified in the "source_network" field. This option can be useful if there are multiple interfaces for the given source subnet.  If you enable this option, the net.inet.ip.choose_ifa_by_ipsrc sysctl should be set.
	// Required: true
	ForceInterface *bool `json:"force_interface"`

	// This field is false if the policy is in its initial sync state and true otherwise.  Setting this field to false will reset the policy's sync state.
	// Required: true
	HasSyncState *bool `json:"has_sync_state"`

	// The system ID given to this sync policy.
	// Required: true
	ID *string `json:"id"`

	// If --schedule is set to When-Source-Modified, the duration to wait after a modification is made before starting a job (default is 0 seconds).
	// Minimum: 0
	JobDelay *int64 `json:"job_delay,omitempty"`

	// This is the state of the most recent job for this policy.
	// Required: true
	LastJobState *string `json:"last_job_state"`

	// The most recent time a job was started for this policy.  Value is null if the policy has never been run.
	// Minimum: 0
	LastStarted *int64 `json:"last_started,omitempty"`

	// Timestamp of last known successfully completed synchronization.  Value is null if the policy has never completed successfully.
	// Minimum: 0
	LastSuccess *int64 `json:"last_success,omitempty"`

	// Severity an event must reach before it is logged.
	// Required: true
	// Enum: [fatal error notice info copy debug trace]
	LogLevel *string `json:"log_level"`

	// If true, the system will log any files or directories that are deleted due to a sync.
	// Required: true
	LogRemovedFiles *bool `json:"log_removed_files"`

	// User-assigned name of this sync policy.
	// Required: true
	Name *string `json:"name"`

	// This is the next time a job is scheduled to run for this policy in Unix epoch seconds.  This field is null if the job is not scheduled.
	NextRun int64 `json:"next_run,omitempty"`

	// Indicates if a password is set for accessing the target cluster. Password value is not shown with GET.
	// Required: true
	PasswordSet *bool `json:"password_set"`

	// Determines the priority level of a policy. Policies with higher priority will have precedence to run over lower priority policies. Valid range is [0, 1]. Default is 0.
	// Required: true
	Priority *int64 `json:"priority"`

	// Length of time (in seconds) a policy report will be stored.
	// Required: true
	// Minimum: 0
	ReportMaxAge *int64 `json:"report_max_age"`

	// Maximum number of policy reports that will be stored on the system.
	// Required: true
	// Maximum: 2000
	// Minimum: 1
	ReportMaxCount *int64 `json:"report_max_count"`

	// If you specify true, and you specify a SmartConnect zone in the "target_host" field, replication policies will connect only to nodes in the specified SmartConnect zone.  If you specify false, replication policies are not restricted to specific nodes on the target cluster.
	// Required: true
	RestrictTargetNetwork *bool `json:"restrict_target_network"`

	// If --schedule is set to a time/date, an alert is created if the specified RPO for this policy is exceeded. The default value is 0, which will not generate RPO alerts.
	// Minimum: 0
	RpoAlert *int64 `json:"rpo_alert,omitempty"`

	// The schedule on which new jobs will be run for this policy.
	// Required: true
	Schedule *string `json:"schedule"`

	// Skip DNS lookup of target IPs.
	// Required: true
	SkipLookup *bool `json:"skip_lookup"`

	// If true and --schedule is set to a time/date, the policy will not run if no changes have been made to the contents of the source directory since the last job successfully completed.
	// Required: true
	SkipWhenSourceUnmodified *bool `json:"skip_when_source_unmodified"`

	// If true, snapshot-triggered syncs will include snapshots taken before policy creation time (requires --schedule when-snapshot-taken).
	// Required: true
	SnapshotSyncExisting *bool `json:"snapshot_sync_existing"`

	// The naming pattern that a snapshot must match to trigger a sync when the schedule is when-snapshot-taken (default is "*").
	// Required: true
	SnapshotSyncPattern *string `json:"snapshot_sync_pattern"`

	// Directories that will be excluded from the sync.  Modifying this field will result in a full synchronization of all data.
	// Required: true
	SourceExcludeDirectories []string `json:"source_exclude_directories"`

	// Directories that will be included in the sync.  Modifying this field will result in a full synchronization of all data.
	// Required: true
	SourceIncludeDirectories []string `json:"source_include_directories"`

	// Restricts replication policies on the local cluster to running on the specified subnet and pool.
	SourceNetwork *SyncPolicySourceNetwork `json:"source_network,omitempty"`

	// The root directory on the source cluster the files will be synced from.  Modifying this field will result in a full synchronization of all data.
	// Required: true
	SourceRootPath *string `json:"source_root_path"`

	// If true, archival snapshots of the source data will be taken on the source cluster before a sync.
	// Required: true
	SourceSnapshotArchive *bool `json:"source_snapshot_archive"`

	// The length of time in seconds to keep snapshots on the source cluster.
	// Required: true
	// Minimum: 0
	SourceSnapshotExpiration *int64 `json:"source_snapshot_expiration"`

	// The name pattern for snapshots taken on the source cluster before a sync.
	// Required: true
	SourceSnapshotPattern *string `json:"source_snapshot_pattern"`

	// If true, the target creates diffs against the original sync.
	// Required: true
	TargetCompareInitialSync *bool `json:"target_compare_initial_sync"`

	// If true, target cluster will detect if files have been changed on the target by legacy tree walk syncs.
	// Required: true
	TargetDetectModifications *bool `json:"target_detect_modifications"`

	// Hostname or IP address of sync target cluster.  Modifying the target cluster host can result in the policy being unrunnable if the new target does not match the current target association.
	// Required: true
	TargetHost *string `json:"target_host"`

	// Absolute filesystem path on the target cluster for the sync destination.
	// Required: true
	TargetPath *string `json:"target_path"`

	// The alias of the snapshot taken on the target cluster after the sync completes. A value of @DEFAULT will reset this field to the default creation value.
	// Required: true
	TargetSnapshotAlias *string `json:"target_snapshot_alias"`

	// If true, archival snapshots of the target data will be taken on the target cluster after successful sync completions.
	// Required: true
	TargetSnapshotArchive *bool `json:"target_snapshot_archive"`

	// The length of time in seconds to keep snapshots on the target cluster.
	// Required: true
	// Minimum: 0
	TargetSnapshotExpiration *int64 `json:"target_snapshot_expiration"`

	// The name pattern for snapshots taken on the target cluster after the sync completes.  A value of @DEFAULT will reset this field to the default creation value.
	// Required: true
	TargetSnapshotPattern *string `json:"target_snapshot_pattern"`

	// The number of worker threads on a node performing a sync.
	// Required: true
	// Maximum: 20
	// Minimum: 1
	WorkersPerNode *int64 `json:"workers_per_node"`
}

// Validate validates this sync policy extended
func (m *SyncPolicyExtended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceleratedFailback(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBurstMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChangelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckIntegrity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudDeepCopy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConflicted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisableFileSplit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisableFofb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisableStf(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpectedDataloss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileMatchingPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForceInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasSyncState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobDelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastJobState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSuccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogRemovedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportMaxCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictTargetNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRpoAlert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkipLookup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkipWhenSourceUnmodified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotSyncExisting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotSyncPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceExcludeDirectories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceIncludeDirectories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceRootPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceSnapshotArchive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceSnapshotExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceSnapshotPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetCompareInitialSync(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetDetectModifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetSnapshotAlias(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetSnapshotArchive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetSnapshotExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetSnapshotPattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkersPerNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SyncPolicyExtended) validateAcceleratedFailback(formats strfmt.Registry) error {

	if err := validate.Required("accelerated_failback", "body", m.AcceleratedFailback); err != nil {
		return err
	}

	return nil
}

var syncPolicyExtendedTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["copy","sync"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncPolicyExtendedTypeActionPropEnum = append(syncPolicyExtendedTypeActionPropEnum, v)
	}
}

const (

	// SyncPolicyExtendedActionCopy captures enum value "copy"
	SyncPolicyExtendedActionCopy string = "copy"

	// SyncPolicyExtendedActionSync captures enum value "sync"
	SyncPolicyExtendedActionSync string = "sync"
)

// prop value enum
func (m *SyncPolicyExtended) validateActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, syncPolicyExtendedTypeActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SyncPolicyExtended) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateBurstMode(formats strfmt.Registry) error {

	if err := validate.Required("burst_mode", "body", m.BurstMode); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateChangelist(formats strfmt.Registry) error {

	if err := validate.Required("changelist", "body", m.Changelist); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateCheckIntegrity(formats strfmt.Registry) error {

	if err := validate.Required("check_integrity", "body", m.CheckIntegrity); err != nil {
		return err
	}

	return nil
}

var syncPolicyExtendedTypeCloudDeepCopyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deny","allow","force"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncPolicyExtendedTypeCloudDeepCopyPropEnum = append(syncPolicyExtendedTypeCloudDeepCopyPropEnum, v)
	}
}

const (

	// SyncPolicyExtendedCloudDeepCopyDeny captures enum value "deny"
	SyncPolicyExtendedCloudDeepCopyDeny string = "deny"

	// SyncPolicyExtendedCloudDeepCopyAllow captures enum value "allow"
	SyncPolicyExtendedCloudDeepCopyAllow string = "allow"

	// SyncPolicyExtendedCloudDeepCopyForce captures enum value "force"
	SyncPolicyExtendedCloudDeepCopyForce string = "force"
)

// prop value enum
func (m *SyncPolicyExtended) validateCloudDeepCopyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, syncPolicyExtendedTypeCloudDeepCopyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SyncPolicyExtended) validateCloudDeepCopy(formats strfmt.Registry) error {

	if err := validate.Required("cloud_deep_copy", "body", m.CloudDeepCopy); err != nil {
		return err
	}

	// value enum
	if err := m.validateCloudDeepCopyEnum("cloud_deep_copy", "body", *m.CloudDeepCopy); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateConflicted(formats strfmt.Registry) error {

	if err := validate.Required("conflicted", "body", m.Conflicted); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateDisableFileSplit(formats strfmt.Registry) error {

	if err := validate.Required("disable_file_split", "body", m.DisableFileSplit); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateDisableFofb(formats strfmt.Registry) error {

	if err := validate.Required("disable_fofb", "body", m.DisableFofb); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateDisableStf(formats strfmt.Registry) error {

	if err := validate.Required("disable_stf", "body", m.DisableStf); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateExpectedDataloss(formats strfmt.Registry) error {

	if err := validate.Required("expected_dataloss", "body", m.ExpectedDataloss); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateFileMatchingPattern(formats strfmt.Registry) error {

	if err := validate.Required("file_matching_pattern", "body", m.FileMatchingPattern); err != nil {
		return err
	}

	if m.FileMatchingPattern != nil {
		if err := m.FileMatchingPattern.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("file_matching_pattern")
			}
			return err
		}
	}

	return nil
}

func (m *SyncPolicyExtended) validateForceInterface(formats strfmt.Registry) error {

	if err := validate.Required("force_interface", "body", m.ForceInterface); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateHasSyncState(formats strfmt.Registry) error {

	if err := validate.Required("has_sync_state", "body", m.HasSyncState); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateJobDelay(formats strfmt.Registry) error {

	if swag.IsZero(m.JobDelay) { // not required
		return nil
	}

	if err := validate.MinimumInt("job_delay", "body", int64(*m.JobDelay), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateLastJobState(formats strfmt.Registry) error {

	if err := validate.Required("last_job_state", "body", m.LastJobState); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateLastStarted(formats strfmt.Registry) error {

	if swag.IsZero(m.LastStarted) { // not required
		return nil
	}

	if err := validate.MinimumInt("last_started", "body", int64(*m.LastStarted), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateLastSuccess(formats strfmt.Registry) error {

	if swag.IsZero(m.LastSuccess) { // not required
		return nil
	}

	if err := validate.MinimumInt("last_success", "body", int64(*m.LastSuccess), 0, false); err != nil {
		return err
	}

	return nil
}

var syncPolicyExtendedTypeLogLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fatal","error","notice","info","copy","debug","trace"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncPolicyExtendedTypeLogLevelPropEnum = append(syncPolicyExtendedTypeLogLevelPropEnum, v)
	}
}

const (

	// SyncPolicyExtendedLogLevelFatal captures enum value "fatal"
	SyncPolicyExtendedLogLevelFatal string = "fatal"

	// SyncPolicyExtendedLogLevelError captures enum value "error"
	SyncPolicyExtendedLogLevelError string = "error"

	// SyncPolicyExtendedLogLevelNotice captures enum value "notice"
	SyncPolicyExtendedLogLevelNotice string = "notice"

	// SyncPolicyExtendedLogLevelInfo captures enum value "info"
	SyncPolicyExtendedLogLevelInfo string = "info"

	// SyncPolicyExtendedLogLevelCopy captures enum value "copy"
	SyncPolicyExtendedLogLevelCopy string = "copy"

	// SyncPolicyExtendedLogLevelDebug captures enum value "debug"
	SyncPolicyExtendedLogLevelDebug string = "debug"

	// SyncPolicyExtendedLogLevelTrace captures enum value "trace"
	SyncPolicyExtendedLogLevelTrace string = "trace"
)

// prop value enum
func (m *SyncPolicyExtended) validateLogLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, syncPolicyExtendedTypeLogLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SyncPolicyExtended) validateLogLevel(formats strfmt.Registry) error {

	if err := validate.Required("log_level", "body", m.LogLevel); err != nil {
		return err
	}

	// value enum
	if err := m.validateLogLevelEnum("log_level", "body", *m.LogLevel); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateLogRemovedFiles(formats strfmt.Registry) error {

	if err := validate.Required("log_removed_files", "body", m.LogRemovedFiles); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validatePasswordSet(formats strfmt.Registry) error {

	if err := validate.Required("password_set", "body", m.PasswordSet); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateReportMaxAge(formats strfmt.Registry) error {

	if err := validate.Required("report_max_age", "body", m.ReportMaxAge); err != nil {
		return err
	}

	if err := validate.MinimumInt("report_max_age", "body", int64(*m.ReportMaxAge), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateReportMaxCount(formats strfmt.Registry) error {

	if err := validate.Required("report_max_count", "body", m.ReportMaxCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("report_max_count", "body", int64(*m.ReportMaxCount), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("report_max_count", "body", int64(*m.ReportMaxCount), 2000, false); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateRestrictTargetNetwork(formats strfmt.Registry) error {

	if err := validate.Required("restrict_target_network", "body", m.RestrictTargetNetwork); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateRpoAlert(formats strfmt.Registry) error {

	if swag.IsZero(m.RpoAlert) { // not required
		return nil
	}

	if err := validate.MinimumInt("rpo_alert", "body", int64(*m.RpoAlert), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSkipLookup(formats strfmt.Registry) error {

	if err := validate.Required("skip_lookup", "body", m.SkipLookup); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSkipWhenSourceUnmodified(formats strfmt.Registry) error {

	if err := validate.Required("skip_when_source_unmodified", "body", m.SkipWhenSourceUnmodified); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSnapshotSyncExisting(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_sync_existing", "body", m.SnapshotSyncExisting); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSnapshotSyncPattern(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_sync_pattern", "body", m.SnapshotSyncPattern); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSourceExcludeDirectories(formats strfmt.Registry) error {

	if err := validate.Required("source_exclude_directories", "body", m.SourceExcludeDirectories); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSourceIncludeDirectories(formats strfmt.Registry) error {

	if err := validate.Required("source_include_directories", "body", m.SourceIncludeDirectories); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSourceNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceNetwork) { // not required
		return nil
	}

	if m.SourceNetwork != nil {
		if err := m.SourceNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source_network")
			}
			return err
		}
	}

	return nil
}

func (m *SyncPolicyExtended) validateSourceRootPath(formats strfmt.Registry) error {

	if err := validate.Required("source_root_path", "body", m.SourceRootPath); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSourceSnapshotArchive(formats strfmt.Registry) error {

	if err := validate.Required("source_snapshot_archive", "body", m.SourceSnapshotArchive); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSourceSnapshotExpiration(formats strfmt.Registry) error {

	if err := validate.Required("source_snapshot_expiration", "body", m.SourceSnapshotExpiration); err != nil {
		return err
	}

	if err := validate.MinimumInt("source_snapshot_expiration", "body", int64(*m.SourceSnapshotExpiration), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateSourceSnapshotPattern(formats strfmt.Registry) error {

	if err := validate.Required("source_snapshot_pattern", "body", m.SourceSnapshotPattern); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetCompareInitialSync(formats strfmt.Registry) error {

	if err := validate.Required("target_compare_initial_sync", "body", m.TargetCompareInitialSync); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetDetectModifications(formats strfmt.Registry) error {

	if err := validate.Required("target_detect_modifications", "body", m.TargetDetectModifications); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetHost(formats strfmt.Registry) error {

	if err := validate.Required("target_host", "body", m.TargetHost); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetPath(formats strfmt.Registry) error {

	if err := validate.Required("target_path", "body", m.TargetPath); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetSnapshotAlias(formats strfmt.Registry) error {

	if err := validate.Required("target_snapshot_alias", "body", m.TargetSnapshotAlias); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetSnapshotArchive(formats strfmt.Registry) error {

	if err := validate.Required("target_snapshot_archive", "body", m.TargetSnapshotArchive); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetSnapshotExpiration(formats strfmt.Registry) error {

	if err := validate.Required("target_snapshot_expiration", "body", m.TargetSnapshotExpiration); err != nil {
		return err
	}

	if err := validate.MinimumInt("target_snapshot_expiration", "body", int64(*m.TargetSnapshotExpiration), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateTargetSnapshotPattern(formats strfmt.Registry) error {

	if err := validate.Required("target_snapshot_pattern", "body", m.TargetSnapshotPattern); err != nil {
		return err
	}

	return nil
}

func (m *SyncPolicyExtended) validateWorkersPerNode(formats strfmt.Registry) error {

	if err := validate.Required("workers_per_node", "body", m.WorkersPerNode); err != nil {
		return err
	}

	if err := validate.MinimumInt("workers_per_node", "body", int64(*m.WorkersPerNode), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("workers_per_node", "body", int64(*m.WorkersPerNode), 20, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SyncPolicyExtended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncPolicyExtended) UnmarshalBinary(b []byte) error {
	var res SyncPolicyExtended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
