/**
 * Created by hyku on 2016/11/27.
 */

"use strict";

function service(Configure, $http, $rootScope) {
    this.self = null;
    var svc = this;
    this.login = function(data) {
        return Configure.getHost()
            .then(function(host) {
                return $http({
                    url: host + "/accounts/login",
                    method: "POST",
                    data: data
                });
            })
            .then(function(res) {
                return res.data;
            })
            .then(function(me) {
                svc.self = me;
                $rootScope.$broadcast("login", me);
                return me;
            })
        ;
    };
    this.logout = function() {
        return Configure.getHost()
            .then(function(host) {
                return $http({
                    url: host + "/accounts/logout",
                    method: "POST"
                });
            })
            .then(function(res) {
                $rootScope.$broadcast("logout");
                return res.data;
            })
        ;
    }
}

module.exports = service;
