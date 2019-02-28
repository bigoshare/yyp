package yyp

/* -------------------------------------------------------------------------
★ 0 ~ 99 验证模块错误码;

★ 1xx 说明性的返回;

★ 200 表示处理成功;

★ 3xx是告诉客户端需要进一步的动作才能完成功能;
★ 4xx是客户端的原因导致的出错;
★ 5xx是服务器端的原因导致的出错.
	500 ~ 549 通用错误
	550 ~ 599 locate 定位服务中转错误

★ 6xx是 popoc 使用

★ 700 ~ 800 会话 session 错误码字段

★ 801 ~ 900 短信服务 错误码字段

★ 返回代码使用uint16_t的数据类型.
   ------------------------------------------------------------------------- */

/*
login 登陆事件
AuthService 验证服务事件
验证错误不能忽略，会引起登陆流程中断。
忽略错误
*/

const AUTH_TICKET_OK = 0 /*密码验证，ticket ok 原来的 AUTH_READY0*/
const AUTH_LINK_OK = 1   /*connection ok  AUTH_READY*/

// AuthService 与 srvbase 交互，用来拆除绑定连接
const AUTH_LINK_DESTROY = 2         /* 链接已经销毁 */
const AUTH_LINK_DESTROY_TIMEOUT = 3 /* 链接由于idle-timeou而销毁, 主要用于服务器 */

// 流程相关
const AUTH_EUIDPASS = 4      /* 用户名密码错误 */
const AUTH_EPROXYUIDPASS = 5 /* 代理用户名密码错误 */

// 其他错误
const AUTH_ENET = 6     /* 网络错误 */
const AUTH_EPROCESS = 7 /* 处理错误 */
const AUTH_DESTROY = 8  /* Auth服务Destroy，XXX内部使用作其他作用 */

const AUTH_ENET_TICKET_TIMO = 9        /*客户端获取ticket超时*/
const AUTH_ENET_LINK_TIMO = 10         /*客户端链接linkd超时*/
const AUTH_EPROXY_NET_TICKET_TIMO = 11 /*客户端代理获取ticket超时*/
const AUTH_EPROXY_NET_LINK_TIMO = 12   /*客户端代理获取链接超时*/
const AUTH_ETICKET = 14                /*服务器拒绝提交的Ticket*/

const AUTH_EMAKECOOKIE = 20 /*服务器拒绝提交的Ticket*/

// srvbase 的登陆细分事件
const PP_MYINFO_OK = 50 // 自己的资料，私有配置获取成功
const PP_MYLIST_OK = 51 // 好友列表同步完成
const PP_LOGIN_OK = 52  // 登陆完成

//////////////////////////////////////////////////////////////////

const RES_SUCCESS = 200   /* 功能成功完成,一切正常,不必担心 */
const RES_ACCEPTED = 202  /* 功能请求已接受,等候处理 */
const RES_NOCHANGED = 204 /* 结果没有变，用于CHECKSUM */

const RES_ERETRY = 300 /* 暂时无法受理,建议稍后再试 */

const RES_EREQUEST = 400       /* 碰到无法理解的功能请求 */
const RES_EAUTH = 401          /* 请求者身份未经认证,因此后台拒绝完成功能 */
const RES_EPERM = 403          /* 对方实体没授权(可不是后台不受理) */
const RES_ENONEXIST = 404      /* 目标(对象或用户)不存在 */
const RES_EACCESS = 405        /* 无权限请求此项功能 */
const RES_EQUOTA = 406         /* 用户数据存储超过了限额 */
const RES_EVOLATILE = 407      /* 某些有生存时间限制的资源已经超时不可用 */
const RES_ETIMEOUT = 408       /* 请求过程超时 */
const RES_ECONFLICT = 409      /* 资源或者对象冲突(比如重名) */
const RES_EPARAM = 414         /* 参数数据出错.(越界,超长,不完整等原因) */
const RES_EDBERROR = 415       /* 数据库操作失败 */
const RES_EDBNVALID = 416      /* 数据库暂时不可用，可能是正在维护 */
const RES_ECONNURS = 417       /* 连接URS的错误 */
const RES_NODEMOVED = 418      /* 节点已经被转移*/
const RES_EUNKNOWN = 500       /* 出错了.但是原因未明,或是不便透露给你 */
const RES_EBUSY = 504          /* 后台忙,拒绝处理 */
const RES_EPROT = 505          /* 后台不支持此协议版本 */
const RES_EOVERTIMES = 453     /* 操作次数太多了 */
const RES_EDATANOSYNC = 506    /* 与数据库数据不同步,client需要重新获得数据 */
const RES_ENOTENOUGH = 507     /* 数量不够 */
const RES_ENONEXIST_REAL = 508 /* 目标(对象或用户)不存在 */

//memcache
const RES_MCNOEXISIT = 509

//owner is Banned
const RES_OWNER_BANNED = 510

/* NOTE:针对某些特定功能的特殊含义待扩充 */

/* 550~599 locate 中转错误 begin */
const RES_ESERVICE = 550  /* 不支持的服务 */
const RES_EDAEMON = 551   /* 服务器未找到 */
const RES_EUNUSABLE = 552 /* 服务暂时不可用 */
const RES_ECONNMISS = 553 /* 内部错误，根据connid找不到连接 */
const RES_EBUFOVR = 554   /* 发送请求给内部服务器时，缓冲不够，瓶颈保护用。*/
const RES_ECONNECT = 555  /* 内部服务连接不上,
                                   通过 IConnectErrorHandle 报告给应用，由应用填写 */
const RES_ESENDREQ = 556  /* 发送请求发生异常 */
const RES_EHASHTYPE = 557 /* 内部配置错误 */
const RES_EPACKET = 558   /* 请求数据包错误 */

const RES_ELOCATE = 559 /* 定位错误，服务器收到不属于自己的请求 */

const RES_LIMITUSER = 580 /* 不允许一台机器上登录两个相同账号 */

/* 599 linkd 错误 */
const RES_ETRUNKCLI = 599 /* 不允许客户端发送 trunk */
/* locate 中转错误 end */

/* 600~699 local client error begin */

// general
const PP_ENETBROKEN = 600  /* 连接断开 */
const PP_EPROCESSEXP = 601 /* 网络回调处理过程中发生异常 */
const PP_ETIMEOUTCTX = 602 /* 上下文等待请求结果超时 */
const PP_EAPPLYCTX = 603   /* XXX. apply 异常，会使用这个错误码，再一次apply */
const PP_ESRVDESTROY = 604 /* 服务销毁的时候，对于某些还存在的上下文，报告这个错误 */

// session
const PP_ESSOPENING = 609      /* session opening */
const PP_ESSNOTREADY = 610     /* session not ready */
const PP_ESSCREATEVIEW = 611   /* 会话，创建UI-View 失败 */
const PP_ESSSAYSELF = 612      /* session 没有任何下接收这，自言自语 */
const PP_ESSCTXINDESTROY = 613 /* 会话销毁中，但上下文还存在 */
const PP_ECTXINPASSIVE = 614   /* 被动模式不能有上下文 */
const PP_ESSIDMISSMATCH = 615  /* 加入会话的 ssid 和返回的 ssid 不匹配 */
const PP_EDUPSSID = 616        /* 创建会话返回的ssid重复了 */
const PP_ECREATESSID = 617     /* 创建会话返回不正确的 ssid */
const PP_ESSNOTEXIST = 618     /* 会话不存在了，一般是群，聊天室 */
const PP_ESSIMNOTTHERE = 619   /* 自己不在列表里面，一般是被踢出去了 */
const PP_ESSIDUNKNOWN = 620    /* 不认识的 ssid */
const PP_EFORUMDISMISSED = 621 /* 话题被解散了 */
const PP_ENOTMEMBER = 622      /* 不是成员，话题，群等，加入失败 */
const PP_ESSDISMISSED = 623    /* 会话被解散了 */
const PP_ESSNOTOPEN = 624      /* 会话没有实例化，根本没有被打开 */

//session manager
const RES_CHANNEL_MOVED = 625     //频道已经迁移到新版
const RES_CHANNEL_NOT_MOVED = 626 //频道未迁移到新版

// parameter
const PP_EBADUID = 630     /* 错误的uid */
const PP_EUNKNOWNEXP = 631 /* 不知道类型的异常 */
const PP_ENWPRES = 632     /* NofityService 不能注册状态的通知 */
const PP_ENWEXIST = 633    /* NofityService 通知(nid) 已经被注册 */
const PP_ENWREVOKE = 634   /* NofityService 注销失败，参数不匹配 */
const PP_EWITHPARENT = 635 /* Parent 必须单独修改，不能和其他属性一起修改 */
const PP_ESSPARAM = 636    /* 一般的参数错误 */
const PP_ESUPEREXIT = 637  /* Group :: Super can't exit */

const PP_ENOTIMPLEMENT = 640 /* 操作没有实现 */

//// operate
//#define PP_OPERR(op, err)  ((op<<16)|err)
//#define PP_GETOP(err)      (((unsigned long)(err))>>16)
//#define PP_GETERR(err)     (err & 0xffff)

const OP_CREATEGINFO = 660 /* 创建群的第一步，Create Ginfo */
const OP_JOINGROUP = 661   /* uglist.add ， 加入群 */
const OP_OPENGLIST = 662   /* OpenGroupList */

/* local client error end */

//700~800 会话session错误码字段
//const RES_SS_ =    7xx     /*  */
//701~739  session公有错误
//741~759 standard session错误
//761~779 room session错误
//781~799 group错误

// 严重错误，客户端应该禁用会话
// unprotected error for client
const SS_ESNOTEXIST = 701   /*  · */
const SS_EUNOTEXIST = 702   /* 用户不在会话中 */
const SS_ETIMEOUT = 705     /* 定时器操时 */
const SS_EBANPENANCE = 765  /* 禁闭室 */
const SS_EJOINTOOFAST = 766 /* 加入退出聊天室过于频繁*/

// 非致命错误，一般是操作引起的，仅表示当前操作失败
// error can be proccessed easily by client
const SS_EDELIVER = 703        /* 后台服务器传输错误 */
const SS_ENOAUTH = 704         /* 请求越权了 */
const SS_EINVALIDOBJ = 706     /* 不合法的操作对象*/
const SS_ETOOMANYUSER = 707    /* 已经达到会话人数上限*/
const SS_EUNOTOPEN = 708       /* 用户没有打开会话就发送其它请求*/
const SS_EBANNED = 761         /* 黑名单 */
const SS_ENESTFORUM = 763      /* 嵌套的话题 */
const SS_EFOLDERNOTEXIST = 764 /* 目录或者话题不存在 */
const SS_ECHATTOOFAST = 767    /* 聊天室里说话太快*/
const SS_EPRIVATEFORUM = 783   /* 私有的话题 */
const SS_EFOLDERNEMPTY = 784   /* 组织结构非空 */
const SS_EPROTECTEDFORUM = 785 /* 受保护的话题 */
const SS_EREJECTAUTO = 786     /* 被设置自动拒绝*/
const SS_EREJECTADMIN = 787    /* 被管理员自动拒绝*/
const SS_ETOGROUPLIMIT = 788   /* 已经达到群人数上限,无法再添加用户*/

/*
srvbase 会话错误处理(从严)
    对于 "非致命错误" 仅仅报告。
	其他错误全部禁用会话。

	TODO 识别更多的非致命错误
	see protocol/popoc/Session.cpp::onError
*/

//800~900 短信sms错误码字段
const RES_SMS_ELIMIT = 801          /* 超过发送限制 */
const RES_SMS_ETOOLONG = 802        /* 文字太长     */
const RES_SMS_EBANNED = 803         /* 手机被禁用   */
const RES_SMS_EINACTIVE = 804       /* 对方没有注册手机 */
const RES_SMS_EINACTIVE_SELF = 805  /* 用户没有注册手机 */
const RES_SMS_ENETWORK = 806        /* 网络错误 */
const RES_SMS_EPEERPERMIT = 807     /* 对方禁止接收 */
const RES_SMS_EFREQ = 808           /* 发送过快 */
const RES_SMS_EBANQF = 809          /* 欠费 */
const RES_SMS_ELIANTONG = 810       /* 对方是联通手机，没有注册 */
const RES_SMS_EORDER = 811          /* 定制没有成功 */
const RES_SMS_ENOTSAME_NET = 812    /* 联通移动不能互通 */
const RES_SMS_ENOENOUGH_PAOBI = 813 /* 泡币不足 */
const RES_SMS_EBADWORD = 814        /* */
const RES_SMS_EPARAMETER = 815      /* 参数错误 */
const RES_SMS_EOTHER_ERROR = 816    /* 其它错误 */

/* roster 900~950 */
const ROSTER_NOTREADY = 901 /* Roster 没有准备好, 一般是转载数据失败 */
const ROSTER_ERRLOGIN = 902 /* NotLogin or NotAnswer or NotOwner */

/*-------------------------------------------
	用最高一位区分操作是否成功，其他位代表原因。
---------------------------------------------*/

//#define ISOK(rc)   (0x8000 & rc) /* 判断返回是否成功 */
//#define RSCODE(rc) (0x7fff & rc) /* 提取返回码 */
//#define RSOK(rc)   (0x8000 | rc) /* 生成成功返回码 */
//#define RSERR(rc)  (0x7fff & rc) /* 生成失败返回码 */
//#define RES_NULL   0

//////////////////////////////////////////////////////////
// 大于 10000 都保留给客户端

//通用错误
const RES_ACCOUNT_TOOSHORT = 30000 // 帐号太短了，必须大于３个字
const RES_ACCOUNT_LIMIT = 30001    // 帐号超长
const RES_ACCOUNT_EMPTY = 30002    // 帐号为空
const RES_ACCOUNT_INVALID = 30003  // 帐号格式出错

const RES_PASS_EMPTY = 30004 // 密码为空
const RES_PASS_LIMIT = 30005 // 密码超长

const RES_NICK_EMPTY = 30006   // 昵称为空
const RES_NICK_INVALID = 30007 // 昵称格式出错
const RES_NICK_LIMIT = 30008   // 昵称超长

const RES_YEAR_INVALID = 30009  // 年份有误
const RES_MONTH_INVALID = 30010 // 月份有误
const RES_DAY_INVALID = 30011   // 日期有误

//登录的专有错误
const RES_LOGIN_OFFLINENOPASS = 30012 //离线登录但没保存密码

//专有错误

//代理服务器设置
const RES_PROXY_SERVEREMPTY = 10009    // 代理服务器的地址为空
const RES_PROXY_SERVER_INVALID = 10010 // 代理服务器的地址出错
const RES_PROXY_PORTEMPTY = 10011      // 代理服务器的端口为空

//个人信息框错误
const RES_UINFO_PROVINCE_EMPTY = 10012   // 个人信息框的省分空白
const RES_UINFO_PROVINCE_INVALID = 10013 // 个人信息框的省分格式出错
const RES_UINFO_PROVINCE_LIMIT = 10014   // 个人信息框的省分超长

const RES_UINFO_NOTENICK_INVALID = 10015 // 个人信息框的备注昵称出错
const RES_UINFO_NOTENICK_LIMIT = 10016   // 个人信息框的备注昵称超长

//系统配置框错误

//创建群的错误
const RES_CLUSTE_NAME_EMPTY = 10017        // 群名称为空
const RES_CLUSTE_NAME_INVALID = 10018      // 群名称格式出错
const RES_CLUSTE_NAME_LIMIT = 10019        // 群名称超长
const RES_CLUSTE_TYPE_EMPTY = 10020        // 群类型为空
const RES_CLUSTE_CARD_NICK_INVALID = 10021 // 群名片的昵称出错
const RES_CLUSTE_CARD_NICK_LIMIT = 10022   // 群名片的昵称超长
const RES_CLUSTE_CARD_NOTE_INVALID = 10023 // 群名片中的备注出错
const RES_CLUSTE_CARD_NOTE_LIMIT = 10024   // 群名片中的备注超长

//查找群的错误
const RES_FINDCLUSTE_ACCOUNT_EMPTY = 10025 // 群搜索的帐号为空
const RES_FINDCLUSTE_ACCOUNT_LIMIT = 10026 // 群搜索的帐号超长

const RES_FINDCLUSTE_NAME_EMPTY = 10027 // 群搜索的名字为空
const RES_FINDCLUSTE_NAME_LIMIT = 10028 // 群搜索的名字超长

const RES_FINDCLUSTE_TYPE_INVALID = 10029 // 群搜索的类型有误

//IM错误
const RES_REACHE_MAX_OFFLINEMSG = 10030  //到达最大离线消息数量
const RES_ALREDY_BUDDY = 10031           //已经存在于自己的好友列表中
const RES_CANNOT_ADD_SELE = 10032        //不能自己加自己为好友
const RES_SNEDMSG_FAIL_NOT_BUDDY = 10033 //不是自己的好友

type RES_CODE uint16
