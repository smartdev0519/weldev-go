// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// generated by genzabbrs.go from
// http://unicode.org/cldr/data/common/supplemental/windowsZones.xml

package time

type abbr struct {
	std string
	dst string
}

var abbrs = map[string]abbr{
	"Egypt Standard Time":             {"EET", "EET"},    // Africa/Cairo
	"Morocco Standard Time":           {"WET", "WEST"},   // Africa/Casablanca
	"South Africa Standard Time":      {"SAST", "SAST"},  // Africa/Johannesburg
	"W. Central Africa Standard Time": {"WAT", "WAT"},    // Africa/Lagos
	"E. Africa Standard Time":         {"EAT", "EAT"},    // Africa/Nairobi
	"Namibia Standard Time":           {"WAT", "WAST"},   // Africa/Windhoek
	"Alaskan Standard Time":           {"AKST", "AKDT"},  // America/Anchorage
	"Paraguay Standard Time":          {"PYT", "PYST"},   // America/Asuncion
	"Bahia Standard Time":             {"BRT", "BRST"},   // America/Bahia
	"SA Pacific Standard Time":        {"COT", "COT"},    // America/Bogota
	"Argentina Standard Time":         {"ART", "ART"},    // America/Buenos_Aires
	"Venezuela Standard Time":         {"VET", "VET"},    // America/Caracas
	"SA Eastern Standard Time":        {"GFT", "GFT"},    // America/Cayenne
	"Central Standard Time":           {"CST", "CDT"},    // America/Chicago
	"Mountain Standard Time (Mexico)": {"MST", "MDT"},    // America/Chihuahua
	"Central Brazilian Standard Time": {"AMT", "AMST"},   // America/Cuiaba
	"Mountain Standard Time":          {"MST", "MDT"},    // America/Denver
	"Greenland Standard Time":         {"WGT", "WGST"},   // America/Godthab
	"Central America Standard Time":   {"CST", "CST"},    // America/Guatemala
	"Atlantic Standard Time":          {"AST", "ADT"},    // America/Halifax
	"US Eastern Standard Time":        {"EST", "EDT"},    // America/Indianapolis
	"SA Western Standard Time":        {"BOT", "BOT"},    // America/La_Paz
	"Pacific Standard Time":           {"PST", "PDT"},    // America/Los_Angeles
	"Central Standard Time (Mexico)":  {"CST", "CDT"},    // America/Mexico_City
	"Montevideo Standard Time":        {"UYT", "UYST"},   // America/Montevideo
	"Eastern Standard Time":           {"EST", "EDT"},    // America/New_York
	"US Mountain Standard Time":       {"MST", "MST"},    // America/Phoenix
	"Canada Central Standard Time":    {"CST", "CST"},    // America/Regina
	"Pacific Standard Time (Mexico)":  {"PST", "PDT"},    // America/Santa_Isabel
	"Pacific SA Standard Time":        {"CLT", "CLST"},   // America/Santiago
	"E. South America Standard Time":  {"BRT", "BRST"},   // America/Sao_Paulo
	"Newfoundland Standard Time":      {"NST", "NDT"},    // America/St_Johns
	"Central Asia Standard Time":      {"ALMT", "ALMT"},  // Asia/Almaty
	"Jordan Standard Time":            {"EET", "EEST"},   // Asia/Amman
	"Arabic Standard Time":            {"AST", "AST"},    // Asia/Baghdad
	"Azerbaijan Standard Time":        {"AZT", "AZST"},   // Asia/Baku
	"SE Asia Standard Time":           {"ICT", "ICT"},    // Asia/Bangkok
	"Middle East Standard Time":       {"EET", "EEST"},   // Asia/Beirut
	"India Standard Time":             {"IST", "IST"},    // Asia/Calcutta
	"Sri Lanka Standard Time":         {"IST", "IST"},    // Asia/Colombo
	"Syria Standard Time":             {"EET", "EEST"},   // Asia/Damascus
	"Bangladesh Standard Time":        {"BDT", "BDT"},    // Asia/Dhaka
	"Arabian Standard Time":           {"GST", "GST"},    // Asia/Dubai
	"North Asia East Standard Time":   {"IRKT", "IRKT"},  // Asia/Irkutsk
	"Israel Standard Time":            {"IST", "IDT"},    // Asia/Jerusalem
	"Afghanistan Standard Time":       {"AFT", "AFT"},    // Asia/Kabul
	"Pakistan Standard Time":          {"PKT", "PKT"},    // Asia/Karachi
	"Nepal Standard Time":             {"NPT", "NPT"},    // Asia/Katmandu
	"North Asia Standard Time":        {"KRAT", "KRAT"},  // Asia/Krasnoyarsk
	"Magadan Standard Time":           {"MAGT", "MAGT"},  // Asia/Magadan
	"E. Europe Standard Time":         {"EET", "EEST"},   // Asia/Nicosia
	"N. Central Asia Standard Time":   {"NOVT", "NOVT"},  // Asia/Novosibirsk
	"Myanmar Standard Time":           {"MMT", "MMT"},    // Asia/Rangoon
	"Arab Standard Time":              {"AST", "AST"},    // Asia/Riyadh
	"Korea Standard Time":             {"KST", "KST"},    // Asia/Seoul
	"China Standard Time":             {"CST", "CST"},    // Asia/Shanghai
	"Singapore Standard Time":         {"SGT", "SGT"},    // Asia/Singapore
	"Taipei Standard Time":            {"CST", "CST"},    // Asia/Taipei
	"West Asia Standard Time":         {"UZT", "UZT"},    // Asia/Tashkent
	"Georgian Standard Time":          {"GET", "GET"},    // Asia/Tbilisi
	"Iran Standard Time":              {"IRST", "IRDT"},  // Asia/Tehran
	"Tokyo Standard Time":             {"JST", "JST"},    // Asia/Tokyo
	"Ulaanbaatar Standard Time":       {"ULAT", "ULAT"},  // Asia/Ulaanbaatar
	"Vladivostok Standard Time":       {"VLAT", "VLAT"},  // Asia/Vladivostok
	"Yakutsk Standard Time":           {"YAKT", "YAKT"},  // Asia/Yakutsk
	"Ekaterinburg Standard Time":      {"YEKT", "YEKT"},  // Asia/Yekaterinburg
	"Caucasus Standard Time":          {"AMT", "AMT"},    // Asia/Yerevan
	"Azores Standard Time":            {"AZOT", "AZOST"}, // Atlantic/Azores
	"Cape Verde Standard Time":        {"CVT", "CVT"},    // Atlantic/Cape_Verde
	"Greenwich Standard Time":         {"GMT", "GMT"},    // Atlantic/Reykjavik
	"Cen. Australia Standard Time":    {"CST", "CST"},    // Australia/Adelaide
	"E. Australia Standard Time":      {"EST", "EST"},    // Australia/Brisbane
	"AUS Central Standard Time":       {"CST", "CST"},    // Australia/Darwin
	"Tasmania Standard Time":          {"EST", "EST"},    // Australia/Hobart
	"W. Australia Standard Time":      {"WST", "WST"},    // Australia/Perth
	"AUS Eastern Standard Time":       {"EST", "EST"},    // Australia/Sydney
	"UTC":                            {"GMT", "GMT"},       // Etc/GMT
	"UTC-11":                         {"GMT+11", "GMT+11"}, // Etc/GMT+11
	"Dateline Standard Time":         {"GMT+12", "GMT+12"}, // Etc/GMT+12
	"UTC-02":                         {"GMT+2", "GMT+2"},   // Etc/GMT+2
	"UTC+12":                         {"GMT-12", "GMT-12"}, // Etc/GMT-12
	"W. Europe Standard Time":        {"CET", "CEST"},      // Europe/Berlin
	"GTB Standard Time":              {"EET", "EEST"},      // Europe/Bucharest
	"Central Europe Standard Time":   {"CET", "CEST"},      // Europe/Budapest
	"Turkey Standard Time":           {"EET", "EEST"},      // Europe/Istanbul
	"Kaliningrad Standard Time":      {"FET", "FET"},       // Europe/Kaliningrad
	"FLE Standard Time":              {"EET", "EEST"},      // Europe/Kiev
	"GMT Standard Time":              {"GMT", "BST"},       // Europe/London
	"Russian Standard Time":          {"MSK", "MSK"},       // Europe/Moscow
	"Romance Standard Time":          {"CET", "CEST"},      // Europe/Paris
	"Central European Standard Time": {"CET", "CEST"},      // Europe/Warsaw
	"Mauritius Standard Time":        {"MUT", "MUT"},       // Indian/Mauritius
	"Samoa Standard Time":            {"WST", "WST"},       // Pacific/Apia
	"New Zealand Standard Time":      {"NZST", "NZDT"},     // Pacific/Auckland
	"Fiji Standard Time":             {"FJT", "FJT"},       // Pacific/Fiji
	"Central Pacific Standard Time":  {"SBT", "SBT"},       // Pacific/Guadalcanal
	"Hawaiian Standard Time":         {"HST", "HST"},       // Pacific/Honolulu
	"West Pacific Standard Time":     {"PGT", "PGT"},       // Pacific/Port_Moresby
	"Tonga Standard Time":            {"TOT", "TOT"},       // Pacific/Tongatapu
}
