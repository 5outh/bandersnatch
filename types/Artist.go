package types

type Artist struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
