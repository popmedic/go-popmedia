package model

/*
TrickPlayFiles an array of TrickPlayFile objects
*/
type TrickPlayFiles []TrickPlayFile

/*
TrickPlayFile object of property content -> trickPlayFiles.

This object represents a single trickplay file.

Trickplay files are the images shown when a user scrubs through a video,
either fast-forwarding or rewinding.

The file must be in the Roku BIF format, as described in Trick Mode Support.
https://sdkdocs.roku.com/display/sdkdoc/Trick+Mode+Support
*/
type TrickPlayFile struct {
	//Required: The URL to the image representing the trickplay file Must be in the Roku BIF format,
	// more information in the Trick Mode Support article.
	//  https://sdkdocs.roku.com/display/sdkdoc/Trick+Mode+Support
	URL string `json:"url"`
	//Required	Must be one of the following:
	//  HD – 720p
	//  FHD – 1080p
	Quality string `json:"quality"`
}

var (
	//TrickPlayFileQualityHD - 720p
	TrickPlayFileQualityHD = "HD"
	//TrickPlayFileQualityFHD - 1080p
	TrickPlayFileQualityFHD = "FHD"
)
