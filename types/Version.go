package types

type Version struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
