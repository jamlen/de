package main
/*
  Do an integration test, to do that we will need:
  * use a temp folder

  * run de git pull
    * expect clone
  * rewind to HEAD~1
  * run de git pull
    * expect pulled 1 commit

  * rewind to HEAD~1
  * make local change
  * run de git pull
    * expect prompt for stash, pull, pop
  * respond Yes
    * expect HEAD commit and local changes still present
*/

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Integration test", func() {
})
