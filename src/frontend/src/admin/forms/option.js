/**
 * Created by hyku on 2016/10/15.
 */

"use strict";

function routine($scope, AdminDataService) {
    $scope.addCase = function() {
        if (!$scope.item.report.cases) {
            $scope.item.report.cases = [];
        }
        $scope.item.report.cases.push({});
    };
    $scope.addDecree = function() {
        if (!$scope.item.report.decrees) {
            $scope.item.report.decrees = [];
        }
        $scope.item.report.decrees.push({});
    };
    if ($scope.ngDialogData.item) {
        $scope.item = $scope.ngDialogData.item;
    } else {
        $scope.item = {
            parent_id: $scope.ngDialogData.parent.id,
            report: {}
        };
    }
    $scope.submit = function() {
        var promise = null;
        if ($scope.ngDialogData.item) {
            promise = AdminDataService
                .optionUpdate($scope.item.id, $scope.item);
        } else {
            promise = AdminDataService
                .optionCreate($scope.item);
        }
        promise
            .then(function() {
                $scope.closeThisDialog({
                    status: "success",
                    data: $scope.item
                });
            })
            .catch(function(error) {
                console.error(error);
            })
        ;
    };
}

module.exports = routine;
