/**
 * Created by hyku on 2016/10/23.
 */

"use strict";

function routine($scope, $q, tools, AdminDataService) {
    if ($scope.ngDialogData.item) {
        $scope.item = $scope.ngDialogData.item;
    } else {
        $scope.item = {
            type: "entry",
            channel_id: $scope.ngDialogData.channel.id
        };
    }
    $scope.submit = function() {
        $scope.formStatus = "submitting";
        var promise = null;
        if ($scope.ngDialogData.item) {
            if ($scope.item.type==='entry') {
                promise = AdminDataService.entryUpdate($scope.item.id, $scope.item);
            }
        } else {
            promise = AdminDataService.entryCreate($scope.item);
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
