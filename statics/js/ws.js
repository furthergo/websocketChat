let VIL = (function () {
    let VIL = {
    };

    function DefaultWebSocket(host, call) {
        let _host = host;
        let _isOpen = false;
        let _call = null;
        if("undefined" !== typeof call && call !== null){
            _call = call
        }else{
            _call = {
                onConnect:function (e) {
                    console.log("connect success ", e);
                },
                onDisconnect:function (e) {
                    console.log("disconnect ", e);
                },
                onMsg:function (data) {
                    console.log("receive message ", data)
                }
            }
        }

        let _socket = new WebSocket(_host);
        _socket.binaryType = "arraybuffer";

        /**
         * 发送消息
         * @param {string | ArrayBuffer } data
         * @constructor
         */
        this.send = function(data){
            if(_isOpen && _socket){
                _socket.send(data);
            }
        };

        this.setName = function(data) {
            data = "NAME:" + data;
            _socket.send(data);
        }

        this.close = function(){
            _socket.close(1000, "normal");
        };

        _socket.onopen = function(even){
            _isOpen = true;
            _call.onConnect(even);
        };

        _socket.onmessage = function(e){
            let data = e.data;
            _call.onMsg(data);
        };

        /**
         * 收到关闭连接
         * @param even
         */
        _socket.onclose = function(even){
            _isOpen = false;
            _call.onDisconnect({host:_host, event:even});
        };

        /**
         * 收到错误
         * @param err
         */
        _socket.onerror = function(err){
            _isOpen = false;
            _call.onDisconnect({host:_host, event:err});
        };
    }

    try{
        VIL.EngineSocket = DefaultWebSocket;
    }catch (e) {
        console.error("VILEngine error ", e);
    }

    return VIL;
})();
