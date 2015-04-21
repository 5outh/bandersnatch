package bandersnatch

import (
    "net/http"
    "net/url"
    "io/ioutil"
    "encoding/xml"
    "github.com/5outh/bandersnatch/types/responses"
    "strconv"
)

type SearchRequest struct {
    SearchTerm string
    Exact bool
}

func NewSearchRequest (searchTerm string, options ...func(*SearchRequest) error) (*SearchRequest, error) {
    searchRequest := &SearchRequest{}

    // Default options
    searchRequest.Exact = false
    searchRequest.SearchTerm = searchTerm

	for _, op := range options {
		err := op(searchRequest)
		if err != nil {
			return nil, err
		}
	}

	return searchRequest, nil
}

func OptionExact(exact bool) func (searchRequest *SearchRequest) error {
    return func (searchRequest *SearchRequest) error {
        searchRequest.Exact = true
        return nil
    }
}

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

func GetSearch(searchRequest *SearchRequest) (*responses.SearchResponse, error) {
    url := &url.URL{
        Scheme: "http",
        Host: "boardgamegeek.com",
        Path: "/xmlapi/search",
    }

    q := url.Query()

    q.Set("search", (*searchRequest).SearchTerm)

    if (*searchRequest).Exact {
        q.Set("exact", "1")
    }

    url.RawQuery = q.Encode()

    apiUrl := url.String()

    resp, err := http.Get(apiUrl)

    if (err != nil) {
        return nil, err
    }

    contents, _ := ioutil.ReadAll(resp.Body)

    var searchResponse responses.SearchResponse

    xml.Unmarshal(contents, &searchResponse)

    return &searchResponse, nil
}
