package template

var index = `<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <title>web压力测试</title>
</head>

<style>
    body {
        margin: 0;
        background-color: #f5f5f5;
    }

    .header {
        background-color: #ffffff;
        width: 100%;
        height: 60px;
        border-bottom: 0.5px solid #778899;
    }

    .header_ul {
        float: left;
        list-style: none;
    }

    .header_ul li {
        font-size: x-large;
        font-family: 方正舒体;
        color: #CD853F;
    }

    .configPanel {
        margin: 1em;
        width: auto;
        display: inline-block;
    }

    .log {
        margin: 1em;
        height: 400px;
        border: 0.5px solid #778899;
        flex-grow: 2;
        overflow: auto;

    }

    .content {
        display: flex;
        flex-direction: row;
        align-items: flex-start;
    }

    #bt_switch {
        margin: auto;
        margin-top: 10px;
        display: block;
        width: 100px;
    }

    #bt_quit {
        margin: auto;
        margin-top: 10px;
        display: block;
        width: 100px;
    }

    /*#bt_control:active{*/
    /*    background-color: red;*/
    /*}*/
    /*#bt_control:hover{*/
    /*      !*background-color: blue;*!*/
    /*  }*/
    .log_item {
        margin: 0px;
        font-family: "微软雅黑 Light";
    }
</style>
<body>
<div class="header">
    <ul class=" header_ul">
        <li class="header_ul_li">
            网站压力测试
        </li>
    </ul>
</div>
<div class="content">
    <div class="configPanel">
        <p style="margin: 0;">url:</p><input id="input_url"/>
        <p style="margin: 0;">interval:</p><input id="input_interval"/>
        <p style="margin: 0;">routine:</p><input id="input_routine"/>
        <button id="bt_switch">start</button>
        <button id="bt_quit">quit</button>
    </div>
    <div class="log" id="logContainer">
    <a href="http://papasen.com" target="_blank" title="papasen.com">欢迎来喷</a>
    </div>
</div>
</body>
<script>
    let isWork = false;
    let btSwitch = document.getElementById("bt_switch");
    let btQuit = document.getElementById("bt_quit");
    let inputUrl = document.getElementById("input_url");
    let inputInterval = document.getElementById("input_interval");
    let inputRoutine = document.getElementById("input_routine");
    let logContainer = document.getElementById("logContainer");
    btSwitch.onclick = onBtSwitch;
    btQuit.onclick = onBtQuit;

    function onBtQuit() {
        let data = {
            "purpose": "quit",
            "url": "",
            "interval": 0,
            "goroutine": 0
        };
        let ajax = new XMLHttpRequest();
        ajax.open("POST", "/work");
        ajax.send(JSON.stringify(data));
        ajax.onreadystatechange = onServerHandle(ajax);

    }

    function onBtSwitch() {
        if (isWork) {
            //正在工作就将其停止
            stop();
            btSwitch.innerText = "start";
            isWork = !isWork;

        } else {
            //开启
            start();
            btSwitch.innerText = "stop";
            isWork = !isWork;
        }
    }

    function start() {
        console.log("start");
        //以百度为例测试
        let data = {
            "purpose": "start",
            "url": inputUrl.value,
            "interval": parseInt(inputInterval.value),
            "goroutine": parseInt(inputRoutine.value)
        };
        let ajax = new XMLHttpRequest();
        ajax.open("POST", "/work");
        ajax.send(JSON.stringify(data));
        ajax.onreadystatechange = onServerHandle(ajax);
    }

    function stop() {
        console.log("stop");
        let data = {
            "purpose": "stop",
            "url": "",
            "interval": 0,
            "goroutine": 0
        };
        let ajax = new XMLHttpRequest();
        ajax.open("POST", "/work");
        ajax.send(JSON.stringify(data));
        ajax.onreadystatechange = onServerHandle(ajax);
    }

    function onServerHandle(ajax) {
        return () => {
            if (ajax.readyState == 4 && ajax.status == 200) {
                console.log(ajax.responseText);
                print(ajax.responseText);

            }
        }
    }


    function print(contentStr) {
        logContainer.innerHTML += " <p class=\"log_item\">" + contentStr + "</p>";
        logContainer.scrollTop = logContainer.scrollHeight;
    }

</script>
</html>`

func GetIndexHtml() string {
	return index
}
