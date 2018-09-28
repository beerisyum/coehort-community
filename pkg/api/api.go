package api

import (
	"http"

	"github.com/erei/coemmunity/pkg/common"
)

type API struct {
	ClientID string
}

type TwitchStream struct {
	Channel struct {
		Status      string
		Game        string
		Name        string
		DisplayName string
		Logo        string
	} `json:"channel"`
	Viewers int
}

const URL = "https://api.twitch.tv/kraken/streams/"

func NewAPI(cid string) *API {
	api := API{
		ClientID: cid,
	}

	return &api
}

func (a *API) GetStreams(ids string) ([]TwitchStream, error) {
	if len(ids) > 100 {

	}
}

func streamsRequest(ids) {

}
