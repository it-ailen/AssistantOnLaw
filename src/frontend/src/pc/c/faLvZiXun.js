/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

function func($scope, ResourceService, ngDialog, toastr, tools, $sce) {
    $scope.data = {};
    function reload() {
        ResourceService.selectClasses()
            .then(function (classes) {
                $scope.data.classes = classes;
            })
            .catch(function (error) {
                toastr.error(error);
            })
        ;
    }

    $scope.loadEntries = function (classId) {
        ResourceService.selectEntries({
            "class_id": classId
        })
            .then(function (entries) {
                $scope.data.entries = entries;
            })
            .catch(function (error) {
                toastr.error(error);
            })
        ;
    };

    $scope.loadQuestions = function (entryId) {
        ResourceService.selectQuestions({
            "entry_id": entryId
        })
            .then(function (questions) {
                $scope.data.questions = questions;
            })
            .catch(function (error) {
                toastr.error(error);
            })
        ;
    };

    $scope.updateClass = function (cls) {
        var promise = ngDialog.open({
            template: require("./v/class.pug"),
            plain: true,
            controller: function ($scope) {
                $scope.item = {
                    name: cls && cls.name || undefined,
                    description: cls && cls.description || undefined,
                    logo: cls && cls.logo || undefined
                };
                $scope.upload = function (file) {
                    return tools.uploadImage(file, "static");
                };
                $scope.submit = function () {
                    var promise = null;
                    if (cls) {
                        promise = ResourceService.updateClass(cls.id, $scope.item);
                    } else {
                        promise = ResourceService.insertClass($scope.item);
                    }
                    promise
                        .then(function (newClass) {
                            $scope.closeThisDialog({
                                success: true,
                                data: newClass
                            });
                        })
                        .catch(function (error) {
                            toastr.error(error);
                        })
                    ;
                }
            }
        }).closePromise;
        promise
            .then(function (data) {
                return data.value;
            })
            .then(function (value) {
                if (value.success) {
                    var newClass = value.data;
                    if (cls) {
                        cls.name = newClass.name;
                        cls.description = newClass.description;
                        cls.logo = newClass.logo;
                    } else {
                        $scope.data.classes.push(newClass);
                    }
                }
            })
        ;
    };
    $scope.updateEntry = function (classId, entry) {
        console.log("updateEntry????")
        var promise = ngDialog.open({
            template: require("./v/entry.pug"),
            plain: true,
            controller: function ($scope) {
                $scope.item = {
                    name: entry && entry.name || undefined,
                    logo: entry && entry.logo || undefined,
                    class_id: classId,
                    layout_type: "single"
                };
                $scope.upload = function (file) {
                    return tools.uploadImage(file, "static");
                };
                $scope.submit = function () {
                    var promise = null;
                    if (entry) {
                        promise = ResourceService.updateEntry(entry.id, $scope.item);
                    } else {
                        promise = ResourceService.insertEntry($scope.item);
                    }
                    promise
                        .then(function (newEntry) {
                            $scope.closeThisDialog({
                                success: true,
                                data: newEntry
                            });
                        })
                        .catch(function (error) {
                            toastr.error(error);
                        })
                    ;
                }
            }
        }).closePromise;
        promise
            .then(function (data) {
                return data.value;
            })
            .then(function (value) {
                if (value.success) {
                    var newEntry = value.data;
                    if (entry) {
                        entry.name = newEntry.name;
                        entry.description = newEntry.description;
                        entry.logo = newEntry.logo;
                    } else {
                        $scope.data.entries.push(newEntry);
                    }
                }
            })
        ;
    };
    $scope.updateQuestion = function (entryId, question) {
        var promise = ngDialog.open({
            template: require("./v/question.pug"),
            plain: true,
            controller: function ($scope) {
                $scope.item = {
                    question: question && question.question || undefined,
                    type: question && question.type || undefined,
                    options: question && question.options || [],
                    entry_id: question && question.entry_id || entryId
                };
                $scope.upload = function (file) {
                    return tools.uploadImage(file, "static");
                };
                $scope.submit = function () {
                    var promise = null;
                    if (question) {
                        promise = ResourceService.updateQuestion(question.id, $scope.item);
                    } else {
                        promise = ResourceService.insertQuestion($scope.item);
                    }
                    promise
                        .then(function (item) {
                            $scope.closeThisDialog({
                                success: true,
                                data: item
                            });
                        })
                        .catch(function (error) {
                            toastr.error(error);
                        })
                    ;
                }
            }
        }).closePromise;
        promise
            .then(function (data) {
                return data.value;
            })
            .then(function (value) {
                if (value.success) {
                    var item = value.data;
                    if (question) {
                        question.question = item.question;
                        question.type = item.type;
                        question.options = item.options;
                    } else {
                        $scope.data.questions.push(item);
                    }
                }
            })
        ;
    };
    $scope.selectionMap = {};
    $scope.search = function (selections) {
        console.log(selections)
        var data = [];
        for (var k in selections) {
            var value = selections[k];
            console.log(k + " -- " + value);
            console.log(typeof value)
            var item = {
                question_id: k,
                selections: []
            };
            if (typeof value === 'object') {
                for (var option in value) {
                    item.selections.push(parseInt(option));
                }
            } else {
                item.selections.push(parseInt(value));
            }
            data.push(item);
        }
        var promise = ngDialog.open({
            template: require("./v/conclusion.pug"),
            plain: true,
            controller: function ($scope) {
                var put = false;
                $scope.conclusion = null;
                ResourceService.selectConclusions({
                    selections: JSON.stringify(data)
                })
                    .then(function (conclusions) {
                        if (conclusions && conclusions.length > 0) {
                            put = true;
                            $scope.conclusion = conclusions[0];
                        } else {
                            $scope.conclusion = {
                                selections: data
                            };
                        }
                    })
                    .catch(function (error) {
                        toastr.error(error);
                    })
                ;
                $scope.summerNoteOptions = {
                    height: 500
                };
                $scope.uploadImage = function (files) {
                    console.log("files");
                    console.log(files);
                    for (var i = 0; i < files.length; i++) {
                        var file = files[i];
                        tools.uploadImage(file, "conclusion")
                            .then(function (path) {
                                console.log($("#article-summernote"));
                                console.log($(".summernote").summernote);
                                $(".summernote")
                                    .summernote('editor.insertImage', path, function ($image) {
                                        console.log($image);
                                    })
                                ;
                            })
                    }
                };
                $scope.submit = function () {
                    var promise = null;
                    if (put) {
                        promise = ResourceService.updateConclusions($scope.conclusion.id, $scope.conclusion);
                    } else {
                        promise = ResourceService.insertConclusions($scope.conclusion);
                    }
                    promise
                        .then(function (item) {
                            $scope.closeThisDialog({
                                success: true,
                                data: item
                            });
                        })
                        .catch(function (error) {
                            toastr.error(error);
                        })
                    ;
                }
            }
        }).closePromise;

    };

    reload();
}

module.exports = func;
