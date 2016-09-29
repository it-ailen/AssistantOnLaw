/**
 * Created by hyku on 16/9/29.
 */
var path = require("path");
var webpack = require("webpack");
const autoprefixer = require('autoprefixer');
const cssnano = require('cssnano');

var config = {
    entry: {
        mobile: "./src/mobile.js"
    },
    output: {
        path: path.resolve(__dirname, "../entries/dist"),
        filename: "[name].js"
    },
    resolve: {
        root: path.resolve(__dirname, "./src")
    },
    module: {
        loaders: [
            {
                test: /\.css$/i,
                loader: "style-loader!css-loader!postcss-loader"
            },
            {
                test: /\.less$/i,
                loader: "style-loader!css-loader!postcss-loader!less-loader"
            },
            {
                test: /\.json$/i,
                loader: "json"
            },
            {
                test: /\.(jpe?g|png|gif|svg)$/i,
                loader: "url-loader?limit=10000&name=images/[name].[ext]"
            }
        ]
    },
    postcss: function () {
        return {
            plugins: [autoprefixer, cssnano]
        };
    },
    plugins: [
        new webpack.optimize.DedupePlugin(),
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery"
        }),
        new webpack.optimize.OccurenceOrderPlugin()
    ]
};

module.exports = config;
