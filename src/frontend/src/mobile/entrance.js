/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require('!!ng-cache?prefix=mobile/view/!./view/home.html');

function register_controllers(app) {
    app
        .controller("mobile.home", require("./home"))
        ;
    app.config(function($routeProvider) {
        route($routeProvider);
    })
}

function route($routeProvider) {
    $routeProvider
        .when("/mobile/home", {
              controller: "mobile.home",
              templateUrl: "mobile/view/home.html"
        })
        .when("/mobile", {
            redirectTo: "/mobile/home"
        })
    ;
}

module.exports = {
    register: register_controllers
};
