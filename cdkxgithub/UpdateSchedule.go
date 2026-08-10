package cdkxgithub


// Schedule preferences.
// Experimental.
type UpdateSchedule struct {
	// Experimental.
	Interval ScheduleInterval `field:"required" json:"interval" yaml:"interval"`
	// Specify an alternative day to check for updates.
	// Experimental.
	Day ScheduleDay `field:"optional" json:"day" yaml:"day"`
	// Specify an alternative time of day to check for updates (format: hh:mm).
	// Experimental.
	Time *string `field:"optional" json:"time" yaml:"time"`
	// The time zone identifier must be from the Time Zone database maintained by IANA.
	// Experimental.
	Timezone Timezone `field:"optional" json:"timezone" yaml:"timezone"`
}

