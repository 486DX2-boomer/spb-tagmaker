# Shoot Point Blank - Tag Maker

**SPB internal use only**

The goal for this is to improve time, efficiency, and ease of use when creating price tags, instead of using an Excel spreadsheet with macros/manually editing PDFs

Feature list:
- Run on a server (access from devices other than the desk computer)
- Query Google Images for manufacturer logos automatically
- Display a list of all tags in queue
- Generate more than one page of price tags at a time
- If less than 10 2x3 / 3 3x5 tags on a page, leave the rest of the slots empty (saves on ink/toner when printing)
- Mix new and used gun tags on the same page

## How It Works

As a web app running on the store's local server or on the internet (although in the latter case, multiple stores cannot use the tag maker at the same time, as they would overwrite each other's tags.)

On opening the web app, there is a list of all tags currently stored in memory. The user can clear them, or add more.

Enter in the data for each tag
- Make
- Model
- Caliber
- Price
- Tag Size (handgun tags 2x3, long gun tags 3x5)
- New or Used

So for example, the program might look something like this...

____

*SPB Tag Maker*

1. Ruger MK II 22 LR 499.99 small new **(edit) (delete)**
2. Colt Government 45 ACP 999.999 small new **(edit) (delete)**
3. Daniel Defense DDM4V7 5.56 1599.99 big new **(edit) (delete)**

[the make, model, caliber, price, tag size, and new or used would be color coded in the browser for easier reading]
   
**(GENERATE PDF)**

**(DELETE ALL)**

____

It will not matter how many tags are entered in, what order they are entered, or whether they are new or used. The program will sort them all when it generates the PDF.

The program pulls manufacturer logos from a local directory. If the logo cannot be found, it will automatically query a logo from Google Images. Then it will save it serverside for later, instead of resaving it every time, which will keep the tags consistent. Filenames must be all lowercase and a jpg, png, or gif.

## The Tags

Handgun tags are 2x3. Long gun tags are 3x5. There are 10 handgun tags per 8.5x11 page and 3 long gun tags per 8.5x11 page.

The font on the tags is Franklin Gothic Medium. 

For the handgun tags, the price is written in approximately 65 pt font with the last two digits (the 99 cents) superscripted and approximately 34 pt font. The model name is approximately 12 pt font and the caliber is approximately 19 pt font. The SPB logo is in the top left and the manufacturer logo is on the top right, with an orangish-red separator bar between the price and the model and caliber. The hex value of the color of the separator bar is #D4582A. The separator bar is 2 pt in thickness.

For the long gun tags, the price is written in approximately 113 pt font with the last two digits superscripted and approximately 60 pt font. The model name is approximately 20 pt font and the caliber is approximately 29 pt font. 

Each tag has a discrete position on the page represented as a set of x, y coordinates. The information on that tag is then drawn to the page based on an offset from that position. So for example, if the first tag on the page starts at (44, 25), then the dollar sign is offset +5, +25. Storing offsets instead of concrete positions on the page allows for greater flexibility.

Cell width and font size will have to be dynamic based by the number of digits in the price and number of characters in the model name and caliber. So "9mm" will have a smaller font size and cell width than "6.5 Creedmoor" and a $4200 gun will have a narrower cell for the price than a $99 used gun. The values for this will have to be hand tuned as I'm not sure how to accomplish that mathematically unless I can determine the exact width of each character based on the font size.

____

## Devlog

10/15/2020
Made quite a bit of progress on the long gun tags. I moved the function to a Tag method called Draw and hand tuned the font size offsets. Now a 3 digit price tag
will fit perfectly, as well as 4. A 5 digit price tag will also work OK, assuming we ever sell a 10,000+ dollar gun. Additionally, the program now checks for the manufacturer logo and will work correctly for a png, gif, or jpg, assuming that the filename is completely lowercase. It will break on uppercase or mixed case filenames.

The final version of the program should scrape Google Images to automatically find the logo and save it to the server.

I still have to make the handgun tag logic, which will also be included in Tag.Draw(). There are still values such as text cell widths that need to be moved to constants.

After that I will need a Build function that will run through all the tags stored in List and called Tag.Draw() on each of them, dynamically adding more pages as needed.

Once all of that is working, it only needs to be hooked up to a web server. The frontend of the web app will be a simple HTML form that adds the tag from the data entered in the form. That will be the most time consuming and difficult part, as I have no idea how to do web apps in Go and the tutorials available aren't much help.

____

## Package References
https://godoc.org/github.com/jung-kurt/gofpdf