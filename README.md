apicube
========

[![GoDoc](https://godoc.org/github.com/ikeikeikeike/apicube?status.svg)](https://godoc.org/github.com/ikeikeikeike/apicube)
[![Build Status](https://api.travis-ci.org/ikeikeikeike/apicube.svg?branch=master)](https://travis-ci.org/ikeikeikeike/apicube)
[![Go Report Card](https://goreportcard.com/badge/github.com/ikeikeikeike/apicube)](https://goreportcard.com/report/github.com/ikeikeikeike/apicube)


Install Packages([dep](https://github.com/golang/dep))

```php
$ make install deps
```

Protocol generation.

```php
$ make protocol
$ make deps  # I think need this cmd when you are going to install with beginning after "make protocol"
```

## Dependencies Library


**protoc@3.6+**

 ```php
$ brew install protobuf
```

## Model migration(Optional)

```php
$ make modelgen
```

Elasticsearch migration

```php
$ make reindex
```

## Development

Build packages and Run server

```php
$ make runner  # Auto build when file is edited

$ export GRPC_TRACE=all
$ export GRPC_VERBOSITY=DEBUG
$ export GRPC_GO_LOG_VERBOSITY_LEVEL=2
$ export GRPC_GO_LOG_SEVERITY_LEVEL=info
$ export APICUBE_ESURL=http://127.0.0.1:9200
$ (cd ./rpc && make run)
```

## Road to the Goish

This is completely Goish code if resolved all of errors.

```php
$ make check
$ make check-completely
```

## Environment Variables

Frequency used below

| Key | Default Value | Description |
|:-----------|:------------|:------------|
| APICUBE_URI | grpc://0.0.0.0:50001 | APICube server: this var is set grpc server's host and port, which means the same as FQDN |
| APICUBE_GWURI | http://0.0.0.0:50000 | APICube server: this var is set grpc gateway server's host and port, which means the same as FQDN |
| APICUBE_RDBURI | redis://127.0.0.1:6379/10 | Redis: this var is set server host and port with db number, that's like DSN |
| APICUBE_DLMURI | redis://127.0.0.1:6379/9 | Redis: DLMURI is set distributed lock server host and port with db number, that's like DSN |
| APICUBE_ESURL | http://127.0.0.1:9200 | Elasticsearch: this means is the same as APICUBE_URI |
| APICUBE_FURI | file://./storage/data.flac | Storage folder: this var is set s3 or file uri |
| APICUBE_DSN | root:@tcp(127.0.0.1:3306)/apicube?parseTime=true | MySQL: Data source name |
| APICUBE_MODE | development | Defines project environment mode: development,staging,production |
| GOOGLE_APPLICATION_CREDENTIALS | None | Google API Credentials Json e.g. apicube-fca6752be359.json |

see [base/util/env.go](base/util/env.go)

