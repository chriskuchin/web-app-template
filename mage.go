//+build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default = Run
)

const (
	dockerImage = "registry.home.cksuperman.com/grpfit"
)

func Run() error {
	fmt.Println("Running API")
	return sh.RunV("go", "run", "main.go", "server")
}

func Build() error {
	mg.Deps(Tidy)
	return sh.Run("docker", "build", "-t", dockerImage, "./")
}

func Publish() error {
	mg.Deps(Build)
	return sh.Run("docker", "push", dockerImage)
}

func Tidy() error {
	return sh.Run("go", "mod", "tidy")
}
