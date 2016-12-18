/**
 * Created by hyku on 2016/11/19.
 */

"use strict";

require("./v/style/frame.less");
require("./v/style/home.less");

function register(mod) {
    mod
        .service("AccountsService", require("./m/accounts"))
        .service("ResourceService", require("./m/resource"))
        .service("Session", function (AccountsService) {
            this.self = AccountsService.checkAuthentication;
        })
        .config(function ($stateProvider, $urlRouterProvider) {
            $stateProvider
                .state("home", {
                    abstract: true,
                    url: "/home",
                    views: {
                        "topBar@": {
                            template: require("./v/top_bar.pug"),
                            controller: require("./c/top_bar")
                        },
                        "main@": {
                            template: require("./v/home.pug"),
                            controller: require("./c/home")
                        },
                        "footer@": {
                            template: require("./v/footer.pug")
                        }
                    }
                })
                .state("home.base", {
                    url: "",
                    views: {
                        "suSongWenShu@home": {
                            template: require("./v/suSongWenShu.pug"),
                            controller: require("./c/suSongWenShu")
                        },
                        "xieYiFanBen@home": {
                            template: require("./v/xieYiFanBen.pug"),
                            controller: require("./c/xieYiFanBen")
                        },
                        "faLvWenDa@home": {
                            template: require("./v/faLvWenDa.pug"),
                            controller: require("./c/faLvWenDa")
                        }
                    }
                })
                .state("home.base.customer", {
                    url: "/customer"
                })
                .state("home.base.super", {
                    url: "/super"
                })
            ;
            $urlRouterProvider.when('/', '/home');
        })
        .run(function ($rootScope, $state, $stateParams) {
            $rootScope.$state = $state;
            $rootScope.$stateParams = $stateParams;
            $rootScope.$on("session.auth_failed", function () {
                console.log($state.current);
                console.log("session.auth_failed occurs");
                if ($state.is("home.base.super") || $state.is("home.base.customer")) {
                    console.log("auth_failed occurs");
                    $state.go("^", $stateParams, {reload: true, notify: false});
                }
            });
        })
    ;
}

module.exports = {
    register: register
};
