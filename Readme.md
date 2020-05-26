Timezone lookup service

tzLookupService provide the timezone name using the latitude and longitude
it uses the data from the project 

https://github.com/evansiroky/timezone-boundary-builder

the service runs with Gin framework https://gin-gonic.com/

and the lookup from the library 

https://github.com/evanoberholster/timezoneLookup

for convenience a makefile target was created to download the data, unzip it and 
process the geojson to create a BoltDB database with the timezones.

the script to create the timezone db comes from timezoneLookup "cmd/timezone.go" 
 
to test the project:

```bash
go get -u https://github.com/jms/tzLookupService
cd ${GOPATH}/src/github.com/jms/tzLookupService
make create-tz-db
make run
curl 'localhost:8080/tz?lat=12&lon=-86'
{"timezone":"America/Managua"}

# for production, change to release mode and set a different port 
export GIN_MODE=release
export PORT=80
```

