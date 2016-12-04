/**
 * Created by hyku on 2016/11/27.
 */

"use strict";

function provider() {
    var self = this;
    this.host = "";
    this.init = function(config) {
        self.host = config.host || "";
    };

    this.$get = function($q) {
        return {
            getHost: function() {
                return $q.when(self.host);
            }
        }
    };
}

module.exports = provider;
