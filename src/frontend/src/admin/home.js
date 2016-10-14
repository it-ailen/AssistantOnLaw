/**
 * Created by hyku on 16/9/30.
 */

"use strict";

require("./view/style/home.less");

function controller($scope, AdminDataService) {
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
                text: "测试"
            },
            {
                type: "report",
                title: "Report test"
            }
        ];
    };
}

module.exports = controller;
