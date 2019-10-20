package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//AuctionSystem : struct for auction collection
type AuctionSystem struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	AuctionID   string             `json:"auction_id" bson:"auction_id,omitempty"`
	AuctionName string             `json:"auction_name" bson:"auction_name,omitempty"`
	BidAmount   float64            `json:"bid_amount" bson:"bid_amount,omitempty"`
	AdRequest   string             `json:"ad_request" bson:"ad_request,omitempty"`
	AdResponse  string             `json:"ad_response" bson:"ad_response,omitempty"`
}
