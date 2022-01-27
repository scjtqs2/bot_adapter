// Package handler 业务代码
package handler

import "go.uber.org/dig"

type AdapterHandler struct {
	ct *dig.Container
}

// NewAdapterHandler 新建handler
func NewAdapterHandler(ct *dig.Container) (*AdapterHandler, error) {
	var adapt AdapterHandler
	adapt.ct = ct
	return &adapt, nil
}
