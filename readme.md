# Shoot Point Blank - Tag Maker

**SPB internal use only**

## Efficient firearm price tag label maker

Feature list:
- Runs on a server as a web app (access from devices other than the desk computer)
- Display a list of all tags in queue
- Add and delete price tags
- Edit price tags
- Upload manufacturer logos
- Query Google Images for manufacturer logos automatically *(Not implemented yet)*
- Generate more than one page of price tags at a time
- If less than 10 2x3 / 3 3x5 tags on a page, leave the rest of the slots empty (saves on ink/toner when printing)
- Mix new and used gun tags on the same page
- Won't break like an Excel spreadsheet (Will probably break in other ways)

## How It Works

The tag maker is a web app intended to run on the store's local server (although it can alternately run on just the main desk machine, or on the internet). In the latter case, only one store can use it at a time, otherwise we'd end up overwriting each other's tags.

On opening the web app, there is a list of all tags currently stored in memory. It will look something like this: 

```
    SPB Tag Maker

    (Add Tag) (Delete All Tags) (Upload Manufacturer Logo) (Generate PDF)

    FN | 509c | 899.99 | 9x19mm | New | Small Tag (delete)
    Smith & Wesson | M&P 380 EZ | 449 | 380 ACP | New | Small Tag (delete)
    Colt | M4 Carbine | 1699 | 5.56x45mm | New | Big Tag (delete)
    CZ | 75B | 539 | 9x19mm | Used | Small Tag (delete)
    Daniel Defense | DDM4V7 | 1699 | 5.56x45 | New | Big Tag (delete)
```

From here, the user can manipulate the list of price tags with the simple web user interface.

Once all desired tags are added to the list, the user need only click Generate PDF to download a very pretty PDF document with all the tags in the correct position and formatting ready to be printed and put in the store.

Manufacturer logos are stored server-side. If the logo for a tag cannot be found, it will automatically query a logo from Google Images *(not implemented yet)*. Alternatively, the user can upload an image of the desired logo. Filenames must be all lowercase and the image must be either a jpg, png, or gif.

## The Tags

Handgun tags are 2x3. Long gun tags are 3x5. There are 10 handgun tags per 8.5x11 page and 3 long gun tags per 8.5x11 page. The font on the tags is Franklin Gothic Medium. The SPB logo is in the top left and the manufacturer logo is on the top right, with an orangish-red separator bar between the price and the model and caliber.The separator bar color is RGB(212, 87, 42). The separator bar is 2.5 pt in thickness.

Font sizes adjust dynamically based on the amount of characters in a field, so most common calibers, prices, and model names will fit on the tag. Only the longest entries will start to spill over (for example, *M&P Shield 380 EZ Performance Center* will be printed in a smaller font size than *1911A1*, and if the model name is long enough it can indeed spill over in to the caliber field.)

## Installation

The tag maker is written in Go and requires the latest version to be installed on the desired machine. Clone the github repo (I can also provide the source code in a zip file if you need it), navigate to the directory and type in `go run .` It is also completely fine to use `go build .` and then run the compiled binary.

The program runs by default on port 8080 (you can modify ListenPort in constants.go, there is no configuration file yet), so you can open your browser to localhost:8080 if running on your local machine. If running on a server, navigate to the correct IP at port 8080.
____

## Package References
https://godoc.org/github.com/jung-kurt/gofpdf

## What's Coming

- CSS styling to make the front end look pretty
- Scrape manufacturer logos from Google Images
- Show a list of uploaded manufacturer logos
- Save tags for later (Handy for commonly stocked guns)
