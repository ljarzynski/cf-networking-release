// This file was generated by counterfeiter
package fakes

import (
	"io"
	"lib/serial"
	"sync"
)

type Serializer struct {
	DecodeAllStub        func(file io.ReadSeeker, outData interface{}) error
	decodeAllMutex       sync.RWMutex
	decodeAllArgsForCall []struct {
		file    io.ReadSeeker
		outData interface{}
	}
	decodeAllReturns struct {
		result1 error
	}
	EncodeAndOverwriteStub        func(file serial.OverwriteableFile, outData interface{}) error
	encodeAndOverwriteMutex       sync.RWMutex
	encodeAndOverwriteArgsForCall []struct {
		file    serial.OverwriteableFile
		outData interface{}
	}
	encodeAndOverwriteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Serializer) DecodeAll(file io.ReadSeeker, outData interface{}) error {
	fake.decodeAllMutex.Lock()
	fake.decodeAllArgsForCall = append(fake.decodeAllArgsForCall, struct {
		file    io.ReadSeeker
		outData interface{}
	}{file, outData})
	fake.recordInvocation("DecodeAll", []interface{}{file, outData})
	fake.decodeAllMutex.Unlock()
	if fake.DecodeAllStub != nil {
		return fake.DecodeAllStub(file, outData)
	} else {
		return fake.decodeAllReturns.result1
	}
}

func (fake *Serializer) DecodeAllCallCount() int {
	fake.decodeAllMutex.RLock()
	defer fake.decodeAllMutex.RUnlock()
	return len(fake.decodeAllArgsForCall)
}

func (fake *Serializer) DecodeAllArgsForCall(i int) (io.ReadSeeker, interface{}) {
	fake.decodeAllMutex.RLock()
	defer fake.decodeAllMutex.RUnlock()
	return fake.decodeAllArgsForCall[i].file, fake.decodeAllArgsForCall[i].outData
}

func (fake *Serializer) DecodeAllReturns(result1 error) {
	fake.DecodeAllStub = nil
	fake.decodeAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *Serializer) EncodeAndOverwrite(file serial.OverwriteableFile, outData interface{}) error {
	fake.encodeAndOverwriteMutex.Lock()
	fake.encodeAndOverwriteArgsForCall = append(fake.encodeAndOverwriteArgsForCall, struct {
		file    serial.OverwriteableFile
		outData interface{}
	}{file, outData})
	fake.recordInvocation("EncodeAndOverwrite", []interface{}{file, outData})
	fake.encodeAndOverwriteMutex.Unlock()
	if fake.EncodeAndOverwriteStub != nil {
		return fake.EncodeAndOverwriteStub(file, outData)
	} else {
		return fake.encodeAndOverwriteReturns.result1
	}
}

func (fake *Serializer) EncodeAndOverwriteCallCount() int {
	fake.encodeAndOverwriteMutex.RLock()
	defer fake.encodeAndOverwriteMutex.RUnlock()
	return len(fake.encodeAndOverwriteArgsForCall)
}

func (fake *Serializer) EncodeAndOverwriteArgsForCall(i int) (serial.OverwriteableFile, interface{}) {
	fake.encodeAndOverwriteMutex.RLock()
	defer fake.encodeAndOverwriteMutex.RUnlock()
	return fake.encodeAndOverwriteArgsForCall[i].file, fake.encodeAndOverwriteArgsForCall[i].outData
}

func (fake *Serializer) EncodeAndOverwriteReturns(result1 error) {
	fake.EncodeAndOverwriteStub = nil
	fake.encodeAndOverwriteReturns = struct {
		result1 error
	}{result1}
}

func (fake *Serializer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.decodeAllMutex.RLock()
	defer fake.decodeAllMutex.RUnlock()
	fake.encodeAndOverwriteMutex.RLock()
	defer fake.encodeAndOverwriteMutex.RUnlock()
	return fake.invocations
}

func (fake *Serializer) recordInvocation(key string, args []interface{}) {
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

var _ serial.Serializer = new(Serializer)
