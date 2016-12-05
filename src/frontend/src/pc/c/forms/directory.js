/**
 * Created by hyku on 2016/12/4.
 */

"use strict";

function func($scope, ResourceService) {
    $scope.item = {
        name: $scope.ngDialogData.item && $scope.ngDialogData.item.properties.name,
        parent: $scope.ngDialogData.parent,
        type: "directory"
    };
    console.log($scope.ngDialogData);
    console.log($scope.item);
    $scope.submit = function () {
        if ($scope.ngDialogData.item) {
            ResourceService.updateFile($scope.ngDialogData.item.properties.id, $scope.item)
                .then(function (data) {
                    $scope.closeThisDialog({
                        status: "ok",
                        data: data
                    });
                })
            ;
        } else {
            ResourceService.createFile($scope.item)
                .then(function (data) {
                    $scope.closeThisDialog({
                        status: "ok",
                        data: data
                    });
                })
            ;
        }
    };
}

module.exports = func;
