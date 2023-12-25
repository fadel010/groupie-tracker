package api

type Api struct {
	Locations string
	Dates string
	Relation string
	Artists string
}

 var ApiLinks = Api {
	Locations: "https://groupietrackers.herokuapp.com/api/locations",
	Relation: "https://groupietrackers.herokuapp.com/api/relation",
	Dates: "https://groupietrackers.herokuapp.com/api/dates",
	Artists: "https://groupietrackers.herokuapp.com/api/artists",
 }
