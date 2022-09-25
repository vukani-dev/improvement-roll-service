# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method, with the owners of this repository before making a change. 

If you are simply adding a category to be shared, this is not needed, and you can submit the request as long as tests are passing

## Pull Request Process (Changing the API)

1. Fork this repo
2. Make your changes
3. Update the README.md with details of the changes
    2a. Write a test in `main_test.go` to confirm your changes are working properly
4. If tests pass, submit the pull request and wait for approval

## Pull Request Process (Adding a category)

1. Fork this repo
2. Add your category json to the `categories` directory
3. Run `go test`. Resolve any issues if present
4. If tests pass, submit the pull request and wait for approval

Once your changes are in main, they are shortly after reflected in production

## Running the app locally

*Developed with go 1.18.1

- Clone the repo
- In the root directory, run `go run main.go`
    - The service with be hosted on `localhost:3000`

## Running the app Docker

- Clone the repo
- `docker build -t imp .`
- `docker run --name imp --publish 3000:3000 -it imp:latest`
    - you should be in the container now
- `./bin/server`
    - The service with be hosted on `localhost:3000`
