# golang-echo-server-example

run server
```bash
make run
```

Login request
```bash
token=$(curl -X POST --data-urlencode 'username=test' -d 'password=user' http://localhost:8080/v1/login | jq '.token' | tr -d '"')
```

Authenticated request
```bash
curl -H "Authorization: Bearer $token" http://localhost:8080/v1/hello
curl -H "Authorization: Bearer $token" http://localhost:8080/v1/hello2
```
