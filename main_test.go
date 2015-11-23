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
    . "github.com/onsi/gomega"

    "crypto/md5"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "time"
)

var _ = Describe("Integration test", func() {
    var (
        tmp    string
        err    error
        output string
    )

    BeforeEach(func() {
        if tmp, err = createTemporaryDir(); err != nil {
            Expect(err).ShouldNot(HaveOccurred())
        }
    })

    Describe("git commands", func() {
        Describe("git pull", func() {
            Context("when path does not exist", func() {
                BeforeEach(func() {
                    output = Invoke("git", "pull", "--all")
                })

                It("clones projects and modules", func() {
                    Expect(output).Should(ContainSubstring("Cloning guzzlerio/deride"))
                })
            })

            PContext("when path exists", func() {
                Context("with no local changes", func() {
                    BeforeEach(func() {
                        // checkout one of the repos
                        p := tmp + "/guzzlerio/deride"
                        //TODO create path in tmp dir as p
                        cmd := exec.Command("git", "-q git@github.com:guzzlerio/deride.git " + p)
                        _, err = cmd.CombinedOutput()
                        Expect(err).ShouldNot(HaveOccurred())
                        output = Invoke("git", "pull", "--all")
                    })

                    It("performs the pull", func() {
                        Expect(output).Should(ContainSubstring("Pulling guzzlerio/deride"))
                    })
                })

                Context("with local changes", func() {
                    BeforeEach(func() {
                        // checkout one of the repos
                        // make a local change
                        output = Invoke("git", "pull", "--all")
                    })
                    Context("choosing to stash/pull/pop", func() {})
                    Context("choosing to skip", func() {})
                })
            })
        })

        Describe("git status", func() {
        })
    })

    AfterEach(func() {
        if err := os.RemoveAll(tmp); err != nil {
            Expect(err).ShouldNot(HaveOccurred())
        }
    })
})

func Invoke(args ...string) string {
    // read test-config.yml
    //data, _ := ioutil.ReadFile(filepath.Abs("./test-config.yml"))

    // replace TEMP_DIR with tmp
    // write out to tmp/test-config.yml
    exePath, err := filepath.Abs("./de")
    Expect(err).To(BeNil())

    // cd into tmp
    cmd := exec.Command(exePath, append(args, "-c", "test-config.yml")...)
    output, err := cmd.CombinedOutput()
    if len(output) > 0 {
        log.Println(fmt.Sprintf("%s", output))
    }
    Expect(err).To(BeNil())
    return string(output)
}

func createTemporaryDir() (string, error) {
    stamp := string(time.Now().UnixNano())
    hashed := md5.Sum([]byte(stamp))
    return ioutil.TempDir(os.TempDir(), fmt.Sprintf("%x", hashed))
}
