/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("angular");
require("angular-ui-router");
require("angular-bootstrap-npm");
require("angular-ui-bootstrap");

require("angular-summernote/dist/angular-summernote");

require("ng-dialog");
require("ng-dialog/css/ngDialog.min.css");
require("ng-dialog/css/ngDialog-theme-default.min.css");

require("angular-bootstrap-contextmenu");


require("angular-treeview");

require("angular-toastr/dist/angular-toastr.css");


var app = angular.module("LvDaJia.pc", [
    // "ngRoute",
    "ui.router",
    "angular.tree",
    "ngDialog",
    "ui.bootstrap",
    require("angular-messages"),
    "ui.bootstrap.contextMenu",
    require("./directives/file-input"),
    require("angular-animate"),
    require("angular-toastr"),
    "summernote",
    require("./directives/svg")
    // "ui.bootstrap"
]);

app
    .run(function($rootScope, $state, $stateParams) {
        $rootScope.$state = $state;
        $rootScope.$stateParams = $stateParams;
    })
    .filter("range", function () {
        return function(emptyArray, count) {
            count = count || 0;
            for (var i = 0; i < count; i++) {
                emptyArray.push(i);
            }
            return emptyArray;
        }
    })
    .service("tools", require("./service/tools"))
    .provider("Configure", require("./service/configure"))
    .config(function($urlRouterProvider, ConfigureProvider) {
        $urlRouterProvider.otherwise('/');
    })
;

require("./pc/entrance").register(app);

module.exports = app;
