
<script>
    $(function(){
        var socket = null;
        var msgBox = $("#chatbox textarea");
        var messages = $("#messages");
        $("#chatbox").submit(function(){
            if (!msgBox.val()) return false;
            if (!socket) {
                alert("오류: 소켓 연결이 없습니다.");
                return false;
            }
            socket.send(msgBox.val());
            msgBox.val("");
            return false;
        });
        if ( !window["WebSocket"]){
            alert("오류: 브로우저가 웹 소켓을 지원하지 않습니다.");
        } else {
            socket = new WebSocket("ws://localhost:8080/room");
            socket.onclose = function() {
                alert("연결이 종료되었습니다.");
            }
            socket.onmessage = function(e) {
                messages.append($("<li>").text(e.data));
            }
        }
    });
</script>