package role

type Role struct {
	value roleType
}

type roleType uint8

const (
	owner roleType = iota
	admin
	member
	unknown
)

func NewOwner() *Role {
	return &Role{owner}
}

func NewAdmin() *Role {
	return &Role{admin}
}

func NewMember() *Role {
	return &Role{member}
}

func (r *Role) IsOwner() bool {
	return r.value == owner
}

func (r *Role) IsAdmin() bool {
	return r.value == admin
}

func (r *Role) IsMember() bool {
	return r.value == member
}

func (r *Role) IsUnknown() bool {
	return r.value == unknown
}
