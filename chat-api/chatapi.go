package getchat

import (
	"chat-api/db_info/query"
	chatapi "chat-api/gen/chatapi"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"goa.design/goa/v3/security"
)

// chatapi service example implementation.
// The example methods log the requests and return zero values.
type chatapisrvc struct {
	logger *log.Logger
}

// NewChatapi returns the chatapi service implementation.
func NewChatapi(logger *log.Logger) chatapi.Service {
	return &chatapisrvc{logger}
}

// APIKeyAuth implements the authorization logic for service "chatapi" for the
// "api_key" security scheme.
func (s *chatapisrvc) APIKeyAuth(ctx context.Context, key string, scheme *security.APIKeyScheme) (context.Context, error) {

	//簡易版、本番環境ではDBからの参照やアクセストークンを使用する方法
	if key != "apikey" {
		return ctx, fmt.Errorf("not implemented")
	}
	return ctx, nil
}

func (s *chatapisrvc) Getchat(ctx context.Context, p *chatapi.GetchatPayload) (res chatapi.GoaChatCollection, err error) {

	dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dbChtrm.Close()

	//リクエストされたルームIDのチャットルームがあるかの確認も兼ねて、ルーム情報をDBから取得する
	selectedChatroom := query.SelectChatroomById(p.ID, dbChtrm)

	Chats := query.SelectAllChatsById(selectedChatroom.Id, dbChtrm)
	fmt.Println(p.ID)
	fmt.Println("successed")
	s.logger.Print("chatAPI.get chat")

	return Chats, nil
}

func (s *chatapisrvc) Postchat(ctx context.Context, p *chatapi.PostchatPayload) (res bool, err error) {
	s.logger.Print("chatapi.postchat")

	dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dbChtrm.Close()

	//POSTされたチャットのルームIDのチャットルームがあるかの確認も兼ねて、ルーム情報をDBから取得する
	roomId, _ := strconv.Atoi(p.ID)
	currentChatroom := query.SelectChatroomById(roomId, dbChtrm)

	dbSession, err := sql.Open("mysql", query.ConStrSession)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dbSession.Close()

	//POSTされたチャットのcookieから、投稿者のユーザーIDを特定する
	cookie, _ := url.QueryUnescape(p.Cookie)
	postedUser := query.SelectSessionBySessionId(cookie, dbSession)

	postedChat := regexp.QuoteMeta(p.Chat)

	postDt := time.Now().UTC().Round(time.Second)

	//投稿者がルーム作成者と同じかどうかで、投稿者を変える
	if postedUser == currentChatroom.UserId {
		//投稿主と部屋作成者が同じ場合
		posted := query.InsertChat(roomId, currentChatroom.UserId, currentChatroom.RoomName, currentChatroom.Member, postedChat, postDt, dbChtrm)
		if posted {
			return true, nil
		}
	} else {
		//投稿主と部屋作成者が違う場合(=メンバーが投稿主の場合)
		posted := query.InsertChat(roomId, currentChatroom.Member, currentChatroom.RoomName, currentChatroom.UserId, postedChat, postDt, dbChtrm)
		if posted {
			return true, nil
		}
	}
	return false, err
}
