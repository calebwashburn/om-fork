// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type JobsConfigurer struct {
	JobsStub        func(productGUID string) (api.JobsOutput, error)
	jobsMutex       sync.RWMutex
	jobsArgsForCall []struct {
		productGUID string
	}
	jobsReturns struct {
		result1 api.JobsOutput
		result2 error
	}
	ConfigureStub        func(api.JobConfigurationInput) error
	configureMutex       sync.RWMutex
	configureArgsForCall []struct {
		arg1 api.JobConfigurationInput
	}
	configureReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *JobsConfigurer) Jobs(productGUID string) (api.JobsOutput, error) {
	fake.jobsMutex.Lock()
	fake.jobsArgsForCall = append(fake.jobsArgsForCall, struct {
		productGUID string
	}{productGUID})
	fake.recordInvocation("Jobs", []interface{}{productGUID})
	fake.jobsMutex.Unlock()
	if fake.JobsStub != nil {
		return fake.JobsStub(productGUID)
	} else {
		return fake.jobsReturns.result1, fake.jobsReturns.result2
	}
}

func (fake *JobsConfigurer) JobsCallCount() int {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return len(fake.jobsArgsForCall)
}

func (fake *JobsConfigurer) JobsArgsForCall(i int) string {
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	return fake.jobsArgsForCall[i].productGUID
}

func (fake *JobsConfigurer) JobsReturns(result1 api.JobsOutput, result2 error) {
	fake.JobsStub = nil
	fake.jobsReturns = struct {
		result1 api.JobsOutput
		result2 error
	}{result1, result2}
}

func (fake *JobsConfigurer) Configure(arg1 api.JobConfigurationInput) error {
	fake.configureMutex.Lock()
	fake.configureArgsForCall = append(fake.configureArgsForCall, struct {
		arg1 api.JobConfigurationInput
	}{arg1})
	fake.recordInvocation("Configure", []interface{}{arg1})
	fake.configureMutex.Unlock()
	if fake.ConfigureStub != nil {
		return fake.ConfigureStub(arg1)
	} else {
		return fake.configureReturns.result1
	}
}

func (fake *JobsConfigurer) ConfigureCallCount() int {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return len(fake.configureArgsForCall)
}

func (fake *JobsConfigurer) ConfigureArgsForCall(i int) api.JobConfigurationInput {
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return fake.configureArgsForCall[i].arg1
}

func (fake *JobsConfigurer) ConfigureReturns(result1 error) {
	fake.ConfigureStub = nil
	fake.configureReturns = struct {
		result1 error
	}{result1}
}

func (fake *JobsConfigurer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.jobsMutex.RLock()
	defer fake.jobsMutex.RUnlock()
	fake.configureMutex.RLock()
	defer fake.configureMutex.RUnlock()
	return fake.invocations
}

func (fake *JobsConfigurer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
