package main

import (
	"io/ioutil"
	"os"

	"github.com/TudorHulban/GoTemplating/cmd/config"
	"github.com/TudorHulban/GoTemplating/internal/blog/blogfile"
	"github.com/TudorHulban/GoTemplating/pkg/httpserve"
	"github.com/TudorHulban/log"
)

func main() {
	cfg, errCfg := config.NewConfiguration("", 3)
	if errCfg != nil {
		l := log.NewLogger(log.DEBUG, os.Stdout, true)
		l.Printf("configuration error %s", errCfg)
		os.Exit(1)
	}

	articles, errGet := getFiles(cfg.L, cfg.ArticlesRAWFolder, config.ExtensionArticleFile)
	if errGet != nil {
		cfg.L.Print(errGet)
		os.Exit(2)
	}

	if len(articles) == 0 {
		cfg.L.Print("could not find article files")
		os.Exit(3)
	}

	blog, errBlog := blogfile.NewBlogFromFiles(cfg, articles...)
	if errBlog != nil {
		cfg.L.Print(errBlog)
		os.Exit(4)
	}

	// render article pages
	errRender := blog.RenderArticles()
	if errRender != nil {
		cfg.L.Print(errRender)
		os.Exit(5)
	}

	// start HTTP server
	c := httpserve.Cfg{
		ListenPort:         cfg.ListeningPort,
		StaticAssetsFolder: cfg.ArticlesRenderToFolder,
	}

	http, errStart := httpserve.NewHTTPServer(c)
	if errStart != nil {
		cfg.L.Info(errStart)
		os.Exit(6)
	}

	http.Start()
}

func getFiles(l *log.Logger, fromFolder, withExtension string) ([]string, error) {
	// TODO: improve validation
	l.Debugf("searching folder %s for files with extension %s", fromFolder, withExtension)

	files, err := ioutil.ReadDir(fromFolder)
	if err != nil {
		return []string{}, err
	}

	result := []string{}

	for _, v := range files {
		if v.IsDir() == false {
			path := fromFolder + "/" + v.Name()
			l.Debugf("adding file %s", path)
			result = append(result, path)
		}
	}

	return result, nil
}
