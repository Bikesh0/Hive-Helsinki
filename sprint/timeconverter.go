package sprint

func TimeConverter(sec int) (int, int, int) {
	hour := sec / 3600 // 1 hour = 60 minutes × 60 seconds = 3600 seconds.
    min := (sec % 3600)/ 60 // 60 // % gives the remainder after removing the complete ho
	secs := sec % 60 // % 60 gives the seconds left after removing complete minutes.

	return hour, min, secs
	  
}
