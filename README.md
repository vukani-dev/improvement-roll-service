# Improvement Roll Sharing Service

This is an API made to easily share categories made on the app [Improvement Roll](https://github.com/vukani-dev/improvement-roll)

Made with Go

## API

### Models
The every call to the API will result in the same object:

```json
{
    "sharedCategories" : [SharedCategory],
    "page" : number,
    "totalPages" : number
}
```

This further breaks down into the following Objects:
```json
SharedCategory:
{
    "category" : {Category},
    "tags" : [string],
    "author" : string,
    "date": string of datetime in MM-DD-YYYY format
}

Category:
{
    "name": string,
    "timeSensitive: bool,
    "description": string,
    "tasks": [Task]
}

Task:
{
    "name" : string,
    "desc" : string,
    "minutes" : number
}
```

### Routes
Every response will be pages by 10 records. So in order to get the next few simply increment the `page` query parameter

##### GET /?page={x}&search={y}&tags={z}&author={p}
All parameters are optional, and can of course be combined

- **page** = The page of results you want
- **search** = only return categories who's name contains the param
- **tags** = only return categories who have a tag that contains the param. This can also be comma delimited
- **author** = only return categories who's author contains the param

## Hosting

### CI/CD
This repository is connected to a build pipeline that hosts the service being used in the app. 
Any changes to main will be reflected in production.


### Running the app locally

*Developed with go 1.18.1

- Clone the repo
- In the root directory, run `go run main.go`
    - The service with be hosted on `localhost:3000`

### Running the app Docker

- Clone the repo
- `docker build -t imp .`
- `docker run --name imp --publish 3000:3000 -it imp:latest`
    - you should be in the container now
- `./bin/server`
    - The service with be hosted on `localhost:3000`
