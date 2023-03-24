# ECS healthcheck

Image with binary to run healthchecks over HTTP. Meant to be used by images running on ECS

## Usage

Add the healthcheck binary to your own image with something like this:

```dockerfile
# This healthcheck image
FROM imageUrl:imageTag as healthcheck

# The image with your app in it
FROM docker-image-with-my-app

# Add health check binary to the final image
COPY --from=healthcheck /healthcheck /

ENTRYPOINT ["/app/example"]
```

Then use it as part of your `container_definition`:

```hcl
...
  healthCheck : {
    command : [
      "CMD",
      "/healthcheck"
    ]
  }
...
```

By default, it'll probe `/status`. To use a different path and/or port, pass it a `-path` argument. E.g.

```hcl
  healthCheck : {
    command : [
      "CMD",
      "/healthcheck",
      "-path=:8080/my-status-path"
    ]
  }
```