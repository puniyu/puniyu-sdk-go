package sender

import (
	"encoding/json"

	"github.com/puniyu/puniyu-sdk-go/pkg/sex"
)

// FriendSender 好友消息发送者
type FriendSender struct {
	data friendSenderData
}

// friendSenderData 好友消息发送者内部数据
type friendSenderData struct {
	// UserID 发送者ID
	UserID string `json:"user_id"`
	// Nick 用户昵称
	Nick *string `json:"nick,omitempty"`
	// Sex 性别，默认 Unknown
	Sex sex.Sex `json:"sex"`
	// Age 年龄
	Age *uint32 `json:"age,omitempty"`
}

// FriendSenderBuilder 好友消息发送者构建器
type FriendSenderBuilder struct {
	data friendSenderData
}

// NewFriendSenderBuilder 创建好友消息发送者构建器
func NewFriendSenderBuilder() *FriendSenderBuilder {
	return &FriendSenderBuilder{
		data: friendSenderData{},
	}
}

// SetUserID 设置发送者ID
func (b *FriendSenderBuilder) SetUserID(userID string) *FriendSenderBuilder {
	b.data.UserID = userID
	return b
}

// SetNick 设置用户昵称
func (b *FriendSenderBuilder) SetNick(nick string) *FriendSenderBuilder {
	b.data.Nick = &nick
	return b
}

// SetSex 设置性别
func (b *FriendSenderBuilder) SetSex(sex sex.Sex) *FriendSenderBuilder {
	b.data.Sex = sex
	return b
}

// SetAge 设置年龄
func (b *FriendSenderBuilder) SetAge(age uint32) *FriendSenderBuilder {
	b.data.Age = &age
	return b
}

// Build 构建好友消息发送者
func (b *FriendSenderBuilder) Build() *FriendSender {
	return &FriendSender{data: b.data}
}

func (f *FriendSender) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.data)
}

func (f *FriendSender) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &f.data)
}

// UserID 获取发送者ID
func (f *FriendSender) UserID() string { return f.data.UserID }

// Sex 获取性别
func (f *FriendSender) Sex() sex.Sex { return f.data.Sex }

// Nick 获取用户昵称
func (f *FriendSender) Nick() *string { return f.data.Nick }

// Age 获取年龄
func (f *FriendSender) Age() *uint32 { return f.data.Age }
