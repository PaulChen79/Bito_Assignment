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

*Remove a user from the matching system so that the user cannot be matched anymore.*


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

*Find the most N possible matched single people, where N is a request parameter. It will return random N matches, if N > all matches then return all matches.*


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

# System design document

I use sort of DDD structure to decouple the service and repo instance.

For in-memory storage, I use hash map to store the data.

I use GIN web framework to build this API server, the flow is as below.

1. **handler** get the http request and only handle data valid stuff then pass data to **service**.

2. **service** handle the business logic, once you have to do CRUD on data, you call **repo** service

3. **repo** handle all operation of data.

## Time complexity of all API

### AddSinglePersonAndMatch

>Time Complexity: **O(N)** Space Complexity: **O(N)**

In this function, there is only one loop that iterates over the entire userPool to find all eligible matches. Since other operations are performed using a hash map, they are essentially O(1). So, the overall time complexity of this function can be concluded as O(N).

The function needs to store a slice `matches := []*domain.User{}` to return the response. In the worst-case scenario, it might include all users in the userPool. Other operations do not require additional space. Therefore, the space complexity of this function is O(N).

### RemoveSinglePerson

>Time Complexity: **O(1)** Space Complexity: **O(1)**

In this function, since there is only a delete operation on the map, the time complexity is O(1). Furthermore, only one temporary variable is used, which is a pointer, and it does not change in size depending on the size of the userPool. So, the space complexity is also O(1).

#### QuerySinglePeople

>Time Complexity: **O(N)** Space Complexity: **O(N)**

In this function, there is only one for loop that iterates over the entire userPool to first find all the suitable matches, so this part is O(N). Subsequently, there is a random shuffling of the order of matches, which is O(M). However, in the worst-case scenario, this is O(N), as the entire UserPool might qualify as matches. Therefore, the time complexity is O(2N), which simplifies to O(N).

The space complexity is also O(N) because the matches list could potentially include the entire userPool. If the requested number exceeds the size of the userPool, the entire matches list would be returned. Hence, the space complexity is also O(N).