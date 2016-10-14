/**
 * Created by hyku on 16/10/1.
 */

"use strict";


function service($http, $q) {
    var svc = this;
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
    this.reloadChannels = function() {
        if (svc.cache && svc.cache.channels) {
            svc.cache.channels = null;
        }
        return this.loadChannel();
    };
    this.loadChannels = function() {
        var defer = $q.defer();
        if (svc.cache && svc.cache.channels) {
            defer.resolve(svc.cache.channels);
        }
        else {
            svc.getConfig("host")
                .then(function(host) {
                    var req = {
                        url: host + "/static/demo/admin-channels.json",
                        method: "GET"
                    };
                    return $http(req)
                })
                .then(function(response) {
                    var data = null;
                    if (typeof response.data === "object") {
                        data = response.data;
                    } else {
                        data = JSON.parse(response.data);
                    }
                    if (!svc.cache) svc.cache = {};
                    svc.cache.channels = data;
                    defer.resolve(svc.cache.channels);
                })
                .catch(function(error) {
                    defer.reject(error);
                })
            ;
        }
        return defer.promise;
    };
    this.loadEntries = function(cid) {
        var defer = $q.defer();
        if (svc.cache && svc.cache.entries && svc.cache.entries[cid]) {
            defer.resolve(svc.cache.entries[cid]);
        }
        else {
            svc.getConfig("host")
                .then(function(host) {
                    var req = {
                        url: host + "/static/demo/admin-entries.json",
                        method: "GET"
                    };
                    return $http(req)
                })
                .then(function(response) {
                    var data = null;
                    if (typeof response.data === "object") {
                        data = response.data;
                    } else {
                        data = JSON.parse(response.data);
                    }
                    if (!svc.cache) svc.cache = {};
                    if (!svc.cache.entries) svc.cache.entries = {};
                    svc.cache.entries[cid] = data;
                    defer.resolve(data);
                })
                .catch(function(error) {
                    defer.reject(error);
                })
            ;
        }
        return defer.promise;

    };
}

module.exports = service;
