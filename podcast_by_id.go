package listennotes

import (
	"encoding/json"
	"net/http"
)

var (
	pocastByIDURL = "podcasts/"
)

func getPodcastByIDRequest(client *http.Client, token string, id string, options Params) (podcast Podcast, err error) {

	builtURL := buildURL(options)
	podcastByID := listenNotesBaseURL + pocastByIDURL + id + "?" + builtURL

	response, err := newGetRequest(podcastByID, token, client)

	if err != nil {
		return podcast, err
	}

	var podcastResp Podcast

	err = json.Unmarshal(response, &podcastResp)

	if err != nil {
		return podcast, err
	}

	return podcastResp, nil
}
