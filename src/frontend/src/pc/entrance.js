/**
 * Created by hyku on 2016/11/19.
 */

"use strict";

require("./v/style/frame.less");
require("./v/style/home.less");

function register(mod) {
    mod
        .service("AccountsService", require("./m/accounts"))
        .service("ResourceService", require("./m/resource"))
        .service("Session", function (AccountsService) {
            this.self = AccountsService.checkAuthentication;
        })
        .config(function ($stateProvider, $urlRouterProvider) {
            $stateProvider
                .state("frame", {
                    abstract: true,
                    url: "",
                    views: {
                        "topBar@": {
                            template: require("./v/top_bar.pug"),
                            controller: require("./c/top_bar")
                        },
                        "footer@": {
                            template: require("./v/footer.pug")
                        }
                    }
                })
                .state("frame.home", {
                    abstract: true,
                    url: "/home",
                    views: {
                        "main@": {
                            template: require("./v/home.pug"),
                            controller: require("./c/home")
                        }
                    }
                })
                .state("frame.home.base", {
                    url: "",
                    views: {
                        "suSongWenShu@frame.home": {
                            template: require("./v/suSongWenShu.pug"),
                            controller: require("./c/suSongWenShu")
                        },
                        "xieYiFanBen@frame.home": {
                            template: require("./v/xieYiFanBen.pug"),
                            controller: require("./c/xieYiFanBen")
                        },
                        "faLvWenDa@frame.home": {
                            template: require("./v/faLvWenDa.pug"),
                            controller: require("./c/faLvWenDa")
                        },
                        "faLvZiXun@frame.home": {
                            template: require("./v/faLvZiXun.pug"),
                            controller: require("./c/faLvZiXun")
                        }
                    }
                })
                .state("frame.home.base.customer", {
                    url: "/customer"
                })
                .state("frame.home.base.super", {
                    url: "/super"
                })
            ;
            $urlRouterProvider.when('/', '/home');
        })
        .run(function ($rootScope, $state, $stateParams) {
            $rootScope.$state = $state;
            $rootScope.$stateParams = $stateParams;
            $rootScope.$on("session.auth_failed", function () {
                console.log($state.current);
                console.log("session.auth_failed occurs");
                if ($state.is("frame.home.base.super") || $state.is("frame.home.base.customer")) {
                    console.log("auth_failed occurs");
                    $state.go("^", $stateParams, {reload: true, notify: false});
                }
            });
        })
        .directive("questionInput", function () {
            return {
                restrict: "A",
                template: require("./v/question.pug"),
                replace: true,
                require: "ngModel",
                scope: {
                    question: "=",
                    choices: "=ngModel"
                },
                link: function ($scope, ele, attr, ngModelCtrl) {
                    $scope.randomId = Math.floor(Math.random() * 1000);
                    $scope.checkboxMap = {};
                    if (!angular.isDefined($scope.choices)) {
                        $scope.choices = [];
                    }
                    if ($scope.choices && $scope.choices.length > 0) {
                        if ($scope.question.type === 'multiple') {
                            angular.forEach($scope.choices, function (item) {
                                $scope.checkboxMap[item] = true;
                            });
                        } else {
                            $scope.checkboxMap.radio = $scope.choices[0];
                        }

                    }
                    $scope.$watch("checkboxMap", function (newValue) {
                        if (angular.isDefined(newValue) && angular.isDefined($scope.choices))
                            $scope.choices.splice(0, $scope.choices.length);
                        console.log("checkboxMap: ");
                        console.log(newValue);
                        if (angular.isDefined(newValue)) {
                            if ($scope.question.type === 'multiple') {
                                angular.forEach(newValue, function (value, key) {
                                    if (value) {
                                        $scope.choices.push(parseInt(key));
                                    }
                                });
                            } else {
                                if (angular.isDefined(newValue.radio)) {
                                    $scope.choices.push(newValue.radio);
                                }
                            }
                        }
                    }, true);
                }
            };
        })
        .filter("indexToChar", function () {
            return function (index) {
                return String.fromCharCode(65 + index);
            }
        })
        .filter("safeHTML", function ($sce) {
            return function (input) {
                if (input) {
                    return $sce.trustAsHtml(input);
                }
                return "";
            };
        })
    ;
}

module.exports = {
    register: register
};
