package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// GLOBALS
// l contains all price tag objects in memory
var l List

func main() {

	fmt.Println("-- SPB Tagmaker --")

	http.HandleFunc("/", listTags) // setting router rule
	http.HandleFunc("/addtagform", addTagForm)
	http.HandleFunc("/addtag", addTag)
	http.HandleFunc("/deletealltags", deleteAllTags)
	http.HandleFunc("/generatepdf", generatePDF)
	http.HandleFunc("/deletetag/", deleteTag)
	http.HandleFunc("/edittag/", editTag)
	http.HandleFunc("/edittagsave/", editTagSave)
	http.HandleFunc("/uploadmanufacturerlogoform/", uploadManufacturerLogoForm)
	http.HandleFunc("/uploadmanufacturerlogo/", uploadManufacturerLogo)
	http.HandleFunc("/logodirectory/", logoDirectory)

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

// listTags is the main menu of the UI
func listTags(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>SPB Tag Maker</h1>")

	// UI Controls
	fmt.Fprintf(w, "<b><a href=/addtagform>(Add Tag)</b></a>        ")
	fmt.Fprintf(w, "<b><a href=/deletealltags> (Delete All Tags)</a></b>        ")
	fmt.Fprintf(w, "<b><a href=/uploadmanufacturerlogoform>(Upload Manufacturer Logo)<a></b>        ")
	fmt.Fprintf(w, "<b><a href=/generatepdf>(Generate PDF)</a></b>        ")
	fmt.Fprintf(w, "<b><a href=/logodirectory>(Logo Directory)</a></b>        ")

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

		if !LogoFound(l[i].Manufacturer) {
			// If manufacturer logo not detected print warning
			fmt.Fprintf(w, `<span style="color:red;"> | Logo not found</span>`)
		}

		fmt.Fprintf(w, ((" <b><a href=/edittag/") + strconv.Itoa(i) + (">(edit)</a></b>")))

		fmt.Fprintf(w, ((" <b><a href=/deletetag/") + strconv.Itoa(i) + (">(delete)</a></b>")))

		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "</p>")
	}

}

func addTag(w http.ResponseWriter, r *http.Request) {

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
	t, err := template.ParseFiles("./html/add_tag.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, t)

}

func editTag(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Editing tag")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Get the index of the tag to be edited
	s := fmt.Sprint(r.URL)                        // Write the r.URL to a string
	editIndex := strings.Split(s, "/edittag/")[1] // Split it to get the index
	i, _ := strconv.Atoi(editIndex)               // convert to integer to access tag data

	var data struct {
		Manufacturer string
		Model        string
		Caliber      string
		Price        string
		New          string
		Used         string
		Big          string
		Small        string
		EditIndex    string
	}
	// grab the data from the Tag and convert it to a series of strings the templater can use
	data.Manufacturer = l[i].Manufacturer
	data.Model = l[i].Model
	data.Caliber = l[i].Caliber
	data.Price = l[i].Price

	// There are two radio button fields on the edit form, New or Used Gun and Big or Small Tag. One of each should be set to "checked" based on the tag data
	if l[i].New == true {
		data.New = "checked"
		data.Used = " "
	} else {
		data.New = " "
		data.Used = "checked"
	}

	if l[i].TagSize == Big {
		data.Big = "checked"
		data.Small = " "
	} else {
		data.Big = " "
		data.Small = "checked"
	}

	// The index of the edited tag is stored here and put in the template, so it can be posted to editTagSave() when the edit tag form is filled out
	data.EditIndex = editIndex

	fmt.Print(" ", i, "\n") // print to the console which tag is being edited
	fmt.Println(data)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles("./html/edit_tag.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)

}

func editTagSave(w http.ResponseWriter, r *http.Request) {

	s := fmt.Sprint(r.URL)                            // Write the r.URL to a string
	editIndex := strings.Split(s, "/edittagsave/")[1] // Split it to get the index
	i, _ := strconv.Atoi(editIndex)                   // convert to integer to access tag data

	// The rest of the logic is the same as adding a tag, but overwriting it in place instead of appending it
	r.ParseForm()
	fmt.Println(r.Form)

	l[i].Manufacturer = r.Form["manufacturer"][0]
	l[i].Model = r.Form["model"][0]
	l[i].Caliber = r.Form["caliber"][0]
	l[i].Price = r.Form["price"][0]

	if r.Form["new"][0] == "New Gun" {
		l[i].New = true
	} else {
		l[i].New = false
	}

	if r.Form["tagsize"][0] == "Big Tag" {
		l[i].TagSize = Big
	} else {
		l[i].TagSize = Small
	}

	// Redirect back to the main menu
	http.Redirect(w, r, "/", 303)

}

// removeTagFromList deletes an index from the price tag list, maintaining the order of the tags. It is called in deleteTag and supplised with an index from the appropriate URL from the listTags main menu
func removeTagFromList(list []Tag, index int) []Tag {
	ret := make([]Tag, 0)
	ret = append(ret, list[:index]...)
	return append(ret, list[index+1:]...)
}

func deleteTag(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Deleting tag")

	// Get the index of the tag to be removed from list
	s := fmt.Sprint(r.URL)                            // Write the r.URL to a string
	deleteIndex := strings.Split(s, "/deletetag/")[1] // Split it to get the index
	i, _ := strconv.Atoi(deleteIndex)                 // convert to integer to supply to RemoveTagFromList()

	fmt.Print(" ", i, "\n") // print to the console which tag is being deleted

	l = removeTagFromList(l, i)

	// Redirect back to the main menu
	http.Redirect(w, r, "/", 303)
}

func deleteAllTags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting all tags")
	l = nil

	// Redirect back to the main menu
	http.Redirect(w, r, "/", 303)
}

func uploadManufacturerLogo(w http.ResponseWriter, r *http.Request) {

	// 10 MB maximum upload size
	r.ParseMultipartForm(10 << 20)

	// get name of manufacturer
	var logoFilename string
	logoFilename = r.FormValue("Manufacturer Name")
	logoFilename = strings.ToLower(logoFilename)

	file, handler, err := r.FormFile("uploadedLogo")
	if err != nil {
		fmt.Println("Error retrieving logo upload: ")
		fmt.Println("err")
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded file: %+v \n", handler.Filename)
	fmt.Printf("Filesize: %+v \n", handler.Size)
	fmt.Printf("MIME Header: %+v \n", handler.Header)

	if (handler.Header["Content-Type"][0] != "image/jpeg") && (handler.Header["Content-Type"][0] != "image/gif") && (handler.Header["Content-Type"][0] != "image/png") {
		fmt.Println("Not an image file, aborting")
		fmt.Fprintf(w, "<h1>Error: not an image file, aborting upload<h1>")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<br><a href=/listtags/><button>Back</button></a>")
		return
	}

	// Get the file extension to construct the full filename
	extension := strings.Split(handler.Header["Content-Type"][0], "image/")[1]
	logoFilename += "."
	logoFilename += extension

	fmt.Println(logoFilename)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create(logoFilename)
	defer f.Close()

	f.Write(fileBytes)
	f.Close()
	// Move the output file to the logos directory
	err = os.Rename(logoFilename, ("./logos/" + logoFilename))

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "An error occurred: %s", err)
		fmt.Fprintf(w, "<br><a href=/listtags/><button>Back</button></a>")
		return
	}

	fmt.Println("Successfully uploaded", logoFilename)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "File upload successful")
	fmt.Fprintf(w, "<br><a href=/listtags/><button>Back</button></a>")

	// The redirect from this point does not work, I get "http: superfluous response.WriteHeader call from main.uploadManufacturerLogo (main.go:391)"
	// http.Redirect(w, r, "/", 303)
	// Instead I added a Back button

}

func uploadManufacturerLogoForm(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload form triggered triggered")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("./html/upload_manufacturer_logo.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, t)
}

func scrapeManufacturerLogo(w http.ResponseWriter, r *http.Request) {
	// Scrape a logo automatically from Google Images
}

func generatePDF(w http.ResponseWriter, r *http.Request) {
	BuildDocument(l, NewDocument())
	http.ServeFile(w, r, "./output.pdf")

}

// logoDirectory displays a list of all logo images stored in /logos
func logoDirectory(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Logo Directory</h1>")
	fmt.Fprintf(w, "<b><a href=/>(Back)</b></a>        ")

	fmt.Fprintf(w, "<p>")

	// step 1: read every file in /logos
	// step 2: add it to a slice of strings
	// step 3: print it to the dom
	// step 4: Rewrite this so it's not redundant because I could just use ReadDir to write files to the DOM without the intermediate step

	var logoList []string

	files, err := ioutil.ReadDir("./logos")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		logoList = append(logoList, file.Name())
	}

	if len(logoList) == 0 {
		fmt.Fprintf(w, "<i>There are no logos uploaded to the directory yet. Add one with the Upload Manufacturer Logo button.</i>")
	}

	// List all tags in memory
	for i := range logoList {

		fmt.Fprintf(w, logoList[i])

		fmt.Fprintf(w, "<br>")
		// fmt.Fprintf(w, "</p>")
	}

}
