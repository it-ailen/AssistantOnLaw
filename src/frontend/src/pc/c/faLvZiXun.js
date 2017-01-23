/**
 * Created by hyku on 2016/12/3.
 */

"use strict";
require("./styles/faLvZiXun.less");

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
        $scope.data.entries = [];
        $scope.data.questions = [];
        $scope.selectionMap = {};
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
        $scope.data.questions = [];
        $scope.selectionMap = {};
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
            width: "80%",
            controller: function ($scope) {
                $scope.item = {
                    question: question && question.question || undefined,
                    type: question && question.type || undefined,
                    options: question && question.options || [],
                    entry_id: question && question.entry_id || entryId,
                    trigger_by: question && question.trigger_by && angular.copy(question.trigger_by)
                };
                $scope.status = {
                    availableQuestions: []
                };
                ResourceService.selectQuestions({
                    entry_id: $scope.item.entry_id
                })
                    .then(function (questions) {
                        if (question) {
                            questions.forEach(function (item, index) {
                                if (item.id === question.id) {
                                    questions.splice(index, 1);
                                }
                            });
                            if (question.trigger_by && question.trigger_by.question_id) {
                                $scope.triggerQuestionChange(question.trigger_by.question_id);
                            }
                        }
                        $scope.status.availableQuestions = questions;
                    })
                    .catch(function (error) {
                        toastr.error("Error on fetching questions");
                    })
                ;
                function getQuestion(id) {
                    var chosen = null;
                    $scope.status.availableQuestions.forEach(function (item, index) {
                        if (id === item.id) {
                            chosen = item;
                        }
                    });
                    return chosen;
                }

                $scope.triggerQuestionChange = function (questionId) {
                    $scope.status.triggerQuestion = getQuestion(questionId);
                    console.log($scope.status.triggerQuestion);
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
    $scope.deleteQuestion = function (question) {
        ResourceService.deleteQuestion(question.id)
            .then(function () {
                $scope.data.questions.forEach(function (item, index) {
                    if (item.id === question.id) {
                        $scope.data.questions.splice(index, 1);
                    }
                });
            })
            .catch(function (error) {
                toastr.error(error);
            })
        ;
    };
    $scope.selectionMap = {};
    $scope.questionClass = function (question) {
        var classes = [];
        if (question.trigger_by) {
            var chosenOptions = $scope.selectionMap[question.trigger_by.question_id];
            if (!chosenOptions || chosenOptions.length <= 0) {
                classes.push("to_be_triggered");
            } else {
                var triggered = true;
                question.trigger_by.options.forEach(function (item, index) {
                    if (chosenOptions.indexOf(item) < 0) {
                        triggered = false;
                    }
                });
                if (triggered) {
                    classes.push("triggered");
                } else {
                    classes.push("to_be_triggered");
                }
            }
        }
        return classes;
    };
    $scope.search = function (selections) {
        // console.log(selections)
        var data = [];
        angular.forEach(selections, function (value, key) {
            var item = {
                question_id: key,
                selections: value
            };
            data.push(item);
        });
        ResourceService.selectConclusions({
            selections: JSON.stringify(data)
        })
            .then(function (conclusions) {
                if (conclusions.length > 0) {
                    ngDialog.open({
                        template: require("./v/conclusion.pug"),
                        plain: true,
                        width: "80%",
                        data: {
                            conclusion: conclusions[0]
                        },
                        controller: function ($scope) {
                            $scope.item = $scope.ngDialogData.conclusion;
                        }
                    });
                }
            })
            .catch(function (error) {
                toastr.error(error);
            })
        ;
    };
    $scope.editConclusion = function (selections) {
        // console.log(selections)
        var data = [];
        angular.forEach(selections, function (value, key) {
            var item = {
                question_id: key,
                selections: value
            };
            data.push(item);
        });
        var promise = ngDialog.open({
            template: require("./v/forms/conclusion.pug"),
            plain: true,
            width: "80%",
            data: {
                data: data
            },
            controller: function ($scope) {
                var put = false;
                $scope.conclusion = null;
                $scope.selections = $scope.ngDialogData.data;
                ResourceService.selectConclusions({
                    selections: JSON.stringify($scope.selections)
                })
                    .then(function (conclusions) {
                        if (conclusions && conclusions.length > 0) {
                            put = true;
                            $scope.conclusion = conclusions[0];
                        } else {
                            $scope.conclusion = {
                                selections: $scope.ngDialogData.data
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
