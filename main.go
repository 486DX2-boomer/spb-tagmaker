package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jung-kurt/gofpdf"
)

// GLOBALS
// l contains all price tag objects in memory
var l List

func main() {

	// Step 1 : Set up a web server
	// Step 2 : Display all tags in memory
	// Step 3 : User clicks "add tag"
	// Step 4 : Collect information from form, add the tag to memory
	// Repeat from step 2 until PRINT TAG is clicked by user
	// Step 5 : Get all tags in memory, sort by handgun, long gun, new and used
	// Step 6 : Generate a pdf with all handgun and long gun tags (handgun and long gun tags are always separate pages)
	// If manufacturer logo isn't found in the /logos folder, query Google Images for a logo. Save the logo on the server for later
	// The web server runs indefinitely, waiting for new tags or for the user to delete all tags in memory and start over.

	// -----------------TESTING STUFF
	// l = append(l, NewTag("Walther", "PPK", "380 Auto", false, "779.99", Small))
	// l = append(l, NewTag("Colt", "M4 Carbine", "5.56mm", false, "1099.99", Big))
	// l = append(l, NewTag("Springfield", "Saint", "5.56mm", true, "879", Big))

	// l = append(l, NewTag("Glock", "G44", "22 LR", false, "389", Small))
	// l = append(l, NewTag("Smith & Wesson", "617", "22 LR", true, "709.99", Small))
	// l = append(l, NewTag("Kel Tec", "PMR 30", "22 WMR", true, "399", Small))

	// l = append(l, NewTag("Ruger", "Wrangler", "22 LR", true, "199.99", Small))
	// l = append(l, NewTag("Taurus", "Judge Tracker", "45 LC/410", true, "469", Small))
	// l = append(l, NewTag("Charter Arms", "Lavender Lady", "38 Special", true, "414", Small))
	// l = append(l, NewTag("Rock Island", "GI Standard CS", "45 ACP", true, "419", Small))
	// l = append(l, NewTag("FN", "FNX-45 Tactical", "45 ACP", true, "1229.99", Small))
	// l = append(l, NewTag("CZ", "97B", "45 ACP", true, "719.99", Small))
	// l = append(l, NewTag("Smith & Wesson", "M&P 40 FDE", "40 S&W", true, "699", Small))
	// l = append(l, NewTag("GSG", "GSG-16 Carbine", "22 LR", false, "349.99", Big))
	// l = append(l, NewTag("HK", "HK 45c", "45 ACP", true, "779.99", Small))

	// BuildDocument(l, NewDocument())
	// -----------------TESTING STUFF

	http.HandleFunc("/", listTags) // setting router rule
	http.HandleFunc("/addtagform", addTagForm)
	http.HandleFunc("/addtag", addTag)
	http.HandleFunc("/deletealltags", deleteAllTags)

	err := http.ListenAndServe(ListenPort, nil) // setting listening port
	if err != nil {
		log.Fatal(err)
	}

}

// BuildDocument accepts a list of tags and builds a full document, sorting the handgun and long gun tags and then drawing each tag to an appropriate page.
// BuildDocument must be supplied with a pointer to a gofpdf object which can be supplied with the NewDocument() function
func BuildDocument(inputList []Tag, pdf *gofpdf.Fpdf) {

	longGunPageCoord := GetLongGunPageCoord()
	handGunPageCoord := GetHandgunPageCoord()

	var handgunTags List
	var longGunTags List

	for i := 0; i < len(inputList); i++ {
		if inputList[i].TagSize == Big {
			longGunTags = append(longGunTags, inputList[i])
		} else if inputList[i].TagSize == Small {
			handgunTags = append(handgunTags, inputList[i])
		}
	}

	for i := 0; i < len(longGunTags); i++ {
		if i%3 == 0 {
			pdf.AddPage()
		}
		longGunTags[i].Draw(longGunPageCoord[(i%3)], pdf)
	}

	for i := 0; i < len(handgunTags); i++ {
		if i%10 == 0 {
			pdf.AddPage()
		}
		handgunTags[i].Draw(handGunPageCoord[(i%10)], pdf)
	}

	err := pdf.OutputFileAndClose("output.pdf")
	fmt.Println(err)
	fmt.Println("PDF generated")

}

// NewDocument initializes a pdf, sets the font and font size, and then returns the pdf. This is meant to be used in conjunction with BuildDocument()
func NewDocument() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "pt", "A4", "")
	pdf.AddUTF8Font(SPBTagFont, "", "framd.ttf")
	pdf.SetFont(SPBTagFont, "", 12) // Call set font with a default size to avoid compiler error
	pdf.SetTextColor(0, 0, 0)

	return pdf
}

func listTags(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>SPB Tag Maker</h1>")

	// UI Controls
	fmt.Fprintf(w, "<b><a href=/addtagform>(Add Tag)</b></a>        ")
	fmt.Fprintf(w, "<b><a href=/deletealltags> (Delete All Tags)</a></b>        ")
	fmt.Fprintf(w, "<b>(Upload Manufacturer Logo)</b>        ")
	fmt.Fprintf(w, "<b>(Generate PDF)</b>        ")

	fmt.Fprintf(w, "<p>")

	if len(l) == 0 {
		fmt.Fprintf(w, "<i>There are no tags in memory yet. Add one with the Add Tag button.</i>")
	}

	// List all tags in memory
	for i := range l {

		fmt.Fprintf(w, l[i].Manufacturer)
		fmt.Fprintf(w, "  |  ")

		fmt.Fprintf(w, l[i].Model)
		fmt.Fprintf(w, "  |  ")

		fmt.Fprintf(w, l[i].Price)
		fmt.Fprintf(w, "  |  ")

		fmt.Fprintf(w, l[i].Caliber)
		fmt.Fprintf(w, "  |  ")

		if l[i].New == true {
			fmt.Fprintf(w, "New")
			fmt.Fprintf(w, "  |  ")

		} else {
			fmt.Fprintf(w, "Used")
			fmt.Fprintf(w, "  |  ")

		}
		if l[i].TagSize == Big {
			fmt.Fprintf(w, "Big Tag")

		} else {
			fmt.Fprintf(w, "Small Tag")

		}
		fmt.Fprintf(w, "<b>        (edit)</b>")
		fmt.Fprintf(w, "<b>         (delete)</b>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "</p>")
	}

}

func addTag(w http.ResponseWriter, r *http.Request) {
	// get form data and add it to list
	// map[caliber:[5.56x45] manufacturer:[Colt] model:[M4 Carbine] new:[New Gun] price:[1199.99] tagsize:[Big Tag]]
	r.ParseForm()
	fmt.Println(r.Form)

	var t Tag
	t.Manufacturer = r.Form["manufacturer"][0]
	t.Model = r.Form["model"][0]
	t.Caliber = r.Form["caliber"][0]
	t.Price = r.Form["price"][0]

	if r.Form["new"][0] == "New Gun" {
		t.New = true
	} else {
		t.New = false
	}

	if r.Form["tagsize"][0] == "Big Tag" {
		t.TagSize = Big
	} else {
		t.TagSize = Small
	}

	l = append(l, t)

	// redirect back to main menu
	http.Redirect(w, r, "/", 303)
}

func addTagForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Addtag triggered")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles(".\\html\\add_tag.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, t)

}

func editTag(w http.ResponseWriter, r *http.Request) {

}

func deleteTag(w http.ResponseWriter, r *http.Request) {

}

func deleteAllTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all tags")
	l = nil

	// Redirect back to the main menu
	http.Redirect(w, r, "/", 303)
}

func uploadManufacturerLogo(w http.ResponseWriter, r *http.Request) {

}

func generatePDF(w http.ResponseWriter, r *http.Request) {

}
