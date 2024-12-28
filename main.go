package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	timezone "github.com/evanoberholster/timezoneLookup"
)

// Location point latitude and longitude
type Location struct {
	Latitude  float32 `form:"lat" binding:"required,gte=-85,lte=85"`
	Longitude float32 `form:"lon" binding:"required,gte=-180,lte=180"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		var location Location
		if err := c.ShouldBindWith(&location, binding.Query); err == nil {
			res, err := lookupTz(location.Latitude, location.Longitude)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "timezone not found"})
			} else {
				c.JSON(http.StatusOK, gin.H{"timezone": res})
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	_ = router.Run()
}

func lookupTz(Lat float32, Lon float32) (string, error) {
	encoding, err := timezone.EncodingFromString("msgpack")
	if err != nil {
		log.Println(err)
	}

	tz, err := timezone.LoadTimezones(timezone.Config{
		DatabaseType: "boltdb",   // memory or boltdb
		DatabaseName: "timezone", // Name without suffix
		Snappy:       true,
		Encoding:     encoding, // json or msgpack
	})

	if err != nil {
		log.Println(err)
	}

	res, err := tz.Query(timezone.Coord{Lat: Lat, Lon: Lon})
	tz.Close()
	return res, err
}
