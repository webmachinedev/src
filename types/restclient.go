package types

type RestClient struct {
	Endpoint string
	Schema Schema
}

func (c *RestClient) Get(resource, id string) ([]byte, error) {
	return nil, nil
}

func (c *RestClient) Set(resource, id string, value []byte) (error) {
	return nil, nil
}

func (c *RestClient) Del(resource, id string) (error) {
	return nil, nil
}
