(function() {
  window.small_world = function(pts) {
    var ret = [];
    for (var i = 0; i < pts.length; i++) {
      var cl3 = [];
      for (var j = 0; j < pts.length; j++) {
        if (i === j) continue;
        cl3.push([j, dist_sq(pts[i], pts[j])]);
        cl3.sort(function(a,b) { return numcmp(b[1], a[1]); });
        cl3.splice(3);
      }
      ret.push([cl3[0][0], cl3[1][0], cl3[2][0]]);
    }
    return ret;
  };

  function dist_sq(p1, p2) {
    return Math.pow(p1[0] - p2[0], 2) + Math.pow(p1[1] - p2[1], 2);
  }
})();
