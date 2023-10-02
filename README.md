# Football Mock Database
Sample mock database application that can store football teams and players. Intended to be an example to demonstrate a repository design pattern. Exposed through a REST API.

## Instructions:
- Build and run the REST API
```
make
./bin/football-database-service-api  
```
- Create a team
```
curl -X POST \
-H "Content-Type: application/json" \
-d '{
    "Name": "Reading FC"
}'\
http://localhost:8080/createTeam\?id\=5
```
- Search for a team
```
curl http://localhost:8080/getTeam\?id\=5
```
- Create a player
```
curl -X POST \
-H "Content-Type: application/json" \
-d '{
    "age": 17,
    "FirstName": "Joe",
    "LastName": "Smith",
    "TeamID": 5,
    "Position": "Forward"
}' \
http://localhost:8080/createPlayer\?id\=1
```
- Search for a player
```
curl http://localhost:8080/getPlayer\?id\=1
```
