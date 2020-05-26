Timezone lookup service

tzLookupService provide the timezone name using the latitude and longitude
it uses the data from the project 

https://github.com/evansiroky/timezone-boundary-builder

and the lookup from the library 

https://github.com/evanoberholster/timezoneLookup

for convenience a makefile target was created to download the data, unzip it and 
process the geojson to create a BoltDB database with the timezones.

to test the project:

```bash
go get -u https://github.com/jms/tzLookupService
```

