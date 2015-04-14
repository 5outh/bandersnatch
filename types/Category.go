package types

type Category struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
