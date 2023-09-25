# Blogging API

This is a simple API that supports CRUD operations on a list of articles (initially empty)

## Usage

To use the API, just clone the repo and make sure you have Golang installed. 
Run `go build -o output` to download all dependencies needed and get an executable file.
Run the output executable file.

## Documentation

### Endpoints

Main endpoint: `/articles`

- Method GET `/get/all` : returns a list of all available articles
- Method GET `/get/{id}` : returns the content of the article with the specified id
- Method POST `/create` : creates a new article with the provided body
- Method DELETE `/delete/{id}` : deletes the specified article
- Method PUT `/update/{id}` : replaces the old article specified by the id with the new attributes provided in the body 

## License

This project is licensed under the GPL3 license
