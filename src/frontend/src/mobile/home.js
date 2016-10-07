/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require("./view/style/home.less");

function controller($scope, MobileDataService, $location) {
    var homeData = {
        url: $location.url()
    };
    MobileDataService.pagePush("home", homeData);
    MobileDataService.loadHome()
        .then(function(layout) {
            $scope.layout = layout;
            console.log(layout);
        })
        .catch(function(err) {
            console.error(err);
        })
    ;
}

module.exports = controller;