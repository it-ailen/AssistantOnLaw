/**
 * Created by hyku on 2016/10/13.
 */

"use strict";

// var angular = require("angular");

require("!!ng-cache?prefix=directive/tree/!./view/tree.html");
require("!!ng-cache?prefix=directive/tree/!./view/node.html");
require("./view/tree.less");
var fileIcon = require("./view/imgs/file.png");
var folderIcon = require("./view/imgs/folder.png");
var closedFolderIcon = require("./view/imgs/folder-closed.png");

var tree = angular.module("tree", []);
tree
    .directive("treeNode", function () {
        return {
            scope: {
                item: "=",
                adapter: "=",
                icon: "=",
                folderOpen: "=",
                folderClose: "=",
                leafClick: "=",
                childrenLoader: "="
            },
            require: [],
            restrict: "E",
            templateUrl: "directive/tree/node.html",
            link: function($scope, element, attributes, controllers) {
                console.log("link tree node");
                $scope.open = false;
                $scope.node_click = function() {
                    if ($scope.item) {
                        var adaptedItem = $scope.adapter && $scope.adapter($scope.item);
                        if (adaptedItem.type === "branch") {
                            if ($scope.open) {
                                $scope.open = false;
                                $scope.folderClose && $scope.folderClose($scope.item);
                            }
                            else {
                                $scope.open = true;
                                $scope.folderOpen && $scope.folderOpen($scope.item);
                                console.log("load subitems now");
                                console.log($scope.childrenLoader);
                                $scope.subNodes = $scope.childrenLoader && $scope.childrenLoader($scope.item);
                                console.log($scope.subNodes);
                            }
                        }
                        else {
                            $scope.leafClick && $scope.leafClick($scope.item);
                        }
                    }
                    return false;
                };
                $scope.resolve_icon = function() {
                    var icon = null;
                    var adaptedItem = $scope.adapter && $scope.adapter($scope.item);
                    if (adaptedItem.type === 'branch') {
                        icon = ($scope.icon && $scope.icon($scope.item, $scope.open))
                            || (!$scope.open && closedFolderIcon)
                            || ($scope.open && folderIcon);
                    }
                    else {
                        icon = ($scope.icon && $scope.icon($scope.item))
                            || fileIcon;
                    }
                    return icon;
                };
                $scope.node_class = function() {
                    var classes = ["node"];
                    var adaptedItem = $scope.adapter && $scope.adapter($scope.item);
                    if (adaptedItem.type === 'branch') {
                        classes.push("branch");
                        if ($scope.open) {
                            classes.push("open");
                        }
                        else {
                            classes.push("closed");
                        }
                    }
                    else {
                        classes.push("leaf");
                    }
                    return classes;
                };
            }
        };
    })
    .directive("tree", function () {
        var link = function($scope, element, attributes, controllers) {
            console.log($scope.root);
            $scope.adapter = $scope.itemAdapter || function(item) {
                    console.log("in tree .adapter");
                    return item;
                };
            $scope.tree_class = function() {
                var classes = ["tree"];
                return classes;
            }
        };
        return {
            scope: {
                root: "=root",
                itemAdapter: "=adapter",
                icon: "=",
                folderOpen: "=",
                folderClose: "=",
                leafClick: "=",
                childrenLoader: "="
            },
            require: [],
            restrict: "E",
            templateUrl: "directive/tree/tree.html",
            link: link
        }
    })
;

module.exports = tree;
