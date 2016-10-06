/**
 * Created by hyku on 16/10/1.
 */

"use strict";

require("./view/style/entry.less");

function routine($routeParams, MobileDataService, $scope, $location) {
    var entryData = {
        url: $location.url(),
        id: $routeParams.id
    };
    MobileDataService.pagePush("entry", entryData);
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
    $scope.back = function() {
        if ($scope.entry.layout === "single-page") {
            history.back();
            return false;
        }
        if ($scope.entry.layout === "multiple-pages") {
            if ($scope.currentStepIndex === 0) {
                history.back();
                return false;
            }
        }
        return false;
    };
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
