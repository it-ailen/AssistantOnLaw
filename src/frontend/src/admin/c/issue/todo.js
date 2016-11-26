/**
 * Created by hyku on 2016/11/15.
 */

"use strict";

function controller($scope, AdminDataService, ngDialog) {
    function init() {
        $scope.data = {};
        $scope.status = {};
        $scope.solution = {};
    }
    function reloadIssues() {
        AdminDataService
            .loadIssues()
            .then(function(list) {
                $scope.data.issues = list;
            })
            .catch(function(error) {
                console.error(error);
            })
        ;
    }
    function reload() {
        init();
        reloadIssues();
    }
    $scope.expandIssue = function(issue) {
        $scope.status.current = issue;
        $scope.solution.solution = issue.solution;
        $scope.solution.tags = issue.tags || [];
    };
    reload();
    $scope.submit = function() {
        var data = {
            solution: $scope.solution.solution,
            tags: $scope.solution.tags.split(" ")
        };
        AdminDataService
            .submitSolution($scope.status.current.id, data)
            .then(function() {
                reloadIssues();
            })
            .catch(function(error) {
                console.error(error);
            })
        ;
    };
}

module.exports = controller;
