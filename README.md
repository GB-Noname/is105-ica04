# is105-ica04

#### Oppgave 1
https://github.com/GB-Noname/is105-ica04/tree/master/lineshift 

b. lineshift.go og kode i /filetester

#### kode-kommentarer

lineshift.go

programmet importerer "os" "files" og "filestester" omsetter tekst via filestobyteslices. Programmet tar et filnavn som inndata

lineshifttester.go

gir riktig resultat for om filen er laget for klassisk MAC, UNIX eller Windows. Den gir oss og et antall linjer i filen.
Vi får en konklusjon om hvilket kode for linjeskift fra hvilket OS en tekstfil bruker. 

textTester.go

importerer pakken Lineshift og starter func FileStatCounter går gjennom hele byteslicen med tekst og så har du en manuell counter som sjekker hver eneste karakter i slicen. 
Hvis counteren er lik 0 så skal det legges til karakteren som er lik eller som er inklinert hvis vi treffer på samme karakteren. 
Elsen henter fra slicen, setter et nr til integeren og plusser på og returnerer til det nye slicen karakteren med tilhørende.
 

#### Oppgave 2 

https://github.com/GB-Noname/is105-ica04/tree/master/fileinfo

#### kode-kommentarer

a. fileinfo.go

Func main funksjonen kjører hele koden. De ulike Variablene bestemmer hva slags informasjon som skal returneres når prosessen er ferdig   

Programmet må først buildes, så kjøres gjennom ./fileinfo "egendefinert fil" (grunnen til at vi ikke har -f er at vi ikke fikk det til å funke, og ikke fant noe dokumentasjon som fikk det til å funke)

#### Oppgave 3 

https://github.com/GB-Noname/is105-ica04/tree/master/files

b. kjøres i mainFiles.go, kode ligger i /files

#### kode-kommentarer

mainFiles.go

Tester om filene er identiske gjennom å hente inn de to textene og bruker deepequal for å sjekke om bytene er identiske. 
Deretter kommer det en melding om filenes tilstand gjennom printLn-valgene som foretaes.

buffReader.go

func IoUtil, ReadFile leser filen du kaller fra (filename) 
WriteFile skriver data til filen nevnt på (filename) Hvis filen ikke eksisterer, så oppretter WriteFile den ved bruk av perm; rettigheten
func Buffio leser filen. Peek(n:5) returnerer de første (n) bytes uten å lese videre på filen, neste (n) bytes kommer ved neste loop helt til break.

c. files_test.go

	func benchmarkFiles åpner filen som er gitt i filename i files.og sjekker tørrelsen og hvor lang tid den bruker på å gå gjennom alle elementene i slicen
	
	func benchmarkIoUtil leser filen som blir referert til i buffReader.go og går gjennom loopen, stopper også starter den igjen.

	func benchmarkFireader åpner input filen i readwrite.go (func Firereader) og skriver og leser en output fil som den oppretter.
defer func buf:= make lager en buffer for å holde på informasjonen som blir lest gjennom (n, err := fi.Read)
