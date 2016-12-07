/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

function func($scope, ResourceService, ngDialog, toastr, tools) {
    function updateClass (c) {
        var promise = ngDialog.open({
            template: require("./v/directory.pug"),
            plain: true,
            controller: function ($scope, ResourceService) {
                $scope.item = {
                    name: $scope.ngDialogData.item && $scope.ngDialogData.item.name || undefined
                };
                $scope.submit = function () {
                    var promise = ($scope.ngDialogData.item) ?
                        ResourceService.updateFaLvWenDaClass($scope.ngDialogData.item.id, $scope.item) :
                        ResourceService.createFaLvWenDaClass($scope.item);
                    promise
                        .then(function (c) {
                            $scope.closeThisDialog({
                                status: "ok",
                                data: c
                            });
                        })
                        .catch(function (error) {
                            console.log(error);
                            toastr.error(error.message, "Error");
                        })
                    ;
                };
            },
            data: {
                item: c
            },
            closeByDocument: false
        }).closePromise;
        promise
            .then(function(res) {
                console.log(res);
                if (res && res.value && res.value.status === 'ok') {
                    var newClass = res.value.data;
                    if (c) {
                        c.name = newClass.name;
                    } else {
                        if (!$scope.data.classes) {
                            $scope.data.classes = [];
                        }
                        $scope.data.classes.push(c);
                    }
                }
            })
        ;
    }
    function updateArticle (article, c) {
        var promise = ngDialog.open({
            template: require("./v/directory.pug"),
            plain: true,
            controller: function ($scope, ResourceService) {
                $scope.item = {
                    name: $scope.ngDialogData.item && $scope.ngDialogData.item.name || undefined,
                    class_id: $scope.ngDialogData.parent && $scope.ngDialogData.parent.id
                };
                if (!$scope.ngDialogData.item) {
                    $scope.item.content = "";
                }
                $scope.submit = function () {
                    var promise = ($scope.ngDialogData.item) ?
                        ResourceService.updateFaLvWenDaArticle($scope.ngDialogData.item.id, $scope.item) :
                        ResourceService.createFaLvWenDaArticle($scope.item);
                    promise
                        .then(function (a) {
                            $scope.closeThisDialog({
                                status: "ok",
                                data: a
                            });
                        })
                        .catch(function (error) {
                            console.log(error);
                            toastr.error(error.message, "Error");
                        })
                    ;
                };
            },
            data: {
                item: article,
                parent: c
            },
            closeByDocument: false
        }).closePromise;
        promise
            .then(function(res) {
                console.log(res);
                if (res && res.value && res.value.status === 'ok') {
                    var newArticle = res.value.data;
                    if (article) {
                        article.name = newArticle.name;
                    } else {
                        if (!c.articles) {
                            c.articles = [];
                        }
                        c.articles.push(newArticle);
                    }
                }
            })
        ;
    }
    $scope.summerNoteOptions = {
        height: 500
    };
    $scope.uploadImage = function(files) {
        console.log("files");
        console.log(files);
        for (var i = 0; i < files.length; i++) {
            var file = files[i];
            tools.uploadImage(file, "fa_lv_wen_da")
                .then(function(path) {
                    console.log($("#article-summernote"));
                    console.log($(".summernote").summernote);
                    $(".summernote")
                        .summernote('editor.insertImage', path, function($image) {
                            console.log($image);
                        })
                    ;
                })
        }
    };
    $scope.current = {};
    $scope.classContextMenu = [
        ["添加文件", function ($itemScope, $event, modelValue, text, $li) {
            console.log($itemScope.c);
            updateArticle(null, $itemScope.c);
        }],
        ["编辑", function ($itemScope, $event, modelValue, text, $li) {
            console.log($itemScope);
            updateClass($itemScope.c);
        }],
        ["删除", function ($itemScope, $event, modelValue, text, $li) {
            console.log(arguments);
        }]
    ];
    $scope.articleContextMenu = [
        ["编辑", function ($itemScope, $event, modelValue, text, $li) {
            console.log($itemScope);
            updateArticle($itemScope.article);
        }],
        ["删除", function ($itemScope, $event, modelValue, text, $li) {
            console.log(arguments);
        }]
    ];

    $scope.updateClass = updateClass;
    $scope.openArticle = function (article) {
        $scope.current.article = article;
    };
    $scope.data = {};
    function reload() {
        ResourceService.loadFaLvWenDaClasses()
            .then(function (classes) {
                console.log(classes);
                $scope.data.classes = classes;
            })
    }

    reload();
}

module.exports = func;
