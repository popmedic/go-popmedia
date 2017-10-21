package model

/*
Content ...
*/
type Content struct {
	//Required: The date the video was added to the library in the ISO 8601 format:
	//  {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}+{TZ}. E.g.: 2015-11-11T22:21:37+00:00
	// This information is used to generate the “Recently Added” category.
	DateAdded string `json:"dateAdded"`
	//Required: One or more video files. For non-adaptive streams, you can specify
	// the same video with different qualities so the Roku player can choose the
	// best one based on bandwidth.
	Videos Videos `json:"videos"`
	//Required: Runtime in seconds.
	Duration int `json:"duration"`
	//Required: One or more caption files. This is required except for short-form videos.
	// Supported formats are described in Closed Caption / Subtitle Support.
	Captions Captions `json:"captions"`
	//Optional: The trickplay file(s) that displays images as a user scrubs through a video,
	// in Roku’s BIF format. Trickplay files in multiple qualities can be provided.
	TrickPlayFiles TrickPlayFiles `json:"trickPlayFiles"`
	//Optional: The language in which the video was originally produced
	// (e.g., “en”, “en-US”, “es”, etc). ISO 639 alpha-2 or alpha-3 language code string.
	Language string `json:"language"`
	//Optional: The date when the content should become available in the ISO 8601 format:
	//  {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}+{TZ}. E.g.: 2015-11-11T22:21:37+00:00
	ValidityPeriodStart string `json:"validityPeriodStart"`
	//Optional: The date when the content is no longer available in the ISO 8601 format:
	//  {YYYY}-{MM}-{DD}T{hh}:{mm}:{ss}+{TZ}. E.g.: 2015-11-11T22:21:37+00:00
	ValidityPeriodEnd string `json:"validityPeriodEnd"`
	//Optional: One or more time codes. Represents a time duration from the beginning of
	// the video where an ad shows up. Conforms to the format: {hh}:{mm}:{ss}.
	AdBreaks string `json:"adBreaks"`
}
