package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

type GitCommand struct {
	executor *Executor
	config   *Config
}

func NewGitCommand(executor *Executor, config *Config) GitCommand {
	return GitCommand{executor, config}
}

func (g *GitCommand) Pull(repo *Repository) {
	if repoPathExists(repo.Path) {
		g.executor.AddItem(ExecutorItem{fmt.Sprintf("Working on %s", repo.Path), "cd", []string{repo.Path}, log.InfoLevel})
		if hasLocalChanges() {
			if continueWithLocalChanges() {
                g.executor.AddItem(ExecutorItem{"Stashing changes", "git stash", []string{"-q"}, log.InfoLevel})
				g.executor.AddItem(ExecutorItem{"Pulling", "git pull", []string{"-q", repo.URL, repo.Path}, log.InfoLevel})
				g.executor.AddItem(ExecutorItem{"Popping the stash", "git stash pop", []string{"-q"}, log.InfoLevel})
			} else {
            }
		} else {
			g.executor.AddItem(ExecutorItem{fmt.Sprintf("Pulling %s", repo.Name), "git pull", []string{"-q", repo.URL, repo.Path}, log.InfoLevel})
		}
	} else {
        g.Clone(repo)
	}
}

func (g *GitCommand) Clone(repo *Repository) {
    g.executor.AddItem(ExecutorItem{fmt.Sprintf("Cloning %s into %s", repo.Name, repo.Path), "git clone", []string{"-q", repo.URL, repo.Path}, log.InfoLevel})
    g.executor.AddItem(ExecutorItem{fmt.Sprintf("cd %s", repo.Path), "cd", []string{repo.Path}, log.DebugLevel})
    if g.config.SetupGitFlow {
        g.executor.AddItem(ExecutorItem{"Initialising for Git Flow", "git flow", []string{"init", "-d"}, log.InfoLevel})
    }
    g.executor.AddItem(ExecutorItem{fmt.Sprintf("Checkout %s", repo.Branch), "git checkout", []string{repo.Branch}, log.InfoLevel})
}

var repoPathExists = func(path string) bool {
	if _, err := os.Stat(path); os.IsExist(err) {
		return true
	}
	return false
}

var hasLocalChanges = func() bool {
	//TODO figure out how to detect local changes
	return false
}

var continueWithLocalChanges = func() bool {
    //TODO prompt user if we should stash/pull/pop
	return true
}
