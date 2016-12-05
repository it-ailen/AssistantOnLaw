/**
 * Created by AilenZou on 2016/11/15.
 */

"use strict";

function controller($scope) {
    $scope.fileNameChanged = function (ele) {
        console.log(ele);
        $scope.file = ele.files[0];
        $scope.submit();
    };
    $scope.cancel = function () {
        $scope.closeThisDialog({
            action: "cancel"
        });
    };
    $scope.submit = function () {
        $scope.closeThisDialog({
            action: "ok",
            data: $scope.file
        });
    };
}

module.exports = controller;
