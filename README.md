# go-in-lambda
Go WebAPI hosted in AWS Lambda


# How to deploy to AWS Lambda
 1. On windows environment, need run Command line in BASH !!!:
 		$ GOARCH=amd64 GOOS=linux go build -o main main.go
 2. zip output main file into main.zip file
 3. Upload main.zip to lambda
 4. Set handler to main in Lambda configuration
