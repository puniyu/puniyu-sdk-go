package role

type Role string

const (
	Unknown Role = "unknown"
	Member  Role = "member"
	Admin   Role = "admin"
	Owner   Role = "owner"
)
func (r Role) ToString() string { return string(r) }
func (r Role) IsUnknown() bool { return r == Unknown }
func (r Role) IsMember() bool  { return r == Member }
func (r Role) IsAdmin() bool   { return r == Admin }
func (r Role) IsOwner() bool   { return r == Owner }

