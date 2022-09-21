package handler

import (
	"context"
	eventsrepo "github.com/cdsrx/et/service/sportsevents/db"
	sportsevents "github.com/cdsrx/et/service/sportsevents/proto"
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
	listEventResponse, err := s.eventsRepo.List(request.GetFilter())
	if err != nil {
		return err
	}
	response.Events = listEventResponse
	return nil
}
