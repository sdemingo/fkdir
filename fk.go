package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"os"
	"flag"
	"unicode"
	"bufio"
)



var words []string

func LoadWords() []string{
	file, err := os.Open("dict.txt")
	if err!=nil{
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	return text
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
	var punct_symbols =[]string{".",";",","}
	text := ""
	if min >= max {
		min = 0
	}
	nwords := rand.Intn(max-min) + min
	for i := 0; i < nwords; i++ {
		// Palabra aleatoria sacada del diccionario
		w := words[rand.Intn(len(words)-1)]
		if i == 0 {
			w = strings.Title(w)
		}
		// 90% probabilidad de que la palabra
		if rand.Intn(100)>90{
			w = RandomAlfaNumberString(5)
		}
		// Probabilidad de insertar signo de puntuación
		if (i % 5 == 0) && (rand.Intn(2)==0) && (text!=""){
			i:=rand.Intn(len(punct_symbols))
			text = text + punct_symbols[i] + " " + w
		}else{
			text = text + " " + w
		}
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




func buildRandomFile(filename string, npar int){
	text := RandomText(npar)
	lines:=SplitStringInLines(text,80)
	text=strings.Join(lines,"\n")
	err := os.WriteFile(filename, []byte(text), 0660)
	if err != nil {
		fmt.Println(err)
	}
	if *flagRandomDate{
		rdate:=RandomDate()
		err=os.Chtimes(filename,rdate,rdate)
		if err!=nil{
			fmt.Println(err)
		}
	}
}


var flagNumFiles = flag.Int("nf", 10, "Número exacto de ficheros que contendrá el directorio")
var flagNumPar = flag.Int("np",5,"Número máximo de párrafos del fichero")
var flagDeltaFiles = flag.Int("af", 0, "Variación de ficheros sobre el número máximo")
var flagSizeFileNames = flag.Int("sf", 10, "Longitud de los nombres de los ficheros")
var flagCharFileNames = flag.Bool("cf",false,"Incluir caracteres alfanuméricos en los nombres de ficheros")
var flagRandomDate = flag.Bool("rd",false,"Fecha de creación aleatoria para los ficheros")
var flagPrefix = flag.String("p","","Prefijo para los ficheros creados")
var flagMode = flag.String("o","dir","Producto objetivo (dir, table, file)")

func main() {
	rand.Seed(time.Now().UnixNano())
	flag.Parse()
	words=LoadWords()
	
	args:=flag.Args()
	if len(args)<1{
		fmt.Println("\n   Uso:   fk [flags] <output name>\n\n Lista de flags:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *flagMode == "dir"{
		dirname:=args[0]
		os.Mkdir(dirname,0755)

		nfiles:=*flagNumFiles
		for i:=0;i<nfiles;i++{
			filename:=""
			if *flagCharFileNames{
				filename=*flagPrefix+RandomAlfaNumberString(*flagSizeFileNames)
			}else{
				filename=*flagPrefix+RandomNumberString(*flagSizeFileNames)
			}
				
			filename=dirname+"/"+filename
			buildRandomFile(filename, *flagNumPar)
		}

	}else if *flagMode == "file" {
		filename:=args[0]
		buildRandomFile(filename, *flagNumPar)		
	}else{
		fmt.Println("Objetivo no reconocido. Tienes que usar -o con dir, table or file")		
	}
}
