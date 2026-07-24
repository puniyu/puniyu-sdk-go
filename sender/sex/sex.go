package sex

type Sex struct {
	value sexType
}

type sexType = uint8

const (
	male sexType = iota
	female
	unknown
)

func NewMale() *Sex {
	return &Sex{male}
}

func NewFemale() *Sex {
	return &Sex{female}
}

func NewUnknown() *Sex {
	return &Sex{unknown}
}

func (s *Sex) IsMale() bool {
	return s.value == male
}

func (s *Sex) IsFemale() bool {
	return s.value == female
}

func (s *Sex) IsUnknown() bool {
	return s.value == unknown
}
