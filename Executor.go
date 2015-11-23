package main

import (
	log "github.com/Sirupsen/logrus"
)

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

type ExecutorItem struct {
	description string
	command     string
	args        []string
	logLevel    log.Level
}

type Executor struct {
	Items   map[string]ExecutorItem
	Command func(repo *Repository)
	Runner  *Runner
}

func NewExecutor(runner Runner) Executor {
	items := make(map[string]ExecutorItem)
	//set log.formatter to custom formatter based upon text logger
	//  if tty attached output should be coloured:
	//       {colour}de{/colour} - {msg} {fields}
	//  else based upon json
	//       time={timestamp} level={level} msg={msg} fields={fields}
	e := Executor{Items: items, Runner: &runner}
	return e
}

func (e *Executor) Execute(repos []Repository) {
	for _, repo := range repos {
		log.Debugf("repo: %+v\n", repo)
		e.Command(&repo)
		for cmd, item := range e.Items {
			Log(item.logLevel, item.description, cmd)
		}
	}
}

func (e *Executor) AddItem(item ExecutorItem) {
	//TODO check if key exists
	e.Items[item.command] = item
}
