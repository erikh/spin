package supervisor

// Interface abstracts out the supervisor. This will eventually encapsulate all
// of systemd, launchd etc configuration and management.
type Interface interface {
	// Review all configuration and start services as necessary
	Reload() error

	// Start a service
	Start(string) error

	// Stop a service
	Stop(string) error
}
