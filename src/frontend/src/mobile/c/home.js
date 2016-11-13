/**
 * Created by hyku on 2016/11/7.
 */

"use strict";

require("../v/style/home.less");

function routine($scope, MobileDataService, $rootScope, self) {
    console.log("home controller...");
    self.walk_back();
    $scope.layout = {};
}

module.exports = routine;
