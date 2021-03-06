# Dev Guide

## Building From Source
In order to build Aptomi from source, you will need Go (the latest 1.10.x) and a couple of external dependencies:
* __glide__ - all Go dependencies for Aptomi are managed via [Glide](https://glide.sh/)
* __docker__ - to run Aptomi in a container, as well as to run a sample LDAP server with user data
* __kubernetes-cli and kubernetes-helm__ - for using Kubernetes with Helm
* __npm__ - to build the UI and automatically generate the table of contents in README.md 
* __telnet, jq__ - for the script which runs smoke tests

If you are on macOS, install [Homebrew](https://brew.sh/) and [Docker For Mac](https://docs.docker.com/docker-for-mac/install/), then run: 
```
brew install go glide docker kubernetes-cli kubernetes-helm npm telnet jq
```

Make sure `$GOPATH` is exported in your shell (e.g. `~/.bash_profile`) and `$PATH` includes `$GOPATH/bin`:
```
# echo $GOPATH
/Users/test/go

# echo $PATH
/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/Users/test/go/bin
``` 

Check out the Aptomi source code from the repo:
```
mkdir -p $GOPATH/src/github.com/Aptomi
cd $GOPATH/src/github.com/Aptomi
git clone https://github.com/Aptomi/aptomi.git
```

In order to build Aptomi, you must first tell Glide to fetch all project dependencies. It will read the list of
dependencies defined in `glide.lock` and fetch them into a local "vendor" folder:
```
make vendor 
```

After that, you can build the project binaries. This will produce `aptomi` and `aptomictl` in the current directory:
```
make
```

You can also get these binaries built and installed into `$GOPATH/bin`:
```
make install
```

## Tests & Code Validation

Command    | Action          | LDAP Required
-----------|-----------------|--------------
```make test```    | Unit tests | No
```make alltest``` | Integration + Unit tests | Yes
```make smoke```   | Smoke tests + Integration + Unit tests | Yes
```make profile-engine```   | Profile engine for CPU usage | No
```make coverage```   | Calculate code coverage by unit tests | No
```make coverage-full```   | Calculate code coverage by unit & integration tests | Yes

Command     | Action          | Description
------------|-----------------|--------------
```make fmt```  | Re-format code | Re-formats all code according to Go standards
```make lint``` | Examine code | Run linters to examine Go source code and reports suspicious constructs

## Web UI
Source code is available in [webui](webui)

Make sure you have latest `node` and `npm`. We have tested with node v8.9.1 and npm 5.5.1 and it's
known to work with these.

Command     | Action
------------|----------
```npm install```  | Install dependencies
```npm run dev``` | Serve with hot reload at localhost:8080
```npm run build``` | Build for production with minification
```npm run build --report``` | Build for production and view the bundle analyzer report
```npm run unit``` | Run unit tests: *coming soon*
```npm run e2e``` | Run e2e tests: *coming soon*
```npm run test``` | Run all tests: *coming soon*

## Creating configs
If you are building from source and want to run our examples, you will need to generate configs for the Aptomi server and client, and add some sample users to Aptomi.

You can do this manually, or you can run an [installation script](install_compact.md) that will install the latest Aptomi release (it will co-exist well with the binaries built from source) and create those configs for you.

## How to release
Use `git tag` and `make release` for creating new release.

1. Create an annotated git tag and push it to the github repo. Use a commit message like `Aptomi v0.1.2`.

```
git tag -a v0.1.2
git push origin v0.1.2
```

1. Create a GitHub API token with the `repo` scope selected to upload artifacts to the GitHub release page. You can create
one [here](https://github.com/settings/tokens/new). This token should be added to the environment variables as `GITHUB_TOKEN`.

1. Run `make release`. This will create everything you need and upload all artifacts to github.

1. Go to https://github.com/Aptomi/aptomi/releases/tag/v0.1.2 and fix the changelog / description if needed.
