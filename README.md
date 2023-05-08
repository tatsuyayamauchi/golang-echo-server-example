# golang-echo-server-example

## run server
```bash
make run
```

## Login request
```bash
token=$(curl -X POST --data-urlencode 'username=test' -d 'password=user' http://localhost:8080/v1/login | jq '.token' | tr -d '"')
```

## Authenticated request
```bash
curl -H "Authorization: Bearer $token" http://localhost:8080/v1/hello
curl -H "Authorization: Bearer $token" http://localhost:8080/v1/hello2
```

# Benchmark

### terminal1
```bash
make run
```

### terminal2
```bash
# see: https://github.com/tsliwowicz/go-wrk
go-wrk -M POST -body 'username=test&password=user' -H 'Content-Type:application/x-www-form-urlencoded' -d 500 http://localhost:8080/v1/login
```

### terminal3
```bash
go tool pprof -raw -output=prof.cpu -seconds=5 http://localhost:8080/debug/pprof/profile

# see: https://formulae.brew.sh/formula/flamegraph
stackcollapse-go.pl prof.cpu| flamegraph.pl > prof.cpu.svg
```
