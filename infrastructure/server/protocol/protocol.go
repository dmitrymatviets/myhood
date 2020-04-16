package protocol

import "encoding/json"

type MetaWithFrom struct {
	From string `json:"from"`
}

type Envelope struct {
	Meta json.RawMessage `json:"meta"`
}

type Request struct {
	Envelope
	Data json.RawMessage `json:"data" binding:"required"`
}

func NewRequest(data interface{}) *Request {
	metaJson, _ := json.Marshal(map[string]interface{}{
		"from": "myhood",
	})
	dataJson, _ := json.Marshal(data)
	return &Request{
		Envelope: Envelope{Meta: metaJson},
		Data:     dataJson,
	}
}

type ResponseSuccess struct {
	Success int `json:"success"`
	Envelope
	Data json.RawMessage `json:"data" binding:"required"`
}

type RError struct {
	Message string `json:"message" binding:"required"`
	Code    string `json:"code" binding:"required"`
}

type ResponseError struct {
	Success int `json:"success"`
	Envelope
	Error RError `json:"error" binding:"required"`
}

type GoodsApiResponse struct {
	Success int `json:"success"`
	Envelope
	Data  json.RawMessage `json:"data"`
	Error interface{}     `json:"error"`
}

type FileResponseDto struct {
	Path string
	Name string
}
