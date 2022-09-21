## Sports Events Service

### Getting Started

#### 1. Install Go (latest).

```bash
brew install go
```

... or [see here](https://golang.org/doc/install).

#### 2. Install micro : see this [getting started](https://micro.dev/getting-started#install) link


#### 3. Start micro
```bash
micro server
```

#### 4. In another terminal window, start the sports events service by running micro and passing the desired service name. e.g. _racing_
```bash
micro run -name racing service/sportsevents
```
This will create a standalone service with its own temporary database.

To provide a path for database, set the DBPATH environment variable.

```bash
micro run --env_vars DBPATH="/path/to/racing.db" --name racing service/sportsevents
```

Add `--env_vars DBSEED="true"` to seed the database with dummy data

```bash
micro run --env_vars DBPATH="/path/to/racing.db" --env_vars DBSEED="true" --name racing service/sportsevents
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

#### 5. Make a request for racing events... 

Request with no filters:
```bash
curl -X "POST" "http://localhost:8080/racing/SportsEvents/listEvents"
```

Request with visible filter (true/false):
```bash
curl -X "POST" "http://localhost:8080/racing/SportsEvents/listEvents" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"visible":true}
}'
```

Request with order by parameter (ASC/DESC default ASC):
```bash
curl -X "POST" "http://localhost:8080/racing/SportsEvents/listEvents" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"visible":false},"orderBy":"DESC"
}'
```

Getting an event via an event ID
```bash
curl -X "POST" "http://localhost:8080/racing/SportsEvents/getEvent" \
     -H 'Content-Type: application/json' \
     -d $'{
  "id":10
}'
```

#### 6. Running multiple services

To run multiple instances of the events service, simply run the micro command with different service names

Example:

Start the Formule One service
```bash
micro run --env_vars DBPATH="/path/to/formulaone.db" --env_vars DBSEED="true" --name formulaone service/sportsevents
```

Start the MotoGP service
```bash
micro run --env_vars DBPATH="/path/to/motogp.db" --env_vars DBSEED="true" --name motogp service/sportsevents
```

Use the service name to make a request to a specific service.

Request to the Formula One service
```bash
curl -X "POST" "http://localhost:8080/formulaone/SportsEvents/listEvents" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"visible":true},"orderBy":"ASC"
}'
```

Request to the MotoGP service
```bash
curl -X "POST" "http://localhost:8080/racing/motogp/getEvent" \
     -H 'Content-Type: application/json' \
     -d $'{
  "id":35
}'
```


#### 7. Stop the service
```bash
micro kill racing
```
```bash
micro kill formulaone
```
```bash
micro kill motogp
```

