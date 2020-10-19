Work on pet project (Music Review Tool)
=======================================

Technologies:
Server: Golang via nginx;
Client: React/CSS3
DB: PostgreSQL

* Integration with API like LastFm
* User review for artist/album
* Categorize music (Use open AI API)
* Jenkins CI

#How to run it

Run at terminal
```
docker-compose build
docker-compose up -d
```
and open at browser http://localhost:8080 to see UI. Access to database can be throw adminer: http://localhost:8081


#Run go test
```
cd backend/scr
go test ./... -count=1 -cover
```
