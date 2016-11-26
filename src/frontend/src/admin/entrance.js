/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require("./view/style/main.less");

function register_controllers(app) {
    app.service("AdminDataService", require("./data-service"))
    ;
    app.config(function($routeProvider) {
        route($routeProvider);
    });
}

function route($routeProvider) {
    $routeProvider
        .when("/home", {
            controller: require("./home"),
            template: require("./view/home.html")
        })
        .when("/issues/todo", {
            controller: require("./c/issue/todo"),
            template: require("./view/issue/todo.html")
        })
        .when("/issues", {
            redirectTo: "/issues/todo"
        })
        .when("/", {
            redirectTo: "/home"
        })
    ;
}

module.exports = {
    register: register_controllers
};
