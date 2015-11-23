package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GitCommand", func() {
	var (
		executor Executor
		repo     Repository
		cmd      GitCommand
	)

	BeforeEach(func() {
        runner := NullRunner{}
		executor = NewExecutor(&runner)
		repo = Repository{"master", "https://github.com/guzzlerio/deride", "./github.com/guzzlerio/deride", "guzzlerio/deride"}
		cmd = NewGitCommand(&executor, &config)
	})

	Describe("Pull", func() {
		Context("when local path does not exist", func() {
			BeforeEach(func() {
				repoPathExists = func(path string) bool {
					return false
				}
				cmd.Pull(&repo)
			})

			It("performs a clone", func() {
				Expect(executor.Items).Should(HaveKey("git clone"))
				Expect(executor.Items["git clone"].args).Should(ContainElement("https://github.com/guzzlerio/deride"))
			})

			It("checks out the branch", func() {
				Expect(executor.Items).Should(HaveKey("git checkout"))
				Expect(executor.Items["git checkout"].args).Should(ContainElement("master"))
			})

			Context("when configured for git flow init", func() {
				BeforeEach(func() {
				})
			})
		})

		Context("when local path does exist", func() {
			BeforeEach(func() {
				repoPathExists = func(path string) bool {
					return true
				}
				cmd.Pull(&repo)
			})

			Context("and no local changes exist", func() {
				BeforeEach(func() {
					hasLocalChanges = func() bool {
						return false
					}
				})

				It("performs a pull", func() {
					Expect(executor.Items).Should(HaveKey("git pull"))
					Expect(executor.Items["git pull"].args).Should(ContainElement("https://github.com/guzzlerio/deride"))
				})
			})

			Context("and local changes exist", func() {
				BeforeEach(func() {
					hasLocalChanges = func() bool {
						return true
					}
				})

				It("prompts to stash, pull, pop", func() {})
				Context("choosing to stash", func() {
					BeforeEach(func() {
						continueWithLocalChanges = func() bool {
							return true
						}
						cmd.Pull(&repo)
					})

					It("performs a stash", func() {
						Expect(executor.Items).Should(HaveKey("git stash"))
					})

					It("performs a pull", func() {
						Expect(executor.Items).Should(HaveKey("git pull"))
					})

					It("performs a stash pop", func() {
						Expect(executor.Items).Should(HaveKey("git stash pop"))
					})
				})

				PContext("choosing not to stash", func() {
					BeforeEach(func() {
						continueWithLocalChanges = func() bool {
							return false
						}
						cmd.Pull(&repo)
					})

					It("logs the pending changes", func() {})
					It("continues to next repo", func() {})
				})
			})
		})
	})
})
