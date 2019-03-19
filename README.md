## Usage as web service ##

1. Build app
    ```bash
    go build -o ./build/gohasher ./cmd/main.go
    ```
1. Run app
    ```bash
    cd ./build
    ./gohasher
    ```
    
    Or run app on certain port:
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

## Usage as golang package ##

1. Import package "github.com/R-Romanov/gohasher" with dep
2. Use gohasher.HashString() to hash string with certain hash method
```gohasher.HashString("any-string", "fnv1")```
