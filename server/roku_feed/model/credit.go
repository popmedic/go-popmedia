package model

/*
Categories an array of Catagory objects
*/
type Credits []Credit

/*
Credit object of property:

	movie
	series
	series -> episodes -> episode
	shortFormVideo
	tvSpecial

  This object represents a single person in the credits of a video content.
*/
type Credit struct {
	//required: name of the person
	Name string `json:"name"`
	//required: role of the person - must be one of the following values:
	// actor
	// anchor
	// host
	// narrator
	// voice
	// director
	// producer
	// screenwriter
	Role string `json:"role"`
	//required: birthdate of the person
	BirthDate string `json:"birthDate"`
}

var (
	//CreditRoleActor - actor
	CreditRoleActor = "actor"
	//CreditRoleAnchor - anchor
	CreditRoleAnchor = "anchor"
	//CreditRoleHost - host
	CreditRoleHost = "host"
	//CreditRoleNarrator - narrator
	CreditRoleNarrator = "narrator"
	//CreditRoleVoice - voice
	CreditRoleVoice = "voice"
	//CreditRoleDirector - director
	CreditRoleDirector = "director"
	//CreditRoleProducer - producer
	CreditRoleProducer = "producer"
	//CreditRoleScreenwriter - screenwriter
	CreditRoleScreenwriter = "screenwriter"
)
