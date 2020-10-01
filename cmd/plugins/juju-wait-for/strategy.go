// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"reflect"
	"strings"
	"time"

	"github.com/juju/errors"

	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/cmd/plugins/juju-wait-for/query"
)

// StrategyFunc defines a way to change the underlying stategy function that
// can be changed depending on the callee.
type StrategyFunc func(string, []params.Delta, query.Query) bool

// Strategy defines a series of instructions to run for a given wait for
// plan.
type Strategy struct {
	Client  WatchAllAPI
	Timeout time.Duration
}

// Run the strategy and return the given result set.
func (s *Strategy) Run(name string, input string, fn StrategyFunc) error {
	q, err := query.Parse(input)
	if err != nil {
		return errors.Trace(err)
	}

	watcher, err := s.Client.WatchAll()
	if err != nil {
		return errors.Trace(err)
	}
	defer func() {
		_ = watcher.Stop()
	}()

	timeout := make(chan struct{})
	go func() {
		select {
		case <-time.After(s.Timeout):
			close(timeout)
			watcher.Stop()
		}
	}()

	for {
		deltas, err := watcher.Next()
		if err != nil {
			select {
			case <-timeout:
				return errors.Errorf("timed out waiting for %q to reach goal state", name)
			default:
				return errors.Trace(err)
			}
		}

		if done := fn(name, deltas, q); done {
			return nil
		}
	}
}

// GenericScope allows the query to introspect an entity.
type GenericScope struct {
	Info params.EntityInfo
}

// GetIdentValue returns the value of the identifier in a given scope.
func (m GenericScope) GetIdentValue(name string) (interface{}, error) {
	refType := reflect.TypeOf(m.Info).Elem()
	for i := 0; i < refType.NumField(); i++ {
		field := refType.Field(i)
		v := strings.Split(field.Tag.Get("json"), ",")[0]
		if v == name {
			refValue := reflect.ValueOf(m.Info).Elem()
			data := refValue.Field(i).Interface()
			return data, nil
		}
	}
	return nil, errors.Errorf("Runtime Error: identifier %q not found on EntityInfo", name)
}
