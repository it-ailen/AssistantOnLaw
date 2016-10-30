/**
 * Created by hyku on 2016/10/30.
 */

"use strict";
function routine($scope, tools) {
    $scope.data = $scope.ngDialogData.data;
    $scope.addPoster = function() {
        $scope.data.posters.push("");
    };

    $scope.uploadHomePoster = function(file) {
        return tools.uploadImage(file, "home.posters", {});
    };
    $scope.srcToUrl = function(src) {
        console.log("src: " + src);
        return src;
    };
}

module.exports = routine;
