# Hacker News API Client — hnapi

This library exposes the Hacker News API, which is provided by Firebase, in an
attempt to make the public Hacker News data available in near real time.

## Examples

### Get User

```go
username, err := reader.ReadString('\n')
if err != nil {
	log.Panicln(err.Error())
}

user := hnapi.GetUser(username)
fmt.Println(user.ID)
```

### Get Top Story IDs

```go
topStores := hnapi.RetrieveTopStoriesItemNumbers()
```

### Print a top story from it's ID

```go
item := hnapi.GetItem(tsNumber)
fmt.Prinln(item.Title)
```
