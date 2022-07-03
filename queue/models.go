package queue

import (
	"time"
)

type AutoForwardMessage struct {
	FromUserId    string  `json:"fromUserId"`
	OnUserId      string  `json:"onUserId"`
	DocumentId    string  `json:"documentId"`
	FromUserInfo  []byte  `json:"fromUserInfo"`
	OnUserInfo    []byte  `json:"onUserInfo"`
	Message       *string `json:"message"`
	Type          *int    `json:"type"`
	SystemCreated *bool   `json:"systemCreated"`
	VisibleToAll  *bool   `json:"visibleToAll"`
}

type LogMessage struct {
	DocumentId string `json:"documentId"`
	UserId     string `json:"userId"`
	UserInfo   []byte `json:"userInfo"`
	Code       string `json:"code"`
	Data       []byte `json:"data"`
	LogText    string `json:"logText"`
}

type SearchIndex struct {
	DocumentId string `json:"documentId"`
}

type NotifyMessage struct {
	DocumentId       string `json:"documentId"`
	DocumentNumber   string `json:"documentNumber"`
	UserId           string `json:"userId"`
	UserInfo         []byte `json:"userInfo"`
	NotificationCode string `json:"notificationCode"`
}

type ExecuteAllApprovalsMessage struct {
	Id              string    `json:"id"`
	UserId          string    `json:"userId"`
	UserInfo        []byte    `json:"userInfo"`
	ExecuteUserId   string    `json:"executeUserId"`
	ExecuteUserInfo []byte    `json:"executeUserInfo"`
	ExecuteDate     time.Time `json:"executeDate"`
	ExecuteMessage  string    `json:"executeMessage"`
	DocumentId      string    `json:"documentId"`
}

type AddPointDateMessage struct {
	PerformerId    string    `json:"performerId"`
	ControlDate    time.Time `json:"controlDate"`
	ProactiveDate  time.Time `json:"proactiveDate"`
	CuratorComment string    `json:"curatorComment"`
}

type AddApprovalMessage struct {
	GroupId       string    `json:"groupId"`
	ParentId      *string   `json:"parentId"`
	DocumentId    string    `json:"documentId"`
	FromUserId    string    `json:"fromUserId"`
	FromUserInfo  []byte    `json:"fromUserInfo"`
	UserId        string    `json:"userId"`
	UserInfo      []byte    `json:"userInfo"`
	IsResponsible *bool     `json:"isResponsible""`
	ControlDate   time.Time `json:"controlDate"`
	Message       *string   `json:"message"`
	NeedResponse  *bool     `json:"needResponse"`
	IsLoop        *bool     `json:"isLoop"`
	Internal      *bool     `json:"internal"`
	EtcdLocalsKey []string  `json:"etcdLocalsKey"`
}

type SendTaskMessage struct {
	Id      string `json:"id"`
	Retries int    `json:"retries"`
}

type AccessIndex struct {
	DocumentId string `json:"documentId"`
}

type LinkedApprovalsOnAcceptMessage struct {
	OriginalDocumentId string `json:"originalDocumentId"`
	DocumentId         string `json:"documentId"`
}

type VotingOnAcceptMessage struct {
	OriginalDocumentId string `json:"originalDocumentId"`
	DocumentId         string `json:"documentId"`
	VoterOrgId         string `json:"voterOrgId"`
}

type UpdateEventDocInfoMessage struct {
	Id          string `json:"id" db:"document_id"`
	Number      string `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
}


type AcceptAllForSendMessage struct {
	DocumentId string                 `json:"documentId"`
	Token      map[string]interface{} `json:"token"`
}

type RemoveForSendData struct {
	DocumentId    string                 `json:"documentId"`
	SenderId      string                 `json:"senderId"`
	ReceiverId    string                 `json:"receiverId"`
	ReceiverLabel string                 `json:"receiverLabel"`
	Reason        string                 `json:"reason"`
	Token         map[string]interface{} `json:"token"`
}

type ExceptionMessage struct {
	ServiceName  string `json:"serviceName"`
	FileName     string `json:"fileName"`
	CallbackName string `json:"callbackName"`
	Line         int64  `json:"line"`
	Error        string `json:"error"`
	Genesis      string `json:"genesis"`
	ErrorType    string `json:"errorType"`
}
