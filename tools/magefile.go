//+build mage

package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	semver "github.com/blang/semver/v4"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/go-swagger/go-swagger/cmd/swagger/commands/generate"
	_ "github.com/golangci/golangci-lint/pkg/commands"
	"github.com/goreleaser/goreleaser/cmd"
	flags "github.com/jessevdk/go-flags"
	"github.com/magefile/mage/mg"
	"github.com/nolte/plumbing/cmd/golang"

	// mage:import
	_ "github.com/nolte/plumbing/cmd/kind"

	plumbing "github.com/nolte/plumbing/pkg"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

func (Build) GoRelease() {
	os.Chdir("../")
	defer os.Chdir("./tools")
	args := []string{"build", "--rm-dist", "--snapshot"}
	cmd.Execute(
		buildVersion(version, commit, date, builtBy),
		os.Exit,
		args,
	)
}

func buildVersion(version, commit, date, builtBy string) string {
	var result = version
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
