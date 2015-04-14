package types

type Subdomain struct {
    ObjectId int `xml:"objectid,attr"`
    Name string `xml:",chardata"`
}
