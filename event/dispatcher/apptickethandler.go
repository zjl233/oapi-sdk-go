package dispatcher

import (
	"context"
	"time"

	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/event"
)

type appTicketEventData struct {
	AppId     string `json:"app_id"`
	Type      string `json:"type"`
	AppTicket string `json:"app_ticket"`
}

type appTicketEvent struct {
	*event.EventBase
	Event *appTicketEventData `json:"event"`
}

type appTicketEventHandler struct {
	event *appTicketEvent
}

func (h *appTicketEventHandler) Event() interface{} {
	h.event = &appTicketEvent{}
	return h.event
}

func (h *appTicketEventHandler) Handle(ctx context.Context, event interface{}) error {
	appTicketEvent := event.(*appTicketEvent)
	return core.GetAppTicketManager().Set(context.Background(),
		appTicketEvent.Event.AppId,
		appTicketEvent.Event.AppTicket, time.Hour*12)
}