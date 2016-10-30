/**
 * Created by hyku on 2016/10/25.
 */

"use strict";
function service($http, $q, ngDialog) {
    var svc = this;
    this.alert = function(message) {
        var dialog = ngDialog.open({
            template: require("./v/alert.html"),
            plain: true,
            controller: function($scope) {},
            data: {
                message: message
            },
            showClose: false,
            closeByDocument: false
        });
        return dialog.closePromise;
    };
    this.confirm = function(message) {
        var defer = $q.defer();
        ngDialog.open({
            template: require("./v/confirm.html"),
            plain: true,
            controller: function($scope) {},
            data: {
                message: message
            },
            showClose: false,
            closeByDocument: false
        }).closePromise
            .then(function(res) {
                defer.resolve(res.value);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.uploadImage = function(file, tag, data) {
        var defer = $q.defer();
        if (data) {
            data.file = file;
        } else {
            data = {
                file: file
            }
        }
        $http({
            url: "/utils/images/" + tag,
            method: "POST",
            data: svc.map2fd(data),
            headers: {
                "Content-Type": undefined
            }
        })
            .then(function(res) {
                defer.resolve(res.data.path);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    };
    this.map2fd = function(m) {
        if (m) {
            var fd = new FormData();
            for (var k in m) {
                fd.append(k, m[k]);
            }
            return fd;
        }
        return null;
    };
}

module.exports = service;
