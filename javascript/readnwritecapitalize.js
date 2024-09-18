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

    // Convert all the text to lowercase
    const lowercaseContent = lines.map(line => line.toLowerCase());

    // Capitalize the word 'doit' (case-insensitive)
    const wordToCapitalize = 'doit';
    const modifiedContent = lowercaseContent.map(line =>
      line.replace(new RegExp(wordToCapitalize, 'gi'), wordToCapitalize.charAt(0).toUpperCase() + wordToCapitalize.slice(1))
    );

    // Write the modified content to the output file
    await writeFile('output.txt', modifiedContent.join('\n'));
    console.log('File processing completed successfully.');
  } catch (err) {
    console.error('Error processing file:', err);
  }
}

// Call the processFile function
processFile();
