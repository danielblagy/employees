# employees

A test application to familiarize yourself with docker.

## Create image
```
docker build -t employees-app .
```

## Run container from created image
```
docker run -p 5000:5000 -tid employees-app
```