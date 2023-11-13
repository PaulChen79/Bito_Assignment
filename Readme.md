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