# Golang Gin API with Bun ORM




## Requirements

* x86-64
* Linux/Unix
* [Golang](https://go.dev/)
* [Docker](https://www.docker.com/products/docker-desktop/)

## Startup

The script "up" creates our DB container, compiles and executes our binary:
```
1. docker-compose -f db/cars/docker-compose.yml up -d
2. go build -o gin_api src/main.go
3. ./gin_api
```

## Shutdown

The script "down" removed our DB container:
```
1. docker-compose -f db/cars/docker-compose.yml down
```

## Postman Collection

The repository includes a Postman collection in the 'postman' directory.


## Load Testing using Bombardier
In this project we utilize [Bombardier](https://github.com/codesenberg/bombardier) to conduct load-testing.

The subject of our tests is the `/car-details` endpoint, which is responsible for retrieving various vehicle-related data from different repositories and responding with an aggregate object.

### Test 1: 10 Concurrent Connections, 1000 Requests (Baseline)
```plaintext
Command: bombardier -m GET localhost:8080/car-details/1 -c 10 -n 1000
Statistics:
  Reqs/sec: 243.22
  Latency: 41.86ms (Avg)
  HTTP codes: 2xx - 1000
  Throughput: 202.82KB/s
  ```

### Test 2: 100 Concurrent Connections, 1000 Requests
```plaintext
Command: bombardier -m GET localhost:8080/car-details/1 -c 100 -n 1000
Statistics:
  Reqs/sec: 1917.87 (↑689.74% from Test 1)
  Latency: 47.26ms (↑12.92% from Test 1)
  HTTP codes: 2xx - 1000
  Throughput: 1.60MB/s (↑687.62% from Test 1)
```

### Test 3: 1000 Concurrent Connections, 2000 Requests
```plaintext
Command: bombardier -m GET localhost:8080/car-details/1 -c 1000 -n 2000
Statistics:
  Reqs/sec: 1244.00 (↑411.63% from Test 1)
  Latency: 742.03ms (↑1676.64% from Test 1)
  HTTP codes: 2xx - 2000
  Throughput: 1.02MB/s (↑404.95% from Test 1)
```

## Test Result Analysis

### Baseline (C10, N1000)
The baseline test (Test 1), which serves as our reference point, simulates a moderate load with 10 concurrent connections and 1000 requests. In this scenario:

* The application achieved a request rate of 243.22 requests per second.
* The average latency was measured at 41.86ms.
* All 1000 requests resulted in successful HTTP 2xx responses.


### Test 2 (C100, N1000)
In Test 2, where 100 concurrent connections and 1000 requests were applied:

* The request rate increased significantly by 689.74% compared to the baseline (Test 1).
* Latency increased slightly by 12.92% compared to the baseline.
* Throughput exhibited an enormous increase of 687.62% compared to that of the baseline.


### Test 3 (C1000, N2000)
In Test 3, with 1000 concurrent connections and 2000 requests:

* The request rate maintained a relatively high level, increasing by 411.63% compared to the baseline.
* The average latency saw a substantial increase of 1676.64% compared to the baseline.
* Throughput showed a considerable increase of 404.95% compared to the baseline.
* This test also successfully processed all 2000 requests with HTTP 2xx responses.

## Conclusion
These findings suggest that this application performs better at moderate to high concurrent loads, where C=100 seems to be the sweet spot in terms of throughput.