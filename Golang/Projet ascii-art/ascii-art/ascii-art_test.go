package ascii

import (
	"fmt"
	"testing"
)

// Fonction de test 1
func TestArt_hello(t *testing.T) {
	input := []string{"", "hello"}
	Arg = input
	want := []string{
		` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`}
	// Fonction dans ascii-art (supprime les faux saut de ligne et gére les retour a la ligne)
	art := AsciiArt(Arg[1])
	// Génére l'ascii art dans ascii-art
	test := TestAsciiArt(art)

	// Si test mauvais, affiche les 2 résultats à l'utilisateur
	if test[0] != want[0] {
		fmt.Print(`AsciiArt("hello") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		// import testing
		t.FailNow()
	}
}

// Fonction de test 2
func TestArt_Hello(t *testing.T) {
	input := []string{"", "HELLO"}
	Arg = input
	want := []string{
		` _    _   ______   _        _         ____   
| |  | | |  ____| | |      | |       / __ \  
| |__| | | |__    | |      | |      | |  | | 
|  __  | |  __|   | |      | |      | |  | | 
| |  | | | |____  | |____  | |____  | |__| | 
|_|  |_| |______| |______| |______|  \____/  
                                             
                                             
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("HELLO") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 3
func TestArt_HeLo_HuMaN(t *testing.T) {
	input := []string{"", "HeLlo HuMaN"}
	Arg = input
	want := []string{
		` _    _          _        _                 _    _           __  __           _   _  
| |  | |        | |      | |               | |  | |         |  \/  |         | \ | | 
| |__| |   ___  | |      | |   ___         | |__| |  _   _  | \  / |   __ _  |  \| | 
|  __  |  / _ \ | |      | |  / _ \        |  __  | | | | | | |\/| |  / _` + "`" + ` | | . ` + "`" + ` | 
| |  | | |  __/ | |____  | | | (_) |       | |  | | | |_| | | |  | | | (_| | | |\  | 
|_|  |_|  \___| |______| |_|  \___/        |_|  |_|  \__,_| |_|  |_|  \__,_| |_| \_| 
                                                                                     
                                                                                     
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("HeLlo HuMaN") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 4
func TestArt_1Hello_2There(t *testing.T) {
	input := []string{"", "1Hello 2There"}
	Arg = input
	want := []string{
		`     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("1Hello 2There") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 5
func TestArt_HelloLineBreak(t *testing.T) {
	input := []string{"", "Hello\nThere"}
	Arg = input
	want := []string{
		` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("Hello\nThere") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 6
func TestArt_HelloDoubleLineBreak(t *testing.T) {
	input := []string{"", "Hello\n\nThere"}
	Arg = input
	want := []string{
		` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("Hello\n\nThere") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 7
func TestArt_Parenthesis(t *testing.T) {
	input := []string{"", "{Hello & There #}"}
	Arg = input
	want := []string{
		`   __  _    _          _   _                                _______   _                                    _  _    __    
  / / | |  | |        | | | |                 ___          |__   __| | |                                 _| || |_  \ \   
 | |  | |__| |   ___  | | | |   ___          ( _ )            | |    | |__     ___   _ __    ___        |_  __  _|  | |  
/ /   |  __  |  / _ \ | | | |  / _ \         / _ \/\          | |    |  _ \   / _ \ | '__|  / _ \        _| || |_    \ \ 
\ \   | |  | | |  __/ | | | | | (_) |       | (_>  <          | |    | | | | |  __/ | |    |  __/       |_  __  _|   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/         \___/\/          |_|    |_| |_|  \___| |_|     \___|         |_||_|    | |  
  \_\                                                                                                              /_/   
                                                                                                                         
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("Hello\nThere") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 8
func TestArt_SingleQuotes(t *testing.T) {
	input := []string{"", "hello There 1 to 2!"}
	Arg = input
	want := []string{
		` _              _   _                 _______   _                                            _                           _  
| |            | | | |               |__   __| | |                                 _        | |                  ____   | | 
| |__     ___  | | | |   ___            | |    | |__     ___   _ __    ___        / |       | |_    ___         |___ \  | | 
|  _ \   / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \       | |       | __|  / _ \          __) | | | 
| | | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/       | |       \ |_  | (_) |        / __/  |_| 
|_| |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|       |_|        \__|  \___/        |_____| (_) 
                                                                                                                            
                                                                                                                            
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("hello There 1 to 2!") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 9
func TestArt_RandomText1(t *testing.T) {
	input := []string{"", "MaD3IrA&LiSboN"}
	Arg = input
	want := []string{
		` __  __           _____            _____                              _        _    _____   _               _   _  
|  \/  |         |  __ \   _____  |_   _|            /\       ___    | |      (_)  / ____| | |             | \ | | 
| \  / |   __ _  | |  | | |___ /    | |    _ __     /  \     ( _ )   | |       _  | (___   | |__     ___   |  \| | 
| |\/| |  / _` + "`" + ` | | |  | |   |_ \    | |   | '__|   / /\ \    / _ \/\ | |      | |  \___ \  | '_ \   / _ \  | . ` + "`" + ` | 
| |  | | | (_| | | |__| |  ___) |  _| |_  | |     / ____ \  | (_>  < | |____  | |  ____) | | |_) | | (_) | | |\  | 
|_|  |_|  \__,_| |_____/  |____/  |_____| |_|    /_/    \_\  \___/\/ |______| |_| |_____/  |_.__/   \___/  |_| \_| 
                                                                                                                   
                                                                                                                   
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("MaD3IrA&LiSboN") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 10
func TestArt_RandomText2(t *testing.T) {
	input := []string{"", "1a\"#FdwHywR&/()="}
	Arg = input
	want := []string{
		`             _ _     _  _     ______       _              _    _                      _____                 __   __ __            
 _          ( | )  _| || |_  |  ____|     | |            | |  | |                    |  __ \    ___        / /  / / \ \   ______  
/ |   __ _   V V  |_  __  _| | |__      __| | __      __ | |__| |  _   _  __      __ | |__) |  ( _ )      / /  | |   | | |______| 
| |  / _` + "`" + ` |        _| || |_  |  __|    / _` + "`" + ` | \ \ /\ / / |  __  | | | | | \ \ /\ / / |  _  /   / _ \/\   / /   | |   | |  ______  
| | | (_| |       |_  __  _| | |      | (_| |  \ V  V /  | |  | | | |_| |  \ V  V /  | | \ \  | (_>  <  / /    | |   | | |______| 
|_|  \__,_|         |_||_|   |_|       \__,_|   \_/\_/   |_|  |_|  \__, |   \_/\_/   |_|  \_\  \___/\/ /_/     | |   | |          
                                                                   __/ /                                        \_\ /_/           
                                                                  |___/                                                           
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("1a\"#FdwHywR&/()=") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 11
func TestArt_ParenthesisWithRandom(t *testing.T) {
	input := []string{"", "{|}~"}
	Arg = input
	want := []string{
		`   __  _  __     /\/| 
  / / | | \ \   |/\/  
 | |  | |  | |        
/ /   | |   \ \       
\ \   | |   / /       
 | |  | |  | |        
  \_\ | | /_/         
      |_|             
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("{|}~") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 12
func TestArt_BracketsWithRandom(t *testing.T) {
	input := []string{"", `[\]^_ 'a`}
	Arg = input
	want := []string{
		` ___  __       ___   /\                  _          
|  _| \ \     |_  | |/\|                ( )         
| |    \ \      | |                     |/    __ _  
| |     \ \     | |                          / _` + "`" + ` | 
| |      \ \    | |                         | (_| | 
| |_      \_\  _| |                          \__,_| 
|___|         |___|       ______                    
                         |______|                   
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("[\]^_ 'a") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 13
func TestArt_RGB(t *testing.T) {
	input := []string{"", "RGB"}
	Arg = input
	want := []string{
		` _____     _____   ____   
|  __ \   / ____| |  _ \  
| |__) | | |  __  | |_) | 
|  _  /  | | |_ | |  _ <  
| | \ \  | |__| | | |_) | 
|_|  \_\  \_____| |____/  
                          
                          
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("RGB") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 14
func TestArt_Punctuation(t *testing.T) {
	input := []string{"", ":;<=>?@"}
	Arg = input
	want := []string{
		`           __          __     ___             
 _   _    / /  ______  \ \   |__ \     ____   
(_) (_)  / /  |______|  \ \     ) |   / __ \  
        < <    ______    > >   / /   / / _` + "`" + ` | 
 _   _   \ \  |______|  / /   |_|   | | (_| | 
(_) ( )   \_\          /_/    (_)    \ \__,_| 
    |/                                \____/  
                                              
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt(":;<=>?@") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 15
func TestArt_AlphabetUpperCase(t *testing.T) {
	input := []string{"", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	Arg = input
	want := []string{
		`            ____     _____   _____    ______   ______    _____   _    _   _____        _   _  __  _        __  __   _   _    ____    _____     ____    _____     _____   _______   _    _  __      __ __          __ __   __ __     __  ______ 
    /\     |  _ \   / ____| |  __ \  |  ____| |  ____|  / ____| | |  | | |_   _|      | | | |/ / | |      |  \/  | | \ | |  / __ \  |  __ \   / __ \  |  __ \   / ____| |__   __| | |  | | \ \    / / \ \        / / \ \ / / \ \   / / |___  / 
   /  \    | |_) | | |      | |  | | | |__    | |__    | |  __  | |__| |   | |        | | | ' /  | |      | \  / | |  \| | | |  | | | |__) | | |  | | | |__) | | (___      | |    | |  | |  \ \  / /   \ \  /\  / /   \ V /   \ \_/ /     / /  
  / /\ \   |  _ <  | |      | |  | | |  __|   |  __|   | | |_ | |  __  |   | |    _   | | |  <   | |      | |\/| | | . ` + "`" + ` | | |  | | |  ___/  | |  | | |  _  /   \___ \     | |    | |  | |   \ \/ /     \ \/  \/ /     > <     \   /     / /   
 / ____ \  | |_) | | |____  | |__| | | |____  | |      | |__| | | |  | |  _| |_  | |__| | | . \  | |____  | |  | | | |\  | | |__| | | |      | |__| | | | \ \   ____) |    | |    | |__| |    \  /       \  /\  /     / . \     | |     / /__  
/_/    \_\ |____/   \_____| |_____/  |______| |_|       \_____| |_|  |_| |_____|  \____/  |_|\_\ |______| |_|  |_| |_| \_|  \____/  |_|       \___\_\ |_|  \_\ |_____/     |_|     \____/      \/         \/  \/     /_/ \_\    |_|    /_____| 
                                                                                                                                                                                                                                               
                                                                                                                                                                                                                                               
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("ABCDEFGHIJKLMNOPQRSTUVWXYZ") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}

// Fonction de test 16
func TestArt_AlphabetLowerCase(t *testing.T) {
	input := []string{"", "abcdefghijklmnopqrstuvwxyz"}
	Arg = input
	want := []string{
		`         _                  _           __           _       _     _          _                                                            _                                                    
        | |                | |         / _|         | |     (_)   (_)  _     | |                                                          | |                                                   
  __ _  | |__     ___    __| |   ___  | |_    __ _  | |__    _     _  | | _  | |  _ __ ___    _ __     ___    _ __     __ _   _ __   ___  | |_   _   _  __   __ __      __ __  __  _   _   ____ 
 / _` + "`" + ` | | '_ \   / __|  / _` + "`" + ` |  / _ \ |  _|  / _` + "`" + ` | |  _ \  | |   | | | |/ / | | | '_ ` + "`" + ` _ \  | '_ \   / _ \  | '_ \   / _` + "`" + ` | | '__| / __| | __| | | | | \ \ / / \ \ /\ / / \ \/ / | | | | |_  / 
| (_| | | |_) | | (__  | (_| | |  __/ | |   | (_| | | | | | | |   | | |   <  | | | | | | | | | | | | | (_) | | |_) | | (_| | | |    \__ \ \ |_  | |_| |  \ V /   \ V  V /   >  <  | |_| |  / /  
 \__,_| |_.__/   \___|  \__,_|  \___| |_|    \__, | |_| |_| |_|   | | |_|\_\ |_| |_| |_| |_| |_| |_|  \___/  | .__/   \__, | |_|    |___/  \__|  \__,_|   \_/     \_/\_/   /_/\_\  \__, | /___| 
                                              __/ |              _/ |                                        | |         | |                                                       __/ /        
                                             |___/              |__/                                         |_|         |_|                                                      |___/         
`}

	art := AsciiArt(Arg[1])
	test := TestAsciiArt(art)

	if test[0] != want[0] {
		fmt.Print(`AsciiArt("abcdefghijklmnopqrstuvwxyz") = `)
		fmt.Println()
		fmt.Println(test[0])
		fmt.Print(`     want match for : `)
		fmt.Println()
		fmt.Println(want[0])
		t.FailNow()
	}
}
