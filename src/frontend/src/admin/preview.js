/**
 * Created by hyku on 2016/10/30.
 */
"use strict";

var m = angular.module("previewer", [])
    .directive("pagePreviewHome", function () {
        return {
            scope: {
                data: "=ngModel"
            },
            restrict: "A",
            controller: function($scope) {},
            template: require("../mobile/view/home.html"),
            link: function($scope, element, attr, controllers) {
            }
        }
    })
;

module.exports = m;
