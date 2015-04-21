package responses

import "github.com/5outh/bandersnatch/types"

type BoardGameResponse struct {
    TermsOfUse string `xml:"termsofuse,attr"`
    BoardGame []types.BoardGame `xml:"boardgame"`
}
