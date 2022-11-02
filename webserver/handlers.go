package webserver

import (
	"context"
	"database/sql"
	stdlog "log"
	"net/http"
	"net/http/cgi"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func GetMux(ws *WebServer, ctx context.Context) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/blog", blogHandler(ws.dbConn, ctx))

	mux.Handle("/cgi-bin/",
		http.StripPrefix("/cgi-bin/",
			http.HandlerFunc(
				cgiHandler(ws.config.WebServer.CgiPath))))

	if !ws.config.UseCaddy {
		log.Info("Serving static files without Caddy")
		mux.Handle("/static/",
			http.StripPrefix("/static/",
				http.FileServer(
					http.Dir(ws.config.WebServer.StaticPath))))
	}

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := loadedTemplates.Render(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func cgiHandler(cgiPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(cgiPath, r.URL.Path)
		l := stdlog.New(log.StandardLogger().WriterLevel(log.ErrorLevel), "", 0)
		h := &cgi.Handler{Path: path, Root: "/cgi-bin/", Logger: l}

		log.Trace("Running CGI script ", path)
		h.ServeHTTP(w, r)
	}
}

func blogHandler(db *sql.Conn, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if len(id) > 0 {
			iid, err := strconv.Atoi(id)
			post, err := GetPost(db, ctx, int64(iid))
			if err != nil {
				http.Error(w, err.Error(), 500)
				log.Error(err)
			}

			err = loadedTemplates.Render(w, "blog-post.html", post)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}
			return
		}

		posts, err := ListPosts(db, ctx)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Error(err)
		}

		err = loadedTemplates.Render(w, "blog-list.html", posts)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
