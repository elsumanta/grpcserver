// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handler

import (
	"context"
	"github.com/elsumanta/grpcserver/server/model"
	"sync"
)

// Ensure, that RepoMock does implement Repo.
// If this is not the case, regenerate this file with moq.
var _ Repo = &RepoMock{}

// RepoMock is a mock implementation of Repo.
//
// 	func TestSomethingThatUsesRepo(t *testing.T) {
//
// 		// make and configure a mocked Repo
// 		mockedRepo := &RepoMock{
// 			RegisterFunc: func(ctx context.Context, req model.Register) (int, error) {
// 				panic("mock out the Register method")
// 			},
// 		}
//
// 		// use mockedRepo in code that requires Repo
// 		// and then make assertions.
//
// 	}
type RepoMock struct {
	// RegisterFunc mocks the Register method.
	RegisterFunc func(ctx context.Context, req model.Register) (int, error)

	// calls tracks calls to the methods.
	calls struct {
		// Register holds details about calls to the Register method.
		Register []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req model.Register
		}
	}
	lockRegister sync.RWMutex
}

// Register calls RegisterFunc.
func (mock *RepoMock) Register(ctx context.Context, req model.Register) (int, error) {
	if mock.RegisterFunc == nil {
		panic("RepoMock.RegisterFunc: method is nil but Repo.Register was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req model.Register
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockRegister.Lock()
	mock.calls.Register = append(mock.calls.Register, callInfo)
	mock.lockRegister.Unlock()
	return mock.RegisterFunc(ctx, req)
}

// RegisterCalls gets all the calls that were made to Register.
// Check the length with:
//     len(mockedRepo.RegisterCalls())
func (mock *RepoMock) RegisterCalls() []struct {
	Ctx context.Context
	Req model.Register
} {
	var calls []struct {
		Ctx context.Context
		Req model.Register
	}
	mock.lockRegister.RLock()
	calls = mock.calls.Register
	mock.lockRegister.RUnlock()
	return calls
}
