package bq

type InMemoryClient struct{}

func NewInMemoryClient() Client {

	return &InMemoryClient{}
}

func (c *InMemoryClient) Insert(tableId string, src interface{}) error {

	return nil
}

func (c *InMemoryClient) GetAnnouncementCount(from int64) (int64, error) {

	return -1, nil
}
