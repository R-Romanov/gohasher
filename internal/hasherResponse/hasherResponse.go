package hasherResponse

type HasherResponse struct {
	HashedString int    `json:"hash,omitempty"`
	Error        string `json:"error,omitempty"`
}

func NewHasherResponse() *HasherResponse {
	ext := &HasherResponse{}
	ext.HashedString = 0
	ext.Error = ""

	return ext
}

func (c *HasherResponse) SetHashedString(hash int) {
	c.HashedString = hash
}

func (c *HasherResponse) SetError(err string) {
	c.Error = err
}
