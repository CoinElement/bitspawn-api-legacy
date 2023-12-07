package server

import (
	"io"

	"github.com/bitspawngg/bitspawn-api/controllers"
	"github.com/bitspawngg/bitspawn-api/services/challenge"
	"github.com/bitspawngg/bitspawn-api/services/cognito"
	"github.com/bitspawngg/bitspawn-api/services/config"
	organizer "github.com/bitspawngg/bitspawn-api/services/match"
	"github.com/bitspawngg/bitspawn-api/services/poa"
	"github.com/bitspawngg/bitspawn-api/services/processor"
	"github.com/bitspawngg/bitspawn-api/services/queue"
	"github.com/bitspawngg/bitspawn-api/services/s3"
	"github.com/bitspawngg/bitspawn-api/services/signURL"
	"github.com/bitspawngg/bitspawn-api/services/tournament"
	userdata "github.com/bitspawngg/bitspawn-api/services/user"
)

type Service struct {
	config            *config.Service
	poaClient         *poa.BitspawnPoaClient
	tournamentService *tournament.TournamentService
	organizerService  *organizer.MatchService
	challengeService  *challenge.ChallengeService
	cognitoClient     *cognito.Auth
	s3UploadClient    *s3.S3UploadClient
	signURLManager    *signURL.SignURLManager
	sqsService        *queue.SQSService
	userService       *userdata.UserService
	txProcessor       *processor.TxProcessor
}

type Controller struct {
	base        *controllers.BaseController
	tournament  *controllers.TournamentController
	challenge   *controllers.ChallengeController
	user        *controllers.UserController
	badge       *controllers.BadgeController
	game        *controllers.GameController
	invitation  *controllers.InvitationController
	application *controllers.ApplicationController
	team        *controllers.TeamController
	preSignURL  *controllers.PreSignedController
	social      *controllers.SocialController
}

func (server *Server) NewService() []io.Closer {
	var service Service

	service.config = config.NewService(server.db, server.log)
	service.poaClient = poa.NewBitspawnPoaClient(server.log, service.config)
	service.organizerService = organizer.NewMatchService(server.log, server.db)
	service.challengeService = challenge.NewChallengeService(server.db, server.log, service.poaClient, service.config)
	awsConfig := server.config.AwsConfig()
	service.cognitoClient = cognito.NewAuth(awsConfig.CognitoRegion, awsConfig.CognitoUserPoolID)
	service.s3UploadClient = s3.NewS3UploadClient(awsConfig.CognitoRegion, awsConfig.S3BucketName)
	service.signURLManager = signURL.NewSignURLManager(awsConfig.CognitoRegion, server.log, awsConfig.S3BucketName)
	service.sqsService = queue.NewSQSService(server.log, awsConfig)
	service.tournamentService = tournament.NewTournamentService(server.db, server.log, service.poaClient, service.sqsService, server.config)
	service.txProcessor = processor.NewTxProcessor(server.db, server.log, awsConfig, service.tournamentService, service.sqsService)
	// add all service that need to be closed
	toClose := []io.Closer{
		service.txProcessor,
	}
	server.service = &service
	return toClose
}

func (server Server) NewController() *Controller {
	var controller Controller
	controller.base = controllers.NewBaseController("base", server.db, server.log, *server.config)
	controller.tournament = controllers.NewTournamentController(controller.base, server.service.tournamentService, server.service.organizerService, server.service.poaClient, server.service.cognitoClient, server.service.s3UploadClient, server.service.signURLManager, server.service.sqsService)
	controller.user = controllers.NewUserController(controller.base, server.service.poaClient, server.service.cognitoClient, server.service.s3UploadClient, server.service.sqsService, server.service.userService)
	controller.challenge = controllers.NewChallengeController(controller.base, server.service.challengeService, server.service.poaClient)
	controller.game = controllers.NewGameController(controller.base)
	controller.badge = controllers.NewBadgeController(controller.base, server.service.s3UploadClient)
	controller.team = controllers.NewTeamController(controller.base, server.service.s3UploadClient, server.service.poaClient)
	controller.invitation = controllers.NewInvitationController(controller.tournament)
	controller.application = controllers.NewApplicationController(controller.tournament)
	controller.preSignURL = controllers.NewPreSignedController(controller.base, server.service.signURLManager)
	controller.social = controllers.NewSocialController(controller.base)
	return &controller
}
