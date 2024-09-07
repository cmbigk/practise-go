import os

# Open the input file for reading
with open("haha.txt", "r") as input_file:
    content = input_file.readlines()

# The word you want to capitalize
word_to_capitalize = 'In'
target_index = 0  # For example, capitalize the second occurrence (index 1)

# Initialize a counter for occurrences across all lines
occurrence_count = 0

# Modify the specific occurrence of the word
modified_content = []
for line in content:
    words = line.split()
    modified_words = []
    for word in words:
        # Check if the word matches the word to capitalize (case-insensitive)
        if word.lower() == word_to_capitalize.lower():
            if occurrence_count == target_index:
                modified_words.append(word_to_capitalize)  # Add the original word
                occurrence_count += 1  # Increment occurrence count after modifying
            else:
                occurrence_count += 1  # Increment occurrence count for each match
                modified_words.append(word.lower())  # Add the lowercase version
        else:
            modified_words.append(word.lower())  # Add the lowercase version for other words
    modified_content.append(' '.join(modified_words))

# Write the modified content to the output file
with open("output.txt", "w") as output_file:
    output_file.write('\n'.join(modified_content))