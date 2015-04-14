package types

type Designer struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
