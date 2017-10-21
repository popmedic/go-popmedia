package model

/*
ExternalIDs an array of Catagory objects
*/
type ExternalIDs []ExternalID

/*
ExternalID object of property:

	movie
	series
	series -> episodes -> episode
	series -> seasons -> episodes -> episode
	shortFormVideo
	tvSpecial

This object represents a third-party metadata provider ID (such as TMS, Rovi, IMDB, EIDR),
 that can provide more information about a specific video content. This information is
 used to optimize your content to be discovered within the Roku Search, and provide more
 details to users.
*/
type ExternalID struct {
	//Required: The third-party metadata provider ID for your video content. For example,
	// in the case of IMDB you would use the last part of the URL of a movie such as
	//  "http://www.imdb.com/title/tt0371724".
	ID string `json:"id"`
	//Required: Must be one of the following:
	//  TMS – A Tribune Metadata Service ID for the content
	//  ROVI - A Rovi ID for the content
	//  IMDB – An Internet Movie Database ID
	//  EIDR – An Entertainment Identifier Registry ID
	IDType string `json:"idType"`
}

var (
	//IDTypeTMS – A Tribune Metadata Service ID for the content
	IDTypeTMS = "TMS"
	//IDTypeROVI - A Rovi ID for the content
	IDTypeROVI = "ROVI"
	//IDTypeIMDB – An Internet Movie Database ID
	IDTypeIMDB = "IMDB"
	//IDTypeEIDR – An Entertainment Identifier Registry ID
	IDTypeEIDR = "EIDR"
)
