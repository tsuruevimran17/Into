package models

import "time"

type AuditLog struct {
	Base
	ActorID    uint      `json:"actor_id"`
	ActorRole  Role      `json:"actor_role"`
	Action     string    `json:"action"`
	EntityType string    `json:"entity_type"`
	EntityID   uint      `json:"entity_id"`
	Meta       string    `json:"meta"`
	IP         string    `json:"ip"`
	OccurredAt time.Time `json:"occurred_at"`
}
