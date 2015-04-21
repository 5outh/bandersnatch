package bandersnatch

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/xml"
    "github.com/5outh/bandersnatch/types/responses"
    "strconv"
)

func GetBoardGame(boardGameId int) (*responses.BoardGameResponse, error) {
    resp, err := http.Get("http://www.boardgamegeek.com/xmlapi/boardgame/" + strconv.Itoa(boardGameId))

    if (err != nil) {
        return nil, err
    }

    contents, _ := ioutil.ReadAll(resp.Body)

    var boardGameResponse responses.BoardGameResponse

	xml.Unmarshal(contents, &boardGameResponse)

    return &boardGameResponse, nil
}

func GetSearch(searchTerm string) (*responses.SearchResponse, error) {
    resp, err := http.Get("http://www.boardgamegeek.com/xmlapi/search?search=" + url.QueryEscape(searchTerm))

    if (err != nil) {
        return nil, err
    }

    contents, _ := ioutil.ReadAll(resp.Body)

    var searchResponse responses.SearchResponse

    xml.Unmarshal(contents, &searchResponse)

    return &searchResponse, nil
}
