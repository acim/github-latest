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
	args := parseArgs()

	client := github.NewClient(httpClient())

	rels, res, err := client.Repositories.ListReleases(context.TODO(), args.owner, args.repo, nil)
	if err != nil {
		if res.StatusCode == http.StatusNotFound {
			fmt.Printf("Repository %s/%s not found\n", args.owner, args.repo)
			os.Exit(1)
		}

		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	var versions []*version.Version

	for _, rel := range rels {
		v, err := version.NewVersion(*rel.TagName)
		if err != nil {
			fmt.Printf("error parsing tag %s\n", *rel.TagName)

			continue
		}

		if args.major != nil && (v.Segments()[0] != *args.major || v.Prerelease() != "") {

			continue
		}

		versions = append(versions, v)
	}

	sort.Sort(version.Collection(versions))

	if len(versions) == 0 {
		fmt.Println("No releases found")
		os.Exit(1)
	}

	fmt.Println(versions[len(versions)-1])
}

func parseArgs() args {
	as := os.Args[1:]

	if len(as) == 0 {
		help()
	}

	parts := strings.Split(as[0], "/")
	if len(parts) != 2 {
		help()
	}

	if len(as) == 1 {
		return args{
			owner: parts[0],
			repo:  parts[1],
			major: nil,
		}
	}

	m, err := strconv.Atoi(as[1])
	if err != nil {
		help()
	}

	return args{
		owner: parts[0],
		repo:  parts[1],
		major: &m,
	}
}

func help() {
	e := fmt.Sprintf("Examples:\t%s %s %d\n", os.Args[0], "helm/helm", 2)
	e += fmt.Sprintf("\t\t%s %s", os.Args[0], "starship/starship")
	fmt.Printf("Usage:\t\t%s %s [%s]\n%s\n", os.Args[0], "owner/repo", "major", e)
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

type args struct {
	owner string
	repo  string
	major *int
}
