# github-latest

[![GitHub](https://img.shields.io/github/license/acim/github-latest)](LICENSE)
[![Build Status](https://drone.ablab.de/api/badges/acim/github-latest/status.svg)](https://drone.ablab.de/acim/github-latest)

Small utility to find out the latest release of a GitHub repository.

## Install

``` bash
go get -u github.com/acim/github-latest
```

## Usage

github-latest owner/repo major

## Example

github-latest helm/helm 2

## Private repositories or rate limit

If you want to access your private repository or you need more than 60 requests per hour,
define GITHUB_ACCESS_TOKEN environment variable.

You can generate your token on GitHub - Settings - Developer settings - Personal access tokens
