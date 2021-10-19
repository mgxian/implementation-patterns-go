// Code generated by MockGen. DO NOT EDIT.
// Source: domain/service/article.go

// Package mock is a generated GoMock package.
package mock

import (
	model "article/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleServiceInterface is a mock of ArticleServiceInterface interface.
type MockArticleServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceInterfaceMockRecorder
}

// MockArticleServiceInterfaceMockRecorder is the mock recorder for MockArticleServiceInterface.
type MockArticleServiceInterfaceMockRecorder struct {
	mock *MockArticleServiceInterface
}

// NewMockArticleServiceInterface creates a new mock instance.
func NewMockArticleServiceInterface(ctrl *gomock.Controller) *MockArticleServiceInterface {
	mock := &MockArticleServiceInterface{ctrl: ctrl}
	mock.recorder = &MockArticleServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleServiceInterface) EXPECT() *MockArticleServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockArticleServiceInterface) CreateArticle(title, description, body string, authorId int) (model.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", title, description, body, authorId)
	ret0, _ := ret[0].(model.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockArticleServiceInterfaceMockRecorder) CreateArticle(title, description, body, authorId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockArticleServiceInterface)(nil).CreateArticle), title, description, body, authorId)
}
