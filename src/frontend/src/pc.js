/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("bootstrap/dist/css/bootstrap.css");
require("angular-ui-router");

// require("angular-ui-bootstrap");
require("ng-dialog");


// require("ng-dialog/css/ngDialog.min.css");
// require("ng-dialog/css/ngDialog-theme-default.min.css");

require("angular-treeview");


var app = angular.module("LvDaJia.pc", [
    // "ngRoute",
    "ui.router",
    // "ui.bootstrap",
    "angular.tree",
    "ngDialog"
]);

app
    .service("tools", require("./service/tools"))
    .config(function($urlRouterProvider) {
        $urlRouterProvider.otherwise('/');
    })
;

require("./pc/entrance").register(app);
// require("./pc/directives/pc-directives").register(app);

module.exports = app;
