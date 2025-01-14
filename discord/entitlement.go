package discord

import (
	"time"

	"github.com/snekROmonoro/snowflake"
)

type Entitlement struct {
	ID            snowflake.ID    `json:"id"`
	SkuID         snowflake.ID    `json:"sku_id"`
	ApplicationID snowflake.ID    `json:"application_id"`
	UserID        *snowflake.ID   `json:"user_id"`
	Type          EntitlementType `json:"type"`
	Deleted       bool            `json:"deleted"`
	StartsAt      *time.Time      `json:"starts_at"`
	EndsAt        *time.Time      `json:"ends_at"`
	GuildID       *snowflake.ID   `json:"guild_id"`
}

type EntitlementType int

const (
	EntitlementTypeApplicationSubscription EntitlementType = 8
)

type TestEntitlementCreate struct {
	SkuID     snowflake.ID         `json:"sku_id"`
	OwnerID   snowflake.ID         `json:"owner_id"`
	OwnerType EntitlementOwnerType `json:"owner_type"`
}

type EntitlementOwnerType int

const (
	EntitlementOwnerTypeGuild EntitlementOwnerType = iota + 1
	EntitlementOwnerTypeUser
)
