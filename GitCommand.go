package main

import (
    "fmt"

	log "github.com/Sirupsen/logrus"
)

type GitCommand struct {
    executor Executor
    config *Config
}

func NewGitCommand(executor Executor, config *Config) GitCommand{
    return GitCommand{executor, config}
}

func (g *GitCommand) Pull() {
	g.executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			fmt.Sprintf("Cloning %s", repo.Name), fmt.Sprintf("git clone -q --branch %s %s %s", repo.Branch, repo.URL, repo.Path), log.InfoLevel,
		}
	})
	g.executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			"", fmt.Sprintf("cd %s", repo.Path), log.PanicLevel,
		}
	})
	g.executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
            "Fixing fetch refs", "git config remote.origin.fetch +refs/heads/*:refs/remotes/origin/*", log.DebugLevel,
		}
	})
	if g.config.InitialiseGitFlow {
		g.executor.AddItem(func(repo *Repository) ExecutorItem {
			return ExecutorItem{
				"Initialising for Git Flow", "git flow init -d", log.InfoLevel,
			}
		})
	}
	g.executor.AddItem(func(repo *Repository) ExecutorItem {
		return ExecutorItem{
			"", fmt.Sprintf("git checkout %s", repo.Branch), log.PanicLevel,
		}
	})
}
