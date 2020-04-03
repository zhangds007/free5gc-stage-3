/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NSSelection

import (
	"github.com/gin-gonic/gin"

	"free5gc/lib/http_wrapper"
	"free5gc/src/nssf/nssf_handler"
	"free5gc/src/nssf/nssf_handler/nssf_message"
)

func ApiNetworkSliceInformationDocument(c *gin.Context) {
	var request interface{}
	req := http_wrapper.NewRequest(c.Request, request)

	message := nssf_message.NewMessage(nssf_message.NSSelectionGet, req)

	nssf_handler.SendMessage(message)
	rsp := <-message.ResponseChan

	httpResponse := rsp.HttpResponse
	c.JSON(httpResponse.Status, httpResponse.Body)
}
