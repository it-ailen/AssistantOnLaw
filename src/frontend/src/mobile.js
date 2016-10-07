/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("angular");
require("angular-route");
require("angular-ui-bootstrap");
require("../node_modules/bootstrap/dist/css/bootstrap.css");

require("./common/common.less");


var app = angular.module("LA.mobile", [
    "ngRoute",
    "ui.bootstrap"
]);

var mobile = require("./mobile/entrance");
mobile.register(app);

// app
//     .config(function($routeProvider) {
//     })
// ;

module.exports = app;
