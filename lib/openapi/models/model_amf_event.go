/*
 * Namf_Communication
 *
 * AMF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AmfEvent struct {
	Type                     AmfEventType           `json:"type"`
	ImmediateFlag            bool                   `json:"immediateFlag,omitempty"`
	AreaList                 []AmfEventArea         `json:"areaList,omitempty"`
	LocationFilterList       []LocationFilter       `json:"locationFilterList,omitempty"`
	SubscribedDataFilterList []SubscribedDataFilter `json:"subscribedDataFilterList,omitempty"`
}