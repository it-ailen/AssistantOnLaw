/**
 * Created by hyku on 16/9/29.
 */
"use strict";

require("angular");
require("angular-route");
require("angular-ui-bootstrap");
require("bootstrap/dist/css/bootstrap.css");

require("./libs/treeview");

require("./common/common.less");
require("./common/grid.less");


var app = angular.module("LA.admin", [
    "ngRoute",
    "ui.bootstrap",
    "tree"
]);

var admin = require("./admin/entrance");
admin.register(app);

module.exports = app;
