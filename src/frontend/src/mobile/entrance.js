/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require('!!ng-cache?prefix=mobile/view/!./view/home.html');
require('!!ng-cache?prefix=mobile/view/!./view/channel.html');
require('!!ng-cache?prefix=mobile/view/!./view/entry.html');

function register_controllers(app) {
    app.service("MobileDataService", require("./data-service"))
        .controller("mobile.home", require("./home"))
        .controller("mobile.channel", require("./channel"))
        .controller("mobile.entry", require("./entry"))
    ;
    app.config(function($routeProvider) {
        route($routeProvider);
    });
}

function route($routeProvider) {
    $routeProvider
        .when("/home", {
            controller: "mobile.home",
            templateUrl: "mobile/view/home.html"
        })
        .when("/channels/:id", {
            controller: "mobile.channel",
            templateUrl: "mobile/view/channel.html"
        })
        .when("/entries/:id", {
            controller: "mobile.entry",
            templateUrl: "mobile/view/entry.html"
        })
        .when("/", {
            redirectTo: "/home"
        })
    ;
}

module.exports = {
    register: register_controllers
};
