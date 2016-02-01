[![Build Status](https://travis-ci.org/danielpalstra/jsyk.svg?branch=master)](https://travis-ci.org/danielpalstra/jsyk)

# JSYK - Just So You Know
JSYK is a tiny, simple application that sends "informative change" events to elasticsearch by using the commandline. Inspired by the prezi changelog application https://github.com/prezi/changelog

# Elasticsearch (& Kibana)1
JSYK uses elasticsearch as the backend. The backend for jsyk can be configured by using the `--url` parameter, but it's recommended to set the `ES_URL` environment variable. When both are not provided http://localhost:9200 will be used by default. jsyk uses the `@timestamp` field so that Kibana is supported

# Usage
Currently jsyk only supports adding new events. Hit the following command for the available arguments
```
jsyk add --help
```

# Building
Clone this repository
```
go get
go install
```

Add your `$GOPATH/bin` to the `$PATH` env variable and start using jsyk from anywhere.


# Testing
Perform a minimal test by setting up a elasticsearch instance (we recommend this docker image) and performing the following commandline
```
go build && ./jsyk add -m "Started deployment awesomeness"
```

# TODO
Too much todo, should start with code cleanup
