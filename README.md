# isod

A simple HTTP server to create config-drive iso images. So I don't have to install mkisofs everywhere.

## Run the docker image

```
docker run -d -p 6500:8080 kelseyhightower/isod:latest
```

## Testing with curl

```
curl -o config-drive.iso -X POST http://107.178.217.37:6500/genisoimage --data-binary @cloud-config.yml
```
