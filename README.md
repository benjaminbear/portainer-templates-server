# Tiny Portainer templates.json file server

![Travis build status](https://travis-ci.com/benjaminbear/portainer-templates-server.svg?branch=master)
![Docker build status](https://img.shields.io/docker/cloud/build/bbaerthlein/portainer-templates-server)
![Docker build automated](https://img.shields.io/docker/cloud/automated/bbaerthlein/portainer-templates-server)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/benjaminbear/portainer-templates-server)
![Go version](https://img.shields.io/github/go-mod/go-version/benjaminbear/portainer-templates-server?filename=go.mod)
![License](https://img.shields.io/github/license/benjaminbear/portainer-templates-server)

This docker service simply serves just one file, a templates.json file

## Usage

You can access the templates.json file via http.
The accessable file name is always "templates.json".

```
http://localhost:8080/templates.json
```

### Flags

```
-f string
      path to templates.json file (default "templates.json")
-p string
      port to serve on (default "8080")
```

### Docker image

Pull repository

```bash
docker pull bbaerthlein/portainer-templates-server
```


Run container:

```bash
docker run -p 8080:8080 bbaerthlein/portainer-templates-server
```

### Binary

Download and extract the binary for your os and architecture from the [release page](https://github.com/benjaminbear/portainer-templates-server/releases/).

Start the binary: (example for linux)

```bash
./portainer-templates-server_linux_x64 -f myfolder/templates.json
```

## Mount template (docker)

If you want to mount a templates.json file:

```
-v /your/local/path/file.json:/templates.json
```

## Docker hub

https://hub.docker.com/r/bbaerthlein/portainer-templates-server