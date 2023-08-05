Timezone lookup service

tzLookupService is a hobby project it provide the timezone name 
using the latitude and longitude is written in go language and expose 
his functionality via web using GIN framework

It also uses the data from the project 

https://github.com/evansiroky/timezone-boundary-builder

and the lookup of the data from the library

https://github.com/evanoberholster/timezoneLookup

The web service runs with Gin framework https://gin-gonic.com/

for convenience a makefile target was created to download the data, unzip it and 
process the geojson to create a BoltDB database with the timezones.

the script to create the timezone db comes from timezoneLookup "cmd/timezone.go" 
 
to run the project:

```bash
go get -u https://github.com/jms/tzLookupService
cd ${GOPATH}/src/github.com/jms/tzLookupService
make create-tz-db
make run
curl 'localhost:8080?lat=12&lon=-86'

{"timezone":"America/Managua"}
```

Configuration
```bash
# changing to release mode and set a different port 
export GIN_MODE=release
export PORT=80
```

To run the project on a container, you can use podman or docker
```bash
podman build -f Dockerfile -t tzlookup
podman run --rm -p 8080:8080 tz_lookup
```



