// Code generated by protoc-gen-go-register. DO NOT EDIT.
// Sources: todo.proto

package pb

import (
	transport "github.com/erda-project/erda-infra/pkg/transport"
	reflect "reflect"
)

// RegisterTODOServiceImp todo.proto
func RegisterTODOServiceImp(regester transport.Register, srv TODOServiceServer, opts ...transport.ServiceOption) {
	_ops := transport.DefaultServiceOptions()
	for _, op := range opts {
		op(_ops)
	}
	RegisterTODOServiceHandler(regester, TODOServiceHandler(srv), _ops.HTTP...)
	RegisterTODOServiceServer(regester, srv, _ops.GRPC...)
}

// ServiceNames return all service names
func ServiceNames(svr ...string) []string {
	return append(svr,
		"erda.monitor.TODOService",
	)
}

var (
	todoserviceClientType  = reflect.TypeOf((*TODOServiceClient)(nil)).Elem()
	todoserviceServerType  = reflect.TypeOf((*TODOServiceServer)(nil)).Elem()
	todoserviceHandlerType = reflect.TypeOf((*TODOServiceHandler)(nil)).Elem()
)

// TODOServiceClientType .
func TODOServiceClientType() reflect.Type { return todoserviceClientType }

// TODOServiceServerType .
func TODOServiceServerType() reflect.Type { return todoserviceServerType }

// TODOServiceHandlerType .
func TODOServiceHandlerType() reflect.Type { return todoserviceHandlerType }

func Types() []reflect.Type {
	return []reflect.Type{
		// client types
		todoserviceClientType,
		// server types
		todoserviceServerType,
		// handler types
		todoserviceHandlerType,
	}
}
