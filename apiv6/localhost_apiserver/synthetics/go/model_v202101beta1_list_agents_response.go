/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

type V202101beta1ListAgentsResponse struct {
	Agents []V202101beta1Agent `json:"agents,omitempty"`

	InvalidAgentsCount int64 `json:"invalidAgentsCount,omitempty"`
}
