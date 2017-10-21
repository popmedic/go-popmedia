package model

/*
Rating object of property:

	movie
	series -> episodes -> episode
	shortFormVideo
	tvSpecial

  This object represents the rating for the video content.
  You can define the parental rating, as well as the source (USA Parental Rating,
  UK Content Provider, etc).
*/
type Rating struct {
	//Required: Must be a value listed in Parental Ratings
	Rating string `json:"rating"`
	//Required: Must be one of the following:
	// BBFC
	// CHVRS
	// CPR
	// MPAA
	// UK_CP
	// USA_PR
	// See Rating Sources for more information.
	//   https://github.com/rokudev/feed-specifications/blob/master/direct-publisher-feed-specification.md#rating-sources
	RatingSource string `json:"ratingSource"`
}

var (
	//RatingSourceBBFC - BBFC
	RatingSourceBBFC = "BBFC"
	//RatingSourceCHVRS - CHVRS
	RatingSourceCHVRS = "CHVRS"
	//RatingSourceCPR - CPR
	RatingSourceCPR = "CPR"
	//RatingSourceMPAA - MPAA
	RatingSourceMPAA = "MPAA"
	//RatingSourceUKCP - UK CP
	RatingSourceUKCP = "UK_CP"
	//RatingSourceUSAPR - USA PR
	RatingSourceUSAPR = "USA_PR"

	//Rating12 - 12
	Rating12 = "12"
	//Rating12A - 12A
	Rating12A = "12A"
	//Rating14UP - 14+
	Rating14UP = "14+"
	//Rating14A - 14A
	Rating14A = "14A"
	//Rating15 - 15
	Rating15 = "15"
	//Rating18 - 18
	Rating18 = "18"
	//Rating18UP - 18+
	Rating18UP = "18+"
	//Rating18A - 18A
	Rating18A = "18A"
	//RatingA - A
	RatingA = "A"
	//RatingAA - AA
	RatingAA = "AA"
	//RatingC - C
	RatingC = "C"
	//RatingC8 - C8
	RatingC8 = "C8"
	//RatingE - E
	RatingE = "E"
	//RatingG - G
	RatingG = "G"
	//RatingNC17 - NC17
	RatingNC17 = "NC17"
	//RatingPG - PG
	RatingPG = "PG"
	//RatingPG13 - PG13
	RatingPG13 = "PG13"
	//RatingR - R
	RatingR = "R"
	//RatingR18 - R18
	RatingR18 = "R18"
	//RatingTV14 - TV14
	RatingTV14 = "TV14"
	//RatingTVG - TVG
	RatingTVG = "TVG"
	//RatingTVMA - TVMA
	RatingTVMA = "TVMA"
	//RatingTVPG - TVPG
	RatingTVPG = "TVPG"
	//RatingTVY - TVY
	RatingTVY = "TVY"
	//RatingTVY14 - TVY14
	RatingTVY14 = "TVY14"
	//RatingTVY7 - TVY7
	RatingTVY7 = "TVY7"
	//RatingU - U
	RatingU = "U"
	//RatingUc - Uc
	RatingUc = "Uc"
	//RatingUNRATED - UNRATED
	RatingUNRATED = "UNRATED"
)
