package kube

import (
	"errors"
)

var (
	ErrNoPodFound     = errors.New("no pods found in namepspace")
	ErrNoIngressFound = errors.New("no ingress found in namepspace")
)
