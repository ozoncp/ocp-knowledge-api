package models

import "fmt"

// Knowledge student's knowledge information.
type Knowledge struct {
	Id     uint64 `db:"id"`
	UserId uint64 `db:"user_id"`
	Topic  uint64 `db:"topic"`
	Text   string `db:"text"`
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
