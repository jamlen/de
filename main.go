package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbosity = 1
	logLevel  = log.FatalLevel
	config    = Config{}

	de      = kingpin.New("de", "Devenv helper")
	all     = de.Flag("all", "Apply to all").Bool()
	verbose = de.Flag("verbose", "Verbose output").Short('v').Action(verbosityCounter).Bool()

	gitCommand = de.Command("git", "Perform git commands")
	gitPull    = gitCommand.Command("pull", "Pull or clone projects and modules")
	gitStatus  = gitCommand.Command("status", "Git status on projects and modules")
	//gitProject    = gitCommand.Arg("project", "Which project to run against").String()

    // only add if start command provided in config
	appCommand = de.Command("app", "Perform app commands")
	appStart   = appCommand.Command("start", "Start an app")
	appStop    = appCommand.Command("stop", "Stop an app")

    // only add if npm is detected
	npmCommand = de.Command("npm", "Perform npm commands")
	npmLink    = npmCommand.Command("link", "Npm link modules into projects")
	npmInstall = npmCommand.Command("install", "Npm install modules")

    // only add if bower is detected
	bowerCommand = de.Command("bower", "Perform bower commands")
	bowerLink    = bowerCommand.Command("link", "Bower link modules into projects")
	bowerInstall = bowerCommand.Command("install", "Bower install modules")

    // only add if test command provided in config
	testCommand = de.Command("test", "Perform test commands")
)

func verbosityCounter(c *kingpin.ParseContext) error {
	verbosity++
	switch verbosity {
	case 1:
		log.SetLevel(log.WarnLevel)
	case 2:
		log.SetLevel(log.InfoLevel)
	case 3:
		log.SetLevel(log.DebugLevel)
	}

	return nil
}

func main() {
    runner := ShellRunner{}
	executor := NewExecutor(&runner)

	kingpin.Version("0.1.0")
	switch kingpin.MustParse(de.Parse(os.Args[1:])) {
	case gitPull.FullCommand():
		g := NewGitCommand(&executor, &config)
		executor.Command = g.Pull
	}
	executor.Execute(getReposToAction())
}

func getReposToAction() []Repository {
	// load config.json file and parse into Repositories
	return []Repository{{"master", "https://github.com/guzzlerio/deride", "./github.com/guzzlerio/deride", "guzzlerio/deride"}}
}

type Config struct {
	InitialiseGitFlow bool
}

type Repository struct {
	Branch string
	URL    string
	Path   string
	Name   string
}

