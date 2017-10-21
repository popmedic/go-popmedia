package model

/*
Categories an array of Catagory objects
*/
type Categories []Category

/*
Category object of root property categories.

The category object defines a new category your channel will display, and the content
included in it based either on a playlist (see object description below), or a query containing
one or multiple tags. You can also create them directly in Direct Publisher.

There are three default categories in every channel: "Continue Watching", "Most Popular",
and "Recently Added".

Each category is displayed as a separate row to end-users.
*/
type Category struct {
	//Required: The name of the category that will show up in the channel.
	Name string `json:"name"`
	//Required: The name of the playlist in this feed that contains the content for this category.
	PlaylistName string `json:"playlistName"`
	//Required: The query that will specify the content for this category.
	// It is a Boolean expression containing tags that you have provided in your content feed.
	// The available operators are:
	//   AND
	//   OR
	// You cannot use both of them in the same query. You can use more than one. For example, if your feed has the tags "romance", "movie", "korean" and "dramas", you could do:
	//   movie AND korean
	//   movie AND korean AND dramas
	//   romance OR dramas
	// The following is NOT supported:
	//   movie AND romance OR dramas
	Query string `json:"query"`
	//Required	The order of the category. Must be one of the following:
	//  manual – For playlists only
	//  most_recent – reverse chronological order
	//  chronological – the order in which the content was published (e.g., Episode 1, Episode 2, etc.)
	//  most_popular – sort by popularity (based on Roku usage data).
	Order string `json:"order"`
}

// OrderByManual for Order string constant
var OrderByManual = "manual"

// OrderByMostRecent for Order string constant
var OrderByMostRecent = "most_recent"

// OrderByChronological for Order string constant
var OrderByChronological = "chronological"

// OrderByMostPopular for Order string constant
var OrderByMostPopular = "most_popular"
