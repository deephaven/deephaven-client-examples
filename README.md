# Deephaven Client Examples

This collection of client examples performs (roughly) the same operations
in several different client languages to compare their interfaces.

The examples do not currently run, because some irrelevant code is omitted.
However, they do compile.

## Python

Ensure a recent version of the pydeephaven package is installed.

```sh
$ cd ./python
$ python3 main.py
```

## C++

First, build and install the C++ client packages according to `deephaven-core/cpp-client/README.md`.
Most of the build steps only need to be done once.
However, `CMAKE_PREFIX_PATH` must always be set appropriately when building the example.

Then, edit `./cpp/CMakeLists.txt` where indicated to set the correct path for the client library.

```sh
$ cd ./cpp
$ mkdir build && cd build
$ cmake ..
$ make
$ ./example
```

## Go

First, ensure you have the Go client present somewhere (checking out the `go-client-api` branch will suffice).

Then, edit `./go/go.mod` where indicated to set the correct path for the client library.

```sh
$ cd ./go
$ go mod tidy
$ go build
$ ./go_client_example
```
