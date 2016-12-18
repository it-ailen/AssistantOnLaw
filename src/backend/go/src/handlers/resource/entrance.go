package resource

import "github.com/gorilla/mux"

func Register(router *mux.Router) error {
	resourceRouter := router.PathPrefix("/resources").Subrouter()

	fileHandler := NewFileHandler()
	router.Handle("/resources", fileHandler)
	resourceRouter.Handle("/{id:[^/]*}", fileHandler)

	resourceRouter.Handle("/tree/{id:[^/]*}", NewFileTreeHandler())

	faLvWenDaRouter := resourceRouter.PathPrefix("/fa_lv_wen_da").Subrouter()
	classHandler := NewFaLvWenDaClassHandler()
	faLvWenDaRouter.Handle("/classes", classHandler)
	faLvWenDaRouter.Handle("/classes/{id:[^/]+}", classHandler)

	articleHandler := NewFaLvWenDaArticleHandler()
	faLvWenDaRouter.Handle("/articles", articleHandler)
	faLvWenDaRouter.Handle("/articles/{id:[^/]+}", articleHandler)

	suSongWenShuRouter := resourceRouter.PathPrefix("/su_song_wen_shu").Subrouter()

	wenShuFileHandler := NewSuSongFileHandler()
	suSongWenShuRouter.Handle("/files", wenShuFileHandler)
	suSongWenShuRouter.Handle("/files/{id:[^/]+}", wenShuFileHandler)

	minShiSuSongHandler := NewMinShiSuSongHandler()
	suSongWenShuRouter.Handle("/min_shi_su_song/{id:[^/]+}", minShiSuSongHandler)


	return nil
}
