
# International Classification Of Diseases Codes -  (ICD10) 

A RESTful API of internationally recognized set of diagnosis codes, diseases, signs, symptoms, abnormal findings, compliants and causes. This API centers specifically on the 10th revision which is tagged ICD-10. 

For further info please refer to: https://en.wikipedia.org/wiki/ICD-10


## Setup

To run this project, please refer to the below listed guide:


- First Clone the git repo
- Then cd into the directory 
- Next, build the docker image by running
- Finally run the docker image. 

```
    $ git clone https://github.com/otyang/icd-10.git
    $ cd icd-10/
    $ docker build -t icd10-go-app .  
    $ docker run -p 3000:3000 -t icd10-go-app
```

At this time, you have a RESTful API server running at http://127.0.0.1:3000. It provides the following endpoints:
## API Reference

#### Response structure
Conventional HTTP response codes are used to indicate the success or failure of an API request. 

| Status Code | Description|
| --- | --- | 
| 200 - OK | Everything worked as expected|
| 201 - Created | Resource was created successfully |
| 400 - Bad Request| Request was unacceptable, often due to missing a required parameter. | 
| 404 - Not Found | The requested resource doesn't exist  | 
| 405 - Not Allowed | Create a new diagnosis code record |
| 409 - Conflict | The request conflicts with another request | 
| 500 - Server Error | Very rare. It means something went wrong on code end. Contact me | 


#### Error response body
````
{
  "success": false,
  "message": "error message",
  "errorCode": "internal_server_error"
}
````


####  Endpoints
| HTTP Verbs | Endpoints | Action |
| --- | --- | --- |
| GET | / | Home or welcome page to the API |
| GET | /icd/:fullCode | Retrieve diagnosis codes by ID  | 
| GET | /icd?limit=20&cursor=A0101 | List diagnosis codes in batches of 20  | 
| POST | /icd | Create a new diagnosis code record |
| PUT | /icd/:fullCode | Edit an existing diagnosis code record | 
| DELETE | /icd/:fullCode | Delete a diagnosis code by ID |
| POST | /icd-upload | for uploading ICD CSV files with up to 10K diagnosis code records |



 


## Project Layout, Architectural Considerations & Tech stack 

The code uses the following project layout:
 
```
.
├── cmd                  main entry point applications of the project
│   └── icd              Entry point for the icd service
├── internal             private application and library code
│   ├── event            pubsub or event related library
│   ├── icd              icd features: entities, handlers, repository & likes
│   ├── entity           entity definitions and domain logic
├── pkg                  public library code. 
    ├── config           configuration library.  
    ├── datastore        helpers for working with database
    ├── middleware       middleware related handlers library
    ├── logger           structured and context-aware logger
    └── response         handles http response, errors and request
    └── validators       helpers to efficiently handle validation 
```

The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example, 
the `icd` directory contains the application logic related with the icd-10 feature. 

Within each feature package, code are organized in layers (API, entity, repository, handlers-for-http, handlers-for-events), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

 
* **[Golang](https://go.dev/)** - An open-source programming language supported by Google. 
* **[Sqlite Database](https://sqlite.org/)** - This is a small, fast, self-contained, high-reliability, portable and most-used, SQL database engine. I selected this to reduce dependencies and for portability.

## Acknowledgment
[@kamillamagna](https://github.com/kamillamagna) - His csv file on the icd10 codes, was what i converted and loaded into a sqlite database.

## Authors

- [@otyang](https://www.github.com/otyang)

## License

[MIT](https://choosealicense.com/licenses/mit/) - This project is available for use under the MIT License.
