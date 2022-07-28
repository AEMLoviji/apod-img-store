# apod-img-store

Project to grab APOD images and store it locally in PostgreSQL and provide image to user when it is requested.

> Project requirements: <https://github.com/justty/golang-units/blob/main/unit-1.md>

## Requirements

Before you spin up API and it's infrastructure make sure you have

- `docker` installed
- `Make` installed

## How to run

> Tips: run `make help` or just `make` to see all available make targets

### Only once

_If you are running the system first time please call below target_

```sh
make build
```

### To spin-up API and Postgresql  

```sh
make up
```

## Request samples

cURL samples can be run from command line

### /image-of-the-day

> Make sure you change the date parameter before sending request 

```sh
curl -X GET 'http://localhost:8080/image-of-the-day?date=YYYY-mm-dd'
```

### /images

```sh
curl -X GET 'http://localhost:8080/images'
```

## TODO

[] Low coupling achived by using interfaces. It will help us in feature to cover code with Tests.
