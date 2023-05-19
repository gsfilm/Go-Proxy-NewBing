package api

import (
	"adams549659584/go-proxy-newbing/common"
	"adams549659584/go-proxy-newbing/web"
	"net/http"
)

func WebStatic(w http.ResponseWriter, r *http.Request) {
	if _, ok := web.WEB_PATH_MAP[r.URL.Path]; ok || r.URL.Path == common.PROXY_WEB_PREFIX_PATH {
		http.StripPrefix(common.PROXY_WEB_PREFIX_PATH, http.FileServer(web.GetWebFS())).ServeHTTP(w, r)
	} else {
		common.NewSingleHostReverseProxy(common.BING_URL).ServeHTTP(w, r)
	}
}
