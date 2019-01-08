# todo app apply service-repository pattern

### Run this project
1. Setup mysql db which has `todo` table
```
type Todo struct {
    ID        int
    Name      string
    Task      string
    Status    int
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}
```
> Notice that table fields will follow snake_case

2. Setup environment by copy then rename `.env.example` to `.env`, fill out necessary VARIABLES

3. Run project
- `go build cmd/serve.go`
- `./serve`