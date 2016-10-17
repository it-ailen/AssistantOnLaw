/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require("./view/style/home.less");

function controller($scope, AdminDataService, $q, ngDialog) {
    $scope.status = {};
    AdminDataService.loadChannels()
        .then(function(channels) {
            $scope.channels = channels;
            console.log($scope.channels);
        })
    ;
    $scope.channelClass = function(channel) {
        if ($scope.status.channel && $scope.status.channel.id === channel.id) {
            return "current";
        }
        return "";
    };
    $scope.channelClick = function(channel) {
        if ($scope.status.channel && $scope.status.channel.id === channel) {
            return;
        }
        $scope.status.channel = channel;
        AdminDataService.loadEntries(channel.id)
            .then(function(entries) {
                $scope.entries = entries;
                console.log($scope.entries);
            })
        ;
    };
    $scope.formatEntry = function(entry) {
        return {
            type: "directory",
            text: entry.title,
            data: entry
        };
    };
    $scope.entry_adapter = function(item) {
        var type = "leaf";
        var text = null;
        switch (item.type) {
            case "entry":
            case "step":
                type = "branch";
                text = item.text;
                break;
            case "report":
                type = "leaf";
                text = item.title;
                break;
        }
        var res = {
            type: type,
            text: text,
            data: item
        };
        return res;
    };
    $scope.load_children = function(item) {
        return [
            {
                type: "step",
                text: "测试",
                id: "sfsadfdf"
            },
            {
                id: rid,
                title: "测试",
                channelId: "123",
                conclusion: {
                    detail: "测试结论"
                },
                decrees: [
                    {
                        source: "《中国人民共和国宪法》",
                        content: "第一款第一条: 测试"
                    }
                ],
                cases: [
                    {
                        intro: "案例简介",
                        link: "http://detail..."
                    }
                ]
            }
        ];
    };
    $scope.add_item = function(parent) {
        var defer = $q.defer();
        console.log(parent);
        var dialog = ngDialog.open({
            template: require("./view/form/step.or.report.html"),
            plain: true,
            controller: require("./form.step.or.report"),
            data: {
                parent: parent
            },
            closeByDocument: false
        });
        dialog.closePromise
            .then(function(data) {
                console.log("!!!!")
                console.log(data);
                defer.resolve(data);
            })
        ;
        return defer.promise;
    };
    $scope.remove_item = function(item) {
        var defer = $q.defer();
        console.log("remove item now");
        console.log(parent);
        defer.resolve();
        return defer.promise;
    };
    $scope.item_click = function(item) {
        console.log("item clicked");
        console.log(item);
        $scope.currentItem = item;
    };
}

module.exports = controller;
