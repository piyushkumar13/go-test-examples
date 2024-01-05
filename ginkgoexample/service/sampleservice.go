package service

//go:generate counterfeiter -o ../mock . Student
type Student interface {
	GetFullName() string
	GetAge() int
	IsFullTime() bool
}

type student struct {
	firstName  string
	lastName   string
	age        int
	isFullTime bool
}

func NewStudent() Student {
	return &student{
		firstName:  "Piyush",
		lastName:   "Kumar",
		age:        34,
		isFullTime: true,
	}
}

func (s *student) GetFullName() string {

	return s.firstName + s.lastName

}

func (s *student) GetAge() int {

	return s.age
}

func (s *student) IsFullTime() bool {

	return s.isFullTime
}

func SomeUtilityMethod(num int) string {

	return string(rune(num))
}
