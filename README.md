# Flight app

1. Copy the contents of `.env.dist` to `.env`
2. Run `docker-compose up -d --build` to start the project locally
4. Visit the documentation page: http://localhost:9000


### Database migrations

#### Install golang-migrations
```
brew install golang-migrate
```
from: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Run migrations
```
 make migrate
```

# Build
```
 make
```

# Run
```
 make run
```

# API requests 

## Add flights

```
curl -X "POST" "http://localhost:9000/v1/flight" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
    "flights": [
        {
            "number": 17,
            "city_departure": "Los Angeles",
            "city_arrival": "New York",
            "time_departure": "2022-01-30T13:00:00.000Z",
            "time_arrival": "2022-01-31T13:00:00.000Z"
        },
        {
            "number": 142,
            "city_departure": "Chicago",
            "city_arrival": "Houston",
            "time_departure": "2022-02-01T13:00:00.000Z",
            "time_arrival": "2022-02-02T13:00:00.000Z"
        },
        {
            "number": 77,
            "city_departure": "Arizona",
            "city_arrival": "Pennsylvania",
            "time_departure": "2022-02-03T13:00:00.000Z",
            "time_arrival": "2022-02-04T13:00:00.000Z"
        },
        {
            "number": 92,
            "city_departure": "Texas",
            "city_arrival": "Florida",
            "time_departure": "2022-01-30T13:00:00.000Z",
            "time_arrival": "2022-01-31T13:00:00.000Z"
        },
        {
            "number": 9,
            "city_departure": "Indiana",
            "city_arrival": "California",
            "time_departure": "2022-01-30T13:00:00.000Z",
            "time_arrival": "2022-01-31T13:00:00.000Z"
        }
    ]
}
'
```
## List of flights

```
curl "http://localhost:9000/v1/flight" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

## List of flights sorted by number with merge sort algorithm

```
curl "http://localhost:9000/v1/flight?sort=number&limit=100" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

## List of flights filter by city departure

```
curl "http://localhost:9000/v1/flight?city_departure=Los Angeles" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

