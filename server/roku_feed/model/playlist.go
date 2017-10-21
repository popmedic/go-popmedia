package model

/*
Playlists an array of Playlist objects
*/
type Playlists []Playlist

/*
Playlist ...
*/
type Playlist struct {
	//Required: The name of the playlist. The name is limited to 20 characters.
	Name string `json:"name"`
	//Required: An ordered list of one or more item IDs. An item ID is the ID of
	// a movie/series/short-form video/TV special.
	ItemIDs []string `json:"itemIds"`
}
