package Gamecontrollers

import (
	"container/list"
	"time"
	"github.com/fv0008/AWS_Russia/server"

	"github.com/astaxie/beego"
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



func (this *GroupChatChan)NewGroupChatMsg(ep lunarhookmodel.EventType, user lunarhookmodel.IM_protocol_user,Socketid uint32, msg string) lunarhookmodel.IM_protocol {
	return ImArchive.IM_protocol{ep, msg,Socketid,user, int(time.Now().Unix())}
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
				lunarhookmodel.IM_EVENT_HEART,
				lunarhookmodel.IM_EVENT_JOIN,
				lunarhookmodel.IM_EVENT_LEAVE,
				lunarhookmodel.IM_EVENT_MESSAGE:
				//this.HeartWebSocket(event)
				beego.Info("ChatChinMsg from", Message.Users.From, ";Msg:", Message.Msg)
				break
			case
				lunarhookmodel.IM_EVENT_BROADCAST_HEART,
				lunarhookmodel.IM_EVENT_BROADCAST_JOIN,
				lunarhookmodel.IM_EVENT_BROADCAST_LEAVE,
				lunarhookmodel.IM_EVENT_BROADCAST_MESSAGE:
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
	this.GroupMsgList = make(chan lunarhookmodel.IM_protocol, 10)
	// Long polling waiting list.
	this.ActiveGroupChatChan = list.New()

	go this.GroupChatChanAcitve()



}



