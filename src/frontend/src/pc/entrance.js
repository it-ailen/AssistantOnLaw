/**
 * Created by hyku on 2016/11/19.
 */

"use strict";

require("./v/style/frame.less");
require("./v/style/home.less");

function register(mod) {
    mod
        .service("AccountsService", require("./m/accounts"))
        .config(function ($stateProvider, $urlRouterProvider) {
            $stateProvider
                .state("home", {
                    abstract: true,
                    controller: require("./c/home"),
                    template: require("./v/home.pug")
                })
                .state("home.inst", {
                    url: "/home",
                    views: {
                        "topBar@": {
                            template: require("./v/top_bar.pug"),
                            controller: require("./c/top_bar")
                        },
                        "footer@": {
                            template: require("./v/footer.pug")
                        },
                        "xieYiFanBen": {
                            template: require("./v/xieYiFanBen.pug"),
                            controller: require("./c/xieYiFanBen")
                        }
                    }
                })
            ;
            $urlRouterProvider.when('/', '/home');
        })
    ;
}

module.exports = {
    register: register
};
