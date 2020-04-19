# github-latest

[![GitHub](https://img.shields.io/github/license/acim/github-latest)](LICENSE)
[![HitCount](http://hits.dwyl.com/acim/github-latest.svg)](http://hits.dwyl.com/acim/github-latest)

Small utility to find out the latest release of a GitHub repository.

## Installation

### From source

``` bash
go get -u github.com/acim/github-latest
```

### Binary release (Linux and macOS)

Download binary from our [releases](https://github.com/acim/github-latest/releases).

## Usage

github-latest owner/repo [major]

major parameter is optional

## Example

github-latest helm/helm 2

## Private repositories or rate limit

If you want to access your private repository or you need more than 60 requests per hour,
define GITHUB_ACCESS_TOKEN environment variable.

You can generate your token on GitHub - Settings - Developer settings - Personal access tokens
