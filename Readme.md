# This is assignment fir bito interview test

## Run test

To run test, clone this project to your local and cd into the folder,

run `go test ./internal/service  -coverpkg=./internal/service -coverprofile ./coverage.out`

It will generate a `coverage.out` file for you to see coverage detail.

then run `go tool cover -func ./coverage.out` to see the details of coverage

## Build docker image

To Build Docker image, you just run `docker build ./Dockerfile` to build the image

## Run docker image

Once you build the image, you can use `docker run -d <imageID> -p 5050:5050` to run the api image.

# API documentation

## AddSinglePersonAndMatch

*Add a new user to the matching system and find any possible matches for the new user.*


- API path: `/add_and_match`
- API method: **POST**
- API request body:
    ```json
        {
            "name": "Ken", // string
            "height": 180, // integer gte 0
            "gender": "male", // string only accept: female, male
            "wanted_dates": 12 // integer gte 0
        }
    ```
- API response body:
    ```json
        // success
        {
            "code": 200, // response code, not http status code
            "data": [
                {
                    "Name": "Sandra",
                    "Height": 167,
                    "Gender": "female",
                    "RemainingDates": 11
                }
            ]
        }

        // failed
        {
            "code": 500,
            "message": "User already exists",
            "time": "1699872014"
        }
    ```

## RemoveSinglePerson

*Remove a user from the matching system so that the user cannot be matched anymore.glePersonAndMatch*


- API path: `/:user_name/remove`
- API method: **DELETE**
- API response body:
    ```json
        // success
        {
            "code": 200, // api code, not http status code
            "data": null
        }

        // failed
        {
            "code": 500,
            "message": "User not found", // api code, not http status code
            "time": "1699872429"
        }
    ```

## QuerySinglePeople

*Find the most N possible matched single people, where N is a request parameter.*


- API path: `/query_single_person/:user_name/:number`
- API method: **GET**
- API response body:
    ```json
        // success
        {
            "code": 200,
            "data": [
                {
                    "Name": "Sandra",
                    "Height": 167,
                    "Gender": "female",
                    "RemainingDates": 10
                },
                {
                    "Name": "Lisa",
                    "Height": 172,
                    "Gender": "female",
                    "RemainingDates": 10
                }
            ]
        }

        // failed
        {
            "code": 500,
            "message": "User not found",
            "time": "1699872014"
        }
    ```
