package main

const (
	indexHtml = `<html>
<head>
  <title>Slide controller</title>
  <meta name="viewport"
        content="width=device-width; initial-scale=1.0; maximum-scale=1.0; user-scalable=0;">
</head>

<style>
  body {
    background: #2b2b2b;
    color: #efefef;
    padding: 16px;
  }

  #left, #right {
    position: absolute;
    display: table;
    font-size: 20rem;
    top: 0;
    height: 100%;
    width: 50%;
    overflow: hidden;
    cursor: pointer;
  }

  #left {
    left: 0;
  }

  #right {
    left: 50%;
  }

  #right:active,
  #left:active {
    background: #efefef;
    color: #2b2b2b;
  }

  .text {
    display: table-cell;
    vertical-align: middle;
    text-align: center;
  }

</style>

<body>
<div id="container">
  <div id="left">
    <div class="text"> <</div>
  </div>
  <div id="right">
    <div class="text"> ></div>
  </div>
</div>

<script>
  var url = "ws://" + window.location.host + "/presento";
  var ws = new WebSocket(url);
  var left = document.getElementById("left");
  var right = document.getElementById("right");

  left.onclick = function () {
    ws.send("left");
  };

  right.onclick = function () {
    ws.send("right");
  };

</script>
</body>
</html>`
)
