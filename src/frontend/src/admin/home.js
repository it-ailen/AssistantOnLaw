/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require("./view/style/home.less");

function controller($scope, AdminDataService, $q, ngDialog, tools) {
    $scope.status = {};
    $scope.home = {
        posters: []
    };
    AdminDataService.loadChannels()
        .then(function(channels) {
            $scope.channels = channels;
            console.log($scope.channels);
        })
    ;
    $scope.editChannel = function(channel) {
        ngDialog.open({
            template: require("./view/form/channel.html"),
            plain: true,
            controller: require("./forms/channel"),
            data: {
                item: channel
            }
        }).closePromise
            .then(function(result) {
                var value = result.value;
                if (value.status === 'success') {
                    if (!channel) {
                        $scope.channels.push(value.data);
                    }
                }
            })
        ;
    };
    $scope.removeChannel = function(index, channel) {
        tools.confirm("删除Channel?" + channel.name)
            .then(function(ensure) {
                if (ensure) {
                    AdminDataService
                        .channelDelete(channel.id)
                        .then(function() {
                            $scope.channels.splice(index, 1);
                        })
                        .catch(function(error) {
                            console.error(error);
                        })
                    ;
                }
            })
            .catch(function(error) {
                console.error(error);
            })
        ;
    };
    $scope.channelClass = function(channel) {
        if ($scope.status.channel && $scope.status.channel.id === channel.id) {
            return "current";
        }
        return "";
    };
    $scope.channelClick = function(channel) {
        console.log("channel clicked: " + channel.name);
        if ($scope.status.channel && $scope.status.channel.id === channel) {
            return;
        }
        $scope.status.channel = channel;
        AdminDataService.loadEntries(channel.id)
            .then(function(entries) {
                $scope.entries = entries;
            })
        ;
    };
    $scope.editEntryOrOption = function(type, channel, item, parent) {
        var defer = $q.defer();
        if (type === 'option' || type === 'report') {
            var dialog = ngDialog.open({
                template: require("./view/form/option.html"),
                plain: true,
                controller: require("./forms/option"),
                data: {
                    parent: parent,
                    item: item
                },
                closeByDocument: false
            });
            dialog.closePromise
                .then(function(data) {
                    defer.resolve(data);
                })
            ;
        } else {
            ngDialog.open({
                template: require("./view/form/entry.html"),
                plain: true,
                controller: require("./forms/entry"),
                data: {
                    channel: channel,
                    item: item
                }
            })
                .closePromise
                .then(function(result) {
                    var value = result.value;
                    if (value.status === 'success') {
                        $scope.entries.push(value.data);
                    }
                    defer.resolve();
                })
                .catch(function(error) {
                    console.error(error);
                })
            ;
        }
        return defer.promise;
    };
    $scope.formatEntry = function(entry) {
        return {
            type: "directory",
            text: entry.title,
            data: entry
        };
    };
    $scope.entryAdapter = function(item) {
        var type = "leaf";
        var text = null;
        switch (item.type) {
            case "entry":
            case "option":
                type = "branch";
                text = item.text;
                break;
            case "report":
                type = "leaf";
                text = item.report.title;
                break;
        }
        return {
            type: type,
            text: text,
            data: item
        };
    };
    $scope.loadChildren = function(item) {
        return AdminDataService
            .loadOptions(item.id)
        ;
    };
    $scope.addItem = function(parent) {
        return $scope.editEntryOrOption("option", null, null, parent);
    };
    $scope.removeItem = function(option) {
        return AdminDataService
            .optionDelete(option)
        ;
    };
    $scope.itemClick = function(item) {
        console.log("item clicked");
        console.log(item);
        $scope.currentItem = item;
    };
    $scope.itemEdit = function(item) {
        console.log(item);
        return $scope.editEntryOrOption(item.type, $scope.status.channel, item);
    };
    $scope.editPageHome = function() {
        ngDialog.open({
            template: require("./view/form/pages/home.html"),
            plain: true,
            controller: require("./forms/pages/home"),
            data: {
                data: $scope.home
            },
            closeByDocument: false
        });
    };
}

module.exports = controller;
