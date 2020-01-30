## Go String Calculator

String Calculator kata as gym for Go lang study.

http://codingdojo.org/kata/StringCalculator/

## Run dev isolated docker environment

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

## Dev Notes

- http://codingdojo.org/kata/StringCalculator/
- https://golang.org/pkg/strconv/
- https://golang.org/pkg/strings/
- https://gobyexample.com/
- https://learnxinyminutes.com/docs/go/
- http://mattyjwilliams.blogspot.com/2013/06/why-im-happy-to-live-without-generics.html
