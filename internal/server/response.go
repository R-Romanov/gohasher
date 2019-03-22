package server

type Response struct {
	HashedString int    `json:"hash,omitempty"`
	Error        string `json:"error,omitempty"`
}

func NewResponse() *Response {
	ext := &Response{}
	ext.HashedString = 0
	ext.Error = ""

	return ext
}

func (c *Response) SetHashedString(hash int) {
	c.HashedString = hash
}

func (c *Response) SetError(err string) {
	c.Error = err
}
