package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/go-version"
)

func main() {
	args := parseArgs()

	client, err := github.NewClient(clientOptions()...)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	rels, res, err := client.Repositories.ListReleases(context.TODO(), args.owner, args.repo, nil)
	if err != nil {
		if res.StatusCode == http.StatusNotFound {
			fmt.Printf("Repository %s/%s not found\n", args.owner, args.repo)
			os.Exit(1)
		}

		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	versions := make([]*version.Version, 0, len(rels))

	for _, rel := range rels {
		ver, err := version.NewVersion(rel.TagName)
		if err != nil {
			fmt.Printf("error parsing tag %s\n", rel.TagName)

			continue
		}

		if args.major != nil && (ver.Segments()[0] != *args.major || ver.Prerelease() != "") {
			continue
		}

		versions = append(versions, ver)
	}

	sort.Sort(version.Collection(versions))

	if len(versions) == 0 {
		fmt.Println("No releases found")
		os.Exit(1)
	}

	fmt.Println(versions[len(versions)-1])
}

func parseArgs() args {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		help()
	}

	parts := strings.Split(arguments[0], "/")
	if len(parts) != 2 {
		help()
	}

	if len(arguments) == 1 {
		return args{
			owner: parts[0],
			repo:  parts[1],
			major: nil,
		}
	}

	maj, err := strconv.Atoi(arguments[1])
	if err != nil {
		help()
	}

	return args{
		owner: parts[0],
		repo:  parts[1],
		major: &maj,
	}
}

func help() {
	e := fmt.Sprintf("Examples:\t%s %s %d\n", os.Args[0], "helm/helm", 2)
	e += fmt.Sprintf("\t\t%s %s", os.Args[0], "starship/starship")
	fmt.Printf("Usage:\t\t%s %s [%s]\n%s\n", os.Args[0], "owner/repo", "major", e)
	os.Exit(1)
}

func clientOptions() []github.ClientOptionsFunc {
	tok := os.Getenv("GITHUB_ACCESS_TOKEN")
	if tok == "" {
		return nil
	}

	return []github.ClientOptionsFunc{github.WithAuthToken(tok)}
}

type args struct {
	owner string
	repo  string
	major *int
}
