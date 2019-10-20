package auction

import (
	"net/http"

	"../../infrastructure"
	"../../models"
	"../../utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

//System ....
type System struct{}

var mongoInfra = infrastructure.MongoDBInfra{}
var _logger = utils.GetLogger()

//AuctionSystem : Api for selecting highest bidder
//Params : auction_id(string)
//Author : Kavya
//Date :20/10/2019
func (as *System) AuctionSystem(auctionID string) (int, map[string]interface{}) {
	var auctionSystem []models.AuctionSystem
	var bidderAmount []float64
	var ids []primitive.ObjectID
	status := http.StatusBadRequest
	response := map[string]interface{}{"status": status, "message": "Bidders not found"}
	err := mongoInfra.FindMany("auctions", bson.M{"auction_id": auctionID}, &auctionSystem)
	if err == nil && len(auctionSystem) > 0 {
		for _, bidder := range auctionSystem {
			ids = append(ids, bidder.ID)
			bidderAmount = append(bidderAmount, bidder.BidAmount)
		}
		max, index := as.FindMax(bidderAmount)
		status = http.StatusOK
		response = map[string]interface{}{"status": status, "bid_value": max, "bidder_id": ids[index]}
	}
	return status, response
}

//FindMax : function to find highest element in a slice
//Params : slice(list of price in float64)
//Author :Kavya
//Date : 20/10/2019
func (as *System) FindMax(bidderAmount []float64) (float64, int) {
	highestBidder := bidderAmount[0]
	var index int
	for i, value := range bidderAmount {
		if value > highestBidder {
			highestBidder = value
			index = i
		}
	}
	return highestBidder, index
}
