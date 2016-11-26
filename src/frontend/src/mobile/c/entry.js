/**
 * Created by hyku on 16/10/1.
 */

"use strict";

require("./style/entry.less");

function routine($routeParams, MobileDataService, $scope, tools) {
    $scope.status = {};
    var entryId = $routeParams.id;
    MobileDataService
        .loadLayoutEntry(entryId)
        .then(function(layout) {
            $scope.entry = layout;
            console.log(layout);
        })
        .catch(function(error) {
            tools.alert(error);
        })
    ;
    $scope.actionClick = function (action) {
        if (action.type === "step") {
            MobileDataService.loadStep(action.id)
                .then(function(step) {
                    $scope.steps.push(step);
                    $scope.currentStep = step;
                })
            ;
        } else if (action.type === "report") {
            location.href = "#/reports/" + action.id;
        }
        return true;
    };
}

module.exports = routine;
