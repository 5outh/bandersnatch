package types

type BoardGames struct {
    TermsOfUse string `xml:"termsofuse,attr"`
    BoardGame []BoardGame `xml:"boardgame"`
}
