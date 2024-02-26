package models

type Response struct {
	Error bool        `json:"error"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
}

type DataResponse map[string]interface{}
