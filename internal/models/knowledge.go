package models

import "fmt"

type Knowledge struct {
	Id     uint64
	UserId uint64
	Topic  uint64
	Text   string
}

// String prints values of all fields.
func (knowledge Knowledge) String() string {
	return fmt.Sprintf(
		"Id: %v\nUser Id: %v\nTopic: %v\nText: %v\n",
		knowledge.Id,
		knowledge.UserId,
		knowledge.Topic,
		knowledge.Text)
}
