package puniyu

// Sender 消息发送者接口
type Sender interface {
	// UserID 获取发送者ID
	UserID() string
	// Nick 获取发送者昵称
	Nick() *string
	// Sex 获取发送者性别
	Sex() Sex
	// Age 获取发送者年龄
	Age() *uint32
}
