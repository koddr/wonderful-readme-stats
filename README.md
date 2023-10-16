# wonderful-readme-stats – A self-hosted solution for repository statistics

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

**English** | [Русский][repo_readme_ru_url] | [简体中文][repo_readme_cn_url] | [Español][repo_readme_es_url]

A wonderful **out-of-the-box** and **self-hosted** solution for displaying statistics (stargazers, contributors, etc.) about your repository in your **README** file.

Features:

- 100% **free** and **Open Source** under the [Apache 2.0][repo_license_url] license;
- For **any** level of developer's knowledge and technical expertise;
- Write on the pure **Go** language, without any overheads;
- Used minimum dependencies, **well-tested** and **well-optimized** for the production;
- Delivered as a **self-hosted** solution in the Docker image, independent of anyone else;
- Has the **Complete user guide** to understand the basic principles and deployment processes;
- **Well-documented**, with a lot of tips and assists from the author.

## ⚡️ Quick start

Feel free to using the latest version of the `wonderful-readme-stats` from our [official Docker image][docker_image_url].

> 💡 Note: See the [Complete user guide][repo_cug_url] to understand the basic principles of the project.

Run the `wonderful-readme-stats` container with your environment variables:

```console
docker run -d \
  -e REPOSITORY_OWNER=<OWNER> \
  -e REPOSITORY_NAME=<NAME> \
  koddr/wonderful-readme-stats:latest
```

After starting, the `wonderful-readme-stats` backend will be available at `http://localhost:9876` on your local machine. To test the backend, open your browser and navigate to:

- `/github/<OWNER>/<NAME>/stargazers.png` to see the stargazers stats of the repo (PNG image).
- `/github/<OWNER>/<NAME>/contributors.png` to see the contributors stats of the repo (PNG image).

That's it! 🔥 A wonderful stats are ready to be deployed to a remote server and added to your repo's README.

### 🛠 Manual way to quick start

If you want to build the image yourself (or change something in the code), just `git clone` this repository.

I made sure that the documentation in the code is **comprehensive** and **covers** as much functionality as possible.

### 📦 Other way to quick start

Download ready-made `exe` files for Windows, `deb`, `rpm`, `apk` or Arch Linux packages from the [Releases][repo_releases_url] page.

## 📖 Complete user guide

To get a complete guide to use and understand the basic principles of the `wonderful-readme-stats` project, I have prepared a comprehensive explanation of each step at once in this README file.

> 💬 From the author: I always treasure your time and want you to start building really great web products on this awesome technology stack as soon as possible!

I hope you find answers to all of your questions! 👌 But, if you do not find needed information, feel free to create an [issue][repo_issues_url] or send a [PR][repo_pull_request_url] to this repository.

> 🔠 Don't forget to switch this page to your language (current is **English**): [Русский][repo_readme_ru_url], [简体中文][repo_readme_cn_url], [Español][repo_readme_es_url].

### Step 1: Configure remote server with Portainer

I recommend using the [Portainer][portainer_url] Community Edition platform to make the process of deploying the `wonderful-readme-stats` backend more comfortable and faster. Almost every cloud provider has a ready-to-use Docker image that can be deployed directly from the dashboard.

Let's take a look at [Timeweb.Cloud][timeweb_cloud_url] as an example:

- Login (or register) to the cloud provider.
- Click to the **Server** link on the left panel.
- Click to the **Create** button on the top right.
- Switch to the **Marketplace** tab and type word `portainer` in the **Search** field:

<img width="600" alt="timeweb cloud" src="https://github.com/koddr/wonderful-readme-stats/assets/11155743/568710e6-8460-426b-ac85-aa361f519791">

- Click to the **Portainer card** and select your preferred GNU/Linux distribution.
- Fill the required fields (region, CPU, RAM, disk, backup and so on).
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
    # Forward the exposed port 9876 on the container to port 9876 on the host machine.
    ports:
      - '9876:9876'
    # Set required environment variables for the backend.
    environment:
      - GITHUB_TOKEN=${GITHUB_TOKEN}
      - REPOSITORY_OWNER=${REPOSITORY_OWNER}
      - REPOSITORY_NAME=${REPOSITORY_NAME}
      - SERVER_PORT=9876
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
- After starting the container, the backend will be available at `http://YOUR-SERVER-IP:9876`.
- To test the `wonderful-readme-stats` backend, open your browser and navigate to:
  - `http://YOUR-SERVER-IP:9876/github/<OWNER>/<NAME>/stargazers.png` to see the stargazers statistics of the repository in the auto-generated PNG image.
  - `http://YOUR-SERVER-IP:9876/github/<OWNER>/<NAME>/contributors.png` to see the contributors statistics of the repository in the auto-generated PNG image.

#### Environment variables explanation

Yes, I create a container with the settings defined from the **environment variables**. This was done on purpose to make it easier to deploy to a remote server so that you don't have to create and store a configuration file.

The full list of the environment variables are used to configure the `wonderful-readme-stats` backend.

Environment variables for the **GitHub API**:

| Environment variable name | Description                                                                    | Type     | Default value |
| ------------------------- | ------------------------------------------------------------------------------ | -------- | ------------- |
| `GITHUB_TOKEN`            | Token for the GitHub API from your [GitHub account][github_token_url] settings | `string` | `""`          |

> ❗️ Warning: Do not leave the token for `GITHUB_TOKEN` exposed as a string, only as a variable! **This is not safe**. If you want to commit this to your repository, make sure you don't leave any secret data in the file first.

> 💡 Note: You can choose not to define `GITHUB_TOKEN`, but then the update time interval of the output image in the `OUTPUT_IMAGE_UPDATE_INTERVAL` parameter **cannot be lower** than the recommended `3600` seconds.
>
> This is because without defining a GitHub token, the `wonderful-readme-stats` backend will work with **public limits** for getting data from the API.

Environment variables for the **repository** name and owner:

| Environment variable name | Description                | Type     | Default value            |
| ------------------------- | -------------------------- | -------- | ------------------------ |
| `REPOSITORY_OWNER`        | Repository owner on GitHub | `string` | `koddr`                  |
| `REPOSITORY_NAME`         | Repository name on GitHub  | `string` | `wonderful-readme-stats` |

Environment variables for the **server** options:

| Environment variable name | Description                                    | Type  | Default value |
| ------------------------- | ---------------------------------------------- | ----- | ------------- |
| `SERVER_PORT`             | Port for the server                            | `int` | `9876`        |
| `SERVER_READ_TIMEOUT`     | HTTP read timeout for the server (in seconds)  | `int` | `5`           |
| `SERVER_WRITE_TIMEOUT`    | HTTP write timeout for the server (in seconds) | `int` | `10`          |

Environment variables for the **user avatar** options (used for the each avatar image):

| Environment variable name  | Description                                                                         | Type     | Default value |
| -------------------------- | ----------------------------------------------------------------------------------- | -------- | ------------- |
| `AVATAR_SHAPE`             | Shape type for the one user avatar (available values: `rounded`, `circular`)        | `string` | `rounded`     |
| `AVATAR_SIZE`              | Size for the one user avatar (in pixels)                                            | `int`    | `64`          |
| `AVATAR_HORIZONTAL_MARGIN` | Horizontal margin for the one user avatar (in pixels)                               | `int`    | `12`          |
| `AVATAR_VERTICAL_MARGIN`   | Vertical margin for the one user avatar (in pixels)                                 | `int`    | `12`          |
| `AVATAR_ROUNDED_RADIUS`    | Radius of corners for the one user avatar (in pixels, required for `rounded` shape) | `float`  | `16.0`        |

Environment variables for the **output image** options:

| Environment variable name      | Description                                          | Type  | Default value |
| ------------------------------ | ---------------------------------------------------- | ----- | ------------- |
| `OUTPUT_IMAGE_MAX_PER_ROW`     | Max number of avatars per row for the output image   | `int` | `16`          |
| `OUTPUT_IMAGE_MAX_ROWS`        | Max number of rows with avatars for the output image | `int` | `2`           |
| `OUTPUT_IMAGE_UPDATE_INTERVAL` | Update interval for the output images (in seconds)   | `int` | `3600`        |

### Step 3: Configure Nginx Proxy Manager

To avoid thinking about configuring [Nginx][nginx_url] proxy and [Let's Encrypt][lets_encrypt_url] SSL certificates, let's install [Nginx Proxy Manager][nginx_proxy_manager_url] on the remote server using Portainer. He's going to do it all for us.

- Go to your **Portainer** dashboard.
- Add a new Docker network called `nginx_proxy_manager_default` in the **Networks** section.
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

> 💡 Note: It is assumed that you already have a domain name purchased and its NS servers configured on your cloud provider, on which we have deployed Portainer and the rest of the tools.

Let's configure the domain name and create the SSL certificate.

- Go to your **Nginx Proxy Manager** dashboard.
- Click to the **Add Proxy Host** button and fill the required fields:
  - **Domain Names** with the purchased domain names (e.g., `example.com` and `www.example.com`).
  - **Scheme** with the HTTP scheme (by default, `http`).
  - **Forward Hostname / IP** with the IP address of your remote server.
  - **Forward Port** with the port of the `wonderful-readme-stats` backend (by default, `9876`).
  - Check the **Cache assets** and **Block Common Exploits** checkboxes.
- Next, go to the **SSL** section:
  - In the **SSL Certificate** field select the **Request a new SSL certificate** option.
  - Check the **Force SSL**, **HTTP/2 Support**, **HSTS Enabled** and **HSTS Subdomains** checkboxes.
  - **Email Address for Let's Encrypt** with your real email address.
  - Check the **I Agree to the Let's Encrypt Terms of Service** checkbox.
- Then, click to the **Save** button and wait for the process to complete.

### Step 5: Add the statistics to your README

Now, you can add the statistics of your repository to the README.

- For the repository **Stargazers** (*users that have starred the repository*):

```bash
![Repository stargazers](https://your-domain.com/github/<OWNER>/<NAME>/stargazers.png)
```

- For the repository **Contributors** (*users that have contributed to the repository*):

```bash
![Repository contributors](https://your-domain.com/github/<OWNER>/<NAME>/contributors.png)
```

- And the final image will be like this:

![Repository stargazers](https://stats.gowebly.org/github/gowebly/gowebly/stargazers.png)

> 💡 Note: In this example I use stargazers statistics of the [`gowebly`][gowebly_url] repository with a default settings.

## 🎯 Motivation to create

I've always loved making beautiful and informative READMEs for my projects.

Yes, there are so many tools out there already, but not all of them can be installed on the remote server (as a self-hosted solutions) and/or have advanced settings to display the final result.

That's why I created a Docker image with `wonderful-readme-stats` that I've been using since a long time. And now, I put it to the Open Source for the whole developer community.

Treasure your time and create only clear and handsome README pages with me! ✨

> 💬 From the author: Earlier, I have already saved the world twice: it was [Create Go App][cgapp_url] and [gowebly][gowebly_url] (yep, that's my projects too). The GitHub stars statistics of these projects can't lie: more than **2.2k** developers of any level and different countries start a new project through these CLI tools.

## 🏆 A win-win cooperation

If you liked the `wonderful-readme-stats` project and found it useful for your tasks, please click a 👁️ **Watch** button to avoid missing notifications about new versions, and give it a 🌟 **GitHub Star**!

It really **motivates** me to make this product **even** better.

![win-win cooperation](https://github.com/koddr/wonderful-readme-stats/assets/11155743/87f0bb0b-3cf2-44a7-a7b2-4bad50117c8e)

And now, I invite you to participate in this project! Let's work **together** to create and popularize the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current.
- Say a few words about the project on your social networks and blogs (Dev.to, Medium, Хабр, and so on).

Your PRs, issues & any words are welcome! Thank you 😘

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
