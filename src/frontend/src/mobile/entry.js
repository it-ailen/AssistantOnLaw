/**
 * Created by hyku on 16/10/1.
 */

"use strict";

require("./view/style/entry.less");

function routine($routeParams, MobileDataService, $scope) {
    MobileDataService.loadEntry($routeParams.id)
        .then(function(entry) {
            $scope.entry = entry;
            console.log($scope.entry);
        })
        .catch(function(err) {
            console.error(err);
        })
    ;
}

module.exports = routine;
