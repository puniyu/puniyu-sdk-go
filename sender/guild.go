package sender

import (
	"github.com/puniyu/puniyu-sdk-go/sender/role"
	"github.com/puniyu/puniyu-sdk-go/sender/sex"
)

type GuildSender struct {
	UserID string    `json:"userID"`
	Nick   *string   `json:"nick,omitempty"`
	Sex    sex.Sex   `json:"sex"`
	Age    *uint32   `json:"age,omitempty"`
	Role   role.Role `json:"role"`
	Card   *string   `json:"card,omitempty"`
	Level  *uint32   `json:"level,omitempty"`
	Title  *string   `json:"title,omitempty"`
}
