package ginkgoexample_test

import (
	ginkgoexample "02ginkgoexample"
	"02ginkgoexample/mock"
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// F in front of It,Context,Describe or any block method means focus - it will run only that block in the suite. Ex - FIt()
// X in front of It,Context,Describe or any block method means exclude that block - it will skip that block in the suite. Ex - XIt()
var _ = Describe("Example", func() {

	BeforeEach(func() {
		fmt.Println("Inside before Each")
	})

	AfterEach(func() {
		fmt.Println("Inside after each")
	})

	Describe("CheckFullName", func() {

		Context("Student Name", func() {

			It("Should return full name", func() {

				student := &mock.FakeStudent{}
				student.GetFullNameReturns("Piyush Kumar")

				receivedName := ginkgoexample.CheckFullName(student)

				Expect(receivedName).To(Equal("My name is ::: Piyush Kumar"))

			})

		})

	})

	Describe("CheckAge", func() {

		When("Student Age", func() {

			It("Should return String age", func() {

				student := &mock.FakeStudent{}
				student.GetAgeReturns(20)

				old := ginkgoexample.SomeUtilityFunc
				defer func() {
					ginkgoexample.SomeUtilityFunc = old
				}()

				ginkgoexample.SomeUtilityFunc = func(num int) string { // Here I am mocking package method. Refer : https://stackoverflow.com/a/42952311

					return "My age is 34"
				}
				actualAgeStr := ginkgoexample.CheckAge(student)

				Expect(actualAgeStr).To(Equal("My age is 34"))

			})

		})

	})
})
