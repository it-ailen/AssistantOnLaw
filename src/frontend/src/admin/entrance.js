/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require("./view/style/main.less");
require('!!ng-cache?prefix=admin/view/!./view/home.html');

function register_controllers(app) {
    app.service("AdminDataService", require("./data-service"))
        .controller("admin.home", require("./home"))
    ;
    app.config(function($routeProvider) {
        route($routeProvider);
    });
}

function route($routeProvider) {
    $routeProvider
        .when("/home", {
            controller: "admin.home",
            templateUrl: "admin/view/home.html"
        })
        .when("/", {
            redirectTo: "/home"
        })
    ;
}

module.exports = {
    register: register_controllers
};
