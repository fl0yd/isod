# isod

A simple HTTP server to create config-drive iso images. So I don't have to install mkisofs everywhere.

## Run the docker image

```
docker run -d -p 80:8080 kelseyhightower/isod:latest
```

## Testing with curl

```
curl -o config-drive.iso -X POST http://isod.36containers.com/genisoimage --data-binary @cloud-config.yml
```
