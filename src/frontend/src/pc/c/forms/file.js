/**
 * Created by hyku on 2016/12/4.
 */

"use strict";

function func($scope, ResourceService, tools) {
    $scope.item = {
        name: $scope.ngDialogData.item && $scope.ngDialogData.item.properties.name,
        reference_uri: $scope.ngDialogData.item && $scope.ngDialogData.item.properties.reference,
        parent: $scope.ngDialogData.parent,
        type: "file"
    };
    $scope.upload = function(file) {
        return tools.uploadImage(file, "static");
    };
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
