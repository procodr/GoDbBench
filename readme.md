# Benchmark Some Db Engines

## 🚀 How to run

### 1️⃣ Run Back-end

Start database servers: `docker compose -f dbs.yaml up`

Start API node: `cd api && go run .`

### 2️⃣ Run Benchmarks using WRK

Unix:

  - PostgreSQL: `wrk -t1 -c5 -d10s --timeout 1 -s benchmark.lua http://127.0.0.1:3000/pgTest`
  - MongoDB: `wrk -t1 -c5 -d10s --timeout 1 -s benchmark.lua http://127.0.0.1:3000/mnTest`
  - SQLite: `wrk -t1 -c5 -d10s --timeout 1 -s benchmark.lua http://127.0.0.1:3000/sqTest`
  - Pebble: `wrk -t1 -c5 -d10s --timeout 1 -s benchmark.lua http://127.0.0.1:3000/pbTest`

Windows (using WRK docker image):

  - PostgreSQL

`docker run --rm --network=host -v $PWD/benchmark.lua:/benchmark.lua:ro williamyeh/wrk -t1 -c5 -d10s --timeout 1 -s /benchmark.lua http://host.docker.internal:3000/pgTest`

  - MongoDB:
 
`docker run --rm --network=host -v $PWD/benchmark.lua:/benchmark.lua:ro williamyeh/wrk -t1 -c5 -d10s --timeout 1 -s /benchmark.lua http://host.docker.internal:3000/mnTest`

  - SQLite:

`docker run --rm --network=host -v $PWD/benchmark.lua:/benchmark.lua:ro williamyeh/wrk -t1 -c5 -d10s --timeout 1 -s /benchmark.lua http://host.docker.internal:3000/sqTest`

  - Pebble:

`docker run --rm --network=host -v $PWD/benchmark.lua:/benchmark.lua:ro williamyeh/wrk -t1 -c5 -d10s --timeout 1 -s /benchmark.lua http://host.docker.internal:3000/pbTest`


## 📇 Results

| Database Engine | Result (req/sec) [1 Thread - 5 Connections] | Result (req/sec) [100 Threads - 500 Connections] |
|-----------------|:-------------------------------------------:|:------------------------------------------------:|
| PostgreSQL      |                     1430                    |                       1501                       |
| MongoDB         |                     3345                    |                       5695                       |
| SQLite          |                     3911                    |                        841                       |
| Pebble          |                     600                     |                       57050                      |