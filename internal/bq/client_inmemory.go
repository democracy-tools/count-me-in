package bq

type InMemoryClient struct{}

func NewInMemoryClient() Client {

	return &InMemoryClient{}
}

func (c *InMemoryClient) Insert(tableId string, src interface{}) error {

	return nil
}
