package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/go-version"
	"golang.org/x/oauth2"
)

func main() {
	owner, repo, major := parseArgs()

	client := github.NewClient(httpClient())

	rels, res, err := client.Repositories.ListReleases(context.TODO(), owner, repo, nil)
	if err != nil {
		if res.StatusCode == http.StatusNotFound {
			fmt.Printf("Repository %s/%s not found\n", owner, repo)
			os.Exit(1)
		}
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	var versions []*version.Version

	for _, rel := range rels {
		v, err := version.NewVersion(*rel.TagName)
		if err != nil {
			fmt.Printf("error parsing tag %s\n", *rel.TagName)
			continue
		}

		if v.Segments()[0] != major || v.Prerelease() != "" {
			continue
		}

		versions = append(versions, v)
	}

	sort.Sort(version.Collection(versions))

	if len(versions) == 0 {
		fmt.Printf("No major %d versions found\n", major)
		os.Exit(1)
	}

	fmt.Println(versions[len(versions)-1])
}

func parseArgs() (string, string, int) {
	args := os.Args[1:]

	if len(args) != 2 {
		help()
	}

	parts := strings.Split(args[0], "/")
	if len(parts) != 2 {
		help()
	}

	m, err := strconv.Atoi(args[1])
	if err != nil {
		help()
	}

	return parts[0], parts[1], m
}

func help() {
	e := fmt.Sprintf("Example: %s %s %d", os.Args[0], "helm/helm", 2)
	fmt.Printf("Usage: %s %s %s\n\t%s\n", os.Args[0], "owner/repo", "major", e)
	os.Exit(1)
}

func httpClient() *http.Client {
	t := os.Getenv("GITHUB_ACCESS_TOKEN")
	if t == "" {
		return http.DefaultClient
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: t},
	)
	return oauth2.NewClient(ctx, ts)
}
