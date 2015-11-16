package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GitCommand", func() {
    var executor Executor

    BeforeEach(func(){
        executor = NewExecutor()
    })

	Describe("Pull", func() {
        Context("when local path does not exist", func() {
            BeforeEach(func() {
                cmd := NewGitCommand(&executor, &config)

                cmd.Pull()
            })

            It("performs a clone", func(){
                Expect(executor.Items).Should(HaveKey("git clone"))
            })
        })

        PContext("when local path does exist", func() {
            It("performs a pull", func(){})
        })
	})
})
