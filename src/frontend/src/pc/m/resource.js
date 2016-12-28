/**
 * Created by hyku on 2016/12/4.
 */

"use strict";
function func(Configure, $http) {
    this.loadDirectoryTree = function (root) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/tree/" + root,
                    method: "GET"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.createFile = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateFile = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.deleteFile = function (id) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/" + id,
                    method: "DELETE"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.loadFaLvWenDaClasses = function () {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/classes",
                    method: "GET"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.createFaLvWenDaClass = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/classes",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateFaLvWenDaClass = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/classes/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.deleteFaLvWenDaClass = function (id) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/classes/" + id,
                    method: "DELETE"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.createFaLvWenDaArticle = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/articles",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateFaLvWenDaArticle = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/articles/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.deleteFaLvWenDaArticle = function (id) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/fa_lv_wen_da/articles/" + id,
                    method: "DELETE"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };

    this.loadMinShiSuSongDetail = function (stepId) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/su_song_wen_shu/min_shi_su_song/" + stepId,
                    method: "GET"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };

    this.saveMinShiSuSongDetail = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/su_song_wen_shu/min_shi_su_song/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (res) {
                return res.data;
            })
            ;
    };

    this.updateSuSongWenShuFile = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/su_song_wen_shu/files/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (res) {
                return res.data;
            })
            ;
    };

    this.addSuSongWenShuFile = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/su_song_wen_shu/files",
                    method: "POST",
                    data: data
                });
            })
            .then(function (res) {
                return res.data;
            })
            ;
    };

    this.deleteSuSongWenShuFile = function (id) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/resources/su_song_wen_shu/files/" + id,
                    method: "DELETE"
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };

    this.selectClasses = function (params) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/classes",
                    method: "GET",
                    params: params
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.insertClass = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/classes",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateClass = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/classes/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.selectEntries = function (params) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/entries",
                    method: "GET",
                    params: params
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.insertEntry = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/entries",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateEntry = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/entries/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.selectQuestions = function (params) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/questions",
                    method: "GET",
                    params: params
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.insertQuestion = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/questions",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateQuestion = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/questions/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.selectConclusions = function (params) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/conclusions",
                    method: "GET",
                    params: params
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.insertConclusions = function (data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/conclusions",
                    method: "POST",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
    this.updateConclusions = function (id, data) {
        return Configure.getHost()
            .then(function (host) {
                return $http({
                    url: host + "/counsels/conclusions/" + id,
                    method: "PUT",
                    data: data
                });
            })
            .then(function (resp) {
                return resp.data;
            })
            ;
    };
}

module.exports = func;
