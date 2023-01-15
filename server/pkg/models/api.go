package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Code    int         `json:"code"`
	Pagination 			`bson:"pagination" json:"pagination"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	Page     int `bson:"page" json:"page"` // Page number, 1-indexed 
	PerPage  int `bson:"perpage" json:"perpage"` // Number of items per page
	LastPage int`bson:"lastpage" json:"lastpage"` //Last page 
}

type NewPost struct {
	Body 			string 				`bson:"body" json:"body"`
	RedditData 		[]RedditData 		`bson:"reddit_data" json:"reddit_data"`		
	Date			time.Time 			`bson:"date" json:"date"`	
}

type Post struct {
	Id 				primitive.ObjectID 	`bson:"_id" json:"_id"`
	Body 			string 				`bson:"body" json:"body"`
	RedditData 		[]RedditData 		`bson:"reddit_data" json:"reddit_data"`		
	Date			time.Time 			`bson:"date" json:"date"`	
}

type RedditData struct {
	URL				string 	`bson:"url" json:"url"`	
	Title			string 	`bson:"title" json:"title"`	
	UserCount		int 	`bson:"subscribers" json:"subscribers"`	
	Description 	string 	`bson:"public_description" json:"public_description"`	
}

