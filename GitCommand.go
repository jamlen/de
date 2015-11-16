package main

import (
    "fmt"

	log "github.com/Sirupsen/logrus"
)

type GitCommand struct {
    executor *Executor
    config *Config
}

func NewGitCommand(executor *Executor, config *Config) GitCommand{
    return GitCommand{executor, config}
}

func (g *GitCommand) Pull() {
	g.executor.AddItem("git clone", func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			fmt.Sprintf("Cloning %s", repo.Name), "git clone", []string{"-q","--branch", repo.Branch, repo.URL, repo.Path}, log.InfoLevel,
		}
	})
	g.executor.AddItem("cd", func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			fmt.Sprintf("cd %s", repo.Path), "cd", []string{repo.Path}, log.DebugLevel,
		}
	})
	g.executor.AddItem("git config", func(repo *Repository) ExecutorItem {
		return ExecutorItem{
            "Fixing fetch refs", "git config", []string{"remote.origin.fetch", "+refs/heads/*:refs/remotes/origin/*"}, log.DebugLevel,
		}
	})
	if g.config.InitialiseGitFlow {
		g.executor.AddItem("git flow", func(repo *Repository) ExecutorItem {
			return ExecutorItem{
				"Initialising for Git Flow", "git flow", []string{"init", "-d"}, log.InfoLevel,
			}
		})
	}
	g.executor.AddItem("git checkout", func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			fmt.Sprintf("Checkout %s", repo.Branch), "git checkout", []string{repo.Branch}, log.InfoLevel,
		}
	})
}
