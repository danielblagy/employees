# employees

A test application to familiarize yourself with docker.

### Create images and start api and postgres containers
```
docker compose up --build
```

### Stops containers and removes containers, networks, volumes, and images created by `up`
```
docker compose down
```

## Previous version where only one service was being implemented

### Create image
```
docker build -t employees-app .
```

### Run container from created image
```
docker run -p 5000:5000 -tid employees-app
```