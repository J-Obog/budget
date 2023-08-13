// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	data "github.com/J-Obog/paidoff/data"
	mock "github.com/stretchr/testify/mock"
)

// CategoryStore is an autogenerated mock type for the CategoryStore type
type CategoryStore struct {
	mock.Mock
}

type CategoryStore_Expecter struct {
	mock *mock.Mock
}

func (_m *CategoryStore) EXPECT() *CategoryStore_Expecter {
	return &CategoryStore_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: id, accountId
func (_m *CategoryStore) Delete(id string, accountId string) (bool, error) {
	ret := _m.Called(id, accountId)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(id, accountId)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(id, accountId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryStore_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CategoryStore_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id string
//   - accountId string
func (_e *CategoryStore_Expecter) Delete(id interface{}, accountId interface{}) *CategoryStore_Delete_Call {
	return &CategoryStore_Delete_Call{Call: _e.mock.On("Delete", id, accountId)}
}

func (_c *CategoryStore_Delete_Call) Run(run func(id string, accountId string)) *CategoryStore_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *CategoryStore_Delete_Call) Return(_a0 bool, _a1 error) *CategoryStore_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryStore_Delete_Call) RunAndReturn(run func(string, string) (bool, error)) *CategoryStore_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAll provides a mock function with given fields:
func (_m *CategoryStore) DeleteAll() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CategoryStore_DeleteAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAll'
type CategoryStore_DeleteAll_Call struct {
	*mock.Call
}

// DeleteAll is a helper method to define mock.On call
func (_e *CategoryStore_Expecter) DeleteAll() *CategoryStore_DeleteAll_Call {
	return &CategoryStore_DeleteAll_Call{Call: _e.mock.On("DeleteAll")}
}

func (_c *CategoryStore_DeleteAll_Call) Run(run func()) *CategoryStore_DeleteAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CategoryStore_DeleteAll_Call) Return(_a0 error) *CategoryStore_DeleteAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CategoryStore_DeleteAll_Call) RunAndReturn(run func() error) *CategoryStore_DeleteAll_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id, accountId
func (_m *CategoryStore) Get(id string, accountId string) (*data.Category, error) {
	ret := _m.Called(id, accountId)

	var r0 *data.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*data.Category, error)); ok {
		return rf(id, accountId)
	}
	if rf, ok := ret.Get(0).(func(string, string) *data.Category); ok {
		r0 = rf(id, accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryStore_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CategoryStore_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id string
//   - accountId string
func (_e *CategoryStore_Expecter) Get(id interface{}, accountId interface{}) *CategoryStore_Get_Call {
	return &CategoryStore_Get_Call{Call: _e.mock.On("Get", id, accountId)}
}

func (_c *CategoryStore_Get_Call) Run(run func(id string, accountId string)) *CategoryStore_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *CategoryStore_Get_Call) Return(_a0 *data.Category, _a1 error) *CategoryStore_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryStore_Get_Call) RunAndReturn(run func(string, string) (*data.Category, error)) *CategoryStore_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: accountId
func (_m *CategoryStore) GetAll(accountId string) ([]data.Category, error) {
	ret := _m.Called(accountId)

	var r0 []data.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]data.Category, error)); ok {
		return rf(accountId)
	}
	if rf, ok := ret.Get(0).(func(string) []data.Category); ok {
		r0 = rf(accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]data.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryStore_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type CategoryStore_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - accountId string
func (_e *CategoryStore_Expecter) GetAll(accountId interface{}) *CategoryStore_GetAll_Call {
	return &CategoryStore_GetAll_Call{Call: _e.mock.On("GetAll", accountId)}
}

func (_c *CategoryStore_GetAll_Call) Run(run func(accountId string)) *CategoryStore_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *CategoryStore_GetAll_Call) Return(_a0 []data.Category, _a1 error) *CategoryStore_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryStore_GetAll_Call) RunAndReturn(run func(string) ([]data.Category, error)) *CategoryStore_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: accountId, name
func (_m *CategoryStore) GetByName(accountId string, name string) (*data.Category, error) {
	ret := _m.Called(accountId, name)

	var r0 *data.Category
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*data.Category, error)); ok {
		return rf(accountId, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *data.Category); ok {
		r0 = rf(accountId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data.Category)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(accountId, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryStore_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type CategoryStore_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - accountId string
//   - name string
func (_e *CategoryStore_Expecter) GetByName(accountId interface{}, name interface{}) *CategoryStore_GetByName_Call {
	return &CategoryStore_GetByName_Call{Call: _e.mock.On("GetByName", accountId, name)}
}

func (_c *CategoryStore_GetByName_Call) Run(run func(accountId string, name string)) *CategoryStore_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *CategoryStore_GetByName_Call) Return(_a0 *data.Category, _a1 error) *CategoryStore_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryStore_GetByName_Call) RunAndReturn(run func(string, string) (*data.Category, error)) *CategoryStore_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// Insert provides a mock function with given fields: category
func (_m *CategoryStore) Insert(category data.Category) error {
	ret := _m.Called(category)

	var r0 error
	if rf, ok := ret.Get(0).(func(data.Category) error); ok {
		r0 = rf(category)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CategoryStore_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type CategoryStore_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - category data.Category
func (_e *CategoryStore_Expecter) Insert(category interface{}) *CategoryStore_Insert_Call {
	return &CategoryStore_Insert_Call{Call: _e.mock.On("Insert", category)}
}

func (_c *CategoryStore_Insert_Call) Run(run func(category data.Category)) *CategoryStore_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(data.Category))
	})
	return _c
}

func (_c *CategoryStore_Insert_Call) Return(_a0 error) *CategoryStore_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CategoryStore_Insert_Call) RunAndReturn(run func(data.Category) error) *CategoryStore_Insert_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: updated
func (_m *CategoryStore) Update(updated data.Category) (bool, error) {
	ret := _m.Called(updated)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(data.Category) (bool, error)); ok {
		return rf(updated)
	}
	if rf, ok := ret.Get(0).(func(data.Category) bool); ok {
		r0 = rf(updated)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(data.Category) error); ok {
		r1 = rf(updated)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CategoryStore_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type CategoryStore_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - updated data.Category
func (_e *CategoryStore_Expecter) Update(updated interface{}) *CategoryStore_Update_Call {
	return &CategoryStore_Update_Call{Call: _e.mock.On("Update", updated)}
}

func (_c *CategoryStore_Update_Call) Run(run func(updated data.Category)) *CategoryStore_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(data.Category))
	})
	return _c
}

func (_c *CategoryStore_Update_Call) Return(_a0 bool, _a1 error) *CategoryStore_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CategoryStore_Update_Call) RunAndReturn(run func(data.Category) (bool, error)) *CategoryStore_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewCategoryStore creates a new instance of CategoryStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCategoryStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *CategoryStore {
	mock := &CategoryStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
