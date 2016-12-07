/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

function func($scope, ResourceService, ngDialog) {
    $scope.current = {};
    $scope.createRoot = function () {
        ngDialog.open({
            template: require("./v/directory.pug"),
            plain: true,
            controller: require("./forms/directory"),
            data: {
                parent: "xie_yi_fan_ben"
            },
            closeByDocument: false
        }).closePromise
            .then(function (data) {
                return data.value;
            })
            .then(function (data) {
                reload();
            })
        ;
    };
    $scope.contextMenu = [
        ["新建文件夹", function ($itemScope, $event, modelValue, text, $li) {
            if ($itemScope.item.properties.type === 'directory') {
                ngDialog.open({
                    template: require("./v/directory.pug"),
                    plain: true,
                    controller: require("./forms/directory"),
                    data: {
                        parent: $itemScope.item.properties.id
                    },
                    closeByDocument: false
                }).closePromise
                    .then(function (data) {
                        console.log(data);
                        return data.value;
                    })
                    .then(function (data) {
                        reload();
                    })
                ;
            }
        }],
        ["添加文件", function ($itemScope, $event, modelValue, text, $li) {
            ngDialog.open({
                template: require("./v/file.pug"),
                plain: true,
                controller: require("./forms/file"),
                data: {
                    parent: $itemScope.item.properties.id
                },
                closeByDocument: false
            })
                .closePromise
                .then(function (data) {
                    console.log(data);
                    return data.value;
                })
                .then(function (data) {
                    console.log(data);
                    reload();
                })
            ;
        }],
        ["编辑", function ($itemScope, $event, modelValue, text, $li) {
            var promise = null;
            if ($itemScope.item.properties.type === 'directory') {
                promise = ngDialog.open({
                    template: require("./v/directory.pug"),
                    plain: true,
                    controller: require("./forms/directory"),
                    data: {
                        item: $itemScope.item
                    },
                    closeByDocument: false
                }).closePromise;
            } else {
                promise = ngDialog.open({
                    template: require("./v/file.pug"),
                    plain: true,
                    controller: require("./forms/file"),
                    data: {
                        item: $itemScope.item
                    },
                    closeByDocument: false
                }).closePromise;
            }
            promise
                .then(function (data) {
                    console.log(data);
                    return data.value;
                })
                .then(function (data) {
                    if (data.status === 'ok') {
                        $itemScope.item.properties = data.data;
                    }
                })
            ;
        }],
        ["删除", function ($itemScope, $event, modelValue, text, $li) {
            console.log(arguments);
        }]
    ];
    $scope.itemClass = function (item) {
        var cls = [item.properties.type];
        if (item.properties.type === 'directory' && item.opened) {
            cls.push("open")
        }
        if (item.indent) {
            cls.push("indent-" + item.indent);
        }
        return cls;
    };
    $scope.toggle = function (item, index) {
        function open(t) {
            var children = t.children || [];
            var childrenCount = children.length;
            for (var i in children) {
                children[i].indent = (item.indent || 1) + 1;
            }
            if (childrenCount > 0) {
                Array.prototype.splice.apply($scope.current.expandedRows,
                    [index + 1, 0].concat(children));
            }
            t.opened = true;
        }

        function close(t) {
            var childrenCount = t.children && t.children.length || 0;
            $scope.current.expandedRows.splice(index + 1, childrenCount);
            t.opened = false;
        }

        if (item.properties.type !== 'directory') {
            return false;
        }
        var children = item.children || [];
        if (item.opened) {
            /* close now */
            for (var i in children) {
                close(children[i]);
            }
            close(item);
        } else {
            open(item);
        }
        return false;
    };
    $scope.data = {};
    function reload() {
        ResourceService.loadDirectoryTree("xie_yi_fan_ben")
            .then(function (tree) {
                $scope.data.whole = tree;
                $scope.current.expandedRows = tree.children;
            })
            .catch(function (error) {
                console.error(error);
            })
        ;
    }

    reload();
}

module.exports = func;
