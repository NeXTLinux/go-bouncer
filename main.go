package main

import "github.com/nextlinux/go-bouncer/cmd"

var (
	version   = "No version provided"
	commit    = "No commit provided"
	buildTime = "No build timestamp provided"
)

func main() {
	cmd.SetVersion(&cmd.Version{
		Version:   version,
		Commit:    commit,
		BuildTime: buildTime,
	})

	cmd.Execute()
}
