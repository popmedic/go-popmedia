package model

/*
Videos an array of Video objects
*/
type Videos []Video

/*
Video object of property content -> videos.

This object represents the details of a single video file.
*/
type Video struct {
	//Required	The URL of the video itself. The video should be served from a CDN (Content Distribution Network). Supported formats are described in Audio and Video Support.
	URL string `json:"url"`
	//Required	Must be one of the following:
	//   HD – 720p
	//   FHD – 1080p
	//   UHD – 4K
	// If your stream uses an adaptive bitrate, set the quality to the highest available.
	Quality string `json:"quality"`
	//Required	Must be one of the following:
	//   HLS
	//   SMOOTH
	//   DASH
	//   MP4
	//   MOV
	//   M4V
	VideoType string `json:"videoType"`
	//Required only for non-ABR streams.	The bitrate in kbps. For non-adaptive streams,
	// this must be provided. It is not needed for an ABR (e.g., HLS) stream.
	Bitrate int `json:"bitrate"`
}

var (
	//VideoQualityHD - 720p
	VideoQualityHD = "HD"
	//VideoQualityFHD - 1080p
	VideoQualityFHD = "FHD"
	//VideoQualityUHD - 4k
	VideoQualityUHD = "UJD"
	//VideoTypeHLS hls video
	VideoTypeHLS = "HLS"
	//VideoTypeSMOOTH smooth video
	VideoTypeSMOOTH = "SMOOTH"
	//VideoTypeDASH dash video
	VideoTypeDASH = "DASH"
	//VideoTypeMP4 mp4 video
	VideoTypeMP4 = "MP4"
	//VideoTypeMOV MOV video
	VideoTypeMOV = "MOV"
	//VideoTypeM4V M4V video
	VideoTypeM4V = "M4V"
)
