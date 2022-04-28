package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"os"
	"flag"
	"unicode"
)

var words = []string{
	"ad",
	"adipisicing",
	"aliqua",
	"aliquip",
	"amet",
	"anim",
	"aute",
	"cillum",
	"commodo",
	"consectetur",
	"consequat",
	"culpa",
	"cupidatat",
	"deserunt",
	"do",
	"dolor",
	"dolore",
	"duis",
	"ea",
	"eiusmod",
	"elit",
	"enim",
	"esse",
	"est",
	"et",
	"eu",
	"ex",
	"excepteur",
	"exercitation",
	"fugiat",
	"id",
	"in",
	"incididunt",
	"ipsum",
	"irure",
	"labore",
	"laboris",
	"laborum",
	"Lorem",
	"magna",
	"minim",
	"mollit",
	"nisi",
	"non",
	"nostrud",
	"nulla",
	"occaecat",
	"officia",
	"pariatur",
	"proident",
	"qui",
	"quis",
	"reprehenderit",
	"sint",
	"sit",
	"sunt",
	"tempor",
	"ullamco",
	"ut",
	"velit",
	"veniam",
	"voluptate",
}

func RandomAlfaNumberString(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func RandomNumberString(len int) string {
	s := ""
	for i := 0; i < len; i++ {
		digit := rand.Intn(9)
		s = fmt.Sprintf("%s%d", s, digit)
	}
	return s
}

// Create a random paragraph with a n words where n is
// between min and max words
func RandomParagraph(min, max int) string {
	text := ""
	if min >= max {
		min = 0
	}
	nwords := rand.Intn(max-min) + min
	for i := 0; i < nwords; i++ {
		w := words[rand.Intn(len(words)-1)]
		if i == 0 {
			w = strings.Title(w)
		}
		text = text + " " + w
	}

	return strings.Trim(text, " ") + "."
}

// Create a random text with max paragraph. 
func RandomText(max int) string{
	n:=rand.Intn(max)+1
	t:=""
	for i:=0;i<n;i++{
		t=t+RandomParagraph(50,250)+"\n\n"
	}
	return t
}

// Create a random date between today and a year before
func RandomDate() time.Time {
	max:=time.Now().Unix()
	min:=time.Now().AddDate(-1,0,0).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}


// Trocea el texto en líneas de, como máximo nchars.
// Respeta el word wrapping
func SplitStringInLines(text string, nchars int) []string {
	lines := make([]string, 0)
	count := 0
	line := ""
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			lines = append(lines, line)
			count = 0
			line = ""
			continue
		} else {
			line += text[i : i+1]
			count++
			if count == nchars {
				nline := strings.TrimRightFunc(line, func(r rune) bool {
					return !unicode.IsSpace(r) && !unicode.IsPunct(r)
				})
				i -= (len(line) - len(nline)) // retraso i la diferencia entre line y nline (longitud del sufijo quitado)
				line = nline
				lines = append(lines, line)
				line = ""
				count = 0
			}
		}
	}
	if len(lines) > 0 {
		lines = append(lines, line)
	}
	return lines
}





var flagNumFiles = flag.Int("nf", 10, "Número exacto de ficheros que contendrá el directorio")
var flagDeltaFiles = flag.Int("af", 0, "Variación de ficheros sobre el número máximo")
var flagSizeFileNames = flag.Int("sf", 10, "Longitud de los nombres de los ficheros")
var flagRandomDate = flag.Bool("rd",false,"Fecha de creación aleatoria para los ficheros")
var flagPrefix = flag.String("p","","Prefijo para los ficheros creados")

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	args:=flag.Args()
	if len(args)<1{
		fmt.Println("\n   Uso:   fkdir [flags] <newdir>\n\n Lista de flags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	dirname:=args[0]
	os.Mkdir(dirname,0755)

	nfiles:=*flagNumFiles
	for i:=0;i<nfiles;i++{
		text := RandomText(5)
		lines:=SplitStringInLines(text,80)
		text=strings.Join(lines,"\n")
		filename:=*flagPrefix+RandomNumberString(*flagSizeFileNames)
		file:=dirname+"/"+filename
		err := os.WriteFile(file, []byte(text), 0660)
		if err != nil {
			fmt.Println(err)
		}
		if *flagRandomDate{
			rdate:=RandomDate()
			err=os.Chtimes(file,rdate,rdate)
			if err!=nil{
				fmt.Println(err)
			}
		}
	}
}
