package sex


type Sex string

const (
	Unknown Sex = "unknown"
	Male    Sex = "male"
	Female  Sex = "female"
)
func (s Sex) ToString() string { return string(s) }

func (s Sex) IsUnknown() bool { return s == Unknown }

func (s Sex) IsMale() bool { return s == Male }

func (s Sex) IsFemale() bool { return s == Female }

