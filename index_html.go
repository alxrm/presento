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
  }

  #left, #right {
    position: absolute;
    display: table;
    font-size: 4rem;
    top: 0;
    width: 50%;
    overflow: hidden;
    cursor: pointer;
    transition: color 0.1s, background 0.1s ease-in-out;
    -webkit-tap-highlight-color: rgba(0, 0, 0, 0);
    -webkit-user-select: none;
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
  }

  #left {
    left: 0;
  }

  #right {
    left: 50%;
  }

  #right:active,
  #left:active {
    background: rgba(0, 0, 0, .8);
    color: white;
  }

  #prompter {
    position: absolute;
    top: 6em;
    bottom: 0;
    overflow-y: scroll;
    font-family: CustomSerif, Georgia, Cambria, 'Times New Roman', serif;
    font-weight: 400;
    font-style: normal;
    font-size: 18px;
    line-height: 1.58;
    box-sizing: border-box;
  }

  .switch {
    display: table-cell;
    vertical-align: middle;
    text-align: center;
  }

  @media screen and (min-width: 768px) {
    body {
      padding: 16px;
    }

    #prompter {
      left: 50%;
      margin-left: -360px;
      max-width: 720px;
    }
  }

  @media screen and (max-width: 768px) {
    body {
      padding: 0;
    }

    #prompter {
      width: calc(100% - 12px);
      padding: 0 6px;
      margin: 0;
    }
  }

</style>

<body>
<div id="container">
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
