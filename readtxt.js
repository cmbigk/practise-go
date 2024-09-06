const fs = require('fs');

fs.readFile('haha.txt', 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading file:', err);
    return;
  }

  console.log(data);
});
