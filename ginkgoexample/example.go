package ginkgoexample

import "ginkgoexample/service"

func CheckFullName(student service.Student) string {

	name := student.GetFullName()

	return "My name is ::: " + name

}

var SomeUtilityFunc = service.SomeUtilityMethod

func CheckAge(student service.Student) string {

	age := student.GetAge()
	ageStr := SomeUtilityFunc(age)

	return ageStr
}
