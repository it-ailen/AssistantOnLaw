/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

var minShiSuSong = require("../flow-graphics/minShiSuSong");

function func($scope, ResourceService, ngDialog, toastr, tools) {
    function initModel(detail) {
        $scope.model = {
            id: detail.id,
            name: detail.name,
            description: detail.description
        };
    }

    var status = {
        current: null,
        editing: false
    };
    $scope.status = status;
    minShiSuSong.$on("click", function (data) {
        console.log("clicked...");
        console.log(data);
        if (status.current && status.current.id === data.id) {
            return;
        }
        ResourceService.loadMinShiSuSongDetail(data.id)
            .then(function (detail) {
                status.current = detail;
                initModel(detail);
                $scope.model.id = data.id;
            })
            .catch(function (error) {
                console.log(error);
                toastr.error(error.status + " - " + error.data.error);
            })
        ;
    });
    minShiSuSong.draw("#svg-minShiSuSong");

    $scope.modifyFile = function (detail, file) {
        console.log("lalallalalal")
        console.log(detail);
        var promise = ngDialog.open({
            template: require("./v/suSongWenShuFile.pug"),
            plain: true,
            controller: function ($scope) {
                $scope.item = {
                    name: file && file.name || undefined,
                    uri: file && file.uri || undefined,
                    flow: detail.flow,
                    step_id: detail.id
                };
                console.log("Modify File???")
                console.log($scope.item);
                $scope.upload = function (file) {
                    return tools.uploadImage(file, "static");
                };
                $scope.submit = function () {
                    var p = (file) ? ResourceService.updateSuSongWenShuFile(file.id, $scope.item) :
                        ResourceService.addSuSongWenShuFile($scope.item);
                    p
                        .then(function (f) {
                            $scope.closeThisDialog({
                                success: true,
                                data: f
                            });
                        })
                        .catch(function (error) {
                            toastr.error(error.status + " - " + error.data.error);
                        })
                    ;
                }
            }
        }).closePromise;
        promise
            .then(function (data) {
                console.log(data);
                return data.value;
            })
            .then(function (value) {
                if (value.success) {
                    detail.files.push(value.data);
                }
            })
        ;
    };

    $scope.removeFile = function (index, file) {
        ResourceService.deleteSuSongWenShuFile(file.id)
            .then(function () {
                status.current.files.splice(index, 1);
            })
            .catch(function (error) {
                toastr.error(error.status + " - " + error.data.error);
            })
        ;
    };

    $scope.submit = function (model) {
        console.log(model);
        ResourceService.saveMinShiSuSongDetail(model.id, model)
            .then(function (detail) {
                status.current = detail;
                status.editing = false;
                initModel(detail);
            })
            .catch(function (error) {
                toastr.error(error.status + " - " + error.data.error);
            })
        ;
    };
}

module.exports = func;
