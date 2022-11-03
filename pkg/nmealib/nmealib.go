package nmealib

import (
	"fmt"
	"strings"
	"time"
)

// Reference: https://docs.novatel.com/OEM7/Content/Logs/GPZDA.htm
//
// Deprecated: https://github.com/rakiyoshi/raspi-virtual-gps-receiver/issues/9
func FromTimeToZDA(t time.Time) string {
	var data []string

	// 1: Log header
	data = append(data, "$GPZDA")
	// 2: UTC time status, hhmmss.ss
	data = append(data, fmt.Sprintf(
		"%02d%02d%02d.%03d",
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1000/1000,
	))
	// 3: Day, 01 to 31
	data = append(data, fmt.Sprintf("%02d", t.Day()))
	// 4: Month, 01 to 12
	data = append(data, fmt.Sprintf("%02d", t.Month()))
	// 5: Year
	data = append(data, fmt.Sprintf("%04d", t.Year()))
	// 6: null (Local zone description - not available)
	data = append(data, "")
	// 7: null (Local zone minutes description - not available)
	data = append(data, "")
	// 8: checksum (optional)

	return strings.Join(data, ",") + "\r\n" // 9: <CR><LF>
}

// Reference: https://docs.novatel.com/OEM7/Content/Logs/GPGGA.htm
func FromTimeToGGA(t time.Time) string {
	var data []string

	//  1: Log header
	data = append(data, "$GPGGA")
	//  2: UTC time status of position (hhmmss or hhmmss.ss)
	data = append(data, fmt.Sprintf(
		"%02d%02d%02d.%03d",
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1000/1000,
	))
	//  3: Latitude (DDmm.mm)
	data = append(data, "3504.528")
	//  4: Latitude direction (N = North, S = South)
	data = append(data, "N")
	//  5: Longitude (DDDmm.mm)
	data = append(data, "13578.504")
	//  6: Longitude direction (E = East, W = West)
	data = append(data, "E")
	//  7: Quality (Ref: https://docs.novatel.com/OEM7/Content/Logs/GPGGA.htm#GPSQualityIndicators)
	data = append(data, "1")
	//  8: Number of satellites in use. May be different to the number in view
	data = append(data, "08")
	//  9: Horizonal dilution of precision
	data = append(data, "01.96")
	// 10: Antenna altitude above/below mean sea level
	data = append(data, "0013")
	// 11: Units of antenna altitude (M = meters)
	data = append(data, "M")
	// 12: Undulation - the relationship between the geoid and the WGS84 ellipsoid
	data = append(data, "0032")
	// 13: Units of undulation (M = meters)
	data = append(data, "M")
	// 14: Age of correction data (in secconds). The maximum age reported here is limited to 99 seconds.
	data = append(data, "")
	// 15: Differential base station ID
	data = append(data, "")
	// 16: checksum (optional)

	return strings.Join(data, ",") + "\r\n" // 17: <CR><LF>
}

// Reference: https://docs.novatel.com/OEM7/Content/Logs/GPRMC.htm
func FromTimeToRMC(t time.Time) string {
	var data []string

	//  1: Log header
	data = append(data, "$GPRMC")
	//  2: UTC of position
	data = append(data, fmt.Sprintf(
		"%02d%02d%02d.%03d",
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/1000/1000,
	))
	//  3: Position status (A = data valid, V = data invalid)
	data = append(data, "A")
	//  4: Latitude (DDmm.mm)
	data = append(data, "3504.528")
	//  5: Latitude direction (N = North, S = South)
	data = append(data, "N")
	//  6: Longitude (DDDmm.mm)
	data = append(data, "13578.504")
	//  7: Longitude direction (E = East, W = West)
	data = append(data, "E")
	//  8: Speed over ground, knots
	data = append(data, "000.0")
	//  9: Track made good, degrees True
	data = append(data, "240.3")
	// 10: Date (ddmmyy)
	data = append(data, fmt.Sprintf("%02d%02d%02d", t.Day(), t.Month(), t.Year()%100))
	// 11: Magnetic variation, degrees
	data = append(data, "")
	// 12: Magnetic variation direction E/W
	data = append(data, "")
	// 13: Positioning system mode indicator (Ref: https://docs.novatel.com/OEM7/Content/Logs/GPRMC.htm#NMEAPositioningSystemModeIndicator)
	data = append(data, "A")
	// 14: checksum (optional)

	return strings.Join(data, ",") + "\r\n" // 15: <CR><LF>
}
