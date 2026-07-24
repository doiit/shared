package events

type EventType string

// Event Types for various user and system actions
const (
	EventUserLogin          EventType = "user.login"
	EventUserLogout         EventType = "user.logout"
	EventUserRegistered     EventType = "user.registered"
	EventUserPasswordReset  EventType = "user.password_reset"
	EventUserProfileUpdated EventType = "user.profile_updated"
	EventJobCreated         EventType = "job.created"
	EventJobUpdated         EventType = "job.updated"
	EventJobDeleted         EventType = "job.deleted"
	EventJobApplied         EventType = "job.applied"
	EventJobViewed          EventType = "job.viewed"
	EventJobCompleted       EventType = "job.completed"
	EventProposalSent       EventType = "proposal.sent"
	EventProposalAccepted   EventType = "proposal.accepted"
	EventProposalRejected   EventType = "proposal.rejected"
	EventPaymentCompleted   EventType = "payment.completed"
	EventPaymentFailed      EventType = "payment.failed"
	EventMessageSent        EventType = "message.sent"
	EventRatingSubmitted    EventType = "rating.submitted"
	EventUserRegister       EventType = "user.registered"
	EventMessageRead        EventType = "message.read"
	EventUserAnalytics      EventType = "user.analytics"
)
