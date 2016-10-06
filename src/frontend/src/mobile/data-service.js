/**
 * Created by hyku on 16/10/1.
 */

"use strict";

var pageTypes = ["home", "channel", "entry", "report"];

function service($http, $q) {
    var svc = this;
    this.pageStack = [];
    function type2index(type) {
        var index = pageTypes.indexOf(type);
        if (index >= 0) {
            return index;
        }
        throw "Unknown type: " + type;
    }
    this.pagePush = function(type, meta) {
        var index = type2index(type);
        var data = {
            type: type,
            meta: meta
        };
        svc.pageStack.splice(index, svc.pageStack.length, data);
    };
    this.pagePopTo = function(type) {
        var index = type2index(type);
        svc.pageStack.splice(index + 1, svc.pageStack.length);
    };
    this.currentMeta = function(type) {
        var index = type2index(type);
        if (svc.pageStack.length > index) {
            return svc.pageStack[index].meta;
        } else {
            throw "Current page is empty: " + type;
        }
    };
    this.getConfig = function(key) {
        var defer = $q.defer();
        if (svc.configs) {
            defer.resolve(svc.configs[key]);
        } else {
            svc.configs = {
                host: "" // 使用当前 host
            };
            defer.resolve(svc.configs[key]);
        }
        return defer.promise;
    };
    this.loadHome = function() {
        var defer = $q.defer();
        if (svc.cache && svc.cache.layout && svc.cache.layout.home) {
            defer.resolve(svc.cache.layout.home);
        }
        else {
            svc.getConfig("host")
                .then(function(host) {
                    var req = {
                        url: host + "/static/layout/home.json",
                        method: "GET"
                    };
                    return $http(req)
                })
                .then(function(response) {
                    var layout = null;
                    if (typeof response.data === "object") {
                        layout = response.data;
                    } else {
                        layout = JSON.parse(response.data);
                    }
                    if (!svc.cache) svc.cache = {};
                    if (!svc.cache.layout) svc.cache.layout = {};
                    svc.cache.layout.home = layout;
                    defer.resolve(svc.cache.layout.home);
                })
                .catch(function(error) {
                    defer.reject(error);
                })
            ;
        }
        return defer.promise;
    };
    this.loadChannel = function(cid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                // TODO load data from server
                var d = $q.defer();
                d.resolve({
                    title: "test",
                    poster: {
                        url: "/ugc/demo.jpg",
                        title: "test",
                        description: "description"
                    },
                    entries: [
                        {
                            id: "2321",
                            icon: "/ugc/icon-demo.jpg",
                            text: "Test"
                        }
                    ]
                });
                return d.promise;
            })
            .then(function(data) {
                defer.resolve(data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.loadEntry = function(eid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                // TODO load data from server
                var d = $q.defer();
                d.resolve({
                    title: "test",
                    layout: "single-page",
                    step: {
                        id: "1312321",
                        title: "测试",
                        actions: [
                            {
                                id: "action12dsf",
                                type: "step",
                                text: "action1"
                            },
                            {
                                id: "action231d",
                                type: "report",
                                text: "action2"
                            }
                        ]

                    }
                });
                return d.promise;
            })
            .then(function(data) {
                defer.resolve(data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.loadStep = function(sid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                // TODO load data from server
                var d = $q.defer();
                d.resolve({
                    id: "1312321" + Math.random(),
                    title: "测试",
                    actions: [
                        {
                            id: "action12dsf",
                            type: "step",
                            text: "action1"
                        },
                        {
                            id: "action231d",
                            type: "report",
                            text: "action2"
                        }
                    ]
                });
                return d.promise;
            })
            .then(function(data) {
                defer.resolve(data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
}

module.exports = service;
