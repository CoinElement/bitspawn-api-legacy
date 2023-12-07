package tournament

type Action string

const (
	DEPLOY     Action = "Deploy"
	REGISTER   Action = "Register"
	UNREGISTER Action = "Unregister"
	CANCEL     Action = "Cancel"
	FUND       Action = "Fund"
	COMPLETE   Action = "Complete"
)

type Status string

const (
	DRAFT        Status = "DRAFT"
	REGISTRATION Status = "REGISTRATION"
	STARTED      Status = "STARTED"
	CANCELLED    Status = "CANCELLED"
	PAYOUT       Status = "PAYOUT"
	COMPLETED    Status = "COMPLETED"
)
