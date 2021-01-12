// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import modelsval "github.com/PedroMFC/EvaluaUGR/internal/microval/modelsval"

// IValSaver is an autogenerated mock type for the IValSaver type
type IValSaver struct {
	mock.Mock
}

// CrearAsignatura provides a mock function with given fields: asignatura
func (_m *IValSaver) CrearAsignatura(asignatura string) {
	_m.Called(asignatura)
}

// GuardarValoracion provides a mock function with given fields: asignatura, val
func (_m *IValSaver) GuardarValoracion(asignatura string, val *modelsval.Valoracion) {
	_m.Called(asignatura, val)
}

// ObtenerAsignaturas provides a mock function with given fields:
func (_m *IValSaver) ObtenerAsignaturas() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// ObtenerValoraciones provides a mock function with given fields: asignatura
func (_m *IValSaver) ObtenerValoraciones(asignatura string) []modelsval.Valoracion {
	ret := _m.Called(asignatura)

	var r0 []modelsval.Valoracion
	if rf, ok := ret.Get(0).(func(string) []modelsval.Valoracion); ok {
		r0 = rf(asignatura)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]modelsval.Valoracion)
		}
	}

	return r0
}
