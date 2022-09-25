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
    "date": string of datetime in RFC-3339 format
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
All parameters are optional, and can ofcourse be combined

- **page** = The page of results you want
- **search** = only return categories whos name contains the param
- **tags** = only return categories who have a tag that contains the param. This can also be comma delimited
- **author** = only return categories whos author contains the param

## Hosting
This repository is connected to a build pipeline that hosts the service being used in the app. 
Any changes to main will be reflected in production.