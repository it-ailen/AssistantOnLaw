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
    this.loadLayoutHome = function() {
        var defer = $q.defer();
        if (svc.cache && svc.cache.layout && svc.cache.layout.home) {
            defer.resolve(svc.cache.layout.home);
        }
        else {
            svc.getConfig("host")
                .then(function(host) {
                    var req = {
                        url: host + "/mobile/layout/home",
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
    this.loadLayoutChannel = function(cid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/mobile/layout/channels/" + cid,
                    method: "GET"
                });
            })
            .then(function(res) {
                return res.data;
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
    this.loadLayoutEntry = function(eid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/mobile/layout/entries/" + eid,
                    method: "GET"
                })
            })
            .then(function(res) {
                return res.data;
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
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/utils/files/self-consultant",
                    method: "POST",
                    data: fd,
                    headers: {
                        "Content-Type": undefined
                    }
                });
            })
            .then(function(res) {
                return res.data;
            })
            .then(function(data) {
                defer.resolve(data.uri);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.post_issue = function(data) {
        console.log(data);
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/mobile/issues",
                    method: "POST",
                    data: data
                });
            })
            .then(function(data) {
                defer.resolve(data.value);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
}

module.exports = service;
