/**
 * Created by hyku on 2016/11/19.
 */

"use strict";

require("./styles/home.less");

function service($scope, $anchorScroll, $location) {
    console.log("home...");
    $scope.entries = [
        {
            text: "诉讼文书",
            icon: "icon-icoo_susongws",
            hash: "suSongWenShu"
        },
        {
            text: "协议范本",
            icon: "icon-icoo_xieyifb",
            hash: "xieYiFanBen"
        },
        {
            text: "法律问答",
            icon: "icon-icoo_wendafl",
            hash: "faLvWenDa"
        },
        {
            text: "法律咨询",
            icon: "icon-icoo_zixunfl",
            hash: "faLvZiXun"
        },
        {
            text: "人工咨询",
            icon: "icon-icoo_kefurg",
            hash: "renGongZiXun"
        }
    ];
    $scope.goto = function(hash) {
        if ($location.hash() !== hash) {
            $location.hash(hash);
        } else {
            $anchorScroll();
        }
    };
}

module.exports = service;
