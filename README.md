# Go UK Maps
HTTP Tile Server for VectorMapDistrict Open Source Data from the Ordnance Survey Data Hub

First import osdata using:
https://github.com/rockwell-uk/go-uk-maps-import

### Usage
```
go build go-uk-maps.go
./go-uk-maps -h
```

### Examples
```
./go-uk-maps -v -dbengine pgsql
./go-uk-maps -v -dbengine mysql -dbhost 127.0.0.1 -dbport 3307
./go-uk-maps -vv -dbengine sqlite -dbfolder ../go-uk-maps-import/db
```

### Run in Docker
```
docker build -t uk-maps .
docker run -p 8080:8080 -v "$PWD":/app uk-maps go run go-uk-maps.go -dbhost host.docker.internal
```

### OSData Copyright
All osdata is copyright © Crown: https://www.ordnancesurvey.co.uk/business-government/licensing-agreements/copyright-acknowledgements
* Contains OS data © Crown copyright [and database right] [year].
* Cynnwys data OS Ⓗ Hawlfraint y Goron [a hawliau cronfa ddata] OS [flwyddyn].

## Author
This software was engineered by [David Boyle](https://github.com/dbx123) @ [Rockwell Consultants Ltd.](https://www.rockwellconsultants.co.uk/)