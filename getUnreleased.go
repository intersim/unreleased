package chchchanges

// import "github.com/go-git/go-git/v5"

func getUnreleased() {
	// get current directory
	// get current branch

	// get repos from config file
	// for each repo:
	// cd to it
	// get unreleased changes
	// log out changes info

	// cd back to current directory
	// change back to current branch
}

func getCurrentDirectory() {
	// https://gist.github.com/arxdsilva/4f73d6b89c9eac93d4ac887521121120
}

func getCurrentBranch() {
	// https://golang.org/pkg/os/exec/
}

func getReposFromFile() {
	// read from chchchanges.config file (filename?)
	// get array of paths to look at, or array of repo names to look for?
}

func getUnreleasedChanges() {
	// function get_unreleased_git_helper {
	// 	(git checkout master && git pull origin master) > /dev/null 2>&1
	// 	repo_name=$1
	// 	last_tag="$(git describe --tag --abbrev=0)"
	// 	num_commits="$(git rev-list $last_tag..HEAD --count)"
	// 	get_unreleased_log_helper $repo_name $num_commits $last_tag
	// }
}

func logUnreleasedChanges() {
	// function get_unreleased_log_helper {
	// 	repo_name=$1
	// 	num_commits=$2
	// 	last_tag=$3
	// 	if [ "$num_commits" -gt "1" ]; then
	// 		echo ":white_check_mark: $repo_name: There are $num_commits unreleased commits! The last release was: $last_tag"
	// 	elif [ "$num_commits" == "1" ]; then
	// 		echo ":white_check_mark: $repo_name: There is $num_commits unreleased commit! The last release was: $last_tag"
	// 	else
	// 		echo ":x: $repo_name: No changes to release. The last release was: $last_tag"
	// 	fi
	// }
}
