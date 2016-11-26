/**
 * Created by hyku on 2016/11/13.
 */

"use strict";

function controller($scope, MobileDataService, $log) {
    $scope.attachment = {};
    $scope.upload = function(file) {
        return MobileDataService
            .upload_file(file);
    };
    $scope.submit = function() {
        $log.debug("attachment: ");
        $log.debug($scope.attachment);
        $scope.closeThisDialog({
            status: "ok",
            data: $scope.attachment
        });
    };
}

module.exports = controller;
