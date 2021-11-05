package pool

type Connection struct {
	Id string
}

func (c *Connection) GetID() string {
	return c.Id
}
