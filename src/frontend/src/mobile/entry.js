/**
 * Created by hyku on 16/10/1.
 */

"use strict";

require("./view/style/entry.less");

function routine($routeParams, MobileDataService, $scope, $location, $window) {
    $scope.currentStep = null;
    MobileDataService.loadEntry($routeParams.id)
        .then(function(entry) {
            $scope.entry = entry;
            console.log($scope.entry);
            $scope.steps = [
                $scope.entry.step
            ];
            $scope.currentStep = $scope.entry.step;
        })
        .catch(function(err) {
            console.error(err);
        })
    ;
    $scope.stepStatus = function(step) {
        if ($scope.currentStep && step.id === $scope.currentStep.id) {
            return "current-step";
        }
        return "";
    };
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
