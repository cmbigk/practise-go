const fs = require('fs');
const readline = require('readline');

// Open the input file for reading
const inputFile = fs.createReadStream('haha.txt');
const outputFile = fs.createWriteStream('output.txt');

// Create a readline interface for reading the input file line by line
const rl = readline.createInterface({
  input: inputFile,
  crlfDelay: Infinity
});

// Read each line from the input file and write it to the output file

// Handle any errors
rl.on('error', (err) => {
  console.error('Error reading input file:', err);
});

outputFile.on('error', (err) => {
  console.error('Error writing to output file:', err);
});
