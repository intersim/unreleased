package unreleased

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// import "github.com/go-git/go-git/v5"

func getCurrentDirectory() {
	// https://gist.github.com/arxdsilva/4f73d6b89c9eac93d4ac887521121120
	cmd := exec.Command("pwd")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

func getCurrentBranch() string {
	// https://golang.org/pkg/os/exec/
	// git rev-parse --abbrev-ref HEAD
	return "master"
}

func getReposFromFile() []string {
	// read from unreleased.config file (filename?)
	// get array of paths to look at, or array of repo names to look for?
	return []string{"/path/to/my_repo"}
}

func getRepoName() string {
	// basename `git rev-parse --show-toplevel`
	return "my_repo"
}

func updateMaster() {
	// (git checkout master && git pull origin master) > /dev/null 2>&1
}

func getLastTag() string {
	// last_tag="$(git describe --tag --abbrev=0)"
	return "my-last-tag"
}

func getNumCommits(repoName string) int {
	// 	num_commits="$(git rev-list $last_tag..HEAD --count)"
	return 1
}

func logUnreleasedChanges(lastTag string, repoName string, numCommits int) {
	var numCommitsMsg string = ""
	var emoji string = ""

	if numCommits > 1 {
		numCommitsMsg = fmt.Sprintf("There are %d unreleased commits!", numCommits)
		emoji = ":white_check_mark:"
	} else if numCommits == 1 {
		numCommitsMsg = fmt.Sprintf("There is %d unreleased commit!", numCommits)
		emoji = ":white_check_mark:"
	} else {
		numCommitsMsg = fmt.Sprintf("No changes to release.")
		emoji = ":x:"
	}

	fmt.Printf("%s %s The last release was: %s", emoji, numCommitsMsg, lastTag)
}
