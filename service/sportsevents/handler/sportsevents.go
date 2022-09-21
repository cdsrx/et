package handler

import (
	"context"
	eventsrepo "github.com/cdsrx/et/service/sportsevents/db"
	sportsevents "github.com/cdsrx/et/service/sportsevents/proto"
	"strings"
)

type Events struct {
	eventsRepo *eventsrepo.EventsRepo
}

// Return a new handler
func New(eventsRepo *eventsrepo.EventsRepo) *Events {
	sportingEvent := &Events{eventsRepo: eventsRepo}

	return sportingEvent
}

func (s *Events) ListEvents(ctx context.Context, request *sportsevents.ListEventsRequest, response *sportsevents.ListEventsResponse) error {
	// Get order by param from the request
	var orderBy string
	orderByParam := strings.ToUpper(request.GetOrderBy())
	switch orderByParam {
	case "ASC":
		fallthrough
	case "DESC":
		orderBy = request.GetOrderBy()
	default:
		orderBy = "ASC" // Default order
	}

	listEventResponse, err := s.eventsRepo.List(request.GetFilter(), orderBy)
	if err != nil {
		return err
	}

	response.Events = listEventResponse
	return nil
}

func (s *Events) GetEvent(ctx context.Context, request *sportsevents.GetEventRequest, response *sportsevents.GetEventsResponse) error {
	event, err := s.eventsRepo.Get(request.GetId())
	if err != nil {
		return err
	}

	response.Event = event
	return nil
}
