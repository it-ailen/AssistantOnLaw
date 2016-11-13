/**
 * Created by Allen on 2016/10/1.
 */

"use strict";

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
        console.log("type: " + type);
        var index = type2index(type);
        console.log("index");
        if (svc.pageStack.length > index) {
            console.log(svc.pageStack);
            console.log(index);
            return svc.pageStack[index].meta;
        }
        throw "Current page is empty: " + type;
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
                    id: cid,
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
                    channelId: "123",
                    id: eid,
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
                                channelId: "123",
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
                    channelId: "123",
                    actions: [
                        {
                            id: "action12dsf",
                            type: "step",
                            text: "action1"
                        },
                        {
                            id: "action231d",
                            type: "report",
                            channelId: "123",
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
    this.loadReport = function(rid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                // TODO load data from server
                var d = $q.defer();
                d.resolve({
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
    this.upload_file = function(file) {
        console.log(file);
        var fd = new FormData();
        fd.append("file", file);
        var defer = $q.defer();
        // svc.getConfig("host")
        //     .then(function(host) {
        //         return $http({
        //             url: host + "/utils/images",
        //             method: "POST",
        //             data: fd,
        //             headers: {
        //                 "Content-Type": undefined
        //             }
        //         });
        //     })
        //     .then(function(res) {
        //         defer.resolve(res.data);
        //     })
        //     .catch(function(error) {
        //         defer.reject(error);
        //     })
        // ;
        defer.resolve("test-id");
        return defer.promise;
    };
    this.post_consulting = function(data) {
        console.log(data);
        var defer = $q.defer();
        // svc.getConfig("host")
        //     .then(function(host) {
        //         return $http({
        //             url: host + "/consultancy",
        //             method: "POST",
        //             data: data
        //         });
        //     })
        //     .then(function(data) {
        //         defer.resolve(data.value);
        //     })
        //     .catch(function(error) {
        //         defer.reject(error);
        //     })
        // ;
        defer.resolve();
        return defer.promise;
    };
}

module.exports = service;
