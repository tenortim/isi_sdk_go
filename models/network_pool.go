// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkPool network pool
// swagger:model NetworkPool
type NetworkPool struct {

	// Name of a valid access zone to map IP address pool to the zone.
	// Min Length: 1
	AccessZone string `json:"access_zone,omitempty"`

	// IP address format.
	// Enum: [ipv4 ipv6]
	AddrFamily string `json:"addr_family,omitempty"`

	// OneFS supports the following NIC aggregation modes.
	// Enum: [roundrobin failover lacp fec]
	AggregationMode string `json:"aggregation_mode,omitempty"`

	// Specifies how IP address allocation is done among pool members.
	// Enum: [dynamic static]
	AllocMethod string `json:"alloc_method,omitempty"`

	// A description of the pool.
	// Max Length: 128
	Description string `json:"description,omitempty"`

	// Name of the groupnet this pool belongs to.
	Groupnet string `json:"groupnet,omitempty"`

	// Unique Pool ID.
	ID string `json:"id,omitempty"`

	// List of interface members in this pool.
	Ifaces []*SubnetsSubnetPoolIface `json:"ifaces"`

	// The name of the pool. It must be unique throughout the given subnet.It's a required field with POST method.
	// Max Length: 32
	Name string `json:"name,omitempty"`

	// List of IP address ranges in this pool.
	Ranges []*SubnetsSubnetPoolRange `json:"ranges"`

	// Rebalance policy..
	// Enum: [auto manual]
	RebalancePolicy string `json:"rebalance_policy,omitempty"`

	// Names of the rules in this pool.
	Rules []string `json:"rules"`

	// Time delay in seconds before a node which has been                 automatically unsuspended becomes usable in SmartConnect                responses for pool zones.
	// Maximum: 86400
	// Minimum: 0
	ScAutoUnsuspendDelay *int64 `json:"sc_auto_unsuspend_delay,omitempty"`

	// SmartConnect client connection balancing policy.
	// Enum: [round_robin conn_count throughput cpu_usage]
	ScConnectPolicy string `json:"sc_connect_policy,omitempty"`

	// SmartConnect zone name for the pool.
	ScDNSZone string `json:"sc_dns_zone,omitempty"`

	// List of SmartConnect zone aliases (DNS names) to the pool.
	ScDNSZoneAliases []string `json:"sc_dns_zone_aliases"`

	// SmartConnect IP failover policy.
	// Enum: [round_robin conn_count throughput cpu_usage]
	ScFailoverPolicy string `json:"sc_failover_policy,omitempty"`

	// Name of SmartConnect service subnet for this pool.
	ScSubnet string `json:"sc_subnet,omitempty"`

	// List of LNNs showing currently suspended nodes in SmartConnect.
	ScSuspendedNodes []int64 `json:"sc_suspended_nodes"`

	// Time to live value for SmartConnect DNS query responses in seconds.
	// Maximum: 2.147483647e+09
	// Minimum: 0
	ScTTL *int64 `json:"sc_ttl,omitempty"`

	// List of interface members in this pool.
	StaticRoutes []*SubnetsSubnetPoolStaticRoute `json:"static_routes"`

	// The name of the subnet.
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this network pool
func (m *NetworkPool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddrFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAggregationMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllocMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRebalancePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScAutoUnsuspendDelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScConnectPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScFailoverPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScTTL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticRoutes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkPool) validateAccessZone(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessZone) { // not required
		return nil
	}

	if err := validate.MinLength("access_zone", "body", string(m.AccessZone), 1); err != nil {
		return err
	}

	return nil
}

var networkPoolTypeAddrFamilyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4","ipv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPoolTypeAddrFamilyPropEnum = append(networkPoolTypeAddrFamilyPropEnum, v)
	}
}

const (

	// NetworkPoolAddrFamilyIPV4 captures enum value "ipv4"
	NetworkPoolAddrFamilyIPV4 string = "ipv4"

	// NetworkPoolAddrFamilyIPV6 captures enum value "ipv6"
	NetworkPoolAddrFamilyIPV6 string = "ipv6"
)

// prop value enum
func (m *NetworkPool) validateAddrFamilyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkPoolTypeAddrFamilyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkPool) validateAddrFamily(formats strfmt.Registry) error {

	if swag.IsZero(m.AddrFamily) { // not required
		return nil
	}

	// value enum
	if err := m.validateAddrFamilyEnum("addr_family", "body", m.AddrFamily); err != nil {
		return err
	}

	return nil
}

var networkPoolTypeAggregationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["roundrobin","failover","lacp","fec"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPoolTypeAggregationModePropEnum = append(networkPoolTypeAggregationModePropEnum, v)
	}
}

const (

	// NetworkPoolAggregationModeRoundrobin captures enum value "roundrobin"
	NetworkPoolAggregationModeRoundrobin string = "roundrobin"

	// NetworkPoolAggregationModeFailover captures enum value "failover"
	NetworkPoolAggregationModeFailover string = "failover"

	// NetworkPoolAggregationModeLacp captures enum value "lacp"
	NetworkPoolAggregationModeLacp string = "lacp"

	// NetworkPoolAggregationModeFec captures enum value "fec"
	NetworkPoolAggregationModeFec string = "fec"
)

// prop value enum
func (m *NetworkPool) validateAggregationModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkPoolTypeAggregationModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkPool) validateAggregationMode(formats strfmt.Registry) error {

	if swag.IsZero(m.AggregationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateAggregationModeEnum("aggregation_mode", "body", m.AggregationMode); err != nil {
		return err
	}

	return nil
}

var networkPoolTypeAllocMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dynamic","static"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPoolTypeAllocMethodPropEnum = append(networkPoolTypeAllocMethodPropEnum, v)
	}
}

const (

	// NetworkPoolAllocMethodDynamic captures enum value "dynamic"
	NetworkPoolAllocMethodDynamic string = "dynamic"

	// NetworkPoolAllocMethodStatic captures enum value "static"
	NetworkPoolAllocMethodStatic string = "static"
)

// prop value enum
func (m *NetworkPool) validateAllocMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkPoolTypeAllocMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkPool) validateAllocMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.AllocMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateAllocMethodEnum("alloc_method", "body", m.AllocMethod); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPool) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 128); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPool) validateIfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Ifaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Ifaces); i++ {
		if swag.IsZero(m.Ifaces[i]) { // not required
			continue
		}

		if m.Ifaces[i] != nil {
			if err := m.Ifaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ifaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkPool) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 32); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPool) validateRanges(formats strfmt.Registry) error {

	if swag.IsZero(m.Ranges) { // not required
		return nil
	}

	for i := 0; i < len(m.Ranges); i++ {
		if swag.IsZero(m.Ranges[i]) { // not required
			continue
		}

		if m.Ranges[i] != nil {
			if err := m.Ranges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ranges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var networkPoolTypeRebalancePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["auto","manual"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPoolTypeRebalancePolicyPropEnum = append(networkPoolTypeRebalancePolicyPropEnum, v)
	}
}

const (

	// NetworkPoolRebalancePolicyAuto captures enum value "auto"
	NetworkPoolRebalancePolicyAuto string = "auto"

	// NetworkPoolRebalancePolicyManual captures enum value "manual"
	NetworkPoolRebalancePolicyManual string = "manual"
)

// prop value enum
func (m *NetworkPool) validateRebalancePolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkPoolTypeRebalancePolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkPool) validateRebalancePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.RebalancePolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateRebalancePolicyEnum("rebalance_policy", "body", m.RebalancePolicy); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPool) validateScAutoUnsuspendDelay(formats strfmt.Registry) error {

	if swag.IsZero(m.ScAutoUnsuspendDelay) { // not required
		return nil
	}

	if err := validate.MinimumInt("sc_auto_unsuspend_delay", "body", int64(*m.ScAutoUnsuspendDelay), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("sc_auto_unsuspend_delay", "body", int64(*m.ScAutoUnsuspendDelay), 86400, false); err != nil {
		return err
	}

	return nil
}

var networkPoolTypeScConnectPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["round_robin","conn_count","throughput","cpu_usage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPoolTypeScConnectPolicyPropEnum = append(networkPoolTypeScConnectPolicyPropEnum, v)
	}
}

const (

	// NetworkPoolScConnectPolicyRoundRobin captures enum value "round_robin"
	NetworkPoolScConnectPolicyRoundRobin string = "round_robin"

	// NetworkPoolScConnectPolicyConnCount captures enum value "conn_count"
	NetworkPoolScConnectPolicyConnCount string = "conn_count"

	// NetworkPoolScConnectPolicyThroughput captures enum value "throughput"
	NetworkPoolScConnectPolicyThroughput string = "throughput"

	// NetworkPoolScConnectPolicyCPUUsage captures enum value "cpu_usage"
	NetworkPoolScConnectPolicyCPUUsage string = "cpu_usage"
)

// prop value enum
func (m *NetworkPool) validateScConnectPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkPoolTypeScConnectPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkPool) validateScConnectPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ScConnectPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateScConnectPolicyEnum("sc_connect_policy", "body", m.ScConnectPolicy); err != nil {
		return err
	}

	return nil
}

var networkPoolTypeScFailoverPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["round_robin","conn_count","throughput","cpu_usage"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkPoolTypeScFailoverPolicyPropEnum = append(networkPoolTypeScFailoverPolicyPropEnum, v)
	}
}

const (

	// NetworkPoolScFailoverPolicyRoundRobin captures enum value "round_robin"
	NetworkPoolScFailoverPolicyRoundRobin string = "round_robin"

	// NetworkPoolScFailoverPolicyConnCount captures enum value "conn_count"
	NetworkPoolScFailoverPolicyConnCount string = "conn_count"

	// NetworkPoolScFailoverPolicyThroughput captures enum value "throughput"
	NetworkPoolScFailoverPolicyThroughput string = "throughput"

	// NetworkPoolScFailoverPolicyCPUUsage captures enum value "cpu_usage"
	NetworkPoolScFailoverPolicyCPUUsage string = "cpu_usage"
)

// prop value enum
func (m *NetworkPool) validateScFailoverPolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkPoolTypeScFailoverPolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NetworkPool) validateScFailoverPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.ScFailoverPolicy) { // not required
		return nil
	}

	// value enum
	if err := m.validateScFailoverPolicyEnum("sc_failover_policy", "body", m.ScFailoverPolicy); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPool) validateScTTL(formats strfmt.Registry) error {

	if swag.IsZero(m.ScTTL) { // not required
		return nil
	}

	if err := validate.MinimumInt("sc_ttl", "body", int64(*m.ScTTL), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("sc_ttl", "body", int64(*m.ScTTL), 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkPool) validateStaticRoutes(formats strfmt.Registry) error {

	if swag.IsZero(m.StaticRoutes) { // not required
		return nil
	}

	for i := 0; i < len(m.StaticRoutes); i++ {
		if swag.IsZero(m.StaticRoutes[i]) { // not required
			continue
		}

		if m.StaticRoutes[i] != nil {
			if err := m.StaticRoutes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("static_routes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkPool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkPool) UnmarshalBinary(b []byte) error {
	var res NetworkPool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}