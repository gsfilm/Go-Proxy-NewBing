package api

import (
	"adams549659584/go-proxy-newbing/common"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, common.PROXY_WEB_PAGE_PATH, http.StatusFound)
	} else {
		common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)
	}
}
