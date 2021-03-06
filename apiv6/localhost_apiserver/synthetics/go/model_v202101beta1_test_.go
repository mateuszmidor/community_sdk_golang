/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

import (
	"time"
)

type V202101beta1Test struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Type string `json:"type,omitempty"`

	DeviceId string `json:"deviceId,omitempty"`

	Status V202101beta1TestStatus `json:"status,omitempty"`

	Settings V202101beta1TestSettings `json:"settings,omitempty"`

	ExpiresOn time.Time `json:"expiresOn,omitempty"`

	Cdate time.Time `json:"cdate,omitempty"`

	Edate time.Time `json:"edate,omitempty"`

	CreatedBy V202101beta1UserInfo `json:"createdBy,omitempty"`

	LastUpdatedBy V202101beta1UserInfo `json:"lastUpdatedBy,omitempty"`
}
