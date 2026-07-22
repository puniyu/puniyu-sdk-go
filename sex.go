package puniyu

type Sex string

const (
	SexUnknown Sex = "unknown"
	SexMale    Sex = "male"
	SexFemale  Sex = "female"
)

func (s Sex) ToString() string { return string(s) }

func (s Sex) IsUnknown() bool { return s == SexUnknown }

func (s Sex) IsMale() bool { return s == SexMale }

func (s Sex) IsFemale() bool { return s == SexFemale }
