/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package logicmonitor

type DowngradeInfo struct {
	CreatedBy string `json:"createdBy,omitempty"`

	Level string `json:"level,omitempty"`

	EndEpoch int64 `json:"endEpoch,omitempty"`

	Description string `json:"description,omitempty"`

	Type_ string `json:"type,omitempty"`

	MajorVersion int32 `json:"majorVersion"`

	MinorVersion int32 `json:"minorVersion"`

	StartEpoch int64 `json:"startEpoch"`
}
