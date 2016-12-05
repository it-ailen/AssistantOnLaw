/**
 * Created by hyku on 2016/12/4.
 */

"use strict";
function func(Configure, $http) {
    this.loadDirectoryTree = function (root) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/tree/" + root,
                    method: "GET"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.createFile = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateFile = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    }
}

module.exports = func;
