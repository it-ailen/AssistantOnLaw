/**
 * Created by hyku on 16/9/30.
 */

"use strict";


function register_controllers(app) {
    app.service("MobileDataService", require("./m/data-service"))
        .service("self", require("./service/self"))
        .service("tools", require("./service/tools"))
        .config(function($routeProvider) {
            route($routeProvider);
        })
    ;
}

function route($routeProvider) {
    $routeProvider
        .when("/home", {
            controller: require("./c/home"),
            template: require("./v/home.html")
        })
        .when("/self-consult", {
            controller: require("./c/self-consult"),
            template: require("./v/self-consult.html")
        })
        .when("/channels/:id", {
            controller: require("./c/channel"),
            template: require("./v/channel.html")
        })
        .when("/entries/:id", {
            controller: require("./c/entry"),
            template: require("./v/entry.html")
        })
        .when("/", {
            redirectTo: "/home"
        })
    ;
}

module.exports = {
    register: register_controllers
};
