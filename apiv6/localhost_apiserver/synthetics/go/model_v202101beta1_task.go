/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

type V202101beta1Task struct {
	Id string `json:"id,omitempty"`

	TestId string `json:"testId,omitempty"`

	DeviceId string `json:"deviceId,omitempty"`

	State V202101beta1TaskState `json:"state,omitempty"`

	Status string `json:"status,omitempty"`

	Family V202101beta1IpFamily `json:"family,omitempty"`

	Ping V202101beta1PingTaskDefinition `json:"ping,omitempty"`

	Traceroute V202101beta1TraceTaskDefinition `json:"traceroute,omitempty"`

	Http V202101beta1HttpTaskDefinition `json:"http,omitempty"`

	Knock V202101beta1KnockTaskDefinition `json:"knock,omitempty"`

	Dns V202101beta1DnsTaskDefinition `json:"dns,omitempty"`

	Shake V202101beta1ShakeTaskDefinition `json:"shake,omitempty"`
}
