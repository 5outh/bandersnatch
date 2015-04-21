package bandersnatch

import (
    "net/http"
    "io/ioutil"
    "encoding/xml"
    "github.com/5outh/bandersnatch/types"
    "strconv"
)

func GetBoardGame(boardGameId int) *types.BoardGames {
    resp, _ := http.Get("http://www.boardgamegeek.com/xmlapi/boardgame/" + strconv.Itoa(boardGameId))

    contents, _ := ioutil.ReadAll(resp.Body)

    var boardGames types.BoardGames

	xml.Unmarshal(contents, &boardGames)

    return &boardGames
}
