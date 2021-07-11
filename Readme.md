# Squares =)

![](https://img.shields.io/badge/go-v1.14-blue) ![](https://img.shields.io/badge/image-golang:alpine-green)

Let n ∈ N Then n can be written as a sum of squares of integers. We define f(n) as the least number of squares that sum to n.
For example  since  f(7) = 4 since 7 = 2*2 + 1*2 + 1*2 + 1*2 and 7 is not a sum of 3 or less squares, f(4) = 1 since 4=2*2 or f(2) = 2 since 2 = 1*2 + 1*2.

---

### Get Started

##### Docker

If you have docker installed, clone this repository anywhere in your machine and in terminal, on this project folder, type the following comands:

```bash
 docker build -t goscraper .
 docker run -it goscraper
```

##### Golang


If you do not have docker installed but have golang in your machine, clone this repo inside your \$GOPATH work directory, them run the application in the terminal on this project folder:

```bash
go run .
```

---

### Extras

install docker: https://docs.docker.com/get-docker/ <br/>
install golang: https://golang.org/doc/install/
