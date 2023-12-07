/*

 */

package server

import (
	"errors"
	"log"

	"github.com/bitspawngg/bitspawn-api/config"
	"github.com/bitspawngg/bitspawn-api/utils/ratelimit"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	socketio "github.com/googollee/go-socket.io"
)

func NewRouter(server *Server, conf config.Config, socketserver *socketio.Server) *gin.Engine {
	gin.SetMode(server.config.GinMode())
	r := gin.Default()

	r.Use(ratelimit.GinMiddleware())
	r.Use(CORSMiddleware())
	r.Use(static.Serve("/", static.LocalFile("/app/Demo-UI", true)))

	r.GET("/ping", server.controller.base.HandlePing)

	v1 := r.Group("/v1")
	authorizedV1 := v1.Group("/")
	authorizedV1.Use(server.GetAuth())
	twoFAV1 := authorizedV1.Group("/")
	twoFAV1.Use(server.Get2FAuth())

	v1Team := v1.Group("/teams")
	v1Team.Use(server.GetAuth())

	v2 := r.Group("/v2")
	v2_authorized := v2.Group("/")
	v2_authorized.Use(server.GetAuth())
	v2_2FA := v2_authorized.Group("/")
	v2_2FA.Use(server.Get2FAuth())

	WithV2UserRoutes(v2_authorized, server, conf)
	WithV2UnauthorizedManualTournamentRoutes(v2, server, conf)

	// todo: add notification socket
	//r.GET("/socket/*any", gin.WrapH(socketserver))
	//r.POST("/socket/*any", gin.WrapH(socketserver))
	v1.POST("/user/signup", server.controller.user.HandleSignup)

	With2FARoutes(twoFAV1, server, conf)

	WithAdminRoutes(v1, server, conf)

	WithAuthorizedManualTournamentRoutes(v2_authorized, server, conf)
	WithSocialRoutes(v2_authorized, server)

	WithInvitationRoutes(v2_authorized, server, conf)

	WithAuthorizedChallengeRoutes(authorizedV1, server, conf)
	WithUnauthorizedChallengeRoutes(v1, server, conf)

	WithUserRoutes(authorizedV1, server, conf)
	WithProfileRoutes(v1, server, conf)

	WithFriendsRoutes(authorizedV1, server, conf)

	WithGameRoutes(authorizedV1, server, conf)
	WithBadgeRoutes(authorizedV1, server, conf)

	WithTeamRoutes(v1Team, server)
	return r
}
func WithSocialRoutes(r *gin.RouterGroup, server *Server) {
	socialRoutes := r

	socialRoutes.GET("/socials", server.controller.social.HandleSocialLinkGetAll)
	socialRoutes.GET("/socials/:social_type", server.controller.social.HandleSocialLinkGet)
	socialRoutes.DELETE("/socials/:social_type", server.controller.social.HandleSocialLinkUnRegister)
	socialRoutes.POST("/socials", server.controller.social.HandleSocialLinkRegister)
	socialRoutes.PUT("/socials/:social_type", server.controller.social.HandleSocialLinkUpdate)
}
func WithTeamRoutes(r *gin.RouterGroup, server *Server) {
	r.GET("/", server.controller.team.HandleTeamList)
	r.POST("/", server.controller.team.HandleTeamRegister)
	r.DELETE("/:team_id", server.controller.team.HandleTeamUnRegister)
	r.GET("/:team_id", server.controller.team.HandleGetTeam)
	r.PUT("/:team_id", server.controller.team.HandleChangeTeam)
	r.POST("/:team_id/members", server.controller.team.HandleAddTeamMember)
	r.DELETE("/:team_id/members/:member_id", server.controller.team.HandleRemoveMember)
	r.PATCH("/:team_id/members/:member_id", server.controller.team.HandleMembershipStatus)
	r.GET("/:team_id/members/:member_id", server.controller.team.HandleGetTeamMember)
	r.GET("/:team_id/members", server.controller.team.HandleTeamMemberList)
	r.PUT("/:team_id/avatar", server.controller.team.HandleTeamAvatarUpload)
	r.PUT("/:team_id/roles", server.controller.team.HandleChangeRole)
	r.POST("/:team_id/transfer", server.controller.team.HandleTransfer)
}
func With2FARoutes(twoFA *gin.RouterGroup, server *Server, conf config.Config) {
	twoFA.POST("/user/2fa/disable", server.controller.user.HandleUser2FADisable)
	twoFA.POST("/user/2fa/enable", server.controller.user.HandleUser2FAEnable)
	twoFA.POST("/user/2fa/verify", server.controller.user.HandleUser2FAVerify)
}

func WithAdminRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	//todo: config admin routes based on config
	admin := r.Group("/admin")
	admin.Use(adminAuth())

	admin.POST("/tournament/cancel", server.controller.tournament.HandleAdminTournamentCancel)
	admin.POST("/tournament/complete", server.controller.tournament.HandleAdminTournamentComplete)
	admin.POST("/tournament/dequeue", server.controller.tournament.HandleAdminTournamentDequeue)
	admin.POST("/challenge/create", server.controller.challenge.HandleAdminChallengeCreate)
	admin.POST("/challenge/cancel", server.controller.challenge.HandleAdminChallengeCancel)
	admin.POST("/challenge/complete", server.controller.challenge.HandleAdminChallengeComplete)
	admin.POST("/challenge/fund", server.controller.challenge.HandleAdminChallengeFund)
	admin.POST("/challenge/reportscore", server.controller.challenge.HandleAdminChallengeReportScore)

	//todo: confirm
	admin.POST("/gameType/createSubType", server.controller.game.HandleGameSubTypeCreate)
	admin.POST("/gameType/create", server.controller.game.HandleGameTypeCreate)

	admin.POST("/creatematchschedule", server.controller.tournament.HandleCreateMatchSchedule)
	admin.POST("/starttournament", server.controller.tournament.HandleStartTournament)
	admin.POST("/preparemanualmatches", server.controller.tournament.HandlePrepareManualMatches)
	admin.POST("/preparematches", server.controller.tournament.HandlePrepareMatches)
	admin.POST("/preparerematches", server.controller.tournament.HandlePrepareRematches)
	admin.POST("/prepareroundone", server.controller.tournament.HandlePrepareRoundOne)
	admin.POST("/replace_platform_username", server.controller.user.HandleReplaceUsername)

}

func WithAuthorizedManualTournamentRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r

	authorized.POST("/tournaments/:tournamentId/matches/:matchId/games", server.controller.tournament.HandleV2TournamentMatchGameCreate)
	authorized.GET("/tournaments/:tournamentId/matches/:matchId/games", server.controller.tournament.HandleV2TournamentMatchGameGet)
	authorized.PATCH("/tournaments/:tournamentId/matches/:matchId/games/:gameNumber", server.controller.tournament.HandleV2TournamentMatchGameUpdate)
	authorized.PUT("/tournament/match/score", server.controller.tournament.HandleV2TournamentMatchScore)

	authorized.POST("/tournament/create", server.controller.tournament.HandleTournamentCreate)
	authorized.PUT("/tournament/upload/banner/:tournamentId", server.controller.tournament.HandleTournamentBannerUpload)
	authorized.PUT("/tournament/upload/logo/:tournamentId", server.controller.tournament.HandleTournamentLogoUpload)
	authorized.PUT("/tournament/upload/thumbnail/:tournamentId", server.controller.tournament.HandleTournamentThumbnailUpload)
	authorized.PUT("/tournament/update", server.controller.tournament.HandleTournamentUpdate)
	authorized.POST("/tournament/publish", server.controller.tournament.HandleTournamentPublish)
	authorized.POST("/tournament/roles/moderators/add", server.controller.tournament.HandleTournamentModeratorAdd)
	authorized.POST("/tournament/roles/moderators/delete", server.controller.tournament.HandleTournamentModeratorDelete)
	authorized.POST("/tournament/team/name", server.controller.tournament.HandleTournamentTeamName)
	authorized.POST("/tournament/team/players/add", server.controller.tournament.HandleTournamentTeamPlayersAdd)
	authorized.POST("/tournament/team/players/remove", server.controller.tournament.HandleTournamentTeamPlayersRemove)
	authorized.GET("/tournament/enum", server.controller.tournament.HandleTournamentEnum)

	authorized.GET("/tournament/user/fetch/:tournamentId", server.controller.tournament.HandleTournamentFetchUserStatus)
	authorized.GET("/tournament/register/status/:tournamentId", server.controller.tournament.HandleV2TournamentRegisterStatus)
	authorized.GET("/tournament/list/user", server.controller.tournament.HandleV2MyTournaments)
	authorized.GET("/tournament/list/registered", server.controller.tournament.HandleV2TournamentListRegistered)
	authorized.GET("/tournament/list/role/moderator", server.controller.tournament.HandleV2MyModeratorTournaments)
	authorized.GET("/tournament/hosted/list/participants", server.controller.tournament.HandleV2TournamentHostedListParticipants)

	authorized.PUT("/tournament/bracket/update", server.controller.tournament.HandleV2TournamentBracketUpdate)
	authorized.POST("/tournament/bracket/teams/swap", server.controller.tournament.HandleV2TournamentBracketTeamsSwap)
	authorized.PUT("/tournament/round/match/bestofn", server.controller.tournament.HandleV2TournamentRoundMatchBestOfN)
	authorized.PUT("/tournament/round/match/time", server.controller.tournament.HandleV2TournamentRoundMatchTime)

	authorized.POST("/tournament/register", server.controller.tournament.HandleTournamentRegister)
	authorized.POST("/tournament/registerteam", server.controller.tournament.HandleTeamRegister)
	authorized.POST("/tournament/unregister", server.controller.tournament.HandleTournamentUnregister)
	authorized.POST("/tournament/unregisterteam", server.controller.tournament.HandleTeamUnregister)
	authorized.POST("/tournament/fund", server.controller.tournament.HandleV2TournamentFund)
	authorized.POST("/tournament/complete", server.controller.tournament.HandleV2TournamentComplete)
	authorized.POST("/tournament/cancel", server.controller.tournament.HandleV2TournamentCancel)
	authorized.POST("/tournament/start", server.controller.tournament.HandleTournamentStart)

}

func WithV2UnauthorizedManualTournamentRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	r.GET("/tournament/count", server.controller.tournament.HandleV2TournamentCount)
	r.GET("/tournament/list/open", server.controller.tournament.HandleV2TournamentListOpen)
	r.GET("/tournament/list/cutoff", server.controller.tournament.HandleV2TournamentListCutoff)
	r.GET("/tournament/list/active", server.controller.tournament.HandleV2TournamentListActive)
	r.GET("/tournament/list/closed", server.controller.tournament.HandleV2TournamentListClosed)
	r.GET("/tournament/details/:tournamentId", server.controller.tournament.HandleV2TournamentDetails)
	r.GET("/tournament/list/participants", server.controller.tournament.HandleV2TournamentListParticipants)
	r.GET("/tournament/list/participants/:tournamentId", server.controller.tournament.HandleV2TournamentListParticipantsById)
	r.GET("/tournament/list/match/:tournamentId", server.controller.tournament.HandleTournamentListMatch)

}

func WithInvitationRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r

	authorized.POST("/tournament/request/invite", server.controller.invitation.HandleInvitationCreate)
	authorized.PUT("/tournament/request/invite/cancel", server.controller.invitation.HandleInvitationCancel)
	authorized.PUT("/tournament/request/invite/accept", server.controller.invitation.HandleInvitationAccept)
	authorized.PUT("/tournament/request/invite/decline", server.controller.invitation.HandleInvitationDecline)
	authorized.GET("/tournament/request/all/:tournamentId", server.controller.invitation.HandleTournamentAllRequestsFetch)
	authorized.GET("/tournament/request/invite/:tournamentId", server.controller.invitation.HandleTournamentInvitationFetch)
	authorized.GET("/tournament/request/join/:tournamentId", server.controller.application.HandleTournamentListApplications)
	authorized.POST("/tournament/request/join", server.controller.application.HandleApplicationCreate)
	authorized.PUT("/tournament/request/join/cancel", server.controller.application.HandleApplicationCancel)
	authorized.PUT("/tournament/request/join/accept", server.controller.application.HandleTournamentApplicationAccept)
	authorized.PUT("/tournament/request/join/decline", server.controller.application.HandleTournamentApplicationDecline)
	authorized.GET("/tournament/requests/user", server.controller.application.HandleTournamentListUserRequests)

}

func WithAuthorizedChallengeRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r
	//authorized.POST("/challenge/create", server.controller.challenge.HandleChallengeCreate)
	authorized.POST("/challenge/register", server.controller.challenge.HandleChallengeRegister)
	authorized.POST("/challenge/cancel", server.controller.challenge.HandleChallengeCancel)
	authorized.POST("/challenge/complete", server.controller.challenge.HandleChallengeComplete)
	authorized.POST("/challenge/fund", server.controller.challenge.HandleChallengeFund)

	authorized.GET("/challenge/user/fetch/:challengeId", server.controller.challenge.HandleChallengeFetchUserStatus)
	authorized.GET("/challenge/list/mychallenges", server.controller.challenge.HandleMyChallenges)
	authorized.GET("/challenge/enum", server.controller.challenge.HandleChallengeEnum)
}

func WithUnauthorizedChallengeRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	r.GET("/challenge/fetch/:challengeId", server.controller.challenge.HandleChallengeFetch)
	r.GET("/challenge/listSponsored", server.controller.challenge.HandleSponsoredChallengeList)
	r.GET("/challenge/listFeatured", server.controller.challenge.HandleFeaturedChallengeList)
	r.GET("/challenge/featured/:title", server.controller.challenge.HandleFeaturedChallengeFetch)

}

func WithV2UserRoutes(authorized *gin.RouterGroup, server *Server, conf config.Config) {
	authorized.GET("/user/games", server.controller.user.HandleGetUserGames)
	authorized.PUT("/user/game", server.controller.user.HandlePutUserGame)
	authorized.PUT("/user/game/disconnect", server.controller.user.HandlePutDisconnectUserGame)
	authorized.PUT("/user/game/check_username", server.controller.user.HandlePutCheckUsername)
}

func WithUserRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r

	authorized.POST("/urls", server.controller.preSignURL.HandleURLRegister)

	authorized.GET("/user/address", server.controller.user.HandleGetAddress)
	authorized.GET("/user/balance", server.controller.user.HandleGetBalance)
	authorized.GET("/user/gameaccounts", server.controller.user.HandleGetGameAccounts)
	authorized.GET("/user/notification", server.controller.user.HandleNotification)
	authorized.GET("/user/profile", server.controller.user.HandleUserProfileGet)

	authorized.POST("/user/feed/create", server.controller.user.HandleUserFeedCreate)
	authorized.PUT("/user/profile", server.controller.user.HandleUserProfilePut)
	authorized.PUT("/user/profile/uploadAvatar", server.controller.user.HandleUserAvatarUpload)
	authorized.PUT("/user/profile/uploadProfileBanner", server.controller.user.HandleUserProfileBannerUpload)
	authorized.PUT("/user/profile/favouriteGame", server.controller.user.HandleUserFavouriteGameUpdate)
	// authorized.GET("/user/profile/badgeget", server.controller.user.HandleUserBadgeGet)
	authorized.GET("/user/playhistory", server.controller.user.HandlePlayHistoryGet)
	authorized.GET("/user/playhistory/:gameType", server.controller.user.HandlePlayHistoryByGameTypeGet)
	authorized.GET("/user/stats", server.controller.user.HandleUserStatsGet)
	authorized.GET("/user/stats/:gameType", server.controller.user.HandleUserStatsByGameTypeGet)

	authorized.GET("/user/depositrecord", server.controller.user.HandleDepositRecordGet)
	authorized.GET("/user/marketplacerecord", server.controller.user.HandleMarketplaceRecordGet)

	authorized.POST("/user/2fa/send", server.controller.user.HandleUser2FASend)
	authorized.POST("/user/phone/reset", server.controller.user.HandleUserPhoneNumberReset)
}

func WithProfileRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	r.GET("/user/existUserNumber", server.controller.user.HandleFetchExistingUsers)
	r.GET("/otheruser/profile/:displayName", server.controller.user.HandleOtherUserProfileGet)
	r.GET("/otheruser/feed/:displayName", server.controller.user.HandleUserFeedList)
	r.GET("/otheruser/stats/:displayname", server.controller.user.HandleOtherUserStatsGet)
}

func WithFriendsRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r
	authorized.POST("/friend/search", server.controller.user.HandleFindFriend)
	authorized.POST("/friend/request/create", server.controller.user.HandleFriendInvitationCreate)
	authorized.GET("/friend/request/list", server.controller.user.HandleFriendInvitationList)
	authorized.PUT("/friend/request/acceptOrReject", server.controller.user.HandleFriendInvitationAcceptAndReject)
	authorized.GET("/friend/list", server.controller.user.HandleFriendListFetch)
	authorized.GET("/friend/status/:displayName", server.controller.user.HandleFriendStatus)
}

func WithBadgeRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r
	authorized.PUT("/badge/upload", server.controller.badge.HandleBadgeUpload)
	authorized.GET("/badge/list", server.controller.badge.HandleAllBadgeGet)
	authorized.GET("/badge/fetch/:badgeId", server.controller.badge.HandleSpecificBadgeGet)
}

func WithGameRoutes(r *gin.RouterGroup, server *Server, conf config.Config) {
	authorized := r
	authorized.GET("/gameType/list", server.controller.game.HandleGameTypeGet)
	authorized.GET("/gameType/getSubType/:gameType", server.controller.game.HandleGameSubTypeGet)
}

func (server Server) GetAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth with Cognito
		token := c.GetHeader("X-Auth-Token")
		if token == "" {
			c.AbortWithStatusJSON(400, gin.H{"message": "token not found"})
			return
		}

		awsConfig := server.config.AwsConfig()
		tokens, err := server.service.cognitoClient.ParseJWT(token, awsConfig.CognitoRegion, awsConfig.CognitoUserPoolID)

		if err != nil {
			log.Printf("error is %v %+v\n", err, tokens)
			// jwt is not valid
			// c.AbortWithStatusJSON(401, gin.H{"message": "token is not valid"})
			// return
		}

		cognitoSub := tokens.Claims.(jwt.MapClaims)["sub"].(string)
		c.Set("sub", cognitoSub)
		// cognitoUserName := tokens.Claims.(jwt.MapClaims)["cognito:username"]
		// username := cognitoUserName.(string)

		//todo: check this in redis
		// userAccount, err := server.db.UpdateOnlineTime(username)
		// if err != nil || userAccount == nil {
		// 	c.AbortWithError(401, errors.New("bad token"))
		// 	return
		// }

		userAccount, err := server.db.UpdateOnlineTime(cognitoSub)
		if err != nil || userAccount == nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "bad token"})
			return
		}
		c.Set("user", userAccount)
	}
}

func (server Server) Get2FAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.GetHeader("Code")
		codeInt, err := validate2FACode(code)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}

		sub, ok := c.Get("sub")
		if !ok {
			c.AbortWithStatusJSON(500, gin.H{"message": "user sub missing"})
			return
		}
		computedCode := compute2FACode(sub.(string))
		if computedCode == -1 {
			c.AbortWithStatusJSON(500, gin.H{"message": "cannot compute 2FA code"})
			return
		}

		if computedCode != codeInt {
			c.AbortWithStatusJSON(401, gin.H{"message": "wrong 2FA code"})
			return
		}
		c.Next()
	}
}

func (server Server) GetAuthLocal() gin.HandlerFunc {
	return func(c *gin.Context) {
		// auth for local test
		userCognitoMap := make(map[string]interface{})
		userCognitoMap["aud"] = "6p2i7g8t1253r973bqs3ravnld"
		userCognitoMap["auth_time"] = 1.582048118e+09
		userCognitoMap["cognitousername"] = "larry@bitspawn.gg" //eric@bitspawn.gg"
		userCognitoMap["email"] = "larry@bitspawn.gg"
		userCognitoMap["email_verified"] = true
		userCognitoMap["event_id"] = "d7ff5955-9f2b-421f-81ba-214178c99f7d"
		userCognitoMap["exp"] = 1.582051719e+09
		userCognitoMap["iat"] = 1.582048119e+09
		userCognitoMap["iss"] = "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_WeNTmaJ8b"
		userCognitoMap["sub"] = "6d9ba864-32f5-44a4-876b-09ab6573d5d1"

		sub := userCognitoMap["sub"].(string)
		// auth for local test Stop here

		userAccount, err := server.db.UpdateOnlineTime(sub)
		if err != nil || userAccount == nil {
			_ = c.AbortWithError(401, errors.New("bad token"))
			return
		}

		c.Set("user", userAccount)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Auth-Token, Authorization, Code, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT , PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func adminAuth() gin.HandlerFunc {
	accounts := gin.Accounts{
		"larry":     "larrykey",
		"scheduler": "schedulerkey",
	}
	return gin.BasicAuth(accounts)
}
