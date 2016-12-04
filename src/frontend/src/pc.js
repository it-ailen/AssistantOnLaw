/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("bootstrap/dist/css/bootstrap.css");
require("angular-ui-router");

require("ng-dialog");
require("ng-dialog/css/ngDialog.min.css");
require("ng-dialog/css/ngDialog-theme-default.min.css");


require("angular-treeview");


var app = angular.module("LvDaJia.pc", [
    // "ngRoute",
    "ui.router",
    "angular.tree",
    "ngDialog",
    require("angular-messages"),
    require("angular-ui-bootstrap")
]);

app
    .run(function($rootScope, $state, $stateParams) {
        $rootScope.$state = $state;
        $rootScope.$stateParams = $stateParams;
    })
    .service("tools", require("./service/tools"))
    .provider("Configure", require("./service/configure"))
    .config(function($urlRouterProvider, ConfigureProvider) {
        $urlRouterProvider.otherwise('/');
    })
;

require("./pc/entrance").register(app);

module.exports = app;
