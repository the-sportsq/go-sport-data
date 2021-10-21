# go-sport-data

go-sport-data is a wrapper for [Sport Data API](https://sportdataapi.com/)

### Examples

#### Get Matches
```go
matches, err := gsd.GetMatches(352)
```

#### Get Teams
```go
teams, err := gsd.GetTeams(48)
```

### Running Tests

Copy example `.env` file, and set your Sport Data API key:

```bash
cp env.example .env
vi .env
```

Then you should be good to run all tests with the command:

```bash
go test
```

