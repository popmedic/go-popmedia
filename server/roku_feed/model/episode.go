package model

/*
Episodes an array of Episode objects
*/
type Episodes []Episode

/*
Episode object represents a single episode in a series or a season.
*/
type Episode struct {
	//Required	Your immutable string reference ID for the series. THIS CANNOT CHANGE.
	//  This should serve as a unique identifier for the movie across different locales.
	ID string `json:"id"`
	//Required	The title of the series. We use this field for matching in Roku Search.
	Title string `json:"title"`
	// Required: The actual video content, such as the URL of the video file, subtitles, etc.
	Content Content `json:"content"`
	//Required	The URL of the thumbnail for the series. This is used within your channel
	//  and in search results. Image dimensions must be at least 800x450
	//  (width x height, 16x9 aspect ratio).
	Thumbnail string `json:"thumbnail"`
	//Required The sequential episode number. E.g.: 3.
	EpisodeNumber int `json:"episodeNumber"`
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
	//Optional	One or more credits. The cast and crew of the series.
	Credits Credits `json:"credits"`
	// Optional: A parental rating for the content.
	Rating Rating `json:"rating"`
	//Optional	One or more third-party metadata provider IDs.
	ExternalIDs ExternalIDs `json:"externalIds"`
}
