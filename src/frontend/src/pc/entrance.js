/**
 * Created by hyku on 2016/11/19.
 */

"use strict";

function register(mod) {
    mod
        .config(function ($stateProvider, $urlRouterProvider) {
            $stateProvider
                .state("home", {
                    abstract: true,
                    url: "/home",
                    template: require("./v/home.html"),
                    controller: require("./c/home")
                })
                .state("home.guest", {
                    url: "/guest",
                    views: {
                        topBar: {
                            template: require("./v/top_bar.html"),
                            controller: require("./c/guest/top_bar")
                        },
                        suSongWenShu: {
                            template: require("./v/su_song_wen_shu.html"),
                            controller: require("./c/su_song_wen_shu")
                        },
                        xieYiFanBen: {
                            template: require("./v/xie_yi_fan_ben.html"),
                            controller: require("./c/xie_yi_fan_ben")
                        },
                        faLvWenDa: {
                            template: require("./v/su_song_wen_shu.html"),
                            controller: require("./c/su_song_wen_shu")
                        },
                        faLvZiXun: {
                            template: require("./v/su_song_wen_shu.html"),
                            controller: require("./c/su_song_wen_shu")
                        },
                        renGongZiXun: {
                            template: require("./v/su_song_wen_shu.html"),
                            controller: require("./c/su_song_wen_shu")
                        }
                    }
                })
            ;
            $urlRouterProvider.when('/', '/home/guest');
        })
    ;
}

module.exports = {
    register: register
};
