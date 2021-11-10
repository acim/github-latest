# github-latest

![check](https://github.com/acim/github-latest/workflows/check/badge.svg)
![release](https://github.com/acim/github-latest/workflows/release/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/acim/github-latest.svg)](https://pkg.go.dev/github.com/acim/github-latest)
[![Go Report](https://goreportcard.com/badge/github.com/acim/github-latest)](https://goreportcard.com/report/github.com/acim/github-latest)

Small utility to find out the latest release of a GitHub repository.

## Installation

### From source

```bash
go install github.com/acim/github-latest@latest
```

### [Binary releases](https://github.com/acim/github-latest/releases)

## Usage

`github-latest owner/repo [major]`

major parameter is optional

## Example

`github-latest helm/helm 2`

## Private repositories or rate limit

If you want to access your private repository or you need more than 60 requests per hour,
define GITHUB_ACCESS_TOKEN environment variable on your local machine.

You can generate your token on **GitHub - Settings - Developer settings - Personal access tokens**
