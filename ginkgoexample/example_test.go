package ginkgoexample_test

import (
	"ginkgoexample"
	"ginkgoexample/mock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example", func() {

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

		Context("Student Age", func() {

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
