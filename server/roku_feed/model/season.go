package model

/*
Seasons an array of Season objects
*/
type Seasons []Season

/*
Season object represents a single season of a series.
*/
type Season struct {
	//Required: Sequential season number. E.g.: 3 or 2015.
	SeasonNumber int `json:"seasonNumber"`
	//Required: One or more episodes of this particular season.
	Episodes Episodes `json:"episodes"`
	//Optional: The URL of the thumbnail for the season.
	//  Image dimensions must be at least 800x450 (width x height, 16x9 aspect ratio).
	Thumbnail string `json:"thumbnail"`
}
