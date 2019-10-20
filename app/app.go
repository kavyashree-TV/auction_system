package app

import (
	"./handlers/auctionsystem"
	"./models"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//App has router and db instances
type App struct{}

var auctionService = auction.System{}

//Init initializes the app with predefined configuration
func (app *App) Init() {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	routes.Use(gzip.Gzip(gzip.DefaultCompression)) //compresses the response time
	routes.Use(gin.Recovery()) //Auto recovers from panic mode
	auctionSystem := routes.Group("auction_system")
	{
		auctionSystem.POST("/auction", app.Auctioneer)
	}
	routes.Run(":8003")
}

//Auctioneer : ...
func (app *App) Auctioneer(c *gin.Context) {
	var auctionSystem models.AuctionSystem
	if c.BindJSON(&auctionSystem) == nil {
		c.SecureJSON(auctionService.AuctionSystem(auctionSystem.AuctionID))
	}
}
