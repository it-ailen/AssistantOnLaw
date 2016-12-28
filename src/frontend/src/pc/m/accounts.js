/**
 * Created by hyku on 2016/11/27.
 */

"use strict";

function service(Configure, $http, $rootScope, $q, $state) {
    this.current = {};
    var status = {};
    var svc = this;
    this.login = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/accounts/login",
                    method: "POST",
                    data: data
                });
            })
            .then(function (res) {
                return res.data;
            })
            .then(function (me) {
                status.self = me;
                $rootScope.$broadcast("session.login", me);
                if (me.type === 'super') {
                    $state.go("frame.home.base.super");
                } else if (me.type === 'customer') {
                    $state.go("frame.home.base.customer");
                }
                return me;
            })
            ;
    };
    this.logout = function () {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/accounts/logout",
                    method: "POST"
                });
            })
            .then(function (res) {
                status.self = null;
                $rootScope.$broadcast("session.logout");
                $state.go("frame.home.base");
                // $rootScope.$emit("session.logout");
                return res.data;
            })
            ;
    };
    this.checkAuthentication = function () {
        var defer = $q.defer();
        if (status.self) {
            defer.resolve(status.self);
        } else {
            Configure.getHost()
                .then(function (host) {
                    return $http({
                        url: host + "/accounts/auth",
                        method: "GET"
                    });
                })
                .then(function (res) {
                    return res.data;
                })
                .then(function (me) {
                    status.self = me;
                    defer.resolve(status.self);
                })
                .catch(function (error) {
                    if (error.status === 401) {
                        $rootScope.$broadcast("session.auth_failed");
                        defer.resolve(null);
                    } else {
                        defer.reject(error)
                    }
                })
            ;
        }
        return defer.promise;
    }
}

module.exports = service;
