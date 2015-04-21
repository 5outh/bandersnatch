package types

type SimpleBoardGame struct {
    ObjectId int `xml:"objectid,attr"`
    YearPublished int `xml:"yearpublished"`
    Name Name `xml:"name"`
}
