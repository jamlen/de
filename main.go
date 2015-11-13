package main

import (
	"fmt"
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
	gitPull    = gitCommand.Command("pull", "Pull projects and modules")
	gitStatus  = gitCommand.Command("status", "Git status on projects and modules")

	appCommand = de.Command("app", "Perform app commands")
	appStart   = appCommand.Command("start", "Start an app")
	appStop    = appCommand.Command("stop", "Stop an app")

	npmCommand = de.Command("npm", "Perform npm commands")
	npmLink    = npmCommand.Command("link", "Npm link modules into projects")
	npmInstall = npmCommand.Command("install", "Npm install modules")

	bowerCommand = de.Command("bower", "Perform bower commands")
	bowerLink    = bowerCommand.Command("link", "Bower link modules into projects")
	bowerInstall = bowerCommand.Command("install", "Bower install modules")

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
	executor := Executor{}

	kingpin.Version("0.1.0")
	switch kingpin.MustParse(de.Parse(os.Args[1:])) {
	case gitPull.FullCommand():
		GitPull(&executor)
	}
	// load config.json file and parse into Repositories
	repos := getReposToAction()
	executor.Execute(repos)
}

func Log(level log.Level, args ...interface{}) {
	switch level {
	case log.DebugLevel:
		log.Debug(args)
	case log.InfoLevel:
		log.Info(args)
	case log.WarnLevel:
		log.Warn(args)
	}
}

func getReposToAction() []Repository {
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

type ExecutorItem struct {
	description string
	command     string
	logLevel    log.Level
}

type ExecuteItemBuilder func(repo *Repository) ExecutorItem

type Executor struct {
	Items []ExecuteItemBuilder
}

func (e *Executor) Execute(repos []Repository) {
	log.Debugf("execute %d\n", verbosity)
	for _, repo := range repos {
		log.Debugf("repo: %+v\n", repo)
		for _, fn := range e.Items {
			item := fn(&repo)
			//fmt.Printf("  item: %+v\n", item)
			//fmt.Printf("%+v\n", item)
            Log(item.logLevel, item.description)
		}
	}
}

func (e *Executor) AddItem(fn ExecuteItemBuilder) {
	e.Items = append(e.Items, fn)
}

func GitPull(executor *Executor) {
	executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			fmt.Sprintf("Cloning %s", repo.Name), fmt.Sprintf("git clone -q --branch %s %s %s", repo.Branch, repo.URL, repo.Path), log.InfoLevel,
		}
	})
	executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			"", fmt.Sprintf("cd %s", repo.Path), log.PanicLevel,
		}
	})
	executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			"Fixing fetch refs", "git config remote.origin.fetch +refs/heads/*:refs/remotes/origin/*", log.DebugLevel,
		}
	})
	if config.InitialiseGitFlow {
		executor.AddItem(func(repo *Repository) ExecutorItem {
			return ExecutorItem{
				"Initialising for Git Flow", "git flow init -d", log.InfoLevel,
			}
		})
	}
	executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			"", fmt.Sprintf("git checkout %s", repo.Branch), log.PanicLevel,
		}
	})
}
