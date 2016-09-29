/**
 * Created by hyku on 16/9/29.
 */

var webpack = require("webpack");
var baseConfig = require("./webpack.config.base");

baseConfig.plugins = (baseConfig.plugins || []).concat([
    new webpack.optimize.UglifyJsPlugin({
        mangle: {
            except: ["$", "jQuery"]
        },
        compress: {
            warnings: false
        }
    })

]);

module.exports = baseConfig;
