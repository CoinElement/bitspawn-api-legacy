/*

 */

package controllers

// Response
type Response struct {
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
	Body  interface{} `json:"body,omitempty"`
}
