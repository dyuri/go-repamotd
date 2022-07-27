package widgets

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var napok = []string{"vasárnap", "hétfő", "kedd", "szerda", "csütörtök", "péntek", "szombat", "vasárnap"}
var honapok = []string{"január", "február", "március", "április", "május", "június", "július", "augusztus", "szeptember", "október", "november", "december"}

var nevnapok = map[string]string{
	"0101": "Fruzsina",
	"0102": "Ábel",
	"0103": "Genovéva, Benjámin",
	"0104": "Titusz, Leona",
	"0105": "Simon",
	"0106": "Boldizsár",
	"0107": "Attila, Ramóna",
	"0108": "Gyöngyvér",
	"0109": "Marcell",
	"0110": "Melánia",
	"0111": "Ágota",
	"0112": "Ernõ",
	"0113": "Veronika",
	"0114": "Bódog",
	"0115": "Lóránt, Loránd",
	"0116": "Gusztáv",
	"0117": "Antal, Antónia",
	"0118": "Piroska",
	"0119": "Sára, Márió",
	"0120": "Fábián, Sebestyén",
	"0121": "Ágnes",
	"0122": "Vince, Artúr",
	"0123": "Zelma, Rajmund",
	"0124": "Timót",
	"0125": "Pál",
	"0126": "Vanda, Paula",
	"0127": "Angelika",
	"0128": "Károly, Karola",
	"0129": "Adél",
	"0130": "Martina, Gerda",
	"0131": "Marcella",
	"0201": "Ignác",
	"0202": "Karolina, Aida",
	"0203": "Balázs",
	"0204": "Ráhel, Csenge",
	"0205": "Ágota, Ingrid",
	"0206": "Dorottya, Dóra",
	"0207": "Tódor, Rómeó",
	"0208": "Aranka",
	"0209": "Abigél, Alex",
	"0210": "Elvira",
	"0211": "Bertold, Marietta",
	"0212": "Lívia, Lídia",
	"0213": "Ella, Linda",
	"0214": "Bálint, Valentin",
	"0215": "Kolos, Georgina",
	"0216": "Julianna, Lilla",
	"0217": "Donát",
	"0218": "Bernadett",
	"0219": "Zsuzsanna",
	"0220": "Aladár, Álmos",
	"0221": "Eleonóra",
	"0222": "Gerzson",
	"0223": "Alfréd",
	"0224": "Mátyás",
	"0225": "Géza",
	"0226": "Edina",
	"0227": "Ákos, Bátor",
	"0228": "Elemér",
	"0229": "Anonymus",
	"0301": "Albin",
	"0302": "Lujza",
	"0303": "Kornélia",
	"0304": "Kázmér",
	"0305": "Adorján, Adrián",
	"0306": "Leonóra, Inez",
	"0307": "Tamás",
	"0308": "Zoltán",
	"0309": "Franciska, Fanni",
	"0310": "Ildikó",
	"0311": "Szilárd",
	"0312": "Gergely",
	"0313": "Krisztián, Ajtony",
	"0314": "Matild",
	"0315": "Kristóf",
	"0316": "Henrietta",
	"0317": "Gertrúd, Patrik",
	"0318": "Sándor, Ede",
	"0319": "József, Bánk",
	"0320": "Klaudia",
	"0321": "Benedek",
	"0322": "Beáta, Izolda",
	"0323": "Emõke",
	"0324": "Gábor, Karina",
	"0325": "Irén, Irisz",
	"0326": "Emánuel",
	"0327": "Hajnalka",
	"0328": "Gedeon, Johanna",
	"0329": "Aguszta",
	"0330": "Zalán",
	"0331": "Árpád",
	"0401": "Hugó",
	"0402": "Áron",
	"0403": "Buda, Richárd",
	"0404": "Izidor",
	"0405": "Vince",
	"0406": "Vilmos, Bíborka",
	"0407": "Herman",
	"0408": "Dénes",
	"0409": "Erhard",
	"0410": "Zsolt",
	"0411": "Leó, Szaniszló",
	"0412": "Gyula",
	"0413": "Ida",
	"0414": "Tibor",
	"0415": "Anasztázia, Tas",
	"0416": "Csongor",
	"0417": "Rudolf",
	"0418": "Andrea, Ilma",
	"0419": "Emma",
	"0420": "Tivadar",
	"0421": "Konrád",
	"0422": "Csilla, Noémi",
	"0423": "Béla",
	"0424": "György",
	"0425": "Márk",
	"0426": "Ervin",
	"0427": "Zita",
	"0428": "Valéria",
	"0429": "Péter",
	"0430": "Katalin, Kitti",
	"0501": "Fülöp, Jakab",
	"0502": "Zsigmond",
	"0503": "Tímea, Irma",
	"0504": "Mónika, Flórián",
	"0505": "Györgyi",
	"0506": "Ivett, Frida",
	"0507": "Gizella",
	"0508": "Mihály",
	"0509": "Gergely",
	"0510": "Ármin, Pálma",
	"0511": "Ferenc",
	"0512": "Pongrác",
	"0513": "Szervác, Imola",
	"0514": "Bonifác",
	"0515": "Zsófia, Szonja",
	"0516": "Mózes, Botond",
	"0517": "Paszkál",
	"0518": "Erik, Alexandra",
	"0519": "Ivó, Milán",
	"0520": "Bernát, Felícia",
	"0521": "Konstantin",
	"0522": "Júlia, Rita",
	"0523": "Dezsõ",
	"0524": "Eszter, Eliza",
	"0525": "Orbán",
	"0526": "Fülöp, Evelin",
	"0527": "Hella",
	"0528": "Emil, Csanád",
	"0529": "Magdolna",
	"0530": "Janka, Zsanett",
	"0531": "Angéla, Petronella",
	"0601": "Tünde",
	"0602": "Kármen, Anita",
	"0603": "Klotild",
	"0604": "Bulcsú",
	"0605": "Fatime",
	"0606": "Norbert, Cintia",
	"0607": "Róbert",
	"0608": "Medárd",
	"0609": "Félix",
	"0610": "Margit, Gréta",
	"0611": "Barnabás",
	"0612": "Villõ",
	"0613": "Antal, Anett",
	"0614": "Vazul",
	"0615": "Jolán, Vid",
	"0616": "Jusztin",
	"0617": "Laura, Alida",
	"0618": "Arnold, Levente",
	"0619": "Gyárfás",
	"0620": "Rafael",
	"0621": "Alajos, Leila",
	"0622": "Paulina",
	"0623": "Zoltán",
	"0624": "Iván",
	"0625": "Vilmos",
	"0626": "János, Pál",
	"0627": "László",
	"0628": "Levente, Irén",
	"0629": "Péter, Pál",
	"0630": "Pál",
	"0701": "Tihamér, Annamária",
	"0702": "Ottó",
	"0703": "Kornél, Soma",
	"0704": "Ulrik",
	"0705": "Emese, Sarolta",
	"0706": "Csaba",
	"0707": "Apollónia",
	"0708": "Ellák",
	"0709": "Lukrécia",
	"0710": "Amália",
	"0711": "Nóra, Lili",
	"0712": "Izabella, Dalma",
	"0713": "Jenõ",
	"0714": "Örs, Stella",
	"0715": "Henrik, Roland",
	"0716": "Valter",
	"0717": "Endre, Elek",
	"0718": "Frigyes",
	"0719": "Emília",
	"0720": "Illés",
	"0721": "Dániel, Daniella",
	"0722": "Magdolna",
	"0723": "Lenke",
	"0724": "Kinga, Kincsõ",
	"0725": "Kristóf, Jakab",
	"0726": "Anna, Anikó",
	"0727": "Olga, Liliána",
	"0728": "Szabolcs",
	"0729": "Márta, Flóra",
	"0730": "Judit, Xénia",
	"0731": "Oszkár",
	"0801": "Boglárka",
	"0802": "Lehel",
	"0803": "Hermina",
	"0804": "Domonkos, Dominika",
	"0805": "Krisztina",
	"0806": "Berta, Bettina",
	"0807": "Ibolya",
	"0808": "László",
	"0809": "Emõd",
	"0810": "Lõrinc",
	"0811": "Zsuzsanna, Tiborc",
	"0812": "Klára",
	"0813": "Ipoly",
	"0814": "Marcell",
	"0815": "Mária",
	"0816": "Ábrahám",
	"0817": "Jácint",
	"0818": "Ilona",
	"0819": "Huba",
	"0820": "István",
	"0821": "Sámuel, Hajna",
	"0822": "Menyhért, Mirjam",
	"0823": "Bence",
	"0824": "Bertalan",
	"0825": "Lajos, Patrícia",
	"0826": "Izsó",
	"0827": "Gáspár",
	"0828": "Ágoston",
	"0829": "Beatrix, Erna",
	"0830": "Rózsa",
	"0831": "Erika, Bella",
	"0901": "Egyed, Egon",
	"0902": "Rebeka, Dorina",
	"0903": "Hilda",
	"0904": "Rozália",
	"0905": "Viktor, Lõrinc",
	"0906": "Zakariás",
	"0907": "Regina",
	"0908": "Mária, Adrienn",
	"0909": "Ádám",
	"0910": "Nikolett, Hunor",
	"0911": "Teodóra",
	"0912": "Mária",
	"0913": "Kornél",
	"0914": "Szeréna, Roxána",
	"0915": "Enikõ, Melitta",
	"0916": "Edit",
	"0917": "Zsófia",
	"0918": "Diána",
	"0919": "Vilhelmina",
	"0920": "Friderika",
	"0921": "Máté, Mirella",
	"0922": "Móric",
	"0923": "Tekla",
	"0924": "Gellért, Mercédesz",
	"0925": "Eufrozina, Kende",
	"0926": "Jusztina",
	"0927": "Adalbert",
	"0928": "Vencel",
	"0929": "Mihály",
	"0930": "Jeromos",
	"1001": "Malvin",
	"1002": "Petra",
	"1003": "Helga",
	"1004": "Ferenc",
	"1005": "Aurél",
	"1006": "Brúnó, Renáta",
	"1007": "Amália",
	"1008": "Koppány",
	"1009": "Dénes",
	"1010": "Gedeon",
	"1011": "Brigitta",
	"1012": "Miksa",
	"1013": "Kálmán, Ede",
	"1014": "Helén",
	"1015": "Teréz",
	"1016": "Gál",
	"1017": "Hedvig",
	"1018": "Lukács",
	"1019": "Nándor",
	"1020": "Vendel",
	"1021": "Orsolya",
	"1022": "Elõd",
	"1023": "Gyöngyi",
	"1024": "Salamon",
	"1025": "Blanka, Bianka",
	"1026": "Dömötör",
	"1027": "Szabina",
	"1028": "Simon, Szimonetta",
	"1029": "Nárcisz",
	"1030": "Alfonz",
	"1031": "Farkas",
	"1101": "Marianna",
	"1102": "Achilles",
	"1103": "Gyõzõ",
	"1104": "Károly",
	"1105": "Imre",
	"1106": "Lénárd",
	"1107": "Rezsõ",
	"1108": "Zsombor",
	"1109": "Tivadar",
	"1110": "Réka",
	"1111": "Márton",
	"1112": "Jónás, Renátó",
	"1113": "Szilvia",
	"1114": "Aliz",
	"1115": "Albert, Lipót",
	"1116": "Ödön",
	"1117": "Hortenzia, Gergõ",
	"1118": "Jenõ",
	"1119": "Erzsébet",
	"1120": "Jolán",
	"1121": "Olivér",
	"1122": "Cecília",
	"1123": "Kelemen, Klementina",
	"1124": "Emma",
	"1125": "Katalin",
	"1126": "Virág",
	"1127": "Virgil",
	"1128": "Stefánia",
	"1129": "Taksony",
	"1130": "András, Andor",
	"1201": "Elza",
	"1202": "Melinda, Vivien",
	"1203": "Ferenc",
	"1204": "Borbála, Barbara",
	"1205": "Vilma",
	"1206": "Miklós",
	"1207": "Ambrus",
	"1208": "Mária",
	"1209": "Natália",
	"1210": "Judit",
	"1211": "Árpád",
	"1212": "Gabriella",
	"1213": "Luca, Otília",
	"1214": "Szilárda",
	"1215": "Valér",
	"1216": "Etelka, Aletta",
	"1217": "Lázár, Olimpia",
	"1218": "Auguszta",
	"1219": "Viola",
	"1220": "Teofil",
	"1221": "Tamás",
	"1222": "Zénó",
	"1223": "Viktória",
	"1224": "Ádám, Éva",
	"1225": "Eugénia",
	"1226": "István",
	"1227": "János",
	"1228": "Kamilla",
	"1229": "Tamás, Tamara",
	"1230": "Dávid",
	"1231": "Szilveszter",
}

// NaptarWidget is a widget that displays calendar information
func NaptarWidget(v *viper.Viper, f formatFn) (string, error) {
	content := strings.Builder{}

	currentTime := time.Now()
	nap := napok[currentTime.Weekday()]
	honap := honapok[currentTime.Month()-1]
	_, iweek := currentTime.ISOWeek()
	lastDayOfYear := time.Date(currentTime.Year(), 12, 31, 0, 0, 0, 0, time.UTC).YearDay()

	holnap := currentTime.AddDate(0, 0, 1)
	holnaputan := currentTime.AddDate(0, 0, 2)
	holnaputan2 := currentTime.AddDate(0, 0, 3)
	nevnapKey := fmt.Sprintf("%02d%02d", currentTime.Month(), currentTime.Day())
	nevnapKey2 := fmt.Sprintf("%02d%02d", holnap.Month(), holnap.Day())
	nevnapKey3 := fmt.Sprintf("%02d%02d", holnaputan.Month(), holnaputan.Day())
	nevnapKey4 := fmt.Sprintf("%02d%02d", holnaputan2.Month(), holnaputan2.Day())

	f1 := f("10", "", true)
	f2 := f("10", "", false)
	f3 := f("11", "", true)
	f4 := f("3", "", false)

	fmt.Fprintf(&content, "Ma %s van, %s\n",
		f1(nap),
		f2(fmt.Sprintf("%d. %s %d.", currentTime.Year(), honap, currentTime.Day())),
	)
	fmt.Fprintf(&content, "Ez az év %s. hete és %s. napja, %s nap van hátra az évből.\n",
		f2(fmt.Sprint(iweek)),
		f2(fmt.Sprint(currentTime.YearDay())),
		f2(fmt.Sprint(lastDayOfYear-currentTime.YearDay())),
	)
	fmt.Fprintf(&content, "Aktuális idő: %s óra %s perc és %s másodperc.\n",
		f2(fmt.Sprintf("%02d", currentTime.Hour())),
		f2(fmt.Sprintf("%02d", currentTime.Minute())),
		f2(fmt.Sprintf("%02d", currentTime.Second())),
	)
	fmt.Fprintf(&content, "Boldog névnapot kedves %s nevű felhasználóinknak!\n", f3(nevnapok[nevnapKey]))
	fmt.Fprintf(&content, "Elkövetkezendő névnapok: %s; %s; %s\n", f4(nevnapok[nevnapKey2]), f4(nevnapok[nevnapKey3]), f4(nevnapok[nevnapKey4]))

	return content.String(), nil
}
