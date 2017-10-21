package model

/*
ShortFormVideos an array of ShortFormVideo objects
*/
type ShortFormVideos []ShortFormVideo

/*
ShortFormVideo ...
*/
type ShortFormVideo struct {
	// Required: Your immutable string reference ID for the movie. THIS CANNOT CHANGE. This should serve as a unique identifier for the movie across different locales.
	ID string `json:"id"`
	// Required: Movie title. We use this value for matching in Roku Search. Please use plain text and don’t include extra information like year, version label, etc.
	Title string `json:"title"`
	// Required: The actual video content, such as the URL of the video file, subtitles, etc.
	Content Content `json:"content"`
	// Required: The genre(s) of the movie. Must be one of the values listed in Genres.
	Genres []string `json:"genre"`
	// Required: The URL of the thumbnail for the movie. This is used within your channel and in search results. Image dimensions must be at least 800x450 (width x height, 16x9 aspect ratio).
	Thumbnail string `json:"thumbnail"`
	// Required: The date the movie was initially released or first aired. Used to sort programs chronologically and grouping related content in Roku Search. Conforms to the ISO 8601 format: {YYYY}-{MM}-{DD}. E.g.: 2015-11-11
	ReleaseDate string `json:"releaseDate"`
	// Required: A movie description that does not exceed 200 characters. The text will be clipped if longer.
	ShortDescription string `json:"shortDescription"`
	// Optional: A longer movie description that does not exceed 500 characters. The text will be clipped if longer. Must be different from shortDescription.
	LongDescription string `json:"longDescription"`
	// Optional: One or more tags (e.g., “dramas”, “korean”, etc.). Each tag is a string and is limited to 20 characters. Tags are used to define what content will be shown within a category.
	Tags []string `json:"tags"`
	// Optional: One or more credits. The cast and crew of the movie.
	Credits Credits `json:"credits"`
	// Optional: A parental rating for the content.
	Rating Rating `json:"rating"`
}
