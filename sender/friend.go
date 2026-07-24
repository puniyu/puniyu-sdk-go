package sender

import (
	"github.com/puniyu/puniyu-sdk-go/sender/sex"
)

// FriendSender 好友消息发送者
type FriendSender struct {
	// ID 发送者ID
	userID string
	// Nick 用户昵称
	nick *string
	// Sex 性别，默认 Unknown
	sex sex.Sex
	// Age 年龄
	age *uint32
}

type FriendSenderBuilder struct {
	friendSender FriendSender
}

func NewFriendSenderBuilder() *FriendSenderBuilder {
	return &FriendSenderBuilder{}
}

func (fb *FriendSenderBuilder) WithUserID(userID string) *FriendSenderBuilder {
	fb.friendSender.userID = userID
	return fb
}

func (fb *FriendSenderBuilder) WithNick(nick string) *FriendSenderBuilder {
	fb.friendSender.nick = &nick
	return fb
}

func (fb *FriendSenderBuilder) WithSex(sex sex.Sex) *FriendSenderBuilder {
	fb.friendSender.sex = sex
	return fb
}

func (fb *FriendSenderBuilder) WithAge(age uint32) *FriendSenderBuilder {
	fb.friendSender.age = &age
	return fb
}

func (fb *FriendSenderBuilder) Build() *FriendSender {
	return &fb.friendSender
}

func (f *FriendSender) UserID() string {
	return f.userID
}

func (f *FriendSender) Nick() *string {
	return f.nick
}

func (f *FriendSender) Sex() sex.Sex {
	return f.sex
}

func (f *FriendSender) Age() *uint32 {
	return f.age
}
