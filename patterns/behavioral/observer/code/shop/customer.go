package shop

import "fmt"

type customer struct {
	id string
}

func NewCustomer(id string) *customer {
	return &customer{
		id: id,
	}
}

func (c *customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) GetID() string {
	return c.id
}
