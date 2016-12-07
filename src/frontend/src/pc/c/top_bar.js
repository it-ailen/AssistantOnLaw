/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

function func($scope, AccountsService, Session, $state) {
    $scope.$on("session.login", function (event, data) {
        console.log("session.login");
        console.log(data);
        $scope.self = data;
    });
    $scope.$on("session.logout", function (event) {
        console.log("session.logout");
        $scope.self = null;
    });
    $scope.$on("session.auth_failed", function (event) {
        console.log("auth_failed....");
        $state.go("home.base");
    });
    console.log($state.current);
    $scope.session = Session;
    console.log($state.includes("customer"));
    console.log($state);
    console.log($state.current);
    console.log("in top_bar");
    $scope.current = AccountsService.current;
    $scope.login = function (data) {
        console.log("login");
        console.log(data);
        AccountsService.login(data);
    };
    $scope.register = function () {
    };

    $scope.logout = function () {
        AccountsService.logout();
    };
    Session.self()
        .then(function(self) {
            if (self && self.type) {
                if (self.type === 'customer' && !$state.is("home.base.customer")) {
                    $state.go("home.base.customer");
                }
                if (self.type === 'super' && !$state.is("home.base.super")) {
                    $state.go("home.base.super");
                }
                $scope.self = self;
            }
            else {
                if (!$state.is("home.base")) {
                    $state.go("home.base");
                }
            }
        })
    ;
}

module.exports = func;
