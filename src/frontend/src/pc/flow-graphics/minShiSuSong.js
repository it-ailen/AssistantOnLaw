/**
 * Created by hyku on 2016/12/17.
 */

"use strict";

var eventHandlerMap = {};

function eventHandlerRegister(event, func) {
    eventHandlerMap[event] = func;
}

function triggerEvent(event, data) {
    eventHandlerMap[event] && eventHandlerMap[event](data);
}

function draw(selector, handlerMap) {
  var C_W = 100,
      C_H = 30,
      C_V_GAP = 40,
      C_H_GAP = 20;

  var s = Snap(selector).attr({
    viewBox: "0 0 1000 1200"
  });

  // 划分为10列(0~9)
  function coordinate(row, col) {
    var x = col * (C_W + C_H_GAP) + C_W / 2;
    var y = row * (C_H + C_V_GAP) + C_H / 2;
    return {
      x: x,
      y: y
    };
  }

  var arrow = s.paper.path("M0,0 L0,2 L3,1 z").attr({
    fill: "#000",
    strokeUnits: "strokeWidth"
  }).marker(0, 0, 3, 2, 0, 1);

  function newStep(text, point, w, h) {
    w = w || C_W;
    h = h || C_H;
    var x = point.x,
        y = point.y;
    var rect = s.paper.rect(x - w / 2, y - h / 2, w, h);
    var textEle = s.paper.text(x, y, text).attr({
      fontSize: (h / 2),
      textAnchor: "middle",
      alignmentBaseline: "middle"
    });
    return s.paper.group(rect, textEle).attr({
      fill: "white",
      stroke: "black"
    });
  }

  function stepBox(ele) {
    return ele.getBBox();
  }

  function pathTo(startEle, endEle) {
    var startBox = stepBox(startEle),
        endBox = stepBox(endEle);
    var arrowLength = 5;
    var startPoint = {
          /* 起始点初始化为上顶点 */
          x: startBox.x + startBox.width / 2,
          y: startBox.y
        },
        endPoint = {
          x: endBox.x + endBox.width / 2,
          y: endBox.y - arrowLength
        };
    var direction = null;
    if (startBox.y + startBox.height < endBox.y) {/* startEle在endEle上方 */
      startPoint.y = startBox.y + startBox.height;
      direction = "down";
    } else if (startBox.y > endBox.y + endBox.height) { /* startEle在endEle下方 */
      endPoint.y = endBox.y + endBox.height;
      direction = "up";
    } else if (startBox.x < endBox.x) {/* startEle在endEle左方 */
      startPoint.x = startBox.x + startBox.width;
      startPoint.y = startBox.y + startBox.height / 2;
      endPoint.x = endBox.x;
      endPoint.y = endPoint.y + endBox.height / 2;
      direction = "right";
    } else {/* startEle在endEle右方 */
      startPoint.x = startBox.x;
      startPoint.y = startBox.y + startBox.height / 2;
      endPoint.x = endBox.x + endBox.width + arrowLength;
      endPoint.y = endBox.y + endBox.height / 2;
      direction = "left";
    }
    var paths = ["M" + startPoint.x + "," + startPoint.y];
    switch (direction) {
      case "down":
      case "up":
        var midY = (endPoint.y + startPoint.y) / 2;
        paths.push("V" + midY);
        paths.push("H" + endPoint.x);
        break;
      case "left":
      case "right":
        var midX = (endPoint.x + startPoint.x) / 2;
        paths.push("H" + midX);
        paths.push("V" + endPoint.y);
        break;
    }
    paths.push("L" + endPoint.x + "," + endPoint.y);
    return s.paper.path(paths.join(" ")).attr({
      markerEnd: arrow,
      fill: "none",
      stroke: "#000",
      strokeWidth: 2
    });
  }

  var s1 = newStep("民事诉讼", coordinate(0, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s1"
    })
  });

  var s2 = newStep("起诉", coordinate(1, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s2"
    })
  });
  pathTo(s1, s2);

  var s3 = newStep("诉前财产保全", coordinate(1, 6)).click(function (e) {
    triggerEvent("click", {
        id: "s3"
    })
  });
  pathTo(s1, s3);

  var s4 = newStep("受理", coordinate(2, 4)).click(function (e) {
    triggerEvent("click", {
        id: "s4"
    })
  });
  pathTo(s2, s4);
  var s5 = newStep("不予受理", coordinate(2, 7)).click(function (e) {
    triggerEvent("click", {
        id: "s5"
    })
  });
  pathTo(s2, s5);

  var s6 = newStep("驳回起诉", coordinate(3, 1)).click(function (e) {
    triggerEvent("click", {
        id: "s6"
    })
  });
  var s7 = newStep("撤诉", coordinate(3, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s7"
    })
  });
  var s8 = newStep("审理", coordinate(3, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s8"
    })
  });

  pathTo(s4, s6);
  pathTo(s4, s7);
  pathTo(s4, s8);

  var s9 = newStep("对不予受理的裁定提起上诉", coordinate(3, 7), 200).click(function (e) {
    triggerEvent("click", {
        id: "s9"
    })
  });
  pathTo(s5, s9);

  var s10 = newStep("回避", coordinate(4, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s10"
    })
  });
  var s11 = newStep("庭前准备", coordinate(4, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s11"
    });
  });
  pathTo(s11, s10);
  pathTo(s8, s11);

  var s12 = newStep("普通程序(6个月审结)", coordinate(5, 2), 200).click(function (e) {
    triggerEvent("click", {
        id: "s12"
    })
  });
  var s13 = newStep("简易程序(3个月审结)", coordinate(5, 6), 200).click(function (e) {
    triggerEvent("click", {
        id: "s13"
    })
  });

  pathTo(s11, s12);
  pathTo(s11, s13);

  var s14 = newStep("财产保全", coordinate(6, 1)).click(function (e) {
    triggerEvent("click", {
        id: "s4"
    })
  });
  var s15 = newStep("诉讼调解", coordinate(6, 2)).click(function (e) {
    triggerEvent("click", {
        id: "s15"
    })
  });
  var s16 = newStep("先予执行", coordinate(6, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s16"
    })
  });
  var s17 = newStep("缺席判决", coordinate(6, 4)).click(function (e) {
    triggerEvent("click", {
        id: "s17"
    })
  });

  var s18 = newStep("程序终结", coordinate(6, 7)).click(function (e) {
    triggerEvent("click", {
        id: "s18"
    })
  });
  pathTo(s12, s14);
  pathTo(s12, s15);
  pathTo(s12, s16);
  pathTo(s12, s17);
  pathTo(s13, s18);

  var s19 = newStep("达成协议", coordinate(7, 1)).click(function (e) {
    triggerEvent("click", {
        id: "s19"
    })
  });
  var s20 = newStep("未达成协议", coordinate(7, 4)).click(function (e) {
    triggerEvent("click", {
        id: "s20"
    })
  });
  pathTo(s15, s19);
  pathTo(s15, s20);

  var s21 = newStep("程序终结", coordinate(8, 1)).click(function (e) {
    triggerEvent("click", {
        id: "s21"
    })
  });
  var s22 = newStep("诉讼终结", coordinate(8, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s22"
    })
  });
  var s23 = newStep("延期审理", coordinate(8, 4)).click(function (e) {
    triggerEvent("click", {
        id: "s23"
    })
  });
  var s24 = newStep("诉讼中止", coordinate(8, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s24"
    })
  });
  var s25 = newStep("判决裁定", coordinate(8, 6)).click(function (e) {
    triggerEvent("click", {
        id: "s25"
    })
  });
  var s26 = newStep("撤诉", coordinate(8, 7)).click(function (e) {
    triggerEvent("click", {
        id: "s26"
    })
  });

  pathTo(s19, s21);
  pathTo(s20, s22);
  pathTo(s20, s23);
  pathTo(s20, s24);
  pathTo(s20, s25);
  pathTo(s20, s26);

  var s27 = newStep("程序终结", coordinate(9, 4)).click(function (e) {
    triggerEvent("click", {
        id: "s27"
    })
  });
  var s28 = newStep("上诉", coordinate(9, 6)).click(function (e) {
    triggerEvent("click", {
        id: "s28"
    })
  });

  pathTo(s25, s27);
  pathTo(s25, s28);
  var s29 = newStep("诉讼调解", coordinate(10, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s29"
    })
  });
  pathTo(s28, s29);

  var s30 = newStep("判决、裁定", coordinate(11, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s30"
    })
  });
  var s31 = newStep("达成协议", coordinate(11, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s31"
    })
  });

  pathTo(s29, s30);
  pathTo(s29, s31);

  var s32 = newStep("维持原判", coordinate(12, 2)).click(function (e) {
    triggerEvent("click", {
        id: "s32"
    })
  });
  var s33 = newStep("依法改判", coordinate(12, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s33"
    })
  });
  var s34 = newStep("发回重审", coordinate(12, 4)).click(function (e) {
    triggerEvent("click", {
        id: "s34"
    })
  });

  pathTo(s30, s32);
  pathTo(s30, s33);
  pathTo(s30, s34);

  var s35 = newStep("提审", coordinate(13, 1)).click(function (e) {
    triggerEvent("click", {
        id: "s35"
    })
  });
  var s36 = newStep("再审", coordinate(13, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s36"
    })
  });
  var s37 = newStep("抗诉", coordinate(13, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s37"
    })
  });
  var s38 = newStep("程序结束", coordinate(13, 7)).click(function (e) {
    triggerEvent("click", {
        id: "s38"
    })
  });

  pathTo(s33, s35);
  pathTo(s33, s36);
  pathTo(s33, s37);
  pathTo(s33, s38);

  var s39 = newStep("审理", coordinate(14, 2)).click(function (e) {
    triggerEvent("click", {
        id: "s39"
    })
  });
  var s40 = newStep("驳回申请", coordinate(14, 4)).click(function (e) {
    triggerEvent("click", {
        id: "s40"
    })
  });
  pathTo(s36, s39);
  pathTo(s36, s40);

  var s41 = newStep("一审程序", coordinate(15, 1)).click(function (e) {
    triggerEvent("click", {
        id: "s41"
    })
  });
  var s42 = newStep("二审程序", coordinate(15, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s42"
    })
  });
  var s43 = newStep("驳回起诉", coordinate(15, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s43"
    })
  });
  pathTo(s39, s41);
  pathTo(s39, s42);
  pathTo(s39, s43);

  var s44 = newStep("维持原判", coordinate(16, 1)).click(function (e) {
    triggerEvent("click", {
        id: "s44"
    })
  });
  var s45 = newStep("改判", coordinate(16, 3)).click(function (e) {
    triggerEvent("click", {
        id: "s45"
    })
  });
  var s46 = newStep("驳回起诉", coordinate(16, 5)).click(function (e) {
    triggerEvent("click", {
        id: "s46"
    })
  });
  var s47 = newStep("撤销原判 发回重审", coordinate(16, 7), 200).click(function (e) {
    triggerEvent("click", {
        id: "s47"
    })
  });

  pathTo(s41, s44);
  pathTo(s41, s45);
  pathTo(s41, s46);
  pathTo(s41, s47);
  pathTo(s42, s44);
  pathTo(s42, s45);
  pathTo(s42, s46);
  pathTo(s42, s47);
}

module.exports = {
    draw: draw,
    $on: eventHandlerRegister
};
