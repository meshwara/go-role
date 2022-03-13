package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/Permify/permify-gorm/collections"
	"github.com/Permify/permify-gorm/models"
	"github.com/Permify/permify-gorm/repositories/scopes"
)

// PermissionRepository is an autogenerated mock type for the PermissionRepository type
type PermissionRepository struct {
	mock.Mock
}

// Migrate provides a mock function.
func (_m *PermissionRepository) Migrate() (err error) {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPermissionByID provides a mock function with given fields: ID
func (_m *PermissionRepository) GetPermissionByID(ID uint) (permission models.Permission, err error) {
	ret := _m.Called(ID)

	var r0 models.Permission
	if rf, ok := ret.Get(0).(func(uint) models.Permission); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPermissionByGuardName provides a mock function with given fields: guardName
func (_m *PermissionRepository) GetPermissionByGuardName(guardName string) (permission models.Permission, err error) {
	ret := _m.Called(guardName)

	var r0 models.Permission
	if rf, ok := ret.Get(0).(func(string) models.Permission); ok {
		r0 = rf(guardName)
	} else {
		r0 = ret.Get(0).(models.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guardName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPermissions provides a mock function with given fields: IDs
func (_m *PermissionRepository) GetPermissions(IDs []uint) (permissions collections.Permission, err error) {
	ret := _m.Called(IDs)

	var r0 collections.Permission
	if rf, ok := ret.Get(0).(func([]uint) collections.Permission); ok {
		r0 = rf(IDs)
	} else {
		r0 = ret.Get(0).(collections.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]uint) error); ok {
		r1 = rf(IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPermissionsByGuardNames provides a mock function with given fields: guardNames
func (_m *PermissionRepository) GetPermissionsByGuardNames(guardNames []string) (permissions collections.Permission, err error) {
	ret := _m.Called(guardNames)

	var r0 collections.Permission
	if rf, ok := ret.Get(0).(func([]string) collections.Permission); ok {
		r0 = rf(guardNames)
	} else {
		r0 = ret.Get(0).(collections.Permission)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(guardNames)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPermissionIDs provides a mock function with given fields: pagination
func (_m *PermissionRepository) GetPermissionIDs(pagination scopes.GormPager) (permissionIDs []uint, totalCount int64, err error) {
	ret := _m.Called(pagination)

	var r0 []uint
	if rf, ok := ret.Get(0).(func(scopes.GormPager) []uint); ok {
		r0 = rf(pagination)
	} else {
		r0 = ret.Get(0).([]uint)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(scopes.GormPager) int64); ok {
		r1 = rf(pagination)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(scopes.GormPager) error); ok {
		r2 = rf(pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDirectPermissionIDsOfUserByID provides a mock function with given fields: userID,  pagination
func (_m *PermissionRepository) GetDirectPermissionIDsOfUserByID(userID uint, pagination scopes.GormPager) (permissionIDs []uint, totalCount int64, err error) {
	ret := _m.Called(userID, pagination)

	var r0 []uint
	if rf, ok := ret.Get(0).(func(uint, scopes.GormPager) []uint); ok {
		r0 = rf(userID, pagination)
	} else {
		r0 = ret.Get(0).([]uint)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(uint, scopes.GormPager) int64); ok {
		r1 = rf(userID, pagination)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uint, scopes.GormPager) error); ok {
		r2 = rf(userID, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPermissionIDsOfRolesByIDs provides a mock function with given fields: roleIDs,  pagination
func (_m *PermissionRepository) GetPermissionIDsOfRolesByIDs(roleIDs []uint, pagination scopes.GormPager) (permissionIDs []uint, totalCount int64, err error) {
	ret := _m.Called(roleIDs, pagination)

	var r0 []uint
	if rf, ok := ret.Get(0).(func([]uint, scopes.GormPager) []uint); ok {
		r0 = rf(roleIDs, pagination)
	} else {
		r0 = ret.Get(0).([]uint)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func([]uint, scopes.GormPager) int64); ok {
		r1 = rf(roleIDs, pagination)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]uint, scopes.GormPager) error); ok {
		r2 = rf(roleIDs, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FirstOrCreate provides a mock function with given fields: permission
func (_m *PermissionRepository) FirstOrCreate(permission *models.Permission) error {
	ret := _m.Called(permission)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Permission) error); ok {
		r0 = rf(permission)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Updates provides a mock function with given fields: permission, updates
func (_m *PermissionRepository) Updates(permission *models.Permission, updates map[string]interface{}) (err error) {
	ret := _m.Called(permission, updates)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Permission, map[string]interface{}) error); ok {
		r0 = rf(permission, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: permission
func (_m *PermissionRepository) Delete(permission *models.Permission) (err error) {
	ret := _m.Called(permission)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Permission) error); ok {
		r0 = rf(permission)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
