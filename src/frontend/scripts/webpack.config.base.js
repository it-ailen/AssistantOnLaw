/**
 * Created by hyku on 16/9/29.
 */
var path = require("path");
var webpack = require("webpack");
const autoprefixer = require('autoprefixer');
const cssnano = require('cssnano');
// var HtmlWebpackPlugin = require("html-webpack-plugin");
var ExtractTextPlugin = require("extract-text-webpack-plugin");
// var PathRewriterPlugin = require('webpack-path-rewriter');
var htmlExtractor = new ExtractTextPlugin("html", "[name].html");

var config = {
    entry: {
        "mobile.js": "./src/mobile.js",
        "admin.js": "./src/admin.js",
        "pc.js": "./src/pc.js",
        "pc": "./src/pc.jade"
    },
    output: {
        path: path.resolve(__dirname, "../entries/dist"),
        publicPath: "./dist/",
        filename: "[name]"
    },
    resolve: {
        root: path.resolve(__dirname, "./src")
    },
    module: {
        loaders: [
            {
                test: /\.less$/i,
                loaders: ["style", "css", "less"]
            },
            {
                test: /\.css$/i,
                loaders: ["style", "css"]
            },
            {
                test: /\.json$/i,
                loader: "json"
            },
            {
                test: /\.(jpe?g|png|gif|svg)$/i,
                loader: "url-loader?limit=10000&name=images/[name].[ext]"
            },
            {
                test: /\.(ttf|eot|woff2?)$/,
                loader: 'file?name=etc/[name].[ext]'
            },
            {
                test: /\.html$/i,
                loaders: ["html"]
            },
            {
                test: /\.jade$/i,
                loaders: ["pug-html"]
            },
            {
                test: /pc\.jade$/i,
                // loader: "source!jade-static"
                // loader: ExtractTextPlugin.extract("pug-html")
                loader: htmlExtractor.extract("html", "apply!pug?pretty=true")
            },
            // {
            //     test: /\.jade$/i,
            //     loaders: ["pug-loader"]
            // }
        ]
    },
    postcss: function () {
        return {
            plugins: [autoprefixer, cssnano]
        };
    },
    plugins: [
        htmlExtractor,
        new webpack.optimize.DedupePlugin(),
        new webpack.ProvidePlugin({
            $: "jquery",
            jQuery: "jquery"
        }),
        new webpack.optimize.OccurenceOrderPlugin(),
    ]
};

module.exports = config;
