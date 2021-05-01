package model

type Artist struct {
	Id               int64
	Name             string
	Image            string
	Members          []string
	CreationDate     uint16
	ConcertLocations []string
	ConcertDates     []string
	Relation         map[string][]string
}
