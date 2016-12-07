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
    };
    this.loadFaLvWenDaClasses = function () {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/classes",
                    method: "GET"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.createFaLvWenDaClass = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/classes",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateFaLvWenDaClass = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/classes/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.createFaLvWenDaArticle = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/articles",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateFaLvWenDaArticle = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/articles/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };

}

module.exports = func;
