package service_test

import (
	"02ginkgoexample/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSampleServiceSuite(t *testing.T) {

	RegisterFailHandler(Fail)         // this line glues code connecting Ginkgo to Gomega.
	RunSpecs(t, "SampleServiceSuite") // This call tells Ginkgo to start the test suite in this package.
}

var _ = Describe("Student", func() {

	Context("Test GetFullName", func() {

		It("Should return full name", func() {

			student := service.NewStudent()

			actualFullName := student.GetFullName()

			Expect(actualFullName).To(Equal("PiyushKumar"))

		})
	})

})
