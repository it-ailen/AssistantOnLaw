/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("angular");
require("angular-route");
require("angular-ui-bootstrap");
require("bootstrap/dist/css/bootstrap.css");
require("ng-dialog");

require("ng-dialog/css/ngDialog.min.css");
require("ng-dialog/css/ngDialog-theme-default.min.css");

// require("./libs/treeview");
require("angular-treeview");
require("angular-easy-input");

require("./common/common.less");
require("./common/grid.less");


var app = angular.module("LA.admin", [
    "ngRoute",
    "ui.bootstrap",
    "angular.tree",
    "ngDialog",
    "angular.easy.input"
]);

app
    .service("tools", require("./service/tools"))
;

var admin = require("./admin/entrance");
admin.register(app);

module.exports = app;
