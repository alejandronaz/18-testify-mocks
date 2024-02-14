# Testing
1. Get the code coverage report by running:
```bash
go test -cover ./app/internal/service -coverprofile=coverage.out
```
2. Generate the function report by running:
```bash
go tool cover -func=coverage.out
```
Output:
```bash
app/app/internal/service/movie.go:5:    NewMovieService 100.0%
app/app/internal/service/movie.go:15:   Save            100.0%
app/app/internal/service/movie.go:34:   Get             0.0%
total:                                  (statements)    85.7%
```
3. Generate the HTML report by running:
```bash
go tool cover -html=coverage.out
```
