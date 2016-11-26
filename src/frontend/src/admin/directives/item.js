/**
 * Created by hyku on 2016/11/15.
 */

"use strict";

var dateFormat = require("dateformat");

var app = angular.module("admin.custom.item", [])
        .filter("ms_2_time", function () {
            return function(input, format) {
                format = format || "yyyy-mm-dd HH:MM:ss Z";
                try {
                    var date = new Date(input);
                    return dateFormat(date, format);
                } catch (e) {
                    return "Invalid";
                }
            };
        })
        .filter("safehtml", function($sce) {
            return function(input) {
                if (input) {
                    return $sce.trustAsHtml(input);
                }
                return "";
            };
        })
        .directive("issue", function() {
        return {
            restrict: "A",
            scope: {
                item: "=issue"
            },
            template: require("./v/issue.html"),
            link: function($scope, ele, attr) {
            }
        };
    })
;

module.exports = app;
