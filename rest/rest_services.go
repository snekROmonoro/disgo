package rest

import (
	"net/http"

	"github.com/DisgoOrg/log"
)

var _ Services = (*servicesImpl)(nil)

func NewServices(logger log.Logger, restClient Client) Services {
	if restClient == nil {
		restClient = NewClient(&DefaultConfig)
	}
	return &servicesImpl{
		logger:               logger,
		restClient:           restClient,
		applicationService:   NewApplicationService(restClient),
		oauth2Service:        NewOAuth2Service(restClient),
		auditLogService:      NewAuditLogService(restClient),
		gatewayService:       NewGatewayService(restClient),
		guildService:         NewGuildService(restClient),
		channelService:       NewChannelService(restClient),
		interactionService:   NewInteractionService(restClient),
		inviteService:        NewInviteService(restClient),
		guildTemplateService: NewGuildTemplateService(restClient),
		userService:          NewUserService(restClient),
		voiceService:         NewVoiceService(restClient),
		webhookService:       NewWebhookService(restClient),
		stageInstanceService: NewStageInstanceService(restClient),
		emojiService:         NewEmojiService(restClient),
		stickerService:       NewStickerService(restClient),
	}
}

// Services is a manager for all of disgo's HTTP requests
type Services interface {
	Close()
	Logger() log.Logger
	RestClient() Client
	HTTPClient() *http.Client
	ApplicationService() ApplicationService
	OAuth2Service() OAuth2Service
	AuditLogService() auditLogService
	GatewayService() GatewayService
	GuildService() GuildService
	ChannelService() ChannelService
	InteractionService() InteractionService
	InviteService() InviteService
	GuildTemplateService() GuildTemplateService
	UserService() UserService
	VoiceService() VoiceService
	WebhookService() WebhookService
	StageInstanceService() StageInstanceService
	EmojiService() EmojiService
	StickerService() StickerService
}

type servicesImpl struct {
	logger     log.Logger
	restClient Client

	applicationService   ApplicationService
	oauth2Service        OAuth2Service
	auditLogService      auditLogService
	gatewayService       GatewayService
	guildService         GuildService
	channelService       ChannelService
	interactionService   InteractionService
	inviteService        InviteService
	guildTemplateService GuildTemplateService
	userService          UserService
	voiceService         VoiceService
	webhookService       WebhookService
	stageInstanceService StageInstanceService
	emojiService         EmojiService
	stickerService       StickerService
}

func (s *servicesImpl) Close() {
	s.restClient.Close()
}

func (s *servicesImpl) Logger() log.Logger {
	return s.logger
}

func (s *servicesImpl) RestClient() Client {
	return s.restClient
}

func (s *servicesImpl) HTTPClient() *http.Client {
	return s.RestClient().HTTPClient()
}

func (s *servicesImpl) ApplicationService() ApplicationService {
	return s.applicationService
}

func (s *servicesImpl) OAuth2Service() OAuth2Service {
	return s.oauth2Service
}

func (s *servicesImpl) AuditLogService() auditLogService {
	return s.auditLogService
}

func (s *servicesImpl) GatewayService() GatewayService {
	return s.gatewayService
}

func (s *servicesImpl) GuildService() GuildService {
	return s.guildService
}

func (s *servicesImpl) ChannelService() ChannelService {
	return s.channelService
}

func (s *servicesImpl) InteractionService() InteractionService {
	return s.interactionService
}

func (s *servicesImpl) InviteService() InviteService {
	return s.inviteService
}

func (s *servicesImpl) GuildTemplateService() GuildTemplateService {
	return s.guildTemplateService
}

func (s *servicesImpl) UserService() UserService {
	return s.userService
}

func (s *servicesImpl) VoiceService() VoiceService {
	return s.voiceService
}

func (s *servicesImpl) WebhookService() WebhookService {
	return s.webhookService
}

func (s *servicesImpl) StageInstanceService() StageInstanceService {
	return s.stageInstanceService
}

func (s *servicesImpl) EmojiService() EmojiService {
	return s.emojiService
}

func (s *servicesImpl) StickerService() StickerService {
	return s.stickerService
}

type Service interface {
	RestClient() Client
}