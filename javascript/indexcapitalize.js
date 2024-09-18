const fs = require('fs');
const util = require('util');

// Promisify the readFile and writeFile functions
const readFile = util.promisify(fs.readFile);
const writeFile = util.promisify(fs.writeFile);

async function processFile() {
  try {
    // Read the entire content of the input file
    const content = await readFile('haha.txt', 'utf8'); // Read the file as a string

    // Split the content into lines
    const lines = content.split('\n');

    // The word you want to capitalize
    const wordToCapitalize = 'In';
    const targetIndex = 0; // For example, capitalize the second occurrence (index 1)

    // Initialize a counter for occurrences across all lines
    let occurrenceCount = 0;

    // Modify the specific occurrence of the word
    const modifiedLines = lines.map(line => {
      // Convert the entire line to lowercase
      const lowercaseLine = line.toLowerCase();

      // Split the lowercase line into words
      const words = lowercaseLine.split(' ');

      const modifiedWords = words.map(word => {
        if (word === wordToCapitalize.toLowerCase()) {
          if (occurrenceCount === targetIndex) {
            occurrenceCount++; // Increment occurrence count after modifying
            return wordToCapitalize; // Return the original capitalized word
          }
          occurrenceCount++; // Increment occurrence count for each match
        }
        return word;
      });

      // Join the modified words back into a line
      return modifiedWords.join(' ');
    });

    // Write the modified content to the output file
    await writeFile('output.txt', modifiedLines.join('\n'));
    console.log('File processing completed successfully.');
  } catch (err) {
    console.error('Error processing file:', err);
  }
}

// Call the processFile function
processFile();