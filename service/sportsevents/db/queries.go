package db

const (
	eventsList = "list"
	event      = "event"
)

func getEventQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM events
		`,
		event: `
			SELECT
				id,
				meeting_id,
				name,
				number,
				visible,
				advertised_start_time
			FROM events
			WHERE id = ?`,
	}
}
