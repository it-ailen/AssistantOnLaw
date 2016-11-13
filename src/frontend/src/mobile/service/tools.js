/**
 * Created by hyku on 2016/11/13.
 */

function service(ngDialog, $q) {
    this.alert = function(message) {
        console.log("alert now...");
        var defer = $q.defer();
        var dialog = ngDialog.open({
            template: '<div style="overflow-x: hidden;">{{ message }}</div>' +
                      '<div style="text-align: center">' +
                      '    <input type="button" class="btn btn-default" value="确定" ng-click="closeThisDialog()">' +
                      '</div>',
            plain: true,
            controller: function($scope) {
                console.log("controller of alert...");
                $scope.message = message;
            },
            showClose: false,
            closeByDocument: false,
            width: "80%"
        });
        console.log(dialog);
        dialog.closePromise
            .then(function(data) {
                defer.resolve(data);
            })
            .catch(function(error) {
                defer.reject(error);
            })
        ;
        return defer.promise;
    }
}

module.exports = service;
