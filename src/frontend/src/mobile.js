/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("angular");
require("angular-route");
require("angular-ui-bootstrap");
require("angular-touch");
require("angular-carousel");
require("ng-dialog");

require("ng-dialog/css/ngDialog.min.css");
require("ng-dialog/css/ngDialog-theme-default.min.css");

require("bootstrap/dist/css/bootstrap.css");

require("./common/common.less");
require("./mobile/directive/etc");


var app = angular.module("LA.mobile", [
    "ngRoute",
    "ui.bootstrap",
    "angular-carousel",
    "ngDialog",
    "etc"
]);

require("./mobile/entrance").register(app);

module.exports = app;
