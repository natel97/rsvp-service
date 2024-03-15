package notifications

import (
	"encoding/json"
	"fmt"

	webpush "github.com/SherClockHolmes/webpush-go"
)

func NewService(privateVapidKey, publicVapidKey string, repository *repository) *service {
	return &service{
		privateVapidKey: privateVapidKey,
		publicVapidKey:  publicVapidKey,
		repository:      repository,
	}
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package=notifications
type Service interface {
	GetIsSubscribed(id string) bool
	NotifyInvite(kind string, body string, invitationID string)
	RemoveByInvitation(id string) error
	NotifyGroup(subscriptionKind string, kind string, body string, url string)
	RegisterForInvitation(invitationID, subscription, kind string)
	NotifyEvent(eventID string, kind string, body string, url string)
	Notify(subscription string, kind string, body string, url string)
}

type service struct {
	privateVapidKey string
	publicVapidKey  string
	repository      *repository
}

type PushBody struct {
	Kind string `json:"kind"`
	Body string `json:"body"`
	URL  string `json:"url"`
}

func (service *service) GetIsSubscribed(id string) bool {
	inSub := service.repository.GetByInvitation(id)

	return inSub.ID != ""
}

func (service *service) NotifyInvite(kind string, body string, invitationID string) {
	val, err := service.repository.GetSubscriptionByInvitation(invitationID)
	if err != nil {
		fmt.Println("Error notifying invite: ", err)
		return
	}

	service.Notify(val.Subscription, kind, body, "/invitation/"+invitationID)
}

func (service *service) RemoveByInvitation(id string) error {
	err := service.repository.DeleteByInvitation(id)
	return err
}

func (service *service) NotifyGroup(subscriptionKind string, kind string, body string, url string) {
	fmt.Println("Notifying group", subscriptionKind)
	val, _ := service.repository.GetByKind(subscriptionKind)

	for _, sub := range val {
		go service.Notify(sub.Subscription, kind, body, url)
	}
}

func (service *service) RegisterForInvitation(invitationID, subscription, kind string) {
	fmt.Println("Registering for updates by invitation: ", invitationID)
	service.repository.SubscribeToInvitation(invitationID, subscription)
}

func (service *service) NotifyEvent(eventID string, kind string, body string, url string) {
	fmt.Println("Notifying event people: ", eventID)
	val, err := service.repository.GetByEvent(eventID)

	if err != nil {
		fmt.Println(err)
	}

	for _, sub := range val {
		go service.Notify(sub.Subscription, kind, body, "/invitation/"+sub.InvitationID)
	}
}

func (service *service) Notify(subscription string, kind string, body string, url string) {
	fmt.Println("Sending Notification ", kind, "body", body, " url: ", url)
	s := &webpush.Subscription{}
	json.Unmarshal([]byte(subscription), s)

	pushBody := PushBody{
		Kind: kind,
		Body: body,
		URL:  url,
	}

	payload, err := json.Marshal(pushBody)
	if err != nil {
		fmt.Println(payload)
	}

	// Send Notification
	resp, err := webpush.SendNotification(payload, s, &webpush.Options{
		VAPIDPublicKey:  service.publicVapidKey,
		VAPIDPrivateKey: service.privateVapidKey,
		TTL:             30,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
}
