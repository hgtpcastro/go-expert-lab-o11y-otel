package test

import "testing"

type TestApp struct{}

type TestAppResult struct{}

func NewTestApp() *TestApp {
	return &TestApp{}
}

func (a *TestApp) Run(t *testing.T) (result *TestAppResult) {
	return &TestAppResult{}
}
