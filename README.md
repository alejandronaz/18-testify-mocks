# Testing

## Service
1. Get the service code coverage report by running:
```bash
go test -cover ./app/internal/service -coverprofile=coverage_service.out
```
2. Generate the function report by running:
```bash
go tool cover -func=coverage_service.out
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
go tool cover -html=coverage_service.out
```

## Repository
1. Get the repository code coverage report by running:
```bash
go test -cover ./app/internal/repository -coverprofile=coverage_repository.out
```
This includes the mock repository, which it should be not considered in the coverage report.
To avoid this, we can exclude the mock files and generate the report again:
```bash
more coverage_repository.out | grep -v "mock" > coverage_repository_filtered.out && rm coverage_repository.out && mv coverage_repository_filtered.out coverage_repository.out
```
> It takes the content of the file and deletes de lines with "mock" (the flag `-v` is used to invert the match, so it will exclude the lines that contain the word "mock"). Then, it creates a new file with the filtered content and deletes the original file.

2. Generate the function report by running:
```bash
go tool cover -func=coverage_repository.out
```
Output:
```bash
app/app/internal/repository/movie.go:5:         NewMovieRepository      100.0%
app/app/internal/repository/movie.go:13:        Save                    100.0%
app/app/internal/repository/movie.go:23:        Get                     0.0%
total:                                          (statements)            55.6%
```
3. Generate the HTML report by running:
```bash
go tool cover -html=coverage_repository.out
```