## Explore Go

Gym for Go lang study

## Dev Notes

Build docker dev env:

```
$ docker build -t expgo .
```

Run it interactly:

```
$ docker run --rm -itv $PWD:/app expgo
```

Run pagkage test inside it:

```
$ cd stringcalculator
$ go test
```

________

From clean docker image if you prefer (no Dockerfile, no image build):

```
$ docker run --rm -itv $PWD:/app golang:1.13-alpine sh
$ apk add gcc libc-dev
$ cd /app
```

Run pagkage test inside it:

```
$ cd stringcalculator
$ go test
```
