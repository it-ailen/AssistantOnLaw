/**
 * Created by hyku on 2016/11/12.
 */

function service() {
    var svc = this;
    this.paths = [];
    this.walk_back = function() {
        svc.paths = [];
    };
    this.walk_backward = function() {
        svc.paths.pop();
    };
    this.walk_forward = function(type, data) {
        svc.paths.push({
            type: type,
            data: data
        });
    };

    this.footprints = function() {
        return svc.paths;
    };
}

module.exports = service;
