/**
 * Created by hyku on 2016/12/11.
 */

"use strict";

angular.module("svg.flow", [])
    .directive("intermediateNode", function ($parse) {
        return {
            restrict: "A",
            type: "svg",
            scope: {
                text: "@",
                borderStroke: "@",
                background: "@",
                textColor: "@",
                fontFamily: "@",
                fontSize: "@",
                data: "=intermediateNode",
                onEvent: "&"
            },
            template: require("./v/step.pug"),
            link: function ($scope, ele, iAttr) {
                angular.element(ele)
                    .on("click mouseover mouseleave", function(e) {
                        console.log("event occurs");
                        // console.log(iAttr.onEvent);
                        // if (iAttr.onEvent) {
                        //     var fn = $parse(iAttr.onEvent);
                        //     console.log(fn);
                        //     fn($scope, {
                        //         $event: e,
                        //         $element: ele,
                        //         $data: $scope.data
                        //     });
                        // }
                        console.log($scope.onEvent);
                        if ($scope.onEvent) {
                            $scope.onEvent({
                                $event: e,
                                // $element: ele,
                                // $data: $scope.data
                            });
                        }
                    })
                ;
            }
        };
    })
;

module.exports = "svg.flow";
