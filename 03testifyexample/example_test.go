package testifyexample

import (
	"03testifyexample/mock"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExampleTestSuite struct {
	suite.Suite
}

func TestExampleSuite(t *testing.T) {

	suite.Run(t, new(ExampleTestSuite))
}

func (s *ExampleTestSuite) SetupTest() {

	fmt.Println("Inside Setup test")
}

func (s *ExampleTestSuite) TearDownTest() {

	fmt.Println("Inside Teardown test")
}

func (s *ExampleTestSuite) TestExample() {

	student := &mock.FakeStudent{}

	student.GetFullNameReturns("ABC")

	actualFullName := CheckFullName(student)

	s.True(actualFullName == "My name is ::: ABC")
	s.Equal(actualFullName, "My name is ::: ABC")
	assert.Equal(s.T(), "My name is ::: ABC", actualFullName)
	assert.Equal(s.T(), student.GetFullNameCallCount(), 1)
	s.NotZero(student.GetFullNameCallCount())
	s.Equal(1, student.GetFullNameCallCount())

	fmt.Println("The result is ::: ", student.GetFullNameCallCount())
}

func (s *ExampleTestSuite) TestExample2() {

	student := &mock.FakeStudent{}

	student.GetFullNameReturns("ABC")

	actualFullName := CheckFullName(student)

	s.True(actualFullName == "My name is ::: ABC")
	s.Equal(actualFullName, "My name is ::: ABC")
	assert.Equal(s.T(), "My name is ::: ABC", actualFullName)
	assert.Equal(s.T(), student.GetFullNameCallCount(), 1)
	s.NotZero(student.GetFullNameCallCount())
	s.Equal(1, student.GetFullNameCallCount())

	fmt.Println("The result is ::: ", student.GetFullNameCallCount())
}
