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
    this.channelCreate = function(data) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/admin/channels",
                    method: "POST",
                    data: data,
                    headers: {
                        "Content-Type": undefined
                    }
                });
            })
            .then(function(res) {
                defer.resolve(res.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.channelUpdate = function(id, data) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/admin/channels/" + id,
                    method: "PUT",
                    data: data,
                    headers: {
                        "Content-Type": undefined
                    }
                });
            })
            .then(function(res) {
                defer.resolve(res.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.channelDelete = function(id) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/admin/channels/" + id,
                    method: "DELETE"
                });
            })
            .then(function(res) {
                defer.resolve(res.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
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
        svc.getConfig("host")
            .then(function(host) {
                var req = {
                    url: host + "/admin/channels",
                    method: "GET"
                };
                return $http(req);
            })
            .then(function(response) {
                var data = null;
                if (typeof response.data === "object") {
                    data = response.data;
                } else {
                    data = JSON.parse(response.data);
                }
                data.type = "entry";
                defer.resolve(data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.entryCreate = function(data) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                var req = {
                    url: host + "/admin/entries",
                    method: "POST",
                    data: data,
                    headers: {
                        "Content-Type": "application/json"
                    }
                };
                return $http(req)
            })
            .then(function(response) {
                defer.resolve(response.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.entryUpdate = function(id, data) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                var req = {
                    url: host + "/admin/entries/" + id,
                    method: "PUT",
                    data: data,
                    headers: {
                        "Content-Type": "application/json"
                    }
                };
                return $http(req)
            })
            .then(function(response) {
                defer.resolve(response.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.loadEntries = function(cid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                var req = {
                    url: host + "/admin/entries",
                    method: "GET",
                    params: {
                        channel_id: cid
                    }
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
                for (var i in data) {
                    data[i].type = "entry";
                }
                defer.resolve(data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.loadOptions = function(pid) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                var req = {
                    url: host + "/admin/options",
                    method: "GET",
                    params: {
                        parent_id: pid
                    }
                };
                return $http(req)
            })
            .then(function(response) {
                defer.resolve(response.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.optionCreate = function(data) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                var req = {
                    url: host + "/admin/options",
                    method: "POST",
                    data: data,
                    headers: {
                        "Content-Type": "application/json"
                    }
                };
                return $http(req)
            })
            .then(function(response) {
                defer.resolve(response.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.optionUpdate = function(id, data) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                var req = {
                    url: host + "/admin/options/" + id,
                    method: "PUT",
                    data: data,
                    headers: {
                        "Content-Type": "application/json"
                    }
                };
                return $http(req)
            })
            .then(function(response) {
                defer.resolve(response.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.optionDelete = function(option) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/admin/options/" + option.id,
                    method: "DELETE"
                });
            })
            .then(function(res) {
                defer.resolve(res.data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.loadIssues = function(params) {
        var defer = $q.defer();
        svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/admin/issues",
                    method: "GET",
                    params: params
                });
            })
            .then(function(res) {
                return res.data;
            })
            .then(function(data) {
                defer.resolve(data.list);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.submitSolution = function(id, data) {
        return svc.getConfig("host")
            .then(function(host) {
                return $http({
                    url: host + "/admin/issues/" + id + "/solutions",
                    method: "PUT",
                    data: data
                });
            })
            .then(function(res) {
                return res.data;
            })
        ;
    };
}

module.exports = service;
