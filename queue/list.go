package queue

const (
	QueueCreateAutoForward              = "createAutoForward"
	QueueCreateLog                      = "createLog"
	QueueCreateSystemLog                = "createSystemLog"
	QueueSendDocument                   = "sendDocument"
	QueueBuildReport                    = "buildReport"
	QueueNotifyUser                     = "notifyUser"
	QueueExecuteAllApprovals            = "executeAllApprovals"
	QueueIndexDocument                  = "indexDocument"
	QueueRemoveFromIndex                = "removeFromIndex"
	QueueAddPointDate                   = "addPointDate"
	QueueAddApproval                    = "addApproval"
	QueueIndexAccess                    = "indexAccess"
	QueueProcessLinkedApprovalsOnAccept = "linkedApprovalsOnAccept"
	QueueProcessVotingOnAccept          = "votingOnAccept"
	QueueUpdateEventDocInfo             = "updateEventDocInfo"
	QueueForSend                        = "forSend"
	QueueAcceptAllSends                 = "acceptAllSends"
	QueueRemoveFromForSend              = "removeForSend"

	QueueException = "exception"
)
