/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/gin-gonic/gin"
	"free5gc/lib/http_wrapper"
	"free5gc/src/udr/udr_handler/udr_message"
)

// QuerySmfSelectData - Retrieves the SMF selection subscription data of a UE
func QuerySmfSelectData(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["servingPlmnId"] = c.Params.ByName("servingPlmnId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventQuerySmfSelectData, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
