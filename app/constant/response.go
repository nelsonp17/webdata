package constant

type Response struct {
	Ok    interface{} `json:"ok,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type Array struct {
	Response map[string]interface{}
}
