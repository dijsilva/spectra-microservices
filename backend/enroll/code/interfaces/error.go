/*
 * @File: models.message.go
 * @Description: Defines Message information will be returned to the clients
 * @Author: Nguyen Truong Duong (seedotech@gmail.com)
 */
package interfaces

// Message defines the response message
type ErrorResponse struct {
	Data ErrorMessage `json:"data"`
}

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
