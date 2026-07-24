package sender

import (
	"github.com/puniyu/puniyu-sdk-go/sender/sex"
	"github.com/vmihailenco/msgpack/v5"
)

// FriendSender 好友消息发送者
type FriendSender struct {
	// ID 发送者ID
	iUserID string
	// Nick 用户昵称
	iNick *string
	// Sex 性别，默认 Unknown
	iSex sex.Sex
	// Age 年龄
	iAge *uint32
}

var _ msgpack.CustomEncoder = (*FriendSender)(nil)
var _ msgpack.CustomDecoder = (*FriendSender)(nil)

func (f *FriendSender) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.EncodeMulti(f.iUserID, f.iNick, f.iSex, f.iAge)
}

func (f *FriendSender) DecodeMsgpack(dec *msgpack.Decoder) error {
	var sexValue int32
	err := dec.DecodeMulti(&f.iUserID, &f.iNick, &sexValue, &f.iAge)
	if err != nil {
		return err
	}
	f.iSex = sex.Sex(sexValue)
	return nil
}

func (f *FriendSender) Marshal() ([]byte, error) {
	return msgpack.Marshal(f)
}

func UnmarshalFriendSender(data []byte) (*FriendSender, error) {
	var fs FriendSender
	err := msgpack.Unmarshal(data, &fs)
	if err != nil {
		return nil, err
	}
	return &fs, nil
}

type FriendSenderBuilder struct {
	friendSender FriendSender
}

func NewFriendSenderBuilder() *FriendSenderBuilder {
	return &FriendSenderBuilder{}
}

func (fb *FriendSenderBuilder) WithUserID(userID string) *FriendSenderBuilder {
	fb.friendSender.iUserID = userID
	return fb
}

func (fb *FriendSenderBuilder) WithNick(nick string) *FriendSenderBuilder {
	fb.friendSender.iNick = &nick
	return fb
}

func (fb *FriendSenderBuilder) WithSex(sex sex.Sex) *FriendSenderBuilder {
	fb.friendSender.iSex = sex
	return fb
}

func (fb *FriendSenderBuilder) WithAge(age uint32) *FriendSenderBuilder {
	fb.friendSender.iAge = &age
	return fb
}

func (fb *FriendSenderBuilder) Build() *FriendSender {
	return &fb.friendSender
}

func (f *FriendSender) UserID() string {
	return f.iUserID
}

func (f *FriendSender) Nick() *string {
	return f.iNick
}

func (f *FriendSender) Sex() sex.Sex {
	return f.iSex
}

func (f *FriendSender) Age() *uint32 {
	return f.iAge
}
