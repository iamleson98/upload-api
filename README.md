Golang upload api

Doc:

  +) run: go run main.go
  +) to get token: curl http://localhost:8000/token
  +) to upload .zip file: curl --form -F file=@[path_to_zip_file] -H 'Authentication:[token_get_from_above]' http://localhost:8000/upload
