package puniyu

type Role string

const (
	RoleUnknown Role = "unknown"
	RoleMember  Role = "member"
	RoleAdmin   Role = "admin"
	RoleOwner   Role = "owner"
)

func (r Role) ToString() string { return string(r) }
func (r Role) IsUnknown() bool  { return r == RoleUnknown }
func (r Role) IsMember() bool   { return r == RoleMember }
func (r Role) IsAdmin() bool    { return r == RoleAdmin }
func (r Role) IsOwner() bool    { return r == RoleOwner }
