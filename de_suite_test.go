package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "De Suite")
}

