/**
 * Created by hyku on 2016/11/8.
 */

"use strict";

require("./style/self-consult.less");

function routine($scope, ngDialog, $log, MobileDataService, tools, $window, self) {
    $scope.problem = {
        paths: self.footprints(),
        attachments: []
    };
    $scope.add_attachment = function() {
        var dialog = ngDialog.open({
            template: require("./forms/attachment.html"),
            plain: true,
            controller: require("./forms/attachment"),
            closeByDocument: true,
            showClose: true,
            width: "80%"
        });
        dialog.closePromise
            .then(function(data) {
                $log.debug(data);
                return data.value;
            })
            .then(function(value) {
                console.log(value);
                if (value.status === "ok") {
                    $scope.problem.attachments.push(value.data.file);
                }
            })
            .catch(function(error) {
                $log.error(error);
            })
        ;
    };
    $scope.remove_attachment = function(index) {
        $scope.problem.attachments.splice(index, 1);
    };

    $scope.submit = function() {
        MobileDataService
            .post_consulting($scope.problem)
            .then(function() {
                console.log("您的问题已提交，我们会尽快给您答复");
                return tools.alert("您的问题已提交，我们会尽快给您答复");
            })
            .then(function() {
                $window.location.href = "#/home";
            })
            .catch(function(error) {
                tools.alert(error);
            })
    };
}

module.exports = routine;
