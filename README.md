# wonderful-readme-stats – A self-hosted solution for repository statistics

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

**English** | [Русский][repo_readme_ru_url] | [简体中文][repo_readme_cn_url] |
[Español][repo_readme_es_url]

A wonderful out-of-the-box and self-hosted solution for displaying statistics (stargazers, contributors, etc.) about your repository in your README file.

Features:

- 100% **free** and **open source** under the [Apache 2.0][repo_license_url]
  license;
- For **any** level of developer's knowledge and technical expertise;
- **Well-documented**, with a lot of tips and assists from the authors;
- ...

## ⚡️ Quick start

Feel free to using the latest version of the `wonderful-readme-stats`
from our [official Docker image][docker_image_url].

You need to create the `docker-compose.yml` file (see the
[Complete user guide][repo_cug_url] below).

Now, just run it by the Docker Compose command in the isolated container
on your local machine (*for testing*) or remote server (*for production*):

```console
docker-compose up
```

### 📦 Other way to quick start

Download ready-made `exe` files for Windows, `deb`, `rpm`, `apk` or Arch
Linux packages from the [Releases][repo_releases_url] page.

## 📖 Complete user guide

To get a complete guide to use and understand the basic principles of the
`wonderful-readme-stats` project, we have prepared a comprehensive explanation
of each step at once in this README file.

> 💬 From the authors: We always treasure your time and want you to start
> building really great web products on this awesome technology stack as
> soon as possible!

We hope you find answers to all of your questions! 👌 But, if you do not find
needed information, feel free to create an [issue][repo_issues_url] or send a
[PR][repo_pull_request_url] to this repository.

Don't forget to switch this page for your language (current is
**English**): [Русский][repo_readme_ru_url], [简体中文][repo_readme_cn_url],
[Español][repo_readme_es_url].

### Step 1: Prepare the Docker Compose file for the backend

Create the `docker-compose.yml` file with the following:

```yaml
version: '3.8'

# Define services.
services:
  # Service for the backend.
  wonderful_readme_stats:
    # Configuration for the Docker image for the service.
    image: koddr/wonderful-readme-stats:latest
    # Set restart rules for the container.
    restart: unless-stopped
    # Forward the exposed port 8080 on the container to port 8080 on the host machine.
    ports:
      - '8080:8080'
    # Set required environment variables for the backend.
    environment:
      - GITHUB_TOKEN=${GITHUB_TOKEN}
      - REPOSITORY_OWNER=${REPOSITORY_OWNER}
      - REPOSITORY_NAME=${REPOSITORY_NAME}
      - SERVER_PORT=8080
      - SERVER_READ_TIMEOUT=5
      - SERVER_WRITE_TIMEOUT=10
      - AVATAR_SHAPE=rounded
      - AVATAR_SIZE=64
      - AVATAR_HORIZONTAL_MARGIN=12
      - AVATAR_VERTICAL_MARGIN=12
      - AVATAR_ROUNDED_RADIUS=16.0
      - OUTPUT_IMAGE_MAX_PER_ROW=16
      - OUTPUT_IMAGE_MAX_ROWS=2
      - OUTPUT_IMAGE_UPDATE_INTERVAL=3600
```

In this file, we create a container with the settings defined from the **environment variables**.

> 💡 Note: This is a deliberate step, as you are supposed to deploy the project on your remote server via Docker Compose. Therefore the backend configuration is in this way.

#### Environment variables explanation

The list of the environment variables are used to configure the `wonderful-readme-stats`.

| Environment variable name      | Description                                                                         | Type     | Default value            |
| ------------------------------ | ----------------------------------------------------------------------------------- | -------- | ------------------------ |
| `GITHUB_TOKEN`                 | Token for the GitHub API from your [GitHub account][github_token_url] settings      | `string` | `""`                     |
| `REPOSITORY_OWNER`             | Repository owner on GitHub                                                          | `string` | `koddr`                  |
| `REPOSITORY_NAME`              | Repository name on GitHub                                                           | `string` | `wonderful-readme-stats` |
| `SERVER_PORT`                  | Port for the server                                                                 | `int`    | `8080`                   |
| `SERVER_READ_TIMEOUT`          | HTTP read timeout for the server (in seconds)                                       | `int`    | `5`                      |
| `SERVER_WRITE_TIMEOUT`         | HTTP write timeout for the server (in seconds)                                      | `int`    | `10`                     |
| `AVATAR_SHAPE`                 | Shape type for the one user avatar (available values: `rounded`, `circular`)        | `string` | `rounded`                |
| `AVATAR_SIZE`                  | Size for the one user avatar (in pixels)                                            | `int`    | `64`                     |
| `AVATAR_HORIZONTAL_MARGIN`     | Horizontal margin for the one user avatar (in pixels)                               | `int`    | `12`                     |
| `AVATAR_VERTICAL_MARGIN`       | Vertical margin for the one user avatar (in pixels)                                 | `int`    | `12`                     |
| `AVATAR_ROUNDED_RADIUS`        | Radius of corners for the one user avatar (in pixels, required for `rounded` shape) | `float`  | `16.0`                   |
| `OUTPUT_IMAGE_MAX_PER_ROW`     | Max number of avatars per row for the output image                                  | `int`    | `16`                     |
| `OUTPUT_IMAGE_MAX_ROWS`        | Max number of rows with avatars for the output image                                | `int`    | `2`                      |
| `OUTPUT_IMAGE_UPDATE_INTERVAL` | Update interval for the output images (in seconds)                                  | `int`    | `3600`                   |

### Step 2: Configure remote server with Portainer

...

## 🎯 Motivation to create

...

> 💬 From the authors: Earlier, we have already saved the world once, it was
> [Create Go App][cgapp_url] (yep, that's our project too). The
> [GitHub stars][cgapp_stars_url] statistics of this project can't lie:
> more than **2.2k** developers of any level and different countries start a
> new project through this CLI tool.

## 🏆 A win-win cooperation

If you liked the `wonderful-readme-stats` project and found it useful for your tasks,
please click a 👁️ **Watch** button to avoid missing notifications about new
versions, and give it a 🌟 **GitHub Star**!

It really **motivates** us to make this product **even** better.

...

And now, I invite you to participate in this project! Let's work **together** to
create and popularize the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.
- Say a few words about the project on your social networks and blogs
  (Dev.to, Medium, Хабр, and so on).

Your PRs, issues & any words are welcome! Thank you 😘

### 🌟 Stargazers

...

## ⚠️ License

[`wonderful-readme-stats`][repo_url] is free and open-source software licensed
under the [Apache 2.0 License][repo_license_url], created and supported by
[Vic Shóstak][author_url] with 🩵 for people and robots. Official logo
distributed under the [Creative Commons License][repo_cc_license_url] (CC BY-SA
4.0 International).

<!-- Go links -->

[go_report_url]: https://goreportcard.com/report/github.com/koddr/wonderful-readme-stats
[go_dev_url]: https://pkg.go.dev/github.com/koddr/wonderful-readme-stats
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_url]: https://codecov.io/gh/koddr/wonderful-readme-stats
[go_code_coverage_img]: https://img.shields.io/codecov/c/gh/koddr/wonderful-readme-stats.svg?logo=codecov&style=for-the-badge
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/koddr/wonderful-readme-stats
[repo_issues_url]: https://github.com/koddr/wonderful-readme-stats/issues
[repo_pull_request_url]: https://github.com/koddr/wonderful-readme-stats/pulls
[repo_releases_url]: https://github.com/koddr/wonderful-readme-stats/releases
[repo_license_url]: https://github.com/koddr/wonderful-readme-stats/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_readme_ru_url]: https://github.com/koddr/wonderful-readme-stats/blob/main/README_RU.md
[repo_readme_cn_url]: https://github.com/koddr/wonderful-readme-stats/blob/main/README_CN.md
[repo_readme_es_url]: https://github.com/koddr/wonderful-readme-stats/blob/main/README_ES.md
[repo_cug_url]: https://github.com/koddr/wonderful-readme-stats#-complete-user-guide

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->
[github_token_url]: https://github.com/settings/tokens
[cgapp_url]: https://github.com/create-go-app/cli
[cgapp_stars_url]: https://github.com/create-go-app/cli/stargazers
[docker_image_url]: https://hub.docker.com/repository/docker/koddr/wonderful-readme-stats
[portainer_url]: https://docs.portainer.io
