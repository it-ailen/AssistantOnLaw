/**
 * Created by hyku on 16/10/1.
 */

"use strict";

require("./style/channel.less");

function routine($routeParams, MobileDataService, $scope, $location, self, tools) {
    var channelId = $routeParams.id;
    self.walk_forward("channel", channelId);
    MobileDataService
        .loadLayoutChannel(channelId)
        .then(function(channel) {
            $scope.channel = channel;
            console.log(channel);
        })
        .catch(function(error) {
            console.error(error);
            tools.alert(error);
        })
    ;
}

module.exports = routine;
