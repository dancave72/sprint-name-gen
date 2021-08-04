package got

import "github.com/flostadler/name-generator/pkg/randomizer"

var(
	gotNames = [...] string {
		"Jon",
		"Snow",
		"Tyrion",
		"Lannister",
		"Arya",
		"Stark",
		"Jaime",
		"Daenerys",
		"Targaryen",
		"Sansa",
		"Bran",
		"Eddard",
		"Catelyn",
		"Cersei",
		"Robb",
		"Samwell",
		"Tarly",
		"Stannis",
		"Baratheon",
		"Joffrey",
		"Theon",
		"Greyjoy",
		"Robert",
		"Brienne",
		"Tarth",
		"Tywin",
		"Davos",
		"Seaworth",
		"Petyr",
		"Baelish",
		"Sandor",
		"Clegane",
		"Barristan",
		"Selmy",
		"Renly",
		"Jorah",
		"Mormont",
		"Jeor",
		"Varys",
		"Tommen",
		"Mance",
		"Rayder",
		"Gregor",
		"Asha",
		"Roose",
		"Bolton",
		"Lysa",
		"Arryn",
		"Bronn",
		"Hodor",
		"Ramsay",
		"Loras",
		"Tyrell",
		"Maester",
		"Aemon",
		"Margaery",
		"Edmure",
		"Tully",
		"Khal",
		"Drogo",
		"Pycelle",
		"Melisandre",
		"Quentyn",
		"Martell",
		"Gendry",
		"Victarion",
		"Hizdahr",
		"Loraq",
		"Luwin",
		"Euron",
		"Meera",
		"Reed",
		"Kevan",
		"Rodrik",
		"Cassel",
		"Gilly",
		"Walder",
		"Frey",
		"Aerys",
		"Connington",
		"Janos",
		"Slynt",
		"Beric",
		"Dondarrion",
		"Rhaegar",
		"Brynden",
		"Podrick",
		"Payne",
		"Jojen",
		"Mace",
		"Grenn",
		"Illyrio",
		"Mopatis",
		"Tormund",
		"Giantsbane",
		"Balon",
		"Hot Pie",
		"Myrcella",
		"Viserys",
		"Aeron",
		"Qhorin",
		"Halfhand",
		"Oberyn",
		"Rickon",
		"Ygritte",
		"Daario",
		"Naharis",
		"Wyman",
		"Manderly",
		"Yoren",
		"Craster",
		"Alliser",
		"Thorne",
		"Shae",
		"Qyburn",
		"Bowen",
		"Marsh",
		"Ilyn",
		"Lancel",
		"Eddison",
		"Tollett",
		"Arianne",
		"Pyp",
		"Osmund",
		"Kettleblack",
		"Skahaz",
		"Kandaq",
		"Meryn",
		"Trant",
		"Belwas",
		"Boros",
		"Blount",
		"Randyll",
		"Haldon",
		"Halfmaester",
		"Old Nan",
		"Penny",
		"Irri",
		"Doran",
		"Lem",
		"Lemoncloak",
		"Jeyne",
		"Poole",
		"Dontos",
		"Hollard",
		"Cleos",
		"Benjen",
		"Jhiqui",
		"Gerris",
		"Drinkwater",
		"Ben",
		"Plumm",
		"Salladhor",
		"Saan",
		"Val",
		"Taena",
		"Merryweather",
		"Thoros",
		"Myr",
		"Vargo",
		"Hoat",
		"Swann",
		"Hoster",
		"Reznak",
		"Xaro",
		"Xhoan",
		"Daxos",
		"Greatjon",
		"Umber",
		"Jory",
		"Donal",
		"Noye",
		"Tom",
		"Sevenstreams",
		"Aegon",
		"Arys",
		"Oakheart",
		"Osha",
		"Chett",
		"Hyle",
		"Hunt",
		"Missandei",
		"Rattleshirt",
		"Axell",
		"Florent",
		"Gyles",
		"Rosby",
		"Dick",
		"Crabb",
		"Cressen",
		"Areo",
		"Hotah",
		"Addam",
		"Marbrand",
		"Septa",
		"Mordane",
		"Osney",
		"Syrio",
		"Forel",
		"Selyse",
		"Rolly",
		"Duckfield",
		"Jhogo",
		"Shagga",
		"Dareon",
		"Varamyr",
		"Sixskins",
		"Lommy",
		"Greenhands",
		"Nestor",
		"Royce",
		"Rickard",
		"Karstark",
		"Amory",
		"Lorch",
		"Satin",
		"Cotter",
		"Pyke",
		"Mya",
		"Stone",
		"Paxter",
		"Redwyne",
		"Aggo",
		"Harwin",
		"High",
		"Septon",
		"Jaqen",
		"H’ghar",
		"Merrett",
		"Mirri",
		"Maz",
		"Duur",
		"Ryman",
		"Anguy",
		"Clydas",
		"Marillion",
		"Pate",
		"Rorge",
		"Elia",
		"Jacelyn",
		"Bywater",
		"Brandon",
		"Yezzan",
		"Qaggaz",
		"Pylos",
		"Meribald",
		"Weese",
		"Justin",
		"Massey",
		"Big",
		"Baelor",
		"Devan",
		"Harys",
		"Swyft",
		"Mandon",
		"Moore",
		"Westerling",
		"Othell",
		"Yarwyck",
		"Galazza",
		"Galare",
		"Grey",
		"Worm",
		"Will",
		"Lothor",
		"Brune",
		"Tytos",
		"Blackwood",
		"Jonos",
		"Bracken",
		"Orton",
		"Small",
		"Paul",
		"Robett",
		"Glover",
		"Tanda",
		"Stokeworth",
		"Sarella",
		"Sand",
		"Biter",
		"Obara",
		"Steelshanks",
		"Walton",
		"Denys",
		"Mallister",
		"Dywen",
		"Galbart",
		"Lollys",
		"Thoren",
		"Smallwood",
		"Kraznys",
		"Nakloz",
		"Polliver",
		"Rakharo",
		"Tristifer",
		"Botley",
		"Aurane",
		"Waters",
		"Emmon",
		"Jarl",
		"Olenna",
		"Harlaw",
		"Shireen",
		"Cortnay",
		"Penrose",
		"Edwyn",
		"Archibald",
		"Yronwood",
		"Mathis",
		"Rowan",
		"Vardis",
		"Egen",
		"Yohn",
		"Barbrey",
		"Ryswell",
		"Tattered",
		"Prince",
		"Tickler",
		"Gerold",
		"Dayne",
		"Harry",
		"Strickland",
		"Three-Finger",
		"Hobb",
		"Tyene",
		"Lemore",
		"Lyn",
		"Corbray",
		"Roslin",
		"Fat",
		"Waymar",
		"Arthur",
		"Doreah",
		"Iron",
		"Emmett",
		"Wendel",
		"Dalla",
		"Godry",
		"Farring",
		"Lothar",
		"Mord",
		"Nymeria",
		"Toad",
		"Cleon",
		"the Great",
		"Dagmer",
		"Cleftjaw",
		"Garlan",
		"Shagwell",
		"Willas",
		"Halder",
		"Jason",
		"Lyanna",
		"Moqorro",
		"Osfryd",
		"Patchface",
		"Rafford",
	}
)

func Generate() string {
	return randomizer.GetRandom(gotNames[:]) + " " + randomizer.GetRandom(gotNames[:])
}
