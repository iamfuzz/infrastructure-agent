// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0
package mocks

import (
	"github.com/newrelic/infrastructure-agent/internal/agent"
	"github.com/newrelic/infrastructure-agent/pkg/config"
	"github.com/newrelic/infrastructure-agent/pkg/entity"
	"github.com/newrelic/infrastructure-agent/pkg/plugins/ids"
	"github.com/newrelic/infrastructure-agent/pkg/sample"
	"github.com/newrelic/infrastructure-agent/pkg/sysinfo/hostname"
	"github.com/stretchr/testify/mock"
)

// AgentContext is an autogenerated mock type for the AgentContext type
type AgentContext struct {
	mock.Mock
}

func (_m *AgentContext) HostnameResolver() hostname.Resolver {
	ret := _m.Called()

	var r0 hostname.Resolver
	if rf, ok := ret.Get(0).(func() hostname.Resolver); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(hostname.Resolver)
	}

	return r0
}

func (_m *AgentContext) IDLookup() agent.IDLookup {
	ret := _m.Called()

	var r0 agent.IDLookup
	if rf, ok := ret.Get(0).(func() agent.IDLookup); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(agent.IDLookup)
	}

	return r0
}

func (_m *AgentContext) AddReconnecting(p agent.Plugin) {
	_m.Called(p)
}

func (_m *AgentContext) Reconnect() {
	_m.Called()
}

func (_m *AgentContext) ActiveEntitiesChannel() chan string {
	ret := _m.Called()

	var r0 chan string
	if rf, ok := ret.Get(0).(func() chan string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan string)
		}
	}

	return r0
}

// AgentIdentifier provides a mock function with given fields:
func (_m *AgentContext) AgentIdentifier() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CacheServicePids provides a mock function with given fields: source, pidMap
func (_m *AgentContext) CacheServicePids(source string, pidMap map[int]string) {
	_m.Called(source, pidMap)
}

// Config provides a mock function with given fields:
func (_m *AgentContext) Config() *config.Config {
	ret := _m.Called()

	var r0 *config.Config
	if rf, ok := ret.Get(0).(func() *config.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*config.Config)
		}
	}

	return r0
}

// GetServiceForPid provides a mock function with given fields: pid
func (_m *AgentContext) GetServiceForPid(pid int) (string, bool) {
	ret := _m.Called(pid)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(pid)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(int) bool); ok {
		r1 = rf(pid)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// SendData provides a mock function with given fields: _a0
func (_m *AgentContext) SendData(_a0 agent.PluginOutput) {
	_m.Called(_a0)
}

// SendEvent provides a mock function with given fields: event, entity
func (_m *AgentContext) SendEvent(e sample.Event, entityKey entity.Key) {
	_m.Called(e, entityKey)
}

// Unregister provides a mock function with given fields: _a0
func (_m *AgentContext) Unregister(_a0 ids.PluginID) {
	_m.Called(_a0)
}

// Version provides a mock function with given fields:
func (_m *AgentContext) Version() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

var _ agent.AgentContext = (*AgentContext)(nil)