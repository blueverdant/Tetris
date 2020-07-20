var socket;
var heart
var reconnectcount = 0
var timeouthandle
var SocketId

function reconnect(){
    clearInterval(heart)
    clearTimeout(timeouthandle)
    reconnectcount++
    if(reconnectcount>10)
    {
        alert("网络异常")
        reconnectcount = 1
    }
    socket = new WebSocket('ws://' + window.location.host + '/IM/ws/socket?Socket='+SocketId);
    config(socket)
}

function heartbeat(){
    heart = setInterval(function() {
        const heart =  {"Type":6,"SocketId":SocketId,"Timestamp": new Date().getTime()}
        //转换数据类型
        var sendinfo =JSON.stringify(heart);
        socket && socket.send(sendinfo)
    }, 1000*20);
}

function config(ret)
{
    ret.onopen =function(event)
    {
        //SocketId = data.SocketId
        reconnectcount = 0
    }
    ret.onclose = function (event) {
        console.log("onclose",event)
        timeouthandle = setTimeout(reconnect,reconnectcount*10000)
        clearInterval(heart)
    }
    ret.onerror = function (event)
    {
        console.log("onerror",event)
        setTimeout(reconnect,reconnectcount*10000)
        clearInterval(heart)
    }
    // Message received on the socket
    ret.onmessage = function (event) {
        var data = JSON.parse(event.data);
        //var li = document.createElement('li');

        console.log(data)
        switch (data.Type) {
            case 0: // JOIN
                SocketId = data.SocketId
                var username = document.getElementById('uname');

                username.innerText = data.SocketId;
                console.log("Join updateSocketId:",data);
                break
            case 1: // BCJOIN
                //li.innerText = data.SocketId + ' joined the chat room.';
                break;
            case 2: // LEAVE
                var username = document.getElementById('uname');
                if (SocketId ==data.SocketId){
                    username.innerText = "";
                }

            case 3: // BCLEAVE

                break;
            case 4: // MESSAGE
                var Notify = document.getElementById('notify');

                Notify.innerText = data.Msg;
                break;
            case 5: // BCMESSAGE

                var action = document.getElementById('action');

                    action.innerText = data.Msg;



                //li.appendChild(username);
                //li.appendChild(document.createTextNode(': '));
                //li.appendChild(content);

                break;
            case 6: // HEART
            case 7: // BCHEART
                return;
                break;
        }

        //$('#chatbox li').first().before(li);
    };
    heartbeat()
}
function postConnect(action) {

    //var content = $('#sendbox').val();
    var info =  {"Msg":action,"Type":5,}
    info.Timestamp = new Date().getTime()
    info.SocketId=SocketId
    info.Users={"To":0, "From":0,"SessKey":0, "ChatChanId":0}
    var sendinfo =JSON.stringify(info);
    socket.send(sendinfo);
    //$('#sendbox').val('');
}
$(document).ready(function () {

    // Create a socket
    reconnect()



    // Send messages.


    $('#sendbtn').click(function () {
        postConnect();
    });


});
