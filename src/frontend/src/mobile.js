/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("angular");
require("angular-route");
require("angular-ui-bootstrap");

var app = angular.module("LA.mobile", [
    "ngRoute",
    "ui.bootstrap"
]);

var mobile = require("./mobile/entrance");
mobile.register(app);

app
    .config(function($routeProvider) {
        $routeProvider
            .when("/", {
                redirectTo: "/mobile"
            })
        ;
    })
;

module.exports = app;
