package main

import "go.mattglei.ch/tlog"

var Cache = tlog.Group("cache", &struct {
	MarshalResponse tlog.Op
	Marshal         struct {
		JSON tlog.Op
		CSV  tlog.Op
	}
}{})

func main() {
	Cache.MarshalResponse.Info("hello world!")
}
