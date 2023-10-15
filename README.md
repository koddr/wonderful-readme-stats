# wonderful-readme-stats – A self-hosted solution for repository statistics

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

**English** | [Русский][repo_readme_ru_url] | [简体中文][repo_readme_cn_url] | [Español][repo_readme_es_url]

A wonderful **out-of-the-box** and **self-hosted** solution for displaying statistics (stargazers, contributors, etc.) about your repository in your **README** file.

Features:

- 100% **free** and **open source** under the [Apache 2.0][repo_license_url] license;
- For **any** level of developer's knowledge and technical expertise;
- **Well-documented**, with a lot of tips and assists from the authors;
- ...

## ⚡️ Quick start

Feel free to using the latest version of the `wonderful-readme-stats` from our [official Docker image][docker_image_url].

> 💡 Note: See the [Complete user guide][repo_cug_url] to understand the basic principles of the project.

Run the `wonderful-readme-stats` container with your environment variables:

```console
docker run -d \
  -p 8080:8080 \
  -e REPOSITORY_OWNER=<OWNER> \
  -e REPOSITORY_NAME=<NAME> \
  koddr/wonderful-readme-stats:latest
```

After starting, the `wonderful-readme-stats` backend will be available at `http://localhost:8080` on your local machine. To test the backend, open your browser and navigate to:

- `http://localhost:8080/github/<OWNER>/<NAME>/stargazers.png` to see the stargazers statistics of the repository in the auto-generated PNG image.
- `http://localhost:8080/github/<OWNER>/<NAME>/contributors.png` to see the contributors statistics of the repository in the auto-generated PNG image.

That's it! 🔥 Your wonderful statistics are ready to be deployed to a remote server and added to your repository's README.

### 📦 Other way to quick start

Download ready-made `exe` files for Windows, `deb`, `rpm`, `apk` or Arch Linux packages from the [Releases][repo_releases_url] page.

## 📖 Complete user guide

To get a complete guide to use and understand the basic principles of the `wonderful-readme-stats` project, we have prepared a comprehensive explanation of each step at once in this README file.

> 💬 From the authors: We always treasure your time and want you to start building really great web products on this awesome technology stack as soon as possible!

We hope you find answers to all of your questions! 👌 But, if you do not find needed information, feel free to create an [issue][repo_issues_url] or send a [PR][repo_pull_request_url] to this repository.

Don't forget to switch this page for your language (current is **English**): [Русский][repo_readme_ru_url], [简体中文][repo_readme_cn_url], [Español][repo_readme_es_url].

### Step 1: Configure remote server with Portainer

We recommend using the [Portainer][portainer_url] Community Edition platform to make the process of deploying the `wonderful-readme-stats` backend more comfortable and faster. Almost every cloud provider has a ready-to-use Docker image that can be deployed directly from the dashboard.

Let's take a look at [Timeweb.Cloud][timeweb_cloud_url] as an example:

- Login (or register) to your cloud provider account.
- Click to the **Server** link on the left panel.
- Click to the **Create** button.
- Switch to the **Marketplace** tab and type `portainer` in the **Search** field:

<img width="640" alt="timeweb cloud" src="https://github.com/koddr/wonderful-readme-stats/assets/11155743/645a4141-f869-4046-9836-3a1dbea083f1">

- Fill the required fields (region, CPU, RAM, disk, and so on).
- Click to the **Order for ...** button and wait for the process to complete.

Now, you're ready to continue configuring the `wonderful-readme-stats` backend.

#### Manual configuration

If you don't want to use the pre-built image provided by your cloud provider, here are instructions on how to manually install Portainer on your server.

> ❗️ Warning: All steps must be performed strongly **after** installing Docker to your server. See [documentation][docker_install_url] page for more information.

- Create a new Docker volume for Portainer data:

```console
docker volume create portainer_data
```

- Start the Portainer container:

```console
docker run -d \
  -p 8000:8000 \
  -p 9443:9443 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -v portainer_data:/data \
  --name portainer --restart=always \
  portainer/portainer-ce:latest
```

- Check the status of the container:

```console
docker ps
```

### Step 2: Prepare the `wonderful-readme-stats` backend

Excellent, let's now set up the `wonderful-readme-stats` backend that will collect and display statistics on the selected repository.

- Go to your **Portainer** dashboard.
- Click to the **Add stack** button.
- Place the following content to the **Web editor** field:

```yaml
version: '3.8'

# Define services.
services:
  # Service for the backend.
  wonderful_readme_stats:
    # Configuration for the Docker image for the service.
    image: 'koddr/wonderful-readme-stats:latest'
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

- Go to **Environment variables** options and click to the **Advanced mode** button.
- Paste your secret data as a plaintext (key-value pairs, separated by a newline) to the field:

```bash
GITHUB_TOKEN=your-secret-github-token-123456789
REPOSITORY_OWNER=your-repo-owner
REPOSITORY_NAME=your-repo-name
```

- Then, click to the **Deploy the stack** button on the bottom of the page.
- After starting the container, the backend will be available at `http://YOUR-SERVER-IP:8080`.
- To test the `wonderful-readme-stats` backend, open your browser and navigate to:
  - `http://YOUR-SERVER-IP:8080/github/<OWNER>/<NAME>/stargazers.png` to see the stargazers statistics of the repository in the auto-generated PNG image.
  - `http://YOUR-SERVER-IP:8080/github/<OWNER>/<NAME>/contributors.png` to see the contributors statistics of the repository in the auto-generated PNG image.

#### Environment variables explanation

Yes, we create a container with the settings defined from the **environment variables**. This was done on purpose to make it easier to deploy to a remote server so that you don't have to create and store a configuration file.

> ❗️ Warning: Do not leave the token for `GITHUB_TOKEN` exposed as a string, only as a variable! **This is not safe**. If you want to commit this to your repository, make sure you don't leave any secret data in the file first.

The full list of the environment variables are used to configure the `wonderful-readme-stats` backend.

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

> 💡 Note: You can choose not to define `GITHUB_TOKEN`, but then the update time interval of the final image in the `OUTPUT_IMAGE_UPDATE_INTERVAL` parameter **cannot be lower** than the recommended `3600` seconds.
>
> This is because without defining a GitHub token, the `wonderful-readme-stats` backend will work with **public limits** for getting data from the API.

### Step 3: Configure Nginx Proxy Manager

To avoid thinking about configuring [Nginx][nginx_url] proxy and [Let's Encrypt][lets_encrypt_url] SSL certificates, let's install [Nginx Proxy Manager][nginx_proxy_manager_url] on the remote server using Portainer. He's going to do it all for us.

- Go to your **Portainer** dashboard.
- Click to the **Add stack** button.
- Place the following content to the **Web editor** field:

```yaml
version: '3.8'

# Define the services.
services:
  # Service for Nginx Proxy Manager.
  nginx_proxy_manager:
    # Configuration for the Docker image for the service.
    image: 'jc21/nginx-proxy-manager:latest'
    # Set restart rules for the container.
    restart: unless-stopped
    # Forward the exposed ports on the container to ports on the host machine.
    ports:
      - '80:80'
      - '81:81' # port for the Nginx Proxy Manager
      - '443:443'
    # Set volumes for the container.
    volumes:
      - ./data:/data
      - ./letsencrypt:/etc/letsencrypt
    # Set networks for the container.
    networks:
      - nginx_proxy_manager_default

# Define networks.
networks:
  # Network for Nginx Proxy Manager.
  nginx_proxy_manager_default:
    external: true # require external access
```

- Then, click to the **Deploy the stack** button on the bottom of the page.
- After starting the container, **Nginx Proxy Manager** will be available at `http://YOUR-SERVER-IP:81`.

> ❗️ Warning: A default email for the first login to the Nginx Proxy Manager is `admin@example.com`, and password is `changeme`. **Do not forget to change this credentials**! See official [documentation][nginx_proxy_manager_url] page for more details.

### Step 4: Configure domain and SSL certificate

After configuring Nginx Proxy Manager, let's configure the domain name and create the SSL certificate.

- Go to your **Nginx Proxy Manager** dashboard.
- Click to the **Add Proxy Host** button and fill the required fields:
  - `Domain Names` with your domain name.
  - `Scheme` with the HTTP scheme.
  - `Forward Hostname / IP` with the IP address of your remote server.
  - `Forward Port` with the port of the `wonderful-readme-stats` backend (by default, `8080`).
  - Check the `Cache assets` and `Block Common Exploits` checkboxes.
- Next, go to the **SSL** section:
  - In `SSL Certificate` field select the `Request a new SSL certificate` option.
  - Check the `Force SSL`, `HTTP/2 Support`, `HSTS Enabled` and `HSTS Subdomains` checkboxes.
  - `Email Address for Let's Encrypt` with your real email address.
  - Check the `I Agree to the Let's Encrypt Terms of Service` checkbox.
- Then, click to the **Save** button and wait for the process to complete.

### Step 5: Add the statistics to your README

Now, you can add the statistics of your repository to the README.

For repository stargazers:

```bash
![Repository stargazers](https://your-domain.com/github/<OWNER>/<NAME>/stargazers.png)
```

For repository contributors:

```bash
![Repository contributors](https://your-domain.com/github/<OWNER>/<NAME>/contributors.png)
```

And the final image will be, for example:

![Repository stargazers](https://stats.gowebly.org/github/gowebly/gowebly/stargazers.png)

## 🎯 Motivation to create

...

> 💬 From the authors: Earlier, we have already saved the world twice, it was [Create Go App][cgapp_url] and [gowebly][gowebly_url] (yep, that's our projects too). The GitHub stars statistics of this project can't lie: more than **2.2k** developers of any level and different countries start a new project through these CLI tools.

## 🏆 A win-win cooperation

If you liked the `wonderful-readme-stats` project and found it useful for your tasks, please click a 👁️ **Watch** button to avoid missing notifications about new versions, and give it a 🌟 **GitHub Star**!

It really **motivates** us to make this product **even** better.

...

And now, I invite you to participate in this project! Let's work **together** to create and popularize the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.
- Say a few words about the project on your social networks and blogs (Dev.to, Medium, Хабр, and so on).

Your PRs, issues & any words are welcome! Thank you 😘

### 🌟 Stargazers

...

## ⚠️ License

[`wonderful-readme-stats`][repo_url] is free and open-source software licensed under the [Apache 2.0 License][repo_license_url], created and supported by [Vic Shóstak][author_url] with 🩵 for people and robots. Official logo distributed under the [Creative Commons License][repo_cc_license_url] (CC BY-SA 4.0 International).

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

[timeweb_cloud_url]: https://timeweb.cloud/r/koddr
[github_token_url]: https://github.com/settings/tokens
[gowebly_url]: https://github.com/gowebly/gowebly
[cgapp_url]: https://github.com/create-go-app/cli
[docker_image_url]: https://hub.docker.com/repository/docker/koddr/wonderful-readme-stats
[docker_compose_url]: https://docs.docker.com/compose
[docker_install_url]: https://docs.docker.com/engine/install/#server
[portainer_url]: https://docs.portainer.io
[nginx_url]: https://nginx.org
[nginx_proxy_manager_url]: https://nginxproxymanager.com/guide/
[lets_encrypt_url]: https://letsencrypt.org