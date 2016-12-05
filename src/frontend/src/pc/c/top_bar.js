/**
 * Created by hyku on 2016/12/3.
 */

"use strict";

function func($scope, AccountsService) {
    console.log("in top_bar")
    $scope.login = function(data) {
        console.log("login");
        console.log(data);
        AccountsService.login(data)
            .then(function(me) {
            })
        ;
    };
    $scope.register = function() {

    };
}

module.exports = func;
