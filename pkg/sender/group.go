package sender

import (
	"github.com/puniyu/puniyu-sdk-go/pkg/role"
	"github.com/puniyu/puniyu-sdk-go/pkg/sex"
)

// GroupSender 群组消息发送者
type GroupSender struct {
	data groupSenderData
}

// groupSenderData 群组消息发送者内部数据
type groupSenderData struct {
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

// GroupSenderBuilder 群组消息发送者构建器
type GroupSenderBuilder struct {
	data groupSenderData
}

// NewGroupSenderBuilder 创建群组消息发送者构建器
func NewGroupSenderBuilder() *GroupSenderBuilder {
	return &GroupSenderBuilder{
		data: groupSenderData{},
	}
}

// SetUserID 设置发送者ID
func (b *GroupSenderBuilder) SetUserID(userID string) *GroupSenderBuilder {
	b.data.UserID = userID
	return b
}

// SetNick 设置用户昵称
func (b *GroupSenderBuilder) SetNick(nick string) *GroupSenderBuilder {
	b.data.Nick = &nick
	return b
}

// SetSex 设置性别
func (b *GroupSenderBuilder) SetSex(sex sex.Sex) *GroupSenderBuilder {
	b.data.Sex = sex
	return b
}

// SetAge 设置年龄
func (b *GroupSenderBuilder) SetAge(age uint32) *GroupSenderBuilder {
	b.data.Age = &age
	return b
}

// SetRole 设置角色
func (b *GroupSenderBuilder) SetRole(role role.Role) *GroupSenderBuilder {
	b.data.Role = role
	return b
}

// SetCard 设置名片
func (b *GroupSenderBuilder) SetCard(card string) *GroupSenderBuilder {
	b.data.Card = &card
	return b
}

// SetLevel 设置等级
func (b *GroupSenderBuilder) SetLevel(level uint32) *GroupSenderBuilder {
	b.data.Level = &level
	return b
}

// SetTitle 设置头衔
func (b *GroupSenderBuilder) SetTitle(title string) *GroupSenderBuilder {
	b.data.Title = &title
	return b
}

// Build 构建群组消息发送者
func (b *GroupSenderBuilder) Build() *GroupSender {
	return &GroupSender{data: b.data}
}

// UserID 获取发送者ID
func (g *GroupSender) UserID() string { return g.data.UserID }

// Sex 获取性别
func (g *GroupSender) Sex() sex.Sex { return g.data.Sex }

// Nick 获取用户昵称
func (g *GroupSender) Nick() *string { return g.data.Nick }

// Age 获取年龄
func (g *GroupSender) Age() *uint32 { return g.data.Age }
