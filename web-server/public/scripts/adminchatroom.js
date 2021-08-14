//取得したチャットは各関数で使用するため、グローバルで定義しておく
let allchats ="";

window.onload = async function() {
  
  //APIで使用するルームIDは、URLから取得
  let url = location.href;
  let roomid = url.replace("http://172.25.0.2/admin/chatrooms/chatroom","");

  const urlForApiGet = "http://172.25.0.3:8000/chatroom/" + roomid;

  const axiosConfig = {
    headers: {
      "Authorization": "apikey",
    }
  };

  async function getChatFromApi() {
    try {
       res = await axios.get(urlForApiGet, axiosConfig);

       allchats = res.data;

      for (const chat of allchats) {
        let textUserNode = document.createTextNode(chat.UserId);
        let textPostDtNode = document.createTextNode(chat.PostDt);
        let textChatUnescaped = unescape(chat.Chat).replace(/\\/g, "");
        let textChatNode = document.createTextNode(textChatUnescaped);

        let elUser = document.createElement("div");
        elUser.appendChild(textUserNode);
        elUser.id ="user";
        elUser.style = "display: inline-block; _display: inline;";
        
        let elPostDt = document.createElement("div");
        elPostDt.appendChild(textPostDtNode);
        elPostDt.id ="postdt";
        elPostDt.style = "display: inline-block; _display: inline;";
        
        let elChat = document.createElement("div");
        elChat.appendChild(textChatNode);
        elChat.id ="chatText";

        let newLi = document.createElement("li");
        newLi.appendChild(elUser);
        newLi.appendChild(elPostDt);
        newLi.appendChild(elChat);
        let chatList = document.getElementById("chats");
        chatList.appendChild(newLi);
      }

      let roomNameUnescaped = allchats[0].RoomName.replace(/\\/g, "");
      let roomnameText = document.createTextNode(roomNameUnescaped);
      let newH2 = document.createElement("h2");
      newH2.appendChild(roomnameText);
      let roomname = document.getElementById("roomname-header");
      roomname.appendChild(newH2);

      return;

    } catch(error){
      const {
        status,
        statusText
      } = error.response;
      console.log(`Error! HTTP Status: ${status} ${statusText}`);
    };
  };

  await getChatFromApi();

  let element = document.documentElement;
  let bottom = element.scrollHeight - element.clientHeight;
  window.scroll(0, bottom);
};

function deleteChtrmFunc(){

	if(window.confirm("本当にこのルームを削除しますか？")){
    this.form.submit();
	} else {
    window.alert("ルーム削除をキャンセルしました");
    return false
  };
};