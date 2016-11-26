/**
 * Created by hyku on 2016/11/7.
 */

"use strict";

require("../v/style/home.less");

function routine($scope, MobileDataService, self) {
    $scope.layout = {};
    console.log("home controller...");
    self.walk_back();
    MobileDataService
        .loadLayoutHome()
        .then(function(home) {
            $scope.layout = home;
        })
    ;
}

module.exports = routine;
