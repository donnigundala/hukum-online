package models

// Invitation is an entity that represents an invitation.
type Invitation struct {
	ID          uint `json:"id" gorm:"primary_key"`
	MemberID    uint `json:"member_id" gorm:"primaryKey"`
	GatheringID uint `json:"gathering_id" gorm:"primaryKey"`
	StatusID    uint `json:"status_id"  gorm:"primaryKey"`

	// BelongsTo
	Member    Member           `json:"member" gorm:"references:ID"`
	Gathering Gathering        `json:"gathering" gorm:"references:ID"`
	Status    InvitationStatus `json:"status" gorm:"references:ID"`
}
