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

The program pulls manufacturer logos from a local directory. If the logo cannot be found, it will automatically query a logo from Google Images. Then it will save it serverside for later, instead of resaving it every time, which will keep the tags consistent.

## The Tags

Handgun tags are 2x3. Long gun tags are 3x5. There are 10 handgun tags per 8.5x11 page and 3 long gun tags per 8.5x11 page.

The font on the tags is Franklin Gothic Medium. 

For the handgun tags, the price is written in approximately 65 pt font with the last two digits (the 99 cents) superscripted and approximately 34 pt font. The model name is approximately 12 pt font and the caliber is approximately 19 pt font. The SPB logo is in the top left and the manufacturer logo is on the top right, with an orangish-red separator bar between the price and the model and caliber. The hex value of the color of the separator bar is #D4582A. The separator bar is 2 pt in thickness.