# Introduction

More information about this application can be found in this [article](http://willsena.dev/using-graceful-shutdown-approach-to-dispose-of-applications/)

# Build image

```shell
docker buildx build -t graceful-shutdown -f graceful_shutdown/Containerfile .

podman build -t graceful-shutdown -f graceful_shutdown/Containerfile .

# old docker way
docker build -t graceful-shutdown -f graceful_shutdown/Containerfile .
```

# Running

```shell
docker run --name graceful-shutdown -d -it --rm graceful-shutdown
docker logs -f graceful-shutdown

# now sent signal to application stop
docker stop graceful-shutdown 

# Using 
# Podman

podman run --name graceful-shutdown -d -it --rm graceful-shutdown
podman logs -f graceful-shutdown

# now sent signal to application stop
podman stop graceful-shutdown 
```

