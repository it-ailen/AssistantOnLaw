/**
 * Created by hyku on 16/10/1.
 */

"use strict";

require("./view/style/report.less");

function routine($routeParams, MobileDataService, $scope) {
    MobileDataService.loadReport($routeParams.id)
        .then(function(report) {
            $scope.report = report;
        })
        .catch(function(err) {
            console.error(err);
        })
    ;
}

module.exports = routine;
