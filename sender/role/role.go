package role

type Role string

const (
	Owner   Role = "owner"
	Admin   Role = "admin"
	Member  Role = "member"
	Unknown Role = "unknown"
)

func (r Role) String() string {
	return string(r)
}

func (r Role) IsOwner() bool {
	return r == Owner
}

func (r Role) IsAdmin() bool {
	return r == Admin
}

func (r Role) IsMember() bool {
	return r == Member
}

func (r Role) IsUnknown() bool {
	return r == Unknown
}
