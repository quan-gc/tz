package tz

import "sync"

// GENERATED FILE DO NOT MODIFY DIRECTLY

var (
	once      sync.Once
	mapped    map[string]Country
	countries = []Country{
		{
			Code: "AF",
			Name: "Afghanistan",
			Zones: []Zone{
				{
					CountryCode: "AF",
					Name:        "Asia/Kabul",
					City:        "kabul",
				},
			},
		},
		{
			Code: "AL",
			Name: "Albania",
			Zones: []Zone{
				{
					CountryCode: "AL",
					Name:        "Europe/Tirane",
					City:        "tirane",
				},
			},
		},
		{
			Code: "DZ",
			Name: "Algeria",
			Zones: []Zone{
				{
					CountryCode: "DZ",
					Name:        "Africa/Algiers",
					City:        "algiers",
				},
			},
		},
		{
			Code: "AS",
			Name: "American Samoa",
			Zones: []Zone{
				{
					CountryCode: "AS",
					Name:        "Pacific/Pago_Pago",
					City:        "pago pago",
				},
			},
		},
		{
			Code: "AD",
			Name: "Andorra",
			Zones: []Zone{
				{
					CountryCode: "AD",
					Name:        "Europe/Andorra",
					City:        "andorra",
				},
			},
		},
		{
			Code: "AO",
			Name: "Angola",
			Zones: []Zone{
				{
					CountryCode: "AO",
					Name:        "Africa/Luanda",
					City:        "luanda",
				},
			},
		},
		{
			Code: "AI",
			Name: "Anguilla",
			Zones: []Zone{
				{
					CountryCode: "AI",
					Name:        "America/Anguilla",
					City:        "anguilla",
				},
			},
		},
		{
			Code: "AQ",
			Name: "Antarctica",
			Zones: []Zone{
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Casey",
					City:        "casey",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Davis",
					City:        "davis",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/DumontDUrville",
					City:        "dumontdurville",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Mawson",
					City:        "mawson",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/McMurdo",
					City:        "mcmurdo",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Palmer",
					City:        "palmer",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Rothera",
					City:        "rothera",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Syowa",
					City:        "syowa",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Troll",
					City:        "troll",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Vostok",
					City:        "vostok",
				},
			},
		},
		{
			Code: "AG",
			Name: "Antigua and Barbuda",
			Zones: []Zone{
				{
					CountryCode: "AG",
					Name:        "America/Antigua",
					City:        "antigua",
				},
			},
		},
		{
			Code: "AR",
			Name: "Argentina",
			Zones: []Zone{
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Buenos_Aires",
					City:        "buenos aires",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Catamarca",
					City:        "catamarca",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Cordoba",
					City:        "cordoba",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Jujuy",
					City:        "jujuy",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/La_Rioja",
					City:        "la rioja",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Mendoza",
					City:        "mendoza",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Rio_Gallegos",
					City:        "rio gallegos",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Salta",
					City:        "salta",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/San_Juan",
					City:        "san juan",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/San_Luis",
					City:        "san luis",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Tucuman",
					City:        "tucuman",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Ushuaia",
					City:        "ushuaia",
				},
			},
		},
		{
			Code: "AM",
			Name: "Armenia",
			Zones: []Zone{
				{
					CountryCode: "AM",
					Name:        "Asia/Yerevan",
					City:        "yerevan",
				},
			},
		},
		{
			Code: "AW",
			Name: "Aruba",
			Zones: []Zone{
				{
					CountryCode: "AW",
					Name:        "America/Aruba",
					City:        "aruba",
				},
			},
		},
		{
			Code: "AU",
			Name: "Australia",
			Zones: []Zone{
				{
					CountryCode: "AU",
					Name:        "Antarctica/Macquarie",
					City:        "macquarie",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Adelaide",
					City:        "adelaide",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Brisbane",
					City:        "brisbane",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Broken_Hill",
					City:        "broken hill",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Currie",
					City:        "currie",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Darwin",
					City:        "darwin",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Eucla",
					City:        "eucla",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Hobart",
					City:        "hobart",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Lindeman",
					City:        "lindeman",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Lord_Howe",
					City:        "lord howe",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Melbourne",
					City:        "melbourne",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Perth",
					City:        "perth",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Sydney",
					City:        "sydney",
				},
			},
		},
		{
			Code: "AT",
			Name: "Austria",
			Zones: []Zone{
				{
					CountryCode: "AT",
					Name:        "Europe/Vienna",
					City:        "vienna",
				},
			},
		},
		{
			Code: "AZ",
			Name: "Azerbaijan",
			Zones: []Zone{
				{
					CountryCode: "AZ",
					Name:        "Asia/Baku",
					City:        "baku",
				},
			},
		},
		{
			Code: "BS",
			Name: "Bahamas",
			Zones: []Zone{
				{
					CountryCode: "BS",
					Name:        "America/Nassau",
					City:        "nassau",
				},
			},
		},
		{
			Code: "BH",
			Name: "Bahrain",
			Zones: []Zone{
				{
					CountryCode: "BH",
					Name:        "Asia/Bahrain",
					City:        "bahrain",
				},
			},
		},
		{
			Code: "BD",
			Name: "Bangladesh",
			Zones: []Zone{
				{
					CountryCode: "BD",
					Name:        "Asia/Dhaka",
					City:        "dhaka",
				},
			},
		},
		{
			Code: "BB",
			Name: "Barbados",
			Zones: []Zone{
				{
					CountryCode: "BB",
					Name:        "America/Barbados",
					City:        "barbados",
				},
			},
		},
		{
			Code: "BY",
			Name: "Belarus",
			Zones: []Zone{
				{
					CountryCode: "BY",
					Name:        "Europe/Minsk",
					City:        "minsk",
				},
			},
		},
		{
			Code: "BE",
			Name: "Belgium",
			Zones: []Zone{
				{
					CountryCode: "BE",
					Name:        "Europe/Brussels",
					City:        "brussels",
				},
			},
		},
		{
			Code: "BZ",
			Name: "Belize",
			Zones: []Zone{
				{
					CountryCode: "BZ",
					Name:        "America/Belize",
					City:        "belize",
				},
			},
		},
		{
			Code: "BJ",
			Name: "Benin",
			Zones: []Zone{
				{
					CountryCode: "BJ",
					Name:        "Africa/Porto-Novo",
					City:        "porto-novo",
				},
			},
		},
		{
			Code: "BM",
			Name: "Bermuda",
			Zones: []Zone{
				{
					CountryCode: "BM",
					Name:        "Atlantic/Bermuda",
					City:        "bermuda",
				},
			},
		},
		{
			Code: "BT",
			Name: "Bhutan",
			Zones: []Zone{
				{
					CountryCode: "BT",
					Name:        "Asia/Thimphu",
					City:        "thimphu",
				},
			},
		},
		{
			Code: "BO",
			Name: "Bolivia, Plurinational State of",
			Zones: []Zone{
				{
					CountryCode: "BO",
					Name:        "America/La_Paz",
					City:        "la paz",
				},
			},
		},
		{
			Code: "BQ",
			Name: "Bonaire, Sint Eustatius and Saba",
			Zones: []Zone{
				{
					CountryCode: "BQ",
					Name:        "America/Kralendijk",
					City:        "kralendijk",
				},
			},
		},
		{
			Code: "BA",
			Name: "Bosnia and Herzegovina",
			Zones: []Zone{
				{
					CountryCode: "BA",
					Name:        "Europe/Sarajevo",
					City:        "sarajevo",
				},
			},
		},
		{
			Code: "BW",
			Name: "Botswana",
			Zones: []Zone{
				{
					CountryCode: "BW",
					Name:        "Africa/Gaborone",
					City:        "gaborone",
				},
			},
		},
		{
			Code:  "BV",
			Name:  "Bouvet Island",
			Zones: []Zone{},
		},
		{
			Code: "BR",
			Name: "Brazil",
			Zones: []Zone{
				{
					CountryCode: "BR",
					Name:        "America/Araguaina",
					City:        "araguaina",
				},
				{
					CountryCode: "BR",
					Name:        "America/Bahia",
					City:        "bahia",
				},
				{
					CountryCode: "BR",
					Name:        "America/Belem",
					City:        "belem",
				},
				{
					CountryCode: "BR",
					Name:        "America/Boa_Vista",
					City:        "boa vista",
				},
				{
					CountryCode: "BR",
					Name:        "America/Campo_Grande",
					City:        "campo grande",
				},
				{
					CountryCode: "BR",
					Name:        "America/Cuiaba",
					City:        "cuiaba",
				},
				{
					CountryCode: "BR",
					Name:        "America/Eirunepe",
					City:        "eirunepe",
				},
				{
					CountryCode: "BR",
					Name:        "America/Fortaleza",
					City:        "fortaleza",
				},
				{
					CountryCode: "BR",
					Name:        "America/Maceio",
					City:        "maceio",
				},
				{
					CountryCode: "BR",
					Name:        "America/Manaus",
					City:        "manaus",
				},
				{
					CountryCode: "BR",
					Name:        "America/Noronha",
					City:        "noronha",
				},
				{
					CountryCode: "BR",
					Name:        "America/Porto_Velho",
					City:        "porto velho",
				},
				{
					CountryCode: "BR",
					Name:        "America/Recife",
					City:        "recife",
				},
				{
					CountryCode: "BR",
					Name:        "America/Rio_Branco",
					City:        "rio branco",
				},
				{
					CountryCode: "BR",
					Name:        "America/Santarem",
					City:        "santarem",
				},
				{
					CountryCode: "BR",
					Name:        "America/Sao_Paulo",
					City:        "sao paulo",
				},
			},
		},
		{
			Code: "IO",
			Name: "British Indian Ocean Territory",
			Zones: []Zone{
				{
					CountryCode: "IO",
					Name:        "Indian/Chagos",
					City:        "chagos",
				},
			},
		},
		{
			Code: "BN",
			Name: "Brunei Darussalam",
			Zones: []Zone{
				{
					CountryCode: "BN",
					Name:        "Asia/Brunei",
					City:        "brunei",
				},
			},
		},
		{
			Code: "BG",
			Name: "Bulgaria",
			Zones: []Zone{
				{
					CountryCode: "BG",
					Name:        "Europe/Sofia",
					City:        "sofia",
				},
			},
		},
		{
			Code: "BF",
			Name: "Burkina Faso",
			Zones: []Zone{
				{
					CountryCode: "BF",
					Name:        "Africa/Ouagadougou",
					City:        "ouagadougou",
				},
			},
		},
		{
			Code: "BI",
			Name: "Burundi",
			Zones: []Zone{
				{
					CountryCode: "BI",
					Name:        "Africa/Bujumbura",
					City:        "bujumbura",
				},
			},
		},
		{
			Code: "KH",
			Name: "Cambodia",
			Zones: []Zone{
				{
					CountryCode: "KH",
					Name:        "Asia/Phnom_Penh",
					City:        "phnom penh",
				},
			},
		},
		{
			Code: "CM",
			Name: "Cameroon",
			Zones: []Zone{
				{
					CountryCode: "CM",
					Name:        "Africa/Douala",
					City:        "douala",
				},
			},
		},
		{
			Code: "CA",
			Name: "Canada",
			Zones: []Zone{
				{
					CountryCode: "CA",
					Name:        "America/Atikokan",
					City:        "atikokan",
				},
				{
					CountryCode: "CA",
					Name:        "America/Blanc-Sablon",
					City:        "blanc-sablon",
				},
				{
					CountryCode: "CA",
					Name:        "America/Cambridge_Bay",
					City:        "cambridge bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Creston",
					City:        "creston",
				},
				{
					CountryCode: "CA",
					Name:        "America/Dawson",
					City:        "dawson",
				},
				{
					CountryCode: "CA",
					Name:        "America/Dawson_Creek",
					City:        "dawson creek",
				},
				{
					CountryCode: "CA",
					Name:        "America/Edmonton",
					City:        "edmonton",
				},
				{
					CountryCode: "CA",
					Name:        "America/Fort_Nelson",
					City:        "fort nelson",
				},
				{
					CountryCode: "CA",
					Name:        "America/Glace_Bay",
					City:        "glace bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Goose_Bay",
					City:        "goose bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Halifax",
					City:        "halifax",
				},
				{
					CountryCode: "CA",
					Name:        "America/Inuvik",
					City:        "inuvik",
				},
				{
					CountryCode: "CA",
					Name:        "America/Iqaluit",
					City:        "iqaluit",
				},
				{
					CountryCode: "CA",
					Name:        "America/Moncton",
					City:        "moncton",
				},
				{
					CountryCode: "CA",
					Name:        "America/Nipigon",
					City:        "nipigon",
				},
				{
					CountryCode: "CA",
					Name:        "America/Pangnirtung",
					City:        "pangnirtung",
				},
				{
					CountryCode: "CA",
					Name:        "America/Rainy_River",
					City:        "rainy river",
				},
				{
					CountryCode: "CA",
					Name:        "America/Rankin_Inlet",
					City:        "rankin inlet",
				},
				{
					CountryCode: "CA",
					Name:        "America/Regina",
					City:        "regina",
				},
				{
					CountryCode: "CA",
					Name:        "America/Resolute",
					City:        "resolute",
				},
				{
					CountryCode: "CA",
					Name:        "America/St_Johns",
					City:        "st johns",
				},
				{
					CountryCode: "CA",
					Name:        "America/Swift_Current",
					City:        "swift current",
				},
				{
					CountryCode: "CA",
					Name:        "America/Thunder_Bay",
					City:        "thunder bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Toronto",
					City:        "toronto",
				},
				{
					CountryCode: "CA",
					Name:        "America/Vancouver",
					City:        "vancouver",
				},
				{
					CountryCode: "CA",
					Name:        "America/Whitehorse",
					City:        "whitehorse",
				},
				{
					CountryCode: "CA",
					Name:        "America/Winnipeg",
					City:        "winnipeg",
				},
				{
					CountryCode: "CA",
					Name:        "America/Yellowknife",
					City:        "yellowknife",
				},
			},
		},
		{
			Code: "CV",
			Name: "Cape Verde",
			Zones: []Zone{
				{
					CountryCode: "CV",
					Name:        "Atlantic/Cape_Verde",
					City:        "cape verde",
				},
			},
		},
		{
			Code: "KY",
			Name: "Cayman Islands",
			Zones: []Zone{
				{
					CountryCode: "KY",
					Name:        "America/Cayman",
					City:        "cayman",
				},
			},
		},
		{
			Code: "CF",
			Name: "Central African Republic",
			Zones: []Zone{
				{
					CountryCode: "CF",
					Name:        "Africa/Bangui",
					City:        "bangui",
				},
			},
		},
		{
			Code: "TD",
			Name: "Chad",
			Zones: []Zone{
				{
					CountryCode: "TD",
					Name:        "Africa/Ndjamena",
					City:        "ndjamena",
				},
			},
		},
		{
			Code: "CL",
			Name: "Chile",
			Zones: []Zone{
				{
					CountryCode: "CL",
					Name:        "America/Punta_Arenas",
					City:        "punta arenas",
				},
				{
					CountryCode: "CL",
					Name:        "America/Santiago",
					City:        "santiago",
				},
				{
					CountryCode: "CL",
					Name:        "Pacific/Easter",
					City:        "easter",
				},
			},
		},
		{
			Code: "CN",
			Name: "China",
			Zones: []Zone{
				{
					CountryCode: "CN",
					Name:        "Asia/Shanghai",
					City:        "shanghai",
				},
				{
					CountryCode: "CN",
					Name:        "Asia/Urumqi",
					City:        "urumqi",
				},
			},
		},
		{
			Code: "CX",
			Name: "Christmas Island",
			Zones: []Zone{
				{
					CountryCode: "CX",
					Name:        "Indian/Christmas",
					City:        "christmas",
				},
			},
		},
		{
			Code: "CC",
			Name: "Cocos (Keeling) Islands",
			Zones: []Zone{
				{
					CountryCode: "CC",
					Name:        "Indian/Cocos",
					City:        "cocos",
				},
			},
		},
		{
			Code: "CO",
			Name: "Colombia",
			Zones: []Zone{
				{
					CountryCode: "CO",
					Name:        "America/Bogota",
					City:        "bogota",
				},
			},
		},
		{
			Code: "KM",
			Name: "Comoros",
			Zones: []Zone{
				{
					CountryCode: "KM",
					Name:        "Indian/Comoro",
					City:        "comoro",
				},
			},
		},
		{
			Code: "CG",
			Name: "Congo",
			Zones: []Zone{
				{
					CountryCode: "CG",
					Name:        "Africa/Brazzaville",
					City:        "brazzaville",
				},
			},
		},
		{
			Code: "CD",
			Name: "Congo, the Democratic Republic of the",
			Zones: []Zone{
				{
					CountryCode: "CD",
					Name:        "Africa/Kinshasa",
					City:        "kinshasa",
				},
				{
					CountryCode: "CD",
					Name:        "Africa/Lubumbashi",
					City:        "lubumbashi",
				},
			},
		},
		{
			Code: "CK",
			Name: "Cook Islands",
			Zones: []Zone{
				{
					CountryCode: "CK",
					Name:        "Pacific/Rarotonga",
					City:        "rarotonga",
				},
			},
		},
		{
			Code: "CR",
			Name: "Costa Rica",
			Zones: []Zone{
				{
					CountryCode: "CR",
					Name:        "America/Costa_Rica",
					City:        "costa rica",
				},
			},
		},
		{
			Code: "HR",
			Name: "Croatia",
			Zones: []Zone{
				{
					CountryCode: "HR",
					Name:        "Europe/Zagreb",
					City:        "zagreb",
				},
			},
		},
		{
			Code: "CU",
			Name: "Cuba",
			Zones: []Zone{
				{
					CountryCode: "CU",
					Name:        "America/Havana",
					City:        "havana",
				},
			},
		},
		{
			Code: "CW",
			Name: "Curaçao",
			Zones: []Zone{
				{
					CountryCode: "CW",
					Name:        "America/Curacao",
					City:        "curacao",
				},
			},
		},
		{
			Code: "CY",
			Name: "Cyprus",
			Zones: []Zone{
				{
					CountryCode: "CY",
					Name:        "Asia/Famagusta",
					City:        "famagusta",
				},
				{
					CountryCode: "CY",
					Name:        "Asia/Nicosia",
					City:        "nicosia",
				},
			},
		},
		{
			Code: "CZ",
			Name: "Czech Republic",
			Zones: []Zone{
				{
					CountryCode: "CZ",
					Name:        "Europe/Prague",
					City:        "prague",
				},
			},
		},
		{
			Code: "CI",
			Name: "Côte d'Ivoire",
			Zones: []Zone{
				{
					CountryCode: "CI",
					Name:        "Africa/Abidjan",
					City:        "abidjan",
				},
			},
		},
		{
			Code: "DK",
			Name: "Denmark",
			Zones: []Zone{
				{
					CountryCode: "DK",
					Name:        "Europe/Copenhagen",
					City:        "copenhagen",
				},
			},
		},
		{
			Code: "DJ",
			Name: "Djibouti",
			Zones: []Zone{
				{
					CountryCode: "DJ",
					Name:        "Africa/Djibouti",
					City:        "djibouti",
				},
			},
		},
		{
			Code: "DM",
			Name: "Dominica",
			Zones: []Zone{
				{
					CountryCode: "DM",
					Name:        "America/Dominica",
					City:        "dominica",
				},
			},
		},
		{
			Code: "DO",
			Name: "Dominican Republic",
			Zones: []Zone{
				{
					CountryCode: "DO",
					Name:        "America/Santo_Domingo",
					City:        "santo domingo",
				},
			},
		},
		{
			Code: "EC",
			Name: "Ecuador",
			Zones: []Zone{
				{
					CountryCode: "EC",
					Name:        "America/Guayaquil",
					City:        "guayaquil",
				},
				{
					CountryCode: "EC",
					Name:        "Pacific/Galapagos",
					City:        "galapagos",
				},
			},
		},
		{
			Code: "EG",
			Name: "Egypt",
			Zones: []Zone{
				{
					CountryCode: "EG",
					Name:        "Africa/Cairo",
					City:        "cairo",
				},
			},
		},
		{
			Code: "SV",
			Name: "El Salvador",
			Zones: []Zone{
				{
					CountryCode: "SV",
					Name:        "America/El_Salvador",
					City:        "el salvador",
				},
			},
		},
		{
			Code: "GQ",
			Name: "Equatorial Guinea",
			Zones: []Zone{
				{
					CountryCode: "GQ",
					Name:        "Africa/Malabo",
					City:        "malabo",
				},
			},
		},
		{
			Code: "ER",
			Name: "Eritrea",
			Zones: []Zone{
				{
					CountryCode: "ER",
					Name:        "Africa/Asmara",
					City:        "asmara",
				},
			},
		},
		{
			Code: "EE",
			Name: "Estonia",
			Zones: []Zone{
				{
					CountryCode: "EE",
					Name:        "Europe/Tallinn",
					City:        "tallinn",
				},
			},
		},
		{
			Code: "ET",
			Name: "Ethiopia",
			Zones: []Zone{
				{
					CountryCode: "ET",
					Name:        "Africa/Addis_Ababa",
					City:        "addis ababa",
				},
			},
		},
		{
			Code: "FK",
			Name: "Falkland Islands (Malvinas)",
			Zones: []Zone{
				{
					CountryCode: "FK",
					Name:        "Atlantic/Stanley",
					City:        "stanley",
				},
			},
		},
		{
			Code: "FO",
			Name: "Faroe Islands",
			Zones: []Zone{
				{
					CountryCode: "FO",
					Name:        "Atlantic/Faroe",
					City:        "faroe",
				},
			},
		},
		{
			Code: "FJ",
			Name: "Fiji",
			Zones: []Zone{
				{
					CountryCode: "FJ",
					Name:        "Pacific/Fiji",
					City:        "fiji",
				},
			},
		},
		{
			Code: "FI",
			Name: "Finland",
			Zones: []Zone{
				{
					CountryCode: "FI",
					Name:        "Europe/Helsinki",
					City:        "helsinki",
				},
			},
		},
		{
			Code: "FR",
			Name: "France",
			Zones: []Zone{
				{
					CountryCode: "FR",
					Name:        "Europe/Paris",
					City:        "paris",
				},
			},
		},
		{
			Code: "GF",
			Name: "French Guiana",
			Zones: []Zone{
				{
					CountryCode: "GF",
					Name:        "America/Cayenne",
					City:        "cayenne",
				},
			},
		},
		{
			Code: "PF",
			Name: "French Polynesia",
			Zones: []Zone{
				{
					CountryCode: "PF",
					Name:        "Pacific/Gambier",
					City:        "gambier",
				},
				{
					CountryCode: "PF",
					Name:        "Pacific/Marquesas",
					City:        "marquesas",
				},
				{
					CountryCode: "PF",
					Name:        "Pacific/Tahiti",
					City:        "tahiti",
				},
			},
		},
		{
			Code: "TF",
			Name: "French Southern Territories",
			Zones: []Zone{
				{
					CountryCode: "TF",
					Name:        "Indian/Kerguelen",
					City:        "kerguelen",
				},
			},
		},
		{
			Code: "GA",
			Name: "Gabon",
			Zones: []Zone{
				{
					CountryCode: "GA",
					Name:        "Africa/Libreville",
					City:        "libreville",
				},
			},
		},
		{
			Code: "GM",
			Name: "Gambia",
			Zones: []Zone{
				{
					CountryCode: "GM",
					Name:        "Africa/Banjul",
					City:        "banjul",
				},
			},
		},
		{
			Code: "GE",
			Name: "Georgia",
			Zones: []Zone{
				{
					CountryCode: "GE",
					Name:        "Asia/Tbilisi",
					City:        "tbilisi",
				},
			},
		},
		{
			Code: "DE",
			Name: "Germany",
			Zones: []Zone{
				{
					CountryCode: "DE",
					Name:        "Europe/Berlin",
					City:        "berlin",
				},
				{
					CountryCode: "DE",
					Name:        "Europe/Busingen",
					City:        "busingen",
				},
			},
		},
		{
			Code: "GH",
			Name: "Ghana",
			Zones: []Zone{
				{
					CountryCode: "GH",
					Name:        "Africa/Accra",
					City:        "accra",
				},
			},
		},
		{
			Code: "GI",
			Name: "Gibraltar",
			Zones: []Zone{
				{
					CountryCode: "GI",
					Name:        "Europe/Gibraltar",
					City:        "gibraltar",
				},
			},
		},
		{
			Code: "GR",
			Name: "Greece",
			Zones: []Zone{
				{
					CountryCode: "GR",
					Name:        "Europe/Athens",
					City:        "athens",
				},
			},
		},
		{
			Code: "GL",
			Name: "Greenland",
			Zones: []Zone{
				{
					CountryCode: "GL",
					Name:        "America/Danmarkshavn",
					City:        "danmarkshavn",
				},
				{
					CountryCode: "GL",
					Name:        "America/Godthab",
					City:        "godthab",
				},
				{
					CountryCode: "GL",
					Name:        "America/Scoresbysund",
					City:        "scoresbysund",
				},
				{
					CountryCode: "GL",
					Name:        "America/Thule",
					City:        "thule",
				},
			},
		},
		{
			Code: "GD",
			Name: "Grenada",
			Zones: []Zone{
				{
					CountryCode: "GD",
					Name:        "America/Grenada",
					City:        "grenada",
				},
			},
		},
		{
			Code: "GP",
			Name: "Guadeloupe",
			Zones: []Zone{
				{
					CountryCode: "GP",
					Name:        "America/Guadeloupe",
					City:        "guadeloupe",
				},
			},
		},
		{
			Code: "GU",
			Name: "Guam",
			Zones: []Zone{
				{
					CountryCode: "GU",
					Name:        "Pacific/Guam",
					City:        "guam",
				},
			},
		},
		{
			Code: "GT",
			Name: "Guatemala",
			Zones: []Zone{
				{
					CountryCode: "GT",
					Name:        "America/Guatemala",
					City:        "guatemala",
				},
			},
		},
		{
			Code: "GG",
			Name: "Guernsey",
			Zones: []Zone{
				{
					CountryCode: "GG",
					Name:        "Europe/Guernsey",
					City:        "guernsey",
				},
			},
		},
		{
			Code: "GN",
			Name: "Guinea",
			Zones: []Zone{
				{
					CountryCode: "GN",
					Name:        "Africa/Conakry",
					City:        "conakry",
				},
			},
		},
		{
			Code: "GW",
			Name: "Guinea-Bissau",
			Zones: []Zone{
				{
					CountryCode: "GW",
					Name:        "Africa/Bissau",
					City:        "bissau",
				},
			},
		},
		{
			Code: "GY",
			Name: "Guyana",
			Zones: []Zone{
				{
					CountryCode: "GY",
					Name:        "America/Guyana",
					City:        "guyana",
				},
			},
		},
		{
			Code: "HT",
			Name: "Haiti",
			Zones: []Zone{
				{
					CountryCode: "HT",
					Name:        "America/Port-au-Prince",
					City:        "port-au-prince",
				},
			},
		},
		{
			Code:  "HM",
			Name:  "Heard Island and McDonald Islands",
			Zones: []Zone{},
		},
		{
			Code: "VA",
			Name: "Holy See (Vatican City State)",
			Zones: []Zone{
				{
					CountryCode: "VA",
					Name:        "Europe/Vatican",
					City:        "vatican",
				},
			},
		},
		{
			Code: "HN",
			Name: "Honduras",
			Zones: []Zone{
				{
					CountryCode: "HN",
					Name:        "America/Tegucigalpa",
					City:        "tegucigalpa",
				},
			},
		},
		{
			Code: "HK",
			Name: "Hong Kong",
			Zones: []Zone{
				{
					CountryCode: "HK",
					Name:        "Asia/Hong_Kong",
					City:        "hong kong",
				},
			},
		},
		{
			Code: "HU",
			Name: "Hungary",
			Zones: []Zone{
				{
					CountryCode: "HU",
					Name:        "Europe/Budapest",
					City:        "budapest",
				},
			},
		},
		{
			Code: "IS",
			Name: "Iceland",
			Zones: []Zone{
				{
					CountryCode: "IS",
					Name:        "Atlantic/Reykjavik",
					City:        "reykjavik",
				},
			},
		},
		{
			Code: "IN",
			Name: "India",
			Zones: []Zone{
				{
					CountryCode: "IN",
					Name:        "Asia/Kolkata",
					City:        "kolkata",
				},
			},
		},
		{
			Code: "ID",
			Name: "Indonesia",
			Zones: []Zone{
				{
					CountryCode: "ID",
					Name:        "Asia/Jakarta",
					City:        "jakarta",
				},
				{
					CountryCode: "ID",
					Name:        "Asia/Jayapura",
					City:        "jayapura",
				},
				{
					CountryCode: "ID",
					Name:        "Asia/Makassar",
					City:        "makassar",
				},
				{
					CountryCode: "ID",
					Name:        "Asia/Pontianak",
					City:        "pontianak",
				},
			},
		},
		{
			Code: "IR",
			Name: "Iran, Islamic Republic of",
			Zones: []Zone{
				{
					CountryCode: "IR",
					Name:        "Asia/Tehran",
					City:        "tehran",
				},
			},
		},
		{
			Code: "IQ",
			Name: "Iraq",
			Zones: []Zone{
				{
					CountryCode: "IQ",
					Name:        "Asia/Baghdad",
					City:        "baghdad",
				},
			},
		},
		{
			Code: "IE",
			Name: "Ireland",
			Zones: []Zone{
				{
					CountryCode: "IE",
					Name:        "Europe/Dublin",
					City:        "dublin",
				},
			},
		},
		{
			Code: "IM",
			Name: "Isle of Man",
			Zones: []Zone{
				{
					CountryCode: "IM",
					Name:        "Europe/Isle_of_Man",
					City:        "isle of man",
				},
			},
		},
		{
			Code: "IL",
			Name: "Israel",
			Zones: []Zone{
				{
					CountryCode: "IL",
					Name:        "Asia/Jerusalem",
					City:        "jerusalem",
				},
			},
		},
		{
			Code: "IT",
			Name: "Italy",
			Zones: []Zone{
				{
					CountryCode: "IT",
					Name:        "Europe/Rome",
					City:        "rome",
				},
			},
		},
		{
			Code: "JM",
			Name: "Jamaica",
			Zones: []Zone{
				{
					CountryCode: "JM",
					Name:        "America/Jamaica",
					City:        "jamaica",
				},
			},
		},
		{
			Code: "JP",
			Name: "Japan",
			Zones: []Zone{
				{
					CountryCode: "JP",
					Name:        "Asia/Tokyo",
					City:        "tokyo",
				},
			},
		},
		{
			Code: "JE",
			Name: "Jersey",
			Zones: []Zone{
				{
					CountryCode: "JE",
					Name:        "Europe/Jersey",
					City:        "jersey",
				},
			},
		},
		{
			Code: "JO",
			Name: "Jordan",
			Zones: []Zone{
				{
					CountryCode: "JO",
					Name:        "Asia/Amman",
					City:        "amman",
				},
			},
		},
		{
			Code: "KZ",
			Name: "Kazakhstan",
			Zones: []Zone{
				{
					CountryCode: "KZ",
					Name:        "Asia/Almaty",
					City:        "almaty",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Aqtau",
					City:        "aqtau",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Aqtobe",
					City:        "aqtobe",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Atyrau",
					City:        "atyrau",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Oral",
					City:        "oral",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Qostanay",
					City:        "qostanay",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Qyzylorda",
					City:        "qyzylorda",
				},
			},
		},
		{
			Code: "KE",
			Name: "Kenya",
			Zones: []Zone{
				{
					CountryCode: "KE",
					Name:        "Africa/Nairobi",
					City:        "nairobi",
				},
			},
		},
		{
			Code: "KI",
			Name: "Kiribati",
			Zones: []Zone{
				{
					CountryCode: "KI",
					Name:        "Pacific/Enderbury",
					City:        "enderbury",
				},
				{
					CountryCode: "KI",
					Name:        "Pacific/Kiritimati",
					City:        "kiritimati",
				},
				{
					CountryCode: "KI",
					Name:        "Pacific/Tarawa",
					City:        "tarawa",
				},
			},
		},
		{
			Code: "KP",
			Name: "Korea, Democratic People's Republic of",
			Zones: []Zone{
				{
					CountryCode: "KP",
					Name:        "Asia/Pyongyang",
					City:        "pyongyang",
				},
			},
		},
		{
			Code: "KR",
			Name: "Korea, Republic of",
			Zones: []Zone{
				{
					CountryCode: "KR",
					Name:        "Asia/Seoul",
					City:        "seoul",
				},
			},
		},
		{
			Code: "KW",
			Name: "Kuwait",
			Zones: []Zone{
				{
					CountryCode: "KW",
					Name:        "Asia/Kuwait",
					City:        "kuwait",
				},
			},
		},
		{
			Code: "KG",
			Name: "Kyrgyzstan",
			Zones: []Zone{
				{
					CountryCode: "KG",
					Name:        "Asia/Bishkek",
					City:        "bishkek",
				},
			},
		},
		{
			Code: "LA",
			Name: "Lao People's Democratic Republic",
			Zones: []Zone{
				{
					CountryCode: "LA",
					Name:        "Asia/Vientiane",
					City:        "vientiane",
				},
			},
		},
		{
			Code: "LV",
			Name: "Latvia",
			Zones: []Zone{
				{
					CountryCode: "LV",
					Name:        "Europe/Riga",
					City:        "riga",
				},
			},
		},
		{
			Code: "LB",
			Name: "Lebanon",
			Zones: []Zone{
				{
					CountryCode: "LB",
					Name:        "Asia/Beirut",
					City:        "beirut",
				},
			},
		},
		{
			Code: "LS",
			Name: "Lesotho",
			Zones: []Zone{
				{
					CountryCode: "LS",
					Name:        "Africa/Maseru",
					City:        "maseru",
				},
			},
		},
		{
			Code: "LR",
			Name: "Liberia",
			Zones: []Zone{
				{
					CountryCode: "LR",
					Name:        "Africa/Monrovia",
					City:        "monrovia",
				},
			},
		},
		{
			Code: "LY",
			Name: "Libya",
			Zones: []Zone{
				{
					CountryCode: "LY",
					Name:        "Africa/Tripoli",
					City:        "tripoli",
				},
			},
		},
		{
			Code: "LI",
			Name: "Liechtenstein",
			Zones: []Zone{
				{
					CountryCode: "LI",
					Name:        "Europe/Vaduz",
					City:        "vaduz",
				},
			},
		},
		{
			Code: "LT",
			Name: "Lithuania",
			Zones: []Zone{
				{
					CountryCode: "LT",
					Name:        "Europe/Vilnius",
					City:        "vilnius",
				},
			},
		},
		{
			Code: "LU",
			Name: "Luxembourg",
			Zones: []Zone{
				{
					CountryCode: "LU",
					Name:        "Europe/Luxembourg",
					City:        "luxembourg",
				},
			},
		},
		{
			Code: "MO",
			Name: "Macao",
			Zones: []Zone{
				{
					CountryCode: "MO",
					Name:        "Asia/Macau",
					City:        "macau",
				},
			},
		},
		{
			Code: "MK",
			Name: "Macedonia, the Former Yugoslav Republic of",
			Zones: []Zone{
				{
					CountryCode: "MK",
					Name:        "Europe/Skopje",
					City:        "skopje",
				},
			},
		},
		{
			Code: "MG",
			Name: "Madagascar",
			Zones: []Zone{
				{
					CountryCode: "MG",
					Name:        "Indian/Antananarivo",
					City:        "antananarivo",
				},
			},
		},
		{
			Code: "MW",
			Name: "Malawi",
			Zones: []Zone{
				{
					CountryCode: "MW",
					Name:        "Africa/Blantyre",
					City:        "blantyre",
				},
			},
		},
		{
			Code: "MY",
			Name: "Malaysia",
			Zones: []Zone{
				{
					CountryCode: "MY",
					Name:        "Asia/Kuala_Lumpur",
					City:        "kuala lumpur",
				},
				{
					CountryCode: "MY",
					Name:        "Asia/Kuching",
					City:        "kuching",
				},
			},
		},
		{
			Code: "MV",
			Name: "Maldives",
			Zones: []Zone{
				{
					CountryCode: "MV",
					Name:        "Indian/Maldives",
					City:        "maldives",
				},
			},
		},
		{
			Code: "ML",
			Name: "Mali",
			Zones: []Zone{
				{
					CountryCode: "ML",
					Name:        "Africa/Bamako",
					City:        "bamako",
				},
			},
		},
		{
			Code: "MT",
			Name: "Malta",
			Zones: []Zone{
				{
					CountryCode: "MT",
					Name:        "Europe/Malta",
					City:        "malta",
				},
			},
		},
		{
			Code: "MH",
			Name: "Marshall Islands",
			Zones: []Zone{
				{
					CountryCode: "MH",
					Name:        "Pacific/Kwajalein",
					City:        "kwajalein",
				},
				{
					CountryCode: "MH",
					Name:        "Pacific/Majuro",
					City:        "majuro",
				},
			},
		},
		{
			Code: "MQ",
			Name: "Martinique",
			Zones: []Zone{
				{
					CountryCode: "MQ",
					Name:        "America/Martinique",
					City:        "martinique",
				},
			},
		},
		{
			Code: "MR",
			Name: "Mauritania",
			Zones: []Zone{
				{
					CountryCode: "MR",
					Name:        "Africa/Nouakchott",
					City:        "nouakchott",
				},
			},
		},
		{
			Code: "MU",
			Name: "Mauritius",
			Zones: []Zone{
				{
					CountryCode: "MU",
					Name:        "Indian/Mauritius",
					City:        "mauritius",
				},
			},
		},
		{
			Code: "YT",
			Name: "Mayotte",
			Zones: []Zone{
				{
					CountryCode: "YT",
					Name:        "Indian/Mayotte",
					City:        "mayotte",
				},
			},
		},
		{
			Code: "MX",
			Name: "Mexico",
			Zones: []Zone{
				{
					CountryCode: "MX",
					Name:        "America/Bahia_Banderas",
					City:        "bahia banderas",
				},
				{
					CountryCode: "MX",
					Name:        "America/Cancun",
					City:        "cancun",
				},
				{
					CountryCode: "MX",
					Name:        "America/Chihuahua",
					City:        "chihuahua",
				},
				{
					CountryCode: "MX",
					Name:        "America/Hermosillo",
					City:        "hermosillo",
				},
				{
					CountryCode: "MX",
					Name:        "America/Matamoros",
					City:        "matamoros",
				},
				{
					CountryCode: "MX",
					Name:        "America/Mazatlan",
					City:        "mazatlan",
				},
				{
					CountryCode: "MX",
					Name:        "America/Merida",
					City:        "merida",
				},
				{
					CountryCode: "MX",
					Name:        "America/Mexico_City",
					City:        "mexico city",
				},
				{
					CountryCode: "MX",
					Name:        "America/Monterrey",
					City:        "monterrey",
				},
				{
					CountryCode: "MX",
					Name:        "America/Ojinaga",
					City:        "ojinaga",
				},
				{
					CountryCode: "MX",
					Name:        "America/Tijuana",
					City:        "tijuana",
				},
			},
		},
		{
			Code: "FM",
			Name: "Micronesia, Federated States of",
			Zones: []Zone{
				{
					CountryCode: "FM",
					Name:        "Pacific/Chuuk",
					City:        "chuuk",
				},
				{
					CountryCode: "FM",
					Name:        "Pacific/Kosrae",
					City:        "kosrae",
				},
				{
					CountryCode: "FM",
					Name:        "Pacific/Pohnpei",
					City:        "pohnpei",
				},
			},
		},
		{
			Code: "MD",
			Name: "Moldova, Republic of",
			Zones: []Zone{
				{
					CountryCode: "MD",
					Name:        "Europe/Chisinau",
					City:        "chisinau",
				},
			},
		},
		{
			Code: "MC",
			Name: "Monaco",
			Zones: []Zone{
				{
					CountryCode: "MC",
					Name:        "Europe/Monaco",
					City:        "monaco",
				},
			},
		},
		{
			Code: "MN",
			Name: "Mongolia",
			Zones: []Zone{
				{
					CountryCode: "MN",
					Name:        "Asia/Choibalsan",
					City:        "choibalsan",
				},
				{
					CountryCode: "MN",
					Name:        "Asia/Hovd",
					City:        "hovd",
				},
				{
					CountryCode: "MN",
					Name:        "Asia/Ulaanbaatar",
					City:        "ulaanbaatar",
				},
			},
		},
		{
			Code: "ME",
			Name: "Montenegro",
			Zones: []Zone{
				{
					CountryCode: "ME",
					Name:        "Europe/Podgorica",
					City:        "podgorica",
				},
			},
		},
		{
			Code: "MS",
			Name: "Montserrat",
			Zones: []Zone{
				{
					CountryCode: "MS",
					Name:        "America/Montserrat",
					City:        "montserrat",
				},
			},
		},
		{
			Code: "MA",
			Name: "Morocco",
			Zones: []Zone{
				{
					CountryCode: "MA",
					Name:        "Africa/Casablanca",
					City:        "casablanca",
				},
			},
		},
		{
			Code: "MZ",
			Name: "Mozambique",
			Zones: []Zone{
				{
					CountryCode: "MZ",
					Name:        "Africa/Maputo",
					City:        "maputo",
				},
			},
		},
		{
			Code: "MM",
			Name: "Myanmar",
			Zones: []Zone{
				{
					CountryCode: "MM",
					Name:        "Asia/Yangon",
					City:        "yangon",
				},
			},
		},
		{
			Code: "NA",
			Name: "Namibia",
			Zones: []Zone{
				{
					CountryCode: "NA",
					Name:        "Africa/Windhoek",
					City:        "windhoek",
				},
			},
		},
		{
			Code: "NR",
			Name: "Nauru",
			Zones: []Zone{
				{
					CountryCode: "NR",
					Name:        "Pacific/Nauru",
					City:        "nauru",
				},
			},
		},
		{
			Code: "NP",
			Name: "Nepal",
			Zones: []Zone{
				{
					CountryCode: "NP",
					Name:        "Asia/Kathmandu",
					City:        "kathmandu",
				},
			},
		},
		{
			Code: "NL",
			Name: "Netherlands",
			Zones: []Zone{
				{
					CountryCode: "NL",
					Name:        "Europe/Amsterdam",
					City:        "amsterdam",
				},
			},
		},
		{
			Code: "NC",
			Name: "New Caledonia",
			Zones: []Zone{
				{
					CountryCode: "NC",
					Name:        "Pacific/Noumea",
					City:        "noumea",
				},
			},
		},
		{
			Code: "NZ",
			Name: "New Zealand",
			Zones: []Zone{
				{
					CountryCode: "NZ",
					Name:        "Pacific/Auckland",
					City:        "auckland",
				},
				{
					CountryCode: "NZ",
					Name:        "Pacific/Chatham",
					City:        "chatham",
				},
			},
		},
		{
			Code: "NI",
			Name: "Nicaragua",
			Zones: []Zone{
				{
					CountryCode: "NI",
					Name:        "America/Managua",
					City:        "managua",
				},
			},
		},
		{
			Code: "NE",
			Name: "Niger",
			Zones: []Zone{
				{
					CountryCode: "NE",
					Name:        "Africa/Niamey",
					City:        "niamey",
				},
			},
		},
		{
			Code: "NG",
			Name: "Nigeria",
			Zones: []Zone{
				{
					CountryCode: "NG",
					Name:        "Africa/Lagos",
					City:        "lagos",
				},
			},
		},
		{
			Code: "NU",
			Name: "Niue",
			Zones: []Zone{
				{
					CountryCode: "NU",
					Name:        "Pacific/Niue",
					City:        "niue",
				},
			},
		},
		{
			Code: "NF",
			Name: "Norfolk Island",
			Zones: []Zone{
				{
					CountryCode: "NF",
					Name:        "Pacific/Norfolk",
					City:        "norfolk",
				},
			},
		},
		{
			Code: "MP",
			Name: "Northern Mariana Islands",
			Zones: []Zone{
				{
					CountryCode: "MP",
					Name:        "Pacific/Saipan",
					City:        "saipan",
				},
			},
		},
		{
			Code: "NO",
			Name: "Norway",
			Zones: []Zone{
				{
					CountryCode: "NO",
					Name:        "Europe/Oslo",
					City:        "oslo",
				},
			},
		},
		{
			Code: "OM",
			Name: "Oman",
			Zones: []Zone{
				{
					CountryCode: "OM",
					Name:        "Asia/Muscat",
					City:        "muscat",
				},
			},
		},
		{
			Code: "PK",
			Name: "Pakistan",
			Zones: []Zone{
				{
					CountryCode: "PK",
					Name:        "Asia/Karachi",
					City:        "karachi",
				},
			},
		},
		{
			Code: "PW",
			Name: "Palau",
			Zones: []Zone{
				{
					CountryCode: "PW",
					Name:        "Pacific/Palau",
					City:        "palau",
				},
			},
		},
		{
			Code: "PS",
			Name: "Palestine, State of",
			Zones: []Zone{
				{
					CountryCode: "PS",
					Name:        "Asia/Gaza",
					City:        "gaza",
				},
				{
					CountryCode: "PS",
					Name:        "Asia/Hebron",
					City:        "hebron",
				},
			},
		},
		{
			Code: "PA",
			Name: "Panama",
			Zones: []Zone{
				{
					CountryCode: "PA",
					Name:        "America/Panama",
					City:        "panama",
				},
			},
		},
		{
			Code: "PG",
			Name: "Papua New Guinea",
			Zones: []Zone{
				{
					CountryCode: "PG",
					Name:        "Pacific/Bougainville",
					City:        "bougainville",
				},
				{
					CountryCode: "PG",
					Name:        "Pacific/Port_Moresby",
					City:        "port moresby",
				},
			},
		},
		{
			Code: "PY",
			Name: "Paraguay",
			Zones: []Zone{
				{
					CountryCode: "PY",
					Name:        "America/Asuncion",
					City:        "asuncion",
				},
			},
		},
		{
			Code: "PE",
			Name: "Peru",
			Zones: []Zone{
				{
					CountryCode: "PE",
					Name:        "America/Lima",
					City:        "lima",
				},
			},
		},
		{
			Code: "PH",
			Name: "Philippines",
			Zones: []Zone{
				{
					CountryCode: "PH",
					Name:        "Asia/Manila",
					City:        "manila",
				},
			},
		},
		{
			Code: "PN",
			Name: "Pitcairn",
			Zones: []Zone{
				{
					CountryCode: "PN",
					Name:        "Pacific/Pitcairn",
					City:        "pitcairn",
				},
			},
		},
		{
			Code: "PL",
			Name: "Poland",
			Zones: []Zone{
				{
					CountryCode: "PL",
					Name:        "Europe/Warsaw",
					City:        "warsaw",
				},
			},
		},
		{
			Code: "PT",
			Name: "Portugal",
			Zones: []Zone{
				{
					CountryCode: "PT",
					Name:        "Atlantic/Azores",
					City:        "azores",
				},
				{
					CountryCode: "PT",
					Name:        "Atlantic/Madeira",
					City:        "madeira",
				},
				{
					CountryCode: "PT",
					Name:        "Europe/Lisbon",
					City:        "lisbon",
				},
			},
		},
		{
			Code: "PR",
			Name: "Puerto Rico",
			Zones: []Zone{
				{
					CountryCode: "PR",
					Name:        "America/Puerto_Rico",
					City:        "puerto rico",
				},
			},
		},
		{
			Code: "QA",
			Name: "Qatar",
			Zones: []Zone{
				{
					CountryCode: "QA",
					Name:        "Asia/Qatar",
					City:        "qatar",
				},
			},
		},
		{
			Code: "RO",
			Name: "Romania",
			Zones: []Zone{
				{
					CountryCode: "RO",
					Name:        "Europe/Bucharest",
					City:        "bucharest",
				},
			},
		},
		{
			Code: "RU",
			Name: "Russian Federation",
			Zones: []Zone{
				{
					CountryCode: "RU",
					Name:        "Asia/Anadyr",
					City:        "anadyr",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Barnaul",
					City:        "barnaul",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Chita",
					City:        "chita",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Irkutsk",
					City:        "irkutsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Kamchatka",
					City:        "kamchatka",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Khandyga",
					City:        "khandyga",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Krasnoyarsk",
					City:        "krasnoyarsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Magadan",
					City:        "magadan",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Novokuznetsk",
					City:        "novokuznetsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Novosibirsk",
					City:        "novosibirsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Omsk",
					City:        "omsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Sakhalin",
					City:        "sakhalin",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Srednekolymsk",
					City:        "srednekolymsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Tomsk",
					City:        "tomsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Ust-Nera",
					City:        "ust-nera",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Vladivostok",
					City:        "vladivostok",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Yakutsk",
					City:        "yakutsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Yekaterinburg",
					City:        "yekaterinburg",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Astrakhan",
					City:        "astrakhan",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Kaliningrad",
					City:        "kaliningrad",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Kirov",
					City:        "kirov",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Moscow",
					City:        "moscow",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Samara",
					City:        "samara",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Saratov",
					City:        "saratov",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Ulyanovsk",
					City:        "ulyanovsk",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Volgograd",
					City:        "volgograd",
				},
			},
		},
		{
			Code: "RW",
			Name: "Rwanda",
			Zones: []Zone{
				{
					CountryCode: "RW",
					Name:        "Africa/Kigali",
					City:        "kigali",
				},
			},
		},
		{
			Code: "RE",
			Name: "Réunion",
			Zones: []Zone{
				{
					CountryCode: "RE",
					Name:        "Indian/Reunion",
					City:        "reunion",
				},
			},
		},
		{
			Code: "BL",
			Name: "Saint Barthélemy",
			Zones: []Zone{
				{
					CountryCode: "BL",
					Name:        "America/St_Barthelemy",
					City:        "st barthelemy",
				},
			},
		},
		{
			Code: "SH",
			Name: "Saint Helena, Ascension and Tristan da Cunha",
			Zones: []Zone{
				{
					CountryCode: "SH",
					Name:        "Atlantic/St_Helena",
					City:        "st helena",
				},
			},
		},
		{
			Code: "KN",
			Name: "Saint Kitts and Nevis",
			Zones: []Zone{
				{
					CountryCode: "KN",
					Name:        "America/St_Kitts",
					City:        "st kitts",
				},
			},
		},
		{
			Code: "LC",
			Name: "Saint Lucia",
			Zones: []Zone{
				{
					CountryCode: "LC",
					Name:        "America/St_Lucia",
					City:        "st lucia",
				},
			},
		},
		{
			Code: "MF",
			Name: "Saint Martin (French part)",
			Zones: []Zone{
				{
					CountryCode: "MF",
					Name:        "America/Marigot",
					City:        "marigot",
				},
			},
		},
		{
			Code: "PM",
			Name: "Saint Pierre and Miquelon",
			Zones: []Zone{
				{
					CountryCode: "PM",
					Name:        "America/Miquelon",
					City:        "miquelon",
				},
			},
		},
		{
			Code: "VC",
			Name: "Saint Vincent and the Grenadines",
			Zones: []Zone{
				{
					CountryCode: "VC",
					Name:        "America/St_Vincent",
					City:        "st vincent",
				},
			},
		},
		{
			Code: "WS",
			Name: "Samoa",
			Zones: []Zone{
				{
					CountryCode: "WS",
					Name:        "Pacific/Apia",
					City:        "apia",
				},
			},
		},
		{
			Code: "SM",
			Name: "San Marino",
			Zones: []Zone{
				{
					CountryCode: "SM",
					Name:        "Europe/San_Marino",
					City:        "san marino",
				},
			},
		},
		{
			Code: "ST",
			Name: "Sao Tome and Principe",
			Zones: []Zone{
				{
					CountryCode: "ST",
					Name:        "Africa/Sao_Tome",
					City:        "sao tome",
				},
			},
		},
		{
			Code: "SA",
			Name: "Saudi Arabia",
			Zones: []Zone{
				{
					CountryCode: "SA",
					Name:        "Asia/Riyadh",
					City:        "riyadh",
				},
			},
		},
		{
			Code: "SN",
			Name: "Senegal",
			Zones: []Zone{
				{
					CountryCode: "SN",
					Name:        "Africa/Dakar",
					City:        "dakar",
				},
			},
		},
		{
			Code: "RS",
			Name: "Serbia",
			Zones: []Zone{
				{
					CountryCode: "RS",
					Name:        "Europe/Belgrade",
					City:        "belgrade",
				},
			},
		},
		{
			Code: "SC",
			Name: "Seychelles",
			Zones: []Zone{
				{
					CountryCode: "SC",
					Name:        "Indian/Mahe",
					City:        "mahe",
				},
			},
		},
		{
			Code: "SL",
			Name: "Sierra Leone",
			Zones: []Zone{
				{
					CountryCode: "SL",
					Name:        "Africa/Freetown",
					City:        "freetown",
				},
			},
		},
		{
			Code: "SG",
			Name: "Singapore",
			Zones: []Zone{
				{
					CountryCode: "SG",
					Name:        "Asia/Singapore",
					City:        "singapore",
				},
			},
		},
		{
			Code: "SX",
			Name: "Sint Maarten (Dutch part)",
			Zones: []Zone{
				{
					CountryCode: "SX",
					Name:        "America/Lower_Princes",
					City:        "lower princes",
				},
			},
		},
		{
			Code: "SK",
			Name: "Slovakia",
			Zones: []Zone{
				{
					CountryCode: "SK",
					Name:        "Europe/Bratislava",
					City:        "bratislava",
				},
			},
		},
		{
			Code: "SI",
			Name: "Slovenia",
			Zones: []Zone{
				{
					CountryCode: "SI",
					Name:        "Europe/Ljubljana",
					City:        "ljubljana",
				},
			},
		},
		{
			Code: "SB",
			Name: "Solomon Islands",
			Zones: []Zone{
				{
					CountryCode: "SB",
					Name:        "Pacific/Guadalcanal",
					City:        "guadalcanal",
				},
			},
		},
		{
			Code: "SO",
			Name: "Somalia",
			Zones: []Zone{
				{
					CountryCode: "SO",
					Name:        "Africa/Mogadishu",
					City:        "mogadishu",
				},
			},
		},
		{
			Code: "ZA",
			Name: "South Africa",
			Zones: []Zone{
				{
					CountryCode: "ZA",
					Name:        "Africa/Johannesburg",
					City:        "johannesburg",
				},
			},
		},
		{
			Code: "GS",
			Name: "South Georgia and the South Sandwich Islands",
			Zones: []Zone{
				{
					CountryCode: "GS",
					Name:        "Atlantic/South_Georgia",
					City:        "south georgia",
				},
			},
		},
		{
			Code: "SS",
			Name: "South Sudan",
			Zones: []Zone{
				{
					CountryCode: "SS",
					Name:        "Africa/Juba",
					City:        "juba",
				},
			},
		},
		{
			Code: "ES",
			Name: "Spain",
			Zones: []Zone{
				{
					CountryCode: "ES",
					Name:        "Africa/Ceuta",
					City:        "ceuta",
				},
				{
					CountryCode: "ES",
					Name:        "Atlantic/Canary",
					City:        "canary",
				},
				{
					CountryCode: "ES",
					Name:        "Europe/Madrid",
					City:        "madrid",
				},
			},
		},
		{
			Code: "LK",
			Name: "Sri Lanka",
			Zones: []Zone{
				{
					CountryCode: "LK",
					Name:        "Asia/Colombo",
					City:        "colombo",
				},
			},
		},
		{
			Code: "SD",
			Name: "Sudan",
			Zones: []Zone{
				{
					CountryCode: "SD",
					Name:        "Africa/Khartoum",
					City:        "khartoum",
				},
			},
		},
		{
			Code: "SR",
			Name: "Suriname",
			Zones: []Zone{
				{
					CountryCode: "SR",
					Name:        "America/Paramaribo",
					City:        "paramaribo",
				},
			},
		},
		{
			Code: "SJ",
			Name: "Svalbard and Jan Mayen",
			Zones: []Zone{
				{
					CountryCode: "SJ",
					Name:        "Arctic/Longyearbyen",
					City:        "longyearbyen",
				},
			},
		},
		{
			Code: "SZ",
			Name: "Swaziland",
			Zones: []Zone{
				{
					CountryCode: "SZ",
					Name:        "Africa/Mbabane",
					City:        "mbabane",
				},
			},
		},
		{
			Code: "SE",
			Name: "Sweden",
			Zones: []Zone{
				{
					CountryCode: "SE",
					Name:        "Europe/Stockholm",
					City:        "stockholm",
				},
			},
		},
		{
			Code: "CH",
			Name: "Switzerland",
			Zones: []Zone{
				{
					CountryCode: "CH",
					Name:        "Europe/Zurich",
					City:        "zurich",
				},
			},
		},
		{
			Code: "SY",
			Name: "Syrian Arab Republic",
			Zones: []Zone{
				{
					CountryCode: "SY",
					Name:        "Asia/Damascus",
					City:        "damascus",
				},
			},
		},
		{
			Code: "TW",
			Name: "Taiwan, Province of China",
			Zones: []Zone{
				{
					CountryCode: "TW",
					Name:        "Asia/Taipei",
					City:        "taipei",
				},
			},
		},
		{
			Code: "TJ",
			Name: "Tajikistan",
			Zones: []Zone{
				{
					CountryCode: "TJ",
					Name:        "Asia/Dushanbe",
					City:        "dushanbe",
				},
			},
		},
		{
			Code: "TZ",
			Name: "Tanzania, United Republic of",
			Zones: []Zone{
				{
					CountryCode: "TZ",
					Name:        "Africa/Dar_es_Salaam",
					City:        "dar es salaam",
				},
			},
		},
		{
			Code: "TH",
			Name: "Thailand",
			Zones: []Zone{
				{
					CountryCode: "TH",
					Name:        "Asia/Bangkok",
					City:        "bangkok",
				},
			},
		},
		{
			Code: "TL",
			Name: "Timor-Leste",
			Zones: []Zone{
				{
					CountryCode: "TL",
					Name:        "Asia/Dili",
					City:        "dili",
				},
			},
		},
		{
			Code: "TG",
			Name: "Togo",
			Zones: []Zone{
				{
					CountryCode: "TG",
					Name:        "Africa/Lome",
					City:        "lome",
				},
			},
		},
		{
			Code: "TK",
			Name: "Tokelau",
			Zones: []Zone{
				{
					CountryCode: "TK",
					Name:        "Pacific/Fakaofo",
					City:        "fakaofo",
				},
			},
		},
		{
			Code: "TO",
			Name: "Tonga",
			Zones: []Zone{
				{
					CountryCode: "TO",
					Name:        "Pacific/Tongatapu",
					City:        "tongatapu",
				},
			},
		},
		{
			Code: "TT",
			Name: "Trinidad and Tobago",
			Zones: []Zone{
				{
					CountryCode: "TT",
					Name:        "America/Port_of_Spain",
					City:        "port of spain",
				},
			},
		},
		{
			Code: "TN",
			Name: "Tunisia",
			Zones: []Zone{
				{
					CountryCode: "TN",
					Name:        "Africa/Tunis",
					City:        "tunis",
				},
			},
		},
		{
			Code: "TR",
			Name: "Turkey",
			Zones: []Zone{
				{
					CountryCode: "TR",
					Name:        "Europe/Istanbul",
					City:        "istanbul",
				},
			},
		},
		{
			Code: "TM",
			Name: "Turkmenistan",
			Zones: []Zone{
				{
					CountryCode: "TM",
					Name:        "Asia/Ashgabat",
					City:        "ashgabat",
				},
			},
		},
		{
			Code: "TC",
			Name: "Turks and Caicos Islands",
			Zones: []Zone{
				{
					CountryCode: "TC",
					Name:        "America/Grand_Turk",
					City:        "grand turk",
				},
			},
		},
		{
			Code: "TV",
			Name: "Tuvalu",
			Zones: []Zone{
				{
					CountryCode: "TV",
					Name:        "Pacific/Funafuti",
					City:        "funafuti",
				},
			},
		},
		{
			Code: "UG",
			Name: "Uganda",
			Zones: []Zone{
				{
					CountryCode: "UG",
					Name:        "Africa/Kampala",
					City:        "kampala",
				},
			},
		},
		{
			Code: "UA",
			Name: "Ukraine",
			Zones: []Zone{
				{
					CountryCode: "UA",
					Name:        "Europe/Kiev",
					City:        "kiev",
				},
				{
					CountryCode: "UA",
					Name:        "Europe/Simferopol",
					City:        "simferopol",
				},
				{
					CountryCode: "UA",
					Name:        "Europe/Uzhgorod",
					City:        "uzhgorod",
				},
				{
					CountryCode: "UA",
					Name:        "Europe/Zaporozhye",
					City:        "zaporozhye",
				},
			},
		},
		{
			Code: "AE",
			Name: "United Arab Emirates",
			Zones: []Zone{
				{
					CountryCode: "AE",
					Name:        "Asia/Dubai",
					City:        "dubai",
				},
			},
		},
		{
			Code: "GB",
			Name: "United Kingdom",
			Zones: []Zone{
				{
					CountryCode: "GB",
					Name:        "Europe/London",
					City:        "london",
				},
			},
		},
		{
			Code: "US",
			Name: "United States",
			Zones: []Zone{
				{
					CountryCode: "US",
					Name:        "America/Adak",
					City:        "adak",
				},
				{
					CountryCode: "US",
					Name:        "America/Anchorage",
					City:        "anchorage",
				},
				{
					CountryCode: "US",
					Name:        "America/Boise",
					City:        "boise",
				},
				{
					CountryCode: "US",
					Name:        "America/Chicago",
					City:        "chicago",
				},
				{
					CountryCode: "US",
					Name:        "America/Denver",
					City:        "denver",
				},
				{
					CountryCode: "US",
					Name:        "America/Detroit",
					City:        "detroit",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Indianapolis",
					City:        "indianapolis",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Knox",
					City:        "knox",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Marengo",
					City:        "marengo",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Petersburg",
					City:        "petersburg",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Tell_City",
					City:        "tell city",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Vevay",
					City:        "vevay",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Vincennes",
					City:        "vincennes",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Winamac",
					City:        "winamac",
				},
				{
					CountryCode: "US",
					Name:        "America/Juneau",
					City:        "juneau",
				},
				{
					CountryCode: "US",
					Name:        "America/Kentucky/Louisville",
					City:        "louisville",
				},
				{
					CountryCode: "US",
					Name:        "America/Kentucky/Monticello",
					City:        "monticello",
				},
				{
					CountryCode: "US",
					Name:        "America/Los_Angeles",
					City:        "los angeles",
				},
				{
					CountryCode: "US",
					Name:        "America/Menominee",
					City:        "menominee",
				},
				{
					CountryCode: "US",
					Name:        "America/Metlakatla",
					City:        "metlakatla",
				},
				{
					CountryCode: "US",
					Name:        "America/New_York",
					City:        "new york",
				},
				{
					CountryCode: "US",
					Name:        "America/Nome",
					City:        "nome",
				},
				{
					CountryCode: "US",
					Name:        "America/North_Dakota/Beulah",
					City:        "beulah",
				},
				{
					CountryCode: "US",
					Name:        "America/North_Dakota/Center",
					City:        "center",
				},
				{
					CountryCode: "US",
					Name:        "America/North_Dakota/New_Salem",
					City:        "new salem",
				},
				{
					CountryCode: "US",
					Name:        "America/Phoenix",
					City:        "phoenix",
				},
				{
					CountryCode: "US",
					Name:        "America/Sitka",
					City:        "sitka",
				},
				{
					CountryCode: "US",
					Name:        "America/Yakutat",
					City:        "yakutat",
				},
				{
					CountryCode: "US",
					Name:        "Pacific/Honolulu",
					City:        "honolulu",
				},
			},
		},
		{
			Code: "UM",
			Name: "United States Minor Outlying Islands",
			Zones: []Zone{
				{
					CountryCode: "UM",
					Name:        "Pacific/Midway",
					City:        "midway",
				},
				{
					CountryCode: "UM",
					Name:        "Pacific/Wake",
					City:        "wake",
				},
			},
		},
		{
			Code: "UY",
			Name: "Uruguay",
			Zones: []Zone{
				{
					CountryCode: "UY",
					Name:        "America/Montevideo",
					City:        "montevideo",
				},
			},
		},
		{
			Code: "UZ",
			Name: "Uzbekistan",
			Zones: []Zone{
				{
					CountryCode: "UZ",
					Name:        "Asia/Samarkand",
					City:        "samarkand",
				},
				{
					CountryCode: "UZ",
					Name:        "Asia/Tashkent",
					City:        "tashkent",
				},
			},
		},
		{
			Code: "VU",
			Name: "Vanuatu",
			Zones: []Zone{
				{
					CountryCode: "VU",
					Name:        "Pacific/Efate",
					City:        "efate",
				},
			},
		},
		{
			Code: "VE",
			Name: "Venezuela, Bolivarian Republic of",
			Zones: []Zone{
				{
					CountryCode: "VE",
					Name:        "America/Caracas",
					City:        "caracas",
				},
			},
		},
		{
			Code: "VN",
			Name: "Viet Nam",
			Zones: []Zone{
				{
					CountryCode: "VN",
					Name:        "Asia/Ho_Chi_Minh",
					City:        "ho chi minh",
				},
			},
		},
		{
			Code: "VG",
			Name: "Virgin Islands, British",
			Zones: []Zone{
				{
					CountryCode: "VG",
					Name:        "America/Tortola",
					City:        "tortola",
				},
			},
		},
		{
			Code: "VI",
			Name: "Virgin Islands, U.S.",
			Zones: []Zone{
				{
					CountryCode: "VI",
					Name:        "America/St_Thomas",
					City:        "st thomas",
				},
			},
		},
		{
			Code: "WF",
			Name: "Wallis and Futuna",
			Zones: []Zone{
				{
					CountryCode: "WF",
					Name:        "Pacific/Wallis",
					City:        "wallis",
				},
			},
		},
		{
			Code: "EH",
			Name: "Western Sahara",
			Zones: []Zone{
				{
					CountryCode: "EH",
					Name:        "Africa/El_Aaiun",
					City:        "el aaiun",
				},
			},
		},
		{
			Code: "YE",
			Name: "Yemen",
			Zones: []Zone{
				{
					CountryCode: "YE",
					Name:        "Asia/Aden",
					City:        "aden",
				},
			},
		},
		{
			Code: "ZM",
			Name: "Zambia",
			Zones: []Zone{
				{
					CountryCode: "ZM",
					Name:        "Africa/Lusaka",
					City:        "lusaka",
				},
			},
		},
		{
			Code: "ZW",
			Name: "Zimbabwe",
			Zones: []Zone{
				{
					CountryCode: "ZW",
					Name:        "Africa/Harare",
					City:        "harare",
				},
			},
		},
		{
			Code: "AX",
			Name: "Åland Islands",
			Zones: []Zone{
				{
					CountryCode: "AX",
					Name:        "Europe/Mariehamn",
					City:        "mariehamn",
				},
			},
		},
	}
)

func init() {
	// load + index countries into map
	// for below functions.

	once.Do(func() {
		mapped = make(map[string]Country)

		for i := 0; i < len(countries); i++ {
			mapped[countries[i].Code] = countries[i]
		}
	})
}

// GetCountries returns an array of all countries.
// Most common use: for loading into a country dropdown
// in HTML.
func GetCountries() []Country {
	return countries
}

// GetCountry returns a single Country that matches the country
// code passed and whether it was found
func GetCountry(code string) (c Country, found bool) {
	c, found = mapped[code]
	return
}
