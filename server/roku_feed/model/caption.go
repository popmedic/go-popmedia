package model

/*
Captions an array of Caption objects
*/
type Captions []Caption

/*
Caption object of property content -> captions.

This object represents a single video caption file of a video content.

The supported formats are described in Closed Caption / Subtitle Support.
https://sdkdocs.roku.com/display/sdkdoc/Closed+Caption+Support
*/
type Caption struct {
	//Required: The URL of the video caption file. Supported formats are described
	// in Closed Caption / Subtitle Support.
	// https://sdkdocs.roku.com/display/sdkdoc/Closed+Caption+Support
	URL string `json:"url"`
	//Required: A language code for the subtitle (e.g., “en”, “es-mx”, “fr”, etc).
	//ISO 639-2 or alpha-3 language code string.
	Language string `json:"language"`
	//Required: A string specifying the type of caption. Default is subtitle.
	// Must be one of the following:
	//  CLOSED_CAPTION
	//  SUBTITLE
	CaptionType string `json:"captionType"`
}

var (
	//CaptionTypeCLOSED closed caption type
	CaptionTypeCLOSED = "CLOSED_CAPTION"
	//CaptionTypeSUB subtitle caption type
	CaptionTypeSUB = "SUBTITLE"
	
	)