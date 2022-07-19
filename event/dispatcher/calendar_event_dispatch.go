/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

// Package dispatcher code generated by oapi sdk gen
package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
)

func (dispatcher *EventDispatcher) OnP2CalendarChangedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarChangedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.changed_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.changed_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.changed_v4"] = larkcalendar.NewP2CalendarChangedV4Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2CalendarAclCreatedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarAclCreatedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.acl.created_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.acl.created_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.acl.created_v4"] = larkcalendar.NewP2CalendarAclCreatedV4Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2CalendarAclDeletedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarAclDeletedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.acl.deleted_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.acl.deleted_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.acl.deleted_v4"] = larkcalendar.NewP2CalendarAclDeletedV4Handler(handler)
	return dispatcher
}
func (dispatcher *EventDispatcher) OnP2CalendarEventChangedV4(handler func(ctx context.Context, event *larkcalendar.P2CalendarEventChangedV4) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["calendar.calendar.event.changed_v4"]
	if existed {
		panic("event: multiple handler registrations for " + "calendar.calendar.event.changed_v4")
	}
	dispatcher.eventType2EventHandler["calendar.calendar.event.changed_v4"] = larkcalendar.NewP2CalendarEventChangedV4Handler(handler)
	return dispatcher
}
