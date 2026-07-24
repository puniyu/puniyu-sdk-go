package sender

import (
	"github.com/puniyu/puniyu-sdk-go/sender/role"
	"github.com/puniyu/puniyu-sdk-go/sender/sex"
)

// GroupSender 群组消息发送者
type GroupSender struct {
	// UserID 发送者ID
	UserID string `json:"user_id"`
	// Nick 用户昵称
	Nick *string `json:"nick,omitempty"`
	// Sex 性别，默认 Unknown
	Sex sex.Sex `json:"sex"`
	// Age 年龄
	Age *uint32 `json:"age,omitempty"`
	// Role 角色
	Role role.Role `json:"role"`
	// Card 名片
	Card *string `json:"card,omitempty"`
	// Level 等级
	Level *uint32 `json:"level,omitempty"`
	// Title 头衔
	Title *string `json:"title,omitempty"`
}
