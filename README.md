Golang upload api

Doc:

  +) run: go run main.go <br />
  +) to get token: curl http://localhost:8000/token <br />
  +) to upload .zip file: curl --form -F file=@<path_to_zip_file> -H 'Authentication:<your_token>' http://localhost:8000/upload

NOTE: Change <path_to_zip_file> and <your_token> to your.