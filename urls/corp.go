package urls

// oauth
const QRCodeAuthorize = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect"

const (
	CorpCgiBinAccessToken  = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	CorpCgiBinAPIDomainIP  = "https://qyapi.weixin.qq.com/cgi-bin/get_api_domain_ip"
	CorpCgiBinCallbackIP   = "https://qyapi.weixin.qq.com/cgi-bin/getcallbackip"
	CorpCgiBinUserInfo     = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo"
	CorpCgiBinUserAuthSucc = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc"
	CorpCginBinCallbackIP  = "https://qyapi.weixin.qq.com/cgi-bin/getcallbackip"
)

// user
const (
	CorpUserCreate               = "https://qyapi.weixin.qq.com/cgi-bin/user/create"
	CorpUserGet                  = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	CorpUserUpdate               = "https://qyapi.weixin.qq.com/cgi-bin/user/update"
	CorpUserDelete               = "https://qyapi.weixin.qq.com/cgi-bin/user/delete"
	CorpUserBatchDelete          = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete"
	CorpUserSimpleList           = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist"
	CorpUserList                 = "https://qyapi.weixin.qq.com/cgi-bin/user/list"
	CorpUserConvertToOpenID      = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid"
	CorpUserConvertToUserID      = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid"
	CorpUserBatchInvite          = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite"
	CorpUserJoinQRCode           = "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode"
	CorpUserActiveStat           = "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat"
	CorpUserGetUserID            = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserid"
	CorpUserDepartmentCreate     = "https://qyapi.weixin.qq.com/cgi-bin/department/create"
	CorpUserDepartmentUpdate     = "https://qyapi.weixin.qq.com/cgi-bin/department/update"
	CorpUserDepartmentDelete     = "https://qyapi.weixin.qq.com/cgi-bin/department/delete"
	CorpUserDepartmentList       = "https://qyapi.weixin.qq.com/cgi-bin/department/list"
	CorpUserDepartmentSimpleList = "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist"
	CorpUserDepartmentGet        = "https://qyapi.weixin.qq.com/cgi-bin/department/get"
	CorpUserTagCreate            = "https://qyapi.weixin.qq.com/cgi-bin/tag/create"
	CorpUserTagUpdate            = "https://qyapi.weixin.qq.com/cgi-bin/tag/update"
	CorpUserTagDelete            = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete"
	CorpUserTagList              = "https://qyapi.weixin.qq.com/cgi-bin/tag/list"
	CorpUserTagGetUser           = "https://qyapi.weixin.qq.com/cgi-bin/tag/get"
	CorpUserTagAddUser           = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers"
	CorpUserTagDeleteUser        = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers"
	CorpUserBatchSyncUser        = "https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser"
	CorpUserBatchReplaceUser     = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser"
	CorpUserBatchReplaceParty    = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty"
	CorpUserGetBatchResult       = "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult"
	CorpUserExportSimpleUser     = "https://qyapi.weixin.qq.com/cgi-bin/export/simple_user"
	CorpUserExportUser           = "https://qyapi.weixin.qq.com/cgi-bin/export/user"
	CorpUserExportDepartment     = "https://qyapi.weixin.qq.com/cgi-bin/export/department"
	CorpUserExportTagUser        = "https://qyapi.weixin.qq.com/cgi-bin/export/taguser"
	CorpUserGetExportResult      = "https://qyapi.weixin.qq.com/cgi-bin/export/get_result"
)

// linkedcorp
const (
	CorpLinkedcorpPermList       = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list"
	CorpLinkedcorpUserGet        = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/get"
	CorpLinkedcorpUserSimpleList = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/simplelist"
	CorpLinkedcorpUserList       = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/list"
	CorpLinkedcorpDepartmentList = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/department/list"
)

// externalcontact
const (
	CorpExternalContactFollowUserList             = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list"
	CorpExternalContactWayAdd                     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way"
	CorpExternalContactWayUpdate                  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way"
	CorpExternalContactWayGet                     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way"
	CorpExternalContactWayList                    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list_contact_way"
	CorpExternalContactWayDelete                  = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way"
	CorpExternalContactCloseTempChat              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/close_temp_chat"
	CorpExternalContactList                       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	CorpExternalContactGet                        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	CorpExternalContactBatchGetByUser             = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user"
	CorpExternalContactRemark                     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark"
	CorpExternalContactCustomerStrategyCreate     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/create"
	CorpExternalContactCustomerStrategyEdit       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/edit"
	CorpExternalContactCustomerStrategyList       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/list"
	CorpExternalContactCustomerStrategyGet        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get"
	CorpExternalContactCustomerStrategyGetRange   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/get_range"
	CorpExternalContactCustomerStrategyDelete     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/customer_strategy/del"
	CorpExternalContactCorpTagList                = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list"
	CorpExternalContactCorpTagAdd                 = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag"
	CorpExternalContactCorpTagEdit                = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag"
	CorpExternalContactCorpTagDelete              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag"
	CorpExternalContactStrategyTagList            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_strategy_tag_list"
	CorpExternalContactStrategyTagAdd             = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_strategy_tag"
	CorpExternalContactStrategyTagEdit            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_strategy_tag"
	CorpExternalContactStrategyTagDelete          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_strategy_tag"
	CorpExternalContactMarkTag                    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag"
	CorpExternalContactTransferCustomer           = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_customer"
	CorpExternalContactTransferResult             = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_result"
	CorpExternalContactGetUnassignedList          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list"
	CorpExternalContactTransferResignedCustomer   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_customer"
	CorpExternalContactResignedTransferResult     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_result"
	CorpExternalContactGroupChatTranster          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer"
	CorpExternalContactGroupChatList              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list"
	CorpExternalContactGroupChatGet               = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get"
	CorpExternalContactOpenGIDToChatID            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/opengid_to_chatid"
	CorpExternalContactAddMomentTask              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_moment_task"
	CorpExternalContactGetMomentTaskResult        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_task_result"
	CorpExternalContactGetMomentList              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_list"
	CorpExternalContactGetMomentTask              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_task"
	CorpExternalContactGetMomentCustomerList      = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_customer_list"
	CorpExternalContactGetMomentSentResult        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_send_result"
	CorpExternalContactGetMomentComments          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_moment_comments"
	CorpExternalContactMomentStrategyList         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/list"
	CorpExternalContactMomentStrategyGet          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get"
	CorpExternalContactMomentStrategyGetRange     = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/get_range"
	CorpExternalContactMomentStrategyCreate       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/create"
	CorpExternalContactMomentStrategyEdit         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/edit"
	CorpExternalContactMomentStrategyDelete       = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/moment_strategy/del"
	CorpExternalContactAddMsgTemplate             = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template"
	CorpExternalContactGetGroupMsgList            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list_v2"
	CorpExternalContactGetGroupMsgTask            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task"
	CorpExternalContactGetGroupMsgSendResult      = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result"
	CorpExternalContactSendWelcomeMsg             = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg"
	CorpExternalContactGroupWelcomeTemplateAdd    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add"
	CorpExternalContactGroupWelcomeTemplateEdit   = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit"
	CorpExternalContactGroupWelcomeTemplateGet    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get"
	CorpExternalContactGroupWelcomeTemplateDelete = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del"
	CorpExternalContactGetUserBehaviorData        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data"
	CorpExternalContactGroupChatStatistic         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic"
	CorpExternalContactGroupChatStatisticByDay    = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day"
	CorpExternalContactProductAlbumAdd            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_product_album"
	CorpExternalContactProductAlbumUpdate         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_product_album"
	CorpExternalContactProductAlbumGet            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_product_album"
	CorpExternalContactProductAlbumList           = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_product_album_list"
	CorpExternalContactProductAlbumDelete         = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/delete_product_album"
	CorpExternalContactInterceptRuleAdd           = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_intercept_rule"
	CorpExternalContactInterceptRuleUpdate        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_intercept_rule"
	CorpExternalContactInterceptRuleGet           = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_intercept_rule"
	CorpExternalContactInterceptRuleList          = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_intercept_rule_list"
	CorpExternalContactInterceptRuleDelete        = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_intercept_rule"
	CorpExternalContactConvertToOpenID            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/convert_to_openid"
	CorpExternalContactUploadAttachment           = "https://qyapi.weixin.qq.com/cgi-bin/media/upload_attachment"
)

// kf
const (
	CorpKFAccountAdd              = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/add"
	CorpKFAccountDelete           = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/del"
	CorpKFAccountUpdate           = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/update"
	CorpKFAccountList             = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/list"
	CorpKFAddContactWay           = "https://qyapi.weixin.qq.com/cgi-bin/kf/add_contact_way"
	CorpKFServicerAdd             = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/add"
	CorpKFServicerDelete          = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/del"
	CorpKFServicerList            = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/list"
	CorpKFServiceStateGet         = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/get"
	CorpKFServiceStateTransfer    = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/trans"
	CorpKFSyncMsg                 = "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg"
	CorpKFSendMsg                 = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg"
	CorpKFSendMsgOnEvent          = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg_on_event"
	CorpKFCustomerBatchGet        = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/batchget"
	CorpKFGetUpgradeServiceConfig = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/get_upgrade_service_config"
	CorpKFUpgradeService          = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/upgrade_service"
	CorpKFCancelUpgradeService    = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/cancel_upgrade_service"
)

// agent
const (
	CorpAgentGet             = "https://qyapi.weixin.qq.com/cgi-bin/agent/get"
	CorpAgentList            = "https://qyapi.weixin.qq.com/cgi-bin/agent/list"
	CorpAgentSet             = "https://qyapi.weixin.qq.com/cgi-bin/agent/set"
	CorpMenuCreate           = "https://qyapi.weixin.qq.com/cgi-bin/menu/create"
	CorpMenuGet              = "https://qyapi.weixin.qq.com/cgi-bin/menu/get"
	CorpMenuDelete           = "https://qyapi.weixin.qq.com/cgi-bin/menu/delete"
	CorpSetWorkbenchTemplate = "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template"
	CorpGetWorkbenchTemplate = "https://qyapi.weixin.qq.com/cgi-bin/agent/get_workbench_template"
	CorpSetWorkbenchData     = "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data"
)

// message
const (
	CorpMessageSend                = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
	CorpMessageUpdateTemplateCard  = "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card"
	CorpMessageRecall              = "https://qyapi.weixin.qq.com/cgi-bin/message/recall"
	CorpAppchatCreate              = "https://qyapi.weixin.qq.com/cgi-bin/appchat/create"
	CorpAppchatUpdate              = "https://qyapi.weixin.qq.com/cgi-bin/appchat/update"
	CorpAppchatGet                 = "https://qyapi.weixin.qq.com/cgi-bin/appchat/get"
	CorpAppchatSend                = "https://qyapi.weixin.qq.com/cgi-bin/appchat/send"
	CorpLinkedcorpMessageSend      = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send"
	CorpExternalContactMessageSend = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/message/send"
)

// media
const (
	CorpMediaUpload    = "https://qyapi.weixin.qq.com/cgi-bin/media/upload"
	CorpMediaUploadImg = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg"
	CorpMediaGet       = "https://qyapi.weixin.qq.com/cgi-bin/media/get"
	CorpMediaGetJSSDK  = "https://qyapi.weixin.qq.com/cgi-bin/media/get/jssdk"
)

// OA
const (
	CorpOAGetCorpCheckinOption                 = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcorpcheckinoption"
	CorpOAGetCheckinOption                     = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinoption"
	CorpOAGetCheckinData                       = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckindata"
	CorpOAGetCheckinDayData                    = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_daydata"
	CorpOAGetCheckinMonthData                  = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckin_monthdata"
	CorpOAGetCheckinScheduleList               = "https://qyapi.weixin.qq.com/cgi-bin/checkin/getcheckinschedulist"
	CorpOASetCheckinScheduleList               = "https://qyapi.weixin.qq.com/cgi-bin/checkin/setcheckinschedulist"
	CorpOAAddCheckinUserFace                   = "https://qyapi.weixin.qq.com/cgi-bin/checkin/addcheckinuserface"
	CorpOAGetHardwareCheckinData               = "https://qyapi.weixin.qq.com/cgi-bin/hardware/get_hardware_checkin_data"
	CorpOAGetTemplateDetail                    = "https://qyapi.weixin.qq.com/cgi-bin/oa/gettemplatedetail"
	CorpOAApplyEvent                           = "https://qyapi.weixin.qq.com/cgi-bin/oa/applyevent"
	CorpOAGetApprovalInfo                      = "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovalinfo"
	CorpOAGetApprovalDetail                    = "https://qyapi.weixin.qq.com/cgi-bin/oa/getapprovaldetail"
	CorpOAGetVacationCorpConf                  = "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getcorpconf"
	CorpOAGetUserVacationQuota                 = "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/getuservacationquota"
	CorpOASetOneUserVacationQuota              = "https://qyapi.weixin.qq.com/cgi-bin/oa/vacation/setoneuserquota"
	CorpOAGetJournalRecordList                 = "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_list"
	CorpOAGetJournalRecordDetail               = "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_record_detail"
	CorpOAGetJournalStatList                   = "https://qyapi.weixin.qq.com/cgi-bin/oa/journal/get_stat_list"
	CorpOAGetOpenApprovalData                  = "https://qyapi.weixin.qq.com/cgi-bin/corp/getopenapprovaldata"
	CorpOAMeetingRoomAdd                       = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/add"
	CorpOAMeetingRoomList                      = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/list"
	CorpOAMeetingRoomEdit                      = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/edit"
	CorpOAMeetingRoomDelete                    = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/del"
	CorpOAGetMeetingRoomBookingInfo            = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info"
	CorpOAMeetingRoomBook                      = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/book"
	CorpOAMeetingRoomCancelBook                = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/cancel_book"
	CorpOAGetMeetingRoomBookingInfoByMeetingID = "https://qyapi.weixin.qq.com/cgi-bin/oa/meetingroom/get_booking_info_by_meeting_id"
	CorpOACallPstncc                           = "https://qyapi.weixin.qq.com/cgi-bin/pstncc/call"
	CorpOAGetPstnccStates                      = "https://qyapi.weixin.qq.com/cgi-bin/pstncc/getstates"
)

// tools
const (
	CorpToolsCalendarAdd              = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/add"
	CorpToolsCalendarUpdate           = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/update"
	CorpToolsCalendarGet              = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/get"
	CorpToolsCalendarDelete           = "https://qyapi.weixin.qq.com/cgi-bin/oa/calendar/del"
	CorpToolsScheduleAdd              = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add"
	CorpToolsScheduleUpdate           = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/update"
	CorpToolsScheduleGet              = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get"
	CorpToolsScheduleDelete           = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del"
	CorpToolsScheduleGetByCalendar    = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/get_by_calendar"
	CorpToolsScheduleAttendeeAdd      = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/add_attendees"
	CorpToolsScheduleAttendeeDelete   = "https://qyapi.weixin.qq.com/cgi-bin/oa/schedule/del_attendees"
	CorpToolsLivingCreate             = "https://qyapi.weixin.qq.com/cgi-bin/living/create"
	CorpToolsLivingModify             = "https://qyapi.weixin.qq.com/cgi-bin/living/modify"
	CorpToolsLivingCancel             = "https://qyapi.weixin.qq.com/cgi-bin/living/cancel"
	CorpToolsLivingDeleteReplayData   = "https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data"
	CorpToolsLivingGetCode            = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_code"
	CorpToolsLivingGetUserAllLivingID = "https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid"
	CorpToolsLivingGetInfo            = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_info"
	CorpToolsLivingGetWatchStat       = "https://qyapi.weixin.qq.com/cgi-bin/living/get_watch_stat"
	CorpToolsLivingGetShareInfo       = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_share_info"
	CorpToolsWedriveSpaceCreate       = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_create"
	CorpToolsWedriveSpaceRename       = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_rename"
	CorpToolsWedriveSpaceDismiss      = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_dismiss"
	CorpToolsWedriveSpaceInfo         = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_info"
	CorpToolsWedriveSpaceAclAdd       = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_acl_add"
	CorpToolsWedriveSpaceAclDelete    = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_acl_del"
	CorpToolsWedriveSpaceSetting      = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_setting"
	CorpToolsWedriveSpaceShare        = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/space_share"
	CorpToolsWedriveFileList          = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_list"
	CorpToolsWedriveFileUpload        = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_upload"
	CorpToolsWedriveFileDownload      = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_download"
	CorpToolsWedriveFileCreate        = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_create"
	CorpToolsWedriveFileRename        = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_rename"
	CorpToolsWedriveFileMove          = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_move"
	CorpToolsWedriveFileDelete        = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_delete"
	CorpToolsWedriveFileInfo          = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_info"
	CorpToolsWedriveFileAclAdd        = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_acl_add"
	CorpToolsWedriveFileAclDelete     = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_acl_del"
	CorpToolsWedriveFileSetting       = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_setting"
	CorpToolsWedriveFileShare         = "https://qyapi.weixin.qq.com/cgi-bin/wedrive/file_share"
	CorpToolsDialRecordGet            = "https://qyapi.weixin.qq.com/cgi-bin/dial/get_dial_record"
)

// payment
const (
	CorpPaymentMerchantAdd    = "https://qyapi.weixin.qq.com/cgi-bin/externalpay/addmerchant"
	CorpPaymentMerchantGet    = "https://qyapi.weixin.qq.com/cgi-bin/externalpay/getmerchant"
	CorpPaymentMerchantDelete = "https://qyapi.weixin.qq.com/cgi-bin/externalpay/delmerchant"
	CorpPaymentMchUseScopeSet = "https://qyapi.weixin.qq.com/cgi-bin/externalpay/set_mch_use_scope"
	CorpPaymentBillListGet    = "https://qyapi.weixin.qq.com/cgi-bin/externalpay/get_bill_list"
)

// cropgroup
const (
	CorpGroupListAppShareInfo     = "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/list_app_share_info"
	CorpGroupGetAccessToken       = "https://qyapi.weixin.qq.com/cgi-bin/corpgroup/corp/gettoken"
	CorpGroupMinipTransferSession = "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/transfer_session"
)

// msgaudit
const (
	CorpMsgAuditGetPermitUserList = "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/get_permit_user_list"
	CorpMsgAuditCheckSingleAgree  = "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/check_single_agree"
	CorpMsgAuditCheckRoomAgree    = "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/check_room_agree"
	CorpMsgAuditGroupChatGet      = "https://qyapi.weixin.qq.com/cgi-bin/msgaudit/groupchat/get"
)

// invoice
const (
	CorpInvoiceGetInfo           = "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/getinvoiceinfo"
	CorpInvoiceBatchGetInfo      = "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch"
	CorpInvoiceUpdateStatus      = "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/updateinvoicestatus"
	CorpInvoiceBatchUpdateStatus = "https://qyapi.weixin.qq.com/cgi-bin/card/invoice/reimburse/updatestatusbatch"
)

// school
const (
	CorpSchoolGetSubscribeQRCode            = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_qr_code"
	CorpSchoolSetSubscribeMode              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/set_subscribe_mode"
	CorpSchoolGetSubscribeMode              = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_subscribe_mode"
	CorpSchoolGetAgentAllowScope            = "https://qyapi.weixin.qq.com/cgi-bin/school/agent/get_allow_scope"
	CorpSchoolStudentCreate                 = "https://qyapi.weixin.qq.com/cgi-bin/school/user/create_student"
	CorpSchoolStudentBatchCreate            = "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_create_student"
	CorpSchoolStudentDelete                 = "https://qyapi.weixin.qq.com/cgi-bin/school/user/delete_student"
	CorpSchoolStudentBatchDelete            = "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_delete_student"
	CorpSchoolStudentUpdate                 = "https://qyapi.weixin.qq.com/cgi-bin/school/user/update_student"
	CorpSchoolStudentBatchUpdate            = "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_update_student"
	CorpSchoolParentCreate                  = "https://qyapi.weixin.qq.com/cgi-bin/school/user/create_parent"
	CorpSchoolParentBatchCreate             = "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_create_parent"
	CorpSchoolParentDelete                  = "https://qyapi.weixin.qq.com/cgi-bin/school/user/delete_parent"
	CorpSchoolParentBatchDelete             = "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_delete_parent"
	CorpSchoolParentUpdate                  = "https://qyapi.weixin.qq.com/cgi-bin/school/user/update_parent"
	CorpSchoolParentBatchUpdate             = "https://qyapi.weixin.qq.com/cgi-bin/school/user/batch_update_parent"
	CorpSchoolUserGet                       = "https://qyapi.weixin.qq.com/cgi-bin/school/user/get"
	CorpSchoolUserList                      = "https://qyapi.weixin.qq.com/cgi-bin/school/user/list"
	CorpSchoolSetArchSyncMode               = "https://qyapi.weixin.qq.com/cgi-bin/school/set_arch_sync_mode"
	CorpSchoolParentList                    = "https://qyapi.weixin.qq.com/cgi-bin/school/user/list_parent"
	CorpSchoolDepartmentCreate              = "https://qyapi.weixin.qq.com/cgi-bin/school/department/create"
	CorpSchoolDepartmentUpdate              = "https://qyapi.weixin.qq.com/cgi-bin/school/department/update"
	CorpSchoolDepartmentDelete              = "https://qyapi.weixin.qq.com/cgi-bin/school/department/delete"
	CorpSchoolDepartmentList                = "https://qyapi.weixin.qq.com/cgi-bin/school/department/list"
	CorpSchoolSetUpgradeInfo                = "https://qyapi.weixin.qq.com/cgi-bin/school/set_upgrade_info"
	CorpSchoolGetHealthReportStat           = "https://qyapi.weixin.qq.com/cgi-bin/health/get_health_report_stat"
	CorpSchoolGetHealthReportJobIDs         = "https://qyapi.weixin.qq.com/cgi-bin/health/get_report_jobids"
	CorpSchoolGetHealthReportJobInfo        = "https://qyapi.weixin.qq.com/cgi-bin/health/get_report_job_info"
	CorpSchoolGetHealthReportAnswer         = "https://qyapi.weixin.qq.com/cgi-bin/health/get_report_answer"
	CorpSchoolGetTeacherCustomizeHealthInfo = "https://qyapi.weixin.qq.com/cgi-bin/school/user/get_teacher_customize_health_info"
	CorpSchoolGetStudentCustomizeHealthInfo = "https://qyapi.weixin.qq.com/cgi-bin/school/user/get_student_customize_health_info"
	CorpSchoolGetHealthQRCode               = "https://qyapi.weixin.qq.com/cgi-bin/school/user/get_health_qrcode"
	CorpSchoolGetUserAllLivingID            = "https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid"
	CorpSchoolGetLivingInfo                 = "https://qyapi.weixin.qq.com/cgi-bin/school/living/get_living_info"
	CorpSchoolGetLivingWatchStat            = "https://qyapi.weixin.qq.com/cgi-bin/school/living/get_watch_stat"
	CorpSchoolGetLivingUnwatchStat          = "https://qyapi.weixin.qq.com/cgi-bin/school/living/get_unwatch_stat"
	CorpSchoolDeleteLivingReplayData        = "https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data"
	CorpSchoolGetPaymentResult              = "https://qyapi.weixin.qq.com/cgi-bin/school/get_payment_result"
	CorpSchoolGetTrade                      = "https://qyapi.weixin.qq.com/cgi-bin/school/get_trade"
)

// report
const (
	CorpReportGridAdd                   = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/add"
	CorpReportGridUpdate                = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/update"
	CorpReportGridDelete                = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/delete"
	CorpReportGridList                  = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/list"
	CorpReportGetUserGridInfo           = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/get_user_grid_info"
	CorpReportGridCataAdd               = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/add_cata"
	CorpReportGridCataUpdate            = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/update_cata"
	CorpReportGridCataDelete            = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/delete_cata"
	CorpReportGridCataList              = "https://qyapi.weixin.qq.com/cgi-bin/report/grid/list_cata"
	CorpReportGetPatrolGridInfo         = "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_grid_info"
	CorpReportGetPatrolCorpStatus       = "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_corp_status"
	CorpReportGetPatrolUserStatus       = "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_user_status"
	CorpReportPatrolCategoryStatistic   = "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/category_statistic"
	CorpReportGetPatrolOrderList        = "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_order_list"
	CorpReportGetPatrolOrderInfo        = "https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_order_info"
	CorpReportGetResidentGridInfo       = "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_grid_info"
	CorpReportGetResidentCorpStatus     = "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_corp_status"
	CorpReportGetResidentUserStatus     = "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_user_status"
	CorpReportResidentCategoryStatistic = "https://qyapi.weixin.qq.com/cgi-bin/report/resident/category_statistic"
	CorpReportGetResidentOrderList      = "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_order_list"
	CorpReportGetResidentOrderInfo      = "https://qyapi.weixin.qq.com/cgi-bin/report/resident/get_order_info"
	CorpReportGetSiteCodeList           = "https://qyapi.weixin.qq.com/cgi-bin/report/sitecode/list"
	CorpReportGetSiteCodeReportInfo     = "https://qyapi.weixin.qq.com/cgi-bin/report/sitecode/get_site_report_info"
	CorpReportGetSiteCodeReportAnswer   = "https://qyapi.weixin.qq.com/cgi-bin/report/sitecode/get_report_answer"
)
