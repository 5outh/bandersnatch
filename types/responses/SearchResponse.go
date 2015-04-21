package responses

import "github.com/5outh/bandersnatch/types"

type SearchResponse struct {
    TermsOfUse string `xml:"termsofuse,attr"`
    BoardGame []types.SimpleBoardGame `xml:"boardgame"`
}
