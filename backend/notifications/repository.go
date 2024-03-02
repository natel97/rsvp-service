package notifications

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	ID           string
	Subscription string
	Kind         string
}

type InvitationSubscription struct {
	gorm.Model
	ID             string
	SubscriptionID string
	InvitationID   string
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) GetByInvitation(id string) InvitationSubscription {
	inSub := InvitationSubscription{}

	err := repo.db.Find(&inSub, "invitation_id = ? AND deleted_at IS NULL", id).Error
	if err != nil {
		fmt.Println(err)
	}

	return inSub
}
func (repo *Repository) DeleteByInvitation(invitationID string) error {
	item := InvitationSubscription{}
	err := repo.db.Delete(&item, "invitation_id = ?", invitationID).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Create(subscription string, kind string) error {
	notify := &Subscription{Kind: kind, Subscription: subscription, ID: uuid.New().String()}

	err := repo.db.Create(notify).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) GetAll() ([]Subscription, error) {
	subscriptions := []Subscription{}
	err := repo.db.Find(&subscriptions).Error
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (repo *Repository) GetByKind(kind string) ([]Subscription, error) {
	subscriptions := []Subscription{}
	err := repo.db.Find(&subscriptions, "kind = ?", kind).Error
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (repo *Repository) SubscribeToInvitation(invitationID string, subscription string) {
	sub := Subscription{}
	err := repo.db.Find(&sub, "subscription = ?", subscription).Error
	if err != nil {
		fmt.Println(err)
	}

	if sub.ID == "" {
		sub.Subscription = subscription
		sub.Kind = "invited"
		sub.ID = uuid.NewString()
		repo.db.Create(&sub)
	}

	invitationSubscription := InvitationSubscription{
		SubscriptionID: sub.ID,
		InvitationID:   invitationID,
	}

	err = repo.db.Table("invitation_subscriptions").
		Where(invitationSubscription).
		Attrs(InvitationSubscription{ID: uuid.NewString()}).
		FirstOrCreate(&invitationSubscription).Error

	if err != nil {
		fmt.Println(err)
	}
}

type SubscriptionWithInvitation struct {
	Subscription,
	InvitationID string
}

func (repo *Repository) GetByEvent(eventID string) ([]SubscriptionWithInvitation, error) {
	subscriptions := []SubscriptionWithInvitation{}
	err := repo.db.Raw(`
	SELECT s.*, inv_sub.invitation_id as invitation_id
	FROM subscriptions s
	LEFT JOIN (
		SELECT *
		FROM invitation_subscriptions
		WHERE invitation_id IN (
			SELECT id
			FROM invitations
			WHERE event_id = ?
			AND deleted_at IS NULL
		)
		AND deleted_at IS NULL
	) inv_sub
	ON inv_sub.subscription_id = s.id
	WHERE inv_sub.subscription_id IS NOT NULL
	OR s.kind = "admin"
	AND inv_sub.deleted_at IS NULL
	AND s.deleted_at IS NULL
	`, eventID).Scan(&subscriptions).Error
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (repo *Repository) GetSubscriptionByInvitation(inviteID string) (*Subscription, error) {
	sub := Subscription{}
	inviteSub := InvitationSubscription{}
	err := repo.db.First(&inviteSub, "invitation_id = ?", inviteID).Error

	if err != nil {
		return nil, err
	}

	err = repo.db.First(&sub, "id = ?", inviteSub.SubscriptionID).Error
	if err != nil {
		return nil, err
	}

	return &sub, nil
}
