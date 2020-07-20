package Gamecontrollers

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/fv0008/AWS_Russia/server"
	"time"
)


type GroupChatChan struct{
	// Channel for new join users.
	JoinGroupChatChan chan SocketId
	// Channel for exit users.
	LeaveGroupChatChan chan UnSocketId
	// Send events here to publish them.
	GroupMsgList chan(server.IM_protocol)
	// Long polling waiting list.
	ActiveGroupChatChan *list.List
}



func (this *GroupChatChan)NewGroupChatMsg(ep server.EventType, user server.IM_protocol_user,Socketid uint32, msg string) server.IM_protocol {
	return server.IM_protocol{ep, msg,Socketid,user, int(time.Now().Unix())}
}


func (this *GroupChatChan)IsInGroupChatChinById(SocketId uint32) (bool) {
	for sub := this.ActiveGroupChatChan.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(SocketInfo).SocketId == SocketId {
			return true
		}
	}
	return false
}

func (this *GroupChatChan) GroupChatChanAcitve() {
	for {
		select {
		case JoinChan := <-this.JoinGroupChatChan:
			if !this.IsInGroupChatChinById(JoinChan.SocketId) {
				this.ActiveGroupChatChan.PushBack(JoinChan) // Add user to the end of list.
				beego.Info("New User socket:", JoinChan.SocketId)
			} else {
				beego.Info("Old User socket:", JoinChan.SocketId )
			}
		case Message := <-this.GroupMsgList:
			//如果是心跳，单发
			switch Message.Type {
			case
				server.IM_EVENT_HEART,
				server.IM_EVENT_JOIN,
				server.IM_EVENT_LEAVE,
				server.IM_EVENT_MESSAGE:
				//this.HeartWebSocket(event)
				beego.Info("ChatChinMsg from", Message.Users.From, ";Msg:", Message.Msg)
				break
			case
				server.IM_EVENT_BROADCAST_HEART,
				server.IM_EVENT_BROADCAST_JOIN,
				server.IM_EVENT_BROADCAST_LEAVE,
				server.IM_EVENT_BROADCAST_MESSAGE:
				//this.broadcastWebSocket(event)
				beego.Info("ChatChinBCMsg from", Message.Users.From, ";Msg:", Message.Msg)
				break
			}
		case LeaveChan := <-this.LeaveGroupChatChan:
			for sub := this.ActiveGroupChatChan.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(SocketInfo).SocketId == LeaveChan.SocketId {
					this.ActiveGroupChatChan.Remove(sub)
					beego.Error("ChatUser Leave:", LeaveChan)
					break
				}
			}
		}
	}
}




func (this *GroupChatChan)init() {
	// Channel for new join users.
	this.JoinGroupChatChan = make(chan SocketId, 10)
	this.LeaveGroupChatChan = make(chan UnSocketId,10)
	// Send events here to publish them.
	this.GroupMsgList = make(chan server.IM_protocol, 10)
	// Long polling waiting list.
	this.ActiveGroupChatChan = list.New()

	go this.GroupChatChanAcitve()



}



