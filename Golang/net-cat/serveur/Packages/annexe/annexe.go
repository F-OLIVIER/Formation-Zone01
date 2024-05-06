package Annexe

import (
	"log"
	"net"
)

// Définition des couleurs
var ColorUser = "\x1b[38;2;0;240;255m"
var ColorMessageBienvenue = "\x1b[38;2;0;255;4m"
var ColorHistorique = "\x1b[38;2;255;151;0m"
var ColorArriveUser = "\x1b[38;2;0;255;4m"
var ColorDepartUser = "\x1b[38;2;255;0;0m"
var yellow = "\x1b[38;2;255;255;0m"
var Fincolor = "\x1b[0m"

// Fonction de conversion : string en int
func Convint(s string) (num int, err string) {
	err = ""
	num = 0
	for _, j := range s {
		if j >= '0' && j <= '9' {
			num = num*10 + int(j-'0')
		} else {
			err = "err"
		}
	}
	return num, err
}

// Fonction de conversion : int en string
func Convstring(num int) (s string) {
	var arrayNum []rune
	for num > 0 {
		arrayNum = append(arrayNum, rune(num%10+'0'))
		num = num / 10
	}
	for a := len(arrayNum) - 1; a >= 0; a-- {
		s = s + string(arrayNum[a])
	}
	return s
}

// Fnction pour obtenir l'adresse IP du serveur
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	conn.Close()
	return localAddr.IP
}

// Fonction qui permet de print le pingouin !!!
func PenguinWelcome() (res string) {
	// res = ColorMessageBienvenue
	// res = res + "Welcome to TCP-Chat!" + "\n"
	// res = res + "         _nnnn_" + "\n"
	// res = res + "        dGGGGMMb" + "\n"
	// res = res + "       @p~qp~~qMb" + "\n"
	// res = res + "       M|@||@) M|" + "\n"
	// res = res + "       @,----.JM|" + "\n"
	// res = res + "      JS^\\__/  qKL" + "\n"
	// res = res + "     dZP        qKRb" + "\n"
	// res = res + "    dZP          qKKb" + "\n"
	// res = res + "   fZP            SMMb" + "\n"
	// res = res + "   HZM            MMMM" + "\n"
	// res = res + "   FqM            MMMM" + "\n"
	// res = res + " __| \".        |\\dS\"qML" + "\n"
	// res = res + " |    `.       | `' \\Zq" + "\n"
	// res = res + "_)      \\.___.,|     .'" + "\n"
	// res = res + "\\____   )MMMMMP|   .'" + "\n"
	// res = res + "     `-'       `--'" + "\n"
	// res = res + Fincolor

	res = res + "\n             Welcome to TCP-Chat!\n\n"
	res = res + "                 .88888888:.\n"
	res = res + "                88888888.88888.\n"
	res = res + "              .8888888888888888.\n"
	res = res + "              888888888888888888\n"
	res = res + "              88' _`88'_  `88888\n"
	res = res + "              88 88 88 88  88888\n"
	res = res + "              88_88_::_88_:88888\n"
	res = res + "              88" + yellow + ":::,::,:::::" + Fincolor + "8888\n"
	res = res + "              88" + yellow + "`:::::::::'`" + Fincolor + "8888\n"
	res = res + "             .88" + yellow + "  `::::'    " + Fincolor + "8:88.\n"
	res = res + "            8888            `8:888.\n"
	res = res + "          .8888'             `888888.\n"
	res = res + "         .8888:..  .::.  ...:'8888888:.\n"
	res = res + "        .8888.'     :'     `'::`88:88888\n"
	res = res + "       .8888        '         `.888:8888.\n"
	res = res + "      888:8         .           888:88888\n"
	res = res + "    .888:88        .:           888:88888:\n"
	res = res + "    8888888.       ::           88:888888\n"
	res = res + "    `" + yellow + ".::." + Fincolor + "888.      ::          .88888888\n"
	res = res + yellow + "   .::::::." + Fincolor + "888.    ::         :::`8888'" + yellow + ".:.\n"
	res = res + "  ::::::::::." + Fincolor + "888   '\"\"         " + yellow + ".::::::::::::\n"
	res = res + "  ::::::::::::." + Fincolor + "8    '      .:8" + yellow + "::::::::::::.\n"
	res = res + " .::::::::::::::." + Fincolor + "        .:888" + yellow + ":::::::::::::\n"
	res = res + " :::::::::::::::" + Fincolor + "88:.__..:88888" + yellow + ":::::::::::'\n"
	res = res + "  `'.:::::::::::" + Fincolor + "88888888888.88" + yellow + ":::::::::'\n"
	res = res + "        `':::_:'" + Fincolor + " -- '' -'-' `" + yellow + "':_::::'`\n" + Fincolor

	return res
}

// Fonction qui permet de print le message d'information de changement de pseudo
func MessageChangeName() (res string) {
	res = ColorMessageBienvenue
	res = res + "┌-------------------------------------------┐\n"
	res = res + "│------- To change your nickname, use ------│\n"
	res = res + "│----------- !name your nickname -----------│\n"
	res = res + "└-------------------------------------------┘\n"
	res = res + Fincolor
	return res
}

// Fonction qui print l'historique de discussion d'un groupe
func SendHistorique(conn net.Conn, data []byte) {
	msg := ColorHistorique
	msg = msg + "┌-------------------------------------------┐\n"
	msg = msg + "│----------► DISCUSSION HISTORY ◄-----------│\n"
	msg = msg + "│-------------------------------------------│\n"
	conn.Write([]byte(msg))
	conn.Write(data)
	msg = "└-------------------------------------------┘\n"
	msg = msg + Fincolor
	conn.Write([]byte(msg))
}

// Fonction qui check le pseudo en cas de changement de pseudo (!name)
func CheckName(msg string) (changement bool, newname string) {
	taille := len(msg)
	newname = ""
	if taille > 5 {
		debutMot := ""
		count := 0
		for _, a := range msg {
			if count < 5 {
				debutMot = debutMot + string(a)
			} else if count < 7 {
				if a != ' ' {
					newname = newname + string(a)
				}
			} else {
				if a != '\n' {
					newname = newname + string(a)
				}
			}
			count++
		}
		if debutMot == "!name" {
			return true, newname
		} else {
			return false, newname
		}
	} else {
		return false, newname
	}
}
