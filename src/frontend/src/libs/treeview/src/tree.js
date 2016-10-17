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
var plusIcon = require("./view/imgs/plus.png");
var removeIcon = require("./view/imgs/remove.png");

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
                nodeClick: "=",
                childrenLoader: "=",
                addItem: "=",
                removeItem: "="
            },
            require: [],
            restrict: "E",
            templateUrl: "directive/tree/node.html",
            link: function($scope, element, attributes, controllers) {
                $scope.open = false;
                $scope.add_btn = plusIcon;
                $scope.remove_btn = removeIcon;
                function load_children() {
                    var children = [];
                    if ($scope.childrenLoader) {
                        children = $scope.childrenLoader($scope.item);
                    }
                    $scope.subNodes = children;
                }
                $scope.wrap_node_click = function() {
                    if ($scope.item) {
                        var adaptedItem = $scope.adapter($scope.item);
                        if (adaptedItem.type === "branch") {
                            if ($scope.open) {
                                $scope.open = false;
                                $scope.folderClose && $scope.folderClose($scope.item);
                            }
                            else {
                                $scope.open = true;
                                $scope.folderOpen && $scope.folderOpen($scope.item);
                                load_children();
                            }
                        }
                        $scope.nodeClick && $scope.nodeClick($scope.item);

                    }
                    return false;
                };
                $scope.resolve_icon = function() {
                    var icon = null;
                    var adaptedItem = $scope.adapter($scope.item);
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
                    var adaptedItem = $scope.adapter($scope.item);
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
                $scope.add_child = function() {
                    console.log("add_child...")
                    if ($scope.addItem) {
                        console.log("call addItemWorker...")
                        $scope.addItem($scope.item)
                            .then(function() {
                                load_children();
                            })
                        ;
                    }
                    return false;
                };
                $scope.remove_self = function() {
                    if ($scope.removeItem) {
                        $scope.removeItem($scope.item)
                            .then(function() {
                                load_children();
                            })
                        ;
                    }
                    return false;
                };
            }
        };
    })
    .directive("tree", function () {
        var link = function($scope, element, attributes, controllers) {
            $scope.itemAdapter = $scope.adapter || function(item) {
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
                adapter: "=",
                icon: "=",
                folderOpen: "=",
                folderClose: "=",
                nodeClick: "=",
                childrenLoader: "=",
                addItem: "=",
                removeItem: "="
            },
            require: [],
            restrict: "E",
            templateUrl: "directive/tree/tree.html",
            link: link
        }
    })
;

module.exports = tree;
