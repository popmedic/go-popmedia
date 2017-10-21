package model

/*
Root object of your feed. It contains basic information such as your company's name,
 when the feed was last updated, and other objects that will describe all your content
 such as TV Shows, Movies, etc.
*/
type Root struct {
	//Required: The name of the feed provider. E.g.: “Acme Productions”.
	ProviderName string `json:"providerName"`
	//Required: The date that the feed was last modified in the ISO 8601 format:
	//  {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}+{TZ}.
	//  E.g.: 2015-11-11T22:21:37+00:00
	LastUpdated string `json:"lastUpdated"`
	//Required: The language the channel uses for all its information and descriptions.
	//  (e.g., “en”, “en-US”, “es”, etc.). ISO 639 alpha-2 or alpha-3 language code string.
	Language string `json:"language"`
	//Required: A list of one or more movies.
	Movies Movies `json:"movies,omitempty"`
	//Required: A list of one or more series. Series are episodic in nature and would include:
	//  TV shows, daily/weekly shows, etc.
	Series Series `json:"series,omitempty"`
	//Required: A list of one or more short-form videos. Short-form videos are usually
	//  less than 20 minutes long and are not TV Shows or Movies.
	ShortFormVideos ShortFormVideos `json:"shortFormVideos,omitempty"`
	//Required: A list of one or more TV Specials. TV Specials are one-time TV programs
	//  that are not part of a series.
	TvSpecials TvSpecials `json:"tvSpecials,omitempty"`
	//Optional: An ordered list of one or more categories that will show up in your Roku Channel.
	//  Categories may also be manually specified within Direct Publisher if you do not want to
	//  provide them directly in the feed. Each time the feed is updated it will refresh the categories.
	Categories Categories `json:"catagories,omitempty"`
	//Optional: A list of one or more playlists. They are useful for creating manually ordered
	//  categories inside your channel.
	Playlists Playlists `json:"playlists,omitempty"`
}
