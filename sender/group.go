package sender

import (
	"github.com/puniyu/puniyu-sdk-go/sender/role"
	"github.com/puniyu/puniyu-sdk-go/sender/sex"
)

// GroupSender 群组消息发送者
type GroupSender struct {
	// userID 发送者ID
	userID string
	// nick 用户昵称
	nick *string
	// sex 性别，默认 Unknown
	sex sex.Sex
	// age 年龄
	age *uint32
	// role 角色
	role role.Role
	// card 名片
	card *string
	// level 等级
	level *uint32
	// title 头衔
	title *string
}

type GroupSenderBuilder struct {
	groupSender GroupSender
}

func NewGroupSenderBuilder() *GroupSenderBuilder {
	return &GroupSenderBuilder{}
}

func (gb *GroupSenderBuilder) WithUserID(userID string) *GroupSenderBuilder {
	gb.groupSender.userID = userID
	return gb
}

func (gb *GroupSenderBuilder) WithNick(nick string) *GroupSenderBuilder {
	gb.groupSender.nick = &nick
	return gb
}

func (gb *GroupSenderBuilder) WithSex(sex sex.Sex) *GroupSenderBuilder {
	gb.groupSender.sex = sex
	return gb
}

func (gb *GroupSenderBuilder) WithAge(age uint32) *GroupSenderBuilder {
	gb.groupSender.age = &age
	return gb
}

func (gb *GroupSenderBuilder) WithRole(role role.Role) *GroupSenderBuilder {
	gb.groupSender.role = role
	return gb
}

func (gb *GroupSenderBuilder) WithCard(card *string) *GroupSenderBuilder {
	gb.groupSender.card = card
	return gb
}

func (gb *GroupSenderBuilder) WithLevel(level uint32) *GroupSenderBuilder {
	gb.groupSender.level = &level
	return gb
}

func (gb *GroupSenderBuilder) WithTitle(title string) *GroupSenderBuilder {
	gb.groupSender.title = &title
	return gb
}

func (gb *GroupSenderBuilder) Build() *GroupSender {
	return &gb.groupSender
}

func (g *GroupSender) UserID() string {
	return g.userID
}

func (g *GroupSender) Nick() *string {
	return g.nick
}

func (g *GroupSender) Sex() sex.Sex {
	return g.sex
}

func (g *GroupSender) Age() *uint32 {
	return g.age
}
