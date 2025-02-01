package enums

type Feature string

const (
	AnimatedBanner                  Feature = "ANIMATED_BANNER"
	AnimatedIcon                    Feature = "ANIMATED_ICON"
	ApplicationCommandPermissionsV2 Feature = "APPLICATION_COMMAND_PERMISSIONS_V2"
	AutoModeration                  Feature = "AUTO_MODERATION"
	Banner                          Feature = "BANNER"
	Community                       Feature = "COMMUNITY"
	MonetizationEnabled             Feature = "CREATOR_MONETIZABLE_PROVISIONAL"
	StorePageEnabled                Feature = "CREATOR_STORE_PAGE"
	SupportServer                   Feature = "DEVELOPER_SUPPORT_SERVER"
	Discoverable                    Feature = "DISCOVERABLE"
	Featurable                      Feature = "FEATURABLE"
	InvitesDisabled                 Feature = "INVITES_DISABLED"
	InviteSplash                    Feature = "INVITE_SPLASH"
	VerificationGateEnabled         Feature = "MEMBER_VERIFICATION_GATE_ENABLED"
	IncreasedSoundCount             Feature = "MORE_SOUNDBOARD"
	IncreasedStickerCount           Feature = "MORE_STICKERS"
	News                            Feature = "NEWS"
	Partnered                       Feature = "PARTNERED"
	PreviewEnabled                  Feature = "PREVIEW_ENABLED"
	RaidAlertsDisabled              Feature = "RAID_ALERTS_DISABLED"
	RoleIcons                       Feature = "ROLE_ICONS"
	SubscriptionRoles               Feature = "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE"
	SubscriptionRolesEnabled        Feature = "ROLE_SUBSCRIPTIONS_ENABLED"
	HasCreatedSoundboardSounds      Feature = "SOUNDBOARD"
	TicketedEventsEnabled           Feature = "TICKETED_EVENTS_ENABLED"
	CustomURL                       Feature = "VANITY_URL"
	Verified                        Feature = "VERIFIED"
	HasAccessToVIPRegions           Feature = "VIP_REGIONS"
	WelcomeScreenEnabled            Feature = "WELCOME_SCREEN_ENABLED"
)
