package db

import (
	"database/sql"
	sportsevents "github.com/cdsrx/et/service/sportsevents/proto"
	"github.com/golang/protobuf/ptypes"
	"strings"
	"sync"
	"time"
)

// EventsRepository provides repository access to events.
type EventsRepository interface {
	// Init will initialise our events repository.
	Init() error

	// List will return a list of events.
	List(filter *sportsevents.ListEventsRequestFilter) ([]*sportsevents.Event, error)
}

type EventsRepo struct {
	db   *sql.DB
	init sync.Once
}

// New returns a new initialised EventsRepo
func New(db *sql.DB) (*EventsRepo, error) {
	eventsRepo := &EventsRepo{db: db}
	err := eventsRepo.Init()
	if err != nil {
		return nil, err
	}
	return eventsRepo, nil
}

// Init prepares the Event repository dummy data.
func (r *EventsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy Events.
		err = r.seed()
	})

	return err
}

func (r *EventsRepo) List(filter *sportsevents.ListEventsRequestFilter) ([]*sportsevents.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventQueries()[eventsList]

	query, args = r.applyFilter(query, filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanEvents(rows)
}

func (r *EventsRepo) applyFilter(query string, filter *sportsevents.ListEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

func (m *EventsRepo) scanEvents(
	rows *sql.Rows,
) ([]*sportsevents.Event, error) {
	var events []*sportsevents.Event

	for rows.Next() {
		var event sportsevents.Event
		var advertisedStart time.Time

		if err := rows.Scan(&event.Id, &event.MeetingId, &event.Name, &event.Number, &event.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		event.AdvertisedStartTime = ts

		events = append(events, &event)
	}

	return events, nil
}
