package services

import (
	demo "goim-pro/api/protos/example"
	protos "goim-pro/api/protos/salty"
	"goim-pro/internal/app/controller/auth"
	"goim-pro/internal/app/controller/group"
	contactsrv "goim-pro/internal/app/services/contact"
	waitersrv "goim-pro/internal/app/services/demowaiter"
	usersrv "goim-pro/internal/app/services/user"
)

type Service struct {
	WaiterServer  demo.WaiterServer
	SMSServer     protos.SMSServiceServer
	UserServer    protos.UserServiceServer
	ContactServer protos.ContactServiceServer
	GroupServer   protos.GroupServiceServer
}

func NewService() *Service {
	return &Service{
		WaiterServer: waitersrv.New(),
		//SMSServer:     authsrv.New(),
		UserServer:    usersrv.New(),
		ContactServer: contactsrv.New(),
		//GroupServer:   groupsrv.New(),
		SMSServer:   auth.New(),
		GroupServer: group.New(),
	}
}
