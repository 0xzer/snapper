package paths

var ACCOUNTS_BASE_URL = "https://accounts.snapchat.com"

var ACCOUNTS_SSO = ACCOUNTS_BASE_URL+"/accounts/sso?client_id=web-calling-corp--prod"

var WEB_BASE_URL = "https://web.snapchat.com"
var APP_VERSION_URL = WEB_BASE_URL+"/version.json"
var GET_FRIENDS = WEB_BASE_URL+"/ami/friends"
var GET_USER_PUBLIC_INFO = WEB_BASE_URL+"/loq/snapchatter_public_info"

var WEB_GRAPHQL_URL = WEB_BASE_URL+"/web-calling-api/graphql"

var NOTIFICATION_SERVICE_URL = WEB_BASE_URL+"/snapchat.notification.notificationdata.PushNotificationDataRegistryService"
var GET_NOTIFICATION_SETTING_URL = NOTIFICATION_SERVICE_URL+"/GetNotificationSetting"
var NOTIFICATIONS_REGISTER_DEVICE_URL = NOTIFICATION_SERVICE_URL+"/RegisterDevice"

var MESSAGING_SERVICE_URL = WEB_BASE_URL+"/messagingcoreservice.MessagingCoreService"
var GET_GROUPS_URL = MESSAGING_SERVICE_URL+"/GetGroups"
var SYNC_CONVERSATIONS = MESSAGING_SERVICE_URL+"/SyncConversations"
var BATCH_DELTA_SYNC_CONVERSATIONS = MESSAGING_SERVICE_URL+"/BatchDeltaSync"
var DELTA_SYNC_CONVERSATIONS = MESSAGING_SERVICE_URL+"/DeltaSync"
var QUERY_CONVERSATIONS = MESSAGING_SERVICE_URL+"/QueryConversations"
var QUERY_MESSAGES = MESSAGING_SERVICE_URL+"/QueryMessages"
var CREATE_CONTENT_MESSAGE = MESSAGING_SERVICE_URL+"/CreateContentMessage"
var UPDATE_CONTENT_MESSAGE = MESSAGING_SERVICE_URL+"/UpdateContentMessage"
var SEND_TYPING_NOTIFICATION = MESSAGING_SERVICE_URL+"/SendTypingNotification"
var CREATE_CONVERSATION = MESSAGING_SERVICE_URL+"/CreateConversation"
var UPDATE_CONVERSATION = MESSAGING_SERVICE_URL+"/UpdateConversation"

var FIDELIUS_SERVICE_URL = WEB_BASE_URL+"/snapchat.fidelius.FideliusIdentityService"
var INITIALIZE_WEB_KEY = FIDELIUS_SERVICE_URL+"/InitializeWebKey"
var GET_FRIEND_KEYS = FIDELIUS_SERVICE_URL+"/GetFriendKeys"