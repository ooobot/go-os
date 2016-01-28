# Go Platform [![GoDoc](https://godoc.org/github.com/micro/go-platform?status.svg)](https://godoc.org/github.com/micro/go-platform) [![Travis CI](https://travis-ci.org/micro/go-platform.svg?branch=master)](https://travis-ci.org/micro/go-platform)

This is a microservice platform library built to be used with micro/go-micro. 
It provides all of the fundamental tooling required to run and manage 
a microservice environment. It's pluggable just like go-micro. It will be 
especially vital for anything above 20+ services and preferred by 
organisations. Developers looking to write standalone services should 
continue to use go-micro. 

Go platform depends on middleware/before-after funcs being added to 
go-micro. While each package can be used independently, it will be 
much more powerful as a whole.

The libraries here are not yet implemented. Discussion for 
the interfaces are welcome.

Note. Go-platform will include 1-3 supported implementations of each feature. 
Further community wide features should be contributed to [go-plugins](https://github.com/micro/go-plugins).

## Features
Each package provides a feature interface that will be pluggable and backed by a 
number of services.

Package     |   Features
-------     |   ---------
[auth](https://godoc.org/github.com/micro/go-platform/auth)        |   authentication and authorisation for users and services	
[config](https://godoc.org/github.com/micro/go-platform/config)      |   dynamic configuration which is namespaced and versioned
db          |   distributed database abstraction
[discovery](https://godoc.org/github.com/micro/go-platform/discovery)   |   extends the go-micro registry to add heartbeating, etc
[event](https://godoc.org/github.com/micro/go-platform/event)   |   platform event publication, subscription and aggregation 
[kv](https://godoc.org/github.com/micro/go-platform/kv)          |   simply key value layered on memcached, etcd, consul 
[log](https://godoc.org/github.com/micro/go-platform/log)         |   structured logging to stdout, logstash, fluentd, pubsub
[monitor](https://godoc.org/github.com/micro/go-platform/monitor)     |   add custom healthchecks measured with distributed systems in mind
[metrics](https://godoc.org/github.com/micro/go-platform/metrics)     |   instrumentation and collation of counters
router      |   global circuit breaking, load balancing, A/B testing
[sync](https://godoc.org/github.com/micro/go-platform/sync)       |   distributed locking, leadership election, etc
[trace](https://godoc.org/github.com/micro/go-platform/trace)       |   distributed tracing of request/response
