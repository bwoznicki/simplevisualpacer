# Simple Visual Pacer (svp)

Terminal based visual pacer for speed reading training.

Most available speedreaders work on the basis of in place replacement of words in text, which does not follow natural patern of reading by following the text ( usualy from left to right ). Using this visual pacer should feel more natural while switching between computer screen and printed text or other medium.

## usage
```GO
$ svp -s 250 -f myBook.txt

$ svp --help
Usage of svp:
  -f string
    	Book file
  -lw int
    	Line width - max number of characters per line (default 80)
  -p int
    	Position in book - line number
  -s int
    	Reading speed in words per minute (default 200)
  -w int
    	Field width - max number of words displayed (default 2)
```

## Some default values:
- Text width is set to **80** characters by default.
- Speed **variable** ( words per minute ) default target is set to **200**. Avarage length of a word in english is 5.1 characters long and your actual speed may vary depending on the length of the words in the text. For more natural feel the speed also changes following comma, full stop and new line. Long words also slow down the pacer and add extra 50 millisends for each letter above 5 characters.
- Visible words are set to 2 with higher speeds you might want to go with 3 or even 4.
- bookmarking is also possible. When you interupt the pacer it will return the line number you finished on and you can resume reading from the same place next time.

## Reading books / articles
You can read any text based file. For best results you might want to consider converting some books from epub to txt.

## Embeded book
Included free book "A tale of two cities" by Charles Dickens was downloaded from www.gutenberg.org

##