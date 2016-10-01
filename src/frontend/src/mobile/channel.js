/**
 * Created by hyku on 16/10/1.
 */

"use strict";

require("./view/style/channel.less");

function routine($routeParams, MobileDataService, $scope) {
    MobileDataService.loadChannel($routeParams.id)
        .then(function(channel) {
            $scope.channel = channel;
            console.log($scope.channel);
        })
        .catch(function(err) {
            console.error(err);
        })
    ;
}

module.exports = routine;
