/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};

/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {

/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId])
/******/ 			return installedModules[moduleId].exports;

/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			exports: {},
/******/ 			id: moduleId,
/******/ 			loaded: false
/******/ 		};

/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);

/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;

/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}


/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;

/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;

/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";

/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ function(module, exports, __webpack_require__) {

	/**
	 * Created by hyku on 2016/10/13.
	 */

	"use strict";

	// var angular = require("angular");

	__webpack_require__(4);
	__webpack_require__(3);
	__webpack_require__(6);
	var fileIcon = __webpack_require__(7);
	var folderIcon = __webpack_require__(9);
	var closedFolderIcon = __webpack_require__(8);

	var tree = angular.module("tree", []);
	tree
	    .directive("treeNode", function () {
	        return {
	            scope: {
	                item: "=",
	                adapter: "=",
	                icon: "=",
	                folderOpen: "=",
	                folderClose: "=",
	                leafClick: "=",
	                childrenLoader: "="
	            },
	            require: [],
	            restrict: "E",
	            templateUrl: "directive/tree/node.html",
	            link: function($scope, element, attributes, controllers) {
	                console.log("link tree node");
	                $scope.open = false;
	                $scope.node_click = function() {
	                    if ($scope.item) {
	                        var adaptedItem = $scope.adapter && $scope.adapter($scope.item);
	                        if (adaptedItem.type === "branch") {
	                            if ($scope.open) {
	                                $scope.open = false;
	                                $scope.folderClose && $scope.folderClose($scope.item);
	                            }
	                            else {
	                                $scope.open = true;
	                                $scope.folderOpen && $scope.folderOpen($scope.item);
	                                console.log("load subitems now");
	                                console.log($scope.childrenLoader);
	                                $scope.subNodes = $scope.childrenLoader && $scope.childrenLoader($scope.item);
	                                console.log($scope.subNodes);
	                            }
	                        }
	                        else {
	                            $scope.leafClick && $scope.leafClick($scope.item);
	                        }
	                    }
	                    return false;
	                };
	                $scope.resolve_icon = function() {
	                    var icon = null;
	                    var adaptedItem = $scope.adapter && $scope.adapter($scope.item);
	                    if (adaptedItem.type === 'branch') {
	                        icon = ($scope.icon && $scope.icon($scope.item, $scope.open))
	                            || (!$scope.open && closedFolderIcon)
	                            || ($scope.open && folderIcon);
	                    }
	                    else {
	                        icon = ($scope.icon && $scope.icon($scope.item))
	                            || fileIcon;
	                    }
	                    return icon;
	                };
	                $scope.node_class = function() {
	                    var classes = ["node"];
	                    var adaptedItem = $scope.adapter && $scope.adapter($scope.item);
	                    if (adaptedItem.type === 'branch') {
	                        classes.push("branch");
	                        if ($scope.open) {
	                            classes.push("open");
	                        }
	                        else {
	                            classes.push("closed");
	                        }
	                    }
	                    else {
	                        classes.push("leaf");
	                    }
	                    return classes;
	                };
	            }
	        };
	    })
	    .directive("tree", function () {
	        var link = function($scope, element, attributes, controllers) {
	            console.log($scope.root);
	            $scope.adapter = $scope.itemAdapter || function(item) {
	                    console.log("in tree .adapter");
	                    return item;
	                };
	            $scope.tree_class = function() {
	                var classes = ["tree"];
	                return classes;
	            }
	        };
	        return {
	            scope: {
	                root: "=root",
	                itemAdapter: "=adapter",
	                icon: "=",
	                folderOpen: "=",
	                folderClose: "=",
	                leafClick: "=",
	                childrenLoader: "="
	            },
	            require: [],
	            restrict: "E",
	            templateUrl: "directive/tree/tree.html",
	            link: link
	        }
	    })
	;

	module.exports = tree;


/***/ },
/* 1 */
/***/ function(module, exports, __webpack_require__) {

	exports = module.exports = __webpack_require__(2)();
	// imports


	// module
	exports.push([module.id, ".tree {\n  overflow: auto;\n}\n.tree .node .directory-level {\n  padding-right: 4px;\n  white-space: nowrap;\n}\n.tree .node .sub-node {\n  padding-left: 14px;\n}\n", ""]);

	// exports


/***/ },
/* 2 */
/***/ function(module, exports) {

	/*
		MIT License http://www.opensource.org/licenses/mit-license.php
		Author Tobias Koppers @sokra
	*/
	// css base code, injected by the css-loader
	module.exports = function() {
		var list = [];

		// return the list of modules as css string
		list.toString = function toString() {
			var result = [];
			for(var i = 0; i < this.length; i++) {
				var item = this[i];
				if(item[2]) {
					result.push("@media " + item[2] + "{" + item[1] + "}");
				} else {
					result.push(item[1]);
				}
			}
			return result.join("");
		};

		// import a list of modules into the list
		list.i = function(modules, mediaQuery) {
			if(typeof modules === "string")
				modules = [[null, modules, ""]];
			var alreadyImportedModules = {};
			for(var i = 0; i < this.length; i++) {
				var id = this[i][0];
				if(typeof id === "number")
					alreadyImportedModules[id] = true;
			}
			for(i = 0; i < modules.length; i++) {
				var item = modules[i];
				// skip already imported module
				// this implementation is not 100% perfect for weird media query combinations
				//  when a module is imported multiple times with different media queries.
				//  I hope this will never occur (Hey this way we have smaller bundles)
				if(typeof item[0] !== "number" || !alreadyImportedModules[item[0]]) {
					if(mediaQuery && !item[2]) {
						item[2] = mediaQuery;
					} else if(mediaQuery) {
						item[2] = "(" + item[2] + ") and (" + mediaQuery + ")";
					}
					list.push(item);
				}
			}
		};
		return list;
	};


/***/ },
/* 3 */
/***/ function(module, exports) {

	var angular=window.angular,ngModule;
	try {ngModule=angular.module(["ng"])}
	catch(e){ngModule=angular.module("ng",[])}
	var v1="<div ng-class=\"node_class()\">\n<div class=\"directory-level\" ng-click=\"node_click()\">\n<img ng-src=\"{{ resolve_icon() }}\">\n<span>{{ adapter(item).text }}</span>\n</div>\n<div class=\"sub-node\" ng-if=\"open\" ng-repeat=\"node in subNodes\">\n<tree-node item=\"node\" adapter=\"adapter\" icon=\"icon\" folder-open=\"folderOpen\" folder-close=\"folderClose\" leaf-click=\"leafClick\" children-loader=\"childrenLoader\">\n</tree-node>\n</div>\n</div>";
	var id1="directive/tree/node.html";
	var inj=angular.element(window.document).injector();
	if(inj){inj.get("$templateCache").put(id1,v1);}
	else{ngModule.run(["$templateCache",function(c){c.put(id1,v1)}]);}
	module.exports=v1;

/***/ },
/* 4 */
/***/ function(module, exports) {

	var angular=window.angular,ngModule;
	try {ngModule=angular.module(["ng"])}
	catch(e){ngModule=angular.module("ng",[])}
	var v1="<div ng-class=\"tree_class()\">\n<tree-node item=\"root\" adapter=\"adapter\" icon=\"icon\" folder-open=\"folderOpen\" folder-close=\"folderClose\" leaf-click=\"leafClick\" children-loader=\"childrenLoader\">\n</tree-node>\n</div>";
	var id1="directive/tree/tree.html";
	var inj=angular.element(window.document).injector();
	if(inj){inj.get("$templateCache").put(id1,v1);}
	else{ngModule.run(["$templateCache",function(c){c.put(id1,v1)}]);}
	module.exports=v1;

/***/ },
/* 5 */
/***/ function(module, exports, __webpack_require__) {

	/*
		MIT License http://www.opensource.org/licenses/mit-license.php
		Author Tobias Koppers @sokra
	*/
	var stylesInDom = {},
		memoize = function(fn) {
			var memo;
			return function () {
				if (typeof memo === "undefined") memo = fn.apply(this, arguments);
				return memo;
			};
		},
		isOldIE = memoize(function() {
			return /msie [6-9]\b/.test(window.navigator.userAgent.toLowerCase());
		}),
		getHeadElement = memoize(function () {
			return document.head || document.getElementsByTagName("head")[0];
		}),
		singletonElement = null,
		singletonCounter = 0,
		styleElementsInsertedAtTop = [];

	module.exports = function(list, options) {
		if(false) {
			if(typeof document !== "object") throw new Error("The style-loader cannot be used in a non-browser environment");
		}

		options = options || {};
		// Force single-tag solution on IE6-9, which has a hard limit on the # of <style>
		// tags it will allow on a page
		if (typeof options.singleton === "undefined") options.singleton = isOldIE();

		// By default, add <style> tags to the bottom of <head>.
		if (typeof options.insertAt === "undefined") options.insertAt = "bottom";

		var styles = listToStyles(list);
		addStylesToDom(styles, options);

		return function update(newList) {
			var mayRemove = [];
			for(var i = 0; i < styles.length; i++) {
				var item = styles[i];
				var domStyle = stylesInDom[item.id];
				domStyle.refs--;
				mayRemove.push(domStyle);
			}
			if(newList) {
				var newStyles = listToStyles(newList);
				addStylesToDom(newStyles, options);
			}
			for(var i = 0; i < mayRemove.length; i++) {
				var domStyle = mayRemove[i];
				if(domStyle.refs === 0) {
					for(var j = 0; j < domStyle.parts.length; j++)
						domStyle.parts[j]();
					delete stylesInDom[domStyle.id];
				}
			}
		};
	}

	function addStylesToDom(styles, options) {
		for(var i = 0; i < styles.length; i++) {
			var item = styles[i];
			var domStyle = stylesInDom[item.id];
			if(domStyle) {
				domStyle.refs++;
				for(var j = 0; j < domStyle.parts.length; j++) {
					domStyle.parts[j](item.parts[j]);
				}
				for(; j < item.parts.length; j++) {
					domStyle.parts.push(addStyle(item.parts[j], options));
				}
			} else {
				var parts = [];
				for(var j = 0; j < item.parts.length; j++) {
					parts.push(addStyle(item.parts[j], options));
				}
				stylesInDom[item.id] = {id: item.id, refs: 1, parts: parts};
			}
		}
	}

	function listToStyles(list) {
		var styles = [];
		var newStyles = {};
		for(var i = 0; i < list.length; i++) {
			var item = list[i];
			var id = item[0];
			var css = item[1];
			var media = item[2];
			var sourceMap = item[3];
			var part = {css: css, media: media, sourceMap: sourceMap};
			if(!newStyles[id])
				styles.push(newStyles[id] = {id: id, parts: [part]});
			else
				newStyles[id].parts.push(part);
		}
		return styles;
	}

	function insertStyleElement(options, styleElement) {
		var head = getHeadElement();
		var lastStyleElementInsertedAtTop = styleElementsInsertedAtTop[styleElementsInsertedAtTop.length - 1];
		if (options.insertAt === "top") {
			if(!lastStyleElementInsertedAtTop) {
				head.insertBefore(styleElement, head.firstChild);
			} else if(lastStyleElementInsertedAtTop.nextSibling) {
				head.insertBefore(styleElement, lastStyleElementInsertedAtTop.nextSibling);
			} else {
				head.appendChild(styleElement);
			}
			styleElementsInsertedAtTop.push(styleElement);
		} else if (options.insertAt === "bottom") {
			head.appendChild(styleElement);
		} else {
			throw new Error("Invalid value for parameter 'insertAt'. Must be 'top' or 'bottom'.");
		}
	}

	function removeStyleElement(styleElement) {
		styleElement.parentNode.removeChild(styleElement);
		var idx = styleElementsInsertedAtTop.indexOf(styleElement);
		if(idx >= 0) {
			styleElementsInsertedAtTop.splice(idx, 1);
		}
	}

	function createStyleElement(options) {
		var styleElement = document.createElement("style");
		styleElement.type = "text/css";
		insertStyleElement(options, styleElement);
		return styleElement;
	}

	function createLinkElement(options) {
		var linkElement = document.createElement("link");
		linkElement.rel = "stylesheet";
		insertStyleElement(options, linkElement);
		return linkElement;
	}

	function addStyle(obj, options) {
		var styleElement, update, remove;

		if (options.singleton) {
			var styleIndex = singletonCounter++;
			styleElement = singletonElement || (singletonElement = createStyleElement(options));
			update = applyToSingletonTag.bind(null, styleElement, styleIndex, false);
			remove = applyToSingletonTag.bind(null, styleElement, styleIndex, true);
		} else if(obj.sourceMap &&
			typeof URL === "function" &&
			typeof URL.createObjectURL === "function" &&
			typeof URL.revokeObjectURL === "function" &&
			typeof Blob === "function" &&
			typeof btoa === "function") {
			styleElement = createLinkElement(options);
			update = updateLink.bind(null, styleElement);
			remove = function() {
				removeStyleElement(styleElement);
				if(styleElement.href)
					URL.revokeObjectURL(styleElement.href);
			};
		} else {
			styleElement = createStyleElement(options);
			update = applyToTag.bind(null, styleElement);
			remove = function() {
				removeStyleElement(styleElement);
			};
		}

		update(obj);

		return function updateStyle(newObj) {
			if(newObj) {
				if(newObj.css === obj.css && newObj.media === obj.media && newObj.sourceMap === obj.sourceMap)
					return;
				update(obj = newObj);
			} else {
				remove();
			}
		};
	}

	var replaceText = (function () {
		var textStore = [];

		return function (index, replacement) {
			textStore[index] = replacement;
			return textStore.filter(Boolean).join('\n');
		};
	})();

	function applyToSingletonTag(styleElement, index, remove, obj) {
		var css = remove ? "" : obj.css;

		if (styleElement.styleSheet) {
			styleElement.styleSheet.cssText = replaceText(index, css);
		} else {
			var cssNode = document.createTextNode(css);
			var childNodes = styleElement.childNodes;
			if (childNodes[index]) styleElement.removeChild(childNodes[index]);
			if (childNodes.length) {
				styleElement.insertBefore(cssNode, childNodes[index]);
			} else {
				styleElement.appendChild(cssNode);
			}
		}
	}

	function applyToTag(styleElement, obj) {
		var css = obj.css;
		var media = obj.media;

		if(media) {
			styleElement.setAttribute("media", media)
		}

		if(styleElement.styleSheet) {
			styleElement.styleSheet.cssText = css;
		} else {
			while(styleElement.firstChild) {
				styleElement.removeChild(styleElement.firstChild);
			}
			styleElement.appendChild(document.createTextNode(css));
		}
	}

	function updateLink(linkElement, obj) {
		var css = obj.css;
		var sourceMap = obj.sourceMap;

		if(sourceMap) {
			// http://stackoverflow.com/a/26603875
			css += "\n/*# sourceMappingURL=data:application/json;base64," + btoa(unescape(encodeURIComponent(JSON.stringify(sourceMap)))) + " */";
		}

		var blob = new Blob([css], { type: "text/css" });

		var oldSrc = linkElement.href;

		linkElement.href = URL.createObjectURL(blob);

		if(oldSrc)
			URL.revokeObjectURL(oldSrc);
	}


/***/ },
/* 6 */
/***/ function(module, exports, __webpack_require__) {

	// style-loader: Adds some css to the DOM by adding a <style> tag

	// load the styles
	var content = __webpack_require__(1);
	if(typeof content === 'string') content = [[module.id, content, '']];
	// add the styles to the DOM
	var update = __webpack_require__(5)(content, {});
	if(content.locals) module.exports = content.locals;
	// Hot Module Replacement
	if(false) {
		// When the styles change, update the <style> tags
		if(!content.locals) {
			module.hot.accept("!!./../../../../../node_modules/css-loader/index.js!./../../../../../node_modules/less-loader/index.js!./tree.less", function() {
				var newContent = require("!!./../../../../../node_modules/css-loader/index.js!./../../../../../node_modules/less-loader/index.js!./tree.less");
				if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
				update(newContent);
			});
		}
		// When the module is disposed, remove the <style> tags
		module.hot.dispose(function() { update(); });
	}

/***/ },
/* 7 */
/***/ function(module, exports) {

	module.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA8AAAAOCAYAAADwikbvAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAOw4AAA7DAZBCL/sAAAAadEVYdFNvZnR3YXJlAFBhaW50Lk5FVCB2My41LjEwMPRyoQAAAHZJREFUOE9j+P//Pxiv7HH8jw/D1CFjOAOqACu4srcGqwFEa8ZmAEma0Q3A1JxbgIrRAH7NBABezT3bF8JxfU08WCwrJwlMgwD1bYbZguxXkBjIVmSboepwaCYA8GoGRQeyvy07S8AYBmhnMyGAUzOxGKLnPwMABE2jDc+cnOsAAAAASUVORK5CYII="

/***/ },
/* 8 */
/***/ function(module, exports) {

	module.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAOCAYAAAAmL5yKAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAadEVYdFNvZnR3YXJlAFBhaW50Lk5FVCB2My41LjEwMPRyoQAAAIhJREFUOE9j+P//P0UYqyApGIVzuN3kPzJGlsOF4Yy1xbpAChVAxVA0oGMwAbbt0TwEvtT8/8GuCjBGdxVYLVYDiAS4DUB2AR48PUoe0wCwX7EoxobRw4U6BszPUAEHHDEYrBbdALC/sCjGhrGGAUgQ7DQsGtAxVgNAGOQ0YjBMPQyjcEjH/xkAhEKsbVNNI1sAAAAASUVORK5CYII="

/***/ },
/* 9 */
/***/ function(module, exports) {

	module.exports = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAOCAYAAAAmL5yKAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAOw4AAA7DAZBCL/sAAAAadEVYdFNvZnR3YXJlAFBhaW50Lk5FVCB2My41LjEwMPRyoQAAAJBJREFUOE9j+P//P0UYqyApGIVzuN3kPzJGlsOF4Yy1xbpAChVAxVA0oGMwAbbt0TwEvtT8/8GuCjBGdxVYLVYDiAS4DUB2AR48PUoe0wCwX9EUgsSIwVgNQOdjxcBwghswP0MFLAAThLHxYbA6mBfA/oIJoinEhmGaUQwgRzPcABAGeYMYDFMPwygc0vF/BgDd66LkDQj2XgAAAABJRU5ErkJggg=="

/***/ }
/******/ ]);