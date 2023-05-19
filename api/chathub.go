package api

import (
	"adams549659584/go-proxy-newbing/common"
	"net/http"
)

func ChatHub(w http.ResponseWriter, r *http.Request) {
	common.NewSingleHostReverseProxy(common.BING_CHAT_URL).ServeHTTP(w, r)
}
