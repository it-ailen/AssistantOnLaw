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
}

module.exports = service;
