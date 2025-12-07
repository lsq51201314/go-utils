package glog

import "testing"

func TestGlog(t *testing.T) {
	Info("Hello World")
	Error("Hello World")
	g, _ := New()
	g.Info("hello")
	g.Error("world")
}
