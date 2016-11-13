/**
 * Created by hyku on 2016/11/13.
 */

"use strict";

require("./v/title-bar.less");

angular.module("etc", [])
    .directive("etcFileUploader", function($parse, $log, $q) {
        return {
            restrict: "A",
            require: "ngModel",
            link: function($scope, element, attr, ngModelCtrl) {
                attr.required = true;
                ngModelCtrl.$validators.required = function(modelValue, viewValue) {
                    return !attr.required || !ngModelCtrl.$isEmpty(viewValue);
                };
                attr.$observe("required", function() {
                    ngModelCtrl.$validate();
                });
                element.on("change", function(e) {
                    $log.debug("onUpload");
                    $log.debug(attr.onUpload);
                    var fn = $parse(attr.onUpload);
                    $log.debug(fn);
                    $log.debug(e);
                    $q.when(fn($scope, {file: e.target.files[0]}))
                        .then(function(id) {
                            $log.debug("setViewValue(" + id + ")");
                            ngModelCtrl.$setViewValue(id);
                        })
                        .catch(function(error) {
                            $log.error(error);
                        })
                    ;
                });
            }
        }
    })
    .directive("titleBar", function($window) {
        return {
            restict: "A",
            template: require("./v/title-bar.html"),
            transclude: true,
            link: function($scope, element, attr) {
                $scope.back = function() {
                    console.log("back now");
                    $window.history.back();
                }
            }
        }
    })
;


