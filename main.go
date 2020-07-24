package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)
import timezone "github.com/evanoberholster/timezoneLookup"

var tz timezone.TimezoneInterface

// Location point latitude and longitude
type Location struct {
	Lat float32 `form:"lat" binding:"required"`
	Lon float32 `form:"lon" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/tz", func(c *gin.Context) {
		var location Location
		if err := c.ShouldBindWith(&location, binding.Query); err == nil {
			res, err := lookupTz(location.Lat, location.Lon)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "timezone not found"})
			}
			c.JSON(http.StatusOK, gin.H{"timezone": res})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	r.Run()
}

func lookupTz(Lat float32, Lon float32) (string, error) {
	tz, err := timezone.LoadTimezones(timezone.Config{
		DatabaseType: "boltdb",   // memory or boltdb
		DatabaseName: "timezone", // Name without suffix
		Snappy:       true,
		Encoding:     "msgpack", // json or msgpack
	})
	if err != nil {
		log.Println(err)
	}
	res, err := tz.Query(timezone.Coord{Lat: Lat, Lon: Lon})
	tz.Close()

	if err != nil {
		log.Println(err)
	}
	return res, err
}
