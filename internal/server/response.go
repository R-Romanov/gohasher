package server

const (
	StatusOk    = "ok"
	StatusError = "error"
)

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Hash    int    `json:"hash,omitempty"`
}

func NewResponse() *Response {
	ext := &Response{}
	ext.Status = StatusOk

	return ext
}

func (c *Response) SetHash(hash int) {
	c.Hash = hash
}

func (c *Response) SetError(err string) {
	c.Status = StatusError
	c.Message = err
}
