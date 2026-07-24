package sex

type Sex string

const (
	Male    Sex = "male"
	Female  Sex = "female"
	Unknown Sex = "unknown"
)

func (s Sex) String() string {
	return string(s)
}

func (s Sex) IsMale() bool {
	return s == Male
}

func (s Sex) IsFemale() bool {
	return s == Female
}

func (s Sex) IsUnknown() bool {
	return s == Unknown
}
