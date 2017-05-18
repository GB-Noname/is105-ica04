# is105-ica04

### Oppgave 1
https://github.com/GB-Noname/is105-ica04/tree/master/lineshift 

b. lineshift.go og kode i /filetester

### Oppgave 2 

https://github.com/GB-Noname/is105-ica04/tree/master/fileinfo

a. fileinfo.go
(kode-kommentarer)

Func main funksjonen kjører hele koden. De ulike Variablene bestemmer hva slags informasjon som skal returneres når prosessen er ferdig   

Programmet må først buildes, så kjøres gjennom ./fileinfo "egendefinert fil" (grunnen til at vi ikke har -f er at vi ikke fikk det til å funke, og ikke fant noe dokumentasjon som fikk det til å funke)
### Oppgave 3 

https://github.com/GB-Noname/is105-ica04/tree/master/files

b. kjøres i mainFiles.go, kode ligger i /files
(kode-kommentarer)

buffReader.go
func IoUtil, ReadFile leser filen du kaller fra (filename) 
WriteFile skriver data til filen nevnt på (filename) Hvis filen ikke eksisterer, så oppretter WriteFile den ved bruk av perm; rettigheten

func Buffio leser filen og Peek returnerer de første (n) bytes uten å lese videre på filen, neste (n) bytes kommer ved neste loop helt til break.

c. files_test.go
(kode-kommentarer)

func benchmarkFiles åpner filen som er gitt i filename i files.go og sjekker tørrelsen og hvor lang tid den bruker på å gå gjennom alle elementene i slicen

func benchmarkIoUtil   


