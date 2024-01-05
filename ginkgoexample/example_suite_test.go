package ginkgoexample_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestExampleSuite(t *testing.T) {

	RegisterFailHandler(Fail)   // this line glues code connecting Ginkgo to Gomega.
	RunSpecs(t, "ExampleSuite") // This call tells Ginkgo to start the test suite in this package.
}
