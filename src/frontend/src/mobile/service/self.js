/**
 * Created by hyku on 2016/11/12.
 */

function service() {
    var svc = this;
    this.paths = [];
    this.walk_back = function() {
        svc.paths = [];
    };
    this.walk = function(path) {
        svc.paths.push(path);
    };
    this.footprints = function() {
        return svc.paths;
    };
}

module.exports = service;
