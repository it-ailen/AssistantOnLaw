/**
 * Created by AilenZou on 2016/11/15.
 */


"use strict";

require("./v/file-input.less");

angular.module("file-input", ["ngDialog"])
    .directive("fileInput", function($parse, $q, ngDialog) {
        return {
            restrict: "A",
            require: "ngModel",
            link: function($scope, element, attr, ngModel) {
                attr.$set("readonly", true);
                var parser = null;
                attr.$observe("urlParser", function(value) {
                    if (value) {
                        parser = $parse(value);
                    }
                });
                var previewer = null;
                element.on("mouseover", function(e) {
                    if (ngModel.$modelValue
                        && angular.isString(ngModel.$modelValue)) {
                        var value = ngModel.$modelValue;
                        if (value.endsWith(".jpg") || value.endsWith(".jpeg")
                            || value.endsWith(".png") || value.endsWith(".gif")) {
                            if (parser) {
                                $q.when(parser($scope, {$value: value}))
                                    .then(function(url) {
                                        console.log("url: " + url);
                                        previewer = angular.element('<img id="previewer" src="' + url + '">');
                                        element.parent().append(previewer);
                                    })
                                ;
                            }
                        }
                    }

                });
                element.on("mouseout", function(e) {
                    if (previewer) {
                        previewer.remove();
                        previewer = null;
                    }
                });
                element.on("click", function () {
                    ngDialog.open({
                        template: require("./v/file-form.html"),
                        plain: true,
                        controller: require("./v/file-form"),
                        closeByDocument: false,
                        showClose: false
                    })
                        .closePromise
                        .then(function(res) {
                            return res.value;
                        })
                        .then(function(value) {
                            console.log(value);
                            if (value.action === 'ok') {
                                var fn = $parse(attr.onUpload);
                                $q.when(fn($scope, {file: value.data}))
                                    .then(function(id) {
                                        ngModel.$setViewValue(id);
                                        ngModel.$render();
                                    })
                                    .catch(function(error) {
                                        console.error(error);
                                        ngModel.$setValidity("upload", false);
                                    })
                                ;
                            }
                        })
                    ;
                });
                ngModel.$viewChangeListeners.push(function() {
                    ngModel.$setValidity("upload", true);
                });
            }
        }
    })
;

module.exports = "file-input";
