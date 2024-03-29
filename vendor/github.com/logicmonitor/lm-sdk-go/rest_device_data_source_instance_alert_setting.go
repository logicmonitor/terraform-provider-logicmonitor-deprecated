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

type RestDeviceDataSourceInstanceAlertSetting struct {
	GlobalAlertExpr string `json:"globalAlertExpr,omitempty"`

	AlertClearInterval int32 `json:"alertClearInterval,omitempty"`

	DisableAlerting bool `json:"disableAlerting,omitempty"`

	AlertExprNote string `json:"alertExprNote,omitempty"`

	DataPointDescription string `json:"dataPointDescription,omitempty"`

	DataSourceInstanceId int32 `json:"dataSourceInstanceId,omitempty"`

	DisableDpAlertHostGroups string `json:"disableDpAlertHostGroups,omitempty"`

	DataPointName string `json:"dataPointName,omitempty"`

	DataPointId int32 `json:"dataPointId,omitempty"`

	DeviceGroupId int32 `json:"deviceGroupId,omitempty"`

	CurrentAlertId int32 `json:"currentAlertId,omitempty"`

	ParentDeviceGroupAlertExprList []DeviceGroupAlertThresholdInfo `json:"parentDeviceGroupAlertExprList,omitempty"`

	AlertingDisabledOn string `json:"alertingDisabledOn,omitempty"`

	Id int32 `json:"id,omitempty"`

	DataSourceInstanceAlias string `json:"dataSourceInstanceAlias,omitempty"`

	DeviceGroupFullPath string `json:"deviceGroupFullPath,omitempty"`

	AlertExpr string `json:"alertExpr,omitempty"`

	AlertTransitionInterval int32 `json:"alertTransitionInterval,omitempty"`
}
