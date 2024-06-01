# Importing Libraries
from gtts import gTTS
import PyPDF2

# Open file Path
pdf_File = open('simple.pdf', 'rb')

# Create PDF Reader Object
pdf_Reader = PyPDF2.PdfReader(pdf_File)
count = len(pdf_Reader.pages) # counts number of pages in pdf
textList = []

# Extracting text data from each page of the pdf file
for i in range(count):
   try:
    page = pdf_Reader.pages[i]
    textList.append(page.extract_text())
   except:
       pass

# Converting multiline text to single line text
textString = " ".join(textList)

# Set language to English (en)
language = 'en'

# Call GTTS
if textString.strip():
    myAudio = gTTS(text=textString, lang=language, slow=False)
    # Save as mp3 file
    myAudio.save("Audio.mp3")
else:
    print("No extractable text found in the PDF.")
