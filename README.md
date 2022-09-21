## Sports Events Service

### Getting Started

1. Install Go (latest).

```bash
brew install go
```

... or [see here](https://golang.org/doc/install).

2. Install micro : see this [getting started](https://micro.dev/getting-started#install) link


3. Start micro
```bash
micro server
```

4. In another terminal window, start the sports events service by running micro and passing the desired service name. e.g. _racing_
```bash
micro run -name racing service/sportsevents
```
This will create a standalone service with its own temporary database.

To provide a path for database, set the DBPATH environment variable.

```bash
micro run --env_vars DBPATH="/path/to/racing.db" --name racing service/sportsevents
```

Make sure the service is running by checking its status
```bash
$ micro status
NAME    VERSION SOURCE                    STATUS  BUILD   UPDATED METADATA
racing  latest  /path/to/sportsevents     running n/a     4s ago  owner=admin, group=micro
```

The service logs can be checked using this command
```bash
micro logs -f racing
```

5. Make a request for racing events... 

```bash
curl -X "POST" "http://localhost:8080/racing/SportsEvents/listEvents" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {}
}'
```
