// This file was generated by counterfeiter
package fake_uploader

import (
	"net/url"
	"sync"

	"github.com/cloudfoundry-incubator/executor/depot/uploader"
)

type FakeUploader struct {
	UploadStub        func(fileLocation string, destinationUrl *url.URL) (int64, error)
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		fileLocation   string
		destinationUrl *url.URL
	}
	uploadReturns struct {
		result1 int64
		result2 error
	}
}

func (fake *FakeUploader) Upload(fileLocation string, destinationUrl *url.URL) (int64, error) {
	fake.uploadMutex.Lock()
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		fileLocation   string
		destinationUrl *url.URL
	}{fileLocation, destinationUrl})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(fileLocation, destinationUrl)
	} else {
		return fake.uploadReturns.result1, fake.uploadReturns.result2
	}
}

func (fake *FakeUploader) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *FakeUploader) UploadArgsForCall(i int) (string, *url.URL) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.uploadArgsForCall[i].fileLocation, fake.uploadArgsForCall[i].destinationUrl
}

func (fake *FakeUploader) UploadReturns(result1 int64, result2 error) {
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

var _ uploader.Uploader = new(FakeUploader)