package main

const (
	indexHtml = `<html>
<head>
  <title>Slide controller</title>
  <meta charset="UTF-8">
  <meta name="viewport"
        content="width=device-width; initial-scale=1.0; maximum-scale=1.0; user-scalable=0;">
  <meta name="theme-color" content="#333333">
</head>

<style>
  body {
    background: white;
    color: rgba(0, 0, 0, .8);
    padding: 0;
    margin: 0;
  }

  #controls {
    position: fixed;
    top: 0;
    width: 100%;
    background: white;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.2);
  }

  #left, #right {
    display: table;
    font-size: 72px;
    width: 50%;
    overflow: hidden;
    cursor: pointer;
    transition: color 0.08s, background 0.08s ease-in-out;
    -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  #left {
    float: left;
  }

  #right {
    float: right;
  }

  #right:active,
  #left:active {
    background: #333;
    color: white;
  }

  #prompter {
    overflow-y: scroll;
    font-family: CustomSerif, Georgia, Cambria, 'Times New Roman', serif;
    font-weight: 400;
    font-style: normal;
    font-size: 18px;
    line-height: 1.58;
    box-sizing: border-box;
    padding: 80px 12px 0;
  }

  .switch {
    display: table-cell;
    vertical-align: middle;
    text-align: center;
  }

  @media screen and (min-width: 768px) {
    #prompter {
      max-width: 736px;
      margin: 0 auto;
    }
  }

  @media screen and (max-width: 768px) {
    #prompter {
      width: calc(100% - 24px);
      margin: 0;
    }
  }

</style>

<body>
<div id="container">
  <div id="controls">
    <div id="left">
      <div class="switch">
        <
      </div>
    </div>
    <div id="right">
      <div class="switch">
        >
      </div>
    </div>
  </div>
  <div id="prompter">
  </div>
</div>

<script>
  var url = "ws://" + window.location.host + "/presento";
  var ws = new WebSocket(url);
  var left = document.getElementById("left");
  var right = document.getElementById("right");
  var prompter = document.getElementById("prompter");

  ws.onmessage = function (msg) {
    prompter.innerHTML = msg.data;
  };

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
