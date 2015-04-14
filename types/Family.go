package types

type Family struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
