/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

function func($scope, ResourceService, ngDialog) {
    $scope.current = {};
    $scope.dirContextMenu = [
        ["添加文件", function ($itemScope, $event, modelValue, text, $li) {
            ngDialog.open({
                template: require("./v/file.pug"),
                plain: true,
                controller: require("./forms/file"),
                data: {
                    parent: $itemScope.item.properties.id
                },
                closeByDocument: false
            })
                .closePromise
                .then(function (data) {
                    console.log(data);
                    return data.value;
                })
                .then(function (data) {
                    console.log(data);
                    reload();
                })
            ;
        }],
        ["编辑", function ($itemScope, $event, modelValue, text, $li) {
            console.log($itemScope)
            var promise = null;
            if ($itemScope.item.properties.type === 'directory') {
                promise = ngDialog.open({
                    template: require("./v/directory.pug"),
                    plain: true,
                    controller: require("./forms/directory"),
                    data: {
                        item: $itemScope.item
                    },
                    closeByDocument: false
                }).closePromise;
            } else {
                promise = ngDialog.open({
                    template: require("./v/file.pug"),
                    plain: true,
                    controller: require("./forms/file"),
                    data: {
                        item: $itemScope.item
                    },
                    closeByDocument: false
                }).closePromise;
            }
            promise
                .then(function (data) {
                    console.log(data);
                    return data.value;
                })
                .then(function (data) {
                    if (data.status === 'ok') {
                        $itemScope.item.properties = data.data;
                    }
                })
            ;
        }],
        ["删除", function ($itemScope, $event, modelValue, text, $li) {
            console.log(arguments);
        }]
    ];
    $scope.fileContextMenu = [
        ["编辑", function ($itemScope, $event, modelValue, text, $li) {
            console.log($itemScope);
            var promise = null;
            if ($itemScope.question.properties.type === 'directory') {
                promise = ngDialog.open({
                    template: require("./v/directory.pug"),
                    plain: true,
                    controller: require("./forms/directory"),
                    data: {
                        item: $itemScope.question
                    },
                    closeByDocument: false
                }).closePromise;
            } else {
                promise = ngDialog.open({
                    template: require("./v/file.pug"),
                    plain: true,
                    controller: require("./forms/file"),
                    data: {
                        item: $itemScope.question
                    },
                    closeByDocument: false
                }).closePromise;
            }
            promise
                .then(function (data) {
                    console.log(data);
                    return data.value;
                })
                .then(function (data) {
                    if (data.status === 'ok') {
                        $itemScope.question.properties = data.data;
                    }
                })
            ;
        }],
        ["删除", function ($itemScope, $event, modelValue, text, $li) {
            console.log(arguments);
        }]
    ];
    $scope.openArticle = function(question) {
        console.log(question);
        $scope.current.question = question;
        console.log($scope.current);
    };
    $scope.data = {};
    function reload() {
        ResourceService.loadDirectoryTree("fa_lv_wen_da")
            .then(function (tree) {
                $scope.data.tree = tree;
            })
            .catch(function (error) {
                console.error(error);
            })
        ;
    }
    reload();
}

module.exports = func;
