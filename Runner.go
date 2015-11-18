package main

type Runner interface {
    Run(string, []string) error
}

type NullRunner struct {}

func (r *NullRunner) Run(string, []string) error {
    return nil
}
