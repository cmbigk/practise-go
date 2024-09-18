import os
import re

# Open the input file for reading
with open("haha.txt", "r") as input_file:
    content = input_file.readlines()

# Convert all the text to lowercase
content = [line.lower() for line in content]

# Capitalize one word (case-insensitive) (for example, the word "doit")
word_to_capitalize = "doit"
capitalized_word = word_to_capitalize.capitalize()
content = [re.sub(word_to_capitalize, capitalized_word, line, flags=re.IGNORECASE) for line in content]

# Create the output file for writing
with open("output.txt", "w") as output_file:
    output_file.writelines(content)

