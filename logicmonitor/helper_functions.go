package logicmonitor

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	lmv1 "github.com/logicmonitor/lm-sdk-go"
)

// add device helper function
func getProperties(d *schema.ResourceData) (t []lmv1.NameAndValue) {
	// interate through hashmap to get custom/system properties
	if r, ok := d.GetOk("properties"); ok {
		for k, v := range r.(map[string]interface{}) {
			t = append(t, lmv1.NameAndValue{Name: k, Value: v.(string)})
		}
	}

	return
}

// get the filters for hostgroup lookup
func getFilters(d *schema.ResourceData) (t string) {
	m := d.Get("filters").(*schema.Set)
	var groupIds []string
	for _, v := range m.List() {
		p := v.(map[string]interface{})

		if p["property"].(string) != "" {
			groupIds = append(groupIds, fmt.Sprintf("%s%s%s", p["property"].(string), p["operator"].(string), p["value"].(string)))
		}

		if p["custom_property_name"].(string) != "" {
			groupIds = append(groupIds, fmt.Sprintf("customProperties.name%s%s,customProperties.value%s%s", p["operator"].(string), p["custom_property_name"].(string), p["operator"].(string), p["custom_property_value"].(string)))
		}
	}
	t = strings.Join(groupIds, ",")
	return
}

// get the filters for collector lookup
func getCollectorFilters(d *schema.ResourceData) (t string) {
	m := d.Get("filters").(*schema.Set)
	var collectorIds []string
	for _, v := range m.List() {
		p := v.(map[string]interface{})
		if p["property"].(string) != "" {
			collectorIds = append(collectorIds, fmt.Sprintf("%s%s%s", p["property"].(string), p["operator"].(string), p["value"].(string)))
		}
	}
	t = strings.Join(collectorIds, ",")
	return
}

// builds the device object with device properties
func makeDeviceObject(m interface{}, d *schema.ResourceData) (output lmv1.RestDevice) {
	// if displayname is not there, we can automatically add ipaddr
	var displayname = d.Get("display_name").(string)
	if displayname == "" {
		displayname = d.Get("ip_addr").(string)
	}

	output = lmv1.RestDevice{
		Name:                 d.Get("ip_addr").(string),
		DisplayName:          displayname,
		DisableAlerting:      d.Get("disable_alerting").(bool),
		HostGroupIds:         d.Get("hostgroup_id").(string),
		PreferredCollectorId: int32(d.Get("collector").(int)),
		CustomProperties:     []lmv1.NameAndValue{},
	}

	return
}

// add hostGroup helper functions
func getGroupProperties(d *schema.ResourceData) (t []lmv1.NameAndValue) {
	// interate through hashmap to get custom/system properties
	if r, ok := d.GetOk("properties"); ok {
		for k, v := range r.(map[string]interface{}) {
			t = append(t, lmv1.NameAndValue{Name: k, Value: v.(string)})
		}
	}

	return
}

// builds the device group object with host group properties
func makeDeviceGroupObject(d *schema.ResourceData) (output lmv1.RestDeviceGroup) {

	output = lmv1.RestDeviceGroup{
		Name:             d.Get("name").(string),
		DisableAlerting:  d.Get("disable_alerting").(bool),
		Description:      d.Get("description").(string),
		AppliesTo:        d.Get("applies_to").(string),
		ParentId:         int32(d.Get("parent_id").(int)),
		CustomProperties: []lmv1.NameAndValue{},
	}

	return
}

// add collector group helper functions
func makeDeviceCollectorGroupObject(d *schema.ResourceData) (output lmv1.RestCollectorGroup) {

	output = lmv1.RestCollectorGroup{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	return
}

func checkStatus(serverResponse int32, serverResponseMessage string, apiResponse int, apiResponseMessage string, err error) error {
	if apiResponse != http.StatusOK {
		return fmt.Errorf("Api Response Error: %s", apiResponseMessage)
	}

	if serverResponse != 200 {
		return fmt.Errorf("%s", serverResponseMessage)
	}

	if err != nil {
		return err
	}

	return nil
}

// function to remove an item from array
func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
