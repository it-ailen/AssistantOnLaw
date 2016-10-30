/**
 * Created by hyku on 2016/10/23.
 */

"use strict";

function routine($scope, $q, tools, AdminDataService) {
    if ($scope.ngDialogData.item) {
        $scope.item = $scope.ngDialogData.item;
    } else {
        $scope.item = {};
    }
    $scope.iconChange = function(e) {
        console.log(this);
        console.log(e);
    };
    $scope.uploadIcon = function(file) {
        return tools.uploadImage(file, "channel", {});
    };
    $scope.srcToUrl = function(src) {
        console.log("src: " + src);
        return src.uri;
    };
    $scope.submit = function() {
        $scope.formStatus = "submitting";
        var promise = null;
        if ($scope.ngDialogData.item) {
            promise = AdminDataService.channelUpdate($scope.item.id, $scope.item);
        } else {
            promise = AdminDataService.channelCreate($scope.item);
        }
        promise
            .then(function() {
                $scope.closeThisDialog({
                    status: "success",
                    data: $scope.item
                });
                $scope.formStatus = "success";
            })
            .catch(function(error) {
                console.error(error);
                $scope.formStatus = "failed";
            })
        ;
    };
}

module.exports = routine;
