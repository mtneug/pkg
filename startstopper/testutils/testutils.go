// Copyright (c) 2016 Matthias Neugebauer
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package testutils contains helper code for testing startstopper.
package testutils

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockStartStopper is a mocked StartStopper.
type MockStartStopper struct {
	mock.Mock
}

// Start implements interface.
func (m *MockStartStopper) Start(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop implements interface.
func (m *MockStartStopper) Stop(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Done implements interface.
func (m *MockStartStopper) Done() <-chan struct{} {
	args := m.Called()
	return args.Get(0).(<-chan struct{})
}

// Err implements interface.
func (m *MockStartStopper) Err(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// MockRunner is a mocked Runner.
type MockRunner struct {
	mock.Mock
}

// Run implements interface.
func (m *MockRunner) Run(ctx context.Context, stopChan <-chan struct{}) error {
	args := m.Called(ctx)
	<-stopChan
	return args.Error(0)
}
