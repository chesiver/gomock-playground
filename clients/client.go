package clients

type Client interface {
	GetData() string
}

type DefaultClient struct {
}

func (c *DefaultClient) GetData() string {
	return "test"
}
func NewDefaultClient() Client {
	return &DefaultClient{}
}
