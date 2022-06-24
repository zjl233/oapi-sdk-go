// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"

	"github.com/larksuite/oapi-sdk-go/service/contact/v3"
)

func (dispatcher *EventDispatcher) OnCustomAttrEventUpdatedV3(handler func(ctx context.Context, event *contact.CustomAttrEventUpdatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.custom_attr_event.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.custom_attr_event.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.custom_attr_event.updated_v3"] = contact.NewCustomAttrEventUpdatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnDepartmentCreatedV3(handler func(ctx context.Context, event *contact.DepartmentCreatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.created_v3"] = contact.NewDepartmentCreatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnDepartmentDeletedV3(handler func(ctx context.Context, event *contact.DepartmentDeletedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.deleted_v3"] = contact.NewDepartmentDeletedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnDepartmentUpdatedV3(handler func(ctx context.Context, event *contact.DepartmentUpdatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.department.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.department.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.department.updated_v3"] = contact.NewDepartmentUpdatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnEmployeeTypeEnumActivedV3(handler func(ctx context.Context, event *contact.EmployeeTypeEnumActivedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.actived_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.actived_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.actived_v3"] = contact.NewEmployeeTypeEnumActivedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnEmployeeTypeEnumCreatedV3(handler func(ctx context.Context, event *contact.EmployeeTypeEnumCreatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.created_v3"] = contact.NewEmployeeTypeEnumCreatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnEmployeeTypeEnumDeactivatedV3(handler func(ctx context.Context, event *contact.EmployeeTypeEnumDeactivatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.deactivated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.deactivated_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.deactivated_v3"] = contact.NewEmployeeTypeEnumDeactivatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnEmployeeTypeEnumDeletedV3(handler func(ctx context.Context, event *contact.EmployeeTypeEnumDeletedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.deleted_v3"] = contact.NewEmployeeTypeEnumDeletedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnEmployeeTypeEnumUpdatedV3(handler func(ctx context.Context, event *contact.EmployeeTypeEnumUpdatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.employee_type_enum.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.employee_type_enum.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.employee_type_enum.updated_v3"] = contact.NewEmployeeTypeEnumUpdatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnScopeUpdatedV3(handler func(ctx context.Context, event *contact.ScopeUpdatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.scope.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.scope.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.scope.updated_v3"] = contact.NewScopeUpdatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnUserCreatedV3(handler func(ctx context.Context, event *contact.UserCreatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.created_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.created_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.created_v3"] = contact.NewUserCreatedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnUserDeletedV3(handler func(ctx context.Context, event *contact.UserDeletedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.deleted_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.deleted_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.deleted_v3"] = contact.NewUserDeletedEventHandler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnUserUpdatedV3(handler func(ctx context.Context, event *contact.UserUpdatedEvent) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["contact.user.updated_v3"]
	if existed {
		panic("event: multiple handler registrations for " + "contact.user.updated_v3")
	}
	dispatcher.eventType2EventHandler["contact.user.updated_v3"] = contact.NewUserUpdatedEventHandler(handler)
	return dispatcher
}