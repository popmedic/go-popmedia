package model

/*
Series an array of Serie objects
*/
type Series []Serie

/*
Serie object represents a series, such as a season of a TV Show or a mini-series.
*/
type Serie struct {
	//Required	Your immutable string reference ID for the series. THIS CANNOT CHANGE.
	//  This should serve as a unique identifier for the movie across different locales.
	ID string `json:"id"`
	//Required	The title of the series. We use this field for matching in Roku Search.
	Title string `json:"title"`
	//Required*	One or more seasons of the series. Seasons should be used if episodes are grouped by seasons.
	Seasons Seasons `json:"seasons"`
	//Required*	One or more episodes of the series. Episodes should be used if they are not grouped by seasons
	//  (e.g., a mini-series).
	Episodes Episodes `json:"episodes"`
	//Required	The genre(s) of the series. Must be one of the values listed in Genres.
	Genres []string `json:"genres"`
	//Required	The URL of the thumbnail for the series. This is used within your channel
	//  and in search results. Image dimensions must be at least 800x450
	//  (width x height, 16x9 aspect ratio).
	Thumbnail string `json:"thumbnail"`
	//Required	The date the series first aired. Used to sort programs chronologically and
	//  grouping related content in Roku Search. Conforms to the ISO 8601 format:
	//   {YYYY}-{MM}-{DD}. E.g.: 2015-11-11
	ReleaseDate string `json:"releaseDate"`
	//Required	A description of the series that does not exceed 200 characters.
	//  The text will be clipped if longer.
	ShortDescription string `json:"shortDescription"`
	//Optional	A longer movie description that does not exceed 500 characters.
	//  The text will be clipped if longer. Must be different from shortDescription.
	LongDescription string `json:"longDescription"`
	//Optional	One or more tags (e.g., “dramas”, “korean”, etc.). Each tag is a string and is limited to 20 characters.
	// Tags are used to define what content will be shown within a category.
	Tags []string `json:"tags"`
	//Optional	One or more credits. The cast and crew of the series.
	Credits Credits `json:"credits"`
	//Optional	One or more third-party metadata provider IDs.
	ExternalIDs ExternalIDs `json:"externalIds"`
}
