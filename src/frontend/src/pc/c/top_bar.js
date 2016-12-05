/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

function func($scope, AccountsService) {
    console.log("in top_bar");
    $scope.current = AccountsService.current;
    $scope.login = function(data) {
        console.log("login");
        console.log(data);
        AccountsService.login(data);
    };
    $scope.register = function() {
    };

    $scope.logout = function() {
        AccountsService.logout();
    };
}

module.exports = func;
