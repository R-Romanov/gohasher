## Use as web server ##

1. Build the app
    ```bash
    go build -o ./build/gohasher ./cmd/main.go
    ```

1. Run local server

    ```bash
    ./build/gohasher
    ```
    
    Run server on certain port:

    ```bash
    ./gohasher -port=9090
    ```
    _* Default port is 8080_
    
### Api ###

Get hashed string:

``` http://localhost:8080/get-hash/?data=...&method=... ```

Where:
- data - String for hashing
- method - Hash method



Get available hash methods:

``` http://localhost:8080/hash-methods-list/ ```

---

If the request is successful, response will be in json format:

```{"hash":-8082402883035101391}```

If the request is not successful, response will be in json format with error message:

```{"error":"unknown hash method"}```

```{"error":"url param 'data' is missing"}```

## Use as package ##

1. Import package:
 
        import "github.com/R-Romanov/gohasher/pkg/hasher"

2. Call hashing function:

        h, err := hasher.Hash("example.com")
