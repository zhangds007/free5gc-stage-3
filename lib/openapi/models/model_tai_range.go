/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type TaiRange struct {
	PlmnId *PlmnId `json:"plmnId" yaml:"plmnId" bson:"plmnId" mapstructure:"PlmnId"`
	TacRangeList []TacRange `json:"tacRangeList" yaml:"tacRangeList" bson:"tacRangeList" mapstructure:"TacRangeList"`
}
